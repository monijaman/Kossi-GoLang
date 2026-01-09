package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter seeds specifications/options for product 'mfe-c5h-gden-dd-inverter'
type SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter() *SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-c5h-gden-dd-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1890": "1890",
		"220-240V~/50Hz": "220-240V~/50Hz",
		"345 Ltr.": "345 Ltr.",
		"358 Ltr.": "358 Ltr.",
		"635": "635",
		"740": "740",
		"76.5 ± Kg": "76.5 ± Kg",
		"84 ± Kg": "84 ± Kg",
		"Electronic": "Electronic",
		"Manual": "Manual",
		"No": "No",
		"Rack Evaporator": "Rack Evaporator",
		"V 0102- R600a V 0301- R600a V 0401- R600a": "V 0102- R600a V 0301- R600a V 0401- R600a",
		"V 0102- RSIR V 0301- RSIR V 0401- BLDC Inverter": "V 0102- RSIR V 0301- RSIR V 0401- BLDC Inverter",
		"Yes": "Yes",
		"Yes/1": "Yes/1",
		"Yes/2": "Yes/2",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-c5h-gden-dd-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdenDdInverter) Seed(db *gorm.DB) error {
	productSlug := "mfe-c5h-gden-dd-inverter"
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
		"Temperature Control (Electronic/  Mechanical)": 158,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor Type": "V 0102- RSIR V 0301- RSIR V 0401- BLDC Inverter",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "740",
		"Door Basket": "No",
		"Gross Volume": "358 Ltr.",
		"Gross Weight": "84 ± Kg",
		"Height/mm": "1890",
		"Ice Case": "Yes/2",
		"Ice Tray": "Yes/1",
		"Net Volume": "345 Ltr.",
		"Net Weight": "76.5 ± Kg",
		"Rated Voltage/ Hz": "220-240V~/50Hz",
		"Refrigerant": "V 0102- R600a V 0301- R600a V 0401- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Rack Evaporator",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
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
