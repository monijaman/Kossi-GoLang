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

print("-- Panasonic")
print("INSERT INTO \"products\" (\"name\", \"slug\", \"category_id\", \"brand_id\", \"model\", \"status\", \"priority\", \"created_at\", \"updated_at\") VALUES")

for i, model in enumerate(models):
    slug = "panasonic-" + model.lower()
    name = f"Panasonic {model}"
    comma = "," if i < len(models) - 1 else ""
    print(f"('{name}', '{slug}', 100, (SELECT id FROM brands WHERE name = 'Panasonic'), '{model}', 1, 1, NOW(), NOW()){comma}")

print("ON CONFLICT (slug) DO NOTHING;")
