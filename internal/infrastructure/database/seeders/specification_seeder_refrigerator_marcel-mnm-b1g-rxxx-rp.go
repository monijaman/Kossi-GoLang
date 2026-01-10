package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp seeds specifications/options for product 'marcel-mnm-b1g-rxxx-rp'
type SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp() *SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp {
	return &SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mnm-b1g-rxxx-rp"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mnm-b1g-rxxx-rp": "মার্সেল-mnm-b1g-rxxx-rp",
		"MNM-B1G-RXXX-RP":        "MNM-B1G-RXXX-RP",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mnm-b1g-rxxx-rp'
func (s *SpecificationSeederRefrigeratorMarcelMnmB1gRxxxRp) Seed(db *gorm.DB) error {
	productSlug := "marcel-mnm-b1g-rxxx-rp"
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
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}
	specs := map[string]string{
		"Bottle Pocket": "GPPS/2",
		"Capillary": "Copper",
		"Chilled Room": "Yes/1",
		"Chilled Room Door": "Yes/1",
		"Climatic Type (SN, N, ST, T)": "T",
		"Compressor": "RSCR",
		"Condenser": "100 % Copper",
		"Cooling Efect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Depth/mm": "645 mm",
		"Egg Tray or Pocket": "Yes/ 1",
		"Gross Volume": "217 Ltr.",
		"Gross Weight": "59.15± 2 Kg",
		"Height/mm": "1500 mm",
		"Ice Box": "Yes/1",
		"Ice tray": "Yes/1",
		"Interior LED Lamp": "N/A",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "83/ 83/ 39",
		"Net Volume": "183 Ltr.",
		"Net Weight": "53.15± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz/ watt": "220-240/ 50 Hz/ 88W",
		"Refrigerant": "R600a",
		"Shelf (Material/ No.)": "GPPS/2",
		"Shelf (Material/No.)": "GPPS/1, Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Non-Frost",
		"Vegetable Crisper": "Yes/1",
		"Width/mm": "550 mm",
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
