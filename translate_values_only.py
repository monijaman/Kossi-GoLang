#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Comprehensive translation of ALL remaining English terms to Bangla
ONLY WITHIN TRANSLATION MAP VALUES (right side of ": ")
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
    ('October', 'অক্টোবর'),
    ('November', 'নভেম্বর'),
    ('December', 'ডিসেম্বর'),
    
    # Brand-specific terms and technologies (MUST come before generic terms)
    ('Corning Gorilla Glass Victus 2', 'কর্নিং গরিলা গ্লাস ভিক্টাস ২'),
    ('Corning Gorilla Glass Victus', 'কর্নিং গরিলা গ্লাস ভিক্টাস'),
    ('Corning Gorilla Glass', 'কর্নিং গরিলা গ্লাস'),
    ('Gorilla Glass', 'গরিলা গ্লাস'),
    ('Google Tensor G4', 'গুগল টেনসর জি৪'),
    ('Google Tensor', 'গুগল টেনসর'),
    ('Dolby Vision', 'ডলবি ভিশন'),
    ('Dolby Atmos', 'ডলবি অ্যাটমস'),
    ('ColorOS', 'কালারওএস'),
    ('headphone jack', 'হেডফোন জ্যাক'),
    ('microSD', 'মাইক্রোএসডি'),
    ('Titanium', 'টাইটানিয়াম'),
    ('Ultimate', 'আল্টিমেট'),
    ('Basalt', 'ব্যাসল্ট'),
    ('Island', 'আইল্যান্ড'),
    ('Dynamic Island', 'ডাইনামিক আইল্যান্ড'),
    
    # Units and measurements
    (' MP', ' মেগাপিক্সেল'),
    (' pixels', ' পিক্সেল'),
    (' pixel', ' পিক্সেল'),
    ('pixels', 'পিক্সেল'),
    (' px', ' পিক্সেল'),
    (' ppi', ' পিপিআই'),
    (' fps', ' এফপিএস'),
    (' GHz', ' গিগাহার্টজ'),
    (' MHz', ' মেগাহার্টজ'),
    (' GB"', ' জিবি"'),
    (' GB ', ' জিবি '),
    (' GB/', ' জিবি/'),
    (' TB', ' টিবি'),
    (' mm', ' মিমি'),
    (' g"', ' গ্রাম"'),
    (' g ', ' গ্রাম '),
    (' nm)', ' ন্যানোমিটার)'),
    (' nm ', ' ন্যানোমিটার '),
    (' nm', ' ন্যানোমিটার'),
    (' inches', ' ইঞ্চি'),
    (' inch', ' ইঞ্চি'),
    
    # Common technical terms
    ('Super Retina XDR', 'সুপার রেটিনা এক্সডিআর'),
    ('Retina XDR', 'রেটিনা এক্সডিআর'),
    ('Retina', 'রেটিনা'),
    ('ProMotion', 'প্রোমোশন'),
    ('Dynamic Island', 'ডাইনামিক আইল্যান্ড'),
    ('Always-On', 'অলওয়েজ-অন'),
    ('Android', 'অ্যান্ড্রয়েড'),
    ('iOS', 'আইওএস'),
    ('upgradable', 'আপগ্রেডযোগ্য'),
    ('AMOLED', 'অ্যামোলেড'),
    ('OLED', 'ওএলইডি'),
    ('IPS LCD', 'আইপিএস এলসিডি'),
    ('LCD', 'এলসিডি'),
    ('HDR10+', 'এইচডিআর১০+'),
    ('HDR10', 'এইচডিআর১০'),
    ('HDR', 'এইচডিআর'),
    
    # Materials
    ('Ceramic Shield', 'সিরামিক শিল্ড'),
    ('ceramic', 'সিরামিক'),
    ('stainless steel', 'স্টেইনলেস স্টিল'),
    ('Glass', 'গ্লাস'),
    ('glass', 'গ্লাস'),
    ('plastic', 'প্লাস্টিক'),
    ('aluminum', 'অ্যালুমিনিয়াম'),
    ('Aluminum', 'অ্যালুমিনিয়াম'),
    ('frame', 'ফ্রেম'),
    ('front', 'সামনে'),
    ('back', 'পেছনে'),
    ('silicone', 'সিলিকন'),
    ('polymer', 'পলিমার'),
    ('leather', 'চামড়া'),
    ('metal', 'ধাতু'),
    
    # Common brand/chip terms
    ('Apple A17 Pro', 'অ্যাপল এ১৭ প্রো'),
    ('Apple A16 Bionic', 'অ্যাপল এ১৬ বায়োনিক'),
    ('Apple A15 Bionic', 'অ্যাপল এ১৫ বায়োনিক'),
    ('Bionic', 'বায়োনিক'),
    ('Apple GPU', 'অ্যাপল জিপিইউ'),
    ('Apple', 'অ্যাপল'),
    ('Mediatek', 'মিডিয়াটেক'),
    ('MediaTek', 'মিডিয়াটেক'),
    ('Qualcomm', 'কোয়ালকম'),
    ('Snapdragon', 'স্ন্যাপড্রাগন'),
    ('Dimensity', 'ডাইমেনসিটি'),
    ('Helio', 'হেলিও'),
    ('Mali', 'মালি'),
    ('Adreno', 'অ্যাড্রেনো'),
    ('PowerVR', 'পাওয়ারভিআর'),
    ('Exynos', 'এক্সিনস'),
    ('Tensor', 'টেনসর'),
    ('Unisoc', 'ইউনিসক'),
    ('Xiaomi', 'শাওমি'),
    ('Dragon Crystal', 'ড্রাগন ক্রিস্টাল'),
    ('Elite', 'এলিট'),
    ('Octa', 'অক্টা'),
    ('Octa-core', 'অক্টা-কোর'),
    
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
    ('Gravity', 'গ্র্যাভিটি'),
    ('Cyber', 'সাইবার'),
    ('Mystery', 'মিস্ট্রি'),
    ('Midnight', 'মিডনাইট'),
    ('Aurora', 'অরোরা'),
    ('Ocean', 'ওশান'),
    ('Phantom', 'ফ্যান্টম'),
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
    ('Starlight', 'স্টারলাইট'),
    ('Starry', 'স্টারি'),
    ('Moonlight', 'মুনলাইট'),
    ('Twilight', 'টোয়াইলাইট'),
    ('Mint', 'মিন্ট'),
    ('Pearl', 'পার্ল'),
    ('Diamond', 'ডায়মন্ড'),
    ('Prism', 'প্রিজম'),
    ('Waterfall', 'ওয়াটারফল'),
    ('Thunder', 'থান্ডার'),
    ('Mercurial', 'মার্কিউরিয়াল'),
    ('Oasis', 'ওয়েসিস'),
    ('Sandy', 'স্যান্ডি'),
    ('Meadow', 'মিডো'),
    ('Alpine', 'আলপাইন'),
    ('Woodland', 'উডল্যান্ড'),
    ('Agate', 'অ্যাগেট'),
    ('Astral', 'অ্যাস্ট্রাল'),
    ('Horizon', 'হরাইজন'),
    ('Martian', 'মার্শিয়ান'),
    ('Blush', 'ব্লাশ'),
    ('DayBreak', 'ডেব্রেক'),
    ('Icy', 'আইসি'),
    ('Coral', 'কোরাল'),
    ('Shadow', 'শ্যাডো'),
    ('Onyx', 'অনিক্স'),
    ('Porcelain', 'পোর্সেলেন'),
    ('Crafted', 'ক্রাফটেড'),
    ('Peach', 'পীচ'),
    ('Sunshine', 'সানশাইন'),
    ('Sky', 'স্কাই'),
    ('Forest', 'ফরেস্ট'),
    
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
    ('GSM', 'জিএসএম'),
    ('HSPA', 'এইচএসপিএ'),
    ('LTE', 'এলটিই'),
    ('SIM', 'সিম'),
    ('Dual', 'ডুয়াল'),
    ('Nano', 'ন্যানো'),
    ('eSIM', 'ইসিম'),
    ('USB Type-C', 'ইউএসবি টাইপ-সি'),
    ('Type-C', 'টাইপ-সি'),
    ('USB-C', 'ইউএসবি-সি'),
    ('Lightning', 'লাইটনিং'),
    ('Bluetooth', 'ব্লুটুথ'),
    ('Wi-Fi', 'ওয়াই-ফাই'),
    ('WiFi', 'ওয়াই-ফাই'),
    ('NFC', 'এনএফসি'),
    ('GPS', 'জিপিএস'),
    
    # Display and video technologies
    ('LTPO', 'এলটিপিও'),
    ('LTPS', 'এলটিপিএস'),
    ('IPS', 'আইপিএস'),
    ('UHD', 'ইউএইচডি'),
    ('AI', 'এআই'),
    
    # Camera and feature terms
    ('optical zoom', 'অপটিক্যাল জুম'),
    ('digital zoom', 'ডিজিটাল জুম'),
    ('telephoto', 'টেলিফটো'),
    ('ultrawide', 'আল্ট্রাওয়াইড'),
    ('ultra-wide', 'আল্ট্রা-ওয়াইড'),
    ('wide', 'ওয়াইড'),
    ('macro', 'ম্যাক্রো'),
    ('depth', 'ডেপথ'),
    (' main', ' প্রধান'),
    ('main', 'প্রধান'),
    ('cover', 'কভার'),
    ('Night mode', 'নাইট মোড'),
    ('Portrait', 'পোর্ট্রেট'),
    ('face unlock', 'ফেস আনলক'),
    ('Side', 'সাইড'),
    ('Reflective', 'রিফ্লেক্টিভ'),
    ('Mirror', 'মিরর'),
    (' light', ' লাইট'),
    ('light,', 'লাইট,'),
    
    # AI and Feature terms
    ('assistant', 'সহায়ক'),
    ('Circle to Search', 'সার্কেল টু সার্চ'),
    ('Magic Eraser', 'ম্যাজিক ইরেজার'),
    ('Best Take', 'বেস্ট টেক'),
    ('Audio', 'অডিও'),
    ('Call Screen', 'কল স্ক্রিন'),
    ('Hold for Me', 'হোল্ড ফর মি'),
    ('Gemini', 'জেমিনি'),
    ('Eraser', 'ইরেজার'),
    ('Screen', 'স্ক্রিন'),
    ('Take', 'টেক'),
    ('Call', 'কল'),
    ('Hold', 'হোল্ড'),
    ('Search', 'সার্চ'),
    ('Circle', 'সার্কেল'),
    (' Me', ' মি'),
    
    # Technical terms
    ('Released', 'প্রকাশিত'),
    ('Available', 'উপলব্ধ'),
    ('Rumored', 'গুজব'),
    ('Upcoming', 'আসন্ন'),
    ('dust tight', 'ধুলো প্রতিরোধী'),
    ('water resistant', 'জল প্রতিরোধী'),
    ('up to', 'পর্যন্ত'),
    (' for ', ' জন্য '),
    (' to ', ' থেকে '),
    (' and ', ' এবং '),
    (' or ', ' বা '),
    ('reverse', 'রিভার্স'),
    ('wired', 'তারযুক্ত'),
    ('wireless', 'বেতার'),
    ('peak', 'পিক'),
    ('skin temperature', 'ত্বকের তাপমাত্রা'),
    ('thermometer', 'থার্মোমিটার'),
    ('barometer', 'ব্যারোমিটার'),
    ('ultrasonic', 'আল্ট্রাসনিক'),
    ('under display', 'ডিসপ্লের নিচে'),
    
    # Other common terms
    ('market dependent', 'বাজার নির্ভরশীল'),
    ('region dependent', 'অঞ্চল নির্ভরশীল'),
    ('core', 'কোর'),
    ('-core', '-কোর'),
    ('Pro', 'প্রো'),
    ('Max', 'ম্যাক্স'),
    ('Plus', 'প্লাস'),
    ('Ultra', 'আল্ট্রা'),
    ('Lite', 'লাইট'),
    ('Mini', 'মিনি'),
    ('Gen', 'জেন'),
]

