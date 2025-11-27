#!/usr/bin/env python3
"""Fix all WiFi band character issues across all seeder files"""

import os
import glob
import re

SEEDERS_DIR = "init-db/seeders"

def fix_wifi_issues(content):
    """Fix WiFi band translation issues"""
    # Replace wrongly translated WiFi bands
    patterns = [
        (r'"ওয়াই-ফাই 802\.11 a/b/গ্রাম/', '"ওয়াই-ফাই 802.11 a/b/g/'),
        (r'"Wi-Fi 802\.11 a/b/গ্রাম/', '"Wi-Fi 802.11 a/b/g/'),
        (r'"Wi‑Fi 802\.11 a/b/গ্রাম/', '"Wi-Fi 802.11 a/b/g/'),
        (r'"802\.11 a/b/গ্রাম/', '"802.11 a/b/g/'),
    ]
    
    for pattern, replacement in patterns:
        content = re.sub(pattern, replacement, content)
    
    return content

def fix_file(filepath):
    """Fix WiFi issues in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        content = fix_wifi_issues(content)
        
        if content != original_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            return True
        return False
    except Exception as e:
        print(f"  ✗ Error: {e}")
        return False

def main():
    print("=" * 70)
    print("FIXING WIFI BAND CHARACTER ISSUES")
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
    print("WIFI BAND FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print("=" * 70)

if __name__ == "__main__":
    main()
