package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerSC250BXLP seeds specifications/options for product 'singer-sc250bxlp'
type SpecificationSeederRefrigeratorSingerSC250BXLP struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerSC250BXLP creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerSC250BXLP() *SpecificationSeederRefrigeratorSingerSC250BXLP {
	return &SpecificationSeederRefrigeratorSingerSC250BXLP{
		BaseSeeder{name: "Specifications for singer-sc250bxlp"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerSC250BXLP) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                       "সিঙ্গার",
		"SC-250BX-LP":                  "SC-250BX-LP",
		"Beverage Cooler":              "বেভারেজ কুলার",
		"241 Liters":                   "241 লিটার",
		"5 Star":                       "৫ তারা",
		"5":                            "৫",
		"1700 x 560 x 540 mm":          "1700 x 560 x 540 মিমি",
		"Pro Smart Inverter":           "প্রো স্মার্ট ইনভার্টার",
		"NoFrost Dual Cooling":         "নো ফ্রস্ট ডুয়াল কুলিং",
		"No":                           "না",
		"220 ~ 240V":                   "২২০ ~ ২৪০ভোল্ট",
		"50":                           "৫০",
		"10 Years Compressor Warranty": "10 বছর কম্প্রেসার ওয়ারেন্টি",
		"10":                           "10",
		"Refrigerant":                  "রেফ্রিজারেন্ট",
		"Gross Volume":                 "মোট ভলিউম",
		"Net Volume":                   "নেট ভলিউম",
		"275 Ltr.":                     "275 লিটার",
		"241 Ltr.":                     "241 লিটার",
		"R-600a":                       "আর-600এ",
		"Special Features":             "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-sc250bxlp'
func (s *SpecificationSeederRefrigeratorSingerSC250BXLP) Seed(db *gorm.DB) error {
	productSlug := "singer-sc250bxlp"
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
		"Brand":                       "Singer",
		"Model Name":                  "SC-250BX-LP",
		"Door Type":                   "Beverage Cooler",
		"Capacity":                    "241 Liters",
		"Refrigerator Capacity":       "241 Liters",
		"Freezer Capacity":            "",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "",
		"Dimensions":                  "1700 x 560 x 540 mm",
		"Weight":                      "",
		"Color":                       "",
		"Compressor Type":             "Pro Smart Inverter",
		"Cooling Technology":          "NoFrost Dual Cooling",
		"Defrost Type":                "No",
		"Temperature Control":         "",
		"Shelf Material":              "",
		"Number of Shelves":           "",
		"Door Bins":                   "",
		"Crisper Drawers":             "",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "275 Ltr.",
		"Net Volume":                  "241 Ltr.",
		"Special Features":            "NoFrost Dual Cooling, Pro Smart Inverter Compressor, Flush Handle, Quick Cooling & Quick Freezing for Longer Freshness, Child & Door Lock, Door Open Alarm, Eco Mode, Holiday Mode, Adjustable Feet, Fast Freezing, Hot-wall Condenser, High Efficiency Compressor, Quick Chilling, Safety Glass Door, Quiet Technique, Advanced cooling system makes product more effective and economic",
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
