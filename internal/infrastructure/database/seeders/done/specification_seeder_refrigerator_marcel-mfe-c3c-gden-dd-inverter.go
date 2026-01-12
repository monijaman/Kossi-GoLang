package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter seeds specifications/options for product 'mfe-c3c-gden-dd-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter() *SpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-c3c-gden-dd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1890": "1890",
		"220- 240/ 50": "220- 240/ 50",
		"293 Ltr.": "293 Ltr.",
		"333 Ltr": "333 Ltr",
		"635": "635",
		"74 ± 2 Kg": "74 ± 2 Kg",
		"740": "740",
		"80 ± 2 Kg": "80 ± 2 Kg",
		"Electronic": "Electronic",
		"Manual": "Manual",
		"No": "No",
		"R600a": "R600a",
		"Rack Evaporator": "Rack Evaporator",
		"V 0501- BLDC Inverter": "V 0501- BLDC Inverter",
		"Wire/2": "Wire/2",
		"Yes/1": "Yes/1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-c3c-gden-dd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC3cGdenDdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfe-c3c-gden-dd-inverter"
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
		"Ice Case": 702,
		"Ice Tray": 702,
		"Net Volume": 710,
		"Net Weight": 80,
		"Rated Voltage/ Hz": 160,
		"Refrigerant": 708,
		"Reversible Door": 142,
		"Shelf (Material/ No.)": 699,
		"Shelf (Material/No.)": 699,
		"Temperature Control (Electronic/  Mechanical)": 158,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor Type": "V 0501- BLDC Inverter",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "740",
		"Door Basket": "No",
		"Gross Volume": "333 Ltr",
		"Gross Weight": "80 ± 2 Kg",
		"Height/mm": "1890",
		"Ice Case": "Yes/1",
		"Ice Tray": "Yes/1",
		"Net Volume": "293 Ltr.",
		"Net Weight": "74 ± 2 Kg",
		"Rated Voltage/ Hz": "220- 240/ 50",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Rack Evaporator",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes/1",
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
