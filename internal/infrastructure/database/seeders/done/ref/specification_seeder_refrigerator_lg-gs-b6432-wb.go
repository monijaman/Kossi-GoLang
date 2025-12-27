package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorLgGsB6432Wb seeds specifications/options for product 'lg-gs-b6432-wb'
type SpecificationSeederRefrigeratorLgGsB6432Wb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorLgGsB6432Wb creates a new seeder instance
func NewSpecificationSeederRefrigeratorLgGsB6432Wb() *SpecificationSeederRefrigeratorLgGsB6432Wb {
	return &SpecificationSeederRefrigeratorLgGsB6432Wb{
		BaseSeeder: BaseSeeder{name: "Specifications for lg-gs-b6432-wb"},
	}
}

func (s *SpecificationSeederRefrigeratorLgGsB6432Wb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                        "এলজি",
		"GS-B6432WB":                "জিএস-বি৬৪৩২ডব্লিউবি",
		"Bottom Mount Freezer":      "বটম মাউন্ট ফ্রিজার",
		"643 Liters":                "৬৪৩ লিটার",
		"430 Liters":                "৪৩০ লিটার",
		"213 Liters":                "২১৩ লিটার",
		"5 Star":                    "৫ তারা",
		"5":                         "৫",
		"290 kWh":                   "২৯০ কিলোওয়াট ঘণ্টা",
		"700 x 800 x 1780 mm":       "৭০০ x ৮০০ x ১৭৮০ মিমি",
		"88 kg":                     "৮৮ কেজি",
		"White":                     "সাদা",
		"Inverter":                  "ইনভার্টার",
		"Frost Free":                "ফ্রস্ট ফ্রি",
		"Electronic":                "ইলেকট্রনিক",
		"Toughened Glass":           "টাফেন্ড গ্লাস",
		"4":                         "৪",
		"2":                         "২",
		"Yes":                       "হ্যাঁ",
		"No":                        "না",
		"40 dB":                     "৪০ ডেসিবেল",
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
		"LED Lighting":              "এলইডি লাইটিং",
		"Door Cooling":              "দরজা কুলিং",
		"Smart Diagnosis":           "স্মার্ট ডায়াগনোসিস",
	}
}

// Seed inserts specification records for the product identified by slug 'lg-gs-b6432-wb'
func (s *SpecificationSeederRefrigeratorLgGsB6432Wb) Seed(db *gorm.DB) error {
	productSlug := "lg-gs-b6432-wb"
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
		"Model Name":                  "GS-B6432WB",
		"Door Type":                   "Bottom Mount Freezer",
		"Capacity":                    "643 Liters",
		"Refrigerator Capacity":       "430 Liters",
		"Freezer Capacity":            "213 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "290 kWh",
		"Dimensions":                  "700 x 800 x 1780 mm",
		"Weight":                      "88 kg",
		"Color":                       "White",
		"Compressor Type":             "Smart Inverter Compressor",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "6",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "643 Ltr.",
		"Net Volume":                  "643 Ltr.",
		"Special Features":            "Multi Airflow, LED Lighting, Door Cooling, Smart Diagnosis",
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
