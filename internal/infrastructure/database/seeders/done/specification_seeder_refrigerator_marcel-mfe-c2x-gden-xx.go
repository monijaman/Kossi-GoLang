package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC2xGdenXx seeds specifications/options for product 'marcel-mfe-c2x-gden-xx'
type SpecificationSeederRefrigeratorMarcelMfeC2xGdenXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC2xGdenXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC2xGdenXx() *SpecificationSeederRefrigeratorMarcelMfeC2xGdenXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC2xGdenXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c2x-gden-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC2xGdenXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                            "মার্সেল",
		"marcel-mfe-c2x-gden-xx":            "মার্সেল-mfe-c2x-gden-xx",
		"MFE-C2X-GDEN-XX":                   "MFE-C2X-GDEN-XX",
		"Direct Cool":                       "ডাইরেক্ট কুল",
		"341 Ltr":                           "৩৪১ লিটার",
		"341 Ltr.":                          "৩৪১ লিটার",
		"320 Ltr.":                          "৩২০ লিটার",
		"67 ± 2 Kg":                         "৬৭ ± ২ কেজি",
		"72 ± 2 Kg":                         "৭২ ± ২ কেজি",
		"N ~ ST":                            "এন ~ এসটি",
		"220-240V~/50Hz":                    "২২০-২৪০V~/৫০Hz",
		"V 0301 - 38~109":                   "V ০৩০১ - ৩৮~১০৯",
		"V 0301 - BLDC":                     "V ০৩০১ - BLDC",
		"Recessed/ Grip":                    "রিসেসড/ গ্রিপ",
		"V 0301 - R600a":                    "V ০৩০১ - R600a",
		"Copper":                            "কপার",
		"RoHS Certified":                    "RoHS সার্টিফায়েড",
		"Cyclopentene":                      "সাইক্লোপেন্টিন",
		"No need to use voltage stabilizer": "স্ট্যাবিলাইজার ব্যবহার করার প্রয়োজন নেই",
		"Wire/2":                            "ওয়্যার/২",
		"PVC/4":                             "পিভিসি/৪",
		"594 x 708 x 1720 mm":               "৫৯৪ x ৭০৮ x ১৭২০ মিমি",
		"635 x 740 x 1790 mm":               "৬৩৫ x ৭৪০ x ১৭৯০ মিমি",
		"77/ 57/ 27":                        "৭৭/ ৫৭/ ২৭",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c2x-gden-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC2xGdenXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c2x-gden-xx"
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
		"Model Name":          "MFE-C2X-GDEN-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "341 Ltr",
		"Net Volume":          "320 Ltr.",
		"Weight":              "67 ± 2 Kg",
		"Voltage":             "220-240V~/50Hz",
		"Compressor Type":     "V 0301 - BLDC",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Refrigerant":         "V 0301 - R600a",
		"Shelf Material":      "Wire",
		"Number of Shelves":   "2",
		"Door Bins":           "PVC/4",
		"Crisper Drawers":     "Yes/1",
		"Dimensions":          "594 x 708 x 1720 mm",
		"Special Features":    "Gross Weight: 72 ± 2 Kg; Climatic Type: N ~ ST; Compressor Input Power: V 0301 - 38~109; Thermostat: RoHS Certified; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene; Recommended stabilizer: No need to use voltage stabilizer; Freezer Shelf: Wire/2; Loading Capacity: 77/ 57/ 27; Packing Dimensions: 635 x 740 x 1790 mm; Interior Lamp: Yes; Can Storage Dispenser: No",
		"Loading Capacity":    "77/ 57/ 27",
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
