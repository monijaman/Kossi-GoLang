package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerFTDS277RG seeds specifications/options for product 'singer-ftds277-rg'
type SpecificationSeederRefrigeratorSingerFTDS277RG struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerFTDS277RG creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerFTDS277RG() *SpecificationSeederRefrigeratorSingerFTDS277RG {
	return &SpecificationSeederRefrigeratorSingerFTDS277RG{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ftds277-rg"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerFTDS277RG) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                    "সিঙ্গার",
		"FTDS277-RG":                "FTDS277-RG",
		"SRREF-SS500-FTDS277ZWF-RG": "SRREF-SS500-FTDS277ZWF-RG",
		"Top Mounted Refrigerator":  "টপ মাউন্টেড রেফ্রিজারেটর",
		"272 Liters":                "272 লিটার",
		"5 Star":                    "৫ তারা",
		"5":                         "৫",
		"250 kWh":                   "২৫০ কিলোওয়াট ঘণ্টা",
		"1799 x 595 x 605 mm":       "1799 x 595 x 605 মিমি",
		"60 kg":                     "60 kg",
		"Red":                       "রেড",
		"Inverter":                  "ইনভার্টার",
		"Direct Cool":               "ডাইরেক্ট কুল",
		"Manual":                    "ম্যানুয়াল",
		"Mechanical":                "মেকানিক্যাল",
		"Tempered Glass Shelves":    "টেম্পার্ড গ্লাস শেল্ফ",
		"4":                         "৪",
		"3":                         "৩",
		"1":                         "১",
		"Yes":                       "হ্যাঁ",
		"No":                        "না",
		"40 dB":                     "৪০ ডেসিবেল",
		"135V":                      "135ভোল্ট",
		"50":                        "৫০",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসর এবং 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"Top Mounted Refrigerator, 45:55 Space Compartment Ratio, 5 Star Energy Rating by BSTI, PowerCool Technology, Real Tempered Glass Finish, Tempered Glass Shelves, NutriLock Technology to Keep Vitamin A & C Longer of Fruits and Vegetable, Fresh-O-Logy Technology to Keep Fruits and Vegitable Fresh upto 20 Days, Built-in Stabilizer, Odour Filter, Bottle Holder, Modular Odour Filter, Modular Bottle Holder, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Keep Fresh Up To 20 Days, Food Grade Plastic Shelf & Crisper": "টপ মাউন্টেড রেফ্রিজারেটর, 45:55 স্পেস কম্পার্টমেন্ট রেশিও, বিএসটিআই দ্বারা 5 তারা এনার্জি রেটিং, পাওয়ারকুল টেকনোলজি, রিয়েল টেম্পার্ড গ্লাস ফিনিশ, টেম্পার্ড গ্লাস শেল্ফ, নুট্রিলক টেকনোলজি ফল এবং শাকসবজিতে ভিটামিন A & C দীর্ঘতর রাখতে, ফ্রেশ-ও-লজি টেকনোলজি ফল এবং শাকসবজি 20 দিন পর্যন্ত তাজা রাখতে, বিল্ট-ইন স্ট্যাবিলাইজার, ওডর ফিল্টার, বটল হোল্ডার, মডুলার ওডর ফিল্টার, মডুলার বটল হোল্ডার, হিডেন কন্ডেন্সার, বিগ ফ্রিজার, ফ্ল্যাশ এর্গোনমিক হ্যান্ডেল, আইস স্ক্রেপার, অ্যান্টি ব্যাকটেরিয়াল গ্যাসকেট, LED লাইট, 20 দিন পর্যন্ত তাজা রাখুন, ফুড গ্রেড প্লাস্টিক শেল্ফ এবং ক্রিস্পার",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"277 Ltr.":         "277 লিটার",
		"R600a":            "আর600এ",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ftds277-rg'
func (s *SpecificationSeederRefrigeratorSingerFTDS277RG) Seed(db *gorm.DB) error {
	productSlug := "singer-ftds277-rg"
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
		"Model Name":                  "SRREF-SS500-FTDS277ZWF-RG",
		"Door Type":                   "Top Mounted Refrigerator",
		"Capacity":                    "272 Liters",
		"Refrigerator Capacity":       "160 Liters",
		"Freezer Capacity":            "112 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "250 kWh",
		"Dimensions":                  "1799 x 595 x 605 mm",
		"Weight":                      "60 kg",
		"Color":                       "Red",
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
		"Special Features":            "Top Mounted Refrigerator, 45:55 Space Compartment Ratio, 5 Star Energy Rating by BSTI, PowerCool Technology, Real Tempered Glass Finish, Tempered Glass Shelves, NutriLock Technology to Keep Vitamin A & C Longer of Fruits and Vegetable, Fresh-O-Logy Technology to Keep Fruits and Vegitable Fresh upto 20 Days, Built-in Stabilizer, Odour Filter, Bottle Holder, Modular Odour Filter, Modular Bottle Holder, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Ice Scraper, Anti Bacterial Gasket, LED Light, Keep Fresh Up To 20 Days, Food Grade Plastic Shelf & Crisper",
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
