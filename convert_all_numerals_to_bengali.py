#!/usr/bin/env python3
"""Convert all ASCII numerals to Bengali numerals in seeder translation values"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# ASCII to Bengali digit mapping
ASCII_TO_BENGALI = {
    '0': '০',
    '1': '১',
    '2': '२',
    '3': '३',
    '4': '४',
    '5': '५',
    '6': '६',
    '7': '७',
    '8': '८',
    '9': '९',
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
            lines = f.readlines()
        
        modified = False
        new_lines = []
        
        for line in lines:
            # Only process lines that are translation map entries (contain ": ")
            if '": "' in line and not line.strip().startswith("//"):
                # Extract the key and value parts
                if line.count('": "') == 1:
                    # Simple case: one key-value pair per line
                    parts = line.split('": "', 1)
                    if len(parts) == 2:
                        key_part = parts[0]
                        value_part = parts[1]
                        
                        # Check if value contains ASCII digits
                        if any(d in value_part for d in '0123456789'):
                            # Convert numerals in the value part only
                            value_part_converted = convert_numerals(value_part)
                            if value_part_converted != value_part:
                                line = key_part + '": "' + value_part_converted
                                modified = True
            
            new_lines.append(line)
        
        if modified:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.writelines(new_lines)
            return True
        return False
    except Exception as e:
        print(f"  ✗ Error in {os.path.basename(filepath)}: {e}")
        return False

def main():
    print("=" * 70)
    print("CONVERTING ALL ASCII NUMERALS TO BENGALI DIGITS")
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
