#!/bin/bash

# List of Panasonic models
models=(
"NR‑CY550QKXZ"
"NR‑BJ226SNSG"
"NR‑BK265S/266"
"NR‑BR347Z"
"NR‑BK305SE"
"NR‑BW415V"
"NR‑F605TT"
"NR‑BT262M"
"NR‑F555TX"
"NR‑BX468XGX3"
"NR-D513"
)

# Template file (using the Walton one as base)
template="specification_seeder_refrigerator_walton-wut-3x8-gere-cx.go"

for model in "${models[@]}"; do
    # Generate slug
    slug="panasonic-$(echo "$model" | tr '[:upper:]' '[:lower:]' | sed 's/‑/-/g' | sed 's/\//-/g')"
    
    # Generate struct name
    # Remove panasonic- and capitalize parts
    struct_slug=$(echo "$slug" | sed 's/panasonic-//')
    IFS='-' read -ra parts <<< "$struct_slug"
    struct_parts=()
    for part in "${parts[@]}"; do
        struct_parts+=("$(echo $part | sed 's/./\U&/')"
    done
    struct_name="SpecificationSeederRefrigeratorPanasonic$(echo "${struct_parts[@]}" | sed 's/ //g')"
    func_name="New${struct_name}"
    file_name="specification_seeder_refrigerator_${slug}.go"
    
    echo "Processing $slug -> $file_name"
    
    # Copy template
    cp "$template" "$file_name"
    
    # Replace in file
    sed -i "s/SpecificationSeederRefrigeratorWaltonWut3x8GereCx/${struct_name}/g" "$file_name"
    sed -i "s/NewSpecificationSeederRefrigeratorWaltonWut3x8GereCx/${func_name}/g" "$file_name"
    sed -i "s/walton-wut-3x8-gere-cx/${slug}/g" "$file_name"
    sed -i "s/WUT-3X8-GERE-CX/${model}/g" "$file_name"
    sed -i "s/Walton/Panasonic/g" "$file_name"
    sed -i "s/ওয়ালটন/প্যানাসনিক/g" "$file_name"
    
    echo "Generated $file_name"
done

echo "Generated ${#models[@]} Panasonic seeder files."
