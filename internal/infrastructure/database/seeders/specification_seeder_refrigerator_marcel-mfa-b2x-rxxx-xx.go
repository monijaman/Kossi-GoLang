package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx seeds specifications/options for product 'marcel-mfa-b2x-rxxx-xx'
type SpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx() *SpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b2x-rxxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b2x-rxxx-xx": "মার্সেল-mfa-b2x-rxxx-xx",
		"MFA-B2X-RXXX-XX":        "MFA-B2X-RXXX-XX",
		"205 Ltr.":               "205 লিটার",
		"220 Ltr.":               "220 লিটার",
		"545 x 635 x 1580 mm":    "545 x 635 x 1580 মিমি",
		"580 x 645 x 1640 mm":    "580 x 645 x 1640 মিমি",
		"48.5 ± 2 Kg":            "48.5 ± 2 কেজি",
		"52 ± 2 Kg":              "52 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"220 ~ 240/50":           "220 ~ 240/50",
		"50":                     "50",
		"5 Star (BDS 1850:2012)": "5 তারা (BDS 1850:2012)",
		"R600a":                  "R600a",
		"Steel":                  "ইস্পাত",
		"Cyclopentene":           "সাইক্লোপেন্টেন",
		"Wire/3":                 "তার/3",
		"Wire/2":                 "তার/2",
		"GPPS/4":                 "GPPS/4",
		"100/ 100/ 50":           "100/ 100/ 50",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b2x-rxxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaB2xRxxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b2x-rxxx-xx"
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
		"Model Name":                      "MFA-B2X-RXXX-XX",
		"Capacity":                        "205 Ltr.",
		"Gross Volume":                    "220 Ltr.",
		"Net Volume":                      "205 Ltr.",
		"Dimensions":                      "545 x 635 x 1580 mm",
		"Packaging Dimensions":            "580 x 645 x 1640 mm",
		"Weight":                          "48.5 ± 2 Kg (Net); 52 ± 2 Kg (Gross)",
		"Compressor Type":                 "RSCR",
		"Compressor Input Power (Watt)":   "108.6 W",
		"Defrost Type":                    "Manual",
		"Temperature Control":             "Mechanical",
		"Refrigerant":                     "R600a",
		"Capillary":                       "Steel",
		"Polyurethane foam blowing agent": "Cyclopentene",
		"Voltage":                         "220 ~ 240/50",
		"Frequency (Hz)":                  "50",
		"Energy Efficiency Rating":        "5 Star (BDS 1850:2012)",
		"Special Features":                "Direct Cool, Lock, Recessed/Grip Handle, Eco-friendly (100% CFC & HCFC Free) Green Technology, Recommended stabilizer: No need (suggested 1000VA if voltage out of 140V-260V)",
		"Number of Shelves":               "Refrigerator: Wire/3; Freezer: Wire/2",
		"Shelf Material":                  "Wire",
		"Door Bins":                       "GPPS/4",
		"Crisper Drawers":                 "1",
		"Interior Lamp":                   "Yes",
		"Vegetable Crisper":               "Yes/1",
		"Vegetable Crisper Cover":         "Yes/1",
		"Egg Tray or Pocket":              "Yes/1",
		"Loading Capacity":                "100/ 100/ 50",
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
