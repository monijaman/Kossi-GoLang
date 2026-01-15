package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaAquaSeeder struct {
	BaseSeeder
}

func NewToyotaAquaSeeder() *ToyotaAquaSeeder {
	return &ToyotaAquaSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Aqua Specifications"},
	}
}

func (taq *ToyotaAquaSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Hybrid":       "১.৫ লিটার হাইব্রিড",
		"111 hp":            "১১১ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"CVT":               "সিভিটি",
		"Hatchback":         "হ্যাচব্যাক",
		"Hybrid":            "হাইব্রিড",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"28 km/L":           "২৮ কিমি/লিটার",
		"30 km/L":           "৩০ কিমি/লিটার",
		"29 km/L":           "২৯ কিমি/লিটার",
		"1070 kg":           "১০৭০ কেজি",
		"Compact":           "কমপ্যাক্ট",
	}
}

func (taq *ToyotaAquaSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-aqua").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Aqua 1.5L Hybrid",
		729: "Aqua",
		730: "2nd Generation",
		731: "Hatchback",
		732: "2024",
		733: "Inline-4 Hybrid",
		734: "4",
		735: "Naturally Aspirated",
		737: "28 km/L",
		738: "30 km/L",
		739: "29 km/L",
		740: "0.8 kWh",
		741: "10 kW",
		742: "40 Nm",
		750: "Yes",
		760: "Yes",
		761: "Yes",
		775: "Yes",
		782: "3 Years",
		784: "8 Years",
		785: "24000 USD",
	}

	banglaTranslations := taq.getBanglaTranslations()

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
