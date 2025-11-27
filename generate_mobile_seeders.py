#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json
import os
import re

def slugify(name):
    """Convert device name to slug format"""
    # Convert to lowercase and replace spaces with hyphens
    slug = name.lower()
    # Replace spaces and special chars with hyphens
    slug = re.sub(r'[^a-z0-9]+', '-', slug)
    # Remove leading/trailing hyphens
    slug = slug.strip('-')
    return slug

def generate_go_file(device_name, specs, output_dir):
    """Generate a Go seeder file for a device"""
    
    slug = slugify(device_name)
    filename = f"specification_seeder_mobile_{slug}.go"
    filepath = os.path.join(output_dir, filename)
    
    # Create class name from device name
    class_name_parts = ''.join(word.capitalize() for word in device_name.replace('-', ' ').split())
    class_name = f"SpecificationSeederMobile{class_name_parts}"
    
    # Normalize class name (remove special chars)
    class_name = re.sub(r'[^a-zA-Z0-9]', '', class_name)
    
    # Extract specs and translations
    spec_map = {}
    bengali_translations = {}
    
    for spec_key, spec_data in specs.items():
        if isinstance(spec_data, dict) and "specification_translations" in spec_data:
            english_value = spec_data["english"]
            spec_map[spec_key] = english_value
            
            # Get Bengali translation
            for trans in spec_data["specification_translations"]:
                if trans["locale"] == "bn":
                    bengali_translations[english_value] = trans["value"]
    
    # Generate Go file content
    go_content = f'''package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// {class_name} seeds specifications for '{slug}'
type {class_name} struct {{
	BaseSeeder
}}

// New{class_name} creates a new seeder instance
func New{class_name}() *{class_name} {{
	return &{class_name}{{BaseSeeder: BaseSeeder{{name: "Specifications for {device_name}"}}}}
}}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *{class_name}) getBanglaTranslations() map[string]string {{
	return map[string]string{{
'''
    
    # Add Bengali translations
    for eng_val, ben_val in sorted(bengali_translations.items()):
        # Escape quotes in strings
        eng_escaped = eng_val.replace('"', '\\"')
        ben_escaped = ben_val.replace('"', '\\"')
        go_content += f'\t\t"{eng_escaped}": "{ben_escaped}",\n'
    
    go_content += '''	}
}

// Seed inserts specification records for the device
func (s *''' + class_name + ''') Seed(db *gorm.DB) error {
	deviceSlug := "''' + slug + '''"

	var prod models.ProductModel
	if err := db.Where("slug = ?", deviceSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := map[string]string{
'''
    
    # Add specifications
    for spec_key, spec_value in sorted(spec_map.items()):
        spec_key_escaped = spec_key.replace('"', '\\"')
        spec_value_escaped = spec_value.replace('"', '\\"')
        go_content += f'\t\t"{spec_key_escaped}": "{spec_value_escaped}",\n'
    
    go_content += '''	}
	banglaTranslations := s.getBanglaTranslations()

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							translation := &models.SpecificationTranslationModel{
								SpecificationID: sModel.ID,
								Locale:          "bn",
								Value:           banglaValue,
							}
							if err := db.Create(translation).Error; err != nil {
								return err
							}
						} else {
							return err
						}
					}
				}
			} else {
				return err
			}
		} else {
			// If specification already exists, check and create Bangla translation if missing
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
'''
    
    return filepath, go_content

# Main execution
print("Loading mobile_specification_translations.json...")
with open("init-db/seeders/mobile_specification_translations.json", 'r', encoding='utf-8') as f:
    mobile_data = json.load(f)

output_dir = "init-db/seeders"

# Create seeders directory if it doesn't exist
os.makedirs(output_dir, exist_ok=True)

print(f"\nGenerating Go seeder files for {len(mobile_data)} devices...\n")

created_count = 0
error_count = 0

for device_name, specs in sorted(mobile_data.items()):
    try:
        filepath, content = generate_go_file(device_name, specs, output_dir)
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(content)
        
        print(f"[OK] Created: {os.path.basename(filepath)}")
        created_count += 1
        
    except Exception as e:
        print(f"[ERROR] Error for {device_name}: {e}")
        error_count += 1

print(f"\n{'='*60}")
print(f"Summary:")
print(f"  Total Created: {created_count}")
print(f"  Total Errors: {error_count}")
print(f"  Output Directory: {output_dir}")
print(f"{'='*60}")
