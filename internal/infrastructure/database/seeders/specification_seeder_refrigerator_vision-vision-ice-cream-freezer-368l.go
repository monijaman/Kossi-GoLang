package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L seeds specifications/options for product 'vision-vision-ice-cream-freezer-368l'
type SpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L() *SpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L {
	return &SpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vision-ice-cream-freezer-368l"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VISION Ice Cream Freezer 368L": "VISION Ice Cream Freezer 368L",
		"Single Door": "Single Door",
		"368 Liters": "368 লিটার",
		"0 Liters": "০ লিটার",
		"5 Star": "৫ তারা",
		"5": "৫",
		"295 kWh": "২৯৫ কিলোওয়াট ঘণ্টা",
		"550 x 570 x 1680 mm": "৫৫০ x ৫৭০ x ১৬৮০ মিমি",
		"418 kg": "418 কেজি",
		"Silver": "Silver",
		"Direct Cool": "Direct Cool",
		"Frost Free": "ফ্রস্ট ফ্রি",
		"Electronic": "ইলেকট্রনিক",
		"Toughened Glass": "টাফেন্ড গ্লাস",
		"3": "৩",
		"2": "২",
		"Yes": "হ্যাঁ",
		"No": "না",
		"40 dB": "৪০ ডেসিবেল",
		"220 ~ 240V": "২২০ ~ ২৪০ভোল্ট",
		"50": "৫০",
		"3 Years": "৩ বছর",
		"10": "১০",
		"R-600a": "আর-৬০০এ",
		"Eco-friendly Technology": "ইকো-ফ্রেন্ডলি টেকনোলজি",
		"220V": "২২০ভোল্ট",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant": "রেফ্রিজারেন্ট",
		"Gross Volume": "মোট ভলিউম",
		"Net Volume": "নেট ভলিউম",
		"BLDC Inverter": "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vision-ice-cream-freezer-368l'
func (s *SpecificationSeederRefrigeratorVisionVISIONIceCreamFreezer368L) Seed(db *gorm.DB) error {
	productSlug := "vision-vision-ice-cream-freezer-368l"
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
		"Brand": "Vision",
		"Model Name": "VISION Ice Cream Freezer 368L",
		"Door Type": "Single Door",
		"Capacity": "368 Liters",
		"Refrigerator Capacity": "368 Liters",
		"Freezer Capacity": "0 Liters",
		"Energy Efficiency Rating": "5 Star",
		"Energy Star Rating": "5",
		"Annual Energy Consumption": "295 kWh",
		"Dimensions": "550 x 570 x 1680 mm",
		"Weight": "418 kg",
		"Color": "Silver",
		"Compressor Type": "Normal",
		"Cooling Technology": "Non-Frost",
		"Defrost Type": "Automatic",
		"Temperature Control": "Electronic",
		"Shelf Material": "Toughened Glass",
		"Number of Shelves": "3",
		"Door Bins": "4",
		"Crisper Drawers": "2",
		"Ice Maker": "Yes",
		"Water Dispenser": "No",
		"Noise Level": "40 dB",
		"Voltage": "220 ~ 240V",
		"Frequency (Hz)": "50",
		"App Control": "No",
		"Voice Assistant Support": "No",
		"Warranty": "3 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant": "R-600a",
		"Gross Volume": "368 Ltr.",
		"Net Volume": "368 Ltr.",
		"Special Features": "Eco-friendly Technology",
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
