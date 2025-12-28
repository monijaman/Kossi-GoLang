package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecossrs72r5001m9d2silver seeds specifications/options for product 'eco-ss-rs72r5001m9-d2-silver'
type SpecificationSeederRefrigeratorecossrs72r5001m9d2silver struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecossrs72r5001m9d2silver creates a new seeder instance
func NewSpecificationSeederRefrigeratorecossrs72r5001m9d2silver() *SpecificationSeederRefrigeratorecossrs72r5001m9d2silver {
	return &SpecificationSeederRefrigeratorecossrs72r5001m9d2silver{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-ss-rs72r5001m9-d2-silver"},
	}
}

func (s *SpecificationSeederRefrigeratorecossrs72r5001m9d2silver) getBanglaTranslations() map[string]string {
	return map[string]string{
		"SAMSUNG":                     "স্যামসাং",
		"SS RS72R5001M9/D2 Silver":    "এসএস আরএস৭২আর৫০০১এম৯/ডি২ সিলভার",
		"VCM":                         "ভিসিএম",
		"647":                         "৬৪৭",
		"418":                         "৪১৮",
		"229":                         "২২৯",
		"700":                         "৭০০",
		"No-frost":                    "নো-ফ্রস্ট",
		"Digital Inverter Compressor": "ডিজিটাল ইনভার্টার কম্প্রেসর",
		"R600a":                       "আর৬০০এ",
		"Silver":                      "সিলভার",
		"1780*912*716":                "১৭৮০*৯১২*৭১৬",
		"20 Years Compressor Warranty, 1 Year Parts and Service Warranty":                                      "২০ বছর কম্প্রেসর ওয়ারেন্টি, ১ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Digital Inverter Compressor, Power Cool Function, Power Freeze Function, Mono Cooling, Vacation Mode": "ডিজিটাল ইনভার্টার কম্প্রেসর, পাওয়ার কুল ফাংশন, পাওয়ার ফ্রিজ ফাংশন, মনো কুলিং, ভ্যাকেশন মোড",
		"Mono Cooling":        "মনো কুলিং",
		"Internal (Ice Blue)": "ইন্টারনাল (আইস ব্লু)",
		"4":                   "৪",
		"5":                   "৫",
		"2":                   "২",
		"Manual Twist":        "ম্যানুয়াল টুইস্ট",
		"Yes":                 "হ্যাঁ",
		"101":                 "১০১",
		"Side-by-Side":        "সাইড-বাই-সাইড",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-ss-rs72r5001m9-d2-silver'
func (s *SpecificationSeederRefrigeratorecossrs72r5001m9d2silver) Seed(db *gorm.DB) error {
	productSlug := "eco-ss-rs72r5001m9-d2-silver"
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
		"Model Name":            "SS RS72R5001M9/D2 Silver",
		"Door Type":             "VCM",
		"Capacity":              "647",
		"Refrigerator Capacity": "418",
		"Freezer Capacity":      "229",
		"Gross Volume":          "700",
		"Net Volume":            "647",
		"Defrost Type":          "No-frost",
		"Compressor Type":       "Digital Inverter Compressor",
		"Refrigerant":           "R600a",
		"Color":                 "Silver",
		"Dimensions":            "1780*912*716",
		"Weight":                "101",
		"Warranty":              "20 Years Compressor Warranty, 1 Year Parts and Service Warranty",
		"Special Features":      "Digital Inverter Compressor, Power Cool Function, Power Freeze Function, Mono Cooling, Vacation Mode",
		"Cooling Technology":    "Mono Cooling",
		"Temperature Control":   "Internal (Ice Blue)",
		"Number of Shelves":     "4",
		"Door Bins":             "5",
		"Crisper Drawers":       "2",
		"Ice Maker":             "Manual Twist",
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
