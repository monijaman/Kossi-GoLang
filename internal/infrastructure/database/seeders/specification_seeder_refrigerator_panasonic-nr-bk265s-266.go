package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorPanasonicNrBk265s266 seeds specifications/options for product 'panasonic-nr-bk265s-266'
type SpecificationSeederRefrigeratorPanasonicNrBk265s266 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorPanasonicNrBk265s266 creates a new seeder instance
func NewSpecificationSeederRefrigeratorPanasonicNrBk265s266() *SpecificationSeederRefrigeratorPanasonicNrBk265s266 {
	return &SpecificationSeederRefrigeratorPanasonicNrBk265s266{
		BaseSeeder: BaseSeeder{name: "Specifications for panasonic-nr-bk265s-266"},
	}
}

func (s *SpecificationSeederRefrigeratorPanasonicNrBk265s266) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":              "প্যানাসনিক",
		"NR‑BK265S/266":     "ডব্লিউইউটি-৩এক্স৮-জিইআরই-সিএক্স",
		"Single Door":         "একটি দরজা",
		"308 Liters":          "৩০৮ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"300 kWh":             "৩০০ কিলোওয়াট ঘণ্টা",
		"600 x 600 x 1700 mm": "৬০০ x ৬০০ x ১৭০০ মিমি",
		"60 kg":               "৬০ কেজি",
		"Silver":              "সিলভার",
		"Inverter":            "ইনভার্টার",
		"Frost Free":          "ফ্রস্ট ফ্রি",
		"Electronic":          "ইলেকট্রনিক",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"3":                   "৩",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"42 dB":               "৪২ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"3 Years":             "৩ বছর",
		"12":                  "১২",
		"Inverter Technology, Large Vegetable Crisper, Egg Tray, Ice Cube Tray": "ইনভার্টার প্রযুক্তি, বড় শাকসবজি ক্রিস্পার, ডিমের ট্রে, বরফের কিউব ট্রে",
		"Refrigerant":                   "রেফ্রিজারেন্ট",
		"Gross Volume":                  "মোট ভলিউম",
		"Net Volume":                    "নেট ভলিউম",
		"596 x 655 x 1860 mm":           "৫৯৬ x ৬৫৫ x ১৮৬০ মিমি",
		"81 ± 2 Kg":                     "৮১ ± ২ কেজি",
		"BLDC Inverter":                 "বিএলডিসি ইনভার্টার",
		"R-600a":                        "আর-৬০০এ",
		"308 Ltr.":                      "৩০৮ লিটার",
		"Eco-friendly Green Technology": "ইকো-ফ্রেন্ডলি গ্রীন টেকনোলজি",
		"220 ~ 240V":                    "২২০ ~ ২৪০ভোল্ট",
		"Special Features":              "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'panasonic-nr-bk265s-266'
func (s *SpecificationSeederRefrigeratorPanasonicNrBk265s266) Seed(db *gorm.DB) error {
	productSlug := "panasonic-nr-bk265s-266"
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
		"Brand":                       "Panasonic",
		"Model Name":                  "NR‑BK265S/266",
		"Door Type":                   "Single Door",
		"Capacity":                    "308 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "308 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "300 kWh",
		"Dimensions":                  "596 x 655 x 1860 mm",
		"Weight":                      "81 ± 2 Kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Non-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "5",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "3 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R-600a",
		"Gross Volume":                "308 Ltr.",
		"Net Volume":                  "308 Ltr.",
		"Special Features":            "Eco-friendly Green Technology",
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
