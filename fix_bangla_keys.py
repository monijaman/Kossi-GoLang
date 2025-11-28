#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Fix keys that have Bangla text in them - these should be English
"""

import os
import re
from pathlib import Path

# Define key corrections (Bangla key -> English key)
KEY_CORRECTIONS = [
    ('"Yes (বাজার নির্ভরশীল)":', '"Yes (market dependent)":'),
    ('"Yes (region নির্ভরশীল)":', '"Yes (region dependent)":'),
    ('"সুপার Silver, মেগা Blue, আল্ট্রা কমলা":', '"Super Silver, Mega Blue, Ultra Orange":'),
    ('"128 GB / 256 GB / 512 GB / 1 টিবি":', '"128 GB / 256 GB / 512 GB / 1 TB":'),
    ('"256 GB / 512 GB / 1 টিবি":', '"256 GB / 512 GB / 1 TB":'),
    ('"iOS 17, আপগ্রেডযোগ্য":', '"iOS 17, upgradable":'),
    ('"LTPO সুপার Retina XDR OLED, 120Hz":', '"LTPO Super Retina XDR OLED, 120Hz":'),
    ('"Android 15, আপগ্রেডযোগ্য Android 22":', '"Android 15, upgradable Android 22":'),
    ('"Black, White, Pink, Teal, আল্ট্রাmarine":', '"Black, White, Pink, Teal, Ultramarine":'),
    ('"সুপার Retina XDR OLED, HDR10, Dolby Vision, 2000 nits (HBM)":', '"Super Retina XDR OLED, HDR10, Dolby Vision, 2000 nits (HBM)":'),
    ('"সুপার Retina XDR LTPO OLED, 120Hz, HDR10, Dolby Vision":', '"Super Retina XDR LTPO OLED, 120Hz, HDR10, Dolby Vision":'),
]

def fix_bangla_keys(file_path):
    """Fix keys that contain Bangla text"""
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    original_content = content
    corrections_made = 0
    
    for bangla_key, english_key in KEY_CORRECTIONS:
        if bangla_key in content:
            content = content.replace(bangla_key, english_key)
            corrections_made += 1
    
    if content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        return corrections_made
    
    return 0

def main():
    """Main execution function"""
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    # Find all mobile seeder files
    seeder_files = sorted(seeders_dir.glob("specification_seeder_mobile_*.go"))
    
    print("=" * 80)
    print("FIXING BANGLA TEXT IN KEYS")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    
    total_files_modified = 0
    total_corrections = 0
    
    for idx, file_path in enumerate(seeder_files, 1):
        corrections = fix_bangla_keys(file_path)
        
        if corrections > 0:
            total_files_modified += 1
            total_corrections += corrections
            print(f"[{idx}/{len(seeder_files)}] OK {file_path.name}: {corrections} keys fixed")
    
    print("\n" + "=" * 80)
    print("FIX COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {total_files_modified}")
    print(f"Total keys fixed: {total_corrections}")
    print("=" * 80)

if __name__ == "__main__":
    main()
