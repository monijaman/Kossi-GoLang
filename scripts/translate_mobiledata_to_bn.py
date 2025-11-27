import json

# Example mapping for demonstration; extend as needed or use an API for full translation
en_bn_map = {
    "inches": "ইঞ্চি",
    "GB": "জিবি",
    "Black": "কালো",
    "White": "সাদা",
    "Gold": "সোনা",
    "Blue": "নীল",
    "Green": "সবুজ",
    "Pink": "গোলাপি",
    "Available": "উপলব্ধ",
    "Yes": "হ্যাঁ",
    "No": "না",
    "Android": "অ্যান্ড্রয়েড",
    "iOS": "আইওএস",
    # Add more mappings as needed
}

def translate_value(val):
    # Simple word replacement; for production use, integrate with a translation API
    for en, bn in en_bn_map.items():
        val = val.replace(en, bn)
    return val

def translate_json(input_path, output_path):
    with open(input_path, 'r', encoding='utf-8') as f:
        data = json.load(f)
    for device, specs in data.items():
        for k, v in specs.items():
            if isinstance(v, str):
                data[device][k] = translate_value(v)
    with open(output_path, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)

if __name__ == "__main__":
    translate_json('i:/GO/kossti/gocrit_server/init-db/seeders/mobiledata.json', 'i:/GO/kossti/gocrit_server/init-db/seeders/mobiledata-bn-translated.json')
