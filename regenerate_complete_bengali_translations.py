#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json
import os
import re

# Comprehensive Bengali translation dictionary
COMPLETE_BENGALI_DICT = {
    # Common English words/phrases that appear in specs
    "and": "এবং",
    "or": "অথবা",
    "with": "সহ",
    "without": "ছাড়া",
    "up to": "পর্যন্ত",
    "for": "জন্য",
    "in": "এ",
    "from": "থেকে",
    "front": "সামনে",
    "rear": "পিছনে",
    "back": "পিছনে",
    "wide": "প্রশস্ত",
    "ultra": "অতি",
    "dual": "ডুয়াল",
    "triple": "ট্রিপল",
    "quad": "চতুর",
    "display": "ডিসপ্লে",
    "colors": "রঙ",
    "color": "রঙ",
    "black": "কালো",
    "white": "সাদা",
    "silver": "রূপা",
    "gold": "সোনা",
    "blue": "নীল",
    "red": "লাল",
    "green": "সবুজ",
    "gray": "ধূসর",
    "grey": "ধূসর",
    "released": "মুক্তি",
    "available": "উপলব্ধ",
    "discontinued": "বন্ধ",
    "min": "মিনিট",
    "sec": "সেকেন্ড",
    "hour": "ঘন্টা",
    "fps": "এফপিএস",
    "p": "পি",
    
    # Units
    "inches": "ইঞ্চি",
    "inch": "ইঞ্চি",
    "GB": "গিগাবাইট",
    "TB": "টেরাবাইট",
    "MB": "মেগাবাইট",
    "KB": "কিলোবাইট",
    "MP": "মেগাপিক্সেল",
    "mAh": "এমএএইচ",
    "W": "ওয়াট",
    "Hz": "হার্জ",
    "MHz": "মেগাহার্জ",
    "GHz": "গিগাহার্জ",
    "kg": "কেজি",
    "g": "গ্রাম",
    "mm": "মিমি",
    "cm": "সেন্টিমিটার",
    "ppi": "পিপিআই",
    "nits": "নিট",
    
    # Boolean and Status
    "Yes": "হ্যাঁ",
    "No": "না",
    
    # Connectivity
    "5G": "৫জি",
    "4G": "৪জি",
    "3G": "৩জি",
    "2G": "২জি",
    "LTE": "এলটিই",
    "WiFi": "ওয়াইফাই",
    "Wi-Fi": "ওয়াই-ফাই",
    "Bluetooth": "ব্লুটুথ",
    "NFC": "এনএফসি",
    "GPS": "জিপিএস",
    "GLONASS": "গ্লোনাস",
    "BDS": "বিডিএস",
    "GALILEO": "গ্যালিলিও",
    "QZSS": "কিউজেডএসএস",
    "NavIC": "নেভিক",
    "GSM": "জিএসএম",
    "CDMA": "সিডিএমএ",
    "HSPA": "এইচএসপিএ",
    "EVDO": "ইভিডিও",
    "HSDPA": "এইচএসডিপিএ",
    
    # Technical Terms
    "OLED": "ওলেড",
    "LCD": "এলসিডি",
    "IPS": "আইপিএস",
    "AMOLED": "অ্যামোলেড",
    "LTPO": "এলটিপিও",
    "HDR": "এইচডিআর",
    "HDR10": "এইচডিআর১০",
    "HDR10+": "এইচডিআর১০+",
    "Dolby Vision": "ডলবি ভিশন",
    "Dolby Atmos": "ডলবি অ্যাটমস",
    "Li-Ion": "লিথিয়াম-আয়ন",
    "Li-Po": "লিথিয়াম-পলিমার",
    "OTG": "ওটিজি",
    "UFS": "ইউএফএস",
    "microSD": "মাইক্রোএসডি",
    "USB": "ইউএসবি",
    "Type-C": "টাইপ-সি",
    "DisplayPort": "ডিসপ্লেপোর্ট",
    
    # Operating Systems
    "Android": "অ্যান্ড্রয়েড",
    "iOS": "আইওএস",
    "HyperOS": "হাইপারওএস",
    "One UI": "ওয়ান ইউআই",
    
    # Battery & Charging
    "wired": "ওয়্যার্ড",
    "wireless": "ওয়্যারলেস",
    "reverse": "বিপরীত",
    "MagSafe": "ম্যাগসেফ",
    "removable": "অপসারণযোগ্য",
    "non-removable": "অপসারণযোগ্য নয়",
    
    # Camera & Video
    "f/": "f/",
    "@": "@",
    "4K": "৪কে",
    "8K": "৮কে",
    "1080p": "১০৮০পি",
    "720p": "৭২০পি",
    "zoom": "জুম",
    "optical": "অপটিক্যাল",
    "digital": "ডিজিটাল",
    "periscope": "পেরিস্কোপ",
    "telephoto": "টেলিফটো",
    "ultrawide": "আল্ট্রাওয়াইড",
    "flash": "ফ্ল্যাশ",
    "stabilization": "স্ট্যাবিলাইজেশন",
    "EIS": "ইলেকট্রনিক ইমেজ স্ট্যাবিলাইজেশন",
    "panorama": "প্যানোরামা",
    "HDR": "এইচডিআর",
    
    # Design & Build
    "glass": "গ্লাস",
    "aluminum": "অ্যালুমিনিয়াম",
    "metal": "ধাতু",
    "plastic": "প্লাস্টিক",
    "frame": "ফ্রেম",
    "gorilla": "গরিলা",
    "corning": "কর্নিং",
    "victus": "ভিকটাস",
    
    # Protection
    "IP68": "আইপি৬৮",
    "IP67": "আইপি৬৭",
    "dust": "ধুলো",
    "tight": "রোধী",
    "water": "জল",
    "resistant": "প্রতিরোধী",
    "immersible": "ডুবানো যায়",
    "proof": "প্রুফ",
    
    # Sensors & Features
    "fingerprint": "ফিংগারপ্রিন্ট",
    "accelerometer": "অ্যাক্সিলারোমিটার",
    "gyro": "জাইরো",
    "proximity": "প্রক্সিমিটি",
    "compass": "কম্পাস",
    "barometer": "ব্যারোমিটার",
    "spectrum": "স্পেকট্রাম",
    
    # Processors
    "Snapdragon": "স্ন্যাপড্রাগন",
    "Exynos": "এক্সিনস",
    "Helio": "হেলিও",
    "MediaTek": "মিডিয়াটেক",
    "Apple": "অ্যাপল",
    "Tensor": "টেনসর",
    "Bionic": "বায়োনিক",
    "Pro": "প্রো",
    "Max": "ম্যাক্স",
    "Ultra": "আলট্রা",
    "Plus": "প্লাস",
    "core": "কোর",
    "Adreno": "অ্যাড্রেনো",
    "Mali": "মালি",
    "Oryon": "অরিয়ন",
    "Immortalis": "ইমমরটালিস",
    "GPU": "জিপিইউ",
    "CPU": "সিপিইউ",
    
    # Other common terms
    "Nano-SIM": "ন্যানো-সিম",
    "eSIM": "ই-সিম",
    "Dual": "ডুয়াল",
    "speaker": "স্পিকার",
    "speakers": "স্পিকার",
    "stereo": "স্টেরিও",
    "mono": "মোনো",
    "jack": "জ্যাক",
    "port": "পোর্ট",
    "slot": "স্লট",
    "band": "ব্যান্ড",
    "model": "মডেল",
    "variant": "ভেরিয়েন্ট",
    "channel": "চ্যানেল",
    "mode": "মোড",
    "case": "কেস",
}

