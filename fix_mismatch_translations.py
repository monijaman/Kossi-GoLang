#!/usr/bin/env python3
"""Fix remaining mismatches - English not properly translated to Bengali"""

import os
import glob

SEEDERS_DIR = "init-db/seeders"

# Critical mismatches to fix
MISMATCH_FIXES = {
    # WiFi "g" incorrectly translated as "gram"
    '"802.11 a/b/g/n"': '"802.11 a/b/g/n"',  # Keep as is - don't translate WiFi bands
    '"802.11 a/b/গ্রাম/n"': '"802.11 a/b/g/n"',
    '"Wi-Fi 802.11 a/b/গ্রাম/n"': '"Wi-Fi 802.11 a/b/g/n"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n"': '"ওয়াই-ফাই 802.11 a/b/g/n"',
    '"Wi-Fi 802.11 a/b/গ্রাম/n/ac"': '"Wi-Fi 802.11 a/b/g/n/ac"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n/ac"': '"ওয়াই-ফাই 802.11 a/b/g/n/ac"',
    '"Wi‑Fi 802.11 a/b/গ্রাম/n/ac"': '"Wi-Fi 802.11 a/b/g/n/ac"',
    '"Wi-Fi 802.11 a/b/গ্রাম/n/ac/6e"': '"Wi-Fi 802.11 a/b/g/n/ac/6e"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n/ac/6e"': '"ওয়াই-ফাই 802.11 a/b/g/n/ac/6e"',
    '"802.11 a/b/গ্রাম/n/ac"': '"802.11 a/b/g/n/ac"',
    '"802.11 a/b/গ্রাম/n/ac/ax"': '"802.11 a/b/g/n/ac/ax"',
    '"Wi-Fi 802.11 a/b/গ্রাম/n/ac/6/6E"': '"Wi-Fi 802.11 a/b/g/n/ac/6/6E"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n/ac/6/6E"': '"ওয়াই-ফাই 802.11 a/b/g/n/ac/6/6E"',
    '"Wi-Fi 802.11 a/b/গ্রাম/n/ac/6e/7, tri-band, Wi-Fi Direct"': '"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct"',
    '"ওয়াই-ফাই 802.11 a/b/গ্রাম/n/ac/6e/7, tri-ব্যান্ড, ওয়াই-ফাই Direct"': '"ওয়াই-ফাই 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct"',
    
    # "glass" not translated in descriptions
    '"Ceramic Shield, harder than any smartphone গ্লাস"': '"Ceramic Shield, harder than any smartphone গ্লাস"',
    
    # "image" and other words not translated
    '"Optical image স্ট্যাবিলাইজেশন"': '"অপটিক্যাল ইমেজ স্ট্যাবিলাইজেশন"',
    '"অপটিক্যাল image স্ট্যাবিলাইজেশন (OIS), Sensor-shift স্ট্যাবিলাইজেশন"': '"অপটিক্যাল ইমেজ স্ট্যাবিলাইজেশন (OIS), সেন্সর-শিফট স্ট্যাবিলাইজেশন"',
    
    # "mode" not translated
    '"Night mode"': '"Night মোড"',
    '"Action mode"': '"Action মোড"',
    '"Portrait mode"': '"Portrait মোড"',
    '"Pro mode"': '"প্রো মোড"',
    
    # "Smart" not translated
    '"Smart এইচডিআর"': '"স্মার্ট এইচডিআর"',
    '"Smart HDR"': '"স্মার্ট এইচডিআর"',
    
    # Display tech partially translated
    '"LED flash, HDR, panorama, OIS on প্রশস্ত এবং টেলিফটো, Night মোড, Portrait মোড"': '"LED ফ্ল্যাশ, এইচডিআর, প্যানোরামা, OIS on প্রশস্ত এবং টেলিফটো, Night মোড, Portrait মোড"',
}

def fix_file(filepath):
    """Fix translation mismatches in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        for old_text, new_text in MISMATCH_FIXES.items():
            if old_text in content:
                content = content.replace(old_text, new_text)
                fixed_count += 1
        
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return fixed_count > 0
        return False
    except Exception as e:
        print(f"  ✗ Error: {e}")
        return False

def main():
    print("=" * 70)
    print("FIXING TRANSLATION MISMATCHES (English-Bengali pairs)")
    print("=" * 70)
    
    seeder_files = sorted(glob.glob(os.path.join(SEEDERS_DIR, "specification_seeder_mobile_*.go")))
    
    fixed_count = 0
    total_count = len(seeder_files)
    
    for i, filepath in enumerate(seeder_files, 1):
        if i % 50 == 0 or i == total_count:
            print(f"\n[Progress] Processing {i}/{total_count} files...")
        
        if fix_file(filepath):
            fixed_count += 1
    
    print("\n" + "=" * 70)
    print("TRANSLATION MISMATCH FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Mismatch patterns fixed: {len(MISMATCH_FIXES)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
