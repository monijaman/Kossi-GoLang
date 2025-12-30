#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Generate Go seeder files for Marcel refrigerators
"""

import os
import re
from pathlib import Path

# List of all Marcel refrigerator models
MARCEL_MODELS = [
    "marcel-mfc-c6e-gdeh-dd-inverter",
    "marcel-mfc-c6e-gdel-xx",
    "marcel-mfc-c6e-gdel-xx-inverter",
    "marcel-mfc-c6e-gdne-xx-inverter",
    "marcel-mfc-c6e-gdxx-xx-inverter",
    "marcel-mfc-c6e-i5-gede-fd-inverter",
    "marcel-mfc-c6e-gdne-xx",
    "marcel-mfc-c6e-nexx-xx",
    "marcel-mfe-c5h-elex-xx",
    "marcel-mfe-c5h-elnx-xx",
    "marcel-mfe-c5h-gdel-xx",
    "marcel-mfe-c5h-gden-dd-inverter",
    "marcel-mfe-c5h-gdxx-xx",
    "marcel-mfc-c4h-gdeh-dd-inverter",
    "marcel-mfc-c4h-gdel-xx-inverter",
    "marcel-mfc-c4h-gdne-xx",
    "marcel-mfc-c4h-gdxx-xx",
    "marcel-mfc-c4h-gdxx-xx-inverter",
    "marcel-mfc-c4h-gjxb-fd-inverter",
    "marcel-mfk-c4g-gdel-xx",
    "marcel-mfk-c4g-gdel-xx-inverter",
    "marcel-mfc-c4h-nexx-xx",
    "marcel-mfe-c3c-gden-dd-inverter",
    "marcel-mfe-c3c-gden-xx",
    "marcel-mfe-c3c-gdxx-xx",
    "marcel-mfe-c2x-crxx-xx",
    "marcel-mfe-c2x-crxx-xx-inverter",
    "marcel-mfe-c2x-gdel-xx",
    "marcel-mfe-c2x-gdel-xx-inverter",
    "marcel-mfe-c2x-gden-dd-inverter",
    "marcel-mfe-c2x-gden-xx",
    "marcel-mfe-c2x-gden-xx-inverter",
    "marcel-mfe-c2x-gdxx-xx",
    "marcel-mfe-c2x-gdxx-xx-inverter",
    "marcel-mfc-c1g-elex-xx",
    "marcel-mfc-c1g-gdne-xx",
    "marcel-mfc-c1g-gdxx-xx",
    "marcel-mfc-c1g-gdxx-xx-inverter",
    "marcel-mfc-c1g-nxxx-xx",
    "marcel-mfc-c1b-crxx-xx",
    "marcel-mfe-c1b-elnx-xx",
    "marcel-mfe-c1b-gdel-xx",
    "marcel-mfe-c1b-gdel-xx-inverter",
    "marcel-mfe-c1b-gden-xx",
    "marcel-mfe-c1b-gden-xx-inverter",
    "marcel-mfe-c1b-gdxx-xx",
    "marcel-mfe-c1b-gdxx-xx-inverter",
    "marcel-mfe-c0n-gdel-xx",
    "marcel-mfe-c0n-gdel-xx-inverter",
    "marcel-mfe-c0n-gdxx-xx",
    "marcel-mfe-c0n-gdxx-xx-inverter",
    "marcel-mfc-c0g-elxx-xx",
    "marcel-mfc-c0g-gdeh-dd-inverter",
    "marcel-mfc-c0g-gdxx-xx",
    "marcel-mfc-c0g-gdxx-xx-inverter",
    "marcel-mfc-c0g-nexx-xx",
    "marcel-mfe-c0n-elnx-xx",
    "marcel-mfe-b9e-crxx-xx",
    "marcel-mfe-b9e-crxx-xx-inverter",
    "marcel-mfe-b9e-elnx-xx",
    "marcel-mfe-b9e-gdel-xx-inverter",
    "marcel-mfb-b5d-gdsh-xx",
    "marcel-mfb-b5d-gdsh-xx-inverter",
    "marcel-mfb-b5d-gdxx-xx",
    "marcel-mfb-b5d-gdxx-xx-inverter",
    "marcel-mfb-b5x-gdeh-sc-inverter",
    "marcel-mfb-b5x-gdeh-xx",
    "marcel-mfb-b5x-gdeh-xx-inverter",
    "marcel-mfb-b5x-gdel-sc",
    "marcel-mfb-b5x-gdel-sc-inverter",
    "marcel-mfb-b5x-gdel-xx",
    "marcel-mfb-b5x-gdel-xx-inverter",
    "marcel-mfb-b5x-gdsh-xx",
    "marcel-mfb-b5x-gdsh-xx-inverter",
    "marcel-mfb-b5x-gdxx-xx",
    "marcel-mfb-b5x-gdxx-xx-inverter",
    "marcel-mfa-b4d-gdeh-xx",
    "marcel-mfa-b4d-gdel-xx",
    "marcel-mfa-b4d-gdxx-xx",
    "marcel-mfa-b4d-nexx-xx",
    "marcel-mfa-b4d-rlxx-xx-inverter",
    "marcel-mfe-b9e-gden-xx",
    "marcel-mfe-b9e-gden-xx-inverter",
    "marcel-mfe-b9e-gdxx-xx",
    "marcel-mfe-b9e-gdxx-xx-inverter",
    "marcel-mfe-b8b-crxx-xx",
    "marcel-mfe-b8b-crxx-xx-inverter",
    "marcel-mfe-b8b-gdel-xx",
    "marcel-mfe-b8b-gdel-xx-inverter",
    "marcel-mfe-b8b-gden-dd-inverter",
    "marcel-mfe-b8b-gden-xx",
    "marcel-mfe-b8b-gden-xx-inverter",
    "marcel-mfe-b8b-gdxx-xx",
    "marcel-mfe-b8b-gdxx-xx-inverter",
    "marcel-mfb-b5d-elxx-xx",
    "marcel-mfb-b5d-gaxb-wd-inverter",
    "marcel-mfb-b5d-gdeh-sc-inverter",
    "marcel-mfb-b5d-gdeh-xx-inverter",
    "marcel-mfb-b5d-gdel-sc-inverter",
    "marcel-mfb-b5d-gdel-xx",
    "marcel-mfa-b4d-rxxx-rp",
    "marcel-mfb-b2f-elxx-xx",
    "marcel-mfb-b2f-gdeh-sc-inverter",
    "marcel-mfb-b2f-gdeh-xx",
    "marcel-mfb-b2f-gdel-sc",
    "marcel-mfb-b2f-gdel-sc-inverter",
    "marcel-mfb-b2f-gdel-xx",
    "marcel-mfb-b2f-gdxx-xx",
    "marcel-mfb-b2f-rnxx-rp",
    "marcel-mfb-b2c-gdsh-xx",
    "marcel-mfa-b2x-gdel-xx",
    "marcel-mfa-b2x-gdxx-xx",
    "marcel-mfa-b2x-rxxx-xx",
    "marcel-mfb-b1h-gdel-xx",
    "marcel-mfb-b1h-gdxx-xx",
    "marcel-mfb-b0a-elxx-xx",
    "marcel-mfb-b0a-rnxx-rp",
    "marcel-mfa-2a3-elxx-xx",
    "marcel-mfa-2a3-gdeh-xx",
    "marcel-mfa-2a3-gdel-sc",
    "marcel-mfa-2a3-gdel-xx",
    "marcel-mfa-2a3-gdsh-xx",
    "marcel-mfa-2a3-gdxx-xx",
    "marcel-mfa-2a3-nexx-xx",
    "marcel-mfa-2a3-rlxx-xx",
    "marcel-mfa-a9c-gdxx-xx",
    "marcel-mnm-b1g-gdel-xx",
    "marcel-mfb-b0a-gdel-xx",
    "marcel-mfb-b0a-gdsh-xx",
    "marcel-mfb-b0a-gdxx-xx",
    "marcel-mfa-a9c-gdsh-xx",
    "marcel-mfb-a8e-elxx-xx",
    "marcel-mfb-a8e-gdel-xx",
    "marcel-mfb-a8e-gdsh-xx",
    "marcel-mfb-a8e-gdxx-xx",
    "marcel-mfb-a8e-rnxx-rp",
    "marcel-mfb-a7g-gdsh-xx",
    "marcel-mfd-a6c-gdel-sc-inverter",
    "marcel-mfd-a6c-gdel-xx",
    "marcel-mfd-a4d-gdeh-xx",
    "marcel-mfd-a4d-gdel-xx",
    "marcel-mfd-a5x-gdel-xx",
    "marcel-mfd-1b6-gdeh-xx",
    "marcel-mfd-1b6-gdel-xx",
    "marcel-mfd-1b6-rdxx-xx",
    "marcel-mfd-1b6-rxxx-xx",
    "marcel-mfo-a0a-rxxx-xx",
    "marcel-mfs-t9c-c2sr-vb",
    "marcel-mfo-jet-rxxx-xx",
    "marcel-mfd-a4d-gdsh-xx",
    "marcel-mfd-a7x-gdeh-xx",
    "marcel-mfd-a7x-gdel-xx",
    "marcel-mfd-a7x-gdsh-xx",
    "marcel-mfd-1b6-gdsh-xx",
    "marcel-mfe-c1b-gden-dd",
    "marcel-mfb-a7g-gdxx-xx",
    "marcel-mfb-b2c-gdeh-sc-inverter",
    "marcel-mfb-b2c-gdeh-xx",
    "marcel-mfb-b2c-gdel-sc-inverter",
    "marcel-mfb-b2c-gdel-xx",
    "marcel-mfb-b2c-gdel-xx-inverter",
    "marcel-mfb-b2c-gdxx-xx",
    "marcel-mfe-c5h-gden-xx",
    "marcel-mfd-a6c-rxxx-xx",
    "marcel-mfd-a4d-rxxx-xx",
    "marcel-mnm-b1g-rxxx-rp",
    "marcel-mni-f1n-gdne-dd",
    "marcel-mni-f1n-gdne-wd",
    "marcel-mni-e6c-gdel-dd",
    "marcel-mni-e6c-gdne-dd",
    "marcel-mnh-d3x-gdel-xx-inverter",
    "marcel-mnh-d3x-hdxx-xx-inverter",
    "marcel-mnh-c8f-gdel-xx-inverter",
    "marcel-mcg-c0t-ddge-xx",
    "marcel-mcf-b0e-gdel-gx",
    "marcel-mcf-b0e-gdel-xx",
    "marcel-mcf-b0e-rrlx-gx",
    "marcel-mcf-b0e-rrlx-xx",
    "marcel-mcg-b5e-ehlc-xx",
    "marcel-mcg-b5e-ehlx-xx",
    "marcel-mcg-b5e-gdel-xx",
    "marcel-mcg-b5e-gdel-xx-inverter",
    "marcel-mcg-b7t-cgxx-xx",
    "marcel-mcg-c0t-ddge-xx-inverter",
    "marcel-mcg-c0t-gddb-xx-inverter",
    "marcel-mcg-c0t-rxlx-xx",
    "marcel-mcf-a2e-gdel-xx",
    "marcel-mcf-a4e-gdel-lx",
    "marcel-mcf-a4e-gdel-xx",
    "marcel-mcf-a4e-gele-dl",
    "marcel-mcf-a4e-rrxx-xx",
    "marcel-mbq-d4x-tdxx-xx",
    "marcel-mba-b6x-gcxb-xx",
    "marcel-mbb-b6x-tdxx-xx"
]

def slug_to_struct_name(slug):
    """Convert slug to Go struct name"""
    # Remove 'marcel-' prefix
    name = slug.replace('marcel-', '')
    # Split by dashes and capitalize each part
    parts = name.split('-')
    # Capitalize each part
    struct_parts = ''.join(part.capitalize() for part in parts)
    return f"SpecificationSeederRefrigeratorMarcel{struct_parts}"

def slug_to_bangla(slug):
    """Convert slug to Bengali transliteration"""
    # Simple transliteration - convert letters to Bengali
    bengali_map = {
        '0': '০', '1': '১', '2': '২', '3': '৩', '4': '৪',
        '5': '৫', '6': '৬', '7': '৭', '8': '৮', '9': '৯',
        'a': 'এ', 'b': 'বি', 'c': 'সি', 'd': 'ডি', 'e': 'ই',
        'f': 'এফ', 'g': 'জি', 'h': 'এইচ', 'i': 'আই', 'j': 'জে',
        'k': 'কে', 'l': 'এল', 'm': 'এম', 'n': 'এন', 'o': 'ও',
        'p': 'পি', 'q': 'কিউ', 'r': 'আর', 's': 'এস', 't': 'টি',
        'u': 'ইউ', 'v': 'ভি', 'w': 'ডব্লিউ', 'x': 'এক্স', 'y': 'ওয়াই', 'z': 'জেড',
        '-': '-'
    }
    
    result = "মার্সেল-"
    name = slug.replace('marcel-', '')
    for char in name:
        result += bengali_map.get(char.lower(), char)
    
    return result

def generate_seeder_content(slug):
    """Generate the complete seeder file content"""
    struct_name = slug_to_struct_name(slug)
    func_name = struct_name
    bangla_slug = slug_to_bangla(slug)
    
    return f'''package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// {struct_name} seeds specifications/options for product '{slug}'
type {struct_name} struct {{
	BaseSeeder
}}

// New{struct_name} creates a new seeder instance
func New{struct_name}() *{struct_name} {{
	return &{struct_name}{{
		BaseSeeder: BaseSeeder{{name: "Specifications for {slug}"}},
	}}
}}

func (s *{struct_name}) getBanglaTranslations() map[string]string {{
	return map[string]string{{
		"Marcel":         "মার্সেল",
		"{slug}":         "{bangla_slug}",
		// Add more translations as needed
	}}
}}

// Seed inserts specification records for the product identified by slug '{slug}'
func (s *{struct_name}) Seed(db *gorm.DB) error {{
	productSlug := "{slug}"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {{
		if err == gorm.ErrRecordNotFound {{
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}}
		return err
	}}
	productID := prod.ID

	existingkeyMapping := map[string]uint{{
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}}

	specs := map[string]string{{
		// Specifications will be populated from the database
		// Add your specifications here as they become available
	}}

	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {{
		specKeyID, exists := existingkeyMapping[key]
		if !exists {{
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
			continue
		}}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {{
			if err == gorm.ErrRecordNotFound {{
				sModel := &models.SpecificationModel{{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}}
				if err := db.Create(sModel).Error; err != nil {{
					return err
				}}
				// Ensure the ID is set after creation
				if sModel.ID == 0 {{
					if err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {{
						return err
					}}
				}}
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {{
					translation := &models.SpecificationTranslationModel{{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}}
					if err := db.Create(translation).Error; err != nil {{
						return err
					}}
				}}
			}} else {{
				return err
			}}
		}} else {{
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {{
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {{
					if err == gorm.ErrRecordNotFound {{
						translation := &models.SpecificationTranslationModel{{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}}
						if err := db.Create(translation).Error; err != nil {{
							return err
						}}
					}} else {{
						return err
					}}
				}}
			}}
		}}
	}}

	return nil
}}
'''

def main():
    """Generate all seeder files"""
    # Get the seeders directory
    script_dir = Path(__file__).parent
    seeders_dir = script_dir / "internal" / "infrastructure" / "database" / "seeders"
    
    print(f"Generating {len(MARCEL_MODELS)} Marcel refrigerator seeder files...")
    
    created = 0
    skipped = 0
    
    for slug in MARCEL_MODELS:
        struct_name = slug_to_struct_name(slug)
        filename = f"specification_seeder_refrigerator_{slug}.go"
        filepath = seeders_dir / filename
        
        # Check if file already exists
        if filepath.exists():
            print(f"🔄 Overwriting {filename}")
            # Remove the file to regenerate
            filepath.unlink()
        
        # Generate and write content
        content = generate_seeder_content(slug)
        filepath.write_text(content, encoding='utf-8')
        print(f"✅ Created {filename}")
        created += 1
    
    print(f"\n📊 Summary:")
    print(f"   Created: {created}")
    print(f"   Skipped: {skipped}")
    print(f"   Total: {len(MARCEL_MODELS)}")

if __name__ == "__main__":
    main()