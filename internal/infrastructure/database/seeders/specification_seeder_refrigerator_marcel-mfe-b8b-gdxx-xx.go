package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx seeds specifications/options for product 'mfe-b8b-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-b8b-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1530": "1530",
		"220 ~ 240/ 50/ 109": "220 ~ 240/ 50/ 109",
		"265 Ltr.": "265 Ltr.",
		"282 Ltr.": "282 Ltr.",
		"54 ± 2 Kg": "54 ± 2 Kg",
		"625": "625",
		"745": "745",
		"Manual": "Manual",
		"Mechanical": "Mechanical",
		"No": "No",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"Wire/2": "Wire/2",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"Yes (Plastic)": "Yes (Plastic)",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-b8b-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "mfe-b8b-gdxx-xx"
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
		"Net Weight": 80,
		"Rated Voltage/ Hz/ watt": 160,
		"Refrigerant": 708,
		"Reversible Door": 142,
		"Shelf (Material/ No.)": 699,
		"Temperature Control": 158,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor": "RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "745",
		"Door Basket": "No",
		"Gross Volume": "282 Ltr.",
		"Height/mm": "1530",
		"Net Volume": "265 Ltr.",
		"Net Weight": "54 ± 2 Kg",
		"Rated Voltage/ Hz/ watt": "220 ~ 240/ 50/ 109",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control": "Mechanical",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
		"Width/mm": "625",
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
