#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Generate Go seeder files for Walton refrigerators
"""

import os
import re
from pathlib import Path

# List of all Walton refrigerator models
WALTON_MODELS = [
    "walton-wnm-2a7-gdel-xx",
    "walton-wnm-2f1-gehe-xx",
    "walton-wfa-2a3-gdel-sc",
    "walton-wfa-2d4-gdel-xx",
    "walton-wfb-1n3-gdxx-xx",
    "walton-wfb-2x1-gdel-xx",
    "walton-wfb-2e4-gdel-xx",
    "walton-wfe-3e8-gden-xx",
    "walton-wfd-1b6-gdeh-xx",
    "walton-wfd-1d4-gdeh-xx",
    "walton-wfd-1b6-gdel-xx",
    "walton-wfd-1d4-gdel-xx",
    "walton-wfd-1f3-gdel-xx",
    "walton-wnr-6d6-gdfs-dd",
    "walton-wnr-6f0-scrc-co",
    "walton-wcg-3j0-rxlx-xx",
    "walton-wcg-3j0-gddb-xx",
    "walton-wcg-3j0-ddge-xx",
    "walton-wcf-2a0-gsre-xx",
    "walton-wfb-2b6-elxx-xx",
    "walton-wfb-2x1-gdsh-xx",
    "walton-wnj-5h5-rxxx-xx",
    "walton-wnm-1n5-rxxx-rp",
    "walton-wfb-1h5-gdxx-xx",
    "walton-wfb-2e0-gdxx-xx",
    "walton-wfe-3a2-gdel-xx",
    "walton-wfc-3d8-gdeh-dd",
    "walton-wcg-2e5-gdel-dd-inverter",
    "walton-wcg-2e5-gdel-dd",
    "walton-wcg-2e5-ehlx-xx",
    "walton-wcg-2e5-ehlc-xx",
    "walton-wcf-2a0-gsre-xx-p",
    "walton-wcf-2t5-rrlx-gx",
    "walton-wcf-2t5-rrlx-xx",
    "walton-wcf-2t5-gdel-xx",
    "walton-wcf-1b5-gdel-xx",
    "walton-wcf-1d5-rrxx-xx",
    "walton-wcf-1d5-gele-dl",
    "walton-wcf-1d5-gdel-xx",
    "walton-wcf-1d5-gdel-lx",
    "walton-wue-3c4-gepb-xx",
    "walton-wut-3x8-gere-cx",
    "walton-wue-2g2-gepb-xx",
    "walton-wfa-1n3-gdxx-xx",
    "walton-wfb-2e0-gjxb-sx-p",
    "walton-wfe-2h2-gdxx-xx",
    "walton-wfc-3f5-gdeh-xx",
    "walton-wnh-3h6-gdel-xx",
    "walton-wni-5f3-gdne-id",
    "walton-wni-5f3-gdel-dd",
    "walton-wni-6a9-gdne-dd",
    "walton-wni-6a9-gdsd-dd",
]

def slug_to_struct_name(slug):
    """Convert slug to Go struct name"""
    # Remove 'walton-' prefix
    name = slug.replace('walton-', '')
    # Split by dashes and capitalize each part
    parts = name.split('-')
    # Capitalize each part
    struct_parts = ''.join(part.capitalize() for part in parts)
    return f"SpecificationSeederRefrigeratorWalton{struct_parts}"

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
    
    result = "ওয়ালটন-"
    name = slug.replace('walton-', '')
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
		"Walton":         "ওয়ালটন",
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

	existingkeyMapping, err := s.LoadExistingSpecificationKeys(db)
	if err != nil {{
		return err
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
    
    print(f"Generating {len(WALTON_MODELS)} Walton refrigerator seeder files...")
    
    created = 0
    skipped = 0
    
    for slug in WALTON_MODELS:
        struct_name = slug_to_struct_name(slug)
        filename = f"specification_seeder_refrigerator_{slug}.go"
        filepath = seeders_dir / filename
        
        # Check if file already exists
        if filepath.exists():
            print(f"⏭️  Skipping {filename} (already exists)")
            skipped += 1
            continue
        
        # Generate and write content
        content = generate_seeder_content(slug)
        filepath.write_text(content, encoding='utf-8')
        print(f"✅ Created {filename}")
        created += 1
    
    print(f"\n📊 Summary:")
    print(f"   Created: {created}")
    print(f"   Skipped: {skipped}")
    print(f"   Total: {len(WALTON_MODELS)}")

if __name__ == "__main__":
    main()
