package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack seeds specifications/options for product 'eco-wp-neo-inv-258gd-crystal-black'
type SpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack creates a new seeder instance
func NewSpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack() *SpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack {
	return &SpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-wp-neo-inv-258gd-crystal-black"},
	}
}

func (s *SpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack) getBanglaTranslations() map[string]string {
	return map[string]string{
		"WHIRLPOOL":                      "হোয়ার্লপুল",
		"WP NEO INV 258GD CRYSTAL BLACK": "ডব্লিউপি নিও আইএনভি ২৫৮জিডি ক্রিস্টাল ব্ল্যাক",
		"Glass Door":                     "গ্লাস ডোর",
		"245":                            "২৪৫",
		"No-frost":                       "নো-ফ্রস্ট",
		"Black":                          "ব্ল্যাক",
		"158.9*56.4*66.7":                "১৫৮.৯*৫৬.৪*৬৬.৭",
		"10 Years Compressor Warranty, 2 Years Parts and Service Warranty": "১০ বছর কম্প্রেসর ওয়ারেন্টি, ২ বছর পার্টস এবং সার্ভিস ওয়ারেন্টি",
		"Up to 30 days preservation for chicken & fish, Up to 14 days garden freshness for fruits & vegetables, Active deo for anti odour action, Chilling gel for cooling retention during power cuts, Microblock technology prevents bacterial growth up to 99%": "চিকেন এবং ফিশের জন্য ৩০ দিন পর্যন্ত সংরক্ষণ, ফল এবং সবজির জন্য ১৪ দিন গার্ডেন ফ্রেশনেস, অ্যান্টি ওডর অ্যাকশনের জন্য অ্যাকটিভ ডিও, পাওয়ার কাটের সময় কুলিং রিটেনশনের জন্য চিলিং গেল, ৯৯% পর্যন্ত ব্যাকটেরিয়াল গ্রোথ প্রতিরোধ করে মাইক্রোব্লক টেকনোলজি",
		"4":   "৪",
		"1":   "১",
		"Yes": "হ্যাঁ",
		"58":  "৫৮",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-wp-neo-inv-258gd-crystal-black'
func (s *SpecificationSeederRefrigeratorecowpneoinv258gdcrystalblack) Seed(db *gorm.DB) error {
	productSlug := "eco-wp-neo-inv-258gd-crystal-black"
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
		"Brand":             "WHIRLPOOL",
		"Model Name":        "WP NEO INV 258GD CRYSTAL BLACK",
		"Door Type":         "Glass Door",
		"Capacity":          "245",
		"Gross Volume":      "245",
		"Defrost Type":      "No-frost",
		"Color":             "Black",
		"Dimensions":        "158.9*56.4*66.7",
		"Weight":            "58",
		"Warranty":          "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Special Features":  "Up to 30 days preservation for chicken & fish, Up to 14 days garden freshness for fruits & vegetables, Active deo for anti odour action, Chilling gel for cooling retention during power cuts, Microblock technology prevents bacterial growth up to 99%",
		"Number of Shelves": "4",
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
