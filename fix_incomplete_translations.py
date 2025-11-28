import os
import re
import glob

# Comprehensive list of common English words/phrases that need Bangla translations
common_translations = {
    # Brands (keep as is, no translation needed)
    "Google": "Google",
    "Apple": "Apple",
    "Samsung": "Samsung",
    "Xiaomi": "Xiaomi",
    "OnePlus": "OnePlus",
    "Oppo": "Oppo",
    "Vivo": "Vivo",
    "Realme": "Realme",
    "Infinix": "Infinix",
    "Tecno": "Tecno",
    "Poco": "Poco",
    "Redmi": "Redmi",
    "iPhone": "iPhone",
    "Galaxy": "Galaxy",
    "Pixel": "Pixel",
    "Tensor": "Tensor",
    "Snapdragon": "Snapdragon",
    "MediaTek": "MediaTek",
    "Mali": "Mali",
    "Adreno": "Adreno",
    "Corning": "Corning",
    "Gorilla Glass": "Gorilla Glass",
    "Victus": "Victus",
    
    # Months
    "January": "জানুয়ারি",
    "February": "ফেব্রুয়ারি",
    "March": "মার্চ",
    "April": "এপ্রিল",
    "May": "মে",
    "June": "জুন",
    "July": "জুলাই",
    "August": "আগস্ট",
    "September": "সেপ্টেম্বর",
    "October": "অক্টোবর",
    "November": "নভেম্বর",
    "December": "ডিসেম্বর",
    
    # Common words
    "Available": "উপলব্ধ",
    "Announced": "ঘোষিত",
    "Released": "মুক্তি",
    "Nona-core": "নোনা-কোর",
    "Octa-core": "অক্টা-কোর",
    "Hexa-core": "হেক্সা-কোর",
    "Quad-core": "কোয়াড-কোর",
    "Dual-core": "ডুয়াল-কোর",
    
    # Colors
    "Obsidian": "অবসিডিয়ান",
    "Hazel": "হ্যাজেল",
    "Rose": "রোজ",
    "Mint": "মিন্ট",
    "Porcelain": "পোর্সেলিন",
    "Wintergreen": "উইন্টারগ্রীন",
    "Peony": "পিওনি",
    "Black": "কালো",
    "White": "সাদা",
    "Blue": "নীল",
    "Gray": "ধূসর",
    "Green": "সবুজ",
    "Pink": "গোলাপী",
    "Red": "লাল",
    "Silver": "রূপালী",
    "Gold": "সোনালী",
    
    # Common phrases with "Available"
    "Available. Released": "উপলব্ধ। মুক্তি",
}

# Find all mobile seeder files
seeder_files = glob.glob("internal/infrastructure/database/seeders/specification_seeder_mobile_*.go")

print(f"Found {len(seeder_files)} mobile seeder files to check")

updated_count = 0
for filepath in seeder_files:
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Find the getBanglaTranslations function
    pattern = r'func \(s \*\w+\) getBanglaTranslations\(\) map\[string\]string \{[\s\S]*?return map\[string\]string\{([\s\S]*?)\n\t\}\n\}'
    match = re.search(pattern, content)
    
    if not match:
        continue
    
    translations_block = match.group(1)
    existing_translations = {}
    
    # Parse existing translations
    trans_pattern = r'"([^"]+)":\s*"([^"]+)"'
    for m in re.finditer(trans_pattern, translations_block):
        existing_translations[m.group(1)] = m.group(2)
    
    # Check if any values need translation
    needs_update = False
    for eng_value, bangla_value in existing_translations.items():
        # Check if the translation contains English words that should be translated
        for eng_word, bengali_word in common_translations.items():
            if eng_word in eng_value and eng_word not in existing_translations:
                # This value contains a word that should have been translated
                if eng_value == eng_word:
                    # Direct match, add translation
                    existing_translations[eng_word] = bengali_word
                    needs_update = True
    
    # Add missing common translations
    for eng, ben in common_translations.items():
        if eng not in existing_translations:
            existing_translations[eng] = ben
            needs_update = True
    
    if needs_update:
        # Rebuild the translations map
        new_translations = []
        for key in sorted(existing_translations.keys()):
            value = existing_translations[key]
            new_translations.append(f'\t\t"{key}": "{value}",')
        
        new_translations_str = "\n".join(new_translations)
        
        # Replace the old translations block
        new_function = f'''func (s *{re.search(r'func \\(s \\*(\\w+)\\) getBanglaTranslations', content).group(1)}) getBanglaTranslations() map[string]string {{
\treturn map[string]string{{
{new_translations_str}
\t}}
}}'''
        
        new_content = re.sub(pattern, new_function, content)
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(new_content)
        
        print(f"✓ Updated: {os.path.basename(filepath)}")
        updated_count += 1

print(f"\n✅ Updated {updated_count} files with missing translations!")
