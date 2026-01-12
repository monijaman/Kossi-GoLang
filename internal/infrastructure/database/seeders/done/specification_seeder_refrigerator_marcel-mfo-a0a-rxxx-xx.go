package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx seeds specifications/options for product 'marcel-mfo-a0a-rxxx-xx'
type SpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx() *SpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfo-a0a-rxxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfo-a0a-rxxx-xx": "মার্সেল-mfo-a0a-rxxx-xx",
		"MFO-A0A-RXXX-XX":        "MFO-A0A-RXXX-XX",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"101 Ltr.":               "১০১ লিটার",
		"93 Ltr.":                "৯৩ লিটার",
		"25.55 \u00b1 2 Kg":      "২৫.৫৫ ± ২ কেজি",
		"28.6 \u00b1 2 Kg":       "২৮.৬ ± ২ কেজি",
		"T":                      "T",
		"R600a":                  "R600a",
		"Mechanical":             "যান্ত্রিক",
		"Manual":                 "ম্যানুয়াল",
		"No":                     "না",
		"Yes":                    "হ্যাঁ",
		"Copper":                 "তামা",
		"CycloPentene [Eco-friendly (HCFC Free) Green Technology]": "সাইক্লোপেন্টিন (পরিবেশবান্ধব, HCFC ফ্রি)",
		"600VA or More": "৬০০ভিএ বা বেশি",
		"Wire/2":        "ওয়্যার/2",
		"Inside temp. 0C to 5C; Preservation of Fresh food":               "ভিতরের তাপমাত্রা 0°C থেকে 5°C; তাজা খাবার সংরক্ষণ",
		"Inside temp. -2C to +3C; Short time preservation of Frozen food": "ভিতরের তাপমাত্রা -2°C থেকে +3°C; স্বল্প সময়ের জন্য জমা খাবার সংরক্ষণ",
		"490 x 525 x 840 mm": "৪৯০ x ৫২৫ x ৮৪০ মিমি",
		"525 x 535 x 870 mm": "৫২৫ x ৫৩৫ x ৮৭০ মিমি",
		"264/ 176/ 88":       "২৬৪/ ১৭৬/ ৮৮",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfo-a0a-rxxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfoA0aRxxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfo-a0a-rxxx-xx"
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
		"Brand":                           "Marcel",
		"Model Name":                      "MFO-A0A-RXXX-XX",
		"Cooling Technology":              "Direct Cool",
		"Gross Volume":                    "101 Ltr.",
		"Net Volume":                      "93 Ltr.",
		"Net Weight":                      "25.55 ± 2 Kg",
		"Gross Weight":                    "28.6 ± 2 Kg",
		"Refrigerant":                     "R600a",
		"Temperature Control":             "Mechanical",
		"Defrosting":                      "Manual",
		"Reversible Door":                 "No",
		"Lock":                            "Yes",
		"Voltage":                         "V 0101-220~240/ 50/ 69\nV 0201-220~240/ 50/ 67",
		"Compressor":                      "RSCR",
		"Capillary":                       "Copper",
		"Polyurethane foam blowing agent": "CycloPentene [Eco-friendly (HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "600VA or More",
		"Shelf (Material/ No.)":                   "Wire/2",
		"Cooling Effect (Refrigerator)":           "Inside temp. 0C to 5C; Preservation of Fresh food",
		"Cooling Effect (Chilled Room)":           "Inside temp. -2C to +3C; Short time preservation of Frozen food",
		"Dimensions":                              "490 x 525 x 840 mm",
		"Packing Dimensions":                      "525 x 535 x 870 mm",
		"Loading Capacity":                        "264/ 176/ 88",
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
