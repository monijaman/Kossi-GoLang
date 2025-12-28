package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecoomega5glg252rpbbhp seeds specifications/options for product 'eco-omega5-gl-g252rpbb-hp'
type SpecificationSeederRefrigeratorecoomega5glg252rpbbhp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecoomega5glg252rpbbhp creates a new seeder instance
func NewSpecificationSeederRefrigeratorecoomega5glg252rpbbhp() *SpecificationSeederRefrigeratorecoomega5glg252rpbbhp {
	return &SpecificationSeederRefrigeratorecoomega5glg252rpbbhp{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-omega5-gl-g252rpbb-hp"},
	}
}

func (s *SpecificationSeederRefrigeratorecoomega5glg252rpbbhp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                  "ইকো+",
		"Omega5 Gl G252rpbb Hp": "Omega5 Gl G252rpbb Hp",
		"Glass Door":            "গ্লাস ডোর",
		"252 Liters":            "252 লিটার",
		"150 Liters":            "150 লিটার",
		"102 Liters":            "102 লিটার",
		"Red Pearl Black":       "রেড পার্ল ব্ল্যাক",
		"R600a":                 "R600a",
		"Mechanical":            "মেকানিক্যাল",
		"Frost":                 "ফ্রস্ট",
		"No":                    "না",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "10 বছর কম্প্রেসার ওয়ারেন্টি, 2 বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"10":             "10",
		"Frameless":      "ফ্রেমলেস",
		"Bulb (side)":    "বাল্ব (সাইড)",
		"2":              "2",
		"1":              "1",
		"Tempered glass": "টেম্পার্ড গ্লাস",
		"220~240 / 50Hz": "220~240 / 50Hz",
		"50":             "50",

		"Frameless Door Design, R600a Refrigerant, Interior Light, Antibacterial Gasket": "ফ্রেমলেস ডোর ডিজাইন, R600a রেফ্রিজারেন্ট, ইন্টেরিয়র লাইট, অ্যান্টিব্যাকটেরিয়াল গ্যাসকেট",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-omega5-gl-g252rpbb-hp'
func (s *SpecificationSeederRefrigeratorecoomega5glg252rpbbhp) Seed(db *gorm.DB) error {
	productSlug := "eco-omega5-gl-g252rpbb-hp"
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
		"Model Name":                  "Omega5 Gl G252rpbb Hp",
		"Door Type":                   "Glass Door",
		"Capacity":                    "252 Liters",
		"Refrigerator Capacity":       "150 Liters",
		"Freezer Capacity":            "102 Liters",
		"Color":                       "Red Pearl Black",
		"Compressor Type":             "R600a",
		"Defrost Type":                "Frost",
		"Temperature Control":         "Mechanical",
		"Ice Maker":                   "1",
		"Water Dispenser":             "No",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Net Volume":                  "252 Liters",
		"Gross Volume":                "252 Liters",
		"Display":                     "Mechanical",
		"LED Light":                   "Bulb (side)",
		"Egg Tray":                    "2",
		"Shelf Material":              "Tempered glass",
		"Number of Shelves":           "2",
		"Voltage":                     "220~240 / 50Hz",
		"Frequency (Hz)":              "50",
		"Refrigerant":                 "R600a",
		"Special Features":            "Frameless Door Design, R600a Refrigerant, Interior Light, Antibacterial Gasket",
		// Add specifications here
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
