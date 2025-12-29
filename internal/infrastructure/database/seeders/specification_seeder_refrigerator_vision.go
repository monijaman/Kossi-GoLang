package seeders

import (
	"log"
	"regexp"
	"strings"

	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVision seeds specifications/options for Vision models
type SpecificationSeederRefrigeratorVision struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVision creates a new seeder instance
func NewSpecificationSeederRefrigeratorVision() *SpecificationSeederRefrigeratorVision {
	return &SpecificationSeederRefrigeratorVision{
		BaseSeeder: BaseSeeder{name: "Specifications for Vision refrigerator models"},
	}
}

func (s *SpecificationSeederRefrigeratorVision) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
	}
}

// helper to create a product slug similar to existing seeders
func (s *SpecificationSeederRefrigeratorVision) slugify(input string) string {
	in := strings.ToLower(strings.TrimSpace(input))
	in = strings.ReplaceAll(in, "+", "plus")
	// replace any non-alphanumeric with hyphen
	re := regexp.MustCompile(`[^a-z0-9]+`)
	out := re.ReplaceAllString(in, "-")
	out = strings.Trim(out, "-")
	// collapse multiple hyphens
	out = regexp.MustCompile(`-+`).ReplaceAllString(out, "-")
	return "vision-" + out
}

// Seed inserts specification records for Vision product models
func (s *SpecificationSeederRefrigeratorVision) Seed(db *gorm.DB) error {
	brand := "Vision"
	modelsList := []string{
		"VISION Mini Refrigerator RE-50L SS",
		"VISION Mini Refrigerator RE-101 Liter SS",
		"VISION Glass Door Refrigerator RE-142 Liter",
		"VISION Glass Door Chest Freeze RE-150 Liter",
		"VISION Glass Door Chest Freezer RE-150 Liter",
		"VISION Glass Door Refrigerator RE-150 Liter",
		"VISION Glass Door Chest Freezer RE-150 Liter",
		"VISION Glass Door Refrigerator RE-160 Liter",
		"VISION Glass Door Refrigerator RE-185 Liter",
		"VISION Glass Door Refrigerator RE-180 Liter",
		"VISION Glass Door Bottom Mount Refrigerator",
		"VISION Glass Door Refrigerator RE-200 Liter",
		"VISION Glass Door Refrigerator RE-216 Liter",
		"VISION Glass Door Refrigerator RE-222 Liter",
		"VSN GD Refrigerator RE-216L Mirror Purple",
		"VISION Glass Door Top Mount Refrigerator",
		"VISION Glass Door Refrigerator RE-208 Liter",
		"VISION Glass Door Chest Freezer RE-250 Liter",
		"VISION Glass Door Refrigerator RE-240 Liter",
		"VISION Glass Door Refrigerator RE-238 Liter",
		"VSN GD Refrigerator RE-240L Mirror Purple",
		"VISION Glass Door Smart Dispenser Refrigerator",
		"VISION Glass Door Refrigerator RE-252L Mirror",
		"VISION Glass Door Refrigerator RE-280 Liter",
		"VISION Glass Door Chest Freezer RE-350 Liter",
		"VISION Glass Door Refrigerator RE-305 Liter",
		"VSN GD Refrigerator RE-305L Mirror Jaba Flower",
		"VISION Glass Door Refrigerator RE-330 Liter",
		"VISION Glass Door Refrigerator RE-356 Liter",
		"VSN GD Refrigerator RE-330L WD Black-BM",
		"VISION Glass Door Refrigerator RE-285 Liter",
		"VISION Ice Cream Freezer 368L",
		"VISION Glass Door Refrigerator RE-309 Liter",
		"VISION Glass Door Refrigerator Side By Side",
		"VSN GD Refrigerator French Door Inverter",
		"VISION Refrigerator Vis-50 Ltr Gray",
		"VSN Mini Refrigerator RE-101L SS",
		"VSN GD Refrigerator RE-142L Digital Rainbow",
		"VISION Glass Door Refrigerator RE-142L Digital",
		"VISION Glass Door Refrigerator RE-150L Digital",
		"VSN GD Chest Freezer VIS-150L Magic Line",
		"VSN GD Chest Freezer RE-150L Maple Leaf",
		"VISION Glass Door Refrigerator Top Mount",
		"VISION Glass Door Bottom Mount Refrigerator",
		"VSN GD Refrigerator RE-150L Digital Red FL",
		"VSN GD Refrigerator RE-160L Red FL-BM",
		"Vision Glass Door Refrigerator RE-160L Blue",
		"VISION Glass Door Refrigerator RE-180L Digital",
		"VISION Glass Door Refrigerator RE-180L Pink",
		"VISION Glass Door Refrigerator RE-180L Platinum",
		"VISION Glass Door Refrigerator RE-185L Digital",
		"VSN GD Refrigerator RE-185L Pink FL-BM",
		"VISION Glass Door Refrigerator RE-196L Pink",
		"VISION Glass Door Refrigerator RE-200L Red",
		"VSN GD Refrigerator RE-200L Digital Dark",
		"VISION Glass Door Refrigerator RE-216L Blue",
		"VSN GD Refrigerator RE-217L Blue Peony",
		"VISION Glass Door Refrigerator RE-222L Mirror",
		"VSN GD Refrigerator RE-222L Lotus Fl-Maroon",
		"VISION Glass Door Chest Freezer RE-250L Maple",
		"VISION Ice Cream Freezer 158L",
		"Vision Refrigerator 267 L Inverter",
		"VISION Beverage Refrigerator RE-275 Litre",
		"VISION Side By Side Door Refrigerator SHR",
		"Vision Ice Cream Freezer VIS 568L",
	}

	existingkeyMapping := map[string]uint{
		"Brand":      310,
		"Model Name": 316,
	}

	banglaTranslations := s.getBanglaTranslations()

	for _, modelName := range modelsList {
		// build product slug using brand + model
		productSlug := s.slugify(modelName)

		var prod models.ProductModel
		if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("⚠️  Product not found for slug: %s (skipping)", productSlug)
				continue
			}
			return err
		}

		productID := prod.ID

		specs := map[string]string{
			"Brand":      brand,
			"Model Name": modelName,
		}

		for key, val := range specs {
			if len(val) > 500 {
				specs[key] = val[:500]
			}
		}

		for key, value := range specs {
			specKeyID, exists := existingkeyMapping[key]
			if !exists {
				log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (product: %s)", key, productSlug)
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
					if err := db.Create(sModel).Last(&sModel).Error; err != nil {
						return err
					}

					// create translation if available
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
				// update if different
				if existing.Value != value {
					existing.Value = value
					if err := db.Save(&existing).Error; err != nil {
						return err
					}
				}

				// handle translation
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
					} else {
						if existingTranslation.Value != banglaValue {
							existingTranslation.Value = banglaValue
							if err := db.Save(&existingTranslation).Error; err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}
