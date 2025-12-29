package seeders

import (
	"log"
	"regexp"
	"strings"

	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorEcoplus seeds specifications/options for ECO+ models
type SpecificationSeederRefrigeratorEcoplus struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorEcoplus creates a new seeder instance
func NewSpecificationSeederRefrigeratorEcoplus() *SpecificationSeederRefrigeratorEcoplus {
	return &SpecificationSeederRefrigeratorEcoplus{
		BaseSeeder: BaseSeeder{name: "Specifications for ECO+ models"},
	}
}

func (s *SpecificationSeederRefrigeratorEcoplus) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+": "ECO+",
	}
}

// helper to create a product slug similar to existing seeders
func (s *SpecificationSeederRefrigeratorEcoplus) slugify(input string) string {
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

// Seed inserts specification records for ECO+ product models
func (s *SpecificationSeederRefrigeratorEcoplus) Seed(db *gorm.DB) error {
	brand := "ECO+"
	modelsList := []string{
		"BCD-202 FL GD Black RED Line WB",
		"BCD-202 FL GD Black Jumping Line WB",
		"BCD-225 FL GD Black Jumping Line WB",
		"BCD-225 FL GD Black Golden Line WB",
		"BCD-235 FL GD Blue Line WB",
		"BCD-252 FL GD Black Golden Line WB",
		"BD-311SG CHAMPAIN",
		"BCD-252 FL GD SPB-Blue 17 WB",
		"BCD-225 FL GD Magnolia Brown WB",
		"BCD-195 FL GD Magnolia Brown WB",
		"BCD-170 FL GD Black Freesia",
		"SBS-566-RWDGD Black",
		"BCD-218 FL GD Magnolia Marron WB",
		"BCD-195 FL GD SPB Blue-17 WB",
		"BCD-252 FL GD Magnolia Brown WB",
		"BCD-225 FL GD SPB Blue 17 WB",
		"BCD-235 FL GD Magnolia Brown WB",
		"BCD-218 FL GD Magnolia Brown WB",
		"BD-198SG GD Black",
		"BCD-202 FL GD Purple Azalea WB",
		"BD-249SG Sky Blue",
		"BD-198SG CHAMPAIGN",
		"BCD-252 FL GD Magnolia Marron WB",
		"BCD-235 FL GD Purple Azalea WB",
		"BCD-202 FL GD Magnolia Brown WB",
		"BCD-260 REC FL GD Black",
		"BD-311SG Sky Blue",
		"BCD-320 REC FL GD Black",
		"BCD-252 FL GD Magnolia Black WB",
		"BCD-225 FL GD Brown WB",
		"BCD-195 FL GD Black Freesia WB",
		"BCD-170 FL GD Purple Azalea",
		"2B332PLCB",
		"R-VG490P8PB GBK",
		"SBS-566-RDGD Black WOD",
		"OMEGA3 GL-G322RLBB PZ",
		"BCD-235 FL GD Magnolia Black WB",
		"BCD-218 FL GD Black Freesia WB",
		"BD-198SG GD Green",
		"WP IF INV 258 STEEL ONYX",
		"OMEGA5 GL-G252RPBB BC",
		"BCD-225 FL GD Purple WB",
		"HRF-362TBG Black",
		"R-VG490P8PB XGR",
		"OMEGA3 GL-G322RVBB AS",
		"BCD-290 REC FL GD Black",
		"BCD-218 FL GD SPB Blue-17 WB",
		"BCD-202 FL GD SPB Blue-17 WB",
		"WP NEO INV 258GD CRYSTAL BLACK",
		"HCF-230 GD Black",
		"OMEGA4 GL-G302RLBB PZ",
		"BCD-225 FL GD Purple Azalea WB",
		"BD-249SG Green",
		"BD-142SG GD Sky Blue",
		"HCF-340N Silver",
		"R-BG410P6PBX XGR",
		"FD-630-BWDVCM Black Steel",
		"OMEGA5 GL-G252RLBB PZ",
		"BCD-235 FL GD Magnolia Marron WB",
		"BCD-202 FL GD Magnolia Marron WB",
		"BD-142SG Green",
		"SS RB21KMFH5UT/D3",
		"BCD-202 FL GD Brown WB",
		"HCF-175 White",
		"GN-304SLBT",
		"2B312PXCB Ebony Sheen",
		"GL-C322RLBB PZ",
		"GCS215SVF.SPZPFLY PS3",
		"OMEGA5 GL-G252RPBB HP",
		"BCD-235 VCM Refrigerator Red 5",
		"BD-198SG SKY BLUE",
		"R-VG490P8PB XRZ",
		"GL-C322RVBB AS",
		"SS RT29HAR9DUT/D3",
		"SS RS72R5001M9/D2 Silver",
		"SS RT34K5532BS/D3 Black",
		"SS RT29HAR9DS8/D3 Silver",
		"BD-142SG Gray",
		"2B392PXBB Ebony Sheen",
		"2B502HXHL Ebony Sheen",
		"SS RT34K5532S8/D3 Silver",
		"BCD-218 VCM Red 5",
		"SS RT37K5532S8/D3 Silver",
		"GCS315SVF.SPZPFLY PS3",
		"SBS-566-RSVCM Silver",
		"SS RB21KMFH5SE/D3 Silver",
		"BD-198SG Gray",
		"GT-M5097PZ Platinum Silver",
		"GCS155SVF.SPZPFLY PS3",
		"RS-06DR/ BC-46",
		"HRF-680BG Black",
		"HRF-460WDBG Black",
		"GCS215SVC.SSWPFLY White",
		"HRF-360WDBG Black",
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
