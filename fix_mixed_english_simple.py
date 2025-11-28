#!/usr/bin/env python3
"""Fix mixed English/Bangla translations using simple string replacements"""

from pathlib import Path

SEEDERS_DIR = Path("internal/infrastructure/database/seeders")

# Simple string replacements for common mixed translations
REPLACEMENTS = [
    # Plastic frame/back patterns
    ('Plastic frame / back', 'প্লাস্টিক ফ্রেম /পেছনে'),
    ('plastic frame / back', 'প্লাস্টিক ফ্রেম /পেছনে'),
    ('Plastic frame & back', 'প্লাস্টিক ফ্রেম & পেছনে'),
    ('plastic frame & back', 'প্লাস্টিক ফ্রেম & পেছনে'),
    ('plastic frame/back', 'প্লাস্টিক ফ্রেম/পেছনে'),
    ('plastic frame, ', 'প্লাস্টিক ফ্রেম, '),
    ('plastic frame"', 'প্লাস্টিক ফ্রেম"'),
    
    # Back patterns
    ('/back,', '/পেছনে,'),
    (' back,', ' পেছনে,'),
    (' back"', ' পেছনে"'),
    ('and back', 'এবং পেছনে'),
    ('& back', '& পেছনে'),
    ('or eco leather back', 'অথবা ইকো লেদার পেছনে'),
    ('eco-leather back', 'ইকো লেদার পেছনে'),
    ('eco leather back', 'ইকো লেদার পেছনে'),
    ('silicone polymer back', 'সিলিকন পলিমার পেছনে'),
    ('glass/silicone polymer back', 'গ্লাস/সিলিকন পলিমার পেছনে'),
    ('ceramic/polymer back', 'সিরামিক/পলিমার পেছনে'),
    
    # Gen patterns in processor names
    (' Gen 1"', ' জেন ১"'),
    (' Gen 1 ', ' জেন ১ '),
    (' Gen 2"', ' জেন ২"'),
    (' Gen 2 ', ' জেন ২ '),
    (' Gen 3"', ' জেন ৩"'),
    (' Gen 3 ', ' জেন ৩ '),
    (' Gen 4"', ' জেন ৪"'),
    (' Gen 4 ', ' জেন ৪ '),
    
    # Headphone jack
    ('headphone jack', 'হেডফোন জ্যাক'),
    
    # Color descriptors
    ('Cool', 'কুল'),
    ('Electric', 'ইলেকট্রিক'),
    ('Satin', 'সাটিন'),
    ('Sparkle', 'স্পার্কল'),
    ('Cyan', 'সায়ান'),
    ('Violet', 'ভায়োলেট'),
    ('Onyx', 'ওনিক্স'),
    ('Marble', 'মার্বেল'),
    ('Cobalt', 'কোবাল্ট'),
    ('Amber', 'অ্যাম্বার'),
    ('Iron', 'আয়রন'),
    ('Luxe', 'লাক্স'),
    ('Coral', 'কোরাল'),
    ('Nebula', 'নেবুলা'),
    ('Celestial', 'সেলেস্টিয়াল'),
    ('Reflective', 'রিফ্লেক্টিভ'),
    ('Honey Dew', 'হানি ডিউ'),
    ('Fusion', 'ফিউশন'),
    ('Graphite', 'গ্রাফাইট'),
    ('Lime', 'লাইম'),
    
    # Status words
    ('Released', 'মুক্তি'),
    
    # Month names
    ('January', 'জানুয়ারি'),
    ('February', 'ফেব্রুয়ারি'),
    ('March', 'মার্চ'),
    ('April', 'এপ্রিল'),
    ('May', 'মে'),
    ('June', 'জুন'),
    ('July', 'জুলাই'),
    ('August', 'আগস্ট'),
    ('September', 'সেপ্টেম্বর'),
    ('October', 'অক্টোবর'),
    ('November', 'নভেম্বর'),
    ('December', 'ডিসেম্বর'),
]

def fix_file(filepath):
    """Apply all string replacements to a file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        changes_made = []
        
        for old_text, new_text in REPLACEMENTS:
            if old_text in content:
                content = content.replace(old_text, new_text)
                changes_made.append(old_text)
        
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return len(changes_made)
        
        return 0
        
    except Exception as e:
        print(f"  ✗ Error: {e}")
        return 0

def main():
    print("=" * 80)
    print("FIXING MIXED ENGLISH/BANGLA TRANSLATIONS")
    print("=" * 80)
    
    if not SEEDERS_DIR.exists():
        print(f"Error: Directory not found: {SEEDERS_DIR}")
        return
    
    seeder_files = sorted(SEEDERS_DIR.glob("specification_seeder_mobile_*.go"))
    
    if not seeder_files:
        print(f"No mobile seeder files found")
        return
    
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    
    fixed_count = 0
    total_replacements = 0
    
    for filepath in seeder_files:
        changes = fix_file(filepath)
        if changes > 0:
            fixed_count += 1
            total_replacements += changes
            print(f"✓ {filepath.name}: {changes} replacements")
    
    print("\n" + "=" * 80)
    print("COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {fixed_count}")
    print(f"Total replacements: {total_replacements}")
    print("=" * 80)

if __name__ == "__main__":
    main()
