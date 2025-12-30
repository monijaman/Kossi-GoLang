package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB2xGdelXx seeds specifications/options for product 'marcel-mfa-b2x-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfaB2xGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB2xGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB2xGdelXx() *SpecificationSeederRefrigeratorMarcelMfaB2xGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfaB2xGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b2x-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB2xGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b2x-gdel-xx": "মার্সেল-mfa-b2x-gdel-xx",
		"MFA-B2X-GDEL-XX":        "MFA-B2X-GDEL-XX",
		"205 Ltr.":               "205 লিটার",
		"220 Ltr.":               "220 লিটার",
		"557 x 635 x 1580 mm":    "557 x 635 x 1580 মিমি",
		"48.5 ± 2 Kg":            "48.5 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"R600a":                  "R600a",
		"220 ~ 240/50":           "220 ~ 240/50",
		"50":                     "50",
		"5 Star":                 "5 তারা",
		"Refrigerator: 3 wire shelves; Freezer: 2 wire shelves": "রেফ্রিজারেটর: 3 তারের তাক; ফ্রিজার: 2 তারের তাক",
		"Wire":   "তার",
		"4 GPPS": "4 GPPS",
		"1":      "1",
		"12":     "12",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b2x-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaB2xGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b2x-gdel-xx"
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
		"Brand":                    "Marcel",
		"Model Name":               "MFA-B2X-GDEL-XX",
		"Capacity":                 "205 Ltr.",
		"Gross Volume":             "220 Ltr.",
		"Net Volume":               "205 Ltr.",
		"Dimensions":               "557 x 635 x 1580 mm",
		"Weight":                   "48.5 ± 2 Kg",
		"Compressor Type":          "RSCR",
		"Defrost Type":             "Manual",
		"Temperature Control":      "Mechanical",
		"Refrigerant":              "R600a",
		"Voltage":                  "220 ~ 240/50",
		"Frequency (Hz)":           "50",
		"Energy Efficiency Rating": "5 Star", "Warranty": "Residential: Compressor 12 Years, Door 3 Years, Spare Parts 4 Years, After Sales Service 5 Years; Commercial: Compressor 4 Years, Door 1 Year, Spare Parts 2 Years, After Sales Service 2 Years",
		"Compressor Warranty (Years)": "12", "Special Features": "Lock, Interior Lamp, Vegetable Crisper (1), Egg Tray (1), Recessed/Grip Handle, Eco-friendly (100% CFC & HCFC Free) Green Technology",
		"Number of Shelves": "Refrigerator: 3 wire shelves; Freezer: 2 wire shelves",
		"Shelf Material":    "Wire",
		"Door Bins":         "4 GPPS",
		"Crisper Drawers":   "1",
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
