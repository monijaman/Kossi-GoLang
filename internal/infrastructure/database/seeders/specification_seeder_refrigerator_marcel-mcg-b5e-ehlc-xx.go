package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx seeds specifications/options for product 'marcel-mcg-b5e-ehlc-xx'
type SpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx() *SpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx {
	return &SpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcg-b5e-ehlc-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                  "মার্সেল",
		"marcel-mcg-b5e-ehlc-xx":  "মার্সেল-mcg-b5e-ehlc-xx",
		"MCG-B5E-EHLC-XX":         "MCG-B5E-EHLC-XX",
		"Direct Cool":             "ডাইরেক্ট কুল",
		"255 Ltr.":                "২৫৫ লিটার",
		"47±2 Kg":                 "৪৭±২ কেজি",
		"220-240V/ 50Hz":          "২২০-২৪০ভি/ ৫০হার্টজ",
		"50Hz":                    "৫০হার্টজ",
		"RSCR, Inverter":          "আরএসসিআর, ইনভার্টার",
		"Mechanical, Electronics": "মেকানিক্যাল, ইলেকট্রনিক্স",
		"Manual":                  "ম্যানুয়াল",
		"R600a":                   "আর৬০০এ",
		"Wire":                    "ওয়্যার",
		"1":                       "১",
		"No":                      "না",
		"1085 x 725 x 855 mm":     "১০৮৫ x ৭২৫ x ৮৫৫ মিমি",
		"Replacement Guarantee: 1 Year (Condition Apply), Main Parts (Compressor): 12 Years, Door: 3 Years *, Spare Parts: 4 Years *, After Sales Service:5 Years *": "রিপ্লেসমেন্ট গ্যারান্টি: ১ বছর (শর্ত প্রযোজ্য), মূল অংশ (কম্প্রেসার): ১২ বছর, দরজা: ৩ বছর *, স্পেয়ার পার্টস: ৪ বছর *, আফটার সেলস সার্ভিস: ৫ বছর *",
		"12": "১২",
		"Lock: Yes, Interior Lamp: Yes, Handle: Yes, Condenser: Steel, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Exterior Material: Painted Steel (PCM), Interior Material: Embossed Aluminium (Al2), Shelf: Wire/1, Basket: Wire/1, Loading quantity: 32/64/96, Climatic Type: V.0301 N~ST, V.0302 N~ST, V.0401 N~T, V.0501 N~T, Compressor Input Power: V.0301-109W, V.0302-109W, V.0401-133W, V.0501-102W, Wide voltage range: V 03.01, Cooling Effect: Freezer Cabinet Less than -180C, Recommended voltage stabilizer capacity: Not Required for 140V to above (most variants), Warranty Note: This warranty does not cover the following cases: 1. Any damage due to accident, electrical fault, natural causes, negligence or improper installation. 2. Any damage or failure caused by unauthorized modification or alteration. 3. Products with original serial numbers that have been removed, distorted or cannot be readily recognized.": "লক: হ্যাঁ, ইন্টেরিয়র ল্যাম্প: হ্যাঁ, হ্যান্ডেল: হ্যাঁ, কনডেনসার: স্টিল, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], এক্সটেরিয়র ম্যাটেরিয়াল: পেইন্টেড স্টিল (পিসিএম), ইন্টেরিয়র ম্যাটেরিয়াল: এমবসড অ্যালুমিনিয়াম (এল২), শেল্ফ: ওয়্যার/১, বাস্কেট: ওয়্যার/১, লোডিং কোয়ান্টিটি: ৩২/৬৪/৯৬, ক্লাইমেটিক টাইপ: ভি.০৩০১ এন~এসটি, ভি.০৩০২ এন~এসটি, ভি.০৪০১ এন~টি, ভি.০৫০১ এন~টি, কম্প্রেসার ইনপুট পাওয়ার: ভি.০৩০১-১০৯ওয়াট, ভি.০৩০২-১০৯ওয়াট, ভি.০৪০১-১৩৩ওয়াট, ভি.০৫০১-১০২ওয়াট, ওয়াইড ভোল্টেজ রেঞ্জ: ভি ০৩.০১, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮০সি, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ১৪০ভি থেকে উপরে প্রয়োজন নেই (বেশিরভাগ ভেরিয়েন্ট), ওয়ারেন্টি নোট: এই ওয়ারেন্টি নিম্নলিখিত ক্ষেত্রে কভার করে না: ১. দুর্ঘটনা, বৈদ্যুতিক ত্রুটি, প্রাকৃতিক কারণ, অবহেলা বা অনুপযুক্ত ইনস্টলেশনের কারণে কোনো ক্ষতি। ২. অননুমোদিত পরিবর্তন বা পরিবর্ধনের কারণে কোনো ক্ষতি বা ব্যর্থতা। ৩. মূল সিরিয়াল নম্বর যেগুলি সরানো, বিকৃত বা সহজেই চেনা যায় না এমন পণ্য।",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Copper": "কপার",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[Eco-friendly (১০০% CFC &HCFC Free) Green Technology]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "Freezer Cabinet Less than -১৮০CRefrigerator Cabinet ০০Cto +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃Refrigerator Cabinet ০℃ to +৫℃",
		"Mechanical": "মেকানিক্যাল",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes": "হ্যাঁ",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"32/64/96": "৩২/৬৪/৯৬",
		"725": "৭২৫",
		"Embossed Aluminium (Al2)": "Embossed Aluminium (Al২)",
		"N/A": "N/A",
		"Painted Steel (PCM)": "Painted Steel (PCM)",
		"Steel": "স্টীল",
		"V 03.01": "V ০৩.০১",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcg-b5e-ehlc-xx'
func (s *SpecificationSeederRefrigeratorMarcelMcgB5eEhlcXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcg-b5e-ehlc-xx"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}
		return err
	}
	productID := prod.ID

	existingkeyMapping := map[string]uint{
		"Brand":                             310,
		"Model Name":                        316,
		"Door Type":                         142,
		"Capacity":                          138,
		"Refrigerator Capacity":             156,
		"Freezer Capacity":                  146,
		"Energy Efficiency Rating":          143,
		"Energy Star Rating":                144,
		"Annual Energy Consumption":         137,
		"Dimensions":                        25,
		"Weight":                            80,
		"Color":                             311,
		"Compressor Type":                   139,
		"Cooling Technology":                698,
		"Defrost Type":                      141,
		"Temperature Control":               158,
		"Shelf Material":                    699,
		"Number of Shelves":                 154,
		"Door Bins":                         700,
		"Crisper Drawers":                   701,
		"Ice Maker":                         702,
		"Water Dispenser":                   703,
		"Noise Level":                       150,
		"Voltage":                           160,
		"Frequency (Hz)":                    704,
		"App Control":                       705,
		"Voice Assistant Support":           385,
		"Warranty":                          323,
		"Compressor Warranty (Years)":       707,
		"Refrigerant":                       708,
		"Gross Volume":                      709,
		"Net Volume":                        710,
		"Special Features":                  69,
		"Capillary":                         711,
		"Climate Type (SN, N, ST, T)":       712,
		"Compressor Input Power (Watt)":     713,
		"Cooling Effect":                    714,
		"Defrosting (Automatic/ Manual)":    715,
		"Deodorizer":                        716,
		"Depth (mm)":                        717,
		"Door Basket":                       718,
		"Egg Tray or Pocket":                719,
		"Handle (Recessed/ Grip)":           720,
		"Interior Lamp":                     721,
		"Polyurethane Foam Blowing Agent":   722,
		"Reversible Door":                   723,
		"Thermostat":                        724,
		"Vegetable Crisper":                 725,
		"Vegetable Crisper Cover":           726,
		"Loading Capacity (40HQ/40Ft/20Ft)": 727,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Basket": "Wire/1",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "V.0301, N~STV.0302, N~STV.0401, N~TV.0501, N~T",
		"Compressor Input Power (Watt)": "V.0301-109V.0302-109V.0401-133V.0501-102",
		"Compressor Type": "V.0301-RSCRV.0302-RSCRV.0401-InverterV.0501-RSCR",
		"Condenser": "Steel",
		"Cooling Effect": "Freezer Cabinet Less than -180C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "725",
		"Drawer": "No",
		"Exterior Material": "Painted Steel (PCM)",
		"Gross Volume": "255 Ltr.",
		"Gross Weight(Kg)": "53±2",
		"Handle (Recessed/ Grip)": "Yes",
		"Height/mm": "855",
		"Ice/ cold water Dispenser": "No",
		"Interior Lamp": "Yes",
		"Interior Material": "Embossed Aluminium (Al2)",
		"Loading quantity(20ft/40ft.40 HQ)": "32/64/96",
		"Lock": "Yes",
		"Net Volume": "255 Ltr.",
		"Net Weight(Kg)": "47±2",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220-240V/ 50Hz",
		"Recommended voltage stabilizer capacity": "V.0301-Not Required for 140V to aboveV.0302-Not Required for 140V to aboveV.0401-Not ApplicableV.0501-Not Required for 140V to above",
		"Refrigerant": "V.0301 - R600aV.0302 - R600aV.0401 - R600aV.0501 - R600a",
		"Reversible Door": "N/A",
		"Shelf (Material/No)": "Wire/1",
		"Temperature Control (Electronic/  Mechanical):": "V.0301-MechanicalV.0302-MechanicalV.0401-MechanicalV.0501-Electronics",
		"Type": "Direct Cool",
		"Wide voltage range": "V 03.01",
		"Width/mm": "1085",
	}





	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
			continue
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Ensure the ID is set after creation
				if sModel.ID == 0 {
					if err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {
						return err
					}
				}
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
