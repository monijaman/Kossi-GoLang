package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx seeds specifications/options for product 'mfe-c3c-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-c3c-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"293 Ltr.": "293 Ltr.",
		"333 Ltr": "333 Ltr",
		"74 ± 2 Kg": "74 ± 2 Kg",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-c3c-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC3cGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "mfe-c3c-gdxx-xx"
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
		"Gross Volume": 709,
		"Net Volume": 710,
		"Net Weight": 80,
	}

	specs := map[string]string{
		"Gross Volume": "333 Ltr",
		"Net Volume": "293 Ltr.",
		"Net Weight": "74 ± 2 Kg",
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
