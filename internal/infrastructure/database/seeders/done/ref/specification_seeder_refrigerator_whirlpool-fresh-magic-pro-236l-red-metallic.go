package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic seeds specifications/options for product 'whirlpool-fresh-magic-pro-236l-red-metallic'
type SpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic creates a new seeder instance
func NewSpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic() *SpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic {
	return &SpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic{
		BaseSeeder: BaseSeeder{name: "Specifications for whirlpool-fresh-magic-pro-236l-red-metallic"},
	}
}

func (s *SpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Whirlpool":                         "হুইরলপুল",
		"Fresh Magic Pro 236L Red Metallic": "ফ্রেশ ম্যাজিক প্রো ২৩৬এল রেড মেটালিক",
		"Single Door":                       "একটি দরজা",
		"236 Liters":                        "২৩৬ লিটার",
		"0 Liters":                          "০ লিটার",
		"5 Star":                            "৫ তারা",
		"5":                                 "৫",
		"220 kWh":                           "২২০ কিলোওয়াট ঘণ্টা",
		"550 x 570 x 1500 mm":               "৫৫০ x ৫৭০ x ১৫০০ মিমি",
		"45 kg":                             "৪৫ কেজি",
		"Red Metallic":                      "রেড মেটালিক",
		"Inverter":                          "ইনভার্টার",
		"Frost Free":                        "ফ্রস্ট ফ্রি",
		"Electronic":                        "ইলেকট্রনিক",
		"Toughened Glass":                   "টাফেন্ড গ্লাস",
		"3":                                 "৩",
		"2":                                 "২",
		"Yes":                               "হ্যাঁ",
		"No":                                "না",
		"40 dB":                             "৪০ ডেসিবেল",
		"220V":                              "২২০ভোল্ট",
		"50":                                "৫০",
		"3 Years":                           "৩ বছর",
		"10":                                "১০",
		"R-600a":                            "আর-৬০০এ",
		"Inverter Technology":               "ইনভার্টার প্রযুক্তি",
		"Refrigerant":                       "রেফ্রিজারেন্ট",
		"Gross Volume":                      "মোট ভলিউম",
		"Net Volume":                        "নেট ভলিউম",
		"BLDC Inverter":                     "বিএলডিসি ইনভার্টার",
		"Fresh Magic Technology":            "ফ্রেশ ম্যাজিক প্রযুক্তি",
		"220 ~ 240V":                        "২২০ ~ ২৪০ভোল্ট",
		"Special Features":                  "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'whirlpool-fresh-magic-pro-236l-red-metallic'
func (s *SpecificationSeederRefrigeratorWhirlpoolFreshMagicPro236lRedMetallic) Seed(db *gorm.DB) error {
	productSlug := "whirlpool-fresh-magic-pro-236l-red-metallic"
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
		"Brand":                       "Whirlpool",
		"Model Name":                  "Fresh Magic Pro 236L Red Metallic",
		"Door Type":                   "Single Door",
		"Capacity":                    "236 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "236 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "220 kWh",
		"Dimensions":                  "550 x 570 x 1500 mm",
		"Weight":                      "45 kg",
		"Color":                       "Red Metallic",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Non-Frost",
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
		"Gross Volume":                "236 Ltr.",
		"Net Volume":                  "236 Ltr.",
		"Special Features":            "Fresh Magic Technology",
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
