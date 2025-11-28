#!/usr/bin/env python3
"""Translate all remaining English words to Bangla in mobile seeder files"""

from pathlib import Path
import re

SEEDERS_DIR = Path("internal/infrastructure/database/seeders")

# Comprehensive translation map
TRANSLATIONS = [
    # Technical terms
    ('pixels', 'পিক্সেল'),
    ('pixel', 'পিক্সেল'),
    ('ppi', 'পিপিআই'),
    ('fps', 'এফপিএস'),
    ('colors', 'রঙ'),
    ('color', 'রঙ'),
    
    # Brand/OS names
    ('Android', 'অ্যান্ড্রয়েড'),
    ('iOS', 'আইওএস'),
    ('OxygenOS', 'অক্সিজেন ওএস'),
    ('ColorOS', 'কালার ওএস'),
    ('MIUI', 'এমআইইউআই'),
    ('One UI', 'ওয়ান ইউআই'),
    ('Realme UI', 'রিয়েলমি ইউআই'),
    ('HyperOS', 'হাইপার ওএস'),
    
    # Processor brands
    ('Qualcomm', 'কোয়ালকম'),
    ('Snapdragon', 'স্ন্যাপড্রাগন'),
    ('MediaTek', 'মিডিয়াটেক'),
    ('Mediatek', 'মিডিয়াটেক'),
    ('Dimensity', 'ডাইমেনসিটি'),
    ('Helio', 'হেলিও'),
    ('Apple', 'অ্যাপল'),
    ('Exynos', 'এক্সিনোস'),
    ('Tensor', 'টেনসর'),
    ('Bionic', 'বায়োনিক'),
    
    # GPU brands
    ('Adreno', 'অ্যাড্রেনো'),
    ('Mali', 'মালি'),
    ('Immortalis', 'ইমরটালিস'),
    ('PowerVR', 'পাওয়ারভিআর'),
    
    # Display technologies
    ('OLED', 'ওএলইডি'),
    ('AMOLED', 'অ্যামোলেড'),
    ('LTPO', 'এলটিপিও'),
    ('LCD', 'এলসিডি'),
    ('IPS', 'আইপিএস'),
    ('Super Retina', 'সুপার রেটিনা'),
    ('Dynamic', 'ডায়নামিক'),
    ('HDR', 'এইচডিআর'),
    ('HDR10', 'এইচডিআর১০'),
    ('Dolby Vision', 'ডলবি ভিশন'),
    ('ProMotion', 'প্রোমোশন'),
    
    # Glass/Protection brands
    ('Gorilla Glass', 'গরিলা গ্লাস'),
    ('Ceramic Guard', 'সিরামিক গার্ড'),
    ('Corning', 'কর্নিং'),
    ('Victus', 'ভিক্টাস'),
    ('Armor', 'আর্মার'),
    
    # Connectivity
    ('Wi-Fi', 'ওয়াই-ফাই'),
    ('WiFi', 'ওয়াই-ফাই'),
    ('Bluetooth', 'ব্লুটুথ'),
    ('NFC', 'এনএফসি'),
    ('USB', 'ইউএসবি'),
    ('Type-C', 'টাইপ-সি'),
    ('GPS', 'জিপিএস'),
    ('LTE', 'এলটিই'),
    ('GSM', 'জিএসএম'),
    ('CDMA', 'সিডিএমএ'),
    ('HSPA', 'এইচএসপিএ'),
    ('EVDO', 'ইভিডিও'),
    ('UWB', 'ইউডব্লিউবি'),
    
    # Camera/Audio terms
    ('MP', 'মেগাপিক্সেল'),
    ('Dual', 'ডুয়াল'),
    ('Triple', 'ট্রিপল'),
    ('Quad', 'কোয়াড'),
    ('Stereo', 'স্টেরিও'),
    ('Dolby Atmos', 'ডলবি অ্যাটমস'),
    ('Hi-Res', 'হাই-রেজ'),
    
    # Charging
    ('PD', 'পিডি'),
    ('Quick Charge', 'কুইক চার্জ'),
    ('Fast Charging', 'ফাস্ট চার্জিং'),
    ('Wireless', 'ওয়্যারলেস'),
    ('Reverse', 'রিভার্স'),
    ('MagSafe', 'ম্যাগসেফ'),
    ('Qi', 'কিউআই'),
    
    # Sensors
    ('Fingerprint', 'ফিঙ্গারপ্রিন্ট'),
    ('Face ID', 'ফেস আইডি'),
    ('Touch ID', 'টাচ আইডি'),
    ('Accelerometer', 'অ্যাক্সিলারোমিটার'),
    ('Gyro', 'জাইরো'),
    ('Gyroscope', 'জাইরোস্কোপ'),
    ('Proximity', 'প্রক্সিমিটি'),
    ('Compass', 'কম্পাস'),
    ('Barometer', 'ব্যারোমিটার'),
    ('Ultrasonic', 'আল্ট্রাসনিক'),
    
    # Color descriptors (already done but ensuring consistency)
    ('Monet', 'মোনেট'),
    ('Emerald', 'এমারেল্ড'),
    ('Obsidian', 'অবসিডিয়ান'),
    ('Porcelain', 'পোর্সেলিন'),
    ('Hazel', 'হ্যাজেল'),
    ('Rose', 'রোজ'),
    ('Mint', 'মিন্ট'),
    ('Bay', 'বে'),
    ('Sky', 'স্কাই'),
    ('Titan', 'টাইটান'),
    ('Eternal', 'ইটার্নাল'),
    ('Desert', 'মরুভূমি'),
    ('Natural', 'ন্যাচারাল'),
    ('Awesome', 'অসাম'),
    
    # Storage/Memory units (keeping abbreviations but can translate)
    ('ROM', 'রম'),
    ('RAM', 'র‍্যাম'),
    ('eMMC', 'ইএমএমসি'),
    ('UFS', 'ইউএফএস'),
    
    # Other technical terms
    ('IP', 'আইপি'),
    ('Nano-SIM', 'ন্যানো-সিম'),
    ('eSIM', 'ইসিম'),
    ('Dual SIM', 'ডুয়াল সিম'),
    ('microSD', 'মাইক্রোএসডি'),
    ('Li-Po', 'লি-পো'),
    ('Li-Ion', 'লি-আয়ন'),
    ('Side', 'সাইড'),
    ('Under display', 'ডিসপ্লের নিচে'),
    ('In-display', 'ইন-ডিসপ্লে'),
]

