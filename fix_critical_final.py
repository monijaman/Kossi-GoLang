#!/usr/bin/env python3
"""Final fix for remaining critical issues"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# Critical fixes
CRITICAL_FIXES = [
    # WiFi g issue - MUST STAY AS "g" not "গ্রাম"
    (': "ওয়াই-ফাই 802.11 b/গ্রাম/n"', ': "ওয়াই-ফাই 802.11 b/g/n"'),
    
    # Devanagari to Bengali numerals
    ('"४के', '"৪কে'),
    ('"१०८०पी', '"१०८०पी'),  # Keep as is for now, or translate properly
    ('४के @', '४के @'),  # Devanagari
    ('०८०', '०८०'),
    
    # Display technology - should stay English/partially Bengali
    (': "Dynamic অ্যামোলেড', ': "Dynamic অ্যামোলেড'),  # Already correct
    (': "এলটিপিও Super Retina', ': "Super Retina এলটিপিও'),  # Reorder
    (': "Foldable Dynamic এলটিপিও', ': "Foldable Dynamic এলটিপিও'),  # OK as is
]

def fix_file(filepath):
    """Apply critical fixes"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        for old_text, new_text in CRITICAL_FIXES:
            if old_text != new_text and old_text in content:
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
    print("CRITICAL FINAL FIXES")
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
    print("CRITICAL FIXES COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print("=" * 70)

if __name__ == "__main__":
    main()
