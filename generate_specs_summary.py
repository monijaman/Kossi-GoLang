import csv
import os
import re
from collections import defaultdict

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
    "Door Bins": 700,
    "Crisper Drawers": 701,
    "Frequency (Hz)": 704,
    "App Control": 705,
    "Compressor Warranty (Years)": 707,
}

def clean_model_name(model):
    """Convert model name to slug format for file naming"""
    # Remove special chars and convert to lowercase
    slug = re.sub(r'[^\w\s-]', '', model.lower())
    # Replace spaces with hyphens
    slug = re.sub(r'\s+', '-', slug)
    # Remove multiple hyphens
    slug = re.sub(r'-+', '-', slug)
    return slug.strip('-')

def get_specs_by_model(csv_file):
    """Read CSV and organize specs by model"""
    csv.field_size_limit(1000000)
    models_specs = defaultdict(dict)
    
    with open(csv_file, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            if row['found'] == 'yes' and row['spec_key'] and row['spec_value']:
                model = row['model']
                key = row['spec_key']
                value = row['spec_value']
                
                # Map the key to database ID if it exists
                db_key = None
                for mapped_key in KEY_MAPPING:
                    if key.strip() == mapped_key.strip():
                        db_key = mapped_key
                        break
                
                if db_key:
                    models_specs[model][db_key] = value
    
    return models_specs

def main():
    csv_file = 'data/marcel_specs.csv'
    models_specs = get_specs_by_model(csv_file)
    
    print(f"Found {len(models_specs)} models with specs")
    print("\nModels with specs:")
    for model in sorted(models_specs.keys()):
        count = len(models_specs[model])
        print(f"  {model}: {count} specs")
    
    # Create a summary file
    with open('models_specs_summary.txt', 'w', encoding='utf-8') as f:
        f.write("MODELS AND THEIR SPECIFICATIONS\n")
        f.write("=" * 80 + "\n\n")
        
        for model in sorted(models_specs.keys()):
            specs = models_specs[model]
            model_slug = clean_model_name(model)
            
            f.write(f"Model: {model}\n")
            f.write(f"Slug: {model_slug}\n")
            f.write(f"Spec Count: {len(specs)}\n")
            f.write(f"Database Keys:\n")
            
            for spec_key, spec_value in sorted(specs.items()):
                key_id = KEY_MAPPING.get(spec_key, "?")
                f.write(f"  {spec_key} ({key_id}): {spec_value[:60]}...\n" if len(spec_value) > 60 else f"  {spec_key} ({key_id}): {spec_value}\n")
            
            f.write("\n" + "-" * 80 + "\n\n")
    
    print(f"\n✓ Summary saved to models_specs_summary.txt")
    print(f"✓ {len(models_specs)} models ready for seeder generation")

if __name__ == '__main__':
    main()
