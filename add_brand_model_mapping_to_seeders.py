#!/usr/bin/env python3
"""Add Brand/Model Name key mappings to car/spec seeders that use existingKeyMapping.

This ensures keys "Brand" (310) and "Model Name" (316) are present
in the existingKeyMapping maps so the specs we already added actually save.
"""

import os

BASE_DIR = "internal/infrastructure/database/seeders"

def add_mapping(filepath: str) -> bool:
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            content = f.read()

        # Only process files that define existingKeyMapping
        marker = "existingKeyMapping := map[string]uint{\n"
        if marker not in content:
            return False

        # If Brand/Model Name already mapped inside existingKeyMapping block, skip
        start = content.index(marker)
        # Heuristic: mapping block ends at the first "}\n" after marker
        end = content.find("}\n", start)
        mapping_block = content[start:end] if end != -1 else content[start:]
        if '"Brand":' in mapping_block and '"Model Name":' in mapping_block:
            return False

        insertion = (
            marker
            + '\t\t"Brand":                             310,\n'
            + '\t\t"Model Name":                        316,\n'
        )

        new_content = content.replace(marker, insertion, 1)
        if new_content == content:
            return False

        with open(filepath, "w", encoding="utf-8") as f:
            f.write(new_content)

        print(f"[OK] Mapping added in {filepath}")
        return True
    except FileNotFoundError:
        return False


def main() -> None:
    updated = 0
    skipped = 0

    for name in os.listdir(BASE_DIR):
        if not name.endswith("_seeder.go"):
            continue
        path = os.path.join(BASE_DIR, name)
        if add_mapping(path):
            updated += 1
        else:
            skipped += 1

    print("=" * 60)
    print(f"Mapping update summary: {updated} files updated, {skipped} skipped")
    print("=" * 60)


if __name__ == "__main__":
    main()
