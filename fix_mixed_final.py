#!/usr/bin/env python3
"""Fix remaining mixed English-Bengali patterns"""

import os
import glob
import re

SEEDERS_DIR = "init-db/seeders"

# Comprehensive replacement patterns
FIXES = {
    # Display technology
    '"LTPO Super Retina XDR OLED, 120Hz": "এলটিপিও Super Retina XDR ওলেড, 120Hz"': '"LTPO Super Retina XDR OLED, 120Hz": "এলটিপিও সুপার রেটিনা এক্সডিআর ওলেড, 120Hz"',
    '"Dynamic AMOLED 2X, 120Hz, HDR10+": "Dynamic অ্যামোলেড 2X, 120Hz, এইচডিআর१०+"': '"Dynamic AMOLED 2X, 120Hz, HDR10+": "ডায়নামিক অ্যামোলেড 2X, 120Hz, এইচডিআর10+"',
    '"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "Foldable Dynamic এলটিপিও অ্যামোলেড 2X, 120Hz, এইচডিআর१०+"': '"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+": "ফোল্ডেবল ডায়নামিক এলটিপিও অ্যামোলেড 2X, 120Hz, এইচডিআর10+"',
    '"Super AMOLED, 90Hz, 800 nits": "Super অ্যামোলেড, 90Hz, 800 নিট"': '"Super AMOLED, 90Hz, 800 nits": "সুপার অ্যামোলেড, 90Hz, 800 নিট"',
    '"Super AMOLED, 120Hz": "Super অ্যামোলেড, 120Hz"': '"Super AMOLED, 120Hz": "সুপার অ্যামোলেড, 120Hz"',
    
    # Google Tensor
    '"Google Tensor G4": "Google টেনসর G4"': '"Google Tensor G4": "গুগল টেনসর G4"',
    '"Google Tensor G4 (4 nm)": "Google টেনসর G4 (4 nm)"': '"Google Tensor G4 (4 nm)": "গুগল টেনসর G4 (4 nm)"',
    '"Google Tensor G3": "Google টেনসর G3"': '"Google Tensor G3": "গুগল টেনসর G3"',
    '"Google Tensor G3 (4 nm)": "Google টেনসর G3 (4 nm)"': '"Google Tensor G3 (4 nm)": "গুগল টেনসর G3 (4 nm)"',
    
    # Snapdragon variations
    '"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"': '"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 জেন 3 (4 nm)"',
    '"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"': '"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 জেন 2 (4 nm)"',
    '"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 Elite (3 nm)"': '"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 এলিট (3 nm)"',
    '"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"': '"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 জেন 3 (4 nm)"',
    '"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"': '"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 জেন 2 (4 nm)"',
    '"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 Elite (3 nm)"': '"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 এলিট (3 nm)"',
    '"Qualcomm Snapdragon 695 5G (6 nm)": "Qualcomm স্ন্যাপড্রাগন 695 ৫জি (6 nm)"': '"Qualcomm Snapdragon 695 5G (6 nm)": "কোয়ালকম স্ন্যাপড্রাগন 695 ৫জি (6 nm)"',
    
    # Other brands
    '"Mecha Black, Amber Gold": "Mecha কালো, Amber সোনা"': '"Mecha Black, Amber Gold": "মেচা কালো, অ্যাম্বার সোনা"',
    '"Mediatek Dimensity 6080 (6 nm)": "মিডিয়াটেক Dimensity 6080 (6 nm)"': '"Mediatek Dimensity 6080 (6 nm)": "মিডিয়াটেক ডাইমেনশিটি 6080 (6 nm)"',
    '"Plastic body": "প্লাস্টিক body"': '"Plastic body": "প্লাস্টিক বডি"',
    '"Octa-core": "Octa-কোর"': '"Octa-core": "অকটা-কোর"',
}

def fix_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    original_content = content
    fixes_applied = 0
    
    for old, new in FIXES.items():
        if old in content:
            content = content.replace(old, new)
            fixes_applied += 1
    
    if content != original_content:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(content)
        return fixes_applied
    return 0

def main():
    print("=" * 70)
    print("FIXING ALL REMAINING MIXED ENGLISH-BENGALI PATTERNS")
    print("=" * 70)
    print()
    
    seeder_files = sorted(glob.glob(os.path.join(SEEDERS_DIR, "specification_seeder_mobile_*.go")))
    
    files_fixed = 0
    total_fixes = 0
    
    for filepath in seeder_files:
        fixes = fix_file(filepath)
        if fixes > 0:
            files_fixed += 1
            total_fixes += fixes
            fname = os.path.basename(filepath)
            print(f"  ✓ {fname}: {fixes} pattern(s) fixed")
    
    print()
    print("=" * 70)
    print(f"FILES FIXED: {files_fixed}/{len(seeder_files)}")
    print(f"TOTAL PATTERNS FIXED: {total_fixes}")
    print("=" * 70)

if __name__ == "__main__":
    main()
