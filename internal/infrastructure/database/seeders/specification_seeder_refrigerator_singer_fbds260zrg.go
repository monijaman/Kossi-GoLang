package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFBDS260ZRG seeds specifications/options for product 'singer-fbds260z-rg'
type SpecificationSeederRefrigeratorSingerFBDS260ZRG struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFBDS260ZRG creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFBDS260ZRG() *SpecificationSeederRefrigeratorSingerFBDS260ZRG {
	return &SpecificationSeederRefrigeratorSingerFBDS260ZRG{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-fbds260z-rg"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFBDS260ZRG) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                      "সিঙ্গার",
		"FBDS260Z-RG":                 "FBDS260Z-RG",
		"Bottom Mounted Refrigerator": "বটম মাউন্টেড রেফ্রিজারেটর",
		"260 Liters":                  "260 লিটার",
		"130 Liters":                  "130 লিটার",
		"5 Star":                      "৫ তারা",
		"5":                           "৫",
		"250 kWh":                     "২৫০ কিলোওয়াট ঘণ্টা",
		"1045 x 605 x 844 mm":         "১০৪৫ x ৬০৫ x ৮৪৪ মিমি",
		"58 kg":                       "58 kg",
		"Red":                         "রেড",
		"Inverter":                    "ইনভার্টার",
		"Direct Cool":                 "ডাইরেক্ট কুল",
		"Manual":                      "ম্যানুয়াল",
		"Mechanical":                  "মেকানিক্যাল",
		"Tempered Glass":              "টেম্পার্ড গ্লাস",
		"2":                           "২",
		"3":                           "৩",
		"1":                           "১",
		"Yes":                         "হ্যাঁ",
		"No":                          "না",
		"40 dB":                       "৪০ ডেসিবেল",
		"220 ~ 240V":                  "২২০ ~ ২৪০ভোল্ট",
		"50":                          "৫০",
		"5 Years Service":             "5 বছর সার্ভিস",
		"Bottom Mounted Refrigerator, Energy Efficient": "বটম মাউন্টেড রেফ্রিজারেটর, এনার্জি এফিশিয়েন্ট",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"260 Ltr.":         "260 লিটার",
		"R600a":            "আর-৬০০এ",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-fbds260z-rg'
func (s *SpecificationSeederRefrigeratorSingerFBDS260ZRG) Seed(db *gorm.DB) error {
	productSlug := "singer-fbds260z-rg"
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
		"Brand":                       "Singer",
		"Model Name":                  "FBDS260Z-RG",
		"Door Type":                   "Bottom Mounted Refrigerator",
		"Capacity":                    "260 Liters",
		"Refrigerator Capacity":       "130 Liters",
		"Freezer Capacity":            "130 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1045 x 605 x 844 mm",
		"Weight":                      "58 kg",
		"Color":                       "Red",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220 ~ 240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years Service",
		"Compressor Warranty (Years)": "2",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "260 Ltr.",
		"Net Volume":                  "260 Ltr.",
		"Special Features":            "Bottom Mounted Refrigerator, Energy Efficient",
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
