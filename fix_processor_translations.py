#!/usr/bin/env python3
"""Fix all processor and chipset translations in seeder files"""

import os
import re
import glob

SEEDERS_DIR = "init-db/seeders"

# Complete dictionary of processor/chipset translations
PROCESSOR_TRANSLATIONS = {
    # Qualcomm Snapdragon with SM codes
    "Qualcomm SM8750-AC Snapdragon 8 Elite": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 Elite",
    "Qualcomm SM8750 Snapdragon 8 Elite": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 Elite",
    "Qualcomm SM8650-AC Snapdragon 8 Gen 3": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 Gen 3",
    "Qualcomm SM8650-AB Snapdragon 8 Gen 3": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 Gen 3",
    "Qualcomm SM8550-AC Snapdragon 8 Gen 2": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 Gen 2",
    "Qualcomm SM8550-AB Snapdragon 8 Gen 2": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 Gen 2",
    
    # Qualcomm Snapdragon without SM codes
    "Qualcomm Snapdragon 8 Gen 4": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 4",
    "Qualcomm Snapdragon 8 Gen 3": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 3",
    "Qualcomm Snapdragon 8 Gen 2": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 2",
    "Qualcomm Snapdragon 8 Gen 1": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 1",
    "Qualcomm Snapdragon 7 Gen 3": "কোয়ালকম স্ন্যাপড্রাগন 7 Gen 3",
    "Qualcomm Snapdragon 7 Gen 2": "কোয়ালকম স্ন্যাপড্রাগন 7 Gen 2",
    "Qualcomm Snapdragon 7s Gen 3": "কোয়ালকম স্ন্যাপড্রাগন 7s Gen 3",
    "Qualcomm Snapdragon 778G 5G": "কোয়ালকম স্ন্যাপড্রাগন 778G 5G",
    "Qualcomm Snapdragon 782G": "কোয়ালকম স্ন্যাপড্রাগন 782G",
    
    # MediaTek Dimensity
    "MediaTek Dimensity 6100+": "মিডিয়াটেক ডাইমেনশিটি 6100+",
    "MediaTek Dimensity 7025 Ultra": "মিডিয়াটেক ডাইমেনশিটি 7025 Ultra",
    "Mediatek Dimensity 6080": "মিডিয়াটেক ডাইমেনশিটি 6080",
    "Mediatek Dimensity 7025 Ultra": "মিডিয়াটেক ডাইমেনশিটি 7025 Ultra",
    
    # MediaTek Helio
    "MediaTek Helio G99": "মিডিয়াটেক হেলিও G99",
    "Mediatek Helio G36": "মিডিয়াটেক হেলিও G36",
    "Helio G36": "হেলিও G36",
    "Helio G99": "হেলিও G99",
    
    # Samsung Exynos
    "Exynos 2400e": "এক্সিনস 2400e",
    "Exynos 2400": "এক্সিনস 2400",
    "Exynos 2200": "এক্সিনস 2200",
    
    # Apple
    "Apple A19 Pro": "অ্যাপল A19 Pro",
    "Apple A19": "অ্যাপল A19",
    "Apple A18 Pro": "অ্যাপল A18 Pro",
    "Apple A18": "অ্যাপল A18",
    "Apple A17 Pro": "অ্যাপল A17 Pro",
    "Apple A16 Bionic": "অ্যাপল A16 Bionic",
    "Apple A15 Bionic": "অ্যাপল A15 Bionic",
    
    # Standalone brand translations (for cases where brand is separate)
    "Qualcomm": "কোয়ালকম",
    "MediaTek": "মিডিয়াটেক",
    "Mediatek": "মিডিয়াটেক",
    "Snapdragon": "স্ন্যাপড্রাগন",
    "Helio": "হেলিও",
    "Dimensity": "ডাইমেনশিটি",
    "Exynos": "এক্সিনস",
    "Bionic": "বায়োনিক",
    "Elite": "এলিট",
    "Pro": "প্রো",
    "Ultra": "আলট্রা",
}

def fix_file(filepath):
    """Fix processor/chipset translations in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        # Sort by length (longest first) to avoid partial replacements
        sorted_translations = sorted(PROCESSOR_TRANSLATIONS.items(), key=lambda x: len(x[0]), reverse=True)
        
        for old_text, new_text in sorted_translations:
            # Use word boundary regex for more precise matching
            pattern = r'"' + re.escape(old_text) + r'"\s*:\s*"'
            replacement = f'"{new_text}": "'
            
            if re.search(pattern, content):
                # Count how many times we'll replace
                matches = len(re.findall(pattern, content))
                content = re.sub(pattern, replacement, content)
                fixed_count += matches
        
        # Write back if changed
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return fixed_count > 0
        return False
    except Exception as e:
        print(f"  ✗ Error in {filepath}: {e}")
        return False

def main():
    print("=" * 70)
    print("FIXING PROCESSOR & CHIPSET TRANSLATIONS")
    print("=" * 70)
    
    seeder_files = sorted(glob.glob(os.path.join(SEEDERS_DIR, "specification_seeder_mobile_*.go")))
    
    fixed_count = 0
    total_count = len(seeder_files)
    
    for i, filepath in enumerate(seeder_files, 1):
        if i % 50 == 0 or i == total_count:
            print(f"\n[Progress] Processing {i}/{total_count} files...")
        
        if fix_file(filepath):
            fixed_count += 1
    
    print("\n" + "=" * 70)
    print("PROCESSOR & CHIPSET FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Translation terms applied: {len(PROCESSOR_TRANSLATIONS)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
