package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorPanasonicA202bwhn seeds specifications/options for product 'panasonic-a202bwhn'
type SpecificationSeederRefrigeratorPanasonicA202bwhn struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorPanasonicA202bwhn creates a new seeder instance
func NewSpecificationSeederRefrigeratorPanasonicA202bwhn() *SpecificationSeederRefrigeratorPanasonicA202bwhn {
	return &SpecificationSeederRefrigeratorPanasonicA202bwhn{
		BaseSeeder: BaseSeeder{name: "Specifications for panasonic-a202bwhn"},
	}
}

func (s *SpecificationSeederRefrigeratorPanasonicA202bwhn) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":           "প্যানাসনিক",
		"A202BWHN":            "A202BWHN",
		"Single Door":         "একটি দরজা",
		"202 Liters":          "২০২ লিটার",
		"0 Liters":            "০ লিটার",
		"3 Star":              "৩ তারা",
		"3":                   "৩",
		"150 kWh":             "১৫০ কিলোওয়াট ঘণ্টা",
		"550 x 1320 x 650 mm": "৫৫০ x ১৩২০ x ৬৫০ মিমি",
		"38 kg":               "৩৮ কেজি",
		"White":               "সাদা",
		"Inverter":            "ইনভার্টার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"2":                   "২",
		"1":                   "১",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"12":                  "১২",
		"LED Lighting, Toughened Glass Shelves, Vegetable Crisper, Stabilizer Free Operation": "এলইডি লাইটিং, টাফেন্ড গ্লাস শেল্ভস, ভেজিটেবল ক্রিস্পার, স্ট্যাবিলাইজার ফ্রি অপারেশন",
		"Refrigerant":                   "রেফ্রিজারেন্ট",
		"Gross Volume":                  "মোট ভলিউম",
		"Net Volume":                    "নেট ভলিউম",
		"BLDC Inverter":                 "বিএলডিসি ইনভার্টার",
		"R-600a":                        "আর-৬০০এ",
		"202 Ltr.":                      "২০২ লিটার",
		"Eco-friendly Green Technology": "ইকো-ফ্রেন্ডলি গ্রীন টেকনোলজি",
		"220 ~ 240V":                    "২২০ ~ ২৪০ভোল্ট",
		"Special Features":              "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'panasonic-a202bwhn'
func (s *SpecificationSeederRefrigeratorPanasonicA202bwhn) Seed(db *gorm.DB) error {
	productSlug := "panasonic-a202bwhn"
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
		"Brand":                       "Panasonic",
		"Model Name":                  "A202BWHN",
		"Door Type":                   "Single Door",
		"Capacity":                    "202 Liters",
		"Refrigerator Capacity":       "202 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "3 Star",
		"Energy Star Rating":          "3",
		"Annual Energy Consumption":   "150 kWh",
		"Dimensions":                  "550 x 1320 x 650 mm",
		"Weight":                      "38 kg",
		"Color":                       "White",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "2",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "202 Ltr.",
		"Net Volume":                  "202 Ltr.",
		"Special Features":            "LED Lighting, Toughened Glass Shelves, Vegetable Crisper, Stabilizer Free Operation",
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
