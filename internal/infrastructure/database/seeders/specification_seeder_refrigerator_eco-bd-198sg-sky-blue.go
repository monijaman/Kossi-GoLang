package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecobd198sgskyblue seeds specifications/options for product 'eco-bd-198sg-sky-blue'
type SpecificationSeederRefrigeratorecobd198sgskyblue struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecobd198sgskyblue creates a new seeder instance
func NewSpecificationSeederRefrigeratorecobd198sgskyblue() *SpecificationSeederRefrigeratorecobd198sgskyblue {
	return &SpecificationSeederRefrigeratorecobd198sgskyblue{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-bd-198sg-sky-blue"},
	}
}

func (s *SpecificationSeederRefrigeratorecobd198sgskyblue) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                   "ইকো+",
		"ECO+ BD-198SG Sky Blue": "ECO+ BD-198SG Sky Blue",
		"Top Opening":            "টপ ওপেনিং",
		"198":                    "১৯৮",
		"189":                    "১৮৯",
		"845*890*560 mm":         "৮৪৫*৮৯০*৫৬০ মিমি",
		"Sky Blue":               "স্কাই ব্লু",
		"R600a":                  "R600a",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical temperature control with adjustable thermostat": "মেকানিক্যাল টেম্পারেচার কন্ট্রোল উইথ অ্যাডজাস্টেবল থার্মোস্ট্যাট",
		"1": "১",
		"4": "৪",
		"R600a Refrigerant, Fast Freezing, Door Lock, Power Indicator, 30° to 60°angle hovering": "R600a রেফ্রিজারেন্ট, ফাস্ট ফ্রিজিং, ডোর লক, পাওয়ার ইন্ডিকেটর, ৩০° টু ৬০° অ্যাঙ্গেল হোভারিং",
		"Yes":     "হ্যাঁ",
		"No":      "না",
		"187~276": "১৮৭~২৭৬",
		"50":      "৫০",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-bd-198sg-sky-blue'
func (s *SpecificationSeederRefrigeratorecobd198sgskyblue) Seed(db *gorm.DB) error {
	productSlug := "eco-bd-198sg-sky-blue"
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
		"Brand":                       "ECO+",
		"Model Name":                  "ECO+ BD-198SG Sky Blue",
		"Door Type":                   "Top Opening",
		"Capacity":                    "198",
		"Freezer Capacity":            "189",
		"Gross Volume":                "198",
		"Net Volume":                  "189",
		"Dimensions":                  "845*890*560 mm",
		"Weight":                      "34",
		"Color":                       "Sky Blue",
		"Compressor Type":             "R600a",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical temperature control with adjustable thermostat",
		"Voltage":                     "187~276",
		"Frequency (Hz)":              "50",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Special Features":            "R600a Refrigerant, Fast Freezing, Door Lock, Power Indicator, 30° to 60°angle hovering",
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
