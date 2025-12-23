package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWue3c4GepbXx seeds specifications/options for product 'walton-wue-3c4-gepb-xx'
type SpecificationSeederRefrigeratorWaltonWue3c4GepbXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWue3c4GepbXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWue3c4GepbXx() *SpecificationSeederRefrigeratorWaltonWue3c4GepbXx {
	return &SpecificationSeederRefrigeratorWaltonWue3c4GepbXx{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wue-3c4-gepb-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWue3c4GepbXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WUE-3C4-GEPB-XX":     "ডব্লিউইউই-৩সি৪-জিইপিবি-এক্সএক্স",
		"Single Door":         "একটি দরজা",
		"321 Liters":          "৩২১ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"300 kWh":             "৩০০ কিলোওয়াট ঘণ্টা",
		"594 x 760 x 1820 mm": "৫৯৪ x ৭৬০ x ১৮২০ মিমি",
		"77 kg":               "৭৭ কেজি",
		"Silver":              "সিলভার",
		"BLDC Inverter":       "বিএলডিসি ইনভার্টার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Electronic":          "ইলেকট্রনিক",
		"Tempered Glass":      "টেম্পার্ড গ্লাস",
		"2":                   "২",
		"0":                   "০",
		"6":                   "৬",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"85-285V":             "৮৫-২৮৫ভোল্ট",
		"50/60":               "৫০/৬০",
		"1 Year":              "১ বছর",
		"12":                  "১২",
		"Intelligent Inverter, Electronic Control with Digital Display, Wide Voltage Design": "ইন্টেলিজেন্ট ইনভার্টার, ইলেকট্রনিক কন্ট্রোল উইথ ডিজিটাল ডিসপ্লে, ওয়াইড ভোল্টেজ ডিজাইন",
		"N/A": "এন/এ",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wue-3c4-gepb-xx'
func (s *SpecificationSeederRefrigeratorWaltonWue3c4GepbXx) Seed(db *gorm.DB) error {
	productSlug := "walton-wue-3c4-gepb-xx"
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
		"Special Features":            69,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
	}

	specs := map[string]string{
		"Brand":                       "Walton",
		"Model Name":                  "WUE-3C4-GEPB-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "321 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "321 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "300 kWh",
		"Dimensions":                  "594 x 760 x 1820 mm",
		"Weight":                      "77 kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "2",
		"Door Bins":                   "0",
		"Crisper Drawers":             "6",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "38 dB",
		"Voltage":                     "85-285V",
		"Frequency (Hz)":              "50/60",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "1 Year",
		"Compressor Warranty (Years)": "12",
		"Special Features":            "Intelligent Inverter, Electronic Control with Digital Display, Wide Voltage Design",
		"Refrigerant":                 "N/A",
		"Gross Volume":                "N/A",
		"Net Volume":                  "N/A",
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
