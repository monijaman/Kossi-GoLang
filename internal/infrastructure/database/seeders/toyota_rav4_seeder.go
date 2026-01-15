package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaRav4Seeder struct {
	BaseSeeder
}

func NewToyotaRav4Seeder() *ToyotaRav4Seeder {
	return &ToyotaRav4Seeder{
		BaseSeeder: BaseSeeder{name: "Toyota RAV4 Specifications"},
	}
}

func (trs *ToyotaRav4Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.5L Petrol": "২.৫ লিটার পেট্রোল",
		"203 hp":      "২০৩ হর্সপাওয়ার",
		"2024":        "২০২৪",
		"Yes":         "হ্যাঁ",
		"No":          "না",
		"8-speed":     "৮-স্পীড",
		"Automatic":   "অটোমেটিক",
		"SUV":         "এসইউভি",
		"Compact SUV": "কমপ্যাক্ট এসইউভি",
		"AWD":         "অল-হুইল ড্রাইভ",
		"23 km/L":     "২৩ কিমি/লিটার",
		"25 km/L":     "২৫ কিমি/লিটার",
		"24 km/L":     "২৪ কিমি/লিটার",
		"1745 kg":     "১৭৪৫ কেজি",
		"Premium":     "প্রিমিয়াম",
		"Dual-zone":   "ডুয়াল-জোন",
		"Touchscreen": "টাচস্ক্রিন",
		"10.5-inch":   "১০.৫-ইঞ্চি",
		"Leather":     "চামড়া",
		"Power":       "পাওয়ার",
		"LED":         "এলইডি",
		"18-inch":     "১৮-ইঞ্চি",
		"Alloy":       "অ্যালয়",
	}
}

func (trs *ToyotaRav4Seeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-rav4").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "RAV4 2.5L",
		729: "RAV4",
		730: "6th Generation",
		731: "SUV",
		732: "2024",
		733: "Inline-4",
		734: "4",
		735: "Naturally Aspirated",
		737: "23 km/L",
		738: "25 km/L",
		739: "24 km/L",
		750: "Yes",
		751: "LED",
		752: "Yes",
		757: "Auto",
		758: "Yes",
		759: "10.5-inch",
		760: "Yes",
		761: "Yes",
		762: "6 Speakers",
		763: "5",
		764: "Yes",
		765: "Yes",
		766: "Yes",
		767: "Yes",
		768: "Yes",
		769: "Yes",
		770: "Rear Camera",
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
		785: "42000 USD",
	}

	banglaTranslations := trs.getBanglaTranslations()

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
