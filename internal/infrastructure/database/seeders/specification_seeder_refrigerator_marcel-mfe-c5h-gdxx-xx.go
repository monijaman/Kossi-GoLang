package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx seeds specifications/options for product 'marcel-mfe-c5h-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx() *SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfe-c5h-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                     "মার্সেল",
		"marcel-mfe-c5h-gdxx-xx":     "মার্সেল-mfe-c5h-gdxx-xx",
		"MFE-C5H-GDXX-XX":            "MFE-C5H-GDXX-XX",
		"Direct Cool":                "ডাইরেক্ট কুল",
		"358 Ltr.":                   "৩৫৮ লিটার",
		"345 Ltr.":                   "৩৪৫ লিটার",
		"76.5 \u00b1 2 Kg":           "৭৬.৫ ± ২ কেজি",
		"84 \u00b1 2 Kg":             "৮৪ ± ২ কেজি",
		"N ~ ST":                     "N ~ ST",
		"220~240/ 50/135":            "২২০~২৪০/৫০/১৩৫",
		"V 0102- 130\nV 0301- 123":   "V 0102- 130\nV 0301- 123",
		"V 0102- RSCR\nV 0301- RSCR": "V 0102- RSCR\nV 0301- RSCR",
		"Freezer Cabinet Less than -18 ̊C; Refrigerator Cabinet 0 ̊C to +5 ̊C": "ফ্রিজার কক্ষে -18℃ নিচে; রেফ্রিজারেটর কক্ষে 0℃ থেকে +5℃",
		"Mechanical": "যান্ত্রিক",
		"Manual":     "ম্যানুয়াল",
		"No":         "না",
		"Recessed":   "রিসেসড",
		"Yes":        "হ্যাঁ",
		"R600a":      "R600a",
		"Copper":     "কপার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "সাইক্লোপেন্টিন (পরিবেশবান্ধব, CFC/HCFC ফ্রি)",
		"No Need To use Voltage stabilizer.":                                  "ভোল্টেজ স্ট্যাবিলাইজার ব্যবহারের প্রয়োজন নেই।",
		"Wire/2":                                                              "ওয়্যার/2",
		"Yes/3":                                                               "হ্যাঁ/3",
		"Yes/4":                                                               "হ্যাঁ/4",
		"594 x 711 x 1820 mm":                                                 "৫৯৪ x ৭১১ x ১৮২০ মিমি",
		"625 x 745 x 1830 mm":                                                 "৬২৫ x ৭৪৫ x ১৮৩০ মিমি",
		"76/ 57/ 27":                                                          "৭৬/ ৫৭/ ২৭",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfe-c5h-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfeC5hGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfe-c5h-gdxx-xx"
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
		"Model Name":          "MFE-C5H-GDXX-XX",
		"Cooling Technology":  "Direct Cool",
		"Gross Volume":        "358 Ltr.",
		"Net Volume":          "345 Ltr.",
		"Weight":              "76.5 ± 2 Kg",
		"Voltage":             "220~240/ 50/135",
		"Compressor Type":     "V 0102- RSCR; V 0301- RSCR",
		"Refrigerant":         "R600a",
		"Temperature Control": "Mechanical",
		"Defrost Type":        "Manual",
		"Shelf Material":      "Wire/2",
		"Door Bins":           "Yes/3",
		"Crisper Drawers":     "Yes/1",
		"Dimensions":          "594 x 711 x 1820 mm",
		"Special Features":    "Gross Weight: 84 ± 2 Kg; Climate Type: N ~ ST; Compressor Input Power: V 0102- 130; V 0301- 123; Cooling Effect: Freezer Cabinet Less than -18 ̊C; Refrigerator Cabinet 0 ̊C to +5 ̊C; Reversible Door: No; Handle: Recessed; Lock: Yes; Capillary: Copper; Polyurethane foam blowing agent: Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]; Recommended voltage stabilizer capacity: No Need To use Voltage stabilizer.; Interior Lamp: Yes; Egg Tray or Pocket: Yes; Freezer - Rack Shelf: Rack Evaporator; Freezer - Drawer: Yes/4; Ice Tray: Yes/1; Ice Case: Yes/2; Ice Remover spoon: Yes/1; Packing Dimensions: 625 x 745 x 1830 mm; Loading Capacity: 76/ 57/ 27",
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
