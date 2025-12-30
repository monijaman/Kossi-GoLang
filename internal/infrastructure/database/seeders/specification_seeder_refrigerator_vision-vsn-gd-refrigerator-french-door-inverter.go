package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter seeds specifications/options for product 'vision-vsn-gd-refrigerator-french-door-inverter'
type SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter() *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter {
	return &SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vsn-gd-refrigerator-french-door-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VSN GD Refrigerator French Door Inverter": "ভিএসএন জিডি রেফ্রিজারেটর ফ্রেঞ্চ ডোর ইনভার্টার",
		"French Door":         "ফ্রেঞ্চ ডোর",
		"558 Liters":          "৫৫৮ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"295 kWh":             "২৯৫ কিলোওয়াট ঘণ্টা",
		"660 x 604 x 1898 mm": "৬৬০ x ৬০৪ x ১৮৯৮ মিমি",
		"100 kg":              "১০০ কেজি",
		"Silver":              "সিলভার",
		"Inverter":            "ইনভার্টার",
		"No Frost":            "নো ফ্রস্ট",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"3":                   "৩",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ভোল্ট",
		"50":                  "৫০",
		"5 Years":             "৫ বছর",
		"10":                  "১০",
		"R600a":               "আর-৬০০এ",
		"Efficient Inverter compressor, No Frost Technology, Multi Air Flow System, Interior LED Light, Twist ice maker, Tempered glass shelves, Strong metal airflow protector, External Water Dispenser, Elegant Interior design, Tropical-grade durability, Wide voltage adaptability": "এফিশিয়েন্ট ইনভার্টার কম্প্রেসর, নো ফ্রস্ট টেকনোলজি, মাল্টি এয়ার ফ্লো সিস্টেম, ইন্টেরিয়র LED লাইট, টুইস্ট আইস মেকার, টেম্পার্ড গ্লাস শেল্ভস, স্ট্রং মেটাল এয়ারফ্লো প্রোটেক্টর, এক্সটার্নাল ওয়াটার ডিসপেনসার, এলিগ্যান্ট ইন্টেরিয়র ডিজাইন, ট্রপিকাল-গ্রেড ডুরাবিলিটি, ওয়াইড ভোল্টেজ অ্যাডাপ্টাবিলিটি",
		"558 Ltr.":   "৫৫৮ লিটার",
		"Mechanical": "মেকানিকাল",
		"Automatic":  "অটোমেটিক",
		"Normal":     "নরমাল",
		"4":          "৪",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vsn-gd-refrigerator-french-door-inverter'
func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorFrenchDoorInverter) Seed(db *gorm.DB) error {
	productSlug := "vision-vsn-gd-refrigerator-french-door-inverter"
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
		"Brand":                       "Vision",
		"Model Name":                  "VSN GD Refrigerator French Door Inverter",
		"Door Type":                   "French Door",
		"Capacity":                    "558 Liters",
		"Refrigerator Capacity":       "558 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "660 x 604 x 1898 mm",
		"Weight":                      "100 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "No Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "Yes",
		"Noise Level":                 "40 dB",
		"Voltage":                     "160 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "558 Ltr.",
		"Net Volume":                  "558 Ltr.",
		"Special Features":            "Efficient Inverter compressor, No Frost Technology, Multi Air Flow System, Interior LED Light, Twist ice maker, Tempered glass shelves, Strong metal airflow protector, External Water Dispenser, Elegant Interior design, Tropical-grade durability, Wide voltage adaptability",
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
