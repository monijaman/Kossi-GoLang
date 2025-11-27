#!/bin/bash

# Fix all ASCII numerals to Bengali numerals in translation values across all seeder files

cd init-db/seeders

echo "=== CONVERTING ALL ASCII NUMERALS TO BENGALI DIGITS ===" 
echo ""

count=0
total=$(ls specification_seeder_mobile_*.go | wc -l)

for file in specification_seeder_mobile_*.go; do
    # Skip if no changes needed
    if ! grep -q '": "[0-9]' "$file"; then
        continue
    fi
    
    # Convert ASCII digits 0-9 to Bengali digits ০-৯
    # Using sed to replace in the entire file
    sed -i \
        -e 's/": "\([^"]*\)0\([^"]*\)"/": "\1০\2"/g' \
        -e 's/": "\([^"]*\)1\([^"]*\)"/": "\1১\2"/g' \
        -e 's/": "\([^"]*\)2\([^"]*\)"/": "\1२\2"/g' \
        -e 's/": "\([^"]*\)3\([^"]*\)"/": "\1३\2"/g' \
        -e 's/": "\([^"]*\)4\([^"]*\)"/": "\1४\2"/g' \
        -e 's/": "\([^"]*\)5\([^"]*\)"/": "\1५\2"/g' \
        -e 's/": "\([^"]*\)6\([^"]*\)"/": "\1६\2"/g' \
        -e 's/": "\([^"]*\)7\([^"]*\)"/": "\1७\2"/g' \
        -e 's/": "\([^"]*\)8\([^"]*\)"/": "\1८\2"/g' \
        -e 's/": "\([^"]*\)9\([^"]*\)"/": "\1९\2"/g' \
        "$file"
    
    ((count++))
done

echo "✓ Converted ASCII numerals to Bengali digits in $count/$total files"
