#!/usr/bin/env python3
"""Fix all remaining mixed English/Bengali translations - Direct replacement"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# Direct string replacements - from mixed to fully Bengali
FIXES = [
    # Camera positions
    ('(Cover)', '(কভার)'),
    ('(Inner)', '(অভ্যন্তরীণ)'),
    ('(Under Display)', '(ডিসপ্লে অধীন)'),
    ('(Main)', '(মূল)'),
    
    # Camera types
    ('(wide)', '(প্রশস্ত)'),
    ('(ultrawide)', '(আল্ট্রাওয়াইড)'),
    ('(ultra-wide)', '(আল্ট্রা-প্রশস্ত)'),
    ('(telephoto', '(টেলিফটো'),
    ('(periscope', '(পেরিস্কোপ'),
    
    # SAR - head and body
    ('(head), ', '(মাথা), '),
    ('(body)', '(শরীর)'),
    
    # Clock speeds
    ('(Performance)', '(পারফরম্যান্স)'),
    ('(Efficiency)', '(দক্ষতা)'),
    
    # Video recording
    ('1080p', '१०८०पी'),
    ('2K', '२के'),
    ('4K', '४के'),
    ('8K', '८के'),
    ('fps)', 'fps)'),
    
    # Zoom and optical
    ('optical', 'অপটিক্যাল'),
    ('Space Zoom', 'Space জুম'),
    ('crop)', 'ক্রপ)'),
    ('via crop', 'মাধ্যমে ক্রপ'),
    
    # Display specifications
    ('ppi density', 'পিপিআই ঘনত্ব'),
    ('ratio)', 'অনুপাত)'),
    ('mm (', 'মিমি ('),
    ('in)', 'ইঞ্চি)'),
    
    # Charging
    ('MagSafe wireless', 'ম্যাগসেফ ওয়্যারলেস'),
    
    # Material
    ('(glass)', '(গ্লাস)'),
    ('(leather)', '(চামড়া)'),
    
    # Common modifiers in camera specs
    ('PDAF', 'PDAF'),
    ('Auto-HDR', 'স্বয়ংক্রিয়-এইচডিআর'),
    ('f/1.7', 'f/1.7'),
    ('f/2.0', 'f/2.0'),
    ('f/2.2', 'f/2.2'),
    ('f/2.4', 'f/2.4'),
    ('26mm', '26mm'),
    ('21mm', '21mm'),
    ('17mm', '17mm'),
    ('1.12µm', '1.12µm'),
    ('dual pixel', 'ডুয়াল pixel'),
    ('Dual Pixel', 'ডুয়াল Pixel'),
]

def fix_file(filepath):
    """Fix mixed translations in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        for old_text, new_text in FIXES:
            # Skip if already correct
            if old_text == new_text:
                continue
            
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
    print("FIXING ALL REMAINING MIXED TRANSLATIONS")
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
    print("COMPREHENSIVE FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Fixes applied: {len(FIXES)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
