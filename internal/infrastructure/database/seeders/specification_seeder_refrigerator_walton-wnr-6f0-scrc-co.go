package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo seeds specifications/options for product 'walton-wnr-6f0-scrc-co'
type SpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo() *SpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo {
	return &SpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wnr-6f0-scrc-co"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WNR-6F0-SCRC-CO":     "ডব্লিউএনআর-৬এফ০-এসসিআরসি-সিও",
		"French Door":         "ফ্রেঞ্চ দরজা",
		"660 Liters":          "৬৬০ লিটার",
		"370 Liters":          "৩৭০ লিটার",
		"243 Liters":          "২৪৩ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"N/A":                 "এন/এ",
		"910 x 756 x 1822 mm": "৯১০ x ৭৫৬ x ১৮২২ মিমি",
		"138 kg":              "১৩৮ কেজি",
		"Depends on Stock":    "স্টকের উপর নির্ভর করে",
		"BDLC Inverter":       "বিডিএলসি ইনভার্টার",
		"No-Frost":            "নো-ফ্রস্ট",
		"Automatic":           "অটোমেটিক",
		"Electronic":          "ইলেকট্রনিক",
		"Glass":               "গ্লাস",
		"4":                   "৪",
		"10":                  "১০",
		"2":                   "২",
		"No":                  "না",
		"Yes":                 "হ্যাঁ",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"1 Year":              "১ বছর",
		"12":                  "১২",
		"Trio Cooling+, Odor Filter, MSO Plus Inverter, LED Lighting, Curved Doors": "ট্রায়ো কুলিং+, ওডর ফিল্টার, এমএসও প্লাস ইনভার্টার, এলইডি লাইটিং, কার্ভড দরজা",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"R600a (CFC Free)": "আর৬০০এ (সিএফসি ফ্রি)",
		"613 Liters":       "৬১৩ লিটার",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wnr-6f0-scrc-co'
func (s *SpecificationSeederRefrigeratorWaltonWnr6f0ScrcCo) Seed(db *gorm.DB) error {
	productSlug := "walton-wnr-6f0-scrc-co"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID
	existingkeyMapping := map[string]uint{
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Special Features":            69,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
	}

	specs := map[string]string{
		"Brand":                       "Walton",
		"Model Name":                  "WNR-6F0-SCRC-CO",
		"Door Type":                   "French Door",
		"Capacity":                    "660 Liters",
		"Refrigerator Capacity":       "370 Liters",
		"Freezer Capacity":            "243 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "910 x 756 x 1822 mm",
		"Weight":                      "138 kg",
		"Color":                       "Depends on Stock",
		"Compressor Type":             "BDLC Inverter",
		"Cooling Technology":          "No-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "10",
		"Crisper Drawers":             "2",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220V",
		"Frequency (Hz)":              "50",
		"App Control":                 "Yes",
		"Voice Assistant Support":     "No",
		"Warranty":                    "1 Year",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R600a (CFC Free)",
		"Gross Volume":                "660 Liters",
		"Net Volume":                  "613 Liters",
		"Special Features":            "Trio Cooling+, Odor Filter, MSO Plus Inverter, LED Lighting, Curved Doors",
	}

	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
			continue
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Last(&sModel).Error; err != nil {
					return err
				}

				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
