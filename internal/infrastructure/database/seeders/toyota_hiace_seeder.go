package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaHiaceSeeder struct {
	BaseSeeder
}

func NewToyotaHiaceSeeder() *ToyotaHiaceSeeder {
	return &ToyotaHiaceSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Hiace Specifications"},
	}
}

func (ths *ToyotaHiaceSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.8L Diesel":      "২.৮ লিটার ডিজেল",
		"177 hp":           "১৭৭ হর্সপাওয়ার",
		"2024":             "২০২৪",
		"Yes":              "হ্যাঁ",
		"No":               "না",
		"6-speed":          "৬-স্পীড",
		"Automatic":        "অটোমেটিক",
		"MPV":              "এমপিভি",
		"Commercial MPV":   "কমার্শিয়াল এমপিভি",
		"Rear-Wheel Drive": "রিয়ার-হুইল ড্রাইভ",
		"10 km/L":          "১০ কিমি/লিটার",
		"12 km/L":          "১২ কিমি/লিটার",
		"11 km/L":          "১১ কিমি/লিটার",
		"2000 kg":          "২০০০ কেজি",
		"9-seater":         "৯-সিটার",
		"4000L":            "৪০০০ লিটার",
		"Commercial":       "কমার্শিয়াল",
		"Manual":           "ম্যানুয়াল",
		"Touchscreen":      "টাচস্ক্রিন",
		"7-inch":           "৭-ইঞ্চি",
		"Fabric":           "ফেব্রিক",
		"Power":            "পাওয়ার",
		"Halogen":          "হ্যালোজেন",
		"16-inch":          "১৬-ইঞ্চি",
		"Steel":            "স্টিল",
	}
}

func (ths *ToyotaHiaceSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-hiace").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Hiace 2.8L Diesel",
		729: "Hiace",
		730: "4th Generation",
		731: "MPV",
		732: "2024",
		733: "Inline-4 Diesel",
		734: "4",
		735: "Turbocharged",
		737: "10 km/L",
		738: "12 km/L",
		739: "11 km/L",
		750: "Yes",
		751: "Halogen",
		752: "No",
		757: "Manual",
		758: "No",
		759: "7-inch",
		760: "Yes",
		761: "No",
		762: "4 Speakers",
		763: "9",
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
		785: "55000 USD",
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
