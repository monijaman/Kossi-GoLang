#!/usr/bin/env python3
import json
import re

# Load the mobiledata.json file
mobiledata_path = "init-db/seeders/mobiledata.json"

# Read the JSON file
with open(mobiledata_path, 'r', encoding='utf-8') as f:
    content = f.read()

# More comprehensive JSON fixing
# 1. Fix trailing commas before closing braces
content = re.sub(r',(\s*[}\]])', r'\1', content)

# 2. Fix missing commas - this is tricky, look for patterns like:
#    "value"\n\t"key" -> "value",\n\t"key"
# Only match if there's a quote after colon followed by newline/space and another quote
lines = content.split('\n')
fixed_lines = []

for i, line in enumerate(lines):
    fixed_lines.append(line)
    # Check if this line ends with a value and next line starts with a quote
    # Pattern: line ends with " and next line starts with whitespace and "
    if i < len(lines) - 1:
        current_stripped = line.rstrip()
        next_line = lines[i + 1] if i + 1 < len(lines) else ""
        next_stripped = next_line.lstrip()
        
        # If current line ends with quote/number and next line starts with quote
        # and current line doesn't end with comma
        if current_stripped and not current_stripped.endswith(',') and not current_stripped.endswith('{') and not current_stripped.endswith('['):
            if (current_stripped[-1] in ['"', '}', ']'] or current_stripped[-1].isdigit()) and next_stripped.startswith('"'):
                # Check if current line has a colon (it's a value, not a key)
                if ':' in current_stripped:
                    fixed_lines[-1] = current_stripped + ',' + line[len(current_stripped):]

content = '\n'.join(fixed_lines)

# Try to parse
try:
    mobile_data = json.loads(content)
    print("JSON fixed successfully!")
except json.JSONDecodeError as e:
    print(f"Still has errors: {e}")
    print(f"Line {e.lineno}, Column {e.colno}")
