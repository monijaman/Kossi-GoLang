package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaProboxSeeder struct {
	BaseSeeder
}

func NewToyotaProboxSeeder() *ToyotaProboxSeeder {
	return &ToyotaProboxSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Probox Specifications"},
	}
}

func (tps *ToyotaProboxSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Petrol":       "১.৫ লিটার পেট্রোল",
		"109 hp":            "১০৯ হর্সপাওয়ার",
		"2024":              "২০২৪",
		"Yes":               "হ্যাঁ",
		"No":                "না",
		"5-speed":           "৫-স্পীড",
		"Manual":            "ম্যানুয়াল",
		"Wagon":             "ওয়াগন",
		"Commercial Wagon":  "কমার্শিয়াল ওয়াগন",
		"Front-Wheel Drive": "ফ্রন্ট-হুইল ড্রাইভ",
		"17 km/L":           "১৭ কিমি/লিটার",
		"19 km/L":           "১৯ কিমি/লিটার",
		"18 km/L":           "১৮ কিমি/লিটার",
		"1200 kg":           "১২০০ কেজি",
		"5-seater":          "৫-সিটার",
		"620L":              "৬২০ লিটার",
		"Commercial":        "কমার্শিয়াল",
		"Touchscreen":       "টাচস্ক্রিন",
		"7-inch":            "৭-ইঞ্চি",
		"Fabric":            "ফেব্রিক",
		"Power":             "পাওয়ার",
		"Halogen":           "হ্যালোজেন",
		"15-inch":           "১৫-ইঞ্চি",
		"Steel":             "স্টিল",
	}
}

func (tps *ToyotaProboxSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-probox").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Probox 1.5L",
		729: "Probox",
		730: "2nd Generation",
		731: "Wagon",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "17 km/L",
		738: "19 km/L",
		739: "18 km/L",
		750: "Yes",
		751: "Halogen",
		752: "No",
		757: "Manual",
		758: "No",
		759: "7-inch",
		760: "Yes",
		761: "No",
		762: "4 Speakers",
		763: "5",
		764: "No",
		765: "Yes",
		766: "Yes",
		767: "No",
		768: "No",
		769: "No",
		770: "Rear Camera",
		771: "No",
		772: "No",
		773: "No",
		774: "No",
		775: "Yes",
		776: "Yes",
		777: "Analog",
		778: "No",
		779: "Normal",
		780: "No",
		781: "No",
		782: "3 Years",
		785: "18000 USD",
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
