package seeders

import (
	"log"
	"regexp"
	"strings"

	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorSinger seeds specifications/options for Singer models
type SpecificationSeederRefrigeratorSinger struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorSinger creates a new seeder instance
func NewSpecificationSeederRefrigeratorSinger() *SpecificationSeederRefrigeratorSinger {
	return &SpecificationSeederRefrigeratorSinger{
		BaseSeeder: BaseSeeder{name: "Specifications for Singer refrigerator models"},
	}
}

func (s *SpecificationSeederRefrigeratorSinger) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Singer": "সিঙ্গার",
	}
}

// helper to create a product slug similar to existing seeders
func (s *SpecificationSeederRefrigeratorSinger) slugify(input string) string {
	in := strings.ToLower(strings.TrimSpace(input))
	in = strings.ReplaceAll(in, "+", "plus")
	// replace any non-alphanumeric with hyphen
	re := regexp.MustCompile(`[^a-z0-9]+`)
	out := re.ReplaceAllString(in, "-")
	out = strings.Trim(out, "-")
	// collapse multiple hyphens
	out = regexp.MustCompile(`-+`).ReplaceAllString(out, "-")
	return out
}

// Seed inserts specification records for Singer product models
func (s *SpecificationSeederRefrigeratorSinger) Seed(db *gorm.DB) error {
	brand := "Singer"
	modelsList := []string{
		"FBDS225Z",
		"FBDS225ZBG",
		"FBDS225",
		"FBDS260BG",
		"FBDS260BG2",
		"FBDS260RG",
		"FBDS260ZBG",
		"FBDS260ZRG",
		"FTDS155",
		"FTDS155BG",
		"FTDS155NSL",
		"FTDS155NSV",
		"FTDS155RG",
		"FTDS155V",
		"FTDS155VBG",
		"FTDS185BG",
		"FTDS185BUG",
		"FTDS185NSV",
		"FTDS185V",
		"FTDS185VBG",
		"FTDS200",
		"FTDS200BG",
		"FTDS200V",
		"FTDS200VBG",
		"FTDS230",
		"FTDS230BG",
		"FTDS230RG",
		"FTDS230Z",
		"FTDS230ZBG",
		"FTDS230ZRG",
		"FTDS257",
		"FTDS257RG",
		"FTDS257ZWF",
		"FTDS257ZWFBG",
		"FTDS260BG",
		"FTDS260V",
		"FTDS277RG",
		"FTDS277ZWF",
		"FTDS277ZWFBG",
		"BCD-243R",
		"ICF256ALP",
		"SC250BXLP",
		"SS100FBDS185NSV",
		"SS300FBDS185",
		"SS300FBDS185BG",
		"SS300FTDS185",
		"SS300FTDS185BG",
		"SS300FTDS185NS",
	}

	existingkeyMapping := map[string]uint{
		"Brand":      310,
		"Model Name": 316,
	}

	banglaTranslations := s.getBanglaTranslations()

	for _, modelName := range modelsList {
		// build product slug using brand + model
		productSlug := s.slugify(brand + " " + modelName)

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
