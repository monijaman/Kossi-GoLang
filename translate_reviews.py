#!/usr/bin/env python3
"""
Translate TV product reviews from English to Bengali
Extracts English reviews and generates comprehensive Bengali translations
"""

import re
import time
import requests
import json

def extract_reviews_from_file(filepath):
    """Extract all English review content from the Go seeder file"""
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Find all product names and their reviews
    product_pattern = r'productName:\s+"([^"]+)"[^`]+review:\s+`([^`]+)`'
    matches = re.findall(product_pattern, content, re.DOTALL)
    
    print(f"Found {len(matches)} English reviews")
    return matches

def translate_text_chunk(translator, text, max_length=4500):
    """Translate text in chunks to handle size limits"""
    if len(text) <= max_length:
        try:
            result = translator.translate(text, src='en', dest='bn')
            return result.text
        except Exception as e:
            print(f"    Error: {e}")
            return text
    
    # Split into smaller chunks
    sentences = re.split(r'([.!?]\s+)', text)
    translated = ''
    current_chunk = ''
    
    for sentence in sentences:
        if len(current_chunk + sentence) < max_length:
            current_chunk += sentence
        else:
            if current_chunk:
                try:
                    result = translator.translate(current_chunk, src='en', dest='bn')
                    translated += result.text
                    time.sleep(0.5)  # Rate limiting
                except Exception as e:
                    print(f"    Chunk error: {e}")
                    translated += current_chunk
            current_chunk = sentence
    
    if current_chunk:
        try:
            result = translator.translate(current_chunk, src='en', dest='bn')
            translated += result.text
        except Exception as e:
            translated += current_chunk
    
    return translated

def translate_html_review(translator, html_content):
    """Translate HTML review content preserving HTML tags"""
    # Split by HTML tags to preserve them
    parts = re.split(r'(<[^>]+>)', html_content)
    
    translated_parts = []
    for part in parts:
        if part.startswith('<') and part.endswith('>'):
            # Keep HTML tags as-is
            translated_parts.append(part)
        elif part.strip():
            # Translate text content
            translated_parts.append(translate_text_chunk(translator, part))
        else:
            translated_parts.append(part)
    
    return ''.join(translated_parts)

def main():
    filepath = 'd:\\GO\\gocrit\\gocrit_server\\internal\\infrastructure\\database\\seeders\\tv_product_review_seeder.go'
    
    print("Initializing translator...")
    translator = Translator()
    
    print("Extracting reviews...")
    products = extract_reviews_from_file(filepath)
    
    if not products:
        print("No reviews found!")
        return
    
    print(f"\nTranslating {len(products)} reviews to Bengali...")
    print("This will take 10-15 minutes due to rate limiting...\n")
    
    translations = []
    for i, (product_name, review) in enumerate(products, 1):
        print(f"[{i}/{len(products)}] Translating: {product_name}")
        try:
            translated = translate_html_review(translator, review)
            translations.append((product_name, translated))
            print(f"  ✓ Completed ({len(translated)} chars)")
            time.sleep(1)  # Rate limiting between products
        except Exception as e:
            print(f"  ✗ Error: {e}")
            translations.append((product_name, review))
    
    # Save translations to file
    output_file = 'd:\\GO\\gocrit\\gocrit_server\\bengali_translations.txt'
    with open(output_file, 'w', encoding='utf-8') as f:
        for i, (product_name, translation) in enumerate(translations, 1):
            f.write(f"=== PRODUCT {i}: {product_name} ===\n")
            f.write(translation)
            f.write("\n\n" + "="*80 + "\n\n")
    
    print(f"\n✓ All translations saved to: {output_file}")
    print("\nNext steps:")
    print("1. Review the translations in bengali_translations.txt")
    print("2. Use the translations to update reviewBn fields in the seeder")
    print("3. Rebuild and run the seeder")

if __name__ == '__main__':
    main()
