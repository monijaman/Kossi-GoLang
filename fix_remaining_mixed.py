#!/usr/bin/env python3
"""Fix all remaining mixed English/Bengali translations"""

import os
import glob
import re

SEEDERS_DIR = "init-db/seeders"

# Dictionary of English terms that should be translated in specifications
TERM_TRANSLATIONS = {
    # Camera positions
    "(Cover)": "(কভার)",
    "(Inner)": "(অভ্যন্তরীণ)",
    "(Under Display)": "(ডিসপ্লে অধীন)",
    
    # Camera types/positions
    "(wide)": "(প্রশস্ত)",
    "(ultrawide)": "(আল্ট্রাওয়াইড)",
    "(telephoto": "(টেলিফটো",
    "(periscope": "(পেরিস্কোপ",
    "(ultra-wide)": "(আল্ট্রা-প্রশস্ত)",
    "(Main)": "(মূল)",
    
    # SAR/Radiation
    "(head)": "(মাথা)",
    "(body)": "(শরীর)",
    
    # Clock speeds
    "(Performance)": "(পারফরম্যান্স)",
    "(Efficiency)": "(দক্ষতা)",
    
    # Video resolution
    "1080p": "১০৮०पी",
    "2K": "२के",
    "4K": "४के",
    "8K": "८के",
    
    # Zoom
    "optical": "অপটিক্যাল",
    "Space": "স্পেস",
    "crop": "ক্রপ",
    "via": "মাধ্যমে",
    "Auto-": "স্বয়ংক্রিয়-",
    
    # Display/Camera settings
    "density": "ঘনত্ব",
    "ratio": "অনুপাত",
    "mm)": "মিমি)",  # Unit fix
    "in)": "ইঞ্চি)",
    
    # Charging types
    "MagSafe": "ম্যাগসেফ",
    "wireless": "ওয়্যারলেস",
    
    # Material
    "(glass)": "(গ্লাস)",
    "(leather)": "(চামড়া)",
    
    # Resolution density
    "ppi density": "পিপিআই ঘনত্ব",
}

def fix_file(filepath):
    """Fix mixed translations in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        # Sort by length (longest first) to avoid partial replacements
        sorted_terms = sorted(TERM_TRANSLATIONS.items(), key=lambda x: len(x[0]), reverse=True)
        
        for english_term, bengali_term in sorted_terms:
            # Only replace in translation values, not keys
            # Pattern: after a colon, in quotes
            if f'": "{bengali_term}"' not in content and english_term in content:
                # Check if it's in a translation value (after ": ")
                pattern = f'": "([^"]*{re.escape(english_term)}[^"]*)"'
                if re.search(pattern, content):
                    content = re.sub(pattern, lambda m: f'": "{m.group(1).replace(english_term, bengali_term)}"', content)
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
    print("FIXING ALL REMAINING MIXED ENGLISH/BENGALI TRANSLATIONS")
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
    print("COMPREHENSIVE MIXED TRANSLATION FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Terms fixed: {len(TERM_TRANSLATIONS)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
