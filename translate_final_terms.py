#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Translate final remaining English terms to Bangla in mobile seeder files
"""

import os
import re
from pathlib import Path

# Define final translation mappings
FINAL_TRANSLATIONS = [
    # Conjunctions and common words
    (' or ', ' বা '),
    ('or ', 'বা '),
    
    # Material terms
    ('silicone', 'সিলিকন'),
    ('polymer', 'পলিমার'),
    ('eco leather', 'ইকো লেদার'),
    ('ceramic', 'সিরামিক'),
    
    # Status terms
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
    ('Moonআলো', 'মুনলাইট'),  # Fix previous bad translation
    ('Twilight', 'টোয়াইলাইট'),
    ('Twiআলো', 'টোয়াইলাইট'),  # Fix previous bad translation
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

def translate_final_terms(file_path):
    """Apply final term translations to a single file"""
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    original_content = content
    translation_count = 0
    
    # Apply each translation
    for english, bangla in FINAL_TRANSLATIONS:
        if english in content:
            # Count occurrences before replacement
            count_before = content.count(english)
            content = content.replace(english, bangla)
            count_after = content.count(english)
            translation_count += (count_before - count_after)
    
    # Only write if changes were made
    if content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        return translation_count
    
    return 0

def main():
    """Main execution function"""
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    # Find all mobile seeder files
    seeder_files = sorted(seeders_dir.glob("specification_seeder_mobile_*.go"))
    
    print("=" * 80)
    print("TRANSLATING FINAL REMAINING TERMS TO BANGLA")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    total_files_modified = 0
    total_translations = 0
    
    for idx, file_path in enumerate(seeder_files, 1):
        translation_count = translate_final_terms(file_path)
        
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
