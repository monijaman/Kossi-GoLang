import csv
import re
from pathlib import Path

BASE = Path('internal/infrastructure/database/seeders')
PATTERN = 'specification_seeder_refrigerator_marcel-'
MANUAL = Path('data/manual_spec_380L.csv')

# Load manual specs as list of (spec_key, spec_value)
manual = []
with MANUAL.open(encoding='utf-8') as f:
    r = csv.reader(f)
    next(r, None)
    for row in r:
        if len(row) < 3:
            continue
        spec_key = row[1].strip()
        spec_value = row[2].strip()
        if spec_key:
            manual.append((spec_key, spec_value))

if not manual:
    print('No manual specs found in', MANUAL)
    raise SystemExit(1)

files = list(BASE.glob(PATTERN + '*.go'))
print('Found', len(files), 'files')

updated = 0

# helper functions

def detect_struct(text):
    m = re.search(r'type\s+(SpecificationSeeder[A-Za-z0-9_]+)\s+struct', text)
    return m.group(1) if m else 'SpecificationSeederRefrigeratorMarcel'

def slug_to_struct(filename):
    # convert filename to struct name by reading the existing type line
    txt = Path(BASE/filename).read_text(encoding='utf-8')
    return detect_struct(txt)

# helper to replace a function block by name
def replace_func_block(text, func_name, new_block):
    # find func with receiver or without
    m = re.search(rf'func\s+\([^)]+\)\s+{re.escape(func_name)}\s*\(', text)
    if not m:
        m = re.search(rf'func\s+{re.escape(func_name)}\s*\(', text)
    if not m:
        return None
    start = m.start()
    brace_pos = text.find('{', m.end())
    if brace_pos == -1:
        return None
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
    return text[:start] + new_block + text[end+1:]

for fp in files:
    s = fp.read_text(encoding='utf-8')
    # find product slug to include in Bangla translations keys if present
    mslug = re.search(r'productSlug := "([^"]+)"', s)
    slug = mslug.group(1) if mslug else ''
    # detect model name to include as Model Name
    modelname = ''
    mmodel = re.search(r'"Model Name"\s*:\s*"([^"]+)"', s)
    if mmodel:
        modelname = mmodel.group(1)
    else:
        # try to infer from filename
        modelname = fp.name.replace('specification_seeder_refrigerator_marcel-','').replace('.go','').upper()

    # build specs block text
    specs_lines = ['\tspecs := map[string]string{']
    for k,v in manual:
        # for keys that include category prefix like 'Capacity - Gross Volume', normalize to existing key names if possible
        key = k
        val = v
        specs_lines.append(f'\t\t"{key}": "{val}",')
    specs_lines.append('\t}')
    specs_block = '\n'.join(specs_lines) + '\n'

    # build bangla translations function
    bangla_lines = []
    # determine struct name
    struct_name = detect_struct(s)
    bangla_lines.append(f'func (s *{struct_name}) getBanglaTranslations() map[string]string {{')
    bangla_lines.append('\treturn map[string]string{')
    # basic entries
    bangla_lines.append('\t\t"Marcel": "মার্সেল",')
    if modelname:
        bangla_lines.append(f'\t\t"{modelname}": "{modelname}",')
    # add mapping from manual values to themselves (no translation available)
    for k,v in manual:
        bangla_lines.append(f'\t\t"{v}": "{v}",')
    bangla_lines.append('\t}')
    bangla_lines.append('}')
    bangla_block = '\n'.join(bangla_lines) + '\n'

    # Replace specs map
    if 'specs := map[string]string' in s:
        m = re.search(r'specs\s*:=\s*map\[string\]string\s*{', s)
        if m:
            i = m.start()
            brace_pos = s.find('{', m.end()-1)
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
            s = s[:i] + specs_block + s[end+1:]
    else:
        # fallback: insert before banglaTranslations assignment
        s = s.replace('\tbanglaTranslations := s.getBanglaTranslations()', '\t'+specs_block+'\tbanglaTranslations := s.getBanglaTranslations()')

    # Replace getBanglaTranslations function
    # find actual function name occurrence without parentheses for matching
    if 'getBanglaTranslations' in s:
        # prepare function name pattern including receiver
        func_name = 'getBanglaTranslations()'
        s2 = replace_func_block(s, 'getBanglaTranslations()', bangla_block)
        if s2:
            s = s2
    else:
        # insert bangla_block above Seed func
        s = s.replace('\nfunc (s *', '\n'+bangla_block+'\nfunc (s *')

    fp.write_text(s, encoding='utf-8')
    updated += 1

print('Updated', updated, 'files')
