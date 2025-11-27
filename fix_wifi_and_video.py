#!/usr/bin/env python3
"""Fix WiFi 'gram' corruption and Devanagari video format corruption"""

import os
import glob
import re

SEEDERS_DIR = "init-db/seeders"

def fix_file(filepath):
    """Fix WiFi and video format issues"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixes = 0
        
        # Fix 1: WiFi "b/গ্রাম/n" → "b/g/n"
        if "b/গ্রাম/n" in content:
            content = content.replace(': "ওয়াই-ফাই 802.11 b/গ্রাম/n"', ': "ওয়াই-ফাই 802.11 b/g/n"')
            fixes += 1
        
        # Fix 2: Devanagari numerals to Bengali numerals
        # "०४के" → "04K" or proper Bengali
        devanagari_fixes = {
            '०': '0',  # Devanagari zero
            '१': '1',  # Devanagari one
            '२': '2',
            '३': '3',
            '४': '4',  # This one appears in "४के"
            '५': '5',
            '६': '6',
            '७': '7',
            '८': '8',
            '९': '9',
            'के': 'K',  # Devanagari KA to English K
            'पी': 'P',  # Devanagari PA to English P
        }
        
        for deva_char, replacement in devanagari_fixes.items():
            if deva_char in content:
                content = content.replace(deva_char, replacement)
                fixes += 1
        
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return fixes > 0
        return False
    except Exception as e:
        print(f"  ✗ Error in {os.path.basename(filepath)}: {e}")
        return False

def main():
    print("=" * 70)
    print("FIXING WiFi 'gram' AND VIDEO FORMAT CORRUPTION")
    print("=" * 70)
    
    seeder_files = sorted(glob.glob(os.path.join(SEEDERS_DIR, "specification_seeder_mobile_*.go")))
    
    fixed_count = 0
    total_count = len(seeder_files)
    
    for i, filepath in enumerate(seeder_files, 1):
        if fix_file(filepath):
            fixed_count += 1
            print(f"  ✓ {os.path.basename(filepath)}")
    
    print("\n" + "=" * 70)
    print("FIXES COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files with fixes applied: {fixed_count}")
    print("=" * 70)

if __name__ == "__main__":
    main()
