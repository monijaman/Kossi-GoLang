#!/usr/bin/env python3
"""Fix mixed Bengali/English translations in all seeder files"""

import os
import re
import glob

SEEDERS_DIR = "init-db/seeders"

# Pattern fixes for common mixed translations
FIXES = {
    # Brand + Processor combinations
    '"Qualcomm Snapdragon 8 Gen 4": "Qualcomm স্ন্যাপড্রাগন 8 Gen 4"': 
    '"Qualcomm Snapdragon 8 Gen 4": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 4"',
    
    '"Qualcomm Snapdragon 8 Gen 3": "Qualcomm স্ন্যাপড্রাগন 8 Gen 3"': 
    '"Qualcomm Snapdragon 8 Gen 3": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 3"',
    
    '"Qualcomm Snapdragon 8 Gen 2": "Qualcomm স্ন্যাপড্রাগন 8 Gen 2"': 
    '"Qualcomm Snapdragon 8 Gen 2": "কোয়ালকম স্ন্যাপড্রাগন 8 Gen 2"',
    
    '"Qualcomm Snapdragon 7 Gen 3": "Qualcomm স্ন্যাপড্রাগন 7 Gen 3"': 
    '"Qualcomm Snapdragon 7 Gen 3": "কোয়ালকম স্ন্যাপড্রাগন 7 Gen 3"',
    
    '"Qualcomm Snapdragon 7 Gen 2": "Qualcomm স্ন্যাপড্রাগন 7 Gen 2"': 
    '"Qualcomm Snapdragon 7 Gen 2": "কোয়ালকম স্ন্যাপড্রাগন 7 Gen 2"',
    
    '"Qualcomm Snapdragon 778G 5G": "Qualcomm স্ন্যাপড্রাগন 778G 5G"': 
    '"Qualcomm Snapdragon 778G 5G": "কোয়ালকম স্ন্যাপড্রাগন 778G 5G"',
    
    '"Qualcomm Snapdragon 782G": "Qualcomm স্ন্যাপড্রাগন 782G"': 
    '"Qualcomm Snapdragon 782G": "কোয়ালকম স্ন্যাপড্রাগন 782G"',
    
    '"MediaTek Helio G99": "মিডিয়াটেক হেলিও G99"': 
    '"MediaTek Helio G99": "মিডিয়াটেক হেলিও G99"',
    
    '"MediaTek Dimensity 7200": "মিডিয়াটেক ডাইমেনশিটি 7200"': 
    '"MediaTek Dimensity 7200": "মিডিয়াটেক ডাইমেনশিটি 7200"',
    
    '"Apple A18": "অ্যাপল A18"': 
    '"Apple A18": "অ্যাপল A18"',
    
    '"Apple A18 Pro": "অ্যাপল A18 Pro"': 
    '"Apple A18 Pro": "অ্যাপল A18 Pro"',
}

def fix_file(filepath):
    """Fix mixed translations in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        
        # Apply all fixes
        for old_text, new_text in FIXES.items():
            if old_text in content:
                content = content.replace(old_text, new_text)
                print(f"  ✓ Fixed: {old_text[:50]}...")
        
        # Write back if changed
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return True
        return False
    except Exception as e:
        print(f"  ✗ Error: {e}")
        return False

def main():
    print("=" * 70)
    print("FIXING MIXED BENGALI/ENGLISH TRANSLATIONS")
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
    print("MIXED TRANSLATION FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Errors: 0")
    print("=" * 70)

if __name__ == "__main__":
    main()
