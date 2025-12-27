package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS200V seeds specifications/options for product 'singer-ftds200v'
type SpecificationSeederRefrigeratorSingerFTDS200V struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS200V creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS200V() *SpecificationSeederRefrigeratorSingerFTDS200V {
	return &SpecificationSeederRefrigeratorSingerFTDS200V{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds200v"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS200V) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                       "সিঙ্গার",
		"FTDS200V":                     "FTDS200V",
		"Top Mounted Refrigerator":     "টপ মাউন্টেড রেফ্রিজারেটর",
		"200 Liters":                   "200 লিটার",
		"116.8 Liters":                 "116.8 লিটার",
		"83.2 Liters":                  "83.2 লিটার",
		"5 Star":                       "৫ তারা",
		"5":                            "৫",
		"250 kWh":                      "২৫০ কিলোওয়াট ঘণ্টা",
		"1630 x 545 x 550 mm":          "1630 x 545 x 550 মিমি",
		"50 kg":                        "50 kg",
		"Blue":                         "ব্লু",
		"Inverter":                     "ইনভার্টার",
		"Direct Cool":                  "ডাইরেক্ট কুল",
		"Manual":                       "ম্যানুয়াল",
		"Mechanical":                   "মেকানিক্যাল",
		"Tempered Glass Shelves":       "টেম্পার্ড গ্লাস শেল্ফ",
		"3":                            "৩",
		"1":                            "১",
		"No":                           "না",
		"40 dB":                        "৪০ ডেসিবেল",
		"135V":                         "135ভোল্ট",
		"50":                           "৫০",
		"10 Years Compressor Warranty": "10 বছর কম্প্রেসর ওয়ারেন্টি",
		"10":                           "10",
		"R600a":                        "আর600এ",
		"202.9 Ltr.":                   "202.9 লিটার",
		"200 Ltr.":                     "200 লিটার",
		"Top Mounted Refrigerator, Low Voltage Compressor, Odor Filter, Bottle Holder, Anti-Bacterial Gasket, Tempered Glass Shelves, Base Drawer": "টপ মাউন্টেড রেফ্রিজারেটর, লো ভোল্টেজ কম্প্রেসর, ওডর ফিল্টার, বটল হোল্ডার, অ্যান্টি-ব্যাকটেরিয়াল গ্যাসকেট, টেম্পার্ড গ্লাস শেল্ফ, বেস ড্রয়ার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds200v'
func (s *SpecificationSeederRefrigeratorSingerFTDS200V) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds200v"
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
		"Model Name":                  "FTDS200V",
		"Door Type":                   "Top Mounted Refrigerator",
		"Capacity":                    "200 Liters",
		"Refrigerator Capacity":       "116.8 Liters",
		"Freezer Capacity":            "83.2 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1630 x 545 x 550 mm",
		"Weight":                      "50 kg",
		"Color":                       "Blue",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass Shelves",
		"Number of Shelves":           "5",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "135V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "202.9 Ltr.",
		"Net Volume":                  "200 Ltr.",
		"Special Features":            "Top Mounted Refrigerator, Low Voltage Compressor, Odor Filter, Bottle Holder, Anti-Bacterial Gasket, Tempered Glass Shelves, Base Drawer",
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
