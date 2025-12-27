package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS230Z seeds specifications/options for product 'singer-ftds230z'
type SpecificationSeederRefrigeratorSingerFTDS230Z struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS230Z creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS230Z() *SpecificationSeederRefrigeratorSingerFTDS230Z {
	return &SpecificationSeederRefrigeratorSingerFTDS230Z{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds230z"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS230Z) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                   "সিঙ্গার",
		"FTDS230Z":                 "FTDS230Z",
		"Top Mounted Refrigerator": "টপ মাউন্টেড রেফ্রিজারেটর",
		"231 Liters":               "231 লিটার",
		"133 Liters":               "133 লিটার",
		"98 Liters":                "98 লিটার",
		"5 Star":                   "৫ তারা",
		"5":                        "৫",
		"250 kWh":                  "২৫০ কিলোওয়াট ঘণ্টা",
		"1580 x 590 x 570 mm":      "1580 x 590 x 570 মিমি",
		"55 kg":                    "55 kg",
		"Black":                    "ব্ল্যাক",
		"Inverter":                 "ইনভার্টার",
		"Direct Cool":              "ডাইরেক্ট কুল",
		"Manual":                   "ম্যানুয়াল",
		"Mechanical":               "মেকানিক্যাল",
		"Tempered Glass Shelves":   "টেম্পার্ড গ্লাস শেল্ফ",
		"3":                        "৩",
		"1":                        "১",
		"No":                       "না",
		"40 dB":                    "৪০ ডেসিবেল",
		"135V":                     "135ভোল্ট",
		"50":                       "৫০",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসর এবং 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"10":       "10",
		"R600a":    "আর600এ",
		"233 Ltr.": "233 লিটার",
		"231 Ltr.": "231 লিটার",
		"Top Mounted Refrigerator, 45:55 Space Compartment Ratio, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, NutriLock & Fresh-O-Logy will keep your food fresh up to 20 Days, Built-in Stabilizer with 135V Running, Odor Filter, Bottle Holder": "টপ মাউন্টেড রেফ্রিজারেটর, 45:55 স্পেস কম্পার্টমেন্ট রেশিও, বিএসটিআই দ্বারা 5 তারা এনার্জি রেটিং, কম পাওয়ার কনসাম্পশন, রিয়েল টেম্পার্ড গ্লাস ফিনিশ, নুট্রিলক এবং ফ্রেশ-ও-লজি আপনার খাবারকে 20 দিন পর্যন্ত তাজা রাখবে, 135ভি রানিং সহ বিল্ট-ইন স্ট্যাবিলাইজার, ওডর ফিল্টার, বটল হোল্ডার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds230z'
func (s *SpecificationSeederRefrigeratorSingerFTDS230Z) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds230z"
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
		"Model Name":                  "FTDS230Z",
		"Door Type":                   "Top Mounted Refrigerator",
		"Capacity":                    "231 Liters",
		"Refrigerator Capacity":       "133 Liters",
		"Freezer Capacity":            "98 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1580 x 590 x 570 mm",
		"Weight":                      "55 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass Shelves",
		"Number of Shelves":           "5",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "135V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "233 Ltr.",
		"Net Volume":                  "231 Ltr.",
		"Special Features":            "Top Mounted Refrigerator, 45:55 Space Compartment Ratio, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, NutriLock & Fresh-O-Logy will keep your food fresh up to 20 Days, Built-in Stabilizer with 135V Running, Odor Filter, Bottle Holder",
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
