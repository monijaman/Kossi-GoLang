package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN seeds specifications/options for product 'vision-vsn-gd-refrigerator-re-150l-digital-red-flower-en'
type SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN() *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN {
	return &SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vsn-gd-refrigerator-re-150l-digital-red-flower-en"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VSN GD Refrigerator RE-150L Digital Red Flower EN": "ভিএসএন জিডি রেফ্রিজারেটর আরই-১৫০এল ডিজিটাল রেড ফ্লাওয়ার ইএন",
		"Single Door":           "সিঙ্গেল ডোর",
		"150 Liters":            "১৫০ লিটার",
		"0 Liters":              "০ লিটার",
		"5 Star":                "৫ তারা",
		"5":                     "৫",
		"295 kWh":               "২৯৫ কিলোওয়াট ঘণ্টা",
		"550 x 570 x 1680 mm":   "৫৫০ x ৫৭০ x ১৬৮০ মিমি",
		"48 kg":                 "৪৮ কেজি",
		"Digital Red Flower EN": "ডিজিটাল রেড ফ্লাওয়ার ইএন",
		"Direct Cool":           "সরাসরি কুলিং",
		"Frost Free":            "ফ্রস্ট ফ্রি",
		"Electronic":            "ইলেকট্রনিক",
		"Toughened Glass":       "টাফেন্ড গ্লাস",
		"3":                     "৩",
		"2":                     "২",
		"Yes":                   "হ্যাঁ",
		"No":                    "না",
		"40 dB":                 "৪০ ডেসিবেল",
		"160 ~ 260V":            "১৬০ ~ ২৬০ভোল্ট",
		"50":                    "৫০",
		"5 Years":               "৫ বছর",
		"10":                    "১০",
		"R600a":                 "আর-৬০০এ",
		"Fabulous ultra-modern glass door design, Very fast cooling speed, Hygienic clean air, Anti-bacterial gasket, Huge storage capacity, Virgin & food grade plastic liner, Low noise compressor, with lock and key, Manual defrost, Mechanical control, Interior LED light, 100% copper condenser": "ফ্যাবুলাস আল্ট্রা-মডার্ন গ্লাস ডোর ডিজাইন, খুব দ্রুত কুলিং গতি, স্বাস্থ্যকর পরিষ্কার বাতাস, অ্যান্টি-ব্যাকটেরিয়াল গ্যাসকেট, বিশাল স্টোরেজ ক্যাপাসিটি, ভার্জিন এবং ফুড গ্রেড প্লাস্টিক লাইনার, লো নয়েজ কম্প্রেসর, লক এবং কী সহ, ম্যানুয়াল ডিফ্রস্ট, মেকানিকাল কন্ট্রোল, ইন্টেরিয়র LED লাইট, ১০০% কপার কন্ডেনসার",
		"175 Ltr.":   "১৭৫ লিটার",
		"150 Ltr.":   "১৫০ লিটার",
		"Mechanical": "মেকানিকাল",
		"Manual":     "ম্যানুয়াল",
		"Normal":     "নরমাল",
		"4":          "৪",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vsn-gd-refrigerator-re-150l-digital-red-flower-en'
func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE150LDigitalRedFlowerEN) Seed(db *gorm.DB) error {
	productSlug := "vision-vsn-gd-refrigerator-re-150l-digital-red-flower-en"
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
		"Model Name":                  "VSN GD Refrigerator RE-150L Digital Red Flower EN",
		"Door Type":                   "Single Door",
		"Capacity":                    "150 Liters",
		"Refrigerator Capacity":       "150 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "550 x 570 x 1680 mm",
		"Weight":                      "48 kg",
		"Color":                       "Digital Red Flower EN",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "160 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "175 Ltr.",
		"Net Volume":                  "150 Ltr.",
		"Special Features":            "Fabulous ultra-modern glass door design, Very fast cooling speed, Hygienic clean air, Anti-bacterial gasket, Huge storage capacity, Virgin & food grade plastic liner, Low noise compressor, with lock and key, Manual defrost, Mechanical control, Interior LED light, 100% copper condenser",
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
