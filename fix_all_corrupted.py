#!/usr/bin/env python3
import subprocess
import os

# Get list of all corrupted files
result = subprocess.run(
    ['bash', '-c', 'cd /d/GO/gocrit/gocrit_server && go build ./internal/infrastructure/database/seeders/... 2>&1 | grep "unexpected EOF" | grep -o "specification_seeder_refrigerator_ecoplus_[^:]*" | sort -u'],
    capture_output=True,
    text=True
)

corrupted_files = [f.strip() for f in result.stdout.strip().split('\n') if f.strip()]

fixed = 0
for filename in corrupted_files:
    filepath = f"d:/GO/gocrit/gocrit_server/internal/infrastructure/database/seeders/{filename}"
    
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Add a closing brace
        content += "}\n"
        
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(content)
        
        fixed += 1
        print(f"✓ Fixed {filename}")
    except Exception as e:
        print(f"✗ Error in {filename}: {e}")

print(f"\n✓ Fixed {fixed} files")
