package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS277ZWFBG seeds specifications/options for product 'singer-ftds277zwf-bg'
type SpecificationSeederRefrigeratorSingerFTDS277ZWFBG struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS277ZWFBG creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS277ZWFBG() *SpecificationSeederRefrigeratorSingerFTDS277ZWFBG {
	return &SpecificationSeederRefrigeratorSingerFTDS277ZWFBG{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds277zwf-bg"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS277ZWFBG) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                    "সিঙ্গার",
		"FTDS277ZWF-BG":             "FTDS277ZWF-BG",
		"SRREF-SS500-FTDS277ZWF-BG": "SRREF-SS500-FTDS277ZWF-BG",
		"Top Mounted Refrigerator":  "টপ মাউন্টেড রেফ্রিজারেটর",
		"272 Liters":                "272 লিটার",
		"5 Star":                    "৫ তারা",
		"5":                         "৫",
		"250 kWh":                   "২৫০ কিলোওয়াট ঘণ্টা",
		"1799 x 595 x 605 mm":       "1799 x 595 x 605 মিমি",
		"221 kg":                    "221 kg",
		"60 kg":                     "60 কেজি",
		"Black":                     "ব্ল্যাক",
		"Inverter":                  "ইনভার্টার",
		"Direct Cool":               "ডাইরেক্ট কুল",
		"Manual":                    "ম্যানুয়াল",
		"Mechanical":                "মেকানিক্যাল",
		"Wire Shelves":              "ওয়্যার শেল্ফ",
		"2":                         "২",
		"3":                         "৩",
		"1":                         "১",
		"Yes":                       "হ্যাঁ",
		"No":                        "না",
		"40 dB":                     "৪০ ডেসিবেল",
		"220 ~ 240V":                "২২০ ~ ২৪০ভোল্ট",
		"135V":                      "135ভোল্ট",
		"50":                        "৫০",
		"5 Years Service":           "5 বছর সার্ভিস",
		"Refrigerant":               "রেফ্রিজারেন্ট",
		"Gross Volume":              "মোট ভলিউম",
		"Net Volume":                "নেট ভলিউম",
		"277 Ltr.":                  "277 লিটার",
		"R-134a":                    "আর-১৩৪এ",
		"R600a":                     "আর-৬০০এ",
		"NutriLock, Fresh-O-Logy, Modular Odour Filter, Modular Bottle Holder, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Keep Fresh Up To 20 Days, Food Grade Plastic Shelf & Crisper": "নুট্রিলক, ফ্রেশ-ও-লজি, মডুলার অডার ফিল্টার, মডুলার বটল হোল্ডার, হিডেন কন্ডেন্সার, বিগ ফ্রিজার, ফ্ল্যাশ এর্গোনমিক হ্যান্ডেল, আইস স্ক্রেপার, অ্যান্টি ব্যাকটেরিয়াল গ্যাসকেট, LED লাইট, 20 দিন পর্যন্ত ফ্রেশ রাখুন, ফুড গ্রেড প্লাস্টিক শেল্ফ এবং ক্রিস্পার",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds277zwf-bg'
func (s *SpecificationSeederRefrigeratorSingerFTDS277ZWFBG) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds277zwf-bg"
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
		"Model Name":                  "SRREF-SS500-FTDS277ZWF-BG",
		"Door Type":                   "Top Mounted Refrigerator",
		"Capacity":                    "272 Liters",
		"Refrigerator Capacity":       "160 Liters",
		"Freezer Capacity":            "112 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1799 x 595 x 605 mm",
		"Weight":                      "60 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass Shelves",
		"Number of Shelves":           "4",
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
		"Gross Volume":                "272 Ltr.",
		"Net Volume":                  "272 Ltr.",
		"Special Features":            "NutriLock, Fresh-O-Logy, Modular Odour Filter, Modular Bottle Holder, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Keep Fresh Up To 20 Days, Food Grade Plastic Shelf & Crisper",
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
