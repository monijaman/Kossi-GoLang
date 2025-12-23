package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd seeds specifications/options for product 'walton-wnr-6d6-gdfs-dd'
type SpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd() *SpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd {
	return &SpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wnr-6d6-gdfs-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":          "ওয়ালটন",
		"WNR-6D6-GDFS-DD": "ডব্লিউএনআর-৬ডি৬-জিডিএফএস-ডিডি",
		"French Door":     "ফ্রেঞ্চ দরজা",
		"606 Liters":      "৬০৬ লিটার",
		"350 Liters":      "৩৫০ লিটার",
		"256 Liters":      "২৫৬ লিটার",
		"5 Star":          "৫ তারা",
		"5":               "৫",
		"N/A":             "এন/এ",
		"Silver":          "সিলভার",
		"Inverter":        "ইনভার্টার",
		"No-Frost":        "নো-ফ্রস্ট",
		"Automatic":       "অটোমেটিক",
		"Electronic":      "ইলেকট্রনিক",
		"Glass":           "গ্লাস",
		"4":               "৪",
		"6":               "৬",
		"2":               "২",
		"Yes":             "হ্যাঁ",
		"No":              "না",
		"220V":            "২২০ভোল্ট",
		"50":              "৫০",
		"2 Years":         "২ বছর",
		"12":              "১২",
		"No-Frost, Inverter Compressor, LED Lighting, Electronic Control": "নো-ফ্রস্ট, ইনভার্টার কম্প্রেসার, এলইডি লাইটিং, ইলেকট্রনিক কন্ট্রোল",
		"Refrigerant":         "রেফ্রিজারেন্ট",
		"Gross Volume":        "মোট ভলিউম",
		"Net Volume":          "নেট ভলিউম",
		"R600a (CFC Free)":    "আর৬০০এ (সিএফসি ফ্রি)",
		"550 Liters":          "৫৫০ লিটার",
		"900 x 700 x 1800 mm": "৯০০ x ৭০০ x ১৮০০ মিমি",
		"120 kg":              "১২০ কেজি",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wnr-6d6-gdfs-dd'
func (s *SpecificationSeederRefrigeratorWaltonWnr6d6GdfsDd) Seed(db *gorm.DB) error {
	productSlug := "walton-wnr-6d6-gdfs-dd"
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
		"Model Name":                  "WNR-6D6-GDFS-DD",
		"Door Type":                   "French Door",
		"Capacity":                    "606 Liters",
		"Refrigerator Capacity":       "350 Liters",
		"Freezer Capacity":            "256 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "900 x 700 x 1800 mm",
		"Weight":                      "120 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "No-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "6",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R600a (CFC Free)",
		"Gross Volume":                "606 Liters",
		"Net Volume":                  "550 Liters",
		"Special Features":            "No-Frost, Inverter Compressor, LED Lighting, Electronic Control",
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
