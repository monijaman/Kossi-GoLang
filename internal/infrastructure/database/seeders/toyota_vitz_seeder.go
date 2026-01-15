package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaVitzSeeder struct {
	BaseSeeder
}

func NewToyotaVitzSeeder() *ToyotaVitzSeeder {
	return &ToyotaVitzSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Vitz Specifications"},
	}
}

func (tvs *ToyotaVitzSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.3L Petrol":          "১.৩ লিটার পেট্রোল",
		"100 hp":               "১০০ হর্সপাওয়ার",
		"2024":                 "২০২৪",
		"Yes":                  "হ্যাঁ",
		"No":                   "না",
		"CVT":                  "সিভিটি",
		"Hatchback":            "হ্যাচব্যাক",
		"Subcompact Hatchback": "সাবকমপ্যাক্ট হ্যাচব্যাক",
		"Front-Wheel Drive":    "ফ্রন্ট-হুইল ড্রাইভ",
		"19 km/L":              "১৯ কিমি/লিটার",
		"21 km/L":              "২১ কিমি/লিটার",
		"20 km/L":              "২০ কিমি/লিটার",
		"1060 kg":              "১০৬০ কেজি",
		"5-seater":             "৫-সিটার",
		"280L":                 "২৮০ লিটার",
		"Entry-level":          "এন্ট্রি-লেভেল",
		"Manual":               "ম্যানুয়াল",
		"Touchscreen":          "টাচস্ক্রিন",
		"7-inch":               "৭-ইঞ্চি",
		"Fabric":               "ফেব্রিক",
		"Power":                "পাওয়ার",
		"Halogen":              "হ্যালোজেন",
		"14-inch":              "১৪-ইঞ্চি",
		"Steel":                "স্টিল",
	}
}

func (tvs *ToyotaVitzSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-vitz").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Vitz 1.3L",
		729: "Vitz",
		730: "4th Generation",
		731: "Hatchback",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "19 km/L",
		738: "21 km/L",
		739: "20 km/L",
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
		785: "15000 USD",
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
