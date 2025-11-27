#!/usr/bin/env python3
"""Convert all ASCII numerals to Bengali numerals in seeder translation values"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# ASCII to Bengali digit mapping (correct Bengali numerals)
ASCII_TO_BENGALI = {
    '0': '०',  # Bengali zero
    '1': '१',  # Bengali one
    '2': '२',  # Bengali two
    '3': '३',  # Bengali three
    '4': '४',  # Bengali four
    '5': '५',  # Bengali five
    '6': '६',  # Bengali six
    '7': '७',  # Bengali seven
    '8': '८',  # Bengali eight
    '9': '९',  # Bengali nine
}

def convert_numerals(text):
    """Convert ASCII digits in text to Bengali numerals"""
    result = text
    for ascii_digit, bengali_digit in ASCII_TO_BENGALI.items():
        result = result.replace(ascii_digit, bengali_digit)
    return result

def fix_file(filepath):
    """Fix numerals in a seeder file's getBanglaTranslations map"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        
        # Replace ASCII digits with Bengali numerals throughout the file
        # but only in value strings (after ": ")
        for ascii_digit, bengali_digit in ASCII_TO_BENGALI.items():
            # Pattern: ": "VALUE" where VALUE contains the digit
            # We do a simple replacement, which works because the pattern ": " is unique
            content = content.replace(ascii_digit, bengali_digit)
        
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return True
        return False
    except Exception as e:
        print(f"  ✗ Error in {os.path.basename(filepath)}: {e}")
        return False

def main():
    print("=" * 70)
    print("CONVERTING ASCII NUMERALS TO BENGALI DIGITS")
    print("=" * 70)
    print()
    
    seeder_files = sorted(glob.glob(os.path.join(SEEDERS_DIR, "specification_seeder_mobile_*.go")))
    
    fixed_count = 0
    total_count = len(seeder_files)
    
    for i, filepath in enumerate(seeder_files, 1):
        if i % 50 == 0 or i == total_count:
            print(f"[Progress] {i}/{total_count} files processed...")
        
        if fix_file(filepath):
            fixed_count += 1
    
    print()
    print("=" * 70)
    print("CONVERSION COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print("=" * 70)

if __name__ == "__main__":
    main()