def translate_value_only(file_path):
    """Apply translations ONLY to values (right side of ': ') in getBanglaTranslations map"""
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
            # Split the line into key and value parts
            match = re.match(r'^(\s*"[^"]*":\s*")(.*?)(",?\s*)$', line)
            if match:
                prefix = match.group(1)  # Everything up to and including ': "'
                value = match.group(2)   # The value part
                suffix = match.group(3)  # The closing quote and comma/newline
                
                original_value = value
                # Apply translations only to the value part
                for english, bangla in COMPREHENSIVE_TRANSLATIONS:
                    value = value.replace(english, bangla)
                
                if value != original_value:
                    translation_count += 1
                
                result_lines.append(prefix + value + suffix)
            else:
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
    print("COMPREHENSIVE TRANSLATION (VALUES ONLY)")
    print("=" * 80)
    print(f"Found {len(seeder_files)} mobile seeder files\n")
    print("Processing files...\n")
    
    total_files_modified = 0
    total_translations = 0
    
    for idx, file_path in enumerate(seeder_files, 1):
        translation_count = translate_value_only(file_path)
        
        if translation_count > 0:
            total_files_modified += 1
            total_translations += translation_count
            print(f"[{idx}/{len(seeder_files)}] OK {file_path.name}: {translation_count} lines modified")
    
    print("\n" + "=" * 80)
    print("TRANSLATION COMPLETE")
    print("=" * 80)
    print(f"Files scanned: {len(seeder_files)}")
    print(f"Files modified: {total_files_modified}")
    print(f"Total lines with translations: {total_translations}")
    print("=" * 80)

if __name__ == "__main__":
    main()
