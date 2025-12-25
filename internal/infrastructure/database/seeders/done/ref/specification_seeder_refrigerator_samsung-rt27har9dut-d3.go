package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSamsungRt27har9dutD3 seeds specifications/options for product 'samsung-rt27har9dut-d3'
type SpecificationSeederRefrigeratorSamsungRt27har9dutD3 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSamsungRt27har9dutD3 creates a new seeder instance
func NewSpecificationSeederRefrigeratorSamsungRt27har9dutD3() *SpecificationSeederRefrigeratorSamsungRt27har9dutD3 {
	return &SpecificationSeederRefrigeratorSamsungRt27har9dutD3{
		BaseSeeder: BaseSeeder{name: "Specifications for samsung-rt27har9dut-d3"},
	}
}

func (s *SpecificationSeederRefrigeratorSamsungRt27har9dutD3) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Samsung":             "স্যামসাং",
		"RT27HAR9DUT/D3":      "আরটি২৭এইচএআর৯ডিইউটি/ডি৩",
		"Double Door":         "ডবল ডোর",
		"275 Liters":          "২৭৫ লিটার",
		"203 Liters":          "২০৩ লিটার",
		"72 Liters":           "৭২ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"265 kWh":             "২৬৫ কিলোওয়াট ঘণ্টা",
		"555 x 637 x 1655 mm": "৫৫৫ x ৬৩৭ x ১৬৫৫ মিমি",
		"54 kg":               "৫৪ কেজি",
		"Luxe Brown":          "লাক্স ব্রাউন",
		"Digital Inverter":    "ডিজিটাল ইনভার্টার",
		"All Around Cooling":  "অল এরাউন্ড কুলিং",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"3":                   "৩",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"R-600a":              "আর-৬০০এ",
		"Digital Inverter Technology, All Around Cooling, Stabilizer Free": "ডিজিটাল ইনভার্টার প্রযুক্তি, অল এরাউন্ড কুলিং, স্ট্যাবিলাইজার ফ্রি",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-rt27har9dut-d3'
func (s *SpecificationSeederRefrigeratorSamsungRt27har9dutD3) Seed(db *gorm.DB) error {
	productSlug := "samsung-rt27har9dut-d3"
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
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Special Features":            69,
	}

	specs := map[string]string{
		"Brand":                       "Samsung",
		"Model Name":                  "RT27HAR9DUT/D3",
		"Door Type":                   "Double Door",
		"Capacity":                    "275 Liters",
		"Refrigerator Capacity":       "203 Liters",
		"Freezer Capacity":            "72 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "265 kWh",
		"Dimensions":                  "555 x 637 x 1655 mm",
		"Weight":                      "54 kg",
		"Color":                       "Luxe Brown",
		"Compressor Type":             "Digital Inverter",
		"Cooling Technology":          "All Around Cooling",
		"Defrost Type":                "Frost Free",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "38 dB",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Special Features":            "Digital Inverter Technology, All Around Cooling, Stabilizer Free",
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
