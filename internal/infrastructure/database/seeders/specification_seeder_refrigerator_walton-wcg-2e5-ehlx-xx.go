package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx seeds specifications/options for product 'walton-wcg-2e5-ehlx-xx'
type SpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx() *SpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx {
	return &SpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wcg-2e5-ehlx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WCG-2E5-EHLX-XX":     "ডব্লিউসিজি-২ই৫-ইএইচএলএক্স-এক্সএক্স",
		"Single Door":         "একটি দরজা",
		"255 Liters":          "২৫৫ লিটার",
		"N/A":                 "এন/এ",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"1085 x 725 x 855 mm": "১০৮৫ x ৭২৫ x ৮৫৫ মিমি",
		"47 ± 2 kg":           "৪৭ ± ২ কেজি",
		"RSCR":                "আরএসসিআর",
		"Wire":                "ওয়াইয়ার",
		"1":                   "১",
		"Yes":                 "হ্যাঁ",
		"220 ~ 240V":          "২২০ ~ ২৪০ভোল্ট",
		"50":                  "৫০",
		"R600a":               "আর৬০০এ",
		"255 Ltr.":            "২৫৫ লিটার",
		"RoHS Certified, Eco-friendly Green Technology": "রোএইচএস সার্টিফাইড, ইকো-ফ্রেন্ডলি গ্রিন টেকনোলজি",
		"Refrigerant":  "রেফ্রিজারেন্ট",
		"Gross Volume": "মোট ভলিউম",
		"Net Volume":   "নেট ভলিউম",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wcg-2e5-ehlx-xx'
func (s *SpecificationSeederRefrigeratorWaltonWcg2e5EhlxXx) Seed(db *gorm.DB) error {
	productSlug := "walton-wcg-2e5-ehlx-xx"
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
		"Model Name":                  "WCG-2E5-EHLX-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "255 Liters",
		"Refrigerator Capacity":       "N/A",
		"Freezer Capacity":            "255 Liters",
		"Energy Efficiency Rating":    "N/A",
		"Energy Star Rating":          "N/A",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "1085 x 725 x 855 mm",
		"Weight":                      "47 ± 2 kg",
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
		"Gross Volume":                "255 Ltr.",
		"Net Volume":                  "255 Ltr.",
		"Special Features":            "RoHS Certified, Eco-friendly Green Technology",
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
