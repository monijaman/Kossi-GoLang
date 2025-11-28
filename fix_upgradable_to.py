#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Fix incorrect "upgradable to" translation
"""

import os
import re
from pathlib import Path

# Define fix
FIX_TRANSLATIONS = [
    ('আপগ্রেডযোগ্য বা', 'আপগ্রেডযোগ্য'),
]

def fix_translation(file_path):
    """Fix incorrect translations in a single file"""
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    original_content = content
    fix_count = 0
    
    # Apply each fix
    for wrong, correct in FIX_TRANSLATIONS:
        if wrong in content:
            count_before = content.count(wrong)
            content = content.replace(wrong, correct)
            count_after = content.count(wrong)
            fix_count += (count_before - count_after)
    
    # Only write if changes were made
    if content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        return fix_count
    
    return 0

def main():
    """Main execution function"""
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    # Find all mobile seeder files
    seeder_files = sorted(seeders_dir.glob("specification_seeder_mobile_*.go"))
    
    print("=" * 80)
    print("FIXING INCORRECT TRANSLATION")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    
    total_files_fixed = 0
    total_fixes = 0
    
    for file_path in seeder_files:
        fix_count = fix_translation(file_path)
        
        if fix_count > 0:
            total_files_fixed += 1
            total_fixes += fix_count
            print(f"✓ {file_path.name}: {fix_count} fixes")
    
    print("\n" + "=" * 80)
    print("FIXES COMPLETE")
    print("=" * 80)
    print(f"Files fixed: {total_files_fixed}")
    print(f"Total fixes applied: {total_fixes}")
    print("=" * 80)

if __name__ == "__main__":
    main()
