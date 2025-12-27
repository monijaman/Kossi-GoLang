package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSharpSjEx585PBk seeds specifications/options for product 'sj‑ex585p‑bk'
type SpecificationSeederRefrigeratorSharpSjEx585PBk struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSharpSjEx585PBk creates a new seeder instance
func NewSpecificationSeederRefrigeratorSharpSjEx585PBk() *SpecificationSeederRefrigeratorSharpSjEx585PBk {
	return &SpecificationSeederRefrigeratorSharpSjEx585PBk{
		BaseSeeder: BaseSeeder{name: "Specifications for sj‑ex585p‑bk"},
	}
}

func (s *SpecificationSeederRefrigeratorSharpSjEx585PBk) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Sharp":               "শার্প",
		"SJ‑EX585P‑BK":        "এসজে-এক্স৫৮৫পি-বিকে",
		"Side by Side":        "সাইড বাই সাইড",
		"585 Liters":          "৫৮৫ লিটার",
		"345 Liters":          "৩৪৫ লিটার",
		"240 Liters":          "২৪০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"380 kWh":             "৩৮০ কিলোওয়াট ঘণ্টা",
		"910 x 650 x 1780 mm": "৯১০ x ৬৫০ x ১৭৮০ মিমি",
		"105 kg":              "১০৫ কেজি",
		"Black":               "ব্ল্যাক",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"4":                   "৪",
		"6":                   "৬",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"10":                  "১০",
		"Inverter Compressor, LED Lighting, Multi Airflow, Humidity Control": "ইনভার্টার কম্প্রেসার, এলইডি লাইটিং, মাল্টি এয়ারফ্লো, হিউমিডিটি কন্ট্রোল",
		"Refrigerant":                  "রেফ্রিজারেন্ট",
		"Gross Volume":                 "মোট ভলিউম",
		"Net Volume":                   "নেট ভলিউম",
		"585 Ltr.":                     "৫৮৫ লিটার",
		"345 Ltr.":                     "৩৪৫ লিটার",
		"240 Ltr.":                     "২৪০ লিটার",
		"Advanced Inverter Technology": "অ্যাডভান্সড ইনভার্টার টেকনোলজি",
		"220 ~ 240V":                   "২২০ ~ ২৪০ভোল্ট",
		"Special Features":             "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'sj‑ex585p‑bk'
func (s *SpecificationSeederRefrigeratorSharpSjEx585PBk) Seed(db *gorm.DB) error {
	productSlug := "sj‑ex585p‑bk"
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
		"Model Name":                  "SJ‑EX585P‑BK",
		"Door Type":                   "Side by Side",
		"Capacity":                    "585 Liters",
		"Refrigerator Capacity":       "345 Liters",
		"Freezer Capacity":            "240 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "380 kWh",
		"Dimensions":                  "910 x 650 x 1780 mm",
		"Weight":                      "105 kg",
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
		"Noise Level":                 "38 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "585 Ltr.",
		"Net Volume":                  "585 Ltr.",
		"Special Features":            "Inverter Compressor, LED Lighting, Multi Airflow, Humidity Control",
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
