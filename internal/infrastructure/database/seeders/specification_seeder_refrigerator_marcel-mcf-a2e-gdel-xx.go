package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcfA2eGdelXx seeds specifications/options for product 'marcel-mcf-a2e-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMcfA2eGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcfA2eGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcfA2eGdelXx() *SpecificationSeederRefrigeratorMarcelMcfA2eGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMcfA2eGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcf-a2e-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcfA2eGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                         "মার্সেল",
		"marcel-mcf-a2e-gdel-xx":         "মার্সেল-mcf-a2e-gdel-xx",
		"MCF-A2E-GDEL-XX":                "MCF-A2E-GDEL-XX",
		"Single Door":                    "সিঙ্গেল ডোর",
		"125 Liters":                     "১২৫ লিটার",
		"0 Liters":                       "০ লিটার",
		"695 x 600 x 855 mm (W x D x H)": "৬৯৫ x ৬০০ x ৮৫৫ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"33 kg":                          "৩৩ কেজি",
		"Silver":                         "সিলভার",
		"RSIR/RSCR":                      "RSIR/RSCR",
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
		"R134a/R600a":                    "R134a/R600a",
		"Single Door Freezer, Manual Defrost, Lock, Interior Lamp": "সিঙ্গেল ডোর ফ্রিজার, ম্যানুয়াল ডিফ্রস্ট, লক, ইন্টেরিয়র ল্যাম্প",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcf-a2e-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMcfA2eGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcf-a2e-gdel-xx"
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
		"Model Name":                  "MCF-A2E-GDEL-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "125 Liters",
		"Refrigerator Capacity":       "0 Liters",
		"Freezer Capacity":            "125 Liters",
		"Gross Volume":                "125 Liters",
		"Net Volume":                  "125 Liters",
		"Dimensions":                  "695 x 600 x 855 mm (W x D x H)",
		"Weight":                      "33 kg",
		"Color":                       "Silver",
		"Compressor Type":             "RSIR/RSCR",
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
		"Refrigerant":                 "R134a/R600a",
		"Special Features":            "Single Door Freezer, Manual Defrost, Lock, Interior Lamp",
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
