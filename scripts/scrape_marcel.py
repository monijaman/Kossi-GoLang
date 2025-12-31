import requests
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import csv

BASE = "https://marcelbd.com"
CATEGORIES = {
    "direct-cool-refrigerator": list(range(1, 10)),  # 1..9
    "no-frost-refrigerator": [1],
    "beverage-cooler": [1],
    "freezer": [1],
}

out_path = "data/marcel_products.csv"
seen = set()
rows = []

headers = {"User-Agent": "Mozilla/5.0 (compatible; scraper/1.0)"}

for cat, pages in CATEGORIES.items():
    for p in pages:
        if cat == "direct-cool-refrigerator":
            url = f"{BASE}/refrigerator-freezer/{cat}?page={p}"
        else:
            url = f"{BASE}/refrigerator-freezer/{cat}"
        try:
            r = requests.get(url, headers=headers, timeout=15)
            r.raise_for_status()
        except Exception as e:
            print(f"Failed to fetch {url}: {e}")
            continue
        soup = BeautifulSoup(r.text, "html.parser")
        # Find anchors that look like product links under this category
        for a in soup.find_all("a", href=True):
            href = a["href"].strip()
            if "/refrigerator-freezer/" in href and href.count("/") >= 2:
                model = a.get_text(strip=True)
                if not model:
                    # sometimes anchor wraps image; try next sibling text
                    continue
                full = urljoin(BASE, href)
                key = (model, full)
                if key in seen:
                    continue
                seen.add(key)
                rows.append((cat, model, full))

# Write CSV
with open(out_path, "w", newline='', encoding='utf-8') as f:
    w = csv.writer(f)
    w.writerow(["category", "model", "url"])
    for r in rows:
        w.writerow(r)

print(f"Wrote {len(rows)} rows to {out_path}")
