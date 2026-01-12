package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC6eGdelXx seeds specifications/options for product 'mfc-c6e-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfcC6eGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC6eGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC6eGdelXx() *SpecificationSeederRefrigeratorMarcelMfcC6eGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfcC6eGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c6e-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC6eGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1910": "1910",
		"220-240V~ and 50Hz": "220-240V~ and 50Hz",
		"365 Ltr.": "365 Ltr.",
		"380 Ltr.": "380 Ltr.",
		"67/76 ± 2": "67/76 ± 2",
		"710": "710",
		"Manual": "Manual",
		"No": "No",
		"RSCR": "RSCR",
		"V 0401- R600a V 0601- R600a V 0701- R600a V 0702- R600a": "V 0401- R600a V 0601- R600a V 0701- R600a V 0702- R600a",
		"Wire/2": "Wire/2",
		"Yes/1": "Yes/1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfc-c6e-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfcC6eGdelXx) Seed(db *gorm.DB) error {
	productSlug := "mfc-c6e-gdel-xx"
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
		"Compressor": 139,
		"Defrosting (Automatic/ Manual)": 141,
		"Depth/mm": 25,
		"Door Basket": 700,
		"Gross Volume": 709,
		"Height/mm": 25,
		"Net Volume": 710,
		"Rated Operating Voltage and Frequency": 160,
		"Refrigerant": 708,
		"Reversible Door": 142,
		"Shelf (Material/ No.)": 699,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Weight/Kg - Net/Packing:": 80,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor": "RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "710",
		"Door Basket": "No",
		"Gross Volume": "380 Ltr.",
		"Height/mm": "1910",
		"Net Volume": "365 Ltr.",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Refrigerant": "V 0401- R600a V 0601- R600a V 0701- R600a V 0702- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes/1",
		"Weight/Kg - Net/Packing:": "67/76 ± 2",
		"Width/mm": "710",
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
