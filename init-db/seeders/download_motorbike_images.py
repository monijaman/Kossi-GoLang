#!/usr/bin/env python3
"""
download_motorbike_images.py

Usage:
  python download_motorbike_images.py --num 3

What it does:
  - Reads `motorbikes.json` in the same folder for model names
  - For each model, searches Wikimedia Commons for images matching the model
    plus common color names and downloads up to N images (default 3) with
    different color query terms (red, blue, black, white, green, yellow).
  - Saves files into: crit_client/public/images/motorbikes/<model-slug>/
  - Produces `attributions.csv` with source URLs and license metadata.

Notes:
  - Run locally. This script does not commit files to git.
  - Wikimedia Commons is used to favor freely-licensed images.
  - If you prefer another source or want the script to auto-commit, tell me.
"""

import argparse
import csv
import json
import os
import re
import sys
import time
from pathlib import Path
from urllib.parse import quote_plus

import requests


COMMON_COLORS = ["red", "blue", "black", "white", "green", "yellow", "silver", "grey", "orange"]


def slugify(s: str) -> str:
    s = s.lower()
    s = re.sub(r"[^a-z0-9]+", "-", s)
    s = s.strip("-")
    return s or "model"


def mw_search_files(query: str, limit: int = 10):
    """Search Wikimedia Commons file namespace for a query. Returns list of titles."""
    url = "https://commons.wikimedia.org/w/api.php"
    params = {
        "action": "query",
        "format": "json",
        "list": "search",
        "srsearch": query,
        "srnamespace": 6,  # File namespace
        "srwhat": "text",
        "srlimit": limit,
    }
    headers = {
        "User-Agent": "MotorbikeImageDownloader/1.0 (https://example.com; contact@example.com)"
    }
    r = requests.get(url, params=params, headers=headers, timeout=30)
    r.raise_for_status()
    data = r.json()
    return [item["title"] for item in data.get("query", {}).get("search", [])]


def mw_get_imageinfo(titles):
    """Get imageinfo for one or more file titles. Returns dict title -> imageinfo dict."""
    url = "https://commons.wikimedia.org/w/api.php"
    joined = "|".join(titles)
    params = {
        "action": "query",
        "format": "json",
        "prop": "imageinfo",
        "iiprop": "url|extmetadata",
        "titles": joined,
    }
    headers = {
        "User-Agent": "MotorbikeImageDownloader/1.0 (https://example.com; contact@example.com)"
    }
    r = requests.get(url, params=params, headers=headers, timeout=30)
    r.raise_for_status()
    data = r.json()
    pages = data.get("query", {}).get("pages", {})
    out = {}
    for pageid, page in pages.items():
        title = page.get("title")
        ii = page.get("imageinfo")
        if ii and isinstance(ii, list):
            out[title] = ii[0]
    return out


def download_url(url, dest_path: Path):
    dest_path.parent.mkdir(parents=True, exist_ok=True)
    with requests.get(url, stream=True, timeout=60) as r:
        r.raise_for_status()
        with open(dest_path, "wb") as f:
            for chunk in r.iter_content(chunk_size=8192):
                if chunk:
                    f.write(chunk)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--num", "-n", type=int, default=3, help="Images per model")
    parser.add_argument("--out", "-o", default="crit_client/public/images/motorbikes/", help="Output base dir")
    parser.add_argument("--models", "-m", default="motorbikes.json", help="Path to motorbikes.json (relative to script)")
    parser.add_argument("--start", type=int, default=0, help="Start index (for batching)")
    parser.add_argument("--end", type=int, default=-1, help="End index (non-inclusive) for batching")
    args = parser.parse_args()

    here = Path(__file__).resolve().parent
    models_path = (here / args.models).resolve()
    if not models_path.exists():
        print("motorbikes.json not found at:", models_path)
        sys.exit(1)

    with open(models_path, "r", encoding="utf-8") as f:
        bikes = json.load(f)

    out_base = Path(args.out)
    out_base.mkdir(parents=True, exist_ok=True)

    attributions = []

    start = args.start
    end = args.end if args.end >= 0 else len(bikes)

    for idx, bike in enumerate(bikes[start:end], start=start):
        name = bike.get("model") or bike.get("name_en") or bike.get("brand", "") + " " + bike.get("model", "")
        print(f"[{idx}] Processing model: {name}")
        model_slug = slugify(name)
        model_dir = out_base / model_slug
        model_dir.mkdir(parents=True, exist_ok=True)

        found = 0
        used_titles = set()
        # Try color variations, but fallback to generic searches
        search_terms = [f"{name} {color}" for color in COMMON_COLORS] + [name]

        for term in search_terms:
            if found >= args.num:
                break
            try:
                titles = mw_search_files(term, limit=8)
            except Exception as e:
                print("  Search failed for", term, "->", e)
                time.sleep(1)
                continue
            if not titles:
                time.sleep(0.3)
                continue

            # Filter out titles we've already used
            titles = [t for t in titles if t not in used_titles]
            if not titles:
                time.sleep(0.2)
                continue

            # Get imageinfo for these titles
            try:
                info_map = mw_get_imageinfo(titles)
            except Exception as e:
                print("  imageinfo failed ->", e)
                time.sleep(0.5)
                continue

            for title, info in info_map.items():
                if found >= args.num:
                    break
                used_titles.add(title)
                img_url = info.get("url")
                if not img_url:
                    continue

                # Determine extension
                ext = os.path.splitext(img_url)[1].split("?")[0]
                if not ext:
                    ext = ".jpg"
                filename = f"{model_slug}_{found+1}{ext}"
                dest = model_dir / filename
                try:
                    print(f"  Downloading ({found+1}) {img_url}")
                    download_url(img_url, dest)
                except Exception as e:
                    print("    download failed ->", e)
                    continue

                # Extract license/attribution info if present
                em = info.get("extmetadata", {})
                license = em.get("LicenseShortName", {}).get("value") if em else None
                artist = em.get("Credit", {}).get("value") or em.get("Artist", {}).get("value") if em else None

                attributions.append({
                    "model": name,
                    "model_slug": model_slug,
                    "filename": str(dest.relative_to(Path.cwd())),
                    "source_title": title,
                    "source_url": img_url,
                    "artist": artist,
                    "license": license,
                })

                found += 1
                # polite pause
                time.sleep(0.4)

        if found == 0:
            print(f"  No images found for {name}")

    # Write attributions CSV
    csv_path = out_base / "attributions.csv"
    with open(csv_path, "w", newline='', encoding="utf-8") as csvfile:
        fieldnames = ["model", "model_slug", "filename", "source_title", "source_url", "artist", "license"]
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()
        for row in attributions:
            writer.writerow(row)

    print("Done. Images saved under:", out_base)
    print("Attributions written to:", csv_path)


if __name__ == "__main__":
    main()
