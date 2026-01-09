package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hElnxXx seeds specifications/options for product 'mfe-c5h-elnx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC5hElnxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hElnxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hElnxXx() *SpecificationSeederRefrigeratorMarcelMfeC5hElnxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hElnxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfe-c5h-elnx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hElnxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1830": "1830",
		"220~240/ 50/135": "220~240/ 50/135",
		"345 Ltr.": "345 Ltr.",
		"358 Ltr.": "358 Ltr.",
		"625": "625",
		"68.5 ± 2 Kg": "68.5 ± 2 Kg",
		"745": "745",
		"Manual": "Manual",
		"No": "No",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"Rack Evaporator": "Rack Evaporator",
		"Yes": "Yes",
		"Yes/1": "Yes/1",
		"Yes/2": "Yes/2",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfe-c5h-elnx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hElnxXx) Seed(db *gorm.DB) error {
	productSlug := "mfe-c5h-elnx-xx"
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
		"Ice Case": 702,
		"Ice Tray": 702,
		"Net Volume": 710,
		"Net Weight": 80,
		"Rated Voltage/ Hz/ watt": 160,
		"Refrigerant": 708,
		"Reversible Door": 142,
		"Shelf (Material/ No.)": 699,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor": "RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "745",
		"Door Basket": "No",
		"Gross Volume": "358 Ltr.",
		"Height/mm": "1830",
		"Ice Case": "Yes/2",
		"Ice Tray": "Yes/1",
		"Net Volume": "345 Ltr.",
		"Net Weight": "68.5 ± 2 Kg",
		"Rated Voltage/ Hz/ watt": "220~240/ 50/135",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Rack Evaporator",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
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
