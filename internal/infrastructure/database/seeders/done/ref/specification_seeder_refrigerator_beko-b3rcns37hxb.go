package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorBekoB3rcns37hxb seeds specifications/options for product 'b3rcns37hxb'
type SpecificationSeederRefrigeratorBekoB3rcns37hxb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorBekoB3rcns37hxb creates a new seeder instance
func NewSpecificationSeederRefrigeratorBekoB3rcns37hxb() *SpecificationSeederRefrigeratorBekoB3rcns37hxb {
	return &SpecificationSeederRefrigeratorBekoB3rcns37hxb{
		BaseSeeder: BaseSeeder{name: "Specifications for b3rcns37hxb"},
	}
}

func (s *SpecificationSeederRefrigeratorBekoB3rcns37hxb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Beko":                "বেকো",
		"B3RCNS37HXB":         "বি৩আরসিএনএস৩৭এইচএক্সবি",
		"Double Door":         "দুইটি দরজা",
		"370 Liters":          "৩৭০ লিটার",
		"240 Liters":          "২৪০ লিটার",
		"130 Liters":          "১৩০ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"220 kWh":             "২২০ কিলোওয়াট ঘণ্টা",
		"595 x 655 x 1850 mm": "৫৯৫ x ৬৫৫ x ১৮৫০ মিমি",
		"65 kg":               "৬৫ কেজি",
		"Black":               "কালো",
		"Standard":            "স্ট্যান্ডার্ড",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"3":                   "৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Frost Free, LED Lighting, Humidity Control": "ফ্রস্ট ফ্রি, এলইডি লাইটিং, হিউমিডিটি কন্ট্রোল",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"R-600a":           "আর-৬০০এ",
		"370 Ltr.":         "৩৭০ লিটার",
		"220 ~ 240V":       "২২০ ~ ২৪০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'b3rcns37hxb'
func (s *SpecificationSeederRefrigeratorBekoB3rcns37hxb) Seed(db *gorm.DB) error {
	productSlug := "b3rcns37hxb"
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
		"Model Name":                  "B3RCNS37HXB",
		"Door Type":                   "Double Door",
		"Capacity":                    "370 Liters",
		"Refrigerator Capacity":       "240 Liters",
		"Freezer Capacity":            "130 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "220 kWh",
		"Dimensions":                  "595 x 655 x 1850 mm",
		"Weight":                      "65 kg",
		"Color":                       "Black",
		"Compressor Type":             "Standard",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "370 Ltr.",
		"Net Volume":                  "370 Ltr.",
		"Special Features":            "Frost Free, LED Lighting, Humidity Control",
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
