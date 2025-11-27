#!/usr/bin/env python3
"""Fix corrupted keys - keys should be in English only"""

import os
import glob
import re

SEEDERS_DIR = "init-db/seeders"

# Map of corrupted keys (Bengali) to correct keys (English)
CORRUPTED_KEYS = {
    '"হেলিও G99"': '"Helio G99"',
    '"হেলিও G36"': '"Helio G36"',
    '"মিডিয়াটেক হেলিও G99"': '"MediaTek Helio G99"',
    '"মিডিয়াটেক ডাইমেনশিটি 6100+"': '"MediaTek Dimensity 6100+"',
    '"এক্সিনস 2400e"': '"Exynos 2400e"',
    '"কোয়ালকম স্ন্যাপড্রাগন 8 Gen 4"': '"Qualcomm Snapdragon 8 Gen 4"',
    '"অ্যাপল A17 Pro"': '"Apple A17 Pro"',
    '"অ্যাপল A18 Pro"': '"Apple A18 Pro"',
    '"অ্যাপল A19 Pro"': '"Apple A19 Pro"',
    '"অ্যাপল A18"': '"Apple A18"',
    '"অ্যাপল A19"': '"Apple A19"',
    '"অ্যাপল A16 Bionic"': '"Apple A16 Bionic"',
    '"অ্যাপল A15 Bionic"': '"Apple A15 Bionic"',
}

def fix_file(filepath):
    """Fix corrupted keys in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        for corrupted_key, correct_key in CORRUPTED_KEYS.items():
            # Only fix keys (left side of colon), not values
            # Pattern: key followed by colon and quote
            pattern = corrupted_key + r':\s*"'
            replacement = correct_key + ': "'
            
            if re.search(pattern, content):
                content = re.sub(pattern, replacement, content)
                fixed_count += 1
        
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return fixed_count > 0
        return False
    except Exception as e:
        print(f"  ✗ Error: {e}")
        return False

def main():
    print("=" * 70)
    print("FIXING CORRUPTED KEYS (Bengali keys to English)")
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
    print("CORRUPTED KEY FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Corrupted keys fixed: {len(CORRUPTED_KEYS)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
