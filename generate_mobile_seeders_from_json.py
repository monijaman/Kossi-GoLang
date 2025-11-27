#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Generate Go seeder files for mobile phones from mobiledata.json
Following template.go pattern with Bengali translations only (no Devanagari)
"""

import json
import re
import os
from pathlib import Path

# Bengali digit mapping
BENGALI_DIGITS = {
    '0': '০', '1': '১', '2': '২', '3': '৩', '4': '৪',
    '5': '৫', '6': '৬', '7': '৭', '8': '৮', '9': '৯'
}

def convert_to_bengali_digits(text):
    """Convert English digits to Bengali digits in text"""
    if not text:
        return text
    result = text
    for eng, ben in BENGALI_DIGITS.items():
        result = result.replace(eng, ben)
    return result

def generate_slug(name):
    """Generate slug from product name"""
    slug = name.lower()
    slug = re.sub(r'[^a-z0-9\s-]', '', slug)
    slug = re.sub(r'\s+', '-', slug)
    slug = re.sub(r'-+', '-', slug)
    return slug.strip('-')

def generate_struct_name(name):
    """Generate struct name from product name"""
    # Remove special chars and split
    clean = re.sub(r'[^a-zA-Z0-9\s]', '', name)
    words = clean.split()
    # Capitalize each word
    struct_name = 'SpecificationSeederMobile' + ''.join(word.capitalize() for word in words)
    return struct_name

def translate_value_to_bangla(value):
    """
    Translate specification value to Bangla.
    Only translate numbers to Bengali digits and common tech terms.
    Keep brand names, technical abbreviations, and proper nouns in English.
    """
    if not value or value.strip() == "":
        return ""
    
    # Skip translation for certain patterns
    skip_patterns = [
        r'^[A-Z][a-zA-Z0-9\s\-]+$',  # Brand/model names like "Apple A18 Pro"
        r'iOS',
        r'Android',
        r'^USB',
        r'^Wi-Fi',
        r'^5G$',
        r'^4G$',
        r'^3G$',
        r'^2G$',
        r'^GSM',
        r'^CDMA',
        r'^LTE',
        r'Bluetooth',
        r'NFC',
        r'GPS',
        r'GLONASS',
        r'GALILEO',
        r'BDS',
        r'QZSS',
        r'LED',
        r'OLED',
        r'AMOLED',
        r'LCD',
        r'IPS',
        r'TFT',
        r'LTPO',
        r'HDR',
        r'Dolby',
        r'ProMotion',
        r'Ceramic Shield',
        r'Corning Gorilla Glass',
        r'IP\d+',
        r'Type-C',
        r'MagSafe',
        r'Face ID',
        r'Touch ID',
        r'ProRes',
        r'Hexa-core',
        r'Octa-core',
        r'Quad-core',
        r'Mali',
        r'Adreno',
        r'Snapdragon',
        r'Dimensity',
        r'Exynos',
        r'Kirin',
        r'Bionic',
    ]
    
    # If value matches skip patterns, only convert numbers
    for pattern in skip_patterns:
        if re.search(pattern, value, re.IGNORECASE):
            return convert_to_bengali_digits(value)
    
    # Common translations
    translations = {
        "Yes": "হ্যাঁ",
        "No": "না",
        "Available": "উপলব্ধ",
        "Coming Soon": "শীঘ্রই আসছে",
        "Discontinued": "বন্ধ",
        "inches": "ইঞ্চি",
        "inch": "ইঞ্চি",
        "mm": "মিমি",
        "mAh": "এমএএইচ",
        "wired": "তারযুক্ত",
        "wireless": "বেতার",
        "Stereo speakers": "স্টেরিও স্পিকার",
        "stereo speakers": "স্টেরিও স্পিকার",
        "No 3.5mm jack": "৩.৫মিমি জ্যাক নেই",
        "non-removable": "অপসারণযোগ্য নয়",
        "Nano-SIM": "ন্যানো-সিম",
        "eSIM": "ই-সিম",
        "Dual SIM": "ডুয়াল সিম",
        "Single SIM": "সিঙ্গেল সিম",
        "accelerometer": "অ্যাক্সিলেরোমিটার",
        "gyro": "জাইরো",
        "proximity": "প্রক্সিমিটি",
        "compass": "কম্পাস",
        "barometer": "ব্যারোমিটার",
        "fingerprint": "ফিঙ্গারপ্রিন্ট",
        "under display": "ডিসপ্লের নিচে",
        "side-mounted": "পাশে মাউন্ট",
        "rear-mounted": "পেছনে মাউন্ট",
        "Li-Ion": "লি-আয়ন",
        "Li-Po": "লি-পো",
        "Fast charging": "ফাস্ট চার্জিং",
        "typical": "সাধারণ",
        "approx": "আনুমানিক",
        "up to": "পর্যন্ত",
        "Aluminum frame": "অ্যালুমিনিয়াম ফ্রেম",
        "Titanium frame": "টাইটানিয়াম ফ্রেম",
        "glass front": "গ্লাস সামনে",
        "glass back": "গ্লাস পেছনে",
        "plastic back": "প্লাস্টিক পেছনে",
        "dust/water resistant": "ধুলো/পানি প্রতিরোধী",
        "dust tight and water resistant": "ধুলো নিরোধক এবং পানি প্রতিরোধী",
        "immersible": "নিমজ্জনযোগ্য",
        "for": "জন্য",
        "min": "মিনিট",
        "Dual-LED": "ডুয়াল-LED",
        "dual-tone flash": "ডুয়াল-টোন ফ্ল্যাশ",
        "HDR": "HDR",
        "photo/panorama": "ফটো/প্যানোরামা",
        "panorama": "প্যানোরামা",
        "Optical Zoom": "অপটিক্যাল জুম",
        "Digital Zoom": "ডিজিটাল জুম",
        "via crop": "ক্রপের মাধ্যমে",
        "Emergency SOS via satellite": "স্যাটেলাইটের মাধ্যমে জরুরি SOS",
        "Black": "কালো",
        "White": "সাদা",
        "Blue": "নীল",
        "Green": "সবুজ",
        "Red": "লাল",
        "Yellow": "হলুদ",
        "Pink": "গোলাপী",
        "Purple": "বেগুনি",
        "Gold": "সোনালী",
        "Silver": "রূপালী",
        "Gray": "ধূসর",
        "Grey": "ধূসর",
        "Rose Gold": "রোজ গোল্ড",
        "Natural": "ন্যাচারাল",
        "Desert": "মরুভূমি",
        "Ultramarine": "আল্ট্রামেরিন",
        "Teal": "টিল",
        "September": "সেপ্টেম্বর",
        "October": "অক্টোবর",
        "November": "নভেম্বর",
        "December": "ডিসেম্বর",
        "January": "জানুয়ারি",
        "February": "ফেব্রুয়ারি",
        "March": "মার্চ",
        "April": "এপ্রিল",
        "May": "মে",
        "June": "জুন",
        "July": "জুলাই",
        "August": "আগস্ট",
    }
    
    # Apply word-level translations
    result = value
    for eng, ban in translations.items():
        # Use word boundaries for better matching
        result = re.sub(r'\b' + re.escape(eng) + r'\b', ban, result, flags=re.IGNORECASE)
    
    # Convert all numbers to Bengali
    result = convert_to_bengali_digits(result)
    
    return result

def escape_go_string(s):
    """Escape special characters for Go string literals"""
    # Replace backslash first
    s = s.replace('\\', '\\\\')
    # Replace double quotes
    s = s.replace('"', '\\"')
    # Replace newlines and tabs
    s = s.replace('\n', '\\n')
    s = s.replace('\r', '\\r')
    s = s.replace('\t', '\\t')
    return s

def generate_seeder_file(product_name, specs, output_dir):
    """Generate a Go seeder file for a mobile product"""
    
    slug = generate_slug(product_name)
    struct_name = generate_struct_name(product_name)
    
    # Build translations map
    translations = {}
    for key, value in specs.items():
        if value and str(value).strip():
            bangla_value = translate_value_to_bangla(str(value))
            if bangla_value and bangla_value != str(value):
                translations[str(value)] = bangla_value
    
    # Build specs assignment
    specs_assignments = []
    for key, value in specs.items():
        if value and str(value).strip():
            # Format the key in title case with spaces
            display_key = key.replace('_', ' ').title()
            escaped_key = escape_go_string(display_key)
            escaped_value = escape_go_string(str(value))
            specs_assignments.append(f'\tspecs["{escaped_key}"] = "{escaped_value}"')
    
    # Build translations map string
    translations_map = []
    for eng_val, ban_val in sorted(translations.items()):
        escaped_eng = escape_go_string(eng_val)
        escaped_ban = escape_go_string(ban_val)
        translations_map.append(f'\t\t"{escaped_eng}": "{escaped_ban}",')
    
    # Generate file content
    content = f'''package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// {struct_name} seeds specifications/options for product '{slug}'
type {struct_name} struct {{
	BaseSeeder
}}

// New{struct_name} creates a new seeder instance
func New{struct_name}() *{struct_name} {{
	return &{struct_name}{{BaseSeeder: BaseSeeder{{name: "Specifications for {product_name}"}}}}
}}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *{struct_name}) getBanglaTranslations() map[string]string {{
	return map[string]string{{
{chr(10).join(translations_map)}
	}}
}}

// Seed inserts specification records for the product identified by slug '{slug}'
func (s *{struct_name}) Seed(db *gorm.DB) error {{
	productSlug := "{slug}"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {{
		if err == gorm.ErrRecordNotFound {{
			return nil
		}}
		return err
	}}
	productID := prod.ID

	specs := DefaultMobileSpecs()
	banglaTranslations := s.getBanglaTranslations()

	// Override model-specific values for {product_name}
{chr(10).join(specs_assignments)}

	for key, value := range specs {{
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {{
			return err
		}}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {{
			if err == gorm.ErrRecordNotFound {{
				sModel := &models.SpecificationModel{{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}}
				if err := db.Create(sModel).Error; err != nil {{
					return err
				}}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {{
					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {{
						if err == gorm.ErrRecordNotFound {{
							translation := &models.SpecificationTranslationModel{{
								SpecificationID: sModel.ID,
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
			}} else {{
				return err
			}}
		}} else {{
			// If specification already exists, check and create Bangla translation if missing
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
    
    # Write to file
    filename = f"specification_seeder_mobile_{slug.replace('-', '_')}.go"
    filepath = os.path.join(output_dir, filename)
    
    with open(filepath, 'w', encoding='utf-8') as f:
        f.write(content)
    
    return filename

def main():
    # Paths
    json_file = Path(__file__).parent / "init-db" / "seeders" / "mobiledata.json"
    output_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    print(f"Reading JSON from: {json_file}")
    print(f"Output directory: {output_dir}")
    
    # Load JSON
    with open(json_file, 'r', encoding='utf-8') as f:
        data = json.load(f)
    
    print(f"\nFound {len(data)} mobile products")
    
    # Generate seeder files
    generated_files = []
    for product_name, specs in data.items():
        try:
            filename = generate_seeder_file(product_name, specs, output_dir)
            generated_files.append(filename)
            print(f"[OK] Generated: {filename}")
        except Exception as e:
            print(f"[ERROR] Error generating {product_name}: {e}")
    
    print(f"\n{'='*60}")
    print(f"Successfully generated {len(generated_files)} seeder files")
    print(f"{'='*60}")
    
    # Create a list file
    list_file = output_dir / "mobile_seeders_generated_list.txt"
    with open(list_file, 'w', encoding='utf-8') as f:
        for filename in sorted(generated_files):
            f.write(f"{filename}\n")
    
    print(f"\nFile list saved to: {list_file}")

if __name__ == "__main__":
    main()
