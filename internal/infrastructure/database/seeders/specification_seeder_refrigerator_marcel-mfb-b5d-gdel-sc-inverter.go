package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter seeds specifications/options for product 'mfb-b5d-gdel-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter() *SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for mfb-b5d-gdel-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1845 mm": "1845 mm",
		"220 ~ 240/ 50": "220 ~ 240/ 50",
		"254 Ltr.": "254 Ltr.",
		"268 Ltr.": "268 Ltr.",
		"5 star (BDS 1850:2012)": "5 star (BDS 1850:2012)",
		"580 mm": "580 mm",
		"62 ± 2 Kg": "62 ± 2 Kg",
		"645 mm": "645 mm",
		"67 ± 2 Kg": "67 ± 2 Kg",
		"Electronic": "Electronic",
		"Manual": "Manual",
		"No": "No",
		"R600a": "R600a",
		"V 05.03-BLDC": "V 05.03-BLDC",
		"Wire/3": "Wire/3",
		"Yes": "Yes",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5d-gdel-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter) Seed(db *gorm.DB) error {
	productSlug := "mfb-b5d-gdel-sc-inverter"
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
		"Door Pocket": 700,
		"Energy Rating": 143,
		"Gross Volume": 709,
		"Gross Weight": 80,
		"Height/mm": 25,
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
		"Compressor Type": "V 05.03-BLDC",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "645 mm",
		"Door Pocket": "No",
		"Energy Rating": "5 star (BDS 1850:2012)",
		"Gross Volume": "268 Ltr.",
		"Gross Weight": "67 ± 2 Kg",
		"Height/mm": "1845 mm",
		"Net Volume": "254 Ltr.",
		"Net Weight": "62 ± 2 Kg",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/3",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Width/mm": "580 mm",
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
