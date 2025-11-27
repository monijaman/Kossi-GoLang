#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json

# Common value translations - simplified but comprehensive
TRANSLATIONS = {
    "Yes": "হ্যাঁ",
    "No": "না", 
    "Available": "উপলব্ধ",
    "Discontinued": "বন্ধ",
    "inches": "ইঞ্চি",
    "GB": "গিগাবাইট",
    "TB": "টেরাবাইট",
    "MP": "মেগাপিক্সেল",
    "mAh": "এমএএইচ",
    "W": "ওয়াট",
    "Hz": "হার্জ",
    "GHz": "গিগাহার্জ",
    "kg": "কেজি",
    "g": "গ্রাম",
    "mm": "মিমি",
}

def translate_to_bengali(text):
    if not text:
        return text
    result = text
    for eng, ben in TRANSLATIONS.items():
        result = result.replace(eng, ben)
    return result

print("Loading data...")
with open("init-db/seeders/mobile_specification_translations.json", 'r', encoding='utf-8') as f:
    data = json.load(f)

print(f"Updating Bengali for {len(data)} devices...")
updated = 0

for device in data.values():
    for spec in device.values():
        if isinstance(spec, dict) and "specification_translations" in spec:
            for trans in spec["specification_translations"]:
                if trans["locale"] == "bn":
                    trans["value"] = translate_to_bengali(spec["english"])
                    updated += 1

print(f"Updated {updated} translations")

with open("init-db/seeders/mobile_specification_translations.json", 'w', encoding='utf-8') as f:
    json.dump(data, f, indent=2, ensure_ascii=False)

print("Done!")
