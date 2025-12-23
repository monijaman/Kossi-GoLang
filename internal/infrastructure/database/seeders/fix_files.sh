#!/bin/bash
for file in specification_seeder_mobile_*.go; do
    if grep -q "getBanglaTranslations" "$file"; then
        # Extract the struct name from the filename
        struct_name=$(basename "$file" .go | sed 's/specification_seeder_mobile_//' | sed 's/_/-/g' | awk -F'-' '{
            result = ""
            for(i=1; i<=NF; i++) {
                word = $i
                if(word == "se") word = "SE"
                else if(word == "pro") word = "Pro"
                else if(word == "max") word = "Max"
                else if(word == "plus") word = "Plus"
                else if(word == "ultra") word = "Ultra"
                else if(word == "lite") word = "Lite"
                else if(word == "fe") word = "FE"
                else word = toupper(substr(word,1,1)) substr(word,2)
                result = result word
            }
            print result
        }')
        
        # Replace the malformed function with the correct one
        sed -i '/\/\/ getBanglaTranslations returns a map of English specification values to their Bangla translations/,/}/ {
            /\/\/ getBanglaTranslations returns a map of English specification values to their Bangla translations/ {
                h
                d
            }
            /}/ {
                g
                a\
func (s *SpecificationSeederMobile'"$struct_name"') getBanglaTranslations() map[string]string {\
return s.GetCommonBanglaTranslations()\
}
                d
            }
            d
        }' "$file"
    fi
done
