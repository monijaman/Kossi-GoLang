package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSharpSjc118Gy seeds specifications/options for product 'sjc‑118‑gy'
type SpecificationSeederRefrigeratorSharpSjc118Gy struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSharpSjc118Gy creates a new seeder instance
func NewSpecificationSeederRefrigeratorSharpSjc118Gy() *SpecificationSeederRefrigeratorSharpSjc118Gy {
	return &SpecificationSeederRefrigeratorSharpSjc118Gy{
		BaseSeeder: BaseSeeder{name: "Specifications for sjc‑118‑gy"},
	}
}

func (s *SpecificationSeederRefrigeratorSharpSjc118Gy) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Sharp":               "শার্প",
		"SJC‑118‑GY":          "এসজিসি-১১৮-জিওয়াই",
		"Single Door":         "সিঙ্গেল দরজা",
		"118 Liters":          "১১৮ লিটার",
		"0 Liters":            "০ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"140 kWh":             "১৪০ কিলোওয়াট ঘণ্টা",
		"450 x 450 x 1000 mm": "৪৫০ x ৪৫০ x ১০০০ মিমি",
		"30 kg":               "৩০ কেজি",
		"Gray":                "গ্রে",
		"Standard":            "স্ট্যান্ডার্ড",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Mechanical":          "মেকানিক্যাল",
		"Wire Shelves":        "ওয়াইয়ার শেল্ফ",
		"2":                   "২",
		"1":                   "১",
		"No":                  "না",
		"35 dB":               "৩৫ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"5":                   "৫",
		"Large Vegetable Crisper, Stabilizer Free Operation": "লার্জ ভেজিটেবল ক্রিস্পার, স্ট্যাবিলাইজার ফ্রি অপারেশন",
		"Refrigerant":                 "রেফ্রিজারেন্ট",
		"Gross Volume":                "মোট ভলিউম",
		"Net Volume":                  "নেট ভলিউম",
		"118 Ltr.":                    "১১৮ লিটার",
		"0 Ltr.":                      "০ লিটার",
		"Energy Efficient Compressor": "এনার্জি এফিশিয়েন্ট কম্প্রেসার",
		"220 ~ 240V":                  "২২০ ~ ২৪০ভোল্ট",
		"Special Features":            "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'sjc‑118‑gy'
func (s *SpecificationSeederRefrigeratorSharpSjc118Gy) Seed(db *gorm.DB) error {
	productSlug := "sjc‑118‑gy"
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
		"Model Name":                  "SJC‑118‑GY",
		"Door Type":                   "Single Door",
		"Capacity":                    "118 Liters",
		"Refrigerator Capacity":       "118 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "140 kWh",
		"Dimensions":                  "450 x 450 x 1000 mm",
		"Weight":                      "30 kg",
		"Color":                       "Gray",
		"Compressor Type":             "Standard",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire Shelves",
		"Number of Shelves":           "2",
		"Door Bins":                   "1",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "35 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "5",
		"Refrigerant":                 "R-134a",
		"Gross Volume":                "118 Ltr.",
		"Net Volume":                  "118 Ltr.",
		"Special Features":            "Large Vegetable Crisper, Stabilizer Free Operation",
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
