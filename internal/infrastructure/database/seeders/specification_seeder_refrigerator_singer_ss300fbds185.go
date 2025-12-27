package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSingerSS300FBDS185 seeds specifications/options for product 'singer-ss300-fbds185'
type SpecificationSeederRefrigeratorSingerSS300FBDS185 struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSingerSS300FBDS185 creates a new seeder instance
func NewSpecificationSeederRefrigeratorSingerSS300FBDS185() *SpecificationSeederRefrigeratorSingerSS300FBDS185 {
	return &SpecificationSeederRefrigeratorSingerSS300FBDS185{
		BaseSeeder: BaseSeeder{name: "Specifications for singer-ss300-fbds185"},
	}
}

func (s *SpecificationSeederRefrigeratorSingerSS300FBDS185) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer":                      "সিঙ্গার",
		"SS300-FBDS185":               "SS300-FBDS185",
		"Bottom Mounted Refrigerator": "বটম মাউন্টেড রেফ্রিজারেটর",
		"183.7 Liters":                "183.7 লিটার",
		"77.7 Liters":                 "77.7 লিটার",
		"106 Liters":                  "106 লিটার",
		"5 Star":                      "৫ তারা",
		"5":                           "৫",
		"1580 x 545 x 550 mm":         "1580 x 545 x 550 মিমি",
		"Black":                       "ব্ল্যাক",
		"Tempered Glass":              "টেম্পার্ড গ্লাস",
		"3":                           "৩",
		"1":                           "১",
		"No":                          "না",
		"135 ~ 254V":                  "135 ~ 254ভোল্ট",
		"50":                          "৫০",
		"10 Years Compressor & 2 Years Spare Parts Warranty": "10 বছর কম্প্রেসার এবং 2 বছর স্পেয়ার পার্টস ওয়ারেন্টি",
		"10":               "10",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Net Volume":       "নেট ভলিউম",
		"183.7 Ltr.":       "183.7 লিটার",
		"R600a":            "আর600এ",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
	}
}

// Seed inserts specification records for the product identified by slug 'singer-ss300-fbds185'
func (s *SpecificationSeederRefrigeratorSingerSS300FBDS185) Seed(db *gorm.DB) error {
	productSlug := "singer-ss300-fbds185"
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
		"Model Name":                  "SS300-FBDS185",
		"Door Type":                   "Bottom Mounted Refrigerator",
		"Capacity":                    "183.7 Liters",
		"Refrigerator Capacity":       "77.7 Liters",
		"Freezer Capacity":            "106 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "",
		"Dimensions":                  "1580 x 545 x 550 mm",
		"Weight":                      "",
		"Color":                       "Black",
		"Compressor Type":             "",
		"Cooling Technology":          "",
		"Defrost Type":                "",
		"Temperature Control":         "",
		"Shelf Material":              "Tempered Glass",
		"Number of Shelves":           "3",
		"Door Bins":                   "3",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "",
		"Voltage":                     "135 ~ 254V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "10 Years Compressor & 2 Years Spare Parts Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "",
		"Net Volume":                  "183.7 Ltr.",
		"Special Features":            "45:55 Space Compartment Ratio, Low power consumption, Real Tempered Glass Finish, Tempered Glass Shelves, Built-in Stabilizer, Odor Filter, Bottom Mounted Refrigerator, Door Lock, Low Voltage Compressor LVS, Base Drawer, Glass Shelves, Modular Odour Filter, Hidden Condenser, Big Freezer, Flash Ergonomic Handle, Anti Bacterial Gasket, LED Light, Food Grade Plastic Shelf & Crisper, Ice Scraper, Modular Bottle Holder",
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
