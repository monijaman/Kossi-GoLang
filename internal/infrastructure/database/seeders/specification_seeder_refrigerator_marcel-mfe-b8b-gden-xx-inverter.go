package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter seeds specifications/options for product 'mfe-b8b-gden-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter() *SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-b8b-gden-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1590": "1590",
		"220-240V~/50Hz": "220-240V~/50Hz",
		"265 Ltr.": "265 Ltr.",
		"282 Ltr.": "282 Ltr.",
		"59 ± 2 Kg": "59 ± 2 Kg",
		"635": "635",
		"66 ± 2 Kg": "66 ± 2 Kg",
		"740": "740",
		"Manual": "Manual",
		"No": "No",
		"V 0501 - BLDC": "V 0501 - BLDC",
		"V 0501 - Mechanical": "V 0501 - Mechanical",
		"V 0501 - R600a": "V 0501 - R600a",
		"Wire/2": "Wire/2",
		"Yes (ABS/ PS)": "Yes (ABS/ PS)",
		"Yes (Plastic)": "Yes (Plastic)",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-b8b-gden-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeB8bGdenXxInverter) Seed(db *gorm.DB) error {
	productSlug := "mfe-b8b-gden-xx-inverter"
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
		"Compressor Type": 139,
		"Defrosting (Automatic/ Manual)": 141,
		"Depth/mm": 25,
		"Door Basket": 700,
		"Gross Volume": 709,
		"Gross Weight": 80,
		"Height/mm": 25,
		"Net Volume": 710,
		"Net Weight": 80,
		"Rated Operating Voltage and Frequency": 160,
		"Refrigerant": 708,
		"Reversible Door": 142,
		"Shelf (Material/ No.)": 699,
		"Temperature Control (Electronic/  Mechanical)": 158,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor Type": "V 0501 - BLDC",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "740",
		"Door Basket": "No",
		"Gross Volume": "282 Ltr.",
		"Gross Weight": "66 ± 2 Kg",
		"Height/mm": "1590",
		"Net Volume": "265 Ltr.",
		"Net Weight": "59 ± 2 Kg",
		"Rated Operating Voltage and Frequency": "220-240V~/50Hz",
		"Refrigerant": "V 0501 - R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "V 0501 - Mechanical",
		"Vegetable Crisper": "Yes (Plastic)",
		"Vegetable Crisper Cover": "Yes (ABS/ PS)",
		"Width/mm": "635",
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
