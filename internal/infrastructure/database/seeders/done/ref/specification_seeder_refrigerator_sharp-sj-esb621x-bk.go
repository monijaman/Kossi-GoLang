package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSharpSjEsb621xBk seeds specifications/options for product 'sj‑esb621x‑bk'
type SpecificationSeederRefrigeratorSharpSjEsb621xBk struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSharpSjEsb621xBk creates a new seeder instance
func NewSpecificationSeederRefrigeratorSharpSjEsb621xBk() *SpecificationSeederRefrigeratorSharpSjEsb621xBk {
	return &SpecificationSeederRefrigeratorSharpSjEsb621xBk{
		BaseSeeder: BaseSeeder{name: "Specifications for sj‑esb621x‑bk"},
	}
}

func (s *SpecificationSeederRefrigeratorSharpSjEsb621xBk) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Sharp":               "শার্প",
		"SJ‑ESB621X‑BK":       "এসজি-ইএসবি৬২১এক্স-বিকে",
		"Bottom Mount":        "বটম মাউন্ট",
		"621 Liters":          "৬২১ লিটার",
		"406 Liters":          "৪০৬ লিটার",
		"215 Liters":          "২১৫ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"420 kWh":             "৪২০ কিলোওয়াট ঘণ্টা",
		"700 x 700 x 1850 mm": "৭০০ x ৭০০ x ১৮৫০ মিমি",
		"85 kg":               "৮৫ কেজি",
		"Black":               "কালো",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Automatic":           "অটোমেটিক",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"4":                   "৪",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"10 Years":            "১০ বছর",
		"10":                  "১০",
		"Inverter Compressor, Multi Airflow Technology, LED Lighting, Door Alarm": "ইনভার্টার কম্প্রেসার, মাল্টি এয়ারফ্লো টেকনোলজি, এলইডি লাইটিং, দরজা অ্যালার্ম",
		"Refrigerant":         "রেফ্রিজারেন্ট",
		"Gross Volume":        "মোট ভলিউম",
		"Net Volume":          "নেট ভলিউম",
		"621 Ltr.":            "৬২১ লিটার",
		"406 Ltr.":            "৪০৬ লিটার",
		"215 Ltr.":            "২১৫ লিটার",
		"Inverter Compressor": "ইনভার্টার কম্প্রেসার",
		"220 ~ 240V":          "২২০ ~ ২৪০ভোল্ট",
		"Special Features":    "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'sj‑esb621x‑bk'
func (s *SpecificationSeederRefrigeratorSharpSjEsb621xBk) Seed(db *gorm.DB) error {
	productSlug := "sj‑esb621x‑bk"
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
		"Model Name":                  "SJ‑ESB621X‑BK",
		"Door Type":                   "Bottom Mount",
		"Capacity":                    "621 Liters",
		"Refrigerator Capacity":       "406 Liters",
		"Freezer Capacity":            "215 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "420 kWh",
		"Dimensions":                  "700 x 700 x 1850 mm",
		"Weight":                      "85 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "621 Ltr.",
		"Net Volume":                  "621 Ltr.",
		"Special Features":            "Inverter Compressor, Multi Airflow Technology, LED Lighting, Door Alarm",
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
