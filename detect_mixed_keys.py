#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Detect and report keys that have mixed Bangla/English text
These keys should be fully in English
"""

import re
from pathlib import Path

def has_bangla(text):
    """Check if text contains Bangla characters"""
    return bool(re.search(r'[অ-হ০-৯]', text))

def has_english_letters(text):
    """Check if text contains English letters"""
    return bool(re.search(r'[A-Za-z]', text))

def find_mixed_keys(file_path):
    """Find keys that contain both Bangla and English"""
    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()
    
    mixed_keys = []
    in_translation_map = False
    
    for line_num, line in enumerate(lines, 1):
        # Detect if we're in the translation map
        if 'getBanglaTranslations()' in line:
            in_translation_map = True
            continue
        
        if in_translation_map and line.strip() == '}':
            in_translation_map = False
            continue
        
        # Check for mixed keys (has both Bangla and English)
        if in_translation_map and '": "' in line:
            match = re.match(r'^\s*"([^"]+)":\s*"', line)
            if match:
                key = match.group(1)
                if has_bangla(key) and has_english_letters(key):
                    mixed_keys.append((line_num, key, line.strip()))
    
    return mixed_keys

def main():
    """Main execution function"""
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    # Find all mobile seeder files
    seeder_files = sorted(seeders_dir.glob("specification_seeder_mobile_*.go"))
    
    print("=" * 80)
    print("DETECTING MIXED BANGLA/ENGLISH IN KEYS")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    
    total_files_with_issues = 0
    total_mixed_keys = 0
    all_unique_patterns = set()
    
    for file_path in seeder_files:
        mixed_keys = find_mixed_keys(file_path)
        
        if mixed_keys:
            total_files_with_issues += 1
            total_mixed_keys += len(mixed_keys)
            print(f"\n{file_path.name}:")
            for line_num, key, full_line in mixed_keys:
                print(f"  Line {line_num}: {key}")
                all_unique_patterns.add(full_line)
    
    print("\n" + "=" * 80)
    print("DETECTION COMPLETE")
    print("=" * 80)
    print(f"Files with issues: {total_files_with_issues}")
    print(f"Total mixed keys: {total_mixed_keys}")
    print(f"Unique patterns: {len(all_unique_patterns)}")
    print("=" * 80)
    
    if all_unique_patterns:
        print("\nAll unique mixed key patterns found:")
        print("-" * 80)
        for pattern in sorted(all_unique_patterns):
            print(pattern)

if __name__ == "__main__":
    main()
