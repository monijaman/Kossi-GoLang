#!/usr/bin/env python3
import re

# Read the original file
mobiledata_path = "init-db/seeders/mobiledata.json"

with open(mobiledata_path, 'r', encoding='utf-8') as f:
    content = f.read()

print("Fixing JSON...")

# Step 1: Remove all trailing commas - be very aggressive
print("Step 1: Removing trailing commas...")
# Remove commas before closing braces (with various whitespace patterns)
content = re.sub(r',\s*}', '}', content)
# Remove commas before closing brackets
content = re.sub(r',\s*]', ']', content)
# Do it again to catch nested cases
content = re.sub(r',\s*}', '}', content)
content = re.sub(r',\s*]', ']', content)

# Step 2: Add missing commas between key-value pairs
print("Step 2: Adding missing commas between properties...")

lines = content.split('\n')
output_lines = []

for i, line in enumerate(lines):
    current = line.rstrip()
    
    if current and i < len(lines) - 1:
        # Look ahead to next non-empty line
        next_non_empty_idx = None
        for j in range(i + 1, len(lines)):
            if lines[j].strip():
                next_non_empty_idx = j
                break
        
        if next_non_empty_idx is not None:
            next_line = lines[next_non_empty_idx].lstrip()
            
            # If current line could need a comma
            if not current.endswith(',') and not current.endswith('{') and not current.endswith('['):
                # And next line starts with a quote (new key) or closing brace
                if next_line and next_line[0] == '"':
                    # And current line has a colon (it's a key-value)
                    if ':' in current and '":' in current:
                        output_lines.append(current + ',')
                    else:
                        output_lines.append(line)
                else:
                    output_lines.append(line)
            else:
                output_lines.append(line)
        else:
            output_lines.append(line)
    else:
        output_lines.append(line)

content = '\n'.join(output_lines)

# Save the fixed content
output_path = "init-db/seeders/mobiledata_fixed.json"
with open(output_path, 'w', encoding='utf-8') as f:
    f.write(content)

print(f"Fixed JSON saved to {output_path}")

# Try to parse it
import json
try:
    data = json.loads(content)
    print(f"\n✓ Successfully parsed! Found {len(data)} devices")
    
    # List first few devices
    devices = list(data.keys())[:5]
    print("First few devices:", devices)
    
except json.JSONDecodeError as e:
    print(f"\n✗ JSON Error at line {e.lineno}, column {e.colno}: {e.msg}")
    # Print context around error
    error_pos = e.pos
    start = max(0, error_pos - 150)
    end = min(len(content), error_pos + 150)
    print("\nContext:")
    print("..." + content[start:end] + "...")
    print(" " * (error_pos - start + 3) + "^")
