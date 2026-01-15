package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaCorollaAxioSeeder struct {
	BaseSeeder
}

func NewToyotaCorollaAxioSeeder() *ToyotaCorollaAxioSeeder {
	return &ToyotaCorollaAxioSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Corolla Axio Specifications"},
	}
}

func (tcas *ToyotaCorollaAxioSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Petrol":       "১.৫ লিটার পেট্রোল",
		"109 hp":            "১০৯ হর্সপাওয়ার",
		"2023":              "২০২৩",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"CVT":               "সিভিটি",
		"5-speed":           "৫-স্পীড",
		"Sedan":             "সেডান",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"18 km/L":           "১৮ কিমি/লিটার",
		"20 km/L":           "২০ কিমি/লিটার",
		"19 km/L":           "১৯ কিমি/লিটার",
		"Compact":           "কমপ্যাক্ট",
		"1130 kg":           "১১৩০ কেজি",
	}
}

func (tcas *ToyotaCorollaAxioSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-corolla-axio").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Axio 1.5L",
		729: "Corolla Axio",
		730: "11th Generation",
		731: "Sedan",
		732: "2023",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "18 km/L",
		738: "20 km/L",
		739: "19 km/L",
		782: "3 Years",
		785: "22000 USD",
	}

	banglaTranslations := tcas.getBanglaTranslations()

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
