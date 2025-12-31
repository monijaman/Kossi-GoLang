package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMnmB1gGdelXx seeds specifications/options for product 'marcel-mnm-b1g-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMnmB1gGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMnmB1gGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMnmB1gGdelXx() *SpecificationSeederRefrigeratorMarcelMnmB1gGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMnmB1gGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mnm-b1g-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMnmB1gGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mnm-b1g-gdel-xx": "মার্সেল-mnm-b1g-gdel-xx",
		"MNM-B1G-GDEL-XX":        "MNM-B1G-GDEL-XX",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mnm-b1g-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMnmB1gGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mnm-b1g-gdel-xx"
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
		"Cooling Features - Type":                             "Direct Cool",
		"Capacity - Gross Volume":                             "380 Ltr.",
		"Capacity - Net Volume":                               "365 Ltr.",
		"Performance - Climatic Type (SN":                     "N",
		"Performance - Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Compressor Input Power":                              "V 0401- 130; V 0601- 130; V 0701- 130; V 0702- 130",
		"Compressor":                                          "RSCR",
		"Cooling Effect":                                      "Freezer Cabinet Less than -18℃; Refrigerator Cabinet 0℃ to +5℃",
		"General Features - Temperature Control":              "Mechanical",
		"General Features - Defrosting":                       "Manual",
		"General Features - Reversible Door":                  "No",
		"General Features - Handle":                           "Recessed/ Grip",
		"General Features - Lock":                             "Yes",
		"Refrigerant":                                         "V 0401- R600a; V 0601- R600a; V 0701- R600a; V 0702- R600a",
		"Capillary":                                           "Copper",
		"Polyurethane foam blowing agent":                     "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity":             "V 0101/ 0201/ 0301/ 0401/ 0601/ 0701/ 0702: 2200VA",
		"Refrigerator Compartment - Shelf (Material/ No.)":    "Wire/3",
		"Refrigerator Compartment - Door Basket":              "4",
		"Refrigerator Compartment - Interior Lamp":            "Yes",
		"Refrigerator Compartment - Vegetable Crisper":        "Yes/1",
		"Refrigerator Compartment - Vegetable Crisper Cover":  "Yes/1",
		"Refrigerator Compartment - Egg Tray":                 "Yes/2",
		"Refrigerator Compartment - Can Storage Dispenser":    "No",
		"Freezer Compartment - Shelf (Material/ No.)":         "Wire/2",
		"Freezer Compartment - Drawer":                        "No",
		"Freezer Compartment - Door Basket":                   "No",
		"Freezer Compartment - Interior Lamp":                 "No",
		"Dimensions (Net) - Width/mm":                         "645",
		"Dimensions (Net) - Depth/mm":                         "645",
		"Dimensions (Net) - Height/mm":                        "1860",
		"Packing - Width/mm":                                  "710",
		"Packing - Depth/mm":                                  "710",
		"Packing - Height/mm":                                 "1910",
		"Weight/Kg - Net/Packing":                             "67/76 ± 2",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":                  "66/ 48/ 24",
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
