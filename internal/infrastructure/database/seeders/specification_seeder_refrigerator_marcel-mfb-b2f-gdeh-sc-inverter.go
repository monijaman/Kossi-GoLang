package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter seeds specifications/options for product 'marcel-mfb-b2f-gdeh-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter() *SpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-gdeh-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfb-b2f-gdeh-sc-inverter": "মার্সেল-mfb-b2f-gdeh-sc-inverter",
		"MFB-B2F-GDEH-SC-INVERTER":        "MFB-B2F-GDEH-SC-INVERTER",

		"Direct Cool":         "সরাসরি কুলিং",
		"177 Ltr.":            "১৭৭ লিটার",
		"175 Ltr.":            "১৭৫ লিটার",
		"50 ± 2 Kg":           "৫০ ± ২ কেজি",
		"R600a":               "R600a",
		"Mechanical":          "ম্যানুয়াল",
		"Manual":              "ম্যানুয়াল",
		"Recessed/ Grip":      "রিসেসড/গ্রিপ",
		"Copper":              "তামা",
		"Cyclopentene":        "সাইক্লোপেন্টেন",
		"220 ~ 240":           "২২০ ~ ২৪০ ভোল্ট",
		"Hermetic":            "হার্মেটিক",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"Compressor: V 01.01-RSCR; V 01.02-RSCR; Compressor Input Power (Watt): V 01.01-88; V 01.02-88; Voltage/Hz: 220 ~ 240/ 50; Stabilizer: 5 Ampere": "কমপ্রেসার: V 01.01-RSCR; V 01.02-RSCR; কমপ্রেসার ইনপুট পাওয়ার (ওয়াট): V 01.01-88; V 01.02-88; ভোল্টেজ/হার্জ: 220 ~ 240/50; স্ট্যাবিলাইজার: ৫ অ্যাম্পিয়ার",
		"66 ± 2 Kg": "৬৬ ± ২ কেজি",
		"5 - Star":  "৫-স্টার",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার অংশ −১৮℃ এর নিচে; রেফ্রিজারেটর অংশ ০℃ থেকে +৫℃",
		"Yes":        "হ্যাঁ",
		"645 mm":     "৬৪৫ মিমি",
		"Electronic": "ইলেকট্রনিক",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "সাইক্লোপেন্টেন (ইকো-ফ্রেন্ডলি, ১০০% CFC/HCFC মুক্ত)",
		"220 ~ 240/ 50":  "২২০ ~ ২৪০ ভোল্ট, ৫০ হার্টজ",
		"HIPS/5":         "HIPS (৫টি)",
		"BLDC Inverter":  "বিএলডিসি ইনভার্টার",
		"Wire/3":         "তারের তাক (৩টি)",
		"RoHS Certified": "RoHS সার্টিফায়েড",
		"1765 mm":        "১৭৬৫ মিমি",
		"580 mm":         "৫৮০ মিমি",
		"V 1101-57~125":  "V ১১০১-৫৭~১২৫",
		"238 Ltr.":       "২৩৮ লিটার",
		"Wire/2":         "তারের তাক (২টি)",
		"N~T":            "N~T",
		"Recessed/Grip":  "রিসেসড/গ্রিপ",
		"252 Ltr.":       "২৫২ লিটার",
		"GPPS/4":         "GPPS (৪টি)",
		"59 ± 2 Kg":      "৫৯ ± ২ কেজি",
		"V 1101-Low Voltage(90~300V), Voltage stabilizer is not required.": "V ১১০১ - লো ভোল্টেজ (৯০~৩০০V), সাধারণত ভোল্টেজ স্ট্যাবিলাইজার প্রয়োজন নেই।",
		"98/ 72/ 36": "৯৮/৭২/৩৬",

		// Special features long mapping
		`Gross Weight: 66 ± 2 Kg; Packing Dimensions: 580 x 645 x 1765 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 98/ 72/ 36; Climatic Type: N~T; Compressor Input Power (Watt): V 1101-57~125; Compressor Models: V 1101 variants; Compressor Type: BLDC Inverter; Energy Rating: 5 - Star; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene [Eco-friendly (100% CFC & HCFC Free)]; Recommended voltage stabilizer capacity: V 1101-Low Voltage(90~300V) — stabilizer usually not required; Refrigerator Compartment: Shelf (Wire/2); Door Basket: GPPS/4; Interior Lamp: Yes; Vegetable Crisper: Yes; Vegetable Crisper Cover: Yes; Egg Tray: Yes; Freezer Compartment: Shelf (Wire/3); Drawer: HIPS/5;`: `গ্রস ওজন: ৬৬ ± ২ কেজি; প্যাকেজিং মাত্রা: ৫৮০ x ৬৪৫ x ১৭৬৫ মিমি; লোডিং ক্যাপাসিটি (40HQ/40Ft/20Ft): ৯৮/৭২/৩৬; আবহাওয়া টাইপ: N~T; কমপ্রেসার ইনপুট পাওয়ার (ওয়াট): V 1101-57~125; কম্প্রেসার মডেল: V 1101 সিরিজ; কম্প্রেসার টাইপ: বিএলডিসি ইনভার্টার; এনার্জি রেটিং: ৫-স্টার; কুলিং পারফরম্যান্স: ফ্রিজার অংশ −১৮℃ এর নিচে; রেফ্রিজারেটর অংশ ০℃ থেকে +৫℃; থার্মোস্ট্যাট: RoHS সার্টিফায়েড; ক্যাপিলারি: তামা; পলিউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন (ইকো-ফ্রেন্ডলি, CFC/HCFC মুক্ত); সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার: V 1101 (লো ভোল্টেজ 90~300V) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন নেই; রেফ্রিজারেটর অংশ: তাক (তারের) (২টি); দরজার বাস্কেট: GPPS (৪টি); অভ্যন্তরীণ বাতি: হ্যাঁ; সবজি ক্রিস্পার: হ্যাঁ; সবজি ক্রিস্পার ঢাকনা: হ্যাঁ; ডিম ধারক: হ্যাঁ; ফ্রিজার অংশ: তাক (তারের) (৩টি); ড্রয়ার: HIPS (৫টি);`,
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-gdeh-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdehScInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-gdeh-sc-inverter"
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
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}

	specs := map[string]string{
		"Brand":                    "Marcel",
		"Model Name":               "MFB-B2F-GDEH-SC-INVERTER",
		"Cooling Technology":       "Direct Cool",
		"Gross Volume":             "252 Ltr.",
		"Net Volume":               "238 Ltr.",
		"Weight":                   "59 ± 2 Kg",
		"Refrigerant":              "R600a",
		"Temperature Control":      "Electronic",
		"Defrost Type":             "Manual",
		"Door Type":                "Recessed/Grip",
		"Voltage":                  "220 ~ 240/ 50",
		"Dimensions":               "555 x 630 x 1720 mm",
		"Compressor Type":          "BLDC Inverter",
		"Energy Efficiency Rating": "5 - Star",
		"Cooling Effect":           "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Shelf Material":           "Wire/2",
		"Door Bins":                "GPPS/4",
		"Crisper Drawers":          "Yes",
		"Special Features":         `Gross Weight: 66 ± 2 Kg; Packing Dimensions: 580 x 645 x 1765 mm; Loading Capacity (40HQ/ 40Ft/ 20Ft): 98/ 72/ 36; Climatic Type: N~T; Compressor Input Power (Watt): V 1101-57~125; Compressor Models: V 1101 variants; Compressor Type: BLDC Inverter; Energy Rating: 5 - Star; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene [Eco-friendly (100% CFC & HCFC Free)]; Recommended voltage stabilizer capacity: V 1101-Low Voltage(90~300V) — stabilizer usually not required; Refrigerator Compartment: Shelf (Wire/2); Door Basket: GPPS/4; Interior Lamp: Yes; Vegetable Crisper: Yes; Vegetable Crisper Cover: Yes; Egg Tray: Yes; Freezer Compartment: Shelf (Wire/3); Drawer: HIPS/5;`,
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
