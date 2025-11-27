#!/usr/bin/env python3
import json
import re
import sys

try:
    # Load the mobiledata.json file
    mobiledata_path = "init-db/seeders/mobiledata.json"
    
    # Read the JSON file
    with open(mobiledata_path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Fix various JSON syntax issues
    # 1. Fix trailing commas: replace ,\n} with \n}
    content = re.sub(r',\s*}', '}', content)
    # 2. Fix trailing commas: replace ,\n] with \n]
    content = re.sub(r',\s*]', ']', content)
    # 3. Fix missing commas between key-value pairs followed by more pairs
    # Match pattern: "key": "value"\n\t"key2" -> "key": "value",\n\t"key2"
    content = re.sub(r'(":\s*"[^"]*")\n(\s*")', r'\1,\n\2', content)
    content = re.sub(r'(":\s*"[^"]*")\r\n(\s*")', r'\1,\n\2', content)
    
    mobile_data = json.loads(content)
    
    # Extract all unique specification keys from all devices
    all_spec_keys = set()
    for device_name, specs in mobile_data.items():
        if isinstance(specs, dict):
            for key in specs.keys():
                all_spec_keys.add(key)
    
    print(f"Found {len(all_spec_keys)} unique specification keys")
    print("\nSpecification keys found:")
    for i, key in enumerate(sorted(all_spec_keys), 1):
        print(f"  {i}. {key}")
    
    # Create the output structure with specification_translations
    output = {}
    for device_name, specs in mobile_data.items():
        output[device_name] = {}
        if isinstance(specs, dict):
            for spec_key, spec_value in specs.items():
                output[device_name][spec_key] = {
                    "english": spec_value,
                    "specification_translations": [
                        {
                            "locale": "en",
                            "value": spec_value
                        },
                        {
                            "locale": "bn",
                            "value": f"[Bengali] {spec_value}"
                        }
                    ]
                }
    
    # Save the output
    output_path = "init-db/seeders/mobile_specification_translations.json"
    with open(output_path, 'w', encoding='utf-8') as f:
        json.dump(output, f, indent=2, ensure_ascii=False)
    
    print(f"\n✅ Transformed JSON saved to: {output_path}")
    print(f"Total devices processed: {len(mobile_data)}")
    
    # Also create a summary of all spec keys
    summary = {
        "total_devices": len(mobile_data),
        "total_unique_specs": len(all_spec_keys),
        "specification_keys": sorted(list(all_spec_keys))
    }
    
    summary_path = "mobile_specs_summary.json"
    with open(summary_path, 'w', encoding='utf-8') as f:
        json.dump(summary, f, indent=2, ensure_ascii=False)
    
    print(f"✅ Summary saved to: {summary_path}")

except Exception as e:
    print(f"❌ Error: {e}")
    import traceback
    traceback.print_exc()
    sys.exit(1)
