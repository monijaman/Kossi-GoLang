#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Translate final remaining English terms to Bangla in mobile seeder files
ONLY WITHIN TRANSLATION MAP VALUES - NOT in code, comments, or identifiers
"""

import os
import re
from pathlib import Path

# Define final translation mappings
FINAL_TRANSLATIONS = [
    # Conjunctions and common words
    (' or ', ' বা '),
    
    # Material terms
    ('silicone', 'সিলিকন'),
    ('polymer', 'পলিমার'),
    ('eco leather', 'ইকো লেদার'),
    ('ceramic', 'সিরামিক'),
    
    # Status terms
    ('upgradable to', 'আপগ্রেডযোগ্য'),
    ('upgradable', 'আপগ্রেডযোগ্য'),
    ('newer', 'নতুন'),
    ('dependent', 'নির্ভরশীল'),
    ('market', 'বাজার'),
    
    # Color and descriptive terms
    ('Meadow', 'মিডো'),
    ('Gleaming', 'ঝকঝকে'),
    ('Funtouch', 'ফানটাচ'),
    ('Agate', 'অ্যাগেট'),
    ('Orange', 'কমলা'),
    ('Moonlight', 'মুনলাইট'),
    ('Twilight', 'টোয়াইলাইট'),
    ('Mint', 'মিন্ট'),
    ('Midnight', 'মিডনাইট'),
    ('Alpine', 'আলপাইন'),
    ('Astral', 'অ্যাস্ট্রাল'),
    ('Horizon', 'হরাইজন'),
    ('Martian', 'মার্শিয়ান'),
    ('Super', 'সুপার'),
    ('Mega', 'মেগা'),
    ('Ultra', 'আল্ট্রা'),
    
    # Other terms
    ('hybrid', 'হাইব্রিড'),
    ('Hybrid', 'হাইব্রিড'),
    ('edition', 'এডিশন'),
    ('Guard', 'গার্ড'),
    
    # Storage unit
    (' TB', ' টিবি'),
]

def translate_in_map_values_only(file_path):
    """Apply translations ONLY to values in getBanglaTranslations map"""
    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()
    
    original_content = ''.join(lines)
    translation_count = 0
    in_translation_map = False
    result_lines = []
    
    for line in lines:
        # Detect if we're entering or leaving the getBanglaTranslations map
        if 'func (s *' in line and 'getBanglaTranslations()' in line:
            in_translation_map = True
            result_lines.append(line)
            continue
        
        if in_translation_map and line.strip() == '}':
            in_translation_map = False
            result_lines.append(line)
            continue
        
        # Only translate if we're inside the translation map and the line has a string value
        if in_translation_map and '": "' in line:
            original_line = line
            # Apply translations to the part after ": "
            for english, bangla in FINAL_TRANSLATIONS:
                if english in line:
                    line = line.replace(english, bangla)
                    translation_count += 1
            result_lines.append(line)
        else:
            result_lines.append(line)
    
    new_content = ''.join(result_lines)
    
    # Only write if changes were made
    if new_content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(new_content)
        return translation_count
    
    return 0

def main():
    """Main execution function"""
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    # Find all mobile seeder files
    seeder_files = sorted(seeders_dir.glob("specification_seeder_mobile_*.go"))
    
    print("=" * 80)
    print("TRANSLATING FINAL TERMS TO BANGLA (MAP VALUES ONLY)")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    total_files_modified = 0
    total_translations = 0
    
    for idx, file_path in enumerate(seeder_files, 1):
        translation_count = translate_in_map_values_only(file_path)
        
        if translation_count > 0:
            total_files_modified += 1
            total_translations += translation_count
            print(f"[{idx}/{len(seeder_files)}] ✓ {file_path.name}: {translation_count} translations")
    
    print("\n" + "=" * 80)
    print("FINAL TRANSLATION COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {total_files_modified}")
    print(f"Total translations applied: {total_translations}")
    print("=" * 80)

if __name__ == "__main__":
    main()