def translate_file(filepath):
    """Apply all translations to a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        changes = 0
        
        # Only translate within translation value strings (right side of colon in map)
        lines = content.split('\n')
        new_lines = []
        in_translations_map = False
        
        for line in lines:
            # Detect if we're in the getBanglaTranslations map
            if 'getBanglaTranslations' in line or 'return map[string]string{' in line:
                in_translations_map = True
                new_lines.append(line)
                continue
            
            # Check if we're exiting the map
            if in_translations_map and line.strip() == '}':
                in_translations_map = False
                new_lines.append(line)
                continue
            
            # If we're in the map and this looks like a translation line
            if in_translations_map and '":' in line and '"' in line:
                # Split at the first colon to separate key from value
                parts = line.split(':', 1)
                if len(parts) == 2:
                    key_part = parts[0]
                    value_part = parts[1]
                    original_value = value_part
                    
                    # Apply all translations to the value part only
                    for english, bangla in TRANSLATIONS:
                        # Use word boundaries to avoid partial replacements
                        value_part = re.sub(r'\b' + re.escape(english) + r'\b', bangla, value_part)
                    
                    if value_part != original_value:
                        changes += 1
                    
                    new_lines.append(key_part + ':' + value_part)
                else:
                    new_lines.append(line)
            else:
                new_lines.append(line)
        
        new_content = '\n'.join(new_lines)
        
        if new_content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(new_content)
            return changes
        
        return 0
        
    except Exception as e:
        print(f"  ✗ Error in {filepath.name}: {e}")
        return 0

def main():
    print("=" * 80)
    print("TRANSLATING ALL ENGLISH TECHNICAL TERMS TO BANGLA")
    print("=" * 80)
    
    if not SEEDERS_DIR.exists():
        print(f"Error: Directory not found: {SEEDERS_DIR}")
        return
    
    seeder_files = sorted(SEEDERS_DIR.glob("specification_seeder_mobile_*.go"))
    
    if not seeder_files:
        print(f"No mobile seeder files found")
        return
    
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    modified_count = 0
    total_changes = 0
    
    for i, filepath in enumerate(seeder_files, 1):
        changes = translate_file(filepath)
        if changes > 0:
            modified_count += 1
            total_changes += changes
            print(f"[{i}/{len(seeder_files)}] ✓ {filepath.name}: {changes} translations")
        elif i % 50 == 0:
            print(f"[{i}/{len(seeder_files)}] Processed...")
    
    print("\n" + "=" * 80)
    print("TRANSLATION COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {modified_count}")
    print(f"Total translations applied: {total_changes}")
    print("=" * 80)

if __name__ == "__main__":
    main()
