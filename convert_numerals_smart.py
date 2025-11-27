#!/usr/bin/env python3
"""Convert ASCII numerals to Bengali numerals in translation VALUES only"""

import os
import re
import glob

SEEDERS_DIR = "init-db/seeders"

# ASCII to Bengali digit mapping
ASCII_TO_BENGALI = {
    '0': '०', '1': '१', '2': '२', '3': '३', '4': '४',
    '5': '५', '6': '६', '7': '७', '8': '८', '9': '९',
}

def convert_digits(text):
    """Convert ASCII digits to Bengali numerals"""
    for ascii_digit, bengali_digit in ASCII_TO_BENGALI.items():
        text = text.replace(ascii_digit, bengali_digit)
    return text

def fix_file(filepath):
    """Fix numerals in getBanglaTranslations map values only"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            lines = f.readlines()
        
        original_lines = lines.copy()
        in_bangla_map = False
        
        for i, line in enumerate(lines):
            # Detect start of getBanglaTranslations function
            if 'func (s *SpecificationSeeder' in line and 'getBanglaTranslations' in line:
                in_bangla_map = True
            
            # Detect end of map (closing brace followed by return or end)
            if in_bangla_map and line.strip() == '}':
                in_bangla_map = False
            
            # Process lines inside the map
            if in_bangla_map and '":' in line and '"' in line:
                # Pattern: "KEY": "VALUE",
                # We want to convert digits in VALUE only
                match = re.match(r'(\s*"[^"]*":\s*")(.*?)(".*)', line)
                if match:
                    prefix = match.group(1)
                    value = match.group(2)
                    suffix = match.group(3)
                    converted_value = convert_digits(value)
                    lines[i] = prefix + converted_value + suffix + '\n'
        
        if lines != original_lines:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.writelines(lines)
            return True
        return False
    except Exception as e:
        print(f"  ✗ Error in {os.path.basename(filepath)}: {e}")
        return False

def main():
    print("=" * 70)
    print("CONVERTING ASCII NUMERALS TO BENGALI (Values Only)")
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
