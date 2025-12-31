package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx seeds specifications/options for product 'marcel-mbb-b6x-tdxx-xx'
type SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx() *SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mbb-b6x-tdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mbb-b6x-tdxx-xx":          "মার্সেল-এমবিবি-বি৬এক্স-টিডিএক্সএক্স-এক্সএক্স",
		"MBB-B6X-TDXX-XX":                 "এমবিবি-বি৬এক্স-টিডিএক্সএক্স-এক্সএক্স",
		"Standard Door":                   "স্ট্যান্ডার্ড দরজা",
		"254 Liters":                      "২৫৪ লিটার",
		"0 Liters":                        "০ লিটার",
		"Not Rated":                       "রেটেড নয়",
		"Not Specified":                   "নির্দিষ্ট নয়",
		"620 x 660 x 1873 mm (W x D x H)": "৬২০ x ৬৬০ x ১৮৭৩ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"78.5 kg":                         "৭৮.৫ কেজি",
		"CSIR":                            "সিএসআইআর",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Manual":                          "ম্যানুয়াল",
		"Mechanical":                      "মেকানিক্যাল",
		"Steel":                           "স্টিল",
		"No":                              "না",
		"220-240V":                        "২২০-২৪০ ভোল্ট",
		"50":                              "৫০",
		"Commercial Use: 4 Years on Compressor, 1 Year on Door, 2 Years on Spare Parts, 2 Years After Sales Service": "কমার্শিয়াল ব্যবহার: কম্প্রেসারে ৪ বছর, দরজায় ১ বছর, স্পেয়ার পার্টসে ২ বছর, আফটার সেলস সার্ভিসে ২ বছর",
		"4":          "৪",
		"R134a":      "আর১৩৪এ",
		"260 Liters": "২৬০ লিটার",
		"Lock, Reversible Door No, Handle Recessed/Grip, Skin Condenser 100% Copper, RoHS Certified Thermostat, Copper Capillary, Eco-friendly PU Foam, IP24 Rating, Recommended for Commercial Use": "লক, রিভার্সিবল দরজা না, হ্যান্ডেল রিসেসড/গ্রিপ, স্কিন কন্ডেন্সার ১০০% কপার, রোএইচএস সার্টিফাইড থার্মোস্ট্যাট, কপার ক্যাপিলারি, ইকো-ফ্রেন্ডলি পিইউ ফোম, আইপি২৪ রেটিং, কমার্শিয়াল ব্যবহারের জন্য সুপারিশকৃত",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mbb-b6x-tdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mbb-b6x-tdxx-xx"
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
        "Model Name":          "MBB-B6X-TDXX-XX",
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
