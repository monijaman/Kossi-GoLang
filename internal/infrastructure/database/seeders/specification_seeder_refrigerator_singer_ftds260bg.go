package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS260BG seeds specifications/options for product 'singer-ftds260-bg'
type SpecificationSeederRefrigeratorSingerFTDS260BG struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS260BG creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS260BG() *SpecificationSeederRefrigeratorSingerFTDS260BG {
	return &SpecificationSeederRefrigeratorSingerFTDS260BG{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds260-bg"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS260BG) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                      "সিঙ্গার",
		"FBDS260-BG":                  "FBDS260-BG",
		"SRREF-SS300-FBDS260-BG":      "SRREF-SS300-FBDS260-BG",
		"Bottom Mounted Refrigerator": "বটম মাউন্টেড রেফ্রিজারেটর",
		"260 Liters":                  "260 লিটার",
		"130 Liters":                  "130 লিটার",
		"5 Star":                      "৫ তারা",
		"5":                           "৫",
		"250 kWh":                     "২৫০ কিলোওয়াট ঘণ্টা",
		"1650 x 560 x 500 mm":         "1650 x 560 x 500 মিমি",
		"55 kg":                       "55 kg",
		"Black":                       "ব্ল্যাক",
		"Inverter":                    "ইনভার্টার",
		"Direct Cool":                 "ডাইরেক্ট কুল",
		"Manual":                      "ম্যানুয়াল",
		"Mechanical":                  "মেকানিক্যাল",
		"Tempered Glass Shelves":      "টেম্পার্ড গ্লাস শেল্ফ",
		"3":                           "৩",
		"1":                           "১",
		"No":                          "না",
		"40 dB":                       "৪০ ডেসিবেল",
		"135V":                        "135ভোল্ট",
		"50":                          "৫০",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসর এবং 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"10":       "10",
		"R600a":    "আর600এ",
		"260 Ltr.": "260 লিটার",
		"Bottom Mounted Refrigerator, 50:50 Space Compartment, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Built-in Stabilizer, Odour Filter, Bottle Holder, Modular Odour Filter, Modular Bottle Holder, Big Freezer, Hidden Condenser, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Food Grade Plastic Shelf & Crisper": "বটম মাউন্টেড রেফ্রিজারেটর, 50:50 স্পেস কম্পার্টমেন্ট, বিএসটিআই দ্বারা 5 তারা এনার্জি রেটিং, কম পাওয়ার কনসাম্পশন, রিয়েল টেম্পার্ড গ্লাস ফিনিশ, টেম্পার্ড গ্লাস শেল্ফ, বিল্ট-ইন স্ট্যাবিলাইজার, ওডর ফিল্টার, বটল হোল্ডার, মডুলার ওডর ফিল্টার, মডুলার বটল হোল্ডার, বিগ ফ্রিজার, হিডেন কন্ডেন্সার, ফ্ল্যাশ এর্গোনমিক হ্যান্ডেল, আইস স্ক্রেপার, অ্যান্টি ব্যাকটেরিয়াল গ্যাসকেট, LED লাইট, ফুড গ্রেড প্লাস্টিক শেল্ফ এবং ক্রিস্পার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds260-bg'
func (s *SpecificationSeederRefrigeratorSingerFTDS260BG) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds260-bg"
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
		"Model Name":                  "SRREF-SS300-FBDS260-BG",
		"Door Type":                   "Bottom Mounted Refrigerator",
		"Capacity":                    "260 Liters",
		"Refrigerator Capacity":       "130 Liters",
		"Freezer Capacity":            "130 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1650 x 560 x 500 mm",
		"Weight":                      "55 kg",
		"Color":                       "Black",
		"Compressor Type":             "Inverter",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Tempered Glass Shelves",
		"Number of Shelves":           "3",
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
		"Gross Volume":                "260 Ltr.",
		"Net Volume":                  "260 Ltr.",
		"Special Features":            "Bottom Mounted Refrigerator, 50:50 Space Compartment, 5 Star Energy Rating by BSTI, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Built-in Stabilizer, Odour Filter, Bottle Holder, Modular Odour Filter, Modular Bottle Holder, Big Freezer, Hidden Condenser, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Food Grade Plastic Shelf & Crisper",
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
