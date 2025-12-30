#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Add Bangla translations for Brand and Model Name to all Marcel refrigerator seeders
"""

import os
import re
from pathlib import Path

def extract_model_from_filename(filename):
    """Extract model name from filename"""
    # filename format: specification_seeder_refrigerator_marcel-{model}.go
    match = re.search(r'specification_seeder_refrigerator_marcel-(.+)\.go', filename)
    if match:
        return match.group(1).upper().replace('-', '-')
    return ""

def update_bangla_translations(filepath):
    """Update Bangla translations to include Brand and Model Name"""
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()

    # Extract model name from filename
    filename = os.path.basename(filepath)
    model_slug = filename.replace('specification_seeder_refrigerator_marcel-', '').replace('.go', '')
    model_name = model_slug.upper().replace('-', '-')

    # Find the getBanglaTranslations function and add Brand and Model Name translations
    pattern = r'("marcel-[^"]*":\s*"[^"]*",\s*\n\s*// Add more translations as needed)'

    replacement = f'''		"marcel-{model_slug}":         "মার্সেল-{model_slug.replace('-', '-')}",
		"{model_name}":   "{model_name}",
		// Add more translations as needed'''

    new_content = re.sub(pattern, replacement, content, flags=re.MULTILINE | re.DOTALL)

    if new_content != content:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)
        print(f"Updated Bangla translations for {filename}")
        return True
    else:
        print(f"No changes needed for {filename}")
        return False

def main():
    """Update Bangla translations for all Marcel seeder files"""
    seeders_dir = Path("internal/infrastructure/database/seeders")
    marcel_files = list(seeders_dir.glob("specification_seeder_refrigerator_marcel-*.go"))

    print(f"Found {len(marcel_files)} Marcel seeder files")

    updated = 0
    for filepath in marcel_files:
        if update_bangla_translations(filepath):
            updated += 1

    print(f"\nUpdated {updated} files")

if __name__ == "__main__":
    main()