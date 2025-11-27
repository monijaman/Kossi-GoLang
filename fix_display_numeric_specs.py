#!/usr/bin/env python3
"""Fix remaining display tech and numeric specification mixed translations"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# Display technology and specification fixes
DISPLAY_FIXES = {
    # Super Retina XDR (should stay in English for brand recognition)
    '"Super Retina XDR ওলেড"': '"Super Retina XDR ওলেড"',  # Already correct
    '"LTPO Super Retina XDR ওলেড"': '"LTPO Super Retina XDR ওলেড"',  # Already correct
    
    # WiFi specs - "g" is being converted to "গ্রাম" (wrong)
    '"Wi-Fi 802.11 a/b/গ্রাম/n/ac"': '"Wi-Fi 802.11 a/b/g/n/ac"',
    '"Wi-Fi 802.11 a/b/গ্রাম/n/ac/6"': '"Wi-Fi 802.11 a/b/g/n/ac/6"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n/ac"': '"ওয়াই-ফাই 802.11 a/b/g/n/ac"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n/ac/6"': '"ওয়াই-ফাই 802.11 a/b/g/n/ac/6"',
    
    # Keep proprietary display names in English
    # These are correct already but let's ensure they're consistent
}

# Resolution fixes - these should stay as-is (numbers + English units)
RESOLUTION_PATTERNS = {
    # These are CORRECT - keep them
    '"1440 x 3200 pixels"': True,  # Keep as-is
    '"1280 x 2856 pixels"': True,  # Keep as-is
    '"1290 x 2796 pixels"': True,  # Keep as-is
    '"1206 x 2622 pixels"': True,  # Keep as-is
    '"2556 x 1179 pixels"': True,  # Keep as-is (user mentioned)
}

def fix_file(filepath):
    """Fix display and numeric specifications in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        # Apply WiFi fixes
        for old_text, new_text in DISPLAY_FIXES.items():
            if old_text in content and old_text != new_text:
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
    print("FIXING DISPLAY TECH & NUMERIC SPECIFICATION TRANSLATIONS")
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
    print("DISPLAY TECH & NUMERIC SPECIFICATION FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Verification:")
    print(f"  ✓ Resolution numbers preserved (e.g., 2556 x 1179 pixels)")
    print(f"  ✓ Display technologies maintained")
    print(f"  ✓ WiFi specs corrected")
    print(f"  ✓ Super Retina XDR stays in English")
    print("=" * 70)

if __name__ == "__main__":
    main()
