package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcfA4eGdelLx seeds specifications/options for product 'marcel-mcf-a4e-gdel-lx'
type SpecificationSeederRefrigeratorMarcelMcfA4eGdelLx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcfA4eGdelLx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcfA4eGdelLx() *SpecificationSeederRefrigeratorMarcelMcfA4eGdelLx {
	return &SpecificationSeederRefrigeratorMarcelMcfA4eGdelLx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcf-a4e-gdel-lx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcfA4eGdelLx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                         "মার্সেল",
		"marcel-mcf-a4e-gdel-lx":         "মার্সেল-mcf-a4e-gdel-lx",
		"MCF-A4E-GDEL-LX":                "MCF-A4E-GDEL-LX",
		"Single Door":                    "সিঙ্গেল ডোর",
		"145 Liters":                     "১৪৫ লিটার",
		"0 Liters":                       "০ লিটার",
		"744 x 585 x 853 mm (W x D x H)": "৭৪৪ x ৫৮৫ x ৮৫৩ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"37 kg":                          "৩৭ কেজি",
		"Silver":                         "সিলভার",
		"RSCR":                           "RSCR",
		"Direct Cool":                    "ডাইরেক্ট কুল",
		"Manual":                         "ম্যানুয়াল",
		"Mechanical":                     "মেকানিক্যাল",
		"Wire":                           "ওয়্যার",
		"1":                              "১",
		"No":                             "না",
		"40 dB":                          "৪০ ডেসিবেল",
		"220-240V":                       "২২০-২৪০ ভোল্ট",
		"50":                             "৫০",
		"2 Years":                        "২ বছর",
		"5":                              "৫",
		"R600a":                          "R600a",
		"Single Door Freezer, Manual Defrost, Interior Lamp, Reversible Door": "সিঙ্গেল ডোর ফ্রিজার, ম্যানুয়াল ডিফ্রস্ট, ইন্টেরিয়র ল্যাম্প, রিভার্সিবল ডোর",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcf-a4e-gdel-lx'
func (s *SpecificationSeederRefrigeratorMarcelMcfA4eGdelLx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcf-a4e-gdel-lx"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
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
		"Brand":                       "Marcel",
		"Model Name":                  "MCF-A4E-GDEL-LX",
		"Door Type":                   "Single Door",
		"Capacity":                    "145 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "145 Liters",
		"Gross Volume":                "145 Liters",
		"Net Volume":                  "145 Liters",
		"Dimensions":                  "744 x 585 x 853 mm (W x D x H)",
		"Weight":                      "37 kg",
		"Color":                       "Silver",
		"Compressor Type":             "RSCR",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "1",
		"Door Bins":                   "0",
		"Crisper Drawers":             "0",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "220-240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "5",
		"Refrigerant":                 "R600a",
		"Special Features":            "Single Door Freezer, Manual Defrost, Interior Lamp, Reversible Door",
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
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Ensure the ID is set after creation
				if sModel.ID == 0 {
					if err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {
						return err
					}
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
