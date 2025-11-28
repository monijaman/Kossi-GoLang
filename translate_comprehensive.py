#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Comprehensive translation of ALL remaining English terms to Bangla
ONLY WITHIN TRANSLATION MAP VALUES
"""

import os
import re
from pathlib import Path

# Define comprehensive translation mappings
COMPREHENSIVE_TRANSLATIONS = [
    # Months (English months that weren't translated yet)
    ('January', 'জানুয়ারি'),
    ('February', 'ফেব্রুয়ারি'),
    ('March', 'মার্চ'),
    ('April', 'এপ্রিল'),
    ('May', 'মে'),
    ('June', 'জুন'),
    ('July', 'জুলাই'),
    ('August', 'আগস্ট'),
    ('September', 'সেপ্টেম্বর'),
    ('November', 'নভেম্বর'),
    ('December', 'ডিসেম্বর'),
    
    # Units and measurements
    (' MP', ' মেগাপিক্সেল'),
    (' pixels', ' পিক্সেল'),
    (' pixel', ' পিক্সেল'),
    ('pixels', 'পিক্সেল'),
    (' GB"', ' জিবি"'),
    (' GB ', ' জিবি '),
    (' mm', ' মিমি'),
    (' g"', ' গ্রাম"'),
    (' g ', ' গ্রাম '),
    (' nm)', ' ন্যানোমিটার)'),
    (' nm ', ' ন্যানোমিটার '),
    (' inches', ' ইঞ্চি'),
    (' inch', ' ইঞ্চি'),
    
    # Common technical terms
    ('Android', 'অ্যান্ড্রয়েড'),
    ('iOS', 'আইওএস'),
    ('AMOLED', 'অ্যামোলেড'),
    ('OLED', 'ওএলইডি'),
    ('IPS LCD', 'আইপিএস এলসিডি'),
    ('LCD', 'এলসিডি'),
    ('Gorilla Glass', 'গরিলা গ্লাস'),
    ('Corning', 'কর্নিং'),
    
    # Materials
    ('Glass', 'গ্লাস'),
    ('glass', 'গ্লাস'),
    ('plastic', 'প্লাস্টিক'),
    ('aluminum', 'অ্যালুমিনিয়াম'),
    ('Aluminum', 'অ্যালুমিনিয়াম'),
    ('frame', 'ফ্রেম'),
    ('front', 'সামনে'),
    ('back', 'পেছনে'),
    
    # Common brand/chip terms
    ('Mediatek', 'মিডিয়াটেক'),
    ('MediaTek', 'মিডিয়াটেক'),
    ('Qualcomm', 'কোয়ালকম'),
    ('Snapdragon', 'স্ন্যাপড্রাগন'),
    ('Dimensity', 'ডাইমেনসিটি'),
    ('Helio', 'হেলিও'),
    ('Mali', 'মালি'),
    ('Adreno', 'অ্যাড্রেনো'),
    ('PowerVR', 'পাওয়ারভিআর'),
    
    # Colors
    ('Black', 'কালো'),
    ('White', 'সাদা'),
    ('Blue', 'নীল'),
    ('Red', 'লাল'),
    ('Green', 'সবুজ'),
    ('Yellow', 'হলুদ'),
    ('Purple', 'বেগুনি'),
    ('Pink', 'গোলাপী'),
    ('Gray', 'ধূসর'),
    ('Grey', 'ধূসর'),
    ('Silver', 'রূপালী'),
    ('Gold', 'সোনালী'),
    ('Bronze', 'ব্রোঞ্জ'),
    ('Copper', 'তামা'),
    ('Violet', 'ভায়োলেট'),
    ('Brown', 'বাদামী'),
    ('Beige', 'বেইজ'),
    ('Cream', 'ক্রিম'),
    ('Ivory', 'আইভরি'),
    ('Rose', 'রোজ'),
    
    # Color descriptors
    ('Magic', 'ম্যাজিক'),
    ('Variable', 'ভেরিয়েবল'),
    ('Obsidian', 'অবসিডিয়ান'),
    ('Hazel', 'হ্যাজেল'),
    ('Olive', 'অলিভ'),
    ('Lake', 'লেক'),
    ('Cosmic', 'কসমিক'),
    ('Space', 'স্পেস'),
    ('Frost', 'ফ্রস্ট'),
    ('Classic', 'ক্লাসিক'),
    ('Navy', 'নেভি'),
    ('Divine', 'ডিভাইন'),
    ('Ice', 'আইস'),
    ('Radium', 'রেডিয়াম'),
    ('Imperial', 'ইম্পেরিয়াল'),
    ('Swamp', 'সোয়াম্প'),
    ('Graphite', 'গ্রাফাইট'),
    ('Jade', 'জেড'),
    ('Jet', 'জেট'),
    ('Natural', 'ন্যাচারাল'),
    ('Desert', 'ডেজার্ট'),
    
    # Other technical terms
    ('colors', 'রঙ'),
    ('color', 'রঙ'),
    ('XOS', 'এক্সওএস'),
    ('MIUI', 'এমআইইউআই'),
    ('HyperOS', 'হাইপার ওএস'),
    ('One UI', 'ওয়ান ইউআই'),
    ('Available', 'উপলব্ধ'),
    ('wired', 'তারযুক্ত'),
    ('wireless', 'বেতার'),
    
    # Storage/Connectivity
    ('SIM', 'সিম'),
    ('Dual', 'ডুয়াল'),
    ('Nano', 'ন্যানো'),
    ('eSIM', 'ইসিম'),
]

def translate_comprehensive(file_path):
    """Apply comprehensive translations ONLY to values in getBanglaTranslations map"""
    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()
    
    original_content = ''.join(lines)
    translation_count = 0
    in_translation_map = False
    result_lines = []
    
    for line in lines:
        # Detect if we're entering or leaving the getBanglaTranslations map
        if 'func (s *' in line and 'getBanglaTranslations()' in line:
            in_translation_map = True
            result_lines.append(line)
            continue
        
        if in_translation_map and line.strip() == '}':
            in_translation_map = False
            result_lines.append(line)
            continue
        
        # Only translate if we're inside the translation map and the line has a string value
        if in_translation_map and '": "' in line:
            original_line = line
            # Apply translations to the part after ": "
            for english, bangla in COMPREHENSIVE_TRANSLATIONS:
                if english in line:
                    line = line.replace(english, bangla)
            if line != original_line:
                translation_count += 1
            result_lines.append(line)
        else:
            result_lines.append(line)
    
    new_content = ''.join(result_lines)
    
    # Only write if changes were made
    if new_content != original_content:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(new_content)
        return translation_count
    
    return 0

def main():
    """Main execution function"""
    seeders_dir = Path(__file__).parent / "internal" / "infrastructure" / "database" / "seeders"
    
    # Find all mobile seeder files
    seeder_files = sorted(seeders_dir.glob("specification_seeder_mobile_*.go"))
    
    print("=" * 80)
    print("COMPREHENSIVE TRANSLATION OF ALL REMAINING TERMS")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    total_files_modified = 0
    total_translations = 0
    
    for idx, file_path in enumerate(seeder_files, 1):
        translation_count = translate_comprehensive(file_path)
        
        if translation_count > 0:
            total_files_modified += 1
            total_translations += translation_count
            print(f"[{idx}/{len(seeder_files)}] ✓ {file_path.name}: {translation_count} lines modified")
    
    print("\n" + "=" * 80)
    print("COMPREHENSIVE TRANSLATION COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {total_files_modified}")
    print(f"Total lines with translations: {total_translations}")
    print("=" * 80)

if __name__ == "__main__":
    main()
