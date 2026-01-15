package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaPriusSeeder struct {
	BaseSeeder
}

func NewToyotaPriusSeeder() *ToyotaPriusSeeder {
	return &ToyotaPriusSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Prius Specifications"},
	}
}

func (tps *ToyotaPriusSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.8L Hybrid":       "১.৮ লিটার হাইব্রিড",
		"172 hp":            "১৭২ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"CVT":               "সিভিটি",
		"Hatchback":         "হ্যাচব্যাক",
		"Hybrid":            "হাইব্রিড",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"25 km/L":           "২৫ কিমি/লিটার",
		"27 km/L":           "২৭ কিমি/লিটার",
		"26 km/L":           "২৬ কিমি/লিটার",
		"Premium":           "প্রিমিয়াম",
		"Dual-zone":         "ডুয়াল-জোন",
		"Touchscreen":       "টাচস্ক্রিন",
		"12.3-inch":         "১২.৩-ইঞ্চি",
		"Leather":           "চামড়া",
		"Power":             "পাওয়ার",
		"LED":               "এলইডি",
		"17-inch":           "১৭-ইঞ্চি",
		"Alloy":             "অ্যালয়",
	}
}

func (tps *ToyotaPriusSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-prius").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Prius 1.8L Hybrid",
		729: "Prius",
		730: "5th Generation",
		731: "Hatchback",
		732: "2024",
		733: "Inline-4 Hybrid",
		734: "4",
		735: "Naturally Aspirated",
		737: "25 km/L",
		738: "27 km/L",
		739: "26 km/L",
		740: "0.75 kWh",
		741: "53 kW",
		742: "163 Nm",
		750: "Yes",
		751: "LED",
		752: "Yes",
		757: "Auto",
		758: "Yes",
		759: "12.3-inch",
		760: "Yes",
		761: "Yes",
		762: "10 Speakers",
		763: "6",
		764: "Yes",
		765: "Yes",
		766: "Yes",
		767: "Yes",
		768: "Yes",
		769: "Yes",
		770: "360° Camera",
		771: "Yes",
		772: "Yes",
		773: "Yes",
		774: "Yes",
		775: "Yes",
		776: "Yes",
		777: "Digital",
		778: "Yes",
		779: "Eco, Normal, Sport, EV",
		780: "Yes",
		781: "Yes",
		782: "3 Years",
		784: "8 Years",
		785: "35000 USD",
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
