#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Translate remaining English technical terms to Bangla in mobile seeder files
"""

import os
import re
from pathlib import Path

# Define comprehensive translation mappings
REMAINING_TRANSLATIONS = [
    # Material and build terms
    ('Titanium', 'টাইটানিয়াম'),
    ('titanium', 'টাইটানিয়াম'),
    ('plastic', 'প্লাস্টিক'),
    
    # Technical specifications
    ('Ultimate', 'আল্টিমেট'),
    ('ultrawide', 'আল্ট্রাওয়াইড'),
    ('audio', 'অডিও'),
    ('Audio', 'অডিও'),
    ('HIOS', 'এইচআইওএস'),
    
    # Common terms
    ('variant', 'ভেরিয়েন্ট'),
    ('variants', 'ভেরিয়েন্ট'),
    ('support', 'সাপোর্ট'),
    ('certification', 'সার্টিফিকেশন'),
    ('Certification', 'সার্টিফিকেশন'),
    
    # Additional terms that might be missed
    ('according to', 'অনুযায়ী'),
    ('only', 'শুধুমাত্র'),
    ('Some', 'কিছু'),
    ('base', 'বেস'),
    ('low', 'কম'),
    ('light', 'আলো'),
]

def translate_remaining_terms(file_path):
    """Apply remaining term translations to a single file"""
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    original_content = content
    translation_count = 0
    
    # Apply each translation
    for english, bangla in REMAINING_TRANSLATIONS:
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
    print("TRANSLATING REMAINING ENGLISH TERMS TO BANGLA")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    total_files_modified = 0
    total_translations = 0
    
    for idx, file_path in enumerate(seeder_files, 1):
        translation_count = translate_remaining_terms(file_path)
        
        if translation_count > 0:
            total_files_modified += 1
            total_translations += translation_count
            print(f"[{idx}/{len(seeder_files)}] ✓ {file_path.name}: {translation_count} translations")
        else:
            # Only show every 10th unchanged file to reduce output
            if idx % 10 == 0:
                print(f"[{idx}/{len(seeder_files)}] - {file_path.name}: no changes")
    
    print("\n" + "=" * 80)
    print("REMAINING TERMS TRANSLATION COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {total_files_modified}")
    print(f"Total translations applied: {total_translations}")
    print("=" * 80)

if __name__ == "__main__":
    main()
