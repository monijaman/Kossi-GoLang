package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaPremioSeeder struct {
	BaseSeeder
}

func NewToyotaPremioSeeder() *ToyotaPremioSeeder {
	return &ToyotaPremioSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Premio Specifications"},
	}
}

func (tps *ToyotaPremioSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.8L Petrol":       "১.৮ লিটার পেট্রোল",
		"140 hp":            "১৪০ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"CVT":               "সিভিটি",
		"Sedan":             "সেডান",
		"Premium Sedan":     "প্রিমিয়াম সেডান",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"14 km/L":           "১৪ কিমি/লিটার",
		"16 km/L":           "১৬ কিমি/লিটার",
		"15 km/L":           "১৫ কিমি/লিটার",
		"1350 kg":           "১৩৫০ কেজি",
	}
}

func (tps *ToyotaPremioSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-premio").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Premio 1.8L",
		729: "Premio",
		730: "3rd Generation",
		731: "Sedan",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "14 km/L",
		738: "16 km/L",
		739: "15 km/L",
		750: "Yes",
		760: "Yes",
		761: "Yes",
		775: "Yes",
		776: "Yes",
		782: "3 Years",
		785: "26000 USD",
	}

	banglaTranslations := tps.getBanglaTranslations()

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
