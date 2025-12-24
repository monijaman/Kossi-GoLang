package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWfa2a3GdelSc seeds specifications/options for product 'walton-wfa-2a3-gdel-sc'
type SpecificationSeederRefrigeratorWaltonWfa2a3GdelSc struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWfa2a3GdelSc creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWfa2a3GdelSc() *SpecificationSeederRefrigeratorWaltonWfa2a3GdelSc {
	return &SpecificationSeederRefrigeratorWaltonWfa2a3GdelSc{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wfa-2a3-gdel-sc"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWfa2a3GdelSc) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WCF-1B5-GDEL-XX":     "ডব্লিউসিএফ-১বি৫-জিডিইএল-এক্সএক্স",
		"Single Door":         "একটি দরজা",
		"150 Liters":          "১৫০ লিটার",
		"125 Liters":          "১২৫ লিটার",
		"25 Liters":           "২৫ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"200 kWh":             "২০০ কিলোওয়াট ঘণ্টা",
		"500 x 520 x 1200 mm": "৫০০ x ৫২০ x ১২০০ মিমি",
		"38 kg":               "৩৮ কেজি",
		"Silver":              "সিলভার",
		"Reciprocating":       "রিসিপ্রোকেটিং",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"2":                   "২",
		"1":                   "১",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"Large Vegetable Crisper, Egg Tray, Ice Cube Tray": "বড় শাকসবজি ক্রিস্পার, ডিমের ট্রে, বরফের কিউব ট্রে",
		"Refrigerant":         "রেফ্রিজারেন্ট",
		"Gross Volume":        "মোট ভলিউম",
		"Net Volume":          "নেট ভলিউম",
		"213 Liters":          "২১৩ লিটার",
		"176 Liters":          "১৭৬ লিটার",
		"RSCR":                "আরএসসিআর",
		"Wire":                "ওয়াইয়ার",
		"GPPS":                "জিপিপিএস",
		"557 x 645 x 1510 mm": "৫৫৭ x ৬৪৫ x ১৫১০ মিমি",
		"45.5 kg":             "৪৫.৫ কেজি",
		"50 kg":               "৫০ কেজি",
		"220 ~ 240V":          "২২০ ~ ২৪০ভোল্ট",
		"R600a":               "আর৬০০এ",
		"N~ST":                "এন~এসটি",
		"5 Star":              "৫ তারা",
		"Electronic":          "ইলেকট্রনিক",
		"Copper":              "কপার",
		"CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "সাইক্লোপেন্টিন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি]",
		"Recessed": "রিসেসড",
		"Yes":      "হ্যাঁ",
		"3":        "৩",
		"5":        "৫",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"Wire/2": "ওয়াইয়ার/২",
		"GPPS/3": "জিপিপিএস/৩",
		"Eco-friendly Green Technology, Vegetable Crisper, Egg Tray": "ইকো-ফ্রেন্ডলি গ্রিন টেকনোলজি, শাকসবজি ক্রিস্পার, ডিমের ট্রে",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wfa-2a3-gdel-sc'
func (s *SpecificationSeederRefrigeratorWaltonWfa2a3GdelSc) Seed(db *gorm.DB) error {
	productSlug := "walton-wfa-2a3-gdel-sc"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
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
		"Brand":                       "Walton",
		"Model Name":                  "WFA-2A3-GDEL-SC",
		"Door Type":                   "Single Door",
		"Capacity":                    "213 Liters",
		"Refrigerator Capacity":       "176 Liters",
		"Freezer Capacity":            "N/A",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "557 x 645 x 1510 mm",
		"Weight":                      "45.5 kg",
		"Color":                       "Silver",
		"Compressor Type":             "RSCR",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "4",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "213 Liters",
		"Net Volume":                  "176 Liters",
		"Special Features":            "Eco-friendly Green Technology, Vegetable Crisper, Egg Tray",
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
				if err := db.Create(sModel).Last(&sModel).Error; err != nil {
					return err
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
