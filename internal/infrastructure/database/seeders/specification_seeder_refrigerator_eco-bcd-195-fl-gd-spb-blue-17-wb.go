package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecobcd195flgdspbblue17wb seeds specifications/options for product 'eco-bcd-195-fl-gd-spb-blue-17-wb'
type SpecificationSeederRefrigeratorecobcd195flgdspbblue17wb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecobcd195flgdspbblue17wb creates a new seeder instance
func NewSpecificationSeederRefrigeratorecobcd195flgdspbblue17wb() *SpecificationSeederRefrigeratorecobcd195flgdspbblue17wb {
	return &SpecificationSeederRefrigeratorecobcd195flgdspbblue17wb{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-bcd-195-fl-gd-spb-blue-17-wb"},
	}
}

func (s *SpecificationSeederRefrigeratorecobcd195flgdspbblue17wb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                         "ইকো+",
		"BCD-195 FL GD SPB Blue-17 WB": "BCD-195 FL GD SPB Blue-17 WB",
		"Glass Door":                   "গ্লাস ডোর",
		"195 Liters":                   "195 লিটার",
		"122 Liters":                   "122 লিটার",
		"73 Liters":                    "73 লিটার",
		"1393 x 550 x 490 mm":          "1393 x 550 x 490 মিমি",
		"Blue":                         "ব্লু",
		"R600a":                        "R600a",
		"Mechanical":                   "মেকানিক্যাল",
		"Frost":                        "ফ্রস্ট",
		"Manual":                       "ম্যানুয়াল",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "10 বছর কম্প্রেসার ওয়ারেন্টি, 2 বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"10":          "10",
		"Frameless":   "ফ্রেমলেস",
		"Bulb (side)": "বাল্ব (সাইড)",
		"2":           "2",
		"1":           "1",
		"2 / 12":      "2 / 12",
		"Frameless Door Design, R600a Refrigerant, Interior Light, Antibacterial Gasket": "ফ্রেমলেস ডোর ডিজাইন, R600a রেফ্রিজারেন্ট, ইন্টেরিয়র লাইট, অ্যান্টিব্যাকটেরিয়াল গ্যাসকেট",
	}
}

// Seed inserts specification records for the product identified by slug 'eco-bcd-195-fl-gd-spb-blue-17-wb'
func (s *SpecificationSeederRefrigeratorecobcd195flgdspbblue17wb) Seed(db *gorm.DB) error {
	productSlug := "eco-bcd-195-fl-gd-spb-blue-17-wb"
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
		"Model Name":                  "BCD-195 FL GD SPB Blue-17 WB",
		"Door Type":                   "Glass Door",
		"Capacity":                    "195 Liters",
		"Refrigerator Capacity":       "122 Liters",
		"Freezer Capacity":            "73 Liters",
		"Dimensions":                  "1393 x 550 x 490 mm",
		"Color":                       "Blue",
		"Compressor Type":             "R600a",
		"Defrost Type":                "Frost",
		"Temperature Control":         "Mechanical",
		"Ice Maker":                   "1",
		"Water Dispenser":             "No",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Net Volume":                  "195 Liters",
		"Gross Volume":                "195 Liters",
		"Display":                     "Mechanical",
		"LED Light":                   "Bulb (side)",
		"Egg Tray":                    "2 / 12",
		"Shelf Material":              "Tempered glass",
		"Number of Shelves":           "2",
		"Voltage":                     "220~240 / 50Hz",
		"Frequency (Hz)":              "50",
		"Refrigerant":                 "R600a",
		"Special Features":            "Frameless Door Design, R600a Refrigerant, Interior Light, Antibacterial Gasket",
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
