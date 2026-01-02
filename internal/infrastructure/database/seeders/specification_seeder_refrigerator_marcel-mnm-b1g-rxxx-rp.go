package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp seeds specifications/options for product 'marcel-mnm-b1g-rxxx-rp'
type SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp() *SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp {
	return &SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mnm-b1g-rxxx-rp"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mnm-b1g-rxxx-rp": "মার্সেল-mnm-b1g-rxxx-rp",
		"MNM-B1G-RXXX-RP":        "MNM-B1G-RXXX-RP",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mnm-b1g-rxxx-rp'
func (s *SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp) Seed(db *gorm.DB) error {
	productSlug := "marcel-mnm-b1g-rxxx-rp"
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
		"Brand":                           "Marcel",
		"Model Name":                      "MNM-B1G-RXXX-RP",
		"Cooling Technology":              "Direct Cool",
		"Gross Volume":                    "380 Ltr.",
		"Net Volume":                      "365 Ltr.",
		"Compressor Type":                 "RSCR; V 0401-130; V 0601-130; V 0701-130; V 0702-130",
		"Temperature Control":             "Mechanical",
		"Defrost Type":                    "Manual",
		"Refrigerant":                     "R600a",
		"Capillary":                       "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "2200VA",
		"Shelf Material":     "Wire",
		"Number of Shelves":  "3",
		"Door Bins":          "4",
		"Crisper Drawers":    "Yes/1",
		"Dimensions":         "645 x 645 x 1860 mm",
		"Packing Dimensions": "710 x 710 x 1910 mm",
		"Weight":             "67/76 ± 2",
		"Loading Capacity":   "66/ 48/ 24",
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
