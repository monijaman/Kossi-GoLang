package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecossrt29har9dutd3 seeds specifications/options for product 'eco-ss-rt29har9dut-d3'
type SpecificationSeederRefrigeratorecossrt29har9dutd3 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecossrt29har9dutd3 creates a new seeder instance
func NewSpecificationSeederRefrigeratorecossrt29har9dutd3() *SpecificationSeederRefrigeratorecossrt29har9dutd3 {
	return &SpecificationSeederRefrigeratorecossrt29har9dutd3{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-ss-rt29har9dut-d3"},
	}
}

func (s *SpecificationSeederRefrigeratorecossrt29har9dutd3) getBanglaTranslations() map[string]string {
	return map[string]string{
		"SAMSUNG":                  "স্যামসাং",
		"SS RT29HAR9DUT/D3 Purple": "এসএস আরটি২৯এইচএআর৯ডিইউটি/ডি৩ পার্পল",
		"VCM":                      "ভিসিএম",
		"255":                      "২৫৫",
		"202":                      "২০২",
		"53":                       "৫৩",
		"275":                      "২৭৫",
		"Frost":                    "ফ্রস্ট",
		"R600a":                    "আর৬০০এ",
		"Purple":                   "পার্পল",
		"1685*555*637 mm":          "১৬৮৫*৫৫৫*৬৩৭ মিমি",
		"20 Years Compressor Warranty, 1 Year Parts and Service Warranty":     "২০ বছর কম্প্রেসর ওয়ারেন্টি, ১ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Humidity Control, Interior LED Light, Multi Storage Box, Multi Flow": "হিউমিডিটি কন্ট্রোল, ইন্টেরিয়র এলইডি লাইট, মাল্টি স্টোরেজ বক্স, মাল্টি ফ্লো",
		"Multi Flow": "মাল্টি ফ্লো",
		"Mechanical": "মেকানিক্যাল",
		"3":          "৩",
		"5":          "৫",
		"1":          "১",
		"Yes":        "হ্যাঁ",
		"50.7":       "৫০.৭",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-ss-rt29har9dut-d3'
func (s *SpecificationSeederRefrigeratorecossrt29har9dutd3) Seed(db *gorm.DB) error {
	productSlug := "eco-ss-rt29har9dut-d3"
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
		"Brand":                 "SAMSUNG",
		"Model Name":            "SS RT29HAR9DUT/D3 Purple",
		"Door Type":             "VCM",
		"Capacity":              "255",
		"Refrigerator Capacity": "202",
		"Freezer Capacity":      "53",
		"Gross Volume":          "275",
		"Net Volume":            "255",
		"Defrost Type":          "Frost",
		"Refrigerant":           "R600a",
		"Color":                 "Purple",
		"Dimensions":            "1685*555*637 mm",
		"Weight":                "50.7",
		"Warranty":              "20 Years Compressor Warranty, 1 Year Parts and Service Warranty",
		"Special Features":      "Humidity Control, Interior LED Light, Multi Storage Box, Multi Flow",
		"Cooling Technology":    "Multi Flow",
		"Temperature Control":   "Mechanical",
		"Number of Shelves":     "3",
		"Door Bins":             "5",
		"Crisper Drawers":       "1",
		"Ice Maker":             "Yes",
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
