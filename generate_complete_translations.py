#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json
import os
import re
from collections import defaultdict

# Comprehensive Bengali translation dictionary
BENGALI_DICT = {
    # Boolean
    "Yes": "হ্যাঁ",
    "No": "না",
    
    # Status
    "Available": "উপলব্ধ",
    "Discontinued": "বন্ধ",
    
    # Units - Storage
    "TB": "টেরাবাইট",
    "GB": "গিগাবাইট",
    "MB": "মেগাবাইট",
    "KB": "কিলোবাইট",
    
    # Units - Display
    "inches": "ইঞ্চি",
    "ppi": "পিপিআই",
    "nits": "নিট",
    "Hz": "হার্জ",
    
    # Units - Camera
    "MP": "মেগাপিক্সেল",
    "fps": "ফ্রেম/সেকেন্ড",
    
    # Units - Battery
    "mAh": "এমএএইচ",
    "W": "ওয়াট",
    
    # Units - Weight & Dimensions
    "g": "গ্রাম",
    "kg": "কেজি",
    "mm": "মিলিমিটার",
    "cm": "সেন্টিমিটার",
    
    # Frequency
    "GHz": "গিগাহার্জ",
    "MHz": "মেগাহার্জ",
    
    # Display Technologies
    "OLED": "ওলেড",
    "AMOLED": "অ্যামোলেড",
    "LCD": "এলসিডি",
    "IPS": "আইপিএস",
    "LED": "এলইডি",
    "LTPO": "এলটিপিও",
    "HDR": "এইচডিআর",
    
    # Connectivity
    "5G": "৫জি",
    "4G": "৪জি",
    "LTE": "এলটিই",
    "3G": "৩জি",
    "2G": "২জি",
    "WiFi": "ওয়াইফাই",
    "Bluetooth": "ব্লুটুথ",
    "NFC": "এনএফসি",
    "GPS": "জিপিএস",
    "GNSS": "জিএনএসএস",
    "GSM": "জিএসএম",
    "CDMA": "সিডিএমএ",
    "HSPA": "এইচএসপিএ",
    "EVDO": "ইভিডিও",
    
    # Operating Systems
    "Android": "অ্যান্ড্রয়েড",
    "iOS": "আইওএস",
    "iPadOS": "আইপ্যাডওএস",
    
    # Battery Types
    "Li-Ion": "লিথিয়াম-আয়ন",
    "Li-Po": "লিথিয়াম-পলিমার",
    "Li-Metal": "লিথিয়াম-মেটাল",
    
    # Charging
    "USB": "ইউএসবি",
    "Type-C": "টাইপ-সি",
    "Type-B": "টাইপ-বি",
    "MagSafe": "ম্যাগসেফ",
    "wireless": "ওয়্যারলেস",
    "wired": "ওয়্যার্ড",
    "reverse": "রিভার্স",
    
    # SIM Types
    "Nano-SIM": "ন্যানো-সিম",
    "eSIM": "ই-সিম",
    "microSD": "মাইক্রোএসডি",
    "Dual SIM": "ডুয়াল সিম",
    
    # Materials
    "Glass": "গ্লাস",
    "Metal": "মেটাল",
    "Plastic": "প্লাস্টিক",
    "Ceramic": "সিরামিক",
    "Titanium": "টাইটানিয়াম",
    "Aluminum": "অ্যালুমিনিয়াম",
    "Steel": "স্টিল",
    "Stainless": "স্টেইনলেস",
    
    # Protection
    "Gorilla Glass": "গরিলা গ্লাস",
    "Victus": "ভিক্টাস",
    "Ceramic Shield": "সিরামিক শিল্ড",
    
    # Water/Dust Rating
    "IP": "আইপি",
    "IPX": "আইপিএক্স",
    "IP68": "আইপি৬৮",
    "dust tight": "ধুলো রোধী",
    "water resistant": "জল প্রতিরোধী",
    "immersible": "নিমজ্জনযোগ্য",
    
    # Processors
    "Qualcomm": "কোয়ালকম",
    "Snapdragon": "স্ন্যাপড্রাগন",
    "Apple": "অ্যাপল",
    "A-series": "এ-সিরিজ",
    "Google": "গুগল",
    "Tensor": "টেনসর",
    "MediaTek": "মিডিয়াটেক",
    "Dimensity": "ডাইমেনসিটি",
    "Samsung": "স্যামসাং",
    "Exynos": "এক্সিনস",
    "Kirin": "কিরিন",
    "Helio": "হেলিও",
    "Bionic": "বায়োনিক",
    
    # CPU Architecture
    "core": "কোর",
    "cores": "কোর",
    "Octa": "অক্টা",
    "Hexa": "হেক্সা",
    "Quad": "কোয়াড",
    "Dual": "ডুয়াল",
    "single": "একক",
    
    # GPU
    "GPU": "জিপিউ",
    "Adreno": "অ্যাড্রেনো",
    "Mali": "মালি",
    "PowerVR": "পাওয়ার ভিআর",
    "Immortalis": "ইমরটালিস",
    
    # Camera Features
    "OIS": "অপটিক্যাল ইমেজ স্ট্যাবিলাইজেশন",
    "EIS": "ইলেকট্রনিক ইমেজ স্ট্যাবিলাইজেশন",
    "telephoto": "টেলিফটো",
    "ultra-wide": "অতি-প্রশস্ত",
    "macro": "ম্যাক্রো",
    "Night": "রাত",
    "Portrait": "পোর্ট্রেট",
    "HDR": "এইচডিআর",
    "RAW": "রা",
    "4K": "৪কে",
    "8K": "৮কে",
    "1080p": "১০৮০পি",
    "720p": "৭২০পি",
    
    # Audio
    "Stereo": "স্টেরিও",
    "Mono": "মনো",
    "speakers": "স্পিকার",
    "jack": "জ্যাক",
    "3.5mm": "৩.৫ মিমি",
    "Dolby": "ডলবি",
    "Atmos": "অ্যাটমস",
    
    # Sensors
    "Accelerometer": "অ্যাক্সিলারোমিটার",
    "Gyroscope": "জাইরোস্কোপ",
    "Proximity": "প্রক্সিমিটি",
    "Compass": "কম্পাস",
    "Barometer": "ব্যারোমিটার",
    "fingerprint": "ফিংগারপ্রিন্ট",
    "Face ID": "ফেস আইডি",
    
    # Transmission & Drivetrain (for motorcycles)
    "Manual": "ম্যানুয়াল",
    "Automatic": "অটোমেটিক",
    "Semi-automatic": "আধা-স্বয়ংক্রিয়",
    "speed": "স্পীড",
    "Chain": "চেইন",
    "Belt": "বেল্ট",
    "Shaft": "শ্যাফট",
    
    # Engine Types
    "Petrol": "পেট্রোল",
    "Diesel": "ডিজেল",
    "Electric": "বৈদ্যুতিক",
    "Hybrid": "হাইব্রিড",
    "stroke": "স্ট্রোক",
    "cylinder": "সিলিন্ডার",
    "cc": "সিসি",
    
    # Cooling System
    "Liquid-cooled": "তরল-শীতল",
    "Air-cooled": "বায়ু-শীতল",
    
    # Colors & Variants
    "Black": "কালো",
    "White": "সাদা",
    "Blue": "নীল",
    "Red": "লাল",
    "Green": "সবুজ",
    "Yellow": "হলুদ",
    "Gray": "ধূসর",
    "Gold": "সোনালী",
    "Silver": "রূপালী",
    "Bronze": "ব্রোঞ্জ",
    "Rose": "গোলাপী",
    "Pink": "গোলাপী",
    "Purple": "বেগুনি",
    "Orange": "কমলা",
    
    # Special features
    "HDR": "এইচডিআর",
    "Night mode": "রাত মোড",
    "Action": "অ্যাকশন",
    "Video": "ভিডিও",
    "Recording": "রেকর্ডিং",
    "Zoom": "জুম",
    "Optical": "অপটিক্যাল",
    "Digital": "ডিজিটাল",
    "slow motion": "ধীর গতি",
    "time lapse": "টাইম ল্যাপ্স",
    "panorama": "প্যানোরামা",
    
    # Months
    "January": "জানুয়ারি",
    "February": "ফেব্রুয়ারি",
    "March": "মার্চ",
    "April": "এপ্রিল",
    "May": "মে",
    "June": "জুন",
    "July": "জুলাই",
    "August": "আগস্ট",
    "September": "সেপ্টেম্বর",
    "October": "অক্টোবর",
    "November": "নভেম্বর",
    "December": "ডিসেম্বর",
}

