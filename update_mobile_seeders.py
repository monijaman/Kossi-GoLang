import os
import re
import glob

# Find all mobile seeder files
seeder_files = glob.glob("internal/infrastructure/database/seeders/specification_seeder_mobile_*.go")

print(f"Found {len(seeder_files)} mobile seeder files to update")

for filepath in seeder_files:
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Check if already using clause import
    if 'gorm.io/gorm/clause' not in content:
        # Add clause import
        content = content.replace(
            'import (\n\t"kossti/internal/infrastructure/database/models"\n\n\t"gorm.io/gorm"\n)',
            'import (\n\t"kossti/internal/infrastructure/database/models"\n\n\t"gorm.io/gorm"\n\t"gorm.io/gorm/clause"\n)'
        )
    
    # Find the Seed function and replace it
    # Pattern to match the entire Seed function
    seed_pattern = r'// Seed inserts specification records.*?^func \(s \*\w+\) Seed\(db \*gorm\.DB\) error \{.*?^\}'
    
    def create_new_seed_function(match):
        original = match.group(0)
        
        # Extract struct name and product slug
        struct_match = re.search(r'func \(s \*(\w+)\) Seed', original)
        if not struct_match:
            return original
        struct_name = struct_match.group(1)
        
        slug_match = re.search(r'productSlug := "([^"]+)"', original)
        if not slug_match:
            return original
        product_slug = slug_match.group(1)
        
        # Extract product name from comment
        comment_match = re.search(r"// Seed inserts specification records for the product identified by slug '([^']+)'", original)
        product_name = comment_match.group(1) if comment_match else product_slug
        
        # Create new function
        new_function = f'''// Seed inserts specification_translations for existing specifications for product '{product_slug}'
func (s *{struct_name}) Seed(db *gorm.DB) error {{
\tproductSlug := "{product_slug}"

\tvar prod models.ProductModel
\tif err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {{
\t\tif err == gorm.ErrRecordNotFound {{
\t\t\treturn nil
\t\t}}
\t\treturn err
\t}}

\tproductID := prod.ID
\tbanglaTranslations := s.getBanglaTranslations()

\t// Get all existing specifications for this product
\tvar existingSpecs []models.SpecificationModel
\tif err := db.Where("product_id = ?", productID).Find(&existingSpecs).Error; err != nil {{
\t\treturn err
\t}}

\t// Insert translations for all existing specifications
\tfor _, spec := range existingSpecs {{
\t\tbanglaValue, exists := banglaTranslations[spec.Value]
\t\tif exists && banglaValue != "" {{
\t\t\ttranslation := &models.SpecificationTranslationModel{{
\t\t\t\tSpecificationID: spec.ID,
\t\t\t\tLocale:          "bn",
\t\t\t\tValue:           banglaValue,
\t\t\t}}
\t\t\t// Use OnConflict to ignore if translation already exists
\t\t\tif err := db.Clauses(clause.OnConflict{{DoNothing: true}}).Create(translation).Error; err != nil {{
\t\t\t\treturn err
\t\t\t}}
\t\t}}
\t}}

\treturn nil
}}'''
        
        return new_function
    
    # Replace the Seed function
    new_content = re.sub(seed_pattern, create_new_seed_function, content, flags=re.MULTILINE | re.DOTALL)
    
    # Write back to file
    if new_content != content:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)
        print(f"✓ Updated: {os.path.basename(filepath)}")
    else:
        print(f"⚠ Skipped (no changes): {os.path.basename(filepath)}")

print("\n✅ All mobile seeder files have been updated!")
