package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorBekoBorefRdne295dse seeds specifications/options for product 'boref-rdne295dse'
type SpecificationSeederRefrigeratorBekoBorefRdne295dse struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorBekoBorefRdne295dse creates a new seeder instance
func NewSpecificationSeederRefrigeratorBekoBorefRdne295dse() *SpecificationSeederRefrigeratorBekoBorefRdne295dse {
	return &SpecificationSeederRefrigeratorBekoBorefRdne295dse{
		BaseSeeder: BaseSeeder{name: "Specifications for boref-rdne295dse"},
	}
}

func (s *SpecificationSeederRefrigeratorBekoBorefRdne295dse) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Beko":                "বেকো",
		"BOREF-RDNE295DSE":    "বোরেফ-আরডিএনই২৯৫ডিএসই",
		"Single Door":         "একটি দরজা",
		"295 Liters":          "২৯৫ লিটার",
		"0 Liters":            "০ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"320 kWh":             "৩২০ কিলোওয়াট ঘণ্টা",
		"545 x 570 x 1520 mm": "৫৪৫ x ৫৭০ x ১৫২০ মিমি",
		"48 kg":               "৪৮ কেজি",
		"Silver":              "সিলভার",
		"Standard":            "স্ট্যান্ডার্ড",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Wire":                "ওয়্যার",
		"2":                   "২",
		"3":                   "৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"Large Vegetable Crisper, Egg Tray, Ice Cube Tray": "বড় শাকসবজি ক্রিস্পার, ডিমের ট্রে, বরফের কিউব ট্রে",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"R-600a":           "আর-৬০০এ",
		"295 Ltr.":         "২৯৫ লিটার",
		"220 ~ 240V":       "২২০ ~ ২৪০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'boref-rdne295dse'
func (s *SpecificationSeederRefrigeratorBekoBorefRdne295dse) Seed(db *gorm.DB) error {
	productSlug := "boref-rdne295dse"
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
		"Model Name":                  "BOREF-RDNE295DSE",
		"Door Type":                   "Single Door",
		"Capacity":                    "295 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "295 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "320 kWh",
		"Dimensions":                  "545 x 570 x 1520 mm",
		"Weight":                      "48 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Standard",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "2",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "295 Ltr.",
		"Net Volume":                  "295 Ltr.",
		"Special Features":            "Large Vegetable Crisper, Egg Tray, Ice Cube Tray",
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
