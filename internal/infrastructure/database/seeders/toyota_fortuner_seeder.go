package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaFortunerSeeder struct {
	BaseSeeder
}

func NewToyotaFortunerSeeder() *ToyotaFortunerSeeder {
	return &ToyotaFortunerSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Fortuner Specifications"},
	}
}

func (tfs *ToyotaFortunerSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.8L Diesel": "২.৮ লিটার ডিজেল",
		"204 hp":      "২০৪ হর্সপাওয়ার",
		"2024":        "২০২৪",
		"Yes":         "হ্যাঁ",
		"No":          "না",
		"6-speed":     "৬-স্পীড",
		"Automatic":   "অটোমেটিক",
		"SUV":         "এসইউভি",
		"Large SUV":   "লার্জ এসইউভি",
		"4WD":         "৪হুইল ড্রাইভ",
		"11 km/L":     "১১ কিমি/লিটার",
		"13 km/L":     "১৩ কিমি/লিটার",
		"12 km/L":     "১২ কিমি/লিটার",
		"2380 kg":     "২৩৮০ কেজি",
		"7-seater":    "৭-সিটার",
		"725L":        "৭২৫ লিটার",
		"Premium":     "প্রিমিয়াম",
		"Dual-zone":   "ডুয়াল-জোন",
		"Touchscreen": "টাচস্ক্রিন",
		"8-inch":      "৮-ইঞ্চি",
		"Leather":     "চামড়া",
		"Power":       "পাওয়ার",
		"LED":         "এলইডি",
		"17-inch":     "১৭-ইঞ্চি",
		"Alloy":       "অ্যালয়",
	}
}

func (tfs *ToyotaFortunerSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-fortuner").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Fortuner 2.8L Diesel",
		729: "Fortuner",
		730: "3rd Generation",
		731: "SUV",
		732: "2024",
		733: "Inline-4 Diesel",
		734: "4",
		735: "Turbocharged",
		737: "11 km/L",
		738: "13 km/L",
		739: "12 km/L",
		750: "Yes",
		751: "LED",
		752: "Yes",
		757: "Auto",
		758: "Yes",
		759: "8-inch",
		760: "Yes",
		761: "Yes",
		762: "6 Speakers",
		763: "7",
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
		779: "Eco, Normal, Sport, Power",
		780: "Yes",
		781: "Yes",
		782: "3 Years",
		785: "48000 USD",
	}

	banglaTranslations := tfs.getBanglaTranslations()

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
