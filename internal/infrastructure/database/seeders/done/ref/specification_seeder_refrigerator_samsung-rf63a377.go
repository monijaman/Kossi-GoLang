package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSamsungRf63a377 seeds specifications/options for product 'samsung-rf63a377'
type SpecificationSeederRefrigeratorSamsungRf63a377 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSamsungRf63a377 creates a new seeder instance
func NewSpecificationSeederRefrigeratorSamsungRf63a377() *SpecificationSeederRefrigeratorSamsungRf63a377 {
	return &SpecificationSeederRefrigeratorSamsungRf63a377{
		BaseSeeder: BaseSeeder{name: "Specifications for samsung-rf63a377"},
	}
}

func (s *SpecificationSeederRefrigeratorSamsungRf63a377) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Samsung":             "স্যামসাং",
		"RF63A377":            "আরএফ৬৩এ৩৭৭",
		"French Door":         "ফ্রেঞ্চ ডোর",
		"630 Liters":          "৬৩০ লিটার",
		"430 Liters":          "৪৩০ লিটার",
		"200 Liters":          "২০০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"400 kWh":             "৪০০ কিলোওয়াট ঘণ্টা",
		"912 x 716 x 1789 mm": "৯১২ x ৭১৬ x ১৭৮৯ মিমি",
		"105 kg":              "১০৫ কেজি",
		"Refined Inox":        "রিফাইন্ড ইনক্স",
		"Digital Inverter":    "ডিজিটাল ইনভার্টার",
		"Twin Cooling Plus":   "টুইন কুলিং প্লাস",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"7":                   "৭",
		"6":                   "৬",
		"3":                   "৩",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"41 dB":               "৪১ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"R-600a":              "আর-৬০০এ",
		"Digital Inverter Compressor, Twin Cooling Plus, FlexZone, Metal Cooling": "ডিজিটাল ইনভার্টার কম্প্রেসর, টুইন কুলিং প্লাস, ফ্লেক্সজোন, মেটাল কুলিং",
	}
}

// Seed inserts specification records for the product identified by slug 'samsung-rf63a377'
func (s *SpecificationSeederRefrigeratorSamsungRf63a377) Seed(db *gorm.DB) error {
	productSlug := "samsung-rf63a377"
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
		"Model Name":                  "RF63A377",
		"Door Type":                   "French Door",
		"Capacity":                    "630 Liters",
		"Refrigerator Capacity":       "430 Liters",
		"Freezer Capacity":            "200 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "400 kWh",
		"Dimensions":                  "912 x 716 x 1789 mm",
		"Weight":                      "105 kg",
		"Color":                       "Refined Inox",
		"Compressor Type":             "Digital Inverter",
		"Cooling Technology":          "Twin Cooling Plus",
		"Defrost Type":                "Frost Free",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "7",
		"Door Bins":                   "6",
		"Crisper Drawers":             "3",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "41 dB",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R-600a",
		"Special Features":            "Digital Inverter Compressor, Twin Cooling Plus, FlexZone, Metal Cooling",
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
