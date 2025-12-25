package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx seeds specifications/options for product 'walton-wcf-1d5-rrxx-xx'
type SpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx() *SpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx {
	return &SpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wcf-1d5-rrxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":                        "ওয়ালটন",
		"WCF-1D5-RRXX-XX":               "ডব্লিউসিএফ-১ডি৫-আরআরএক্সএক্স-এক্সএক্স",
		"Single Door":                   "একটি দরজা",
		"145 Liters":                    "১৪৫ লিটার",
		"N/A":                           "এন/এ",
		"Direct Cool":                   "ডাইরেক্ট কুল",
		"Manual":                        "ম্যানুয়াল",
		"Mechanical":                    "মেকানিক্যাল",
		"744 x 585 x 845 mm":            "৭৪৪ x ৫৮৫ x ৮৪৫ মিমি",
		"34 ± 2 kg":                     "৩৪ ± ২ কেজি",
		"RSCR":                          "আরএসসিআর",
		"Wire":                          "ওয়াইয়ার",
		"1":                             "১",
		"Yes":                           "হ্যাঁ",
		"220 ~ 240V":                    "২২০ ~ ২৪০ভোল্ট",
		"50":                            "৫০",
		"R600a":                         "আর৬০০এ",
		"145 Ltr.":                      "১৪৫ লিটার",
		"Eco-friendly Green Technology": "ইকো-ফ্রেন্ডলি গ্রিন টেকনোলজি",
		"Refrigerant":                   "রেফ্রিজারেন্ট",
		"Gross Volume":                  "মোট ভলিউম",
		"Net Volume":                    "নেট ভলিউম",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wcf-1d5-rrxx-xx'
func (s *SpecificationSeederRefrigeratorWaltonWcf1d5RrxxXx) Seed(db *gorm.DB) error {
	productSlug := "walton-wcf-1d5-rrxx-xx"
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
		"Model Name":                  "WCF-1D5-RRXX-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "145 Liters",
		"Refrigerator Capacity":       "N/A",
		"Freezer Capacity":            "145 Liters",
		"Energy Efficiency Rating":    "N/A",
		"Energy Star Rating":          "N/A",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "744 x 585 x 845 mm",
		"Weight":                      "34 ± 2 kg",
		"Color":                       "N/A",
		"Compressor Type":             "RSCR",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "1",
		"Door Bins":                   "N/A",
		"Crisper Drawers":             "N/A",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "N/A",
		"Compressor Warranty (Years)": "N/A",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "145 Ltr.",
		"Net Volume":                  "145 Ltr.",
		"Special Features":            "Eco-friendly Green Technology",
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
