package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS185BUG seeds specifications/options for product 'singer-ftds185-bug'
type SpecificationSeederRefrigeratorSingerFTDS185BUG struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS185BUG creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS185BUG() *SpecificationSeederRefrigeratorSingerFTDS185BUG {
	return &SpecificationSeederRefrigeratorSingerFTDS185BUG{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds185-bug"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS185BUG) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                   "সিঙ্গার",
		"SRREF-SS300-FTDS185-BUG":  "SRREF-SS300-FTDS185-BUG",
		"Top Mounted Refrigerator": "টপ মাউন্টেড রেফ্রিজারেটর",
		"180 Liters":               "180 লিটার",
		"3":                        "৩",
		"2":                        "২",
		"5 Star":                   "৫ তারা",
		"5":                        "৫",
		"10":                       "10",
		"250 kWh":                  "২৫০ কিলোওয়াট ঘণ্টা",
		"1520 x 545 x 544 mm":      "১৫২০ x ৫৪৫ x ৫৪৪ মিমি",
		"45 kg":                    "45 kg",
		"Blue":                     "ব্লু",
		"Inverter":                 "ইনভার্টার",
		"Direct Cool":              "ডাইরেক্ট কুল",
		"Manual":                   "ম্যানুয়াল",
		"Mechanical":               "মেকানিক্যাল",
		"Tempered Glass":           "টেম্পার্ড গ্লাস",
		"1":                        "১",
		"Yes":                      "হ্যাঁ",
		"No":                       "না",
		"40 dB":                    "৪০ ডেসিবেল",
		"220~240V":                 "২২০~২৪০ভোল্ট",
		"50":                       "৫০",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসার & 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"180 Ltr Capacity, 45:55 Space Compartment Ratio, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Built-in Stabilizer, Odor Filter, Top Mounted Refrigerator, Modular Bottle Holder, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Food Grade Plastic Shelf & Crisper": "180 লিটার ক্যাপাসিটি, 45:55 স্পেস কম্পার্টমেন্ট রেশিও, BSTI দ্বারা 5 তারা এনার্জি রেটিং, লো পাওয়ার কনসাম্পশন, রিয়েল টেম্পার্ড গ্লাস ফিনিশ, টেম্পার্ড গ্লাস শেল্ফ, বিল্ট-ইন স্ট্যাবিলাইজার, ওডর ফিল্টার, টপ মাউন্টেড রেফ্রিজারেটর, মডুলার বটল হোল্ডার, হিডেন কন্ডেন্সার, বিগ ফ্রিজার, ফ্ল্যাশ এর্গোনমিক হ্যান্ডেল, আইস স্ক্রেপার, অ্যান্টি ব্যাকটেরিয়াল গ্যাসকেট, LED লাইট, ফুড গ্রেড প্লাস্টিক শেল্ফ এবং ক্রিস্পার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"180 Ltr.":         "180 লিটার",
		"R600a":            "আর-৬০০এ",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds185-bug'
func (s *SpecificationSeederRefrigeratorSingerFTDS185BUG) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds185-bug"
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
		"Model Name":                  "SRREF-SS300-FTDS185-BUG",
		"Door Type":                   "Top Mounted Refrigerator",
		"Capacity":                    "180 Liters",
		"Refrigerator Capacity":       "3",
		"Freezer Capacity":            "2",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1520 x 545 x 544 mm",
		"Weight":                      "45 kg",
		"Color":                       "Blue",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220~240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "180 Ltr.",
		"Net Volume":                  "180 Ltr.",
		"Special Features":            "180 Ltr Capacity, 45:55 Space Compartment Ratio, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Built-in Stabilizer, Odor Filter, Top Mounted Refrigerator, Modular Bottle Holder, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Food Grade Plastic Shelf & Crisper",
	}

	banglaTranslations := s.getBanglaTranslations()
for key, val := range specs {
    if len(val) > 500 {
        specs[key] = val[:500]
    }
}
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
			// Update the value if different
			if existing.Value != value {
				existing.Value = value
				if err := db.Save(&existing).Error; err != nil {
					return err
				}
			}
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
				} else {
					// Update translation if different
					if existingTranslation.Value != banglaValue {
						existingTranslation.Value = banglaValue
						if err := db.Save(&existingTranslation).Error; err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
