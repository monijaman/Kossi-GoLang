package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWfc3f5GdelXx seeds specifications/options for product 'walton-wfc-3f5-gdel-xx'
type SpecificationSeederRefrigeratorWaltonWfc3f5GdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWfc3f5GdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWfc3f5GdelXx() *SpecificationSeederRefrigeratorWaltonWfc3f5GdelXx {
	return &SpecificationSeederRefrigeratorWaltonWfc3f5GdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wfc-3f5-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWfc3f5GdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":                          "ওয়ালটন",
		"WFC-3F5-GDEL-XX":                 "ডব্লিউএফসি-৩এফ৫-জিডিইএল-এক্সএক্স",
		"Single Door":                     "একটি দরজা",
		"380 Liters":                      "৩৮০ লিটার",
		"365 Liters":                      "৩৬৫ লিটার",
		"N/A":                             "এন/এ",
		"5 Star":                          "৫ তারা",
		"5":                               "৫",
		"Silver":                          "সিলভার",
		"Intelligent Inverter":            "ইন্টেলিজেন্ট ইনভার্টার",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Manual":                          "ম্যানুয়াল",
		"Mechanical":                      "মেকানিক্যাল",
		"Toughened Glass":                 "টাফেন্ড গ্লাস",
		"2":                               "২",
		"4":                               "৪",
		"1":                               "১",
		"No":                              "না",
		"220V":                            "২২০ভোল্ট",
		"50":                              "৫০",
		"2 Years":                         "২ বছর",
		"10":                              "১০",
		"Wide Voltage Design, Glass Door": "ওয়াইড ভোল্টেজ ডিজাইন, গ্লাস দরজা",
		"Refrigerant":                     "রেফ্রিজারেন্ট",
		"Gross Volume":                    "মোট ভলিউম",
		"Net Volume":                      "নেট ভলিউম",
		"BLDC/RSCR":                       "বিএলডিসি/আরএসসিআর",
		"Wire":                            "ওয়্যার",
		"3":                               "৩",
		"6":                               "৬",
		"Yes":                             "হ্যাঁ",
		"220-240V":                        "২২০-২৪০ভোল্ট",
		"R600a":                           "আর৬০০এ",
		"67 kg":                           "৬৭ কেজি",
		"650 x 650 x 1860 mm":             "৬৫০ x ৬৫০ x ১৮৬০ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wfc-3f5-gdel-xx'
func (s *SpecificationSeederRefrigeratorWaltonWfc3f5GdelXx) Seed(db *gorm.DB) error {
	productSlug := "walton-wfc-3f5-gdel-xx"
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
		"Brand":                       "Walton",
		"Model Name":                  "WFC-3F5-GDEL-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "380 Liters",
		"Refrigerator Capacity":       "365 Liters",
		"Freezer Capacity":            "N/A",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "650 x 650 x 1860 mm",
		"Weight":                      "67 kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC/RSCR",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "5",
		"Door Bins":                   "6",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220-240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "380 Liters",
		"Net Volume":                  "365 Liters",
		"Special Features":            "Wide Voltage Design, Glass Door",
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
