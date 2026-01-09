package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx seeds specifications/options for product 'mfc-c4h-gdne-xx'
type SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfcC4hGdneXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfcC4hGdneXx() *SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx {
	return &SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx{
		BaseSeeder: BaseSeeder{name: "Specifications for mfc-c4h-gdne-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1790": "1790",
		"220-240V~ and 50Hz": "220-240V~ and 50Hz",
		"333 Ltr.": "333 Ltr.",
		"348 Ltr": "348 Ltr",
		"5": "5",
		"66/ 73 ± 2": "66/ 73 ± 2",
		"710": "710",
		"Glass/2": "Glass/2",
		"Manual": "Manual",
		"No": "No",
		"V 0201- R600a V 0301- R600a V 0302- R600a": "V 0201- R600a V 0301- R600a V 0302- R600a",
		"V 0201- RSCR V 0301- RSIR V 0302- RSIR": "V 0201- RSCR V 0301- RSIR V 0302- RSIR",
		"V 0201/0301/0302: Wide Voltage Design (140V-260V) N.B.: If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 2100VA.": "V 0201/0301/0302: Wide Voltage Design (140V-260V) N.B.: If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 2100VA.",
		"Wire/2": "Wire/2",
		"Yes(Glass/ Plastic)": "Yes(Glass/ Plastic)",
		"Yes/1": "Yes/1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'mfc-c4h-gdne-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfcC4hGdneXx) Seed(db *gorm.DB) error {
	productSlug := "mfc-c4h-gdne-xx"
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
		"Operating voltage": 160,
		"Rated Operating Voltage and Frequency": 160,
		"Refrigerant": 708,
		"Reversible Door": 142,
		"Shelf (Material/ No.)": 699,
		"Shelf (Material/No.)": 699,
		"Star rating (BDS 1850:2012)": 144,
		"Vegetable Crisper": 701,
		"Vegetable Crisper Cover": 701,
		"Weight/Kg - Net/Packing": 80,
		"Width/mm": 25,
	}

	specs := map[string]string{
		"Compressor Type": "V 0201- RSCR V 0301- RSIR V 0302- RSIR",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "710",
		"Door Basket": "No",
		"Gross Volume": "348 Ltr",
		"Height/mm": "1790",
		"Net Volume": "333 Ltr.",
		"Operating voltage": "V 0201/0301/0302: Wide Voltage Design (140V-260V) N.B.: If out of voltage range(140V-260V), then suggested voltage stabilizer capacity is 2100VA.",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Refrigerant": "V 0201- R600a V 0301- R600a V 0302- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Glass/2",
		"Star rating (BDS 1850:2012)": "5",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes(Glass/ Plastic)",
		"Weight/Kg - Net/Packing": "66/ 73 ± 2",
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
