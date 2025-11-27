#!/usr/bin/env python3
"""Fix all processor and chipset translations - Direct string replacement"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# Complete fixes - old (incorrect) to new (correct)
FIXES = [
    # Full chipset strings with SM codes
    ('"Qualcomm SM8750-AC Snapdragon 8 Elite": "Qualcomm SM8750-AC স্ন্যাপড্রাগন 8 Elite"',
     '"Qualcomm SM8750-AC Snapdragon 8 Elite": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 Elite"'),
    
    ('"Qualcomm SM8750 Snapdragon 8 Elite": "Qualcomm SM8750 স্ন্যাপড্রাগন 8 Elite"',
     '"Qualcomm SM8750 Snapdragon 8 Elite": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 Elite"'),
    
    ('"Qualcomm SM8650-AC Snapdragon 8 Gen 3": "Qualcomm SM8650-AC স্ন্যাপড্রাগন 8 Gen 3"',
     '"Qualcomm SM8650-AC Snapdragon 8 Gen 3": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 Gen 3"'),
    
    ('"Qualcomm SM8650-AB Snapdragon 8 Gen 3": "Qualcomm SM8650-AB স্ন্যাপড্রাগন 8 Gen 3"',
     '"Qualcomm SM8650-AB Snapdragon 8 Gen 3": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 Gen 3"'),
    
    ('"Qualcomm SM8550-AC Snapdragon 8 Gen 2": "Qualcomm SM8550-AC স্ন্যাপড্রাগন 8 Gen 2"',
     '"Qualcomm SM8550-AC Snapdragon 8 Gen 2": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 Gen 2"'),
    
    ('"Qualcomm SM8550-AB Snapdragon 8 Gen 2": "Qualcomm SM8550-AB স্ন্যাপড্রাগন 8 Gen 2"',
     '"Qualcomm SM8550-AB Snapdragon 8 Gen 2": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 Gen 2"'),
     
    # Full chipset strings with parentheses
    ('"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "Qualcomm SM8750-AC স্ন্যাপড্রাগন 8 Elite (3 nm)"',
     '"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 Elite (3 nm)"'),
    
    ('"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "Qualcomm SM8750 স্ন্যাপড্রাগন 8 Elite (3 nm)"',
     '"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 Elite (3 nm)"'),
    
    ('"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM8650-AC স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"',
     '"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"'),
    
    ('"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "Qualcomm SM8650-AB স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"',
     '"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"'),
    
    ('"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM8550-AC স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"',
     '"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"'),
    
    ('"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "Qualcomm SM8550-AB স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"',
     '"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"'),
    
    # MediaTek with partial translations
    ('"MediaTek Dimensity 6100+": "মিডিয়াটেক Dimensity 6100+"',
     '"MediaTek Dimensity 6100+": "মিডিয়াটেক ডাইমেনশিটি 6100+"'),
    
    ('"Mediatek Helio G36 (12 nm)": "মিডিয়াটেক হেলিও G36 (12 nm)"',
     '"Mediatek Helio G36 (12 nm)": "মিডিয়াটেক হেলিও G36 (12 nm)"'),
]

def fix_file(filepath):
    """Fix translations in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        for old_text, new_text in FIXES:
            if old_text in content:
                content = content.replace(old_text, new_text)
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
    print("FIXING REMAINING PROCESSOR TRANSLATIONS")
    print("=" * 70)
    
    seeder_files = sorted(glob.glob(os.path.join(SEEDERS_DIR, "specification_seeder_mobile_*.go")))
    
    fixed_count = 0
    total_count = len(seeder_files)
    total_fixes = 0
    
    for i, filepath in enumerate(seeder_files, 1):
        if i % 50 == 0 or i == total_count:
            print(f"\n[Progress] Processing {i}/{total_count} files...")
        
        if fix_file(filepath):
            fixed_count += 1
    
    print("\n" + "=" * 70)
    print("REMAINING TRANSLATION FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Total fixes applied: {len(FIXES)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
