#!/bin/bash

# Fix all remaining mixed English-Bengali patterns in seeder files

cd init-db/seeders

echo "=== FIXING ALL MIXED LANGUAGE PATTERNS ===" 
echo ""

# Array of replacement patterns (old -> new)
declare -a FIXES=(
    # Display technology - keep technical terms but translate descriptors
    's/"LTPO Super Retina XDR OLED, 120Hz": "এলটিপিও Super Retina XDR ওলেড, 120Hz"/"LTPO Super Retina XDR OLED, 120Hz": "এলটিপিও সুপার রেটিনা এক্সডিআর ওলেড, 120Hz"/g'
    's/"Dynamic AMOLED 2X, 120Hz, HDR10+"/"Dynamic অ্যামোলেড 2X, 120Hz, HDR10+"/g'
    's/"Foldable Dynamic LTPO AMOLED 2X, 120Hz, HDR10+"/"Foldable Dynamic এলটিপিও অ্যামোলেড 2X, 120Hz, HDR10+"/g'
    's/"Super AMOLED, 90Hz, 800 nits": "Super অ্যামোলেড, 90Hz, 800 নিট"/"Super AMOLED, 90Hz, 800 nits": "সুপার অ্যামোলেড, 90Hz, 800 নিট"/g'
    's/"Super AMOLED, 120Hz": "Super অ্যামোলেড, 120Hz"/"Super AMOLED, 120Hz": "সুপার অ্যামোলেড, 120Hz"/g'
    
    # Processor specs - keep chip model, translate architecture
    's/"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "Octa-কোর (2×2.0 গিগাহার্জ Cortex-A75 + 6×1.8 গিগাহহার্জ Cortex-A55)"/"Octa-core (2×2.0 GHz Cortex-A75 + 6×1.8 GHz Cortex-A55)": "অকটা-কোর (2×2.0 গিগাহার্জ Cortex-A75 + 6×1.8 গিগাহার्ज Cortex-A55)"/g'
    's/"Octa-core (2×2.6 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-কোর (2×2.6 গিগাহার্জ Cortex-A76 + 6×2.0 গিগাহহার्ज Cortex-A55)"/"Octa-core (2×2.6 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "অকটা-কোর (2×2.6 গিগাহার्ज Cortex-A76 + 6×2.0 গিগাহার्ज Cortex-A55)"/g'
    's/"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "Octa-কোর (2×2.4 গিগাহার्ज Cortex-A78 + 6×2.0 গিগাহার्ज Cortex-A55)"/"Octa-core (2×2.4 GHz Cortex-A78 + 6×2.0 GHz Cortex-A55)": "অকটা-কোর (2×2.4 গিগাহार्ज Cortex-A78 + 6×2.0 গিগাহার्ज Cortex-A55)"/g'
    's/"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "Octa-কোর (2×2.2 গিগাহার्ज Cortex-A76 + 6×2.0 গিগাহার्ज Cortex-A55)"/"Octa-core (2×2.2 GHz Cortex-A76 + 6×2.0 GHz Cortex-A55)": "অকটা-কোর (2×2.2 গিগাহার्ज Cortex-A76 + 6×2.0 গিগাহार्ज Cortex-A55)"/g'
    
    # Snapdragon with model numbers - keep chip codes
    's/"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"/"Qualcomm SM8650-AB Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AB স্ন্যাপড্রাগন 8 জেন 3 (4 nm)"/g'
    's/"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"/"Qualcomm SM8550-AB Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AB স্ন্যাপড্রাগন 8 জেন 2 (4 nm)"/g'
    's/"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 Elite (3 nm)"/"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750 স্ন্যাপড্রাগন 8 এলিট (3 nm)"/g'
    's/"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 Gen 3 (4 nm)"/"Qualcomm SM8650-AC Snapdragon 8 Gen 3 (4 nm)": "কোয়ালকম SM8650-AC স্ন্যাপড্রাগন 8 জেন 3 (4 nm)"/g'
    's/"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 Gen 2 (4 nm)"/"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম SM8550-AC স্ন্যাপড্রাগন 8 জেন 2 (4 nm)"/g'
    's/"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 Elite (3 nm)"/"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "কোয়ালকম SM8750-AC স্ন্যাপড্রাগন 8 এলিট (3 nm)"/g'
    's/"Qualcomm Snapdragon 695 5G (6 nm)": "Qualcomm স্ন্যাপড্রাগন 695 ৫জি (6 nm)"/"Qualcomm Snapdragon 695 5G (6 nm)": "কোয়ালকম স্ন্যাপড্রাগন 695 ৫জি (6 nm)"/g'
    
    # Google Tensor - keep Gen designation
    's/"Google Tensor G4": "Google টেনসর G4"/"Google Tensor G4": "গুগল টেনসর G4"/g'
    's/"Google Tensor G4 (4 nm)": "Google টেনসর G4 (4 nm)"/"Google Tensor G4 (4 nm)": "গুগল টেনসর G4 (4 nm)"/g'
    's/"Google Tensor G3": "Google টেনসর G3"/"Google Tensor G3": "গুগল টেনসর G3"/g'
    's/"Google Tensor G3 (4 nm)": "Google টেনসর G3 (4 nm)"/"Google Tensor G3 (4 nm)": "গুগল টেনসর G3 (4 nm)"/g'
    
    # Camera features - fix "ma" typo and translate properly
    's/"OIS (maইঞ্চি), LED flash": "OIS (maইঞ্চি), LED ফ্ল্যাশ"/"OIS, LED flash": "অপটিক্যাল ইমেজ স্ট্যাবিলাইজেশন, LED ফ্ল্যাশ"/g'
)

files_fixed=0

for pattern in "${FIXES[@]}"; do
    echo "Applying fix..."
    find specification_seeder_mobile_*.go -exec sed -i "$pattern" {} +
    ((files_fixed++))
done

echo ""
echo "=== FIX COMPLETE ==="
echo "Patterns applied: ${#FIXES[@]}"
echo ""
echo "Verifying remaining mixed patterns..."
REMAINING=$(grep -h '": "[^"]*[A-Za-z][^"]*[ঁ-ৌ].*[A-Za-z][^"]*"' specification_seeder_mobile_*.go | wc -l)
echo "Remaining mixed patterns: $REMAINING"
