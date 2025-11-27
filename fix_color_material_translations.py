#!/usr/bin/env python3
"""Fix ALL remaining mixed English/Bengali translations in specifications"""

import os
import glob
import re

SEEDERS_DIR = "init-db/seeders"

# Complete dictionary of mixed translations to fix
MIXED_FIXES = {
    # iPhone colors - mixed
    '"Black Titanium, White Titanium, Natural Titanium, Desert Titanium": "কালো Titanium, সাদা Titanium, Natural Titanium, Desert Titanium"': 
    '"Black Titanium, White Titanium, Natural Titanium, Desert Titanium": "কালো Titanium, সাদা Titanium, প্রাকৃতিক Titanium, মরুভূমি Titanium"',
    
    '"Black Titanium, White Titanium, Blue Titanium, Natural Titanium": "কালো Titanium, সাদা Titanium, নীল Titanium, Natural Titanium"':
    '"Black Titanium, White Titanium, Blue Titanium, Natural Titanium": "কালো Titanium, সাদা Titanium, নীল Titanium, প্রাকৃতিক Titanium"',
    
    # iPhone colors with unmixed colors
    '"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, Yellow, Pink"':
    '"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপি"',
    
    '"Black, White, Blue": "কালো, সাদা, নীল"': 
    '"Black, White, Blue": "কালো, সাদা, নীল"',  # Already correct
    
    # Rose Gold and similar iPhone colors
    '"Black, White, Gold, Rose Gold, Blue, Deep Purple": "কালো, সাদা, সোনা, Rose সোনা, নীল, Deep Purple"':
    '"Black, White, Gold, Rose Gold, Blue, Deep Purple": "কালো, সাদা, সোনা, গোলাপি সোনা, নীল, গভীর বেগুনি"',
    
    # iPhone with unmixed colors needing fixes
    '"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, Yellow, Pink"':
    '"Black, Blue, Green, Yellow, Pink": "কালো, নীল, সবুজ, হলুদ, গোলাপি"',
    
    # Titanium frame mixed
    '"Titanium frame, glass front/back": "Titanium ফ্রেম, গ্লাস সামনে/পিছনে"':
    '"Titanium frame, glass front/back": "Titanium ফ্রেম, গ্লাস সামনে/পিছনে"',  # Already correct
    
    # Color names with descriptors
    '"Midnight Black, White Titanium, Desert Titanium": "Midnight কালো, সাদা Titanium, Desert Titanium"':
    '"Midnight Black, White Titanium, Desert Titanium": "মধ্যরাত কালো, সাদা Titanium, মরুভূমি Titanium"',
    
    # OnePlus colors
    '"Flowy Emerald, Silky Black, Silver": "Flowy Emerald, Silky কালো, রূপা"':
    '"Flowy Emerald, Silky Black, Silver": "Flowy Emerald, রেশমী কালো, রূপা"',
    
    '"Midnight Ocean, Black Eclipse, Arctic Dawn, White Dew (China only)": "Midnight Ocean, কালো Eclipse, Arctic Dawn, সাদা Dew (China only)"':
    '"Midnight Ocean, Black Eclipse, Arctic Dawn, White Dew (China only)": "মধ্যরাত Ocean, কালো Eclipse, Arctic শুরু, সাদা Dew (শুধুমাত্র চীন)"',
    
    '"Mercurial Silver, Oasis Green, Obsidian Midnight": "Mercurial রূপা, Oasis সবুজ, Obsidian Midnight"':
    '"Mercurial Silver, Oasis Green, Obsidian Midnight": "Mercurial রূপা, ঘাসের মাঠ সবুজ, Obsidian মধ্যরাত"',
    
    # Realme colors
    '"Twilight Purple, Woodland Green": "Twilight Purple, Woodland সবুজ"':
    '"Twilight Purple, Woodland Green": "Twilight বেগুনি, Woodland সবুজ"',
    
    # Samsung Galaxy colors
    '"Thunder Black, Mint Green, Blush Pink": "Thunder কালো, Mint সবুজ, Blush Pink"':
    '"Thunder Black, Mint Green, Blush Pink": "থান্ডার কালো, পুদিনা সবুজ, Blush গোলাপি"',
    
    '"Phantom Black, Green, Cream, Lavender": "Phantom কালো, সবুজ, Cream, Lavender"':
    '"Phantom Black, Green, Cream, Lavender": "Phantom কালো, সবুজ, ক্রিম, ল্যাভেন্ডার"',
    
    '"Titanium Gray, Titanium Black, Titanium Violet, Titanium Yellow": "Titanium ধূসর, Titanium কালো, Titanium Violet, Titanium Yellow"':
    '"Titanium Gray, Titanium Black, Titanium Violet, Titanium Yellow": "Titanium ধূসর, Titanium কালো, Titanium বেগুনি, Titanium হলুদ"',
    
    # More Samsung colors
    '"Black, Green, Blue": "কালো, সবুজ, নীল"':
    '"Black, Green, Blue": "কালো, সবুজ, নীল"',  # Already correct
    
    '"Midnight Black, Aurora Purple, Ocean Teal": "Midnight কালো, Aurora Purple, Ocean Teal"':
    '"Midnight Black, Aurora Purple, Ocean Teal": "মধ্যরাত কালো, Aurora বেগুনি, Ocean Teal"',
    
    # Color descriptors
    '"Onyx Gray, Mint Green, Ice Blue": "Onyx ধূসর, Mint সবুজ, Ice নীল"':
    '"Onyx Gray, Mint Green, Ice Blue": "Onyx ধূসর, পুদিনা সবুজ, বরফ নীল"',
    
    '"Graphite Gray, Coral Green, Sky Blue": "Graphite ধূসর, Coral সবুজ, Sky নীল"':
    '"Graphite Gray, Coral Green, Sky Blue": "গ্রাফাইট ধূসর, প্রবাল সবুজ, আকাশ নীল"',
    
    # Additional descriptors
    '"Icy Blue, Phantom Black, Cream, Gray, Blue": "Icy নীল, Phantom কালো, Cream, ধূসর, নীল"':
    '"Icy Blue, Phantom Black, Cream, Gray, Blue": "বরফ নীল, Phantom কালো, ক্রিম, ধূসর, নীল"',
    
    '"Space Black, Titanium Silver, Cosmic Gold, Ice Blue, Radium Green": "Space কালো, Titanium রূপা, Cosmic সোনা, Ice নীল, Radium সবুজ"':
    '"Space Black, Titanium Silver, Cosmic Gold, Ice Blue, Radium Green": "মহাকাশ কালো, Titanium রূপা, মহাজাগতিক সোনা, বরফ নীল, Radium সবুজ"',
    
    '"Sandy Purple, Midnight Black, Titan Gray": "Sandy Purple, Midnight কালো, Titan ধূসর"':
    '"Sandy Purple, Midnight Black, Titan Gray": "বালুকাময় বেগুনি, মধ্যরাত কালো, Titan ধূসর"',
    
    # Material mixed
    '"Titanium frame, glass front/back": "Titanium ফ্রেম, গ্লাস সামনে/পিছনে"':
    '"Titanium frame, glass front/back": "Titanium ফ্রেম, গ্লাস সামনে/পিছনে"',  # Already correct
    
    # Xiaomi colors
    '"Racing Black, Surfing Green": "Racing কালো, Surfing সবুজ"':
    '"Racing Black, Surfing Green": "রেসিং কালো, সার্ফিং সবুজ"',
    
    '"Timber Black, Shiny Gold": "Timber কালো, Shiny সোনা"':
    '"Timber Black, Shiny Gold": "টাইম্বার কালো, চকচকে সোনা"',
    
    '"Titan Black, Eternal Green": "Titan কালো, Eternal সবুজ"':
    '"Titan Black, Eternal Green": "Titan কালো, চিরস্থায়ী সবুজ"',
    
    # General brand + English color
    '"Black Satin, Cyan Sparkle": "কালো Satin, Cyan Sparkle"':
    '"Black Satin, Cyan Sparkle": "কালো সাটিন, Cyan Sparkle"',
    
    '"Blue Oasis, Celestial Black": "নীল Oasis, Celestial কালো"':
    '"Blue Oasis, Celestial Black": "নীল Oasis, Celestial কালো"',
    
    '"Prism Black, Space Silver": "Prism কালো, Space রূপা"':
    '"Prism Black, Space Silver": "Prism কালো, Space রূপা"',
    
    # More text mixed
    '"Mint, Graphite, Cream, Lavender, Gray, Blue, Green, Yellow": "Mint, Graphite, Cream, Lavender, ধূসর, নীল, সবুজ, Yellow"':
    '"Mint, Graphite, Cream, Lavender, Gray, Blue, Green, Yellow": "পুদিনা, Graphite, ক্রিম, ল্যাভেন্ডার, ধূসর, নীল, সবুজ, হলুদ"',
    
    # More complex descriptors
    '"Radium Green, Cosmic Gold, Honey Dew Green": "Radium সবুজ, Cosmic সোনা, Honey Dew সবুজ"':
    '"Radium Green, Cosmic Gold, Honey Dew Green": "Radium সবুজ, মহাজাগতিক সোনা, মধু শিরা সবুজ"',
    
    '"Luxe Violet, Coral Red, Onyx Black": "Luxe Violet, Coral লাল, Onyx কালো"':
    '"Luxe Violet, Coral Red, Onyx Black": "বিলাসবহুল Violet, প্রবাল লাল, Onyx কালো"',
    
    '"Meteorite Grey, Comet Green": "Meteorite ধূসর, Comet সবুজ"':
    '"Meteorite Grey, Comet Green": "Meteorite ধূসর, ধুমকেতু সবুজ"',
    
    '"Imperial Purple, Cosmic Gold, Titanium Gray, Graphite Black, Swamp Green": "Imperial Purple, Cosmic সোনা, Titanium ধূসর, Graphite কালো, Swamp সবুজ"':
    '"Imperial Purple, Cosmic Gold, Titanium Gray, Graphite Black, Swamp Green": "সাম্রাজ্য বেগুনি, মহাজাগতিক সোনা, Titanium ধূসর, Graphite কালো, জলাভূমি সবুজ"',
    
    '"Sprint Blue, Nitro Blue, Shade Black": "Sprint নীল, Nitro নীল, Shade কালো"':
    '"Sprint Blue, Nitro Blue, Shade Black": "Sprint নীল, Nitro নীল, ছায়া কালো"',
    
    # Stealth, Arctic, Prism mixed
    '"Stealth Black, Arctic White, Prism Gold": "Stealth কালো, Arctic সাদা, Prism সোনা"':
    '"Stealth Black, Arctic White, Prism Gold": "Stealth কালো, আর্কটিক সাদা, Prism সোনা"',
    
    # Brave, Optimistic mixed
    '"Brave Black, Optimistic Blue, Magical Blue, Personality Yellow": "Brave কালো, Optimistic নীল, Magical নীল, Personality Yellow"':
    '"Brave Black, Optimistic Blue, Magical Blue, Personality Yellow": "সাহসী কালো, আশাবাদী নীল, জাদুকরী নীল, ব্যক্তিত্ব হলুদ"',
    
    # Glass and material mixed
    '"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame": "গ্লাস সামনে (গরিলা Armor 2), গ্লাস পিছনে (ভিকটাস 2), titanium ফ্রেম"':
    '"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame": "গ্লাস সামনে (Gorilla Armor 2), গ্লাস পিছনে (Victus 2), titanium ফ্রেম"',
}

def fix_file(filepath):
    """Fix mixed translations in a single file"""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        original_content = content
        fixed_count = 0
        
        for old_text, new_text in MIXED_FIXES.items():
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
    print("FIXING ALL REMAINING MIXED COLOR & MATERIAL TRANSLATIONS")
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
    print("MIXED COLOR & MATERIAL TRANSLATION FIX COMPLETE")
    print("=" * 70)
    print(f"Total files: {total_count}")
    print(f"Files fixed: {fixed_count}")
    print(f"Fix patterns applied: {len(MIXED_FIXES)}")
    print("=" * 70)

if __name__ == "__main__":
    main()
