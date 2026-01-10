package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5dGdshXx seeds specifications/options for product 'marcel-mfb-b5d-gdsh-xx'
type SpecificationSeederRefrigeratorMarcelMfbB5dGdshXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5dGdshXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5dGdshXx() *SpecificationSeederRefrigeratorMarcelMfbB5dGdshXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB5dGdshXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5d-gdsh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGdshXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b5d-gdsh-xx": "মার্সেল-mfb-b5d-gdsh-xx",
		"MFB-B5D-GDSH-XX":        "এমএফবি-বি৫ডি-জিডিএসএইচ-এক্সএক্স",
		"1845 mm":                "১৮৪৫ মিমি",
		"220 ~ 240/ 50":          "২২০ ~ ২৪০/ ৫০",
		"254 Ltr.":               "২৫৪ লিটার",
		"268 Ltr.":               "২৬৮ লিটার",
		"5 star (BDS 1850:2012)": "৫ তারা (বিডিএস ১৮৫০:২০১২)",
		"580 mm":                 "৫৮০ মিমি",
		"62 ± 2 Kg":              "৬২ ± ২ কেজি",
		"645 mm":                 "৬৪৫ মিমি",
		"Manual":                 "ম্যানুয়াল",
		"No":                     "না",
		"R600a":                  "আর৬০০এ",
		"V 02.01-RSCR; V 03.01-RSCR; V 03.02-RSCR; V 04.01-RSCR": "ভি ০২.০১-আরএসসিআর; ভি ০৩.০১-আরএসসিআর; ভি ০৩.০২-আরএসসিআর; ভি ০৪.০১-আরএসসিআর",
		"Wire/3": "তার/৩",
		"Yes":    "হ্যাঁ",
		"Type: Direct Cool; Gross Weight: 67 ± 2 Kg; Performance: Climatic Type: N~ST; Compressor Input Power (Watt): V 02.01-132; V 03.01-118; V 03.02-118; V 04.01-127; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Temperature Control: Mechanical; Handle: Recessed/ Grip; Lock: Yes; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Recommended voltage stabilizer capacity: V 02.01-1000VA; V 03.01,V 03.02-Low Voltage(140~260V); For V 03.01,V 03.02 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended; V 04.01-Low Voltage(150~260V); For V 04.01 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended; Interior Lamp: Yes; Egg Tray: Yes; Freezer Shelf: Wire/3; Freezer Drawer: HIPS/5; Freezer Door Pocket: No; Freezer Interior Lamp: No; Packaging Dimensions: 580 x 645 x 1845 mm; Loading Capacity- 40HQ/ 40Ft/ 20Ft: 97/ 72/ 36": "টাইপ: ডিরেক্ট কুল; গ্রস ওজন: ৬৭ ± ২ কেজি; পারফরম্যান্স: ক্লাইম্যাটিক টাইপ: এন~এসটি; কম্প্রেসার ইনপুট পাওয়ার (ওয়াট): ভি ০২.০১-১৩২; ভি ০৩.০১-১১৮; ভি ০৩.০২-১১৮; ভি ০৪.০১-১২৭; কুলিং ইফেক্ট: ফ্রিজার কেবিনেট -১৮℃ এর কম; রেফ্রিজেরেটর কেবিনেট ০℃ থেকে +৫℃; তাপমাত্রা নিয়ন্ত্রণ: ম্যানুয়াল/মেকানিক্যাল; হ্যান্ডেল: রিসেসড/গ্রিপ; লক: হ্যাঁ; থার্মোস্ট্যাট: রোএইচএস সনদপ্রাপ্ত; ক্যাপিলারি: কপার; পলিউরেথেন ফোম ব্লোয়িং এজেন্ট: সাইক্লোপেন্টেন [পরিবেশবান্ধব (১০০% সিএফসি ও এইচসিএফসি মুক্ত) সবুজ প্রযুক্তি]; সুপারিশকৃত ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ভি ০২.০১-১০০০ভিএ; ভি ০৩.০১,ভি ০৩.০২-নিম্ন ভোল্টেজ(১৪০~২৬০ভি); ভি ০৩.০১,ভি ০৩.০২ এর জন্য - ওয়াইড ভোল্টেজ রেঞ্জ (১৪০ভ্যাক - ২৬০ভ্যাক). ভোল্টেজ স্ট্যাবিলাইজার প্রয়োজন নেই। এই রেঞ্জের বাইরে ভোল্টেজের ক্ষেত্রে, ১০০০ভিএ সুপারিশকৃত; ভি ০৪.০১-নিম্ন ভোল্টেজ(১৫০~২৬০ভি); ভি ০৪.০১ এর জন্য - ওয়াইড ভোল্টেজ রেঞ্জ (১৫০ভ্যাক - ২৬০ভ্যাক). ভোল্টেজ স্ট্যাবিলাইজার প্রয়োজন নেই। এই রেঞ্জের বাইরে ভোল্টেজের ক্ষেত্রে, ১০০০ভিএ সুপারিশকৃত; ইন্টারিয়র ল্যাম্প: হ্যাঁ; ডিম ট্রে: হ্যাঁ; ফ্রিজার শেলফ: তার/৩; ফ্রিজার ড্রয়ার: এইচআইপিএস/৫; ফ্রিজার ডোর পকেট: না; ফ্রিজার ইন্টারিয়র ল্যাম্প: না; প্যাকেজিং ডাইমেনশনস: ৫৮০ x ৬৪৫ x ১৮৪৫ মিমি; লোডিং ক্যাপাসিটি- ৪০এইচকিউ/ ৪০এফটি/ ২০এফটি: ৯৭/ ৭২/ ৩৬",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5d-gdsh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGdshXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5d-gdsh-xx"
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
		"Brand":                          "Marcel",
		"Model Name":                     "MFB-B5D-GDSH-XX",
		"Compressor Type":                "V 02.01-RSCR; V 03.01-RSCR; V 03.02-RSCR; V 04.01-RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm":                       "645 mm",
		"Door Pocket":                    "No",
		"Energy Rating":                  "5 star (BDS 1850:2012)",
		"Gross Volume":                   "268 Ltr.",
		"Height/mm":                      "1845 mm",
		"Net Volume":                     "254 Ltr.",
		"Net Weight":                     "62 ± 2 Kg",
		"Rated Voltage/ Hz":              "220 ~ 240/ 50",
		"Refrigerant":                    "R600a",
		"Reversible Door":                "No",
		"Shelf (Material/ No.)":          "Wire/3",
		"Vegetable Crisper":              "Yes",
		"Vegetable Crisper Cover":        "Yes",
		"Width/mm":                       "580 mm",
		"Special Features":               "Type: Direct Cool; Gross Weight: 67 ± 2 Kg; Performance: Climatic Type: N~ST; Compressor Input Power (Watt): V 02.01-132; V 03.01-118; V 03.02-118; V 04.01-127; Cooling Effect: Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃; Temperature Control: Mechanical; Handle: Recessed/ Grip; Lock: Yes; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Recommended voltage stabilizer capacity: V 02.01-1000VA; V 03.01,V 03.02-Low Voltage(140~260V); For V 03.01,V 03.02 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended; V 04.01-Low Voltage(150~260V); For V 04.01 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended; Interior Lamp: Yes; Egg Tray: Yes; Freezer Shelf: Wire/3; Freezer Drawer: HIPS/5; Freezer Door Pocket: No; Freezer Interior Lamp: No; Packaging Dimensions: 580 x 645 x 1845 mm; Loading Capacity- 40HQ/ 40Ft/ 20Ft: 97/ 72/ 36",
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
