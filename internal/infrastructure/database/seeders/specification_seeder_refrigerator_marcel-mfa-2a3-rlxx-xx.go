package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx seeds specifications/options for product 'marcel-mfa-2a3-rlxx-xx'
type SpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx() *SpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-2a3-rlxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-2a3-rlxx-xx": "মার্সেল-mfa-2a3-rlxx-xx",
		"MFA-2A3-RLXX-XX":        "MFA-2A3-RLXX-XX",
		"176 Ltr.":               "176 লিটার",
		"213 Ltr.":               "213 লিটার",
		"542 x 618 x 1500 mm":    "542 x 618 x 1500 মিমি",
		"45.5 ± 2 Kg":            "45.5 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"220 ~ 240/ 50 Hz":       "220 ~ 240/ 50 হার্জ",
		"50":                     "50",
		"12":                     "12",
		"2 wire shelves":         "2 তারের তাক",
		"Wire":                   "তার",
		"2 PVC":                  "2 PVC",
		"1":                      "1",
		"V 1101 - 102":           "V 1101 - 102",
		"V 1102 - 102":           "V 1102 - 102",
		"V 1301 - 108.6":         "V 1301 - 108.6",
		"V 1302 - 108.6":         "V 1302 - 108.6",
		"V 1401 - 102":           "V 1401 - 102",
		"V 1501 - 99.4":          "V 1501 - 99.4",
		"V 1601 - 108.6":         "V 1601 - 108.6",
		"Copper":                 "তামা",
		"CycloPentane":           "সাইক্লোপেনটেন",
		"580 x 645 x 1530 mm":    "580 x 645 x 1530 মিমি",
		"102/ 102 /50":           "102/ 102 /50",
		"No Need":                "প্রয়োজন নেই",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-2a3-rlxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfa2a3RlxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-2a3-rlxx-xx"
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
		"Model Name":                      "MFA-2A3-RLXX-XX",
		"Capacity":                        "176 Ltr.",
		"Gross Volume":                    "213 Ltr.",
		"Net Volume":                      "176 Ltr.",
		"Dimensions":                      "542 x 618 x 1500 mm",
		"Weight":                          "45.5 ± 2 Kg",
		"Compressor Type":                 "RSCR",
		"Compressor Input Power (Watt)":   "V 1101 - 102; V 1102 - 102; V 1301 - 108.6; V 1302 - 108.6; V 1401 - 102; V 1501 - 99.4; V 1601 - 108.6",
		"Defrost Type":                    "Manual",
		"Refrigerant":                     "R600a",
		"Capillary":                       "Copper",
		"Polyurethane foam blowing agent": "CycloPentane",
		"Packaging Dimensions":            "580 x 645 x 1530 mm",
		"Loading Capacity":                "102/ 102 /50",
		"Frequency (Hz)":                  "50",
		"Warranty":                        "Residential: Replacement Guarantee 1 Year, Compressor 12 Years, Door 3 Years, Spare Parts 4 Years, After Sales Service 5 Years. Commercial: Compressor 4 Years, Door 1 Year, Spare Parts 2 Years, After Sales Service 2 Years.",
		"Compressor Warranty (Years)":     "12",
		"Special Features":                "Lock, Interior Lamp, Vegetable Crisper (1), Egg Tray (1-2), Recessed/Grip Handle, Door Baskets PVC/2, Eco-friendly (100% CFC & HCFC Free) Green Technology",
		"Number of Shelves":               "2 wire shelves",
		"Shelf Material":                  "Wire",
		"Door Bins":                       "2 PVC",
		"Crisper Drawers":                 "1",
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
