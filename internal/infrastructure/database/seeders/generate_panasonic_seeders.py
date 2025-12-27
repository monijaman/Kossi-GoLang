import os
import re

# List of Panasonic models
models = [
    "NR-TG368CPKN",
    "NR-TG368CPAN",
    "NR-TG366CVHN",
    "NR-TG368BPAN",
    "NR-TG368BVHN",
    "NR-TG338CPKN",
    "NR-TG338CVHN",
    "NR-TG338BPAN",
    "NR-TG338BVHN",
    "NR-TG336BVHN",
    "NR-TG292BVHN",
    "NR-BK418BGMN",
    "NR-BK468BGMN",
    "NR-BK415BQKN",
    "NR-BK465BQKN",
    "NR-BK418BQKN",
    "NR-BK468BQKN",
    "NR-TH292CVHN",
    "NR-TH292BVHN",
    "NR-TH272CVHN",
    "NR-TG352BUSN",
    "NR-TG322BVHN",
    "NR-TH292BUHN",
    "NR-TH292BPAN",
    "NR-TH292BPRN",
    "NR-TH292BDRN",
    "NR-TH292CPKN",
    "NR-TH292CDAN",
    "NR-TH272CPRN",
    "NR-TH272CPAN",
    "NR-TH272CPKN",
    "NR-TH272CDAN",
    "NR-TH272CDRN",
    "A202BURN",
    "A202BWHN",
    "A202BFAN",
    "A222BWHN",
    "A222BFRN",
    "A242CFRN",
    "A242CFAN"
]

# Template file
template_file = "specification_seeder_refrigerator_walton-wut-3x8-gere-cx.go"

# Read template
with open(template_file, 'r', encoding='utf-8') as f:
    template = f.read()

for model in models:
    # Generate slug
    slug = "panasonic-" + model.lower().replace('‑', '-').replace('/', '-')
    
    # Generate struct name
    struct_slug = slug.replace('panasonic-', '')
    parts = struct_slug.split('-')
    struct_parts = [part.capitalize() for part in parts]
    struct_name = "SpecificationSeederRefrigeratorPanasonic" + ''.join(struct_parts)
    func_name = "New" + struct_name
    file_name = f"specification_seeder_refrigerator_{slug}.go"
    
    print(f"Processing {slug} -> {file_name}")
    
    # Replace in template
    content = template
    content = content.replace("SpecificationSeederRefrigeratorWaltonWut3x8GereCx", struct_name)
    content = content.replace("NewSpecificationSeederRefrigeratorWaltonWut3x8GereCx", func_name)
    content = content.replace("walton-wut-3x8-gere-cx", slug)
    content = content.replace("WUT-3X8-GERE-CX", model)
    content = content.replace("Walton", "Panasonic")
    content = content.replace("ওয়ালটন", "প্যানাসনিক")
    
    # Write file
    with open(file_name, 'w', encoding='utf-8') as f:
        f.write(content)
    
    print(f"Generated {file_name}")

print(f"Generated {len(models)} Panasonic seeder files.")
