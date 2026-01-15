package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaVoxySeeder struct {
	BaseSeeder
}

func NewToyotaVoxySeeder() *ToyotaVoxySeeder {
	return &ToyotaVoxySeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Voxy Specifications"},
	}
}

func (tvs *ToyotaVoxySeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Petrol":       "২.০ লিটার পেট্রোল",
		"152 hp":            "১৫২ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"8-speed":           "৮-স্পীড",
		"Automatic":         "অটোমেটিক",
		"MPV":               "এমপিভি",
		"Mid-size MPV":      "মিড-সাইজ এমপিভি",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"13 km/L":           "১৩ কিমি/লিটার",
		"15 km/L":           "১৫ কিমি/লিটার",
		"14 km/L":           "১৪ কিমি/লিটার",
		"1650 kg":           "১৬৫০ কেজি",
		"8-seater":          "৮-সিটার",
		"900L":              "৯০০ লিটার",
		"Premium":           "প্রিমিয়াম",
		"Dual-zone":         "ডুয়াল-জোন",
		"Touchscreen":       "টাচস্ক্রিন",
		"10.5-inch":         "১০.৫-ইঞ্চি",
		"Leather":           "চামড়া",
		"Power":             "পাওয়ার",
		"LED":               "এলইডি",
		"16-inch":           "১৬-ইঞ্চি",
		"Alloy":             "অ্যালয়",
	}
}

func (tvs *ToyotaVoxySeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-voxy").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Voxy 2.0L",
		729: "Voxy",
		730: "3rd Generation",
		731: "MPV",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "13 km/L",
		738: "15 km/L",
		739: "14 km/L",
		750: "Yes",
		751: "LED",
		752: "Yes",
		757: "Auto",
		758: "Yes",
		759: "10.5-inch",
		760: "Yes",
		761: "Yes",
		762: "6 Speakers",
		763: "8",
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
		785: "38000 USD",
	}

	banglaTranslations := tvs.getBanglaTranslations()

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
