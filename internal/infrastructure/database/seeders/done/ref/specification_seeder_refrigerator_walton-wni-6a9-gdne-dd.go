package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWni6a9GdneDd seeds specifications/options for product 'walton-wni-6a9-gdne-dd'
type SpecificationSeederRefrigeratorWaltonWni6a9GdneDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWni6a9GdneDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWni6a9GdneDd() *SpecificationSeederRefrigeratorWaltonWni6a9GdneDd {
	return &SpecificationSeederRefrigeratorWaltonWni6a9GdneDd{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wni-6a9-gdne-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWni6a9GdneDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WNI-6A9-GDNE-DD":     "ডব্লিউএনআই-৬এ৯-জিডিএনই-ডিডি",
		"Double Door":         "ডাবল দরজা",
		"591 Liters":          "৫৯১ লিটার",
		"865 x 725 x 1780 mm": "৮৬৫ x ৭২৫ x ১৭৮০ মিমি",
		"113 kg":              "১১৩ কেজি",
		"BLDC Inverter":       "বিএলডিসি ইনভার্টার",
		"No-Frost":            "নো-ফ্রস্ট",
		"Automatic":           "অটোমেটিক",
		"Electronic":          "ইলেকট্রনিক",
		"Glass":               "গ্লাস",
		"5":                   "৫",
		"4":                   "৪",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"220-240V":            "২২০-২৪০ভোল্ট",
		"50":                  "৫০",
		"1 Year":              "১ বছর",
		"12":                  "১২",
		"R600a":               "আর৬০০এ",
		"548 Liters":          "৫৪৮ লিটার",
		"Refrigerant":         "রেফ্রিজারেন্ট",
		"Gross Volume":        "মোট ভলিউম",
		"Net Volume":          "নেট ভলিউম",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wni-6a9-gdne-dd'
func (s *SpecificationSeederRefrigeratorWaltonWni6a9GdneDd) Seed(db *gorm.DB) error {
	productSlug := "walton-wni-6a9-gdne-dd"
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
		"Model Name":                  "WNI-6A9-GDNE-DD",
		"Door Type":                   "Double Door",
		"Capacity":                    "591 Liters",
		"Refrigerator Capacity":       "",
		"Freezer Capacity":            "",
		"Energy Efficiency Rating":    "",
		"Energy Star Rating":          "",
		"Annual Energy Consumption":   "",
		"Dimensions":                  "865 x 725 x 1780 mm",
		"Weight":                      "113 kg",
		"Color":                       "",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "No-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Glass",
		"Number of Shelves":           "5",
		"Door Bins":                   "4",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "",
		"Voltage":                     "220-240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "",
		"Voice Assistant Support":     "",
		"Warranty":                    "1 Year",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "591 Liters",
		"Net Volume":                  "548 Liters",
		"Special Features":            "",
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
