package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaAllionSeeder struct {
	BaseSeeder
}

func NewToyotaAllionSeeder() *ToyotaAllionSeeder {
	return &ToyotaAllionSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Allion Specifications"},
	}
}

func (tas *ToyotaAllionSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Petrol":       "১.৫ লিটার পেট্রোল",
		"109 hp":            "১০৯ হর্সপাওয়ার",
		"2023":              "২০২৩",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"CVT":               "সিভিটি",
		"Sedan":             "সেডান",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"17 km/L":           "১৭ কিমি/লিটার",
		"19 km/L":           "১৯ কিমি/লিটার",
		"18 km/L":           "১৮ কিমি/লিটার",
	}
}

func (tas *ToyotaAllionSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-allion").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Allion 1.5L",
		729: "Allion",
		730: "3rd Generation",
		731: "Sedan",
		732: "2023",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "17 km/L",
		738: "19 km/L",
		739: "18 km/L",
		782: "3 Years",
		785: "23000 USD",
	}

	banglaTranslations := tas.getBanglaTranslations()

	for keyID, value := range specs {
		var existingSpec models.SpecificationModel
		result := db.Where("product_id = ? AND specification_key_id = ?", product.ID, keyID).First(&existingSpec)

		if result.Error == gorm.ErrRecordNotFound {
			spec := models.SpecificationModel{
				ProductID:          product.ID,
				SpecificationKeyID: keyID,
				Value:              value,
				Status:             1,
			}
			if err := db.Create(&spec).Error; err != nil {
				continue
			}

			if banglaValue, ok := banglaTranslations[value]; ok {
				translation := models.SpecificationTranslationModel{
					SpecificationID: spec.ID,
					Locale:          "bn",
					Value:           banglaValue,
				}
				db.Create(&translation)
			}
		}
	}

	return nil
}
