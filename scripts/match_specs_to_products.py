import csv
from pathlib import Path
import re

PROD = Path('data/marcel_products.csv')
SPECS = Path('data/marcel_specs.csv')
OUT = Path('data/marcel_specs_with_url.csv')

def normalize(s):
    if s is None:
        return ''
    s = s.strip()
    s = re.sub(r"\s+\(.*\)$", "", s)  # remove trailing parenthesis like (Inverter)
    s = s.replace('\u00a0',' ')
    s = s.strip()
    return s

# load products map: try multiple keys
prod_map = {}
alt_map = {}
with PROD.open(encoding='utf-8', newline='') as f:
    r = csv.reader(f)
    hdr = next(r, None)
    for row in r:
        if len(row) < 3:
            continue
        category, model, url = row[0].strip(), row[1].strip(), row[2].strip()
        if not model:
            continue
        n = normalize(model).upper()
        prod_map[n] = url
        # also store variants: remove spaces, parentheses, compress dashes
        alt = re.sub(r'[^A-Z0-9\-]', '', n)
        alt_map[alt] = url

# process specs
with SPECS.open(encoding='utf-8', newline='') as inf, OUT.open('w', encoding='utf-8', newline='') as outf:
    r = csv.reader(inf)
    w = csv.writer(outf)
    hdr = next(r, None)
    # ensure header exists
    if hdr is None:
        hdr = ['category','model','url','spec_key','spec_value','found']
    w.writerow(hdr)

    for row in r:
        if len(row) < 6:
            # pad
            row = (row + ['']*6)[:6]
        category, model, url, spec_key, spec_value, found = row
        model_n = normalize(model).upper()
        if (not model_n) and spec_value:
            # try to extract model from spec_value if it looks like a model token
            pass
        filled_url = url.strip()
        # lookup
        if not filled_url and model_n:
            if model_n in prod_map:
                filled_url = prod_map[model_n]
            else:
                alt = re.sub(r'[^A-Z0-9\-]', '', model_n)
                filled_url = alt_map.get(alt, '')
        # write row with filled url and normalized model
        w.writerow([category, model.strip(), filled_url, spec_key, spec_value, found])

print('Wrote', OUT)
