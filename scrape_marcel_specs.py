import csv
import urllib.request
import urllib.error
from html.parser import HTMLParser
import time
import re

class SpecTableParser(HTMLParser):
    """Parse spec tables from HTML"""
    def __init__(self):
        super().__init__()
        self.specs = []
        self.in_table = False
        self.in_row = False
        self.in_cell = False
        self.cell_data = []
        self.row_data = []
        self.current_section = ""
        self.skip_cell = False
        
    def handle_starttag(self, tag, attrs):
        attrs_dict = dict(attrs)
        
        if tag == 'table':
            self.in_table = True
        elif tag == 'tr' and self.in_table:
            self.in_row = True
            self.row_data = []
        elif tag in ['td', 'th'] and self.in_row:
            self.in_cell = True
            self.cell_data = []
            # Check for colspan="4" which indicates section header
            if attrs_dict.get('colspan') == '4':
                self.skip_cell = True
            else:
                self.skip_cell = False
                
    def handle_endtag(self, tag):
        if tag in ['td', 'th'] and self.in_cell:
            self.in_cell = False
            text = ''.join(self.cell_data).strip()
            self.row_data.append(text)
            
        elif tag == 'tr' and self.in_row:
            self.in_row = False
            # Process row data
            if self.row_data:
                if self.skip_cell and len(self.row_data) >= 1:
                    # This is a section header
                    section_text = self.row_data[0].rstrip(':').strip()
                    self.current_section = section_text
                elif len(self.row_data) >= 2 and self.row_data[0] and self.row_data[1]:
                    # This is a spec row
                    self.specs.append((self.row_data[0], self.row_data[1]))
                    
        elif tag == 'table':
            self.in_table = False
            
    def handle_data(self, data):
        if self.in_cell:
            self.cell_data.append(data)

def scrape_product_specs(url, category, model):
    """Scrape technical specifications from a Marcel product page"""
    try:
        print(f"Scraping: {url}")
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
        }
        req = urllib.request.Request(url, headers=headers)
        with urllib.request.urlopen(req, timeout=10) as response:
            html_content = response.read().decode('utf-8', errors='ignore')
        
        specs = []
        
        # Find tech_spec div to get just the table content
        start_idx = html_content.find('id="tech_spec"')
        if start_idx != -1:
            # Get content from tech_spec div
            tech_spec_start = html_content.rfind('<div', 0, start_idx)
            tech_spec_end = html_content.find('</div>', start_idx) + 6
            table_html = html_content[tech_spec_start:tech_spec_end]
        else:
            table_html = html_content
        
        # Parse the HTML
        parser = SpecTableParser()
        try:
            parser.feed(table_html)
        except:
            pass
        
        # Convert parsed specs to the required format
        if parser.specs:
            for key, value in parser.specs:
                # Clean up the value by removing excess whitespace
                cleaned_value = ' '.join(value.split())
                specs.append({
                    'category': category,
                    'model': model,
                    'url': url,
                    'spec_key': key,
                    'spec_value': cleaned_value,
                    'found': 'yes'
                })
        
        # If no specs found via parser, try regex-based extraction with looser patterns
        if not specs:
            # Try to find any <td> tags that look like key-value pairs
            # Pattern 1: <td>key</td><td>value</td>
            pattern1 = r'<td[^>]*>\s*([^<]+?)\s*</td>\s*<td[^>]*>\s*([^<]+?)\s*</td>'
            matches = re.findall(pattern1, table_html, re.DOTALL | re.IGNORECASE)
            
            for key, value in matches:
                # Clean up the values
                key_clean = ' '.join(key.split()).strip()
                value_clean = ' '.join(value.split()).strip()
                
                # Filter out very short keys (likely not real specs) and empty values
                if len(key_clean) > 2 and len(value_clean) > 0 and key_clean.lower() not in ['<', '>', '|<', '1']:
                    specs.append({
                        'category': category,
                        'model': model,
                        'url': url,
                        'spec_key': key_clean,
                        'spec_value': value_clean,
                        'found': 'yes'
                    })

        
        if not specs:
            # Mark as not found if no specs were extracted
            specs.append({
                'category': category,
                'model': model,
                'url': url,
                'spec_key': '',
                'spec_value': '',
                'found': 'no'
            })
        
        return specs
        
    except urllib.error.URLError as e:
        print(f"URL Error scraping {url}: {str(e)}")
        return [{
            'category': category,
            'model': model,
            'url': url,
            'spec_key': '',
            'spec_value': '',
            'found': 'error'
        }]
    except Exception as e:
        print(f"Error scraping {url}: {str(e)}")
        return [{
            'category': category,
            'model': model,
            'url': url,
            'spec_key': '',
            'spec_value': '',
            'found': 'error'
        }]


def main():
    input_csv = 'data/marcel_products.csv'
    output_csv = 'data/marcel_specs.csv'
    
    # Read the input CSV
    products = []
    with open(input_csv, 'r', encoding='utf-8') as f:
        reader = csv.DictReader(f)
        for row in reader:
            products.append(row)
    
    print(f"Found {len(products)} products to scrape")
    
    # Open output CSV for writing
    with open(output_csv, 'w', newline='', encoding='utf-8') as f:
        fieldnames = ['category', 'model', 'url', 'spec_key', 'spec_value', 'found']
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        
        # Scrape specs for each product
        for i, product in enumerate(products, 1):
            url = product['url']
            category = product['category']
            model = product['model']
            
            print(f"\n[{i}/{len(products)}] Processing: {model}")
            
            specs = scrape_product_specs(url, category, model)
            
            # Write specs to CSV
            for spec in specs:
                writer.writerow(spec)
            
            # Be polite - wait between requests
            time.sleep(2)
        
        print(f"\n✓ Scraped data saved to {output_csv}")
        print(f"Total products processed: {len(products)}")

if __name__ == '__main__':
    main()

