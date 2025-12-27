package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorLgGlG322Rvbb seeds specifications/options for product 'lg-gl-g322-rvbb'
type SpecificationSeederRefrigeratorLgGlG322Rvbb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorLgGlG322Rvbb creates a new seeder instance
func NewSpecificationSeederRefrigeratorLgGlG322Rvbb() *SpecificationSeederRefrigeratorLgGlG322Rvbb {
	return &SpecificationSeederRefrigeratorLgGlG322Rvbb{
		BaseSeeder: BaseSeeder{name: "Specifications for lg-gl-g322-rvbb"},
	}
}

func (s *SpecificationSeederRefrigeratorLgGlG322Rvbb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                        "এলজি",
		"GL-G322RVBB":               "জিএল-জি৩২২আরভিবিবি",
		"Single Door":               "একটি দরজা",
		"322 Liters":                "৩২২ লিটার",
		"0 Liters":                  "০ লিটার",
		"5 Star":                    "৫ তারা",
		"5":                         "৫",
		"190 kWh":                   "১৯০ কিলোওয়াট ঘণ্টা",
		"600 x 650 x 1600 mm":       "৬০০ x ৬৫০ x ১৬০০ মিমি",
		"50 kg":                     "৫০ কেজি",
		"Violet Bloom":              "ভায়োলেট ব্লুম",
		"Inverter":                  "ইনভার্টার",
		"Direct Cool":               "ডাইরেক্ট কুল",
		"Manual":                    "ম্যানুয়াল",
		"Toughened Glass":           "টাফেন্ড গ্লাস",
		"3":                         "৩",
		"2":                         "২",
		"Yes":                       "হ্যাঁ",
		"No":                        "না",
		"38 dB":                     "৩৮ ডেসিবেল",
		"220V":                      "২২০ভোল্ট",
		"50":                        "৫০",
		"3 Years":                   "৩ বছর",
		"10":                        "১০",
		"R-600a":                    "আর-৬০০এ",
		"Inverter Technology":       "ইনভার্টার প্রযুক্তি",
		"Refrigerant":               "রেফ্রিজারেন্ট",
		"Gross Volume":              "মোট ভলিউম",
		"Net Volume":                "নেট ভলিউম",
		"Linear Compressor":         "লিনিয়ার কম্প্রেসর",
		"Smart Inverter Compressor": "স্মার্ট ইনভার্টার কম্প্রেসর",
		"220 ~ 240V":                "২২০ ~ ২৪০ভোল্ট",
		"Special Features":          "বিশেষ বৈশিষ্ট্য",
		"Antibacterial Gasket":      "অ্যান্টিব্যাকটেরিয়াল গ্যাসকেট",
		"Stabilizer Free Operation": "স্ট্যাবিলাইজার ফ্রি অপারেশন",
	}
}

// Seed inserts specification records for the product identified by slug 'lg-gl-g322-rvbb'
func (s *SpecificationSeederRefrigeratorLgGlG322Rvbb) Seed(db *gorm.DB) error {
	productSlug := "lg-gl-g322-rvbb"
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
		"Brand":                       "LG",
		"Model Name":                  "GL-G322RVBB",
		"Door Type":                   "Single Door",
		"Capacity":                    "322 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "322 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "190 kWh",
		"Dimensions":                  "600 x 650 x 1600 mm",
		"Weight":                      "50 kg",
		"Color":                       "Violet Bloom",
		"Compressor Type":             "Smart Inverter Compressor",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Manual",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "38 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "322 Ltr.",
		"Net Volume":                  "322 Ltr.",
		"Special Features":            "Antibacterial Gasket, Stabilizer Free Operation",
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
