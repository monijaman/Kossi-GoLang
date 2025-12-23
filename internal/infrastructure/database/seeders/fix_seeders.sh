#!/bin/bash
for file in specification_seeder_refrigerator_*.go; do
    if grep -q "Ensure the ID is set after creation" "$file"; then
        # Remove the wrong block inside the if
        sed -i '/if err := db\.Create(sModel)\.Error; err != nil {/,/}/ {
            /\/\/ Ensure the ID is set after creation/,/}/ d
        }' "$file"
        # Add the correct block after the if
        sed -i '/if err := db\.Create(sModel)\.Error; err != nil {/a\
\t\t\t\treturn err\n\t\t\t\t}\n\t\t\t\t\/\/ Ensure the ID is set after creation\n\t\t\t\tif sModel.ID == 0 {\n\t\t\t\t\tif err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {\n\t\t\t\t\t\treturn err\n\t\t\t\t\t}\n\t\t\t\t}\n' "$file"
    fi
done
