import os
import re

dir_path = '.'
for filename in os.listdir(dir_path):
    if filename.startswith('specification_seeder_refrigerator_vision-') and filename.endswith('.go'):
        filepath = os.path.join(dir_path, filename)
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        # Find the specs map and empty it
        start = content.find('specs := map[string]string{')
        if start != -1:
            end = content.find('\n\t}', start)
            if end != -1:
                # Replace the content between { and }
                new_content = content[:start+len('specs := map[string]string{')] + '\n\t}' + content[end+3:]
                with open(filepath, 'w', encoding='utf-8') as f:
                    f.write(new_content)
                print(f'Updated {filename}')
