package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS155VBG seeds specifications/options for product 'singer-ftds155v-bg'
type SpecificationSeederRefrigeratorSingerFTDS155VBG struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS155VBG creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS155VBG() *SpecificationSeederRefrigeratorSingerFTDS155VBG {
	return &SpecificationSeederRefrigeratorSingerFTDS155VBG{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds155v-bg"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS155VBG) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                   "সিঙ্গার",
		"FTDS155-BG":               "FTDS155-BG",
		"Top Mounted Refrigerator": "টপ মাউন্টেড রেফ্রিজারেটর",
		"156.5 Liters":             "156.5 লিটার",
		"90.5 Liters":              "90.5 লিটার",
		"66 Liters":                "66 লিটার",
		"5 Star":                   "৫ তারা",
		"5":                        "৫",
		"10":                       "10",
		"250 kWh":                  "২৫০ কিলোওয়াট ঘণ্টা",
		"1380 x 545 x 544 mm":      "1380 x 545 x 544 মিমি",
		"40 kg":                    "40 kg",
		"Black":                    "ব্ল্যাক",
		"Inverter":                 "ইনভার্টার",
		"Direct Cool":              "ডাইরেক্ট কুল",
		"Manual":                   "ম্যানুয়াল",
		"Mechanical":               "মেকানিক্যাল",
		"Tempered Glass":           "টেম্পার্ড গ্লাস",
		"3":                        "৩",
		"2":                        "২",
		"1":                        "১",
		"Yes":                      "হ্যাঁ",
		"No":                       "না",
		"40 dB":                    "৪০ ডেসিবেল",
		"165-254V":                 "১৬৫-২৫৪ভোল্ট",
		"50":                       "৫০",
		"5 Years Service":          "5 বছর সার্ভিস",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসার & 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"157 Ltr Capacity, Top Mounted Refrigerator, 45:55 Space Compartment Ratio, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Anti-Bacterial Gas Kit, Built-in Stabilizer, Odor Filter, Bottle Holder": "157 লিটার ক্যাপাসিটি, টপ মাউন্টেড রেফ্রিজারেটর, 45:55 স্পেস কম্পার্টমেন্ট রেশিও, লো পাওয়ার কনসাম্পশন, রিয়েল টেম্পার্ড গ্লাস ফিনিশ, টেম্পার্ড গ্লাস শেল্ফ, অ্যান্টি-ব্যাকটেরিয়াল গ্যাস কিট, বিল্ট-ইন স্ট্যাবিলাইজার, ওডর ফিল্টার, বটল হোল্ডার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"157 Ltr.":         "157 লিটার",
		"156.5 Ltr.":       "156.5 লিটার",
		"R600":             "আর-৬০০",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds155v-bg'
func (s *SpecificationSeederRefrigeratorSingerFTDS155VBG) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds155v-bg"
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
		"Model Name":                  "FTDS155-BG",
		"Door Type":                   "Top Mounted Refrigerator",
		"Capacity":                    "156.5 Liters",
		"Refrigerator Capacity":       "90.5 Liters",
		"Freezer Capacity":            "66 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1380 x 545 x 544 mm",
		"Weight":                      "40 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "165-254V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600",
		"Gross Volume":                "157 Ltr.",
		"Net Volume":                  "156.5 Ltr.",
		"Special Features":            "157 Ltr Capacity, Top Mounted Refrigerator, 45:55 Space Compartment Ratio, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Anti-Bacterial Gas Kit, Built-in Stabilizer, Odor Filter, Bottle Holder",
	}

	banglaTranslations := s.getBanglaTranslations()
for key, val := range specs {
    if len(val) > 500 {
        specs[key] = val[:500]
    }
}
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
			// Update the value if different
			if existing.Value != value {
				existing.Value = value
				if err := db.Save(&existing).Error; err != nil {
					return err
				}
			}
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
				} else {
					// Update translation if different
					if existingTranslation.Value != banglaValue {
						existingTranslation.Value = banglaValue
						if err := db.Save(&existingTranslation).Error; err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
