package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ToyotaLandCruiserSeeder struct {
	BaseSeeder
}

func NewToyotaLandCruiserSeeder() *ToyotaLandCruiserSeeder {
	return &ToyotaLandCruiserSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Land Cruiser Specifications"},
	}
}

func (tlcs *ToyotaLandCruiserSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.5L Diesel V8": "৩.৫ লিটার ডিজেল ভি৮",
		"282 hp":         "২৮২ হর্সপাওয়ার",
		"2024":           "২০২৪",
		"Yes":            "হ্যাঁ",
		"No":             "না",
		"10-speed":       "১০-স্পীড",
		"Automatic":      "অটোমেটিক",
		"SUV":            "এসইউভি",
		"Full-size SUV":  "ফুল-সাইজ এসইউভি",
		"4WD":            "৪হুইল ড্রাইভ",
		"10 km/L":        "১০ কিমি/লিটার",
		"12 km/L":        "১২ কিমি/লিটার",
		"11 km/L":        "১১ কিমি/লিটার",
		"2700 kg":        "২৭০০ কেজি",
		"8-seater":       "৮-সিটার",
		"850L":           "৮৫০ লিটার",
		"Premium":        "প্রিমিয়াম",
		"Four-zone":      "ফোর-জোন",
		"Touchscreen":    "টাচস্ক্রিন",
		"12.3-inch":      "১২.৩-ইঞ্চি",
		"Leather":        "চামড়া",
		"Power":          "পাওয়ার",
		"LED":            "এলইডি",
		"18-inch":        "১৮-ইঞ্চি",
		"Alloy":          "অ্যালয়",
	}
}

func (tlcs *ToyotaLandCruiserSeeder) Seed(db *gorm.DB) error {
	var product models.ProductModel
	if err := db.Where("slug = ?", "toyota-land-cruiser").First(&product).Error; err != nil {
		return err
	}

	specs := map[uint]string{
		728: "Land Cruiser 3.5L Diesel V8",
		729: "Land Cruiser",
		730: "300 Series",
		731: "SUV",
		732: "2024",
		733: "V8 Diesel",
		734: "4",
		735: "Turbocharged",
		737: "10 km/L",
		738: "12 km/L",
		739: "11 km/L",
		750: "Yes",
		751: "LED",
		752: "Yes",
		757: "Auto",
		758: "Yes",
		759: "12.3-inch",
		760: "Yes",
		761: "Yes",
		762: "14 Speakers",
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
		779: "Eco, Normal, Sport, Sand, Rock",
		780: "Yes",
		781: "Yes",
		782: "3 Years",
		785: "120000 USD",
	}

	banglaTranslations := tlcs.getBanglaTranslations()

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
