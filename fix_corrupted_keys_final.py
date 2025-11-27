#!/usr/bin/env python3
"""Fix corrupted keys and remaining mixed translations"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# Keys that got corrupted - fix them back to English
KEY_FIXES = [
    # Resolution specifications
    ('"1080 x 2400 pixels (~393 পিপিআই ঘনত্ব)"', '"1080 x 2400 pixels (~393 ppi density)"'),
    ('"1080 x 2400 pixels (~399 পিপিআই ঘনত্ব)"', '"1080 x 2400 pixels (~399 ppi density)"'),
    ('"1080 x 2400 pixels (~405 পিপিআই ঘনত্ব)"', '"1080 x 2400 pixels (~405 ppi density)"'),
    ('"1080 x 2408 pixels (~400 পিপিআই ঘনত্ব)"', '"1080 x 2408 pixels (~400 ppi density)"'),
    ('"1080 x 2412 pixels (~394 পিপিআই ঘনত্ব)"', '"1080 x 2412 pixels (~394 ppi density)"'),
    ('"1344 x 2992 pixels (~486 পিপিআই ঘনত্ব)"', '"1344 x 2992 pixels (~486 ppi density)"'),
    ('"1440 x 3120 pixels (~516 পিপিআই ঘনত্ব)"', '"1440 x 3120 pixels (~516 ppi density)"'),
    ('"1440 x 3120 pixels, 19.5:9 ratio (~498 পিপিআই ঘনত্ব)"', '"1440 x 3120 pixels, 19.5:9 ratio (~498 ppi density)"'),
    ('"1440 x 3168 pixels (~510 পিপিআই ঘনত্ব)"', '"1440 x 3168 pixels (~510 ppi density)"'),
    ('"1440 x 3200 pixels (~522 পিপিআই ঘনত্ব)"', '"1440 x 3200 pixels (~522 ppi density)"'),
    ('"2556 x 1179 pixels (~460 পিপিআই ঘনত্ব)"', '"2556 x 1179 pixels (~460 ppi density)"'),
    ('"2556 x 1179 pixels (~461 পিপিআই ঘনত্ব)"', '"2556 x 1179 pixels (~461 ppi density)"'),
    ('"2796 x 1290 pixels (~456 পিপিআই ঘনত্ব)"', '"2796 x 1290 pixels (~456 ppi density)"'),
    ('"2796 x 1290 pixels (~460 পিপিআই ঘনত্ব)"', '"2796 x 1290 pixels (~460 ppi density)"'),
    
    # Dimensions with corrupted units
    ('"153.7 x 73.6 x 8.2 মিমি (6.05 x 2.90 x 0.32 ইঞ্চি)"', '"153.7 x 73.6 x 8.2 mm (6.05 x 2.90 x 0.32 in)"'),
    ('"161.8 x 73.9 x 8.4 মিমি (6.37 x 2.91 x 0.33 ইঞ্চি)"', '"161.8 x 73.9 x 8.4 mm (6.37 x 2.91 x 0.33 in)"'),
    ('"162.3 x 75.6 x 7.7 মিমি (6.39 x 2.98 x 0.30 ইঞ্চি)"', '"162.3 x 75.6 x 7.7 mm (6.39 x 2.98 x 0.30 in)"'),
    ('"162.8 x 74.5 x 8.3 মিমি (6.41 x 2.93 x 0.33 ইঞ্চি)"', '"162.8 x 74.5 x 8.3 mm (6.41 x 2.93 x 0.33 in)"'),
    ('"162.8 x 76.6 x 8.5 মিমি (6.41 x 3.02 x 0.33 ইঞ্চি)"', '"162.8 x 76.6 x 8.5 mm (6.41 x 3.02 x 0.33 in)"'),
    ('"162.8 x 77.6 x 8.2 মিমি (6.41 x 3.06 x 0.32 ইঞ্চি)"', '"162.8 x 77.6 x 8.2 mm (6.41 x 3.06 x 0.32 in)"'),
    ('"165.3 x 76.3 x 8.5 মিমি (6.51 x 3.00 x 0.33 ইঞ্চি)"', '"165.3 x 76.3 x 8.5 mm (6.51 x 3.00 x 0.33 in)"'),
    
    # Processor speeds
    ('"1x3.1 গিগাহার্জ + 3x2.6 গিগাহার্জ + 4x1.92 গিগাহার্জ"', '"1x3.1 GHz + 3x2.6 GHz + 4x1.92 GHz"'),
    ('"1x3.25 গিগাহার্জ + 3x2.85 গিগাহার্জ + 4x2.0 গিগাহার্জ"', '"1x3.25 GHz + 3x2.85 GHz + 4x2.0 GHz"'),
    ('"2x4.32 গিগাহার্জ + 6x3.53 গিগাহার্জ"', '"2x4.32 GHz + 6x3.53 GHz"'),
    ('"2x4.47 গিগাহার্জ (Performance) + 6x3.53 গিগাহার্জ (Efficiency)"', '"2x4.47 GHz (Performance) + 6x3.53 GHz (Efficiency)"'),
    
    # Camera specs with corrupted mm
    ('"12 MP, f/2.2, 26মিমি (প্রশস্ত), 1.12µm, ডুয়াল pixel PDAF"', '"12 MP, f/2.2, 26mm (wide), 1.12µm, dual pixel PDAF"'),
    ('"12 MP, f/2.2, 26মিমি (প্রশস্ত), ডুয়াল Pixel PDAF"', '"12 MP, f/2.2, 26mm (wide), Dual Pixel PDAF"'),
    
    # Other mm replacements
    ('"200MP (wide, f/1.7) + 10MP (টেলিফটো 3x, f/2.4) + 50MP (পেরিস্কোপ 5x, f/3.4) + 50MP (ultrawide, f/1.9)"', '"200MP (wide, f/1.7) + 10MP (telephoto 3x, f/2.4) + 50MP (periscope 5x, f/3.4) + 50MP (ultrawide, f/1.9)"'),
]

# Value fixes - these stay as is (already Bengali translations)
VALUE_FIXES = [
    (': "100W ওয়্যার্ড (100% এ 36 ম', ': "100W ওয়্যার্ড (100% এ 36 মিনিট)'),  # Was corrupted with "mইঞ্চি"
    (': "20W ওয়্যার্ড, 15W ম্যাগসেফ ওয়্যারলেস"', ': "20W ওয়্যার্ড, 15W ম্যাগসেফ ওয়্যারলেস"'),
]

def fix_file(filepath):
    """Fix corrupted keys and remaining mixed translations"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        # Fix keys first
        for old_key, new_key in KEY_FIXES:
            if old_key in content:
                content = content.replace(old_key, new_key)
                fixed_count += 1
        
        # Fix values
        for old_val, new_val in VALUE_FIXES:
            if old_val in content:
                content = content.replace(old_val, new_val)
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
    print("FIXING CORRUPTED KEYS AND REMAINING MIXED TRANSLATIONS")
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
    print("CORRUPTED KEY FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Key/Value fixes: {len(KEY_FIXES) + len(VALUE_FIXES)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
