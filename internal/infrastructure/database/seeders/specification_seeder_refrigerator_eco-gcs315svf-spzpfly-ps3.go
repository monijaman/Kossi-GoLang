package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecogcs315svfspzpflyps3 seeds specifications/options for product 'eco-gcs315svf-spzpfly-ps3'
type SpecificationSeederRefrigeratorecogcs315svfspzpflyps3 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecogcs315svfspzpflyps3 creates a new seeder instance
func NewSpecificationSeederRefrigeratorecogcs315svfspzpflyps3() *SpecificationSeederRefrigeratorecogcs315svfspzpflyps3 {
	return &SpecificationSeederRefrigeratorecogcs315svfspzpflyps3{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-gcs315svf-spzpfly-ps3"},
	}
}

func (s *SpecificationSeederRefrigeratorecogcs315svfspzpflyps3) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                       "LG",
		"LG GCS315SVF.SPZPFLY PS3": "LG GCS315SVF.SPZPFLY PS3",
		"Top Opening":              "টপ ওপেনিং",
		"280":                      "২৮০",
		"190":                      "১৯০",
		"845*954*560 mm":           "৮৪৫*৯৫৪*৫৬০ মিমি",
		"Freezer Silver":           "ফ্রিজার সিলভার",
		"Recipro":                  "রেসিপ্রো",
		"Frost":                    "ফ্রস্ট",
		"E-Control":                "ই-কন্ট্রোল",
		"1":                        "১",
		"Mechanical temperature control, Water disposal device, Key and Lock, Door LED": "মেকানিক্যাল টেম্পারেচার কন্ট্রোল, ওয়াটার ডিসপোজাল ডিভাইস, কী অ্যান্ড লক, ডোর LED",
		"Yes":     "হ্যাঁ",
		"No":      "না",
		"220~240": "২২০~২৪০",
		"50":      "৫০",
		"134a":    "134a",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-gcs315svf-spzpfly-ps3'
func (s *SpecificationSeederRefrigeratorecogcs315svfspzpflyps3) Seed(db *gorm.DB) error {
	productSlug := "eco-gcs315svf-spzpfly-ps3"
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
		"Brand":                       "LG",
		"Model Name":                  "LG GCS315SVF.SPZPFLY PS3",
		"Door Type":                   "Top Opening",
		"Capacity":                    "280",
		"Freezer Capacity":            "190",
		"Gross Volume":                "280",
		"Net Volume":                  "190",
		"Dimensions":                  "845*954*560 mm",
		"Weight":                      "40",
		"Color":                       "Freezer Silver",
		"Compressor Type":             "Recipro",
		"Defrost Type":                "Frost",
		"Temperature Control":         "E-Control",
		"Voltage":                     "220~240",
		"Frequency (Hz)":              "50",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "134a",
		"Special Features":            "Mechanical temperature control, Water disposal device, Key and Lock, Door LED",
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
