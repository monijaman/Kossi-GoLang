package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSharpSjEsb631Xn seeds specifications/options for product 'sj‑esb631xn'
type SpecificationSeederRefrigeratorSharpSjEsb631Xn struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSharpSjEsb631Xn creates a new seeder instance
func NewSpecificationSeederRefrigeratorSharpSjEsb631Xn() *SpecificationSeederRefrigeratorSharpSjEsb631Xn {
	return &SpecificationSeederRefrigeratorSharpSjEsb631Xn{
		BaseSeeder: BaseSeeder{name: "Specifications for sj‑esb631xn"},
	}
}

func (s *SpecificationSeederRefrigeratorSharpSjEsb631Xn) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Sharp":               "শার্প",
		"SJ‑ESB631XN":         "এসজে-ইএসবি৬৩১এক্সএন",
		"Single Door":         "সিঙ্গেল দরজা",
		"631 Liters":          "৬৩১ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"380 kWh":             "৩৮০ কিলোওয়াট ঘণ্টা",
		"650 x 650 x 1850 mm": "৬৫০ x ৬৫০ x ১৮৫০ মিমি",
		"85 kg":               "৮৫ কেজি",
		"Silver":              "সিলভার",
		"Inverter":            "ইনভার্টার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Mechanical":          "মেকানিক্যাল",
		"Wire Shelves":        "ওয়াইয়ার শেল্ফ",
		"4":                   "৪",
		"2":                   "২",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, LED Lighting, Large Vegetable Crisper": "ইনভার্টার কম্প্রেসার, এলইডি লাইটিং, লার্জ ভেজিটেবল ক্রিস্পার",
		"Refrigerant":                  "রেফ্রিজারেন্ট",
		"Gross Volume":                 "মোট ভলিউম",
		"Net Volume":                   "নেট ভলিউম",
		"631 Ltr.":                     "৬৩১ লিটার",
		"0 Ltr.":                       "০ লিটার",
		"Advanced Inverter Technology": "অ্যাডভান্সড ইনভার্টার টেকনোলজি",
		"220 ~ 240V":                   "২২০ ~ ২৪০ভোল্ট",
		"Special Features":             "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'sj‑esb631xn'
func (s *SpecificationSeederRefrigeratorSharpSjEsb631Xn) Seed(db *gorm.DB) error {
	productSlug := "sj‑esb631xn"
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
		"Brand":                       "Sharp",
		"Model Name":                  "SJ‑ESB631XN",
		"Door Type":                   "Single Door",
		"Capacity":                    "631 Liters",
		"Refrigerator Capacity":       "631 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "380 kWh",
		"Dimensions":                  "650 x 650 x 1850 mm",
		"Weight":                      "85 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire Shelves",
		"Number of Shelves":           "4",
		"Door Bins":                   "2",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "631 Ltr.",
		"Net Volume":                  "631 Ltr.",
		"Special Features":            "Inverter Compressor, LED Lighting, Large Vegetable Crisper",
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
				if err := db.Create(sModel).Error; err != nil {
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
