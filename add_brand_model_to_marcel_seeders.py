#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Add Brand and Model Name to all empty Marcel refrigerator seeders
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

def update_seeder_file(filepath):
    """Update a seeder file to add Brand and Model Name"""
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()

    # Extract model name from filename
    filename = os.path.basename(filepath)
    model_name = extract_model_from_filename(filename)

    if not model_name:
        print(f"Could not extract model name from {filename}")
        return False

    # Find the empty specs map
    pattern = r'(specs := map\[string\]string\{\s*\n\s*// Specifications will be populated from the database\s*\n\s*// Add your specifications here as they become available\s*\n\s*\})'

    replacement = f'''specs := map[string]string{{
		"Brand":                       "Marcel",
		"Model Name":                  "{model_name}",
		// Add your specifications here as they become available
	}}'''

    new_content = re.sub(pattern, replacement, content, flags=re.MULTILINE | re.DOTALL)

    if new_content != content:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)
        print(f"Updated {filename}")
        return True
    else:
        print(f"No changes needed for {filename}")
        return False

def main():
    """Update all Marcel seeder files"""
    seeders_dir = Path("internal/infrastructure/database/seeders")
    marcel_files = list(seeders_dir.glob("specification_seeder_refrigerator_marcel-*.go"))

    print(f"Found {len(marcel_files)} Marcel seeder files")

    updated = 0
    for filepath in marcel_files:
        if update_seeder_file(filepath):
            updated += 1

    print(f"\nUpdated {updated} files")

if __name__ == "__main__":
    main()