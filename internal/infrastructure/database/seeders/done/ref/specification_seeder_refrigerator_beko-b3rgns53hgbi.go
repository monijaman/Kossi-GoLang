package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorBekoB3rgns53hgbi seeds specifications/options for product 'b3rgns53hgbi'
type SpecificationSeederRefrigeratorBekoB3rgns53hgbi struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorBekoB3rgns53hgbi creates a new seeder instance
func NewSpecificationSeederRefrigeratorBekoB3rgns53hgbi() *SpecificationSeederRefrigeratorBekoB3rgns53hgbi {
	return &SpecificationSeederRefrigeratorBekoB3rgns53hgbi{
		BaseSeeder: BaseSeeder{name: "Specifications for b3rgns53hgbi"},
	}
}

func (s *SpecificationSeederRefrigeratorBekoB3rgns53hgbi) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Beko":                "বেকো",
		"B3RGNS53HGBI":        "বি৩আরজিএনএস৫৩এইচজিবিআই",
		"Side by Side":        "সাইড বাই সাইড",
		"530 Liters":          "৫৩০ লিটার",
		"320 Liters":          "৩২০ লিটার",
		"210 Liters":          "২১০ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"280 kWh":             "২৮০ কিলোওয়াট ঘণ্টা",
		"700 x 750 x 1850 mm": "৭০০ x ৭৫০ x ১৮৫০ মিমি",
		"85 kg":               "৮৫ কেজি",
		"Black":               "কালো",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"6":                   "৬",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"12":                  "১২",
		"Inverter Technology, Frost Free, Ice Dispenser, Water Dispenser, LED Lighting": "ইনভার্টার প্রযুক্তি, ফ্রস্ট ফ্রি, আইস ডিসপেন্সার, ওয়াটার ডিসপেন্সার, এলইডি লাইটিং",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"R-600a":           "আর-৬০০এ",
		"530 Ltr.":         "৫৩০ লিটার",
		"220 ~ 240V":       "২২০ ~ ২৪০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'b3rgns53hgbi'
func (s *SpecificationSeederRefrigeratorBekoB3rgns53hgbi) Seed(db *gorm.DB) error {
	productSlug := "b3rgns53hgbi"
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
		"Brand":                       "Beko",
		"Model Name":                  "B3RGNS53HGBI",
		"Door Type":                   "Side by Side",
		"Capacity":                    "530 Liters",
		"Refrigerator Capacity":       "320 Liters",
		"Freezer Capacity":            "210 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "280 kWh",
		"Dimensions":                  "700 x 750 x 1850 mm",
		"Weight":                      "85 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "6",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "Yes",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "530 Ltr.",
		"Net Volume":                  "530 Ltr.",
		"Special Features":            "Inverter Technology, Frost Free, Ice Dispenser, Water Dispenser, LED Lighting",
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
