package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaHarrierSeeder struct {
	BaseSeeder
}

func NewToyotaHarrierSeeder() *ToyotaHarrierSeeder {
	return &ToyotaHarrierSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Harrier Specifications"},
	}
}

func (ths *ToyotaHarrierSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Petrol":       "২.০ লিটার পেট্রোল",
		"169 hp":            "১৬৯ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"CVT":               "সিভিটি",
		"SUV":               "এসইউভি",
		"Mid-size SUV":      "মিড-সাইজ এসইউভি",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"AWD":               "অল-হুইল ড্রাইভ",
		"19 km/L":           "১৯ কিমি/লিটার",
		"21 km/L":           "২১ কিমি/লিটার",
		"20 km/L":           "২০ কিমি/লিটার",
		"Premium":           "প্রিমিয়াম",
		"Dual-zone":         "ডুয়াল-জোন",
		"Touchscreen":       "টাচস্ক্রিন",
		"12.3-inch":         "১২.৩-ইঞ্চি",
		"Leather":           "চামড়া",
		"Power":             "পাওয়ার",
		"LED":               "এলইডি",
		"18-inch":           "১৮-ইঞ্চি",
		"Alloy":             "অ্যালয়",
	}
}

func (ths *ToyotaHarrierSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-harrier").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Harrier 2.0L",
		729: "Harrier",
		730: "3rd Generation",
		731: "SUV",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "19 km/L",
		738: "21 km/L",
		739: "20 km/L",
		750: "Yes",
		751: "LED",
		752: "Yes",
		757: "Auto",
		758: "Yes",
		759: "12.3-inch",
		760: "Yes",
		761: "Yes",
		762: "8 Speakers",
		763: "5",
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
		779: "Eco, Normal, Sport",
		780: "Yes",
		781: "Yes",
		782: "3 Years",
		785: "45000 USD",
	}

	banglaTranslations := ths.getBanglaTranslations()

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
