#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Fix JSON syntax errors in mobiledata.json
"""

import json
import re
from pathlib import Path

def fix_json_file(filepath):
    """Fix common JSON syntax errors"""
    print(f"Reading: {filepath}")
    
    with open(filepath, 'r', encoding='utf-8') as f:
        lines = f.readlines()
    
    print(f"Total lines: {len(lines)}")
    
    # Fix line by line
    fixed_lines = []
    for i, line in enumerate(lines):
        # Check if this line is just } without comma and next line starts a new object
        if i < len(lines) - 1:
            if re.match(r'^\s*}\s*$', line):
                next_line = lines[i + 1].strip()
                # If next line starts with a quote (new key), add comma
                if next_line.startswith('"'):
                    line = line.rstrip() + ',\n'
                    print(f"Added comma at line {i+1}")
        
        # Remove trailing commas before closing braces/brackets
        line = re.sub(r',(\s*[}\]])', r'\1', line)
        
        fixed_lines.append(line)
    
    # Write back
    with open(filepath, 'w', encoding='utf-8') as f:
        f.writelines(fixed_lines)
    
    print("Applied fixes")
    
    # Now try to parse and handle duplicates
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            data = json.load(f)
        
        # Check for duplicates
        seen = {}
        duplicates = []
        for key in data.keys():
            if key in seen:
                duplicates.append(key)
            else:
                seen[key] = True
        
        if duplicates:
            print(f"\nWarning: Found {len(duplicates)} duplicate keys:")
            for dup in duplicates:
                print(f"  - {dup}")
            
            # Remove duplicates (keep first occurrence)
            cleaned_data = {}
            for key in data.keys():
                if key not in cleaned_data:
                    cleaned_data[key] = data[key]
            
            # Write cleaned data
            with open(filepath, 'w', encoding='utf-8') as f:
                json.dump(cleaned_data, f, indent=2, ensure_ascii=False)
            
            print(f"\nRemoved duplicates, kept {len(cleaned_data)} unique products")
        else:
            print("No duplicates found")
        
        print(f"\n✓ JSON is now valid with {len(cleaned_data if duplicates else data)} products")
        
    except json.JSONDecodeError as e:
        print(f"\n✗ JSON still has errors: {e}")
        return False
    
    return True

def main():
    json_file = Path(__file__).parent / "init-db" / "seeders" / "mobiledata.json"
    
    success = fix_json_file(json_file)
    
    if success:
        print(f"\n{'='*60}")
        print("JSON file successfully fixed!")
        print(f"{'='*60}")
    else:
        print(f"\n{'='*60}")
        print("JSON file still has errors - manual inspection needed")
        print(f"{'='*60}")

if __name__ == "__main__":
    main()
