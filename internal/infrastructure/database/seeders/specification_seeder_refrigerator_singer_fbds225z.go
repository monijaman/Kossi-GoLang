package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFBDS225Z seeds specifications/options for product 'singer-fbds225z'
type SpecificationSeederRefrigeratorSingerFBDS225Z struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFBDS225Z creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFBDS225Z() *SpecificationSeederRefrigeratorSingerFBDS225Z {
	return &SpecificationSeederRefrigeratorSingerFBDS225Z{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-fbds225z"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFBDS225Z) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                      "সিঙ্গার",
		"FBDS225Z":                    "FBDS225Z",
		"Bottom Mounted Refrigerator": "বটম মাউন্টেড রেফ্রিজারেটর",
		"215 Liters":                  "215 লিটার",
		"120 Liters":                  "120 লিটার",
		"95 Liters":                   "95 লিটার",
		"5 Star":                      "৫ তারা",
		"5":                           "৫",
		"10":                          "10",
		"250 kWh":                     "২৫০ কিলোওয়াট ঘণ্টা",
		"1045 x 605 x 844 mm":         "১০৪৫ x ৬০৫ x ৮৪৪ মিমি",
		"48 kg":                       "48 kg",
		"Red":                         "রেড",
		"Inverter":                    "ইনভার্টার",
		"Direct Cool":                 "ডাইরেক্ট কুল",
		"Manual":                      "ম্যানুয়াল",
		"Mechanical":                  "মেকানিক্যাল",
		"Tempered Glass":              "টেম্পার্ড গ্লাস",
		"3":                           "৩",
		"2":                           "২",
		"1":                           "১",
		"Yes":                         "হ্যাঁ",
		"No":                          "না",
		"40 dB":                       "৪০ ডেসিবেল",
		"220 ~ 240V":                  "২২০ ~ ২৪০ভোল্ট",
		"50":                          "৫০",
		"2 Years Spare Parts":         "2 বছর স্পেয়ার পার্টস",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসার & 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"215 Ltr Capacity, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Nutilock & Freshology Technology, Built-in Stabilizer, Odour Filter, Bottle Holder, Bottom Mounted Refrigerator": `215 লিটার ক্যাপাসিটি, বিএসটিআই দ্বারা 5 তারা এনার্জি রেটিং, লো পাওয়ার কনসাম্পশন, রিয়েল টেম্পার্ড গ্লাস ফিনিশ, টেম্পার্ড গ্লাস শেল্ফ, নুট্রিলক & ফ্রেশোলজি টেকনোলজি, বিল্ট-ইন স্ট্যাবিলাইজার, ওডর ফিল্টার, বটল হোল্ডার, বটম মাউন্টেড রেফ্রিজারেটর`,
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"215 Ltr.":         "215 লিটার",
		"R600a":            "আর-৬০০এ",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-fbds225z'
func (s *SpecificationSeederRefrigeratorSingerFBDS225Z) Seed(db *gorm.DB) error {
	productSlug := "singer-fbds225z"
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
		"Brand":                       "Singer",
		"Model Name":                  "FBDS225Z",
		"Door Type":                   "Bottom Mounted Refrigerator",
		"Capacity":                    "215 Liters",
		"Refrigerator Capacity":       "120 Liters",
		"Freezer Capacity":            "95 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1045 x 605 x 844 mm",
		"Weight":                      "48 kg",
		"Color":                       "Red",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "215 Ltr.",
		"Net Volume":                  "215 Ltr.",
		"Special Features":            "215 Ltr Capacity, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Nutilock & Freshology Technology, Built-in Stabilizer, Odour Filter, Bottle Holder, Bottom Mounted Refrigerator",
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
