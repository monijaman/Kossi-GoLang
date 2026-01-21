#!/bin/bash

# List of slugs
slugs=(
"walton-wnm-2a7-gdel-xx"
"walton-wnm-2f1-gehe-xx"
"walton-wfa-2a3-gdel-sc"
"walton-wfa-2d4-gdel-xx"
"walton-wfb-1n3-gdxx-xx"
"walton-wfb-2x1-gdel-xx"
"walton-wfb-2e4-gdel-xx"
"walton-wfe-3e8-gden-xx"
"walton-wfd-1b6-gdeh-xx"
"walton-wfd-1d4-gdeh-xx"
"walton-wfd-1b6-gdel-xx"
"walton-wfd-1d4-gdel-xx"
"walton-wfd-1f3-gdel-xx"
"walton-wnr-6d6-gdfs-dd"
"walton-wnr-6f0-scrc-co"
"walton-wcg-3j0-rxlx-xx"
"walton-wcg-3j0-gddb-xx"
"walton-wcg-3j0-ddge-xx"
"walton-wcf-2a0-gsre-xx"
"walton-wfb-2b6-elxx-xx"
"walton-wfb-2x1-gdsh-xx"
"walton-wnj-5h5-rxxx-xx"
"walton-wnm-1n5-rxxx-rp"
"walton-wfb-1h5-gdxx-xx"
"walton-wfb-2e0-gdxx-xx"
"walton-wfe-3a2-gdel-xx"
"walton-wfc-3d8-gdeh-dd"
"walton-wcg-2e5-gdel-dd-inverter"
"walton-wcg-2e5-gdel-dd"
"walton-wcg-2e5-ehlx-xx"
"walton-wcg-2e5-ehlc-xx"
"walton-wcf-2a0-gsre-xx-p"
"walton-wcf-2t5-rrlx-gx"
"walton-wcf-2t5-rrlx-xx"
"walton-wcf-2t5-gdel-xx"
"walton-wcf-1d5-rrxx-xx"
"walton-wcf-1d5-gele-dl"
"walton-wcf-1d5-gdel-xx"
"walton-wcf-1d5-gdel-lx"
"walton-wue-3c4-gepb-xx"
"walton-wut-3x8-gere-cx"
"walton-wue-2g2-gepb-xx"
"walton-wfa-1n3-gdxx-xx"
"walton-wfb-2e0-gjxb-sx-p"
"walton-wfe-2h2-gdxx-xx"
"walton-wfc-3f5-gdeh-xx"
"walton-wnh-3h6-gdel-xx"
"walton-wni-5f3-gdne-id"
"walton-wni-5f3-gdel-dd"
"walton-wni-6a9-gdne-dd"
"walton-wni-6a9-gdsd-dd"
)

# Template file
template="specification_seeder_refrigerator_walton-wcf-1b5-gdel-xx.go"

for slug in "${slugs[@]}"; do
    echo "Processing $slug"
    # Generate struct name: capitalize each part
    # Split by -
    IFS='-' read -ra parts <<< "$slug"
    struct_parts=()
    for part in "${parts[@]}"; do
        # Capitalize first letter
        struct_parts+=("$(echo $part | sed 's/./\U&/')"
    done
    struct_name="SpecificationSeederRefrigerator$(echo "${struct_parts[@]}" | sed 's/ //g')"
    func_name="New${struct_name}"
    file_name="specification_seeder_refrigerator_${slug}.go"
    
    # Copy template
    cp "$template" "$file_name"
    
    # Replace in file
    sed -i "s/SpecificationSeederRefrigeratorWaltonWcf1b5GdelXx/${struct_name}/g" "$file_name"
    sed -i "s/NewSpecificationSeederRefrigeratorWaltonWcf1b5GdelXx/${func_name}/g" "$file_name"
    sed -i "s/walton-wcf-1b5-gdel-xx/${slug}/g" "$file_name"
done

echo "Generated ${#slugs[@]} seeder files."