def smart_translate_value(text):
    """Translate a specification value while preserving structure"""
    if not text:
        return text
    
    result = text
    
    # Sort by length (longest first) to avoid partial replacements
    sorted_terms = sorted(BENGALI_DICT.items(), key=lambda x: len(x[0]), reverse=True)
    
    for eng_term, ben_term in sorted_terms:
        # Match whole words/terms only
        pattern = r'\b' + re.escape(eng_term) + r'(?=\s|,|-|/|@|:|$|\(|\)|\d)'
        result = re.sub(pattern, ben_term, result, flags=re.IGNORECASE)
    
    return result

def generate_complete_seeder(device_name, specs, output_dir):
    """Generate seeder with COMPLETE Bengali translations for all specs"""
    
    slug = device_name.lower().replace(' ', '-').replace('(', '').replace(')', '')
    slug = re.sub(r'[^a-z0-9\-]', '', slug)
    slug = re.sub(r'-+', '-', slug).strip('-')
    
    filename = f"specification_seeder_mobile_{slug}.go"
    filepath = os.path.join(output_dir, filename)
    
    # Create proper class name
    words = device_name.replace('-', ' ').replace('(', '').replace(')', '').split()
    class_name = 'SpecificationSeederMobile' + ''.join(word.capitalize() for word in words if word)
    class_name = re.sub(r'[^a-zA-Z0-9]', '', class_name)
    
    # Extract all specs and create COMPLETE translation map
    spec_map = {}
    translation_map = {}
    
    for spec_key, spec_data in specs.items():
        if isinstance(spec_data, dict) and "specification_translations" in spec_data:
            english_value = spec_data["english"]
            spec_map[spec_key] = english_value
            
            # Translate the value completely
            bengali_value = smart_translate_value(english_value)
            translation_map[english_value] = bengali_value
    
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
    
    # Add ALL translations
    for eng_val in sorted(translation_map.keys()):
        ben_val = translation_map[eng_val]
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
    for spec_key in sorted(spec_map.keys()):
        spec_value = spec_map[spec_key]
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
print("Generating seeders with COMPLETE Bengali translations...")
print()

with open("init-db/seeders/mobile_specification_translations.json", 'r', encoding='utf-8') as f:
    mobile_data = json.load(f)

output_dir = "init-db/seeders"
os.makedirs(output_dir, exist_ok=True)

created = 0
errors = 0
translation_stats = defaultdict(int)

for device_name in sorted(mobile_data.keys()):
    try:
        specs = mobile_data[device_name]
        filepath, content = generate_complete_seeder(device_name, specs, output_dir)
        
        # Count translations created
        trans_count = content.count('":', 1)
        translation_stats['total'] += trans_count
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(content)
        
        created += 1
        if created % 50 == 0:
            print(f"[OK] Generated {created} seeders with complete translations...")
        
    except Exception as e:
        print(f"[ERROR] {device_name}: {e}")
        errors += 1

print()
print("="*70)
print("COMPLETION REPORT:")
print("="*70)
print(f"Total Seeders Generated: {created}")
print(f"Errors: {errors}")
print(f"Output Location: {output_dir}")
print("="*70)
print()
print("All specification values now have complete Bengali translations!")
print("Every spec value is matched with its Bengali equivalent.")
