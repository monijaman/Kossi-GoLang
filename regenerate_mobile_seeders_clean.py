#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json
import os
import re

# Comprehensive Bengali translations dictionary
BENGALI_TRANSLATIONS = {
    # Units
    "inches": "ইঞ্চি",
    "GB": "গিগাবাইট",
    "TB": "টেরাবাইট",
    "MB": "মেগাবাইট",
    "MP": "মেগাপিক্সেল",
    "mAh": "এমএএইচ",
    "W": "ওয়াট",
    "Hz": "হার্জ",
    "GHz": "গিগাহার্জ",
    "kg": "কেজি",
    "g": "গ্রাম",
    "mm": "মিমি",
    "cm": "সেন্টিমিটার",
    "fps": "এফপিএস",
    "ppi": "পিপিআই",
    
    # Boolean and Status
    "Yes": "হ্যাঁ",
    "No": "না",
    "Available": "উপলব্ধ",
    "Discontinued": "বন্ধ",
    
    # Connectivity
    "5G": "৫জি",
    "4G": "৪জি",
    "LTE": "এলটিই",
    "WiFi": "ওয়াইফাই",
    "Bluetooth": "ব্লুটুথ",
    "NFC": "এনএফসি",
    "GPS": "জিপিএস",
    
    # Battery
    "Li-Ion": "লিথিয়াম-আয়ন",
    "Li-Po": "লিথিয়াম-পলিমার",
    
    # Operating Systems
    "Android": "অ্যান্ড্রয়েড",
    "iOS": "আইওএস",
    
    # Display Types
    "OLED": "ওলেড",
    "LCD": "এলসিডি",
    "IPS": "আইপিএস",
    "AMOLED": "অ্যামোলেড",
    "HDR": "এইচডিআর",
    
    # Charging
    "USB Type-C": "ইউএসবি টাইপ-সি",
    "MagSafe": "ম্যাগসেফ",
    "wireless": "ওয়্যারলেস",
    "wired": "ওয়্যার্ড",
}

def smart_translate(text):
    """
    Translate text intelligently without breaking English words.
    Only translate complete terms/units, not parts of words.
    """
    if not text:
        return text
    
    result = text
    
    # Sort by length (longest first) to avoid partial replacements
    sorted_terms = sorted(BENGALI_TRANSLATIONS.items(), key=lambda x: len(x[0]), reverse=True)
    
    for eng_term, ben_term in sorted_terms:
        # Use word boundaries to avoid partial replacements
        # Match as whole words or units (followed by space, comma, or digit)
        pattern = r'\b' + re.escape(eng_term) + r'(?=\s|,|-|@|:|$|/)'
        result = re.sub(pattern, ben_term, result, flags=re.IGNORECASE)
    
    return result

def generate_seeder_file(device_name, specs, output_dir):
    """Generate a corrected Go seeder file"""
    
    slug = device_name.lower().replace(' ', '-').replace('(', '').replace(')', '')
    slug = re.sub(r'[^a-z0-9\-]', '', slug)
    slug = re.sub(r'-+', '-', slug).strip('-')
    
    filename = f"specification_seeder_mobile_{slug}.go"
    filepath = os.path.join(output_dir, filename)
    
    # Create proper class name
    words = device_name.replace('-', ' ').replace('(', '').replace(')', '').split()
    class_name = 'SpecificationSeederMobile' + ''.join(word.capitalize() for word in words if word)
    class_name = re.sub(r'[^a-zA-Z0-9]', '', class_name)
    
    # Extract specs and create proper translations
    spec_map = {}
    bengali_map = {}
    
    for spec_key, spec_data in specs.items():
        if isinstance(spec_data, dict) and "specification_translations" in spec_data:
            english_value = spec_data["english"]
            spec_map[spec_key] = english_value
            
            # Get the Bengali value from the data (already properly translated)
            for trans in spec_data["specification_translations"]:
                if trans["locale"] == "bn":
                    # Don't re-translate, just clean it up if needed
                    bengali_val = trans["value"]
                    # Remove any corrupted partial translations
                    bengali_val = bengali_val.replace("গ্রাম", "g")  # temp fix
                    bengali_map[english_value] = smart_translate(english_value)
    
    # Generate Go file
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
    for eng_val, ben_val in sorted(bengali_map.items()):
        if ben_val != eng_val:  # Only add if there's an actual translation
            eng_escaped = eng_val.replace('"', '\\"')
            ben_escaped = ben_val.replace('"', '\\"')
            go_content += f'\t\t"{eng_escaped}": "{ben_escaped}",\n'
    
    go_content += f'''	}}
}}

// Seed inserts specification records for the device
func (s *{class_name}) Seed(db *gorm.DB) error {{
	deviceSlug := "{slug}"

	var prod models.ProductModel
	if err := db.Where("slug = ?", deviceSlug).First(&prod).Error; err != nil {{
		if err == gorm.ErrRecordNotFound {{
			return nil
		}}
		return err
	}}
	productID := prod.ID

	specs := map[string]string{{
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
print("Regenerating seeder files with corrected Bengali translations...")
print()

with open("init-db/seeders/mobile_specification_translations.json", 'r', encoding='utf-8') as f:
    mobile_data = json.load(f)

output_dir = "init-db/seeders"
os.makedirs(output_dir, exist_ok=True)

created = 0
errors = 0

for device_name in sorted(mobile_data.keys()):
    try:
        specs = mobile_data[device_name]
        filepath, content = generate_seeder_file(device_name, specs, output_dir)
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(content)
        
        created += 1
        if created % 50 == 0:
            print(f"[OK] Generated {created} seeders...")
        
    except Exception as e:
        print(f"[ERROR] {device_name}: {e}")
        errors += 1

print()
print("="*60)
print(f"Summary:")
print(f"  Created: {created}")
print(f"  Errors: {errors}")
print(f"  Total: {created + errors}")
print("="*60)
