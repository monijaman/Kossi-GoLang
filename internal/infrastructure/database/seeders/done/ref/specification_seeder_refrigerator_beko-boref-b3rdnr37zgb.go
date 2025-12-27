package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb seeds specifications/options for product 'boref-b3rdnr37zgb'
type SpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb creates a new seeder instance
func NewSpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb() *SpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb {
	return &SpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb{
		BaseSeeder: BaseSeeder{name: "Specifications for boref-b3rdnr37zgb"},
	}
}

func (s *SpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Beko":                "বেকো",
		"BOREF-B3RDNR37ZGB":   "বোরেফ-বি৩আরডিএনআর৩৭জেডজিবি",
		"Bottom Freezer":      "নিচের ফ্রিজার",
		"370 Liters":          "৩৭০ লিটার",
		"120 Liters":          "১২০ লিটার",
		"250 Liters":          "২৫০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"280 kWh":             "২৮০ কিলোওয়াট ঘণ্টা",
		"595 x 600 x 1850 mm": "৫৯৫ x ৬০০ x ১৮৫০ মিমি",
		"62 kg":               "৬২ কেজি",
		"Black":               "কালো",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"3":                   "৩",
		"4":                   "৪",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"12":                  "১২",
		"Inverter Technology, Frost Free, LED Lighting": "ইনভার্টার প্রযুক্তি, ফ্রস্ট ফ্রি, এলইডি লাইটিং",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"R-600a":           "আর-৬০০এ",
		"370 Ltr.":         "৩৭০ লিটার",
		"120 Ltr.":         "১২০ লিটার",
		"250 Ltr.":         "২৫০ লিটার",
		"220 ~ 240V":       "২২০ ~ ২৪০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'boref-b3rdnr37zgb'
func (s *SpecificationSeederRefrigeratorBekoBorefB3rdnr37zgb) Seed(db *gorm.DB) error {
	productSlug := "boref-b3rdnr37zgb"
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
		"Model Name":                  "BOREF-B3RDNR37ZGB",
		"Door Type":                   "Bottom Freezer",
		"Capacity":                    "370 Liters",
		"Refrigerator Capacity":       "250 Liters",
		"Freezer Capacity":            "120 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "280 kWh",
		"Dimensions":                  "595 x 600 x 1850 mm",
		"Weight":                      "62 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Frost Free",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "38 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "370 Ltr.",
		"Net Volume":                  "370 Ltr.",
		"Special Features":            "Inverter Technology, Frost Free, LED Lighting",
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
