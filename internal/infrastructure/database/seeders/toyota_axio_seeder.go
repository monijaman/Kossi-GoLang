package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaAxioSeeder struct {
	BaseSeeder
}

func NewToyotaAxioSeeder() *ToyotaAxioSeeder {
	return &ToyotaAxioSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Axio Specifications"},
	}
}

func (tax *ToyotaAxioSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Petrol":       "১.৫ লিটার পেট্রোল",
		"109 hp":            "১০৯ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"Manual":            "ম্যানুয়াল",
		"CVT":               "সিভিটি",
		"Sedan":             "সেডান",
		"Compact Sedan":     "কমপ্যাক্ট সেডান",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"16 km/L":           "১৬ কিমি/লিটার",
		"18 km/L":           "১৮ কিমি/লিটার",
		"17 km/L":           "১৭ কিমি/লিটার",
	}
}

func (tax *ToyotaAxioSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-axio").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Axio 1.5L",
		729: "Axio",
		730: "2nd Generation",
		731: "Sedan",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "16 km/L",
		738: "18 km/L",
		739: "17 km/L",
		750: "Yes",
		775: "Yes",
		782: "3 Years",
		785: "21000 USD",
	}

	banglaTranslations := tax.getBanglaTranslations()

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
