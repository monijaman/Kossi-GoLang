package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWni6a9GdsdDd seeds specifications/options for product 'walton-wni-6a9-gdsd-dd'
type SpecificationSeederRefrigeratorWaltonWni6a9GdsdDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWni6a9GdsdDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWni6a9GdsdDd() *SpecificationSeederRefrigeratorWaltonWni6a9GdsdDd {
	return &SpecificationSeederRefrigeratorWaltonWni6a9GdsdDd{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wni-6a9-gdsd-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWni6a9GdsdDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WNI-6A9-GDSD-DD":     "ডব্লিউএনআই-৬এ৯-জিডিএসডি-ডিডি",
		"Double Door":         "ডাবল দরজা",
		"619 Liters":          "৬১৯ লিটার",
		"582 Liters":          "৫৮২ লিটার",
		"N/A":                 "এন/এ",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"865 x 755 x 1780 mm": "৮৬৫ x ৭৫৫ x ১৭৮০ মিমি",
		"116 kg":              "১১৬ কেজি",
		"Depends on Stock":    "স্টকের উপর নির্ভর করে",
		"BLDC Inverter":       "বিএলডিসি ইনভার্টার",
		"No-Frost":            "নো-ফ্রস্ট",
		"Automatic":           "অটোমেটিক",
		"Electronic":          "ইলেকট্রনিক",
		"Glass":               "গ্লাস",
		"10":                  "১০",
		"1":                   "১",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"12":                  "১২",
		"Electronic Temperature Control, LED Interior Lamp, Vegetable Crisper, Egg Tray, Ice Case, Ice Box": "ইলেকট্রনিক তাপমাত্রা নিয়ন্ত্রণ, এলইডি ইন্টেরিয়র ল্যাম্প, শাকসবজি ক্রিস্পার, ডিমের ট্রে, বরফের কেস, বরফের বাক্স",
		"Refrigerant":  "রেফ্রিজারেন্ট",
		"Gross Volume": "মোট ভলিউম",
		"Net Volume":   "নেট ভলিউম",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wni-6a9-gdsd-dd'
func (s *SpecificationSeederRefrigeratorWaltonWni6a9GdsdDd) Seed(db *gorm.DB) error {
	productSlug := "walton-wni-6a9-gdsd-dd"
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
		"Model Name":                  "WNI-6A9-GDSD-DD",
		"Door Type":                   "Double Door",
		"Capacity":                    "619 Liters",
		"Refrigerator Capacity":       "N/A",
		"Freezer Capacity":            "N/A",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "865 x 755 x 1780 mm",
		"Weight":                      "116 kg",
		"Color":                       "Depends on Stock",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "No-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Glass",
		"Number of Shelves":           "10",
		"Door Bins":                   "10",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "N/A",
		"Gross Volume":                "N/A",
		"Net Volume":                  "N/A",
		"Special Features":            "Electronic Temperature Control, LED Interior Lamp, Vegetable Crisper, Egg Tray, Ice Case, Ice Box",
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
