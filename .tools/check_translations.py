import re
import sys
import os

files = []
for f in os.listdir(r"I:/GO/kossti/gocrit_server/internal/infrastructure/database/seeders"):
    if f.startswith("specification_seeder_mobile_vivo-") and f.endswith(".go"):
        files.append(os.path.join(r"I:/GO/kossti/gocrit_server/internal/infrastructure/database/seeders", f))

pair_re = re.compile(r'"([^"\\]*(?:\\.[^"\\]*)*)"\s*:\s*"([^"\\]*(?:\\.[^"\\]*)*)"')
map_start_re = re.compile(r'(?:specs|translations)\s*:?"]?\s*:=\s*map\[string\]string\s*\{')

results = {}

for path in files:
    try:
        with open(path, 'r', encoding='utf-8') as f:
            txt = f.read()
    except Exception as e:
        print(f"ERROR reading {path}: {e}")
        continue
    # find specs map block
    specs = {}
    translations = {}

    # crude: find 'specs :=' and capture the brace content balancing braces
    def extract_block(name):
        m = re.search(rf'{name}\s*:=\s*map\[string\]string\s*\{{', txt)
        if not m:
            return ''
        start = m.end()
        depth = 1
        i = start
        while i < len(txt):
            c = txt[i]
            if c == '{':
                depth += 1
            elif c == '}':
                depth -= 1
                if depth == 0:
                    return txt[start:i]
            i += 1
        return ''

    specs_block = extract_block('specs')
    translations_block = extract_block('translations')

    def parse_map(block):
        d = {}
        for k, v in pair_re.findall(block):
            key = k.encode('utf-8').decode('unicode_escape')
            val = v.encode('utf-8').decode('unicode_escape')
            d[key] = val
        return d

    specs = parse_map(specs_block)
    translations = parse_map(translations_block)

    # Now compare: for each spec value check translation
    report = []
    total = 0
    ok = 0
    equal = 0
    longer = 0
    missing = 0
    for k, v in specs.items():
        total += 1
        en = v
        bn = translations.get(en)
        if bn is None:
            missing += 1
            report.append((en, 'MISSING', len(en), None))
        else:
            ln_en = len(en)
            ln_bn = len(bn)
            if ln_bn < ln_en:
                ok += 1
                report.append((en, 'OK', ln_en, ln_bn))
            elif ln_bn == ln_en:
                equal += 1
                report.append((en, 'EQUAL', ln_en, ln_bn))
            else:
                longer += 1
                report.append((en, 'LONGER', ln_en, ln_bn))
    results[path] = dict(total=total, ok=ok, equal=equal, longer=longer, missing=missing, details=report)

# Print concise report
for path, r in results.items():
    print('\nFILE:', os.path.basename(path))
    print(f"Total specs: {r['total']}, OK: {r['ok']}, equal: {r['equal']}, longer: {r['longer']}, missing: {r['missing']}")
    if r['longer'] or r['missing'] or r['equal']:
        print('Details (value, status, len_en, len_bn):')
        for item in r['details']:
            status = item[1]
            if status != 'OK':
                print(item)

# exit code non-zero if any missing/longer
any_issue = any(r['longer']>0 or r['missing']>0 for r in results.values())
sys.exit(1 if any_issue else 0)
