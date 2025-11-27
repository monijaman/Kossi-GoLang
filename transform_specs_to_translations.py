#!/usr/bin/env python3
import json

print("Loading fixed mobile data...")

# Load the fixed JSON
with open("init-db/seeders/mobiledata_fixed.json", 'r', encoding='utf-8') as f:
    mobile_data = json.load(f)

print(f"Loaded {len(mobile_data)} devices\n")

# Extract all unique specification keys
all_spec_keys = set()
for device_name, specs in mobile_data.items():
    for key in specs.keys():
        all_spec_keys.add(key)

print(f"Found {len(all_spec_keys)} unique specification keys:")
for i, key in enumerate(sorted(all_spec_keys), 1):
    print(f"  {i}. {key}")

# Create the transformed output
# Format: device_name -> spec_key -> { english, specification_translations: [{locale, value}] }
output = {}

for device_name, specs in mobile_data.items():
    output[device_name] = {}
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

# Save the transformed output
output_path = "init-db/seeders/mobile_specification_translations.json"
with open(output_path, 'w', encoding='utf-8') as f:
    json.dump(output, f, indent=2, ensure_ascii=False)

print(f"\n✓ Transformed JSON saved to: {output_path}")
print(f"  Total devices: {len(mobile_data)}")

# Create a summary
summary = {
    "total_devices": len(mobile_data),
    "total_unique_specs": len(all_spec_keys),
    "specification_keys": sorted(list(all_spec_keys)),
    "sample_device": list(mobile_data.keys())[0],
    "sample_specs": list(list(mobile_data.values())[0].items())[:3]
}

summary_path = "mobile_specs_summary.json"
with open(summary_path, 'w', encoding='utf-8') as f:
    json.dump(summary, f, indent=2, ensure_ascii=False)

print(f"✓ Summary saved to: {summary_path}")

# Show sample of transformed output
print("\nSample of transformed structure:")
first_device = list(output.keys())[0]
first_spec = list(output[first_device].keys())[0]
print(f"\nDevice: {first_device}")
print(f"Spec: {first_spec}")
print(f"Transformed: {json.dumps(output[first_device][first_spec], indent=2)}")
