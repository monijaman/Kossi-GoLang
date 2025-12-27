package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS260V seeds specifications/options for product 'singer-ftds260v'
type SpecificationSeederRefrigeratorSingerFTDS260V struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS260V creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS260V() *SpecificationSeederRefrigeratorSingerFTDS260V {
	return &SpecificationSeederRefrigeratorSingerFTDS260V{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds260v"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS260V) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                          "সিঙ্গার",
		"FBDS260V":                        "FBDS260V",
		"Bottom Mounted Refrigerator":     "বটম মাউন্টেড রেফ্রিজারেটর",
		"260 Liters":                      "260 Liters",
		"5 Star":                          "৫ তারা",
		"5":                               "৫",
		"250 kWh":                         "২৫০ কিলোওয়াট ঘণ্টা",
		"1045 x 605 x 844 mm":             "১০৪৫ x ৬০৫ x ৮৪৪ মিমি",
		"208 kg":                          "208 kg",
		"Silver":                          "সিলভার",
		"Inverter":                        "ইনভার্টার",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Manual":                          "ম্যানুয়াল",
		"Mechanical":                      "মেকানিক্যাল",
		"Wire Shelves":                    "ওয়্যার শেল্ফ",
		"2":                               "২",
		"3":                               "৩",
		"1":                               "১",
		"Yes":                             "হ্যাঁ",
		"No":                              "না",
		"40 dB":                           "৪০ ডেসিবেল",
		"220 ~ 240V":                      "২২০ ~ ২৪০ভোল্ট",
		"50":                              "৫০",
		"5 Years Service":                 "5 বছর সার্ভিস",
		"Chest Freezer, Energy Efficient": "চেস্ট ফ্রিজার, এনার্জি এফিশিয়েন্ট",
		"Refrigerant":                     "রেফ্রিজারেন্ট",
		"Gross Volume":                    "মোট ভলিউম",
		"Net Volume":                      "নেট ভলিউম",
		"260 Ltr.":                        "260 লিটার",
		"R-134a":                          "আর-১৩৪এ",
		"Special Features":                "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds260v'
func (s *SpecificationSeederRefrigeratorSingerFTDS260V) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds260v"
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
		"Model Name":                  "FBDS260V",
		"Door Type":                   "Bottom Mounted Refrigerator",
		"Capacity":                    "260 Liters",
		"Refrigerator Capacity":       "130 Liters",
		"Freezer Capacity":            "130 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1650 x 560 x 500 mm",
		"Weight":                      "55 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass Shelves",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "135V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "260 Ltr.",
		"Net Volume":                  "260 Ltr.",
		"Special Features":            "Bottom Mounted Refrigerator, 50:50 Space Compartment Ratio, 5 Star Energy Rating by BSTI, Low power consumption, NutiLock & Freshology Technology, VCM Door, Tempered Glass Shelves, Built-in Stabilizer with 135V Running",
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
