package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter seeds specifications/options for product 'mfb-b5x-gdel-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter() *SpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfb-b5x-gdel-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1725": "1725",
		"220 ~ 240/ 50": "220 ~ 240/ 50",
		"244 Ltr.": "244 Ltr.",
		"250 Ltr.": "250 Ltr.",
		"50.5/55.5 ± 2": "50.5/55.5 ± 2",
		"54.5 ± 2 Kg": "54.5 ± 2 Kg",
		"580": "580",
		"59.5 ± 2 Kg": "59.5 ± 2 Kg",
		"645": "645",
		"BLDC Inverter": "BLDC Inverter",
		"GPPS/3": "GPPS/3",
		"Manual": "Manual",
		"Mechanical": "Mechanical",
		"R600a": "R600a",
		"Wire/2": "Wire/2",
		"Yes": "Yes",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5x-gdel-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdelXxInverter) Seed(db *gorm.DB) error {
	productSlug := "mfb-b5x-gdel-xx-inverter"
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
		"Defrosting (Automatic/ Manual):": 141,
		"Depth/mm": 25,
		"Door Basket": 700,
		"Gross Volume": 709,
		"Gross Weight": 80,
		"Height/mm": 25,
		"Net Volume": 710,
		"Net Weight": 80,
		"Rated Voltage/ Hz": 160,
		"Refrigerant": 708,
		"Shelf (Material/ No.)": 699,
		"Shelf (Material/No.)": 699,
		"Temperature Control (Electronic/  Mechanical)": 158,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Weight/Kg - Net/Packing": 80,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor Type": "BLDC Inverter",
		"Defrosting (Automatic/ Manual):": "Manual",
		"Depth/mm": "645",
		"Door Basket": "GPPS/3",
		"Gross Volume": "250 Ltr.",
		"Gross Weight": "59.5 ± 2 Kg",
		"Height/mm": "1725",
		"Net Volume": "244 Ltr.",
		"Net Weight": "54.5 ± 2 Kg",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Refrigerant": "R600a",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Weight/Kg - Net/Packing": "50.5/55.5 ± 2",
		"Width/mm": "580",
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
