package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx seeds specifications/options for product 'marcel-mbq-d4x-tdxx-xx'
type SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx() *SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mbq-d4x-tdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mbq-d4x-tdxx-xx":          "মার্সেল-mbq-d4x-tdxx-xx",
		"MBQ-D4X-TDXX-XX":                 "MBQ-D4X-TDXX-XX",
		"Glass Door":                      "গ্লাস ডোর",
		"396 Liters":                      "৩৯৬ লিটার",
		"0 Liters":                        "০ লিটার",
		"440 Liters":                      "৪৪০ লিটার",
		"643 x 738 x 1834 mm (W x D x H)": "৬৪৩ x ৭৩৮ x ১৮৩৪ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"117.25 kg":                       "১১৭.২৫ কেজি",
		"Black":                           "কালো",
		"V 0101-CSR":                      "V 0101-CSR",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Automatic":                       "অটোমেটিক",
		"Mechanical":                      "মেকানিক্যাল",
		"Steel":                           "স্টিল",
		"1":                               "১",
		"No":                              "না",
		"42 dB":                           "৪২ ডেসিবেল",
		"220-240V":                        "২২০-২৪০ ভোল্ট",
		"50":                              "৫০",
		"2 Years":                         "২ বছর",
		"5":                               "৫",
		"R134a":                           "R134a",
		"Beverage Cooler, Commercial Use, External Condenser, Lock, Interior Lamp": "বেভারেজ কুলার, কমার্শিয়াল ব্যবহার, এক্সটার্নাল কন্ডেনসার, লক, ইন্টেরিয়র ল্যাম্প",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mbq-d4x-tdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mbq-d4x-tdxx-xx"
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
		"Brand":                       "Marcel",
		"Model Name":                  "MBQ-D4X-TDXX-XX",
		"Door Type":                   "Glass Door",
		"Capacity":                    "396 Liters",
		"Refrigerator Capacity":       "396 Liters",
		"Freezer Capacity":            "0 Liters",
		"Gross Volume":                "440 Liters",
		"Net Volume":                  "396 Liters",
		"Dimensions":                  "643 x 738 x 1834 mm (W x D x H)",
		"Weight":                      "117.25 kg",
		"Color":                       "Black",
		"Compressor Type":             "V 0101-CSR",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Steel",
		"Number of Shelves":           "1",
		"Door Bins":                   "0",
		"Crisper Drawers":             "0",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "42 dB",
		"Voltage":                     "220-240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "5",
		"Refrigerant":                 "R134a",
		"Special Features":            "Beverage Cooler, Commercial Use, External Condenser, Lock, Interior Lamp",
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
