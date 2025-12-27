package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorHitachiRw720Puc1 seeds specifications/options for product 'hitachi-rw720puc1'
type SpecificationSeederRefrigeratorHitachiRw720Puc1 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorHitachiRw720Puc1 creates a new seeder instance
func NewSpecificationSeederRefrigeratorHitachiRw720Puc1() *SpecificationSeederRefrigeratorHitachiRw720Puc1 {
	return &SpecificationSeederRefrigeratorHitachiRw720Puc1{
		BaseSeeder: BaseSeeder{name: "Specifications for hitachi-rw720puc1"},
	}
}

func (s *SpecificationSeederRefrigeratorHitachiRw720Puc1) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hitachi":             "হিটাচি",
		"RW720PUC1":           "আরডব্লিউ৭২০পিইউসি১",
		"Wine Cooler":         "ওয়াইন কুলার",
		"720 Liters":          "৭২০ লিটার",
		"0 Liters":            "০ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"450 kWh":             "৪৫০ কিলোওয়াট ঘণ্টা",
		"680 x 750 x 1850 mm": "৬৮০ x ৭৫০ x ১৮৫০ মিমি",
		"120 kg":              "১২০ কেজি",
		"Black":               "ব্ল্যাক",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Wooden Shelves":      "ওডেন শেল্ফ",
		"8":                   "৮",
		"2":                   "২",
		"No":                  "না",
		"35 dB":               "৩৫ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Temperature Controlled Storage, LED Display, Vibration Dampening": "টেম্পারেচার কন্ট্রোল্ড স্টোরেজ, এলইডি ডিসপ্লে, ভাইব্রেশন ড্যাম্পেনিং",
		"Refrigerant":                  "রেফ্রিজারেন্ট",
		"Gross Volume":                 "মোট ভলিউম",
		"Net Volume":                   "নেট ভলিউম",
		"120 ± 2 Kg":                   "১২০ ± ২ কেজি",
		"BLDC Inverter":                "বিএলডিসি ইনভার্টার",
		"R-600a":                       "আর-৬০০এ",
		"720 Ltr.":                     "৭২০ লিটার",
		"Wine Preservation Technology": "ওয়াইন প্রিজারভেশন টেকনোলজি",
		"220 ~ 240V":                   "২২০ ~ ২৪০ভোল্ট",
		"Special Features":             "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'hitachi-rw720puc1'
func (s *SpecificationSeederRefrigeratorHitachiRw720Puc1) Seed(db *gorm.DB) error {
	productSlug := "hitachi-rw720puc1"
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
		"Brand":                       "Hitachi",
		"Model Name":                  "RW720PUC1",
		"Door Type":                   "Wine Cooler",
		"Capacity":                    "720 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "720 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "450 kWh",
		"Dimensions":                  "680 x 750 x 1850 mm",
		"Weight":                      "120 ± 2 Kg",
		"Color":                       "Black",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Wooden Shelves",
		"Number of Shelves":           "8",
		"Door Bins":                   "2",
		"Crisper Drawers":             "0",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "35 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "720 Ltr.",
		"Net Volume":                  "720 Ltr.",
		"Special Features":            "Temperature Controlled Storage, LED Display, Vibration Dampening",
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
