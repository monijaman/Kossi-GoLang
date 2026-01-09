import csv
import os
import re
from collections import defaultdict
from datetime import datetime

# Specification key mapping from database
KEY_MAPPING = {
    "Gross Volume": 709,
    "Gross Volume (Outer Dimension, Manufacturer declared)": 709,
    "Net Volume": 710,
    "Weight": 80,
    "Net Weight": 80,
    "Gross Weight": 80,
    "Weight/Kg - Net/Packing": 80,
    "Weight/Kg - Net/Packing:": 80,
    "Rated Operating Voltage and Frequency": 160,
    "Rated Voltage/ Hz": 160,
    "Rated Voltage/ Hz/ watt": 160,
    "Operating voltage": 160,
    "Voltage": 160,
    "Compressor Type": 139,
    "Compressor": 139,
    "Temperature Control (Electronic/ Mechanical)": 158,
    "Temperature Control (Electronic/  Mechanical)": 158,
    "Temperature Control": 158,
    "Defrosting (Automatic/ Manual)": 141,
    "Defrosting (Automatic/ Manual):": 141,
    "Defrost Type": 141,
    "Reversible Door": 142,
    "Door Type": 142,
    "Refrigerant": 708,
    "Shelf (Material/ No.)": 699,
    "Shelf (Material/No.)": 699,
    "Rack Shelf (Material/No)": 699,
    "Shelf Material": 699,
    "Door Basket": 700,
    "Door Bins": 700,
    "Door Pocket": 700,
    "Vegetable Crisper": 701,
    "Crisper Drawers": 701,
    "Vegetable Crisper Cover": 701,
    "Water Dispenser": 703,
    "Ice Maker": 702,
    "Ice Case": 702,
    "Ice Tray": 702,
    "Energy Efficiency Rating": 143,
    "Energy Rating": 143,
    "Energy Star Rating": 144,
    "Star rating (BDS 1850:2012)": 144,
    "Dimensions": 25,
    "Width/mm": 25,
    "Depth/mm": 25,
    "Height/mm": 25,
    "Color": 311,
    "Brand": 310,
    "Model Name": 316,
    "Capacity": 138,
    "Number of Shelves": 154,
    "Warranty": 323,
    "Voice Assistant Support": 385,
    "Cooling Technology": 698,
    "Frequency (Hz)": 704,
    "App Control": 705,
    "Compressor Warranty (Years)": 707,
}

def clean_model_name(model):
    """Convert model name to slug format"""
    slug = re.sub(r'[^\w\s-]', '', model.lower())
    slug = re.sub(r'\s+', '-', slug)
    slug = re.sub(r'-+', '-', slug)
    return slug.strip('-')

def to_go_identifier(name):
    """Convert spec key name to Go identifier (PascalCase)"""
    # Remove special chars and split by space/dash
    parts = re.split(r'[\s\-/]+', name)
    # Capitalize first letter of each part and join
    return ''.join(word.capitalize() for word in parts if word)

def escape_go_string(s):
    """Escape special characters for Go string"""
    if not s:
        return s
    s = s.replace('\\', '\\\\')
    s = s.replace('"', '\\"')
    return s

