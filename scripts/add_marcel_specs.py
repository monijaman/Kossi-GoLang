import re
import io
from pathlib import Path

BASE = Path('internal/infrastructure/database/seeders')
pattern = 'specification_seeder_refrigerator_marcel-'

# canonical blocks
bangla_block = '''func (s *{struct}) getBanglaTranslations() map[string]string {{
    return map[string]string{{
        "Marcel":                 "মার্সেল",
        "{MODEL_UPPER}":        "{MODEL_UPPER}",
        "{MODEL_UPPER_SLUG}": "{MODEL_UPPER_SLUG}",

        "Direct Cool":         "ডিরেক্ট কুল",
        "Gross Volume":        "১৭৭ লিটার",
        "Net Volume":          "১৭৫ লিটার",
        "Weight":              "৫০ ± ২ কেজি",
        "R600a":               "R600a",
        "Mechanical":          "ম্যানুয়াল/মেকানিক্যাল",
        "Manual":              "ম্যানুয়াল",
        "220 ~ 240":           "২২0 ~ ২৪0",
        "555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
        "580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
    }}
}}
'''

specs_block = '''
    specs := map[string]string{{
        "Brand":               "Marcel",
        "Model Name":          "{MODEL_UPPER_SLUG}",
        "Cooling Technology":  "Direct Cool",
        "Gross Volume":        "177 Ltr.",
        "Net Volume":          "175 Ltr.",
        "Weight":              "50 ± 2 Kg",
        "Refrigerant":         "R600a",
        "Temperature Control": "Mechanical",
        "Voltage":             "220 ~ 240",
        "Dimensions":          "555 x 630 x 1410 mm",
        "Packing Dimensions":  "580 x 645 x 1455 mm",
    }}
'''

# helpers to replace a function block by name

def replace_func_block(text, func_name, new_block):
    # Find start of func by regex
    m = re.search(rf'func\s+\([^)]+\)\s+{re.escape(func_name)}\s*\(', text)
    if not m:
        # try without receiver parentheses (unlikely)
        m = re.search(rf'func\s+{re.escape(func_name)}\s*\(', text)
    if not m:
        return None
    start = m.start()
    # find first '{' after this
    brace_pos = text.find('{', m.end())
    if brace_pos == -1:
        return None
    # walk to matching brace
    i = brace_pos
    depth = 0
    while i < len(text):
        if text[i] == '{':
            depth += 1
        elif text[i] == '}':
            depth -= 1
            if depth == 0:
                end = i
                break
        i += 1
    else:
        return None
    new_text = text[:start] + new_block + text[end+1:]
    return new_text


files = list(BASE.glob(pattern + '*.go'))
print(f'Found {len(files)} files to process')

for fp in files:
    s = fp.read_text(encoding='utf-8')
    # infer struct name from file: last part after 'seeder_' without .go and non-alnum
    # simpler: find type SpecificationSeeder... struct { BaseSeeder }
    m = re.search(r'type\s+(SpecificationSeeder[A-Za-z0-9_]+)\s+struct', s)
    struct = m.group(1) if m else 'SpecificationSeederRefrigeratorMarcel'
    # infer model upper from lines that mention slug or Model Name in getBanglaTranslations or elsewhere
    m2 = re.search(r'"([A-Za-z0-9\-]+)":\s*"([^"]+)"', s)
    model_upper = None
    if 'Model Name' in s:
        m3 = re.search(r'Model Name"\s*:\s*"([^"]+)"', s)
        if m3:
            model_upper = m3.group(1)
    if not model_upper:
        # try from filename
        model_upper = fp.name.replace('specification_seeder_refrigerator_marcel-','').replace('.go','').upper()
        model_upper = model_upper.replace('-', '-').upper()
    # produce replacements
    new_bangla = bangla_block.format(struct=struct, MODEL_UPPER=model_upper, MODEL_UPPER_SLUG=model_upper)
    # replace existing getBanglaTranslations function if present
    if 'getBanglaTranslations' in s:
        s2 = replace_func_block(s, 'getBanglaTranslations()', new_bangla)
        if s2:
            s = s2
    else:
        # insert before Seed func
        s = s.replace('\n// Seed inserts', '\n'+new_bangla+'\n// Seed inserts')
    # replace specs block
    if 'specs := map[string]string' in s:
        # find start of specs := map and its closing brace
        m = re.search(r'specs\s*:=\s*map\[string\]string\s*{', s)
        if m:
            i = m.start()
            brace_pos = s.find('{', m.end()-1)
            # find matching
            j = brace_pos
            depth = 0
            while j < len(s):
                if s[j] == '{':
                    depth += 1
                elif s[j] == '}':
                    depth -= 1
                    if depth == 0:
                        end = j
                        break
                j += 1
            new_specs = specs_block.format(MODEL_UPPER_SLUG=model_upper)
            s = s[:i] + new_specs + s[end+1:]
    else:
        # insert before for loop that iterates specs
        s = s.replace('\tbanglaTranslations := s.getBanglaTranslations()', '\t'+specs_block.format(MODEL_UPPER_SLUG=model_upper)+'\n\tbanglaTranslations := s.getBanglaTranslations()')
    fp.write_text(s, encoding='utf-8')

print('Done')
