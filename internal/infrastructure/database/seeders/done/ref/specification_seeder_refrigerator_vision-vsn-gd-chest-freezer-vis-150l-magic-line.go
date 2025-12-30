package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine seeds specifications/options for product 'vision-vsn-gd-chest-freezer-vis-150l-magic-line'
type SpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine() *SpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine {
	return &SpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vsn-gd-chest-freezer-vis-150l-magic-line"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VSN GD Chest Freezer VIS-150L Magic Line": "ভিএসএন জিডি চেস্ট ফ্রিজার ভিআইএস-১৫০এল ম্যাজিক লাইন",
		"Chest":              "চেস্ট",
		"150 Liters":         "১৫০ লিটার",
		"0 Liters":           "০ লিটার",
		"5 Star":             "৫ তারা",
		"5":                  "৫",
		"295 kWh":            "২৯৫ কিলোওয়াট ঘণ্টা",
		"692 x 570 x 835 mm": "৬৯২ x ৫৭০ x ৮৩৫ মিমি",
		"37 kg":              "৩৭ কেজি",
		"Silver":             "সিলভার",
		"Direct Cool":        "সরাসরি কুলিং",
		"Frost Free":         "ফ্রস্ট ফ্রি",
		"Electronic":         "ইলেকট্রনিক",
		"Tempered Glass":     "টেম্পার্ড গ্লাস",
		"3":                  "৩",
		"2":                  "২",
		"Yes":                "হ্যাঁ",
		"No":                 "না",
		"40 dB":              "৪০ ডেসিবেল",
		"220V":               "২২০ভোল্ট",
		"50":                 "৫০",
		"5 Years":            "৫ বছর",
		"10":                 "১০",
		"R600a":              "আর-৬০০এ",
		"Real Tempered glass, Royal modern color, Faster cooling speed, Huge inside space, Thick foaming/side wall preserves long time cooling, Quick freezer indicator, Transferrin glass salve Free, Three layers PCM anti corrosive body, Free wire basket, Low noise compressor, with lock and key, Adjustable thermostat, Interior LED light, Eco-friendly (100% CFC & HCFC Free) Green Technology": "রিয়েল টেম্পার্ড গ্লাস, রয়াল মডার্ন কালার, ফাস্টার কুলিং স্পিড, হিউজ ইনসাইড স্পেস, থিক ফোমিং/সাইড ওয়াল প্রিজার্ভস লং টাইম কুলিং, কুইক ফ্রিজার ইন্ডিকেটর, ট্রান্সফেরিন গ্লাস সালভ ফ্রি, থ্রি লেয়ারস পিসিএম অ্যান্টি করোসিভ বডি, ফ্রি ওয়্যার বাস্কেট, লো নয়েজ কম্প্রেসর, লক এবং কী সহ, অ্যাডজাস্টেবল থার্মোস্ট্যাট, ইন্টেরিয়র LED লাইট, ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি",
		"150 Ltr.":   "১৫০ লিটার",
		"Mechanical": "মেকানিকাল",
		"Manual":     "ম্যানুয়াল",
		"Normal":     "নরমাল",
		"4":          "৪",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vsn-gd-chest-freezer-vis-150l-magic-line'
func (s *SpecificationSeederRefrigeratorVisionVSNGDChestFreezerVIS150LMagicLine) Seed(db *gorm.DB) error {
	productSlug := "vision-vsn-gd-chest-freezer-vis-150l-magic-line"
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
		"Model Name":                  "VSN GD Chest Freezer VIS-150L Magic Line",
		"Door Type":                   "Chest",
		"Capacity":                    "150 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "150 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "692 x 570 x 835 mm",
		"Weight":                      "37 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "4",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "150 Ltr.",
		"Net Volume":                  "150 Ltr.",
		"Special Features":            "Real Tempered glass, Royal modern color, Faster cooling speed, Huge inside space, Thick foaming/side wall preserves long time cooling, Quick freezer indicator, Transferrin glass salve Free, Three layers PCM anti corrosive body, Free wire basket, Low noise compressor, with lock and key, Adjustable thermostat, Interior LED light, Eco-friendly (100% CFC & HCFC Free) Green Technology",
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
