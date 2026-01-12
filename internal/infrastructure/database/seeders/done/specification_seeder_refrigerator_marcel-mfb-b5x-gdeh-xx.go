package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdehXx seeds specifications/options for product 'mfb-b5x-gdeh-xx'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdehXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdehXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdehXx() *SpecificationSeederRefrigeratorMarcelMfbB5xGdehXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdehXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfb-b5x-gdeh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdehXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1725 mm": "1725 mm",
		"220 ~ 240/ 50": "220 ~ 240/ 50",
		"250 Ltr.": "250 Ltr.",
		"254 Ltr.": "254 Ltr.",
		"54.5 ± 2 Kg": "54.5 ± 2 Kg",
		"580 mm": "580 mm",
		"645 mm": "645 mm",
		"GPPS/3": "GPPS/3",
		"Manual": "Manual",
		"RSCR": "RSCR",
		"V 01.01-R600a V 02.01-R600a V 03.01-R600a V 03.02-R600a": "V 01.01-R600a V 02.01-R600a V 03.01-R600a V 03.02-R600a",
		"Wire/2": "Wire/2",
		"Yes": "Yes",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfb-b5x-gdeh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdehXx) Seed(db *gorm.DB) error {
	productSlug := "mfb-b5x-gdeh-xx"
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
		"Height/mm": 25,
		"Net Volume": 710,
		"Net Weight": 80,
		"Rated Voltage/ Hz": 160,
		"Refrigerant": 708,
		"Shelf (Material/ No.)": 699,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor Type": "RSCR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "645 mm",
		"Door Basket": "GPPS/3",
		"Gross Volume": "250 Ltr.",
		"Height/mm": "1725 mm",
		"Net Volume": "254 Ltr.",
		"Net Weight": "54.5 ± 2 Kg",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Refrigerant": "V 01.01-R600a V 02.01-R600a V 03.01-R600a V 03.02-R600a",
		"Shelf (Material/ No.)": "Wire/2",
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
