#!/usr/bin/env python3
"""Fix mixed Bengali/English translations in all mobile seeder files"""

import os
import re
from pathlib import Path

SEEDERS_DIR = Path("internal/infrastructure/database/seeders")

def fix_file(filepath):
    """Fix mixed translations in a single file using regex patterns"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        changes = 0
        
        # Only process lines within getBanglaTranslations map (translation values after colon)
        def fix_translation_line(match):
            nonlocal changes
            full_line = match.group(0)
            key_part = match.group(1)  # Everything before colon including the key
            value_part = match.group(2)  # Everything after colon
            
            original_value = value_part
            
            # Fix "back" when mixed with Bangla
            value_part = re.sub(r'/\s*back\b', '/পেছনে', value_part)
            value_part = re.sub(r'\s+back,', ' পেছনে,', value_part)
            value_part = re.sub(r'\s+back"', ' পেছনে"', value_part)
            value_part = re.sub(r'\band back\b', 'এবং পেছনে', value_part)
            value_part = re.sub(r'&\s*back\b', '& পেছনে', value_part)
            
            # Fix plastic frame
            value_part = re.sub(r'\bplastic frame\b', 'প্লাস্টিক ফ্রেম', value_part, flags=re.IGNORECASE)
            value_part = re.sub(r'\bPlastic frame\b', 'প্লাস্টিক ফ্রেম', value_part)
            
            # Fix eco leather back
            value_part = re.sub(r'\beco[- ]leather back\b', 'ইকো লেদার পেছনে', value_part, flags=re.IGNORECASE)
            
            # Fix silicone polymer back
            value_part = re.sub(r'\bsilicone polymer back\b', 'সিলিকন পলিমার পেছনে', value_part, flags=re.IGNORECASE)
            value_part = re.sub(r'\bglass/silicone polymer back\b', 'গ্লাস/সিলিকন পলিমার পেছনে', value_part, flags=re.IGNORECASE)
            
            # Fix ceramic/polymer back
            value_part = re.sub(r'\bceramic/polymer back\b', 'সিরামিক/পলিমার পেছনে', value_part, flags=re.IGNORECASE)
            
            # Fix "Gen X" in processor names (when mixed with Bangla numbers)
            value_part = re.sub(r'\bGen\s+(\d)', r'জেন \1', value_part)
            
            # Fix headphone jack
            value_part = re.sub(r'\bheadphone jack\b', 'হেডফোন জ্যাক', value_part, flags=re.IGNORECASE)
            
            # Fix color names when mixed
            value_part = re.sub(r'\bViolet\b', 'ভায়োলেট', value_part)
            value_part = re.sub(r'\bCyan\b', 'সায়ান', value_part)
            value_part = re.sub(r'\bSatin\b', 'সাটিন', value_part)
            value_part = re.sub(r'\bSparkle\b', 'স্পার্কল', value_part)
            value_part = re.sub(r'\bElectric\b', 'ইলেকট্রিক', value_part)
            value_part = re.sub(r'\bCool\b', 'কুল', value_part)
            value_part = re.sub(r'\bIron\b', 'আয়রন', value_part)
            value_part = re.sub(r'\bOnyx\b', 'ওনিক্স', value_part)
            value_part = re.sub(r'\bMarble\b', 'মার্বেল', value_part)
            value_part = re.sub(r'\bCobalt\b', 'কোবাল্ট', value_part)
            value_part = re.sub(r'\bAmber\b', 'অ্যাম্বার', value_part)
            value_part = re.sub(r'\bLuxe\b', 'লাক্স', value_part)
            value_part = re.sub(r'\bCoral\b', 'কোরাল', value_part)
            value_part = re.sub(r'\bNebula\b', 'নেবুলা', value_part)
            value_part = re.sub(r'\bCelestial\b', 'সেলেস্টিয়াল', value_part)
            value_part = re.sub(r'\bReflective\b', 'রিফ্লেক্টিভ', value_part)
            value_part = re.sub(r'\bHoney Dew\b', 'হানি ডিউ', value_part)
            value_part = re.sub(r'\bFusion\b', 'ফিউশন', value_part)
            value_part = re.sub(r'\bGraphite\b', 'গ্রাফাইট', value_part)
            value_part = re.sub(r'\bLime\b', 'লাইম', value_part)
            
            # Fix "Released" when mixed
            value_part = re.sub(r'\bReleased\b', 'মুক্তি', value_part)
            
            # Fix month names
            value_part = re.sub(r'\bJanuary\b', 'জানুয়ারি', value_part)
            value_part = re.sub(r'\bFebruary\b', 'ফেব্রুয়ারি', value_part)
            value_part = re.sub(r'\bMarch\b', 'মার্চ', value_part)
            value_part = re.sub(r'\bApril\b', 'এপ্রিল', value_part)
            value_part = re.sub(r'\bMay\b', 'মে', value_part)
            value_part = re.sub(r'\bJune\b', 'জুন', value_part)
            value_part = re.sub(r'\bJuly\b', 'জুলাই', value_part)
            value_part = re.sub(r'\bAugust\b', 'আগস্ট', value_part)
            value_part = re.sub(r'\bSeptember\b', 'সেপ্টেম্বর', value_part)
            value_part = re.sub(r'\bOctober\b', 'অক্টোবর', value_part)
            value_part = re.sub(r'\bNovember\b', 'নভেম্বর', value_part)
            value_part = re.sub(r'\bDecember\b', 'ডিসেম্বর', value_part)
            
            if value_part != original_value:
                changes += 1
            
            return key_part + value_part
        
        # Match translation map entries:  "key": "value"
        # Capture everything before colon and everything after (including colon)
        content = re.sub(
            r'^(\s*"[^"]+"\s*:\s*)("[^"]*".*)$',
            fix_translation_line,
            content,
            flags=re.MULTILINE
        )
        
        # Write back if changes were made
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return changes
        
        return 0
        
    except Exception as e:
        print(f"  ✗ Error processing {filepath.name}: {e}")
        return 0

def main():
    print("=" * 80)
    print("FIXING MIXED BENGALI/ENGLISH TRANSLATIONS IN MOBILE SEEDERS")
    print("=" * 80)
    
    if not SEEDERS_DIR.exists():
        print(f"Error: Directory not found: {SEEDERS_DIR}")
        return
    
    # Find all mobile specification seeder files
    seeder_files = sorted(SEEDERS_DIR.glob("specification_seeder_mobile_*.go"))
    
    if not seeder_files:
        print(f"No mobile seeder files found in {SEEDERS_DIR}")
        return
    
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    
    fixed_count = 0
    total_changes = 0
    
    for filepath in seeder_files:
        changes = fix_file(filepath)
        if changes > 0:
            fixed_count += 1
            total_changes += changes
            print(f"✓ {filepath.name}: {changes} translations fixed")
    
    print("\n" + "=" * 80)
    print("MIXED TRANSLATION FIX COMPLETE")
    print("=" * 80)
    print(f"Total files scanned: {len(seeder_files)}")
    print(f"Files modified: {fixed_count}")
    print(f"Total translations fixed: {total_changes}")
    print("=" * 80)

if __name__ == "__main__":
    main()