def get_specs_by_model(csv_file):
    """Read CSV and organize specs by model"""
    csv.field_size_limit(1000000)
    models_specs = defaultdict(dict)
    
    with open(csv_file, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            if row['found'] == 'yes' and row['spec_key'] and row['spec_value']:
                model = row['model']
                key = row['spec_key'].strip()
                value = row['spec_value'].strip()
                
                # Find matching key in mapping
                db_key_id = None
                original_key = None
                for mapped_key, key_id in KEY_MAPPING.items():
                    if key == mapped_key:
                        db_key_id = key_id
                        original_key = mapped_key
                        break
                
                if db_key_id:
                    models_specs[model][original_key] = (value, db_key_id)
    
    return models_specs

def generate_seeder_file(model_name, specs, output_dir):
    """Generate a Go seeder file for a model"""
    model_slug = clean_model_name(model_name)
    filename = f"specification_seeder_refrigerator_marcel-{model_slug}.go"
    filepath = os.path.join(output_dir, filename)
    
    # Prepare specs mapping
    key_mapping = {}
    specs_dict = {}
    bangla_translations = {}
    
    for spec_key, (spec_value, key_id) in specs.items():
        key_mapping[spec_key] = key_id
        specs_dict[spec_key] = spec_value
        # Store original values for potential translations
        bangla_translations[spec_value] = spec_value  # Placeholder
    
    # Generate Go code
    struct_name = f"SpecificationSeederRefrigerator{to_go_identifier('Marcel-' + model_slug)}"
    receiver = struct_name[0].lower()
    
    content = f'''package seeders

import (
\t"kossti/internal/infrastructure/database/models"
\t"log"

\t"gorm.io/gorm"
)

// {struct_name} seeds specifications/options for product '{model_slug}'
type {struct_name} struct {{
\tBaseSeeder
}}

// New{struct_name} creates a new seeder instance
func New{struct_name}() *{struct_name} {{
\treturn &{struct_name}{{
\t\tBaseSeeder: BaseSeeder{{name: "Specifications for {model_slug}"}},
\t}}
}}

func ({receiver} *{struct_name}) getBanglaTranslations() map[string]string {{
\treturn map[string]string{{
'''
    
    # Add Bangla translations
    for value in sorted(set(specs_dict.values())):
        escaped_value = escape_go_string(value)
        content += f'\t\t"{escaped_value}": "{escaped_value}",\n'
    
    content += f'''\t\t// Add more translations as needed
\t}}
}}

// Seed inserts specification records for the product identified by slug '{model_slug}'
func ({receiver} *{struct_name}) Seed(db *gorm.DB) error {{
\tproductSlug := "{model_slug}"
\tvar prod models.ProductModel
\tif err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {{
\t\tif err == gorm.ErrRecordNotFound {{
\t\t\tlog.Printf("⚠️  Product not found: %s", productSlug)
\t\t\treturn nil
\t\t}}
\t\treturn err
\t}}
\tproductID := prod.ID

\texistingkeyMapping := map[string]uint{{
'''
    
    # Add key mapping
    for spec_key in sorted(key_mapping.keys()):
        key_id = key_mapping[spec_key]
        content += f'\t\t"{spec_key}": {key_id},\n'
    
    content += f'''\t}}

\tspecs := map[string]string{{
'''
    
    # Add specs
    for spec_key in sorted(specs_dict.keys()):
        spec_value = escape_go_string(specs_dict[spec_key])
        content += f'\t\t"{spec_key}": "{spec_value}",\n'
    
    content += f'''\t}}

\tbanglaTranslations := {receiver}.getBanglaTranslations()
\tfor key, value := range specs {{
\t\tspecKeyID, exists := existingkeyMapping[key]
\t\tif !exists {{
\t\t\tlog.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
\t\t\tcontinue
\t\t}}

\t\tvar existing models.SpecificationModel
\t\tif err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {{
\t\t\tif err == gorm.ErrRecordNotFound {{
\t\t\t\tsModel := &models.SpecificationModel{{
\t\t\t\t\tProductID:          productID,
\t\t\t\t\tSpecificationKeyID: specKeyID,
\t\t\t\t\tValue:              value,
\t\t\t\t\tStatus:             1,
\t\t\t\t}}
\t\t\t\tif err := db.Create(sModel).Error; err != nil {{
\t\t\t\t\treturn err
\t\t\t\t}}
\t\t\t\t// Ensure the ID is set after creation
\t\t\t\tif sModel.ID == 0 {{
\t\t\t\t\tif err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {{
\t\t\t\t\t\treturn err
\t\t\t\t\t}}
\t\t\t\t}}
\t\t\t\tbanglaValue, exists := banglaTranslations[value]
\t\t\t\tif exists && banglaValue != "" {{
\t\t\t\t\ttranslation := &models.SpecificationTranslationModel{{
\t\t\t\t\t\tSpecificationID: sModel.ID,
\t\t\t\t\t\tLocale:          "bn",
\t\t\t\t\t\tValue:           banglaValue,
\t\t\t\t\t}}
\t\t\t\t\tif err := db.Create(translation).Error; err != nil {{
\t\t\t\t\t\treturn err
\t\t\t\t\t}}
\t\t\t\t}}
\t\t\t}} else {{
\t\t\t\treturn err
\t\t\t}}
\t\t}} else {{
\t\t\tbanglaValue, exists := banglaTranslations[value]
\t\t\tif exists && banglaValue != "" {{
\t\t\t\tvar existingTranslation models.SpecificationTranslationModel
\t\t\t\tif err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {{
\t\t\t\t\tif err == gorm.ErrRecordNotFound {{
\t\t\t\t\t\ttranslation := &models.SpecificationTranslationModel{{
\t\t\t\t\t\t\tSpecificationID: existing.ID,
\t\t\t\t\t\t\tLocale:          "bn",
\t\t\t\t\t\t\tValue:           banglaValue,
\t\t\t\t\t\t}}
\t\t\t\t\t\tif err := db.Create(translation).Error; err != nil {{
\t\t\t\t\t\t\treturn err
\t\t\t\t\t\t}}
\t\t\t\t\t}} else {{
\t\t\t\t\t\treturn err
\t\t\t\t\t}}
\t\t\t\t}}
\t\t\t}}
\t\t}}
\t}}

\treturn nil
}}
'''
    
    # Write file
    with open(filepath, 'w', encoding='utf-8') as f:
        f.write(content)
    
    return filepath, len(specs)

def main():
    csv_file = 'data/marcel_specs.csv'
    output_dir = 'internal/infrastructure/database/seeders'
    
    # Ensure output directory exists
    os.makedirs(output_dir, exist_ok=True)
    
    models_specs = get_specs_by_model(csv_file)
    
    print(f"Generating Go seeder files for {len(models_specs)} models...")
    print("=" * 80)
    
    generated_files = []
    total_specs = 0
    
    for model in sorted(models_specs.keys()):
        specs = models_specs[model]
        filepath, spec_count = generate_seeder_file(model, specs, output_dir)
        model_slug = clean_model_name(model)
        filename = os.path.basename(filepath)
        
        generated_files.append((filename, model, spec_count))
        total_specs += spec_count
        
        print(f"✓ {filename:<70} ({spec_count:2d} specs)")
    
    print("=" * 80)
    print(f"\n✓ Generated {len(generated_files)} seeder files")
    print(f"✓ Total specifications: {total_specs}")
    print(f"✓ Output directory: {output_dir}")
    
    # Create an index file listing all seeders
    index_content = "// Auto-generated seeder index\n// Generated on " + datetime.now().isoformat() + "\n\n"
    index_content += "// Specification Seeders for Marcel Refrigerators\n"
    index_content += "// ============================================\n\n"
    index_content += "package seeders\n\n"
    index_content += "// RegisterAllMarcelRefrigeratorSeeders registers all Marcel refrigerator specification seeders\n"
    index_content += "func RegisterAllMarcelRefrigeratorSeeders() []Seeder {\n"
    index_content += "\treturn []Seeder{\n"
    
    for filename, model, spec_count in generated_files:
        struct_name = f"SpecificationSeederRefrigerator{to_go_identifier('Marcel-' + clean_model_name(model))}"
        index_content += f"\t\tNew{struct_name}(),\n"
    
    index_content += "\t}\n"
    index_content += "}\n"
    
    index_filepath = os.path.join(output_dir, 'specification_seeders_marcel_index.go')
    with open(index_filepath, 'w', encoding='utf-8') as f:
        f.write(index_content)
    
    print(f"\n✓ Index file created: {os.path.basename(index_filepath)}")

if __name__ == '__main__':
    main()