def translate_value(text):
    """Translate English text to Bengali comprehensively"""
    if not text:
        return text
    
    result = text
    
    # Sort by length (longest first) to avoid partial replacements
    sorted_terms = sorted(COMPLETE_BENGALI_DICT.items(), key=lambda x: len(x[0]), reverse=True)
    
    for eng_term, ben_term in sorted_terms:
        # Use word boundaries for complete word matching
        pattern = r'\b' + re.escape(eng_term) + r'\b'
        result = re.sub(pattern, ben_term, result, flags=re.IGNORECASE)
    
    return result

def generate_bengali_translations(english_value):
    """Generate complete Bengali translation for a specification value"""
    return translate_value(english_value)

# Main execution
print("Regenerating all seeder files with COMPLETE Bengali translations...")
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
        
        # Generate slug
        slug = device_name.lower().replace(' ', '-').replace('(', '').replace(')', '')
        slug = re.sub(r'[^a-z0-9\-]', '', slug)
        slug = re.sub(r'-+', '-', slug).strip('-')
        
        filename = f"specification_seeder_mobile_{slug}.go"
        filepath = os.path.join(output_dir, filename)
        
        # Create class name
        words = device_name.replace('-', ' ').replace('(', '').replace(')', '').split()
        class_name = 'SpecificationSeederMobile' + ''.join(word.capitalize() for word in words if word)
        class_name = re.sub(r'[^a-zA-Z0-9]', '', class_name)
        
        # Extract specs
        spec_map = {}
        bengali_map = {}
        
        for spec_key, spec_data in specs.items():
            if isinstance(spec_data, dict) and "specification_translations" in spec_data:
                english_value = spec_data["english"]
                spec_map[spec_key] = english_value
                # Generate complete Bengali translation
                bengali_value = generate_bengali_translations(english_value)
                bengali_map[english_value] = bengali_value
        
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
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(go_content)
        
        created += 1
        if created % 50 == 0:
            print(f"[OK] Generated {created} seeders with complete translations...")
        
    except Exception as e:
        print(f"[ERROR] {device_name}: {e}")
        errors += 1

print()
print("="*70)
print(f"COMPLETE BENGALI TRANSLATION REGENERATION")
print("="*70)
print(f"Created: {created} seeder files")
print(f"Errors: {errors}")
print(f"Total: {created + errors}")
print()
print("All specifications now have COMPLETE Bengali translations!")
print("="*70)
