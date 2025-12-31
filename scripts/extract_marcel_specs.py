import csv
import requests
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import time

IN = 'data/marcel_products.csv'
OUT = 'data/marcel_specs.csv'

headers = {'User-Agent':'Mozilla/5.0 (compatible; extractor/1.0)'}

def extract_specs_from_soup(soup):
    # Heuristics to find specification content
    # 1. common OpenCart: div id 'tab-specification' or '#tab-specification'
    selectors = ["#tab-specification", "#tab-specifications", ".tab-specification", ".specification", ".product-specification", ".specifications", ".product-info .specification"]
    for sel in selectors:
        node = soup.select_one(sel)
        if node:
            specs = parse_spec_container(node)
            if specs:
                return specs
    # 2. find heading with 'Technical Specification' or 'Specifications' and take next element
    headings = soup.find_all(lambda t: t.name in ['h1','h2','h3','h4','strong','b'] and t.get_text(strip=True).lower().find('technical specification')!=-1)
    if not headings:
        headings = soup.find_all(lambda t: t.name in ['h1','h2','h3','h4','strong','b'] and 'specification' in t.get_text(strip=True).lower())
    for h in headings:
        nxt = h.find_next()
        if nxt:
            specs = parse_spec_container(nxt)
            if specs:
                return specs
    # 3. fallback: search for any table with two columns near keywords
    for table in soup.find_all('table'):
        specs = parse_table(table)
        if specs:
            return specs
    return None


def parse_spec_container(node):
    # if node contains a table
    table = node.find('table')
    if table:
        return parse_table(table)
    # dl dt/dd
    dl = node.find('dl')
    if dl:
        items = []
        for dt, dd in zip(dl.find_all('dt'), dl.find_all('dd')):
            k = dt.get_text(strip=True)
            v = dd.get_text(strip=True)
            items.append((k, v))
        if items:
            return items
    # lists
    lis = node.find_all('li')
    items = []
    for li in lis:
        text = li.get_text(separator=' ', strip=True)
        if ':' in text:
            k, v = [p.strip() for p in text.split(':',1)]
            items.append((k, v))
    if items:
        return items
    # paragraphs with colon
    ps = node.find_all('p')
    for p in ps:
        t = p.get_text(strip=True)
        if ':' in t:
            parts = [pt.strip() for pt in t.split(':',1)]
            if len(parts)==2:
                return [(parts[0], parts[1])]
    return None


def parse_table(table):
    rows = []
    # try rows with th/td pairs
    for tr in table.find_all('tr'):
        cols = tr.find_all(['th','td'])
        if len(cols) >= 2:
            k = cols[0].get_text(strip=True)
            v = cols[1].get_text(' ', strip=True)
            rows.append((k, v))
        else:
            # sometimes key: value in single td
            text = tr.get_text(' ', strip=True)
            if ':' in text:
                k, v = [p.strip() for p in text.split(':',1)]
                rows.append((k, v))
    return rows if rows else None


with open(IN, newline='', encoding='utf-8') as inf, open(OUT, 'w', newline='', encoding='utf-8') as outf:
    r = csv.reader(inf)
    w = csv.writer(outf)
    w.writerow(['category','model','url','spec_key','spec_value','found'])
    next(r, None)
    for i, row in enumerate(r, start=1):
        if len(row) < 3:
            continue
        category, model, url = row[0].strip(), row[1].strip(), row[2].strip()
        # skip category pages accidentally listed as model names
        if url.endswith('/refrigerator-freezer/direct-cool-refrigerator') or 'page=' in url and 'refrigerator-freezer/direct-cool-refrigerator' in url:
            continue
        try:
            res = requests.get(url, headers=headers, timeout=15)
            res.raise_for_status()
            soup = BeautifulSoup(res.text, 'html.parser')
            specs = extract_specs_from_soup(soup)
            if specs:
                for k, v in specs:
                    w.writerow([category, model, url, k, v, 'yes'])
            else:
                w.writerow([category, model, url, '', '', 'no'])
        except Exception as e:
            w.writerow([category, model, url, '', f'error: {e}', 'error'])
        # be polite
        time.sleep(0.4)

print('Done; results written to', OUT)
