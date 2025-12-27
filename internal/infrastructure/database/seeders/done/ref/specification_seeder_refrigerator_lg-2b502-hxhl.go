package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorLg2b502Hxhl seeds specifications/options for product 'lg-2b502-hxhl'
type SpecificationSeederRefrigeratorLg2b502Hxhl struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorLg2b502Hxhl creates a new seeder instance
func NewSpecificationSeederRefrigeratorLg2b502Hxhl() *SpecificationSeederRefrigeratorLg2b502Hxhl {
	return &SpecificationSeederRefrigeratorLg2b502Hxhl{
		BaseSeeder: BaseSeeder{name: "Specifications for lg-2b502-hxhl"},
	}
}

func (s *SpecificationSeederRefrigeratorLg2b502Hxhl) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                        "এলজি",
		"2B502HXHL":                 "২বি৫০২এইচএক্সএইচএল",
		"Side by Side":              "সাইড বাই সাইড",
		"502 Liters":                "৫০২ লিটার",
		"320 Liters":                "৩২০ লিটার",
		"182 Liters":                "১৮২ লিটার",
		"5 Star":                    "৫ তারা",
		"5":                         "৫",
		"250 kWh":                   "২৫০ কিলোওয়াট ঘণ্টা",
		"700 x 900 x 1780 mm":       "৭০০ x ৯০০ x ১৭৮০ মিমি",
		"85 kg":                     "৮৫ কেজি",
		"Stainless Steel":           "স্টেইনলেস স্টিল",
		"Inverter":                  "ইনভার্টার",
		"Frost Free":                "ফ্রস্ট ফ্রি",
		"Electronic":                "ইলেকট্রনিক",
		"Toughened Glass":           "টাফেন্ড গ্লাস",
		"4":                         "৪",
		"2":                         "২",
		"Yes":                       "হ্যাঁ",
		"No":                        "না",
		"42 dB":                     "৪২ ডেসিবেল",
		"220V":                      "২২০ভোল্ট",
		"50":                        "৫০",
		"3 Years":                   "৩ বছর",
		"10":                        "১০",
		"R-600a":                    "আর-৬০০এ",
		"Inverter Technology":       "ইনভার্টার প্রযুক্তি",
		"Refrigerant":               "রেফ্রিজারেন্ট",
		"Gross Volume":              "মোট ভলিউম",
		"Net Volume":                "নেট ভলিউম",
		"Linear Compressor":         "লিনিয়ার কম্প্রেসর",
		"Smart Inverter Compressor": "স্মার্ট ইনভার্টার কম্প্রেসর",
		"220 ~ 240V":                "২২০ ~ ২৪০ভোল্ট",
		"Special Features":          "বিশেষ বৈশিষ্ট্য",
		"Multi Airflow":             "মাল্টি এয়ারফ্লো",
		"Ice & Water Dispenser":     "আইস এবং ওয়াটার ডিসপেন্সার",
		"LED Lighting":              "এলইডি লাইটিং",
	}
}

// Seed inserts specification records for the product identified by slug 'lg-2b502-hxhl'
func (s *SpecificationSeederRefrigeratorLg2b502Hxhl) Seed(db *gorm.DB) error {
	productSlug := "lg-2b502-hxhl"
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
		"Model Name":                  "2B502HXHL",
		"Door Type":                   "Side by Side",
		"Capacity":                    "502 Liters",
		"Refrigerator Capacity":       "320 Liters",
		"Freezer Capacity":            "182 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "700 x 900 x 1780 mm",
		"Weight":                      "85 kg",
		"Color":                       "Stainless Steel",
		"Compressor Type":             "Smart Inverter Compressor",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "6",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "Yes",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "502 Ltr.",
		"Net Volume":                  "502 Ltr.",
		"Special Features":            "Multi Airflow, Ice & Water Dispenser, LED Lighting",
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
