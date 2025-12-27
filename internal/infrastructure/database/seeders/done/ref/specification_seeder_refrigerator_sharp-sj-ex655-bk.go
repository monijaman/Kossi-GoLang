package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSharpSjEx655Bk seeds specifications/options for product 'sj‑ex655‑bk'
type SpecificationSeederRefrigeratorSharpSjEx655Bk struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSharpSjEx655Bk creates a new seeder instance
func NewSpecificationSeederRefrigeratorSharpSjEx655Bk() *SpecificationSeederRefrigeratorSharpSjEx655Bk {
	return &SpecificationSeederRefrigeratorSharpSjEx655Bk{
		BaseSeeder: BaseSeeder{name: "Specifications for sj‑ex655‑bk"},
	}
}

func (s *SpecificationSeederRefrigeratorSharpSjEx655Bk) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Sharp":               "শার্প",
		"SJ‑EX655‑BK":         "এসজে-এক্স৬৫৫-বিকে",
		"Side by Side":        "সাইড বাই সাইড",
		"655 Liters":          "৬৫৫ লিটার",
		"385 Liters":          "৩৮৫ লিটার",
		"270 Liters":          "২৭০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"420 kWh":             "৪২০ কিলোওয়াট ঘণ্টা",
		"910 x 700 x 1780 mm": "৯১০ x ৭০০ x ১৭৮০ মিমি",
		"110 kg":              "১১০ কেজি",
		"Black":               "ব্ল্যাক",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"4":                   "৪",
		"6":                   "৬",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, LED Lighting, Multi Airflow, Ice Dispenser": "ইনভার্টার কম্প্রেসার, এলইডি লাইটিং, মাল্টি এয়ারফ্লো, আইস ডিসপেন্সার",
		"Refrigerant":                  "রেফ্রিজারেন্ট",
		"Gross Volume":                 "মোট ভলিউম",
		"Net Volume":                   "নেট ভলিউম",
		"655 Ltr.":                     "৬৫৫ লিটার",
		"385 Ltr.":                     "৩৮৫ লিটার",
		"270 Ltr.":                     "২৭০ লিটার",
		"Advanced Inverter Technology": "অ্যাডভান্সড ইনভার্টার টেকনোলজি",
		"220 ~ 240V":                   "২২০ ~ ২৪০ভোল্ট",
		"Special Features":             "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'sj‑ex655‑bk'
func (s *SpecificationSeederRefrigeratorSharpSjEx655Bk) Seed(db *gorm.DB) error {
	productSlug := "sj‑ex655‑bk"
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
		"Model Name":                  "SJ‑EX655‑BK",
		"Door Type":                   "Side by Side",
		"Capacity":                    "655 Liters",
		"Refrigerator Capacity":       "385 Liters",
		"Freezer Capacity":            "270 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "420 kWh",
		"Dimensions":                  "910 x 700 x 1780 mm",
		"Weight":                      "110 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "6",
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
		"Gross Volume":                "655 Ltr.",
		"Net Volume":                  "655 Ltr.",
		"Special Features":            "Inverter Compressor, LED Lighting, Multi Airflow, Ice Dispenser",
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
