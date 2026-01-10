package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx seeds specifications/options for product 'marcel-mfb-b2f-gdeh-xx'
type SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fGdehXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fGdehXx() *SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-gdeh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b2f-gdeh-xx": "মার্সেল-mfb-b2f-gdeh-xx",
		"MFB-B2F-GDEH-XX":        "MFB-B2F-GDEH-XX",

		"Direct Cool":   "সরাসরি কুলিং",
		"252 Ltr.":      "২৫২ লিটার",
		"238 Ltr.":      "২৩৮ লিটার",
		"59 ± 2 Kg":     "৫৯ ± ২ কেজি",
		"66 ± 2 Kg":     "৬৬ ± ২ কেজি",
		"N~ST":          "N~ST",
		"220 ~ 240/ 50": "২২০ ~ ২৪০ ভোল্ট, ৫০ হার্টজ",
		"V 06.01-124":   "V ০৬.০১-১২৪",
		"V 07.01-109":   "V ০৭.০১-১০৯",
		"V 07.02-109":   "V ০৭.০২-১০৯",
		"V 07.03-109":   "V ০৭.০৩-১০৯",
		"V 08.01-127":   "V ০৮.০১-১২৭",
		"RSCR":          "RSCR",
		"N/A":           "N/A",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার অংশ −১৮℃ এর নিচে; রেফ্রিজারেটর অংশ ০℃ থেকে +৫℃",
		"Mechanical":     "ম্যানুয়াল",
		"Manual":         "ম্যানুয়াল",
		"Recessed/Grip":  "রিসেসড/গ্রিপ",
		"Lock":           "হ্যাঁ",
		"R600a":          "R600a",
		"RoHS Certified": "RoHS সার্টিফায়েড",
		"Copper":         "তামা",
		"Cyclopentene":   "সাইক্লোপেন্টেন",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "সাইক্লোপেন্টেন (ইকো-ফ্রেন্ডলি, ১০০% CFC/HCFC মুক্ত)",
		"555 x 630 x 1720 mm": "৫৫৫ x ৬৩০ x ১৭২০ মিমি",
		"580 x 645 x 1765 mm": "৫৮০ x ৬৪৫ x ১৭৬৫ মিমি",
		"Wire/2":              "তারের তাক (২টি)",
		"Wire/3":              "তারের তাক (৩টি)",
		"GPPS/4":              "GPPS (৪টি)",
		"HIPS/5":              "HIPS (৫টি)",
		"Yes":                 "হ্যাঁ",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"1765 mm":             "১৭৬৫ মিমি",
		"580 mm":              "৫৮০ মিমি",
		"645 mm":              "৬৪৫ মিমি",
		"98/ 72/ 36":          "৯৮/৭২/৩৬",

		// Special features long translation
		`Compressor Input Power (Watt): V 06.01-124; V 07.01-109; V 07.02-109; V 07.03-109; V 08.01-127. Recommended voltage stabilizer capacity: V 06.01 - Low Voltage(130~260V) (Wide Range 130Vac–260Vac) — stabilizer usually not required; V 07.01/V 07.02/V 07.03 - Low Voltage(140~260V) (Wide Range 140Vac–260Vac) — stabilizer usually not required; V 08.01 - Low Voltage(150~260V) (Wide Range 150Vac–260Vac) — stabilizer usually not required; In case voltages exceed these ranges, 1000VA recommended; Refrigerator Compartment: Shelf (Wire/2); Door Basket: GPPS/4; Interior Lamp: Yes; Vegetable Crisper: Yes; Vegetable Crisper Cover: Yes; Egg Tray: Yes; Freezer Compartment: Shelf (Wire/3); Drawer: HIPS/5; Loading Capacity: 98/ 72/ 36`: `কম্প্রেসার ইনপুট পাওয়ার (ওয়াট): V ০৬.০১-১২৪; V ০৭.০১-১০৯; V ০৭.০২-১০৯; V ০৭.০৩-১০৯; V ০৮.০১-১২৭। সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার: V ০৬.০১ - লো ভোল্টেজ (১৩০~২৬০V) (ওয়াইড রেঞ্জ ১৩০Vac–২৬০Vac) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন হয় না; V ০৭.০১/০৭.০২/০৭.০৩ - লো ভোল্টেজ (১৪০~২৬০V) (ওয়াইড রেঞ্জ ১৪০Vac–২৬০Vac) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন হয় না; V ০৮.০১ - লো ভোল্টেজ (১৫০~২৬০V) (ওয়াইড রেঞ্জ ১৫০Vac–২৬০Vac) — সাধারণত স্ট্যাবিলাইজার প্রয়োজন হয় না; যদি ভোল্টেজ এই রেঞ্জের বাইরে যায়, তবে ১০০০VA স্ট্যাবিলাইজার ব্যবহারের পরামর্শ দেওয়া হয়; রেফ্রিজারেটর অংশ: তাক (তারের) (২টি); দরজার বাস্কেট: GPPS (৪টি); অভ্যন্তরীণ বাতি: হ্যাঁ; সবজি ক্রিস্পার: হ্যাঁ; সবজি ক্রিস্পার ঢাকনা: হ্যাঁ; ডিম ধারক: হ্যাঁ; ফ্রিজার অংশ: তাক (তারের) (৩টি); ড্রয়ার: HIPS (৫টি); লোডিং ক্যাপাসিটি: ৯৮/৭২/৩৬`,
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-gdeh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-gdeh-xx"
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
		"Capillary":                          "Copper",
		"Climatic Type (SN, N, ST, T)":       "N~ST",
		"Compressor Input Power (Watt)":      "V 06.01-124V 07.01-109V 07.02-109V 07.03-109V 08.01-127",
		"Compressor Type":                    "RSCR",
		"Cooling Effect":                     "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual):":    "Manual",
		"Depth/mm":                           "630 mm",
		"Door Basket":                        "GPPS/4",
		"Drawer":                             "HIPS/5",
		"Egg Tray":                           "Yes",
		"Energy Rating":                      "N/A",
		"Gross Volume":                       "252 Ltr.",
		"Gross Weight":                       "66 ± 2 Kg",
		"Handle (Recessed/ Grip)":            "Recessed/Grip",
		"Height/mm":                          "1720 mm",
		"Interior Lamp":                      "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "98/ 72/ 36",
		"Lock":                               "Yes",
		"Net Volume":                         "238 Ltr.",
		"Net Weight":                         "59 ± 2 Kg",
		"Polyurethane foam blowing agent":    "Cyclopentene [Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz":                  "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 06.01-Low Voltage(130~260V)For  V 06.01- Wide Voltage Range (130Vac - 260Vac). Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended.  V 07.01,V 07.02,V 07.03-Low Voltage(140~260V)For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac - 260Vac).Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommendedV 08.01-Low Voltage(150~260V)For V 08.01 - Wide Voltage Range (150Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended.", "Refrigerant": "R600a",
		"Shelf (Material/ No.)":                         "Wire/3",
		"Shelf (Material/No.)":                          "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat":                                    "RoHS Certified",
		"Type":                                          "Direct Cool",
		"Vegetable Crisper":                             "Yes",
		"Vegetable Crisper Cover":                       "Yes",
		"Width/mm":                                      "555 mm",
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
