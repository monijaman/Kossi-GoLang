package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaPassoSeeder struct {
	BaseSeeder
}

func NewToyotaPassoSeeder() *ToyotaPassoSeeder {
	return &ToyotaPassoSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Passo Specifications"},
	}
}

func (tps *ToyotaPassoSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.0L Petrol":             "১.০ লিটার পেট্রোল",
		"67 hp":                   "৬৭ হর্সপাওয়ার",
		"2024":                    "২০২৪",
		"Yes":                     "হ্যাঁ",
		"No":                      "না",
		"CVT":                     "সিভিটি",
		"Hatchback":               "হ্যাচব্যাক",
		"Ultra-compact Hatchback": "আল্ট্রা-কমপ্যাক্ট হ্যাচব্যাক",
		"Front-Wheel Drive":       "ফ্রন্ট-হুইল ড্রাইভ",
		"21 km/L":                 "২১ কিমি/লিটার",
		"23 km/L":                 "২৩ কিমি/লিটার",
		"22 km/L":                 "২২ কিমি/লিটার",
		"920 kg":                  "৯২০ কেজি",
		"4-seater":                "৪-সিটার",
		"268L":                    "২৬৮ লিটার",
		"Budget":                  "বাজেট",
		"Manual":                  "ম্যানুয়াল",
		"Touchscreen":             "টাচস্ক্রিন",
		"7-inch":                  "৭-ইঞ্চি",
		"Fabric":                  "ফেব্রিক",
		"Power":                   "পাওয়ার",
		"Halogen":                 "হ্যালোজেন",
		"13-inch":                 "১৩-ইঞ্চি",
		"Steel":                   "স্টিল",
	}
}

func (tps *ToyotaPassoSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-passo").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Passo 1.0L",
		729: "Passo",
		730: "2nd Generation",
		731: "Hatchback",
		732: "2024",
		733: "Inline-3",
		734: "4",
		735: "Naturally Aspirated",
		737: "21 km/L",
		738: "23 km/L",
		739: "22 km/L",
		750: "Yes",
		751: "Halogen",
		752: "No",
		757: "Manual",
		758: "No",
		759: "7-inch",
		760: "Yes",
		761: "No",
		762: "4 Speakers",
		763: "4",
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
		785: "12000 USD",
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
