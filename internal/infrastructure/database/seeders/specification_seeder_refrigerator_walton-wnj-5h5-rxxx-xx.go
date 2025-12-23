package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx seeds specifications/options for product 'walton-wnj-5h5-rxxx-xx'
type SpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx() *SpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx {
	return &SpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wnj-5h5-rxxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WNJ-5H5-RXXX-XX":     "ডব্লিউএনজে-৫এইচ৫-আরএক্সএক্সএক্স-এক্সএক্স",
		"Single Door":         "একটি দরজা",
		"555 Liters":          "৫৫৫ লিটার",
		"400 Liters":          "৪০০ লিটার",
		"155 Liters":          "১৫৫ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"200 kWh":             "২০০ কিলোওয়াট ঘণ্টা",
		"700 x 650 x 1700 mm": "৭০০ x ৬৫০ x ১৭০০ মিমি",
		"80 kg":               "৮০ কেজি",
		"Silver":              "সিলভার",
		"Inverter":            "ইনভার্টার",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"2":                   "২",
		"1":                   "১",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"Large Vegetable Crisper, Egg Tray, Ice Cube Tray": "বড় শাকসবজি ক্রিস্পার, ডিমের ট্রে, বরফের কিউব ট্রে",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"R600a (CFC Free)": "আর৬০০এ (সিএফসি ফ্রি)",
		"500 Liters":       "৫০০ লিটার",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wnj-5h5-rxxx-xx'
func (s *SpecificationSeederRefrigeratorWaltonWnj5h5RxxxXx) Seed(db *gorm.DB) error {
	productSlug := "walton-wnj-5h5-rxxx-xx"
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
		"Model Name":                  "WNJ-5H5-RXXX-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "555 Liters",
		"Refrigerator Capacity":       "400 Liters",
		"Freezer Capacity":            "155 Liters",
		"Energy Efficiency Rating":    "4 Star",
		"Energy Star Rating":          "4",
		"Annual Energy Consumption":   "200 kWh",
		"Dimensions":                  "700 x 650 x 1700 mm",
		"Weight":                      "80 kg",
		"Color":                       "Silver",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Toughened Glass",
		"Number of Shelves":           "2",
		"Door Bins":                   "4",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "38 dB",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a (CFC Free)",
		"Gross Volume":                "555 Liters",
		"Net Volume":                  "500 Liters",
		"Special Features":            "Large Vegetable Crisper, Egg Tray, Ice Cube Tray",
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
