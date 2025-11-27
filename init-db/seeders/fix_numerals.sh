#!/bin/bash

# Mapping function for digit conversion
convert_file() {
    local file="$1"
    
    # Use sed to convert ASCII digits to Bengali numerals in values only
    # Pattern: "KEY": "VALUE" - we want to replace digits only in VALUE
    sed -i '
    s/": "\([^"]*\)0\([^"]*\)"/": "\1०\2"/g
    s/": "\([^"]*\)1\([^"]*\)"/": "\1१\2"/g
    s/": "\([^"]*\)2\([^"]*\)"/": "\1२\2"/g
    s/": "\([^"]*\)3\([^"]*\)"/": "\1३\2"/g
    s/": "\([^"]*\)4\([^"]*\)"/": "\1४\2"/g
    s/": "\([^"]*\)5\([^"]*\)"/": "\1५\2"/g
    s/": "\([^"]*\)6\([^"]*\)"/": "\1६\2"/g
    s/": "\([^"]*\)7\([^"]*\)"/": "\1७\2"/g
    s/": "\([^"]*\)8\([^"]*\)"/": "\1८\2"/g
    s/": "\([^"]*\)9\([^"]*\)"/": "\1९\2"/g
    ' "$file"
}

total=0
fixed=0

for file in specification_seeder_mobile_*.go; do
    total=$((total + 1))
    convert_file "$file"
    fixed=$((fixed + 1))
done

echo "Converted $fixed / $total files"
