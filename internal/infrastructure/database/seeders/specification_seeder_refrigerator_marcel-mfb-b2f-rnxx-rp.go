package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp seeds specifications/options for product 'marcel-mfb-b2f-rnxx-rp'
type SpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp() *SpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-rnxx-rp"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b2f-rnxx-rp": "মার্সেল-mfb-b2f-rnxx-rp",
		"MFB-B2F-RNXX-RP":        "MFB-B2F-RNXX-RP",

		"Direct Cool":         "ডিরেক্ট কুল",
		"177 Ltr.":            "১৭৭ লিটার",
		"175 Ltr.":            "১৭৫ লিটার",
		"50 ± 2 Kg":           "৫০ ± ২ কেজি",
		"R600a":               "R600a",
		"Mechanical":          "ম্যানুয়াল/মেকানিক্যাল",
		"Manual":              "ম্যানুয়াল",
		"Recessed/ Grip":      "রিসেসড/গ্রিপ",
		"Copper":              "কপার",
		"Cyclopentene":        "সাইক্লোপেন্টেন",
		"220 ~ 240":           "২২0 ~ ২৪0",
		"Hermetic":            "হার্মেটিক",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"Compressor: V 01.01-RSCR; V 01.02-RSCR; Compressor Input Power (Watt): V 01.01-88; V 01.02-88; Voltage/Hz: 220 ~ 240/ 50; Stabilizer: 5 Ampere": "কমপ্রেসার: V 01.01-RSCR; V 01.02-RSCR; কমপ্রেসার ইনপুট পাওয়ার (ওয়াট): V 01.01-88; V 01.02-88; ভোল্টেজ/হার্জ: 220 ~ 240/50; স্ট্যাবিলাইজার: ৫ অ্যাম্পিয়ার",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-rnxx-rp'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fRnxxRp) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-rnxx-rp"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
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
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}

	
    		
    specs := map[string]string{
        "Brand":               "Marcel",
        "Model Name":          "MFB-B2F-RNXX-RP",
        "Cooling Technology":  "Direct Cool",
        "Gross Volume":        "177 Ltr.",
        "Net Volume":          "175 Ltr.",
        "Weight":              "50 ± 2 Kg",
        "Refrigerant":         "R600a",
        "Temperature Control": "Mechanical",
        "Voltage":             "220 ~ 240",
        "Dimensions":          "555 x 630 x 1410 mm",
        "Packing Dimensions":  "580 x 645 x 1455 mm",
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
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Ensure the ID is set after creation
				if sModel.ID == 0 {
					if err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {
						return err
					}
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
