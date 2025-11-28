#!/usr/bin/env python3
"""Translate remaining technical units and terms to complete Bangla"""

from pathlib import Path
import re

SEEDERS_DIR = Path("internal/infrastructure/database/seeders")

# Additional translations for units and remaining terms
UNIT_TRANSLATIONS = [
    # Units
    (' nm)', ' ন্যানোমিটার)'),
    (' nm ', ' ন্যানোমিটার '),
    ('(7 nm)', '(৭ ন্যানোমিটার)'),
    ('(6 nm)', '(৬ ন্যানোমিটার)'),
    ('(5 nm)', '(৫ ন্যানোমিটার)'),
    ('(4 nm)', '(৪ ন্যানোমিটার)'),
    ('(3 nm)', '(৩ ন্যানোমিটার)'),
    ('(12 nm)', '(১২ ন্যানোমিটার)'),
    (' GB"', ' জিবি"'),
    (' GB ', ' জিবি '),
    (' GB/', ' জিবি/'),
    ('GB /', 'জিবি /'),
    (' GHz ', ' গিগাহার্জ '),
    ('GHz Cortex', 'গিগাহার্জ Cortex'),
    (' ppi)', ' পিপিআই)'),
    (' fps', ' এফপিএস'),
    ('30fps', '৩০এফপিএস'),
    ('60fps', '৬০এফপিএস'),
    ('120fps', '১২০এফপিএস'),
    ('240fps', '২৪০এফপিএস'),
    ('@30fps', '@৩০এফপিএস'),
    ('@60fps', '@৬০এফপিএস'),
    
    # Technical terms
    ('Octa-core', 'অক্টা-কোর'),
    ('Hexa-core', 'হেক্সা-কোর'),
    ('Quad-core', 'কোয়াড-কোর'),
    ('Deca-core', 'ডেকা-কোর'),
    ('Nona-core', 'নোনা-কোর'),
    
    # Brand series names
    ('Neo ', 'নিও '),
    
    # Other remaining terms
    (' px ', ' পিক্সেল '),
    (' px)', ' পিক্সেল)'),
    (' px,', ' পিক্সেল,'),
    ('× ', '× '),
]

def translate_file(filepath):
    """Apply unit translations to a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        changes = 0
        
        # Only translate within translation value strings
        lines = content.split('\n')
        new_lines = []
        in_translations_map = False
        
        for line in lines:
            if 'getBanglaTranslations' in line or 'return map[string]string{' in line:
                in_translations_map = True
                new_lines.append(line)
                continue
            
            if in_translations_map and line.strip() == '}':
                in_translations_map = False
                new_lines.append(line)
                continue
            
            if in_translations_map and '":' in line and '"' in line:
                parts = line.split(':', 1)
                if len(parts) == 2:
                    key_part = parts[0]
                    value_part = parts[1]
                    original_value = value_part
                    
                    # Apply all unit translations
                    for english, bangla in UNIT_TRANSLATIONS:
                        value_part = value_part.replace(english, bangla)
                    
                    if value_part != original_value:
                        changes += 1
                    
                    new_lines.append(key_part + ':' + value_part)
                else:
                    new_lines.append(line)
            else:
                new_lines.append(line)
        
        new_content = '\n'.join(new_lines)
        
        if new_content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(new_content)
            return changes
        
        return 0
        
    except Exception as e:
        print(f"  ✗ Error in {filepath.name}: {e}")
        return 0

def main():
    print("=" * 80)
    print("TRANSLATING REMAINING UNITS AND TECHNICAL TERMS TO BANGLA")
    print("=" * 80)
    
    if not SEEDERS_DIR.exists():
        print(f"Error: Directory not found: {SEEDERS_DIR}")
        return
    
    seeder_files = sorted(SEEDERS_DIR.glob("specification_seeder_mobile_*.go"))
    
    if not seeder_files:
        print(f"No mobile seeder files found")
        return
    
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    modified_count = 0
    total_changes = 0
    
    for i, filepath in enumerate(seeder_files, 1):
        changes = translate_file(filepath)
        if changes > 0:
            modified_count += 1
            total_changes += changes
            print(f"[{i}/{len(seeder_files)}] ✓ {filepath.name}: {changes} translations")
        elif i % 50 == 0:
            print(f"[{i}/{len(seeder_files)}] Processed...")
    
    print("\n" + "=" * 80)
    print("UNIT TRANSLATION COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {modified_count}")
    print(f"Total translations applied: {total_changes}")
    print("=" * 80)

if __name__ == "__main__":
    main()
