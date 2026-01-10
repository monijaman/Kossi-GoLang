package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaA9cGdshXx seeds specifications/options for product 'marcel-mfa-a9c-gdsh-xx'
type SpecificationSeederRefrigeratorMarcelMfaA9cGdshXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaA9cGdshXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaA9cGdshXx() *SpecificationSeederRefrigeratorMarcelMfaA9cGdshXx {
	return &SpecificationSeederRefrigeratorMarcelMfaA9cGdshXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-a9c-gdsh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaA9cGdshXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-a9c-gdsh-xx": "মার্সেল-mfa-a9c-gdsh-xx",
		"MFA-A9C-GDSH-XX":        "MFA-A9C-GDSH-XX",
		"175 Ltr.":               "175 লিটার",
		"193 Ltr.":               "193 লিটার",
		"538 x 600 x 1230 mm":    "538 x 600 x 1230 মিমি",
		"45 ± 2 Kg":              "45 ± 2 কেজি",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"R600a":                  "R600a",
		"220-240 V/ 50 Hz":       "220-240 ভোল্ট/ 50 হার্জ",
		"50":                     "50",
		"3 glass shelves":        "3 গ্লাস তাক",
		"Glass":                  "গ্লাস",
		"5 PS":                   "5 PS",
		"1":                      "1",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-a9c-gdsh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaA9cGdshXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-a9c-gdsh-xx"
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
		"Brand":                             310,
		"Model Name":                        316,
		"Door Type":                         142,
		"Capacity":                          138,
		"Refrigerator Capacity":             156,
		"Freezer Capacity":                  146,
		"Energy Efficiency Rating":          143,
		"Energy Star Rating":                144,
		"Annual Energy Consumption":         137,
		"Dimensions":                        25,
		"Weight":                            80,
		"Color":                             311,
		"Compressor Type":                   139,
		"Cooling Technology":                698,
		"Defrost Type":                      141,
		"Temperature Control":               158,
		"Shelf Material":                    699,
		"Number of Shelves":                 154,
		"Door Bins":                         700,
		"Crisper Drawers":                   701,
		"Ice Maker":                         702,
		"Water Dispenser":                   703,
		"Noise Level":                       150,
		"Voltage":                           160,
		"Frequency (Hz)":                    704,
		"App Control":                       705,
		"Voice Assistant Support":           385,
		"Warranty":                          323,
		"Compressor Warranty (Years)":       707,
		"Refrigerant":                       708,
		"Gross Volume":                      709,
		"Net Volume":                        710,
		"Special Features":                  69,
		"Capillary":                         711,
		"Climate Type (SN, N, ST, T)":       712,
		"Compressor Input Power (Watt)":     713,
		"Cooling Effect":                    714,
		"Defrosting (Automatic/ Manual)":    715,
		"Deodorizer":                        716,
		"Depth (mm)":                        717,
		"Door Basket":                       718,
		"Egg Tray or Pocket":                719,
		"Handle (Recessed/ Grip)":           720,
		"Interior Lamp":                     721,
		"Polyurethane Foam Blowing Agent":   722,
		"Reversible Door":                   723,
		"Thermostat":                        724,
		"Vegetable Crisper":                 725,
		"Vegetable Crisper Cover":           726,
		"Loading Capacity (40HQ/40Ft/20Ft)": 727,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"2.25 L Bottle Accommodation": "Yes",
		"Bottle Basket": "PS/5",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "T",
		"Compressor": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "600",
		"Door Basket": "No",
		"Double Layer Freezer Door": "Yes",
		"Drawer": "No",
		"Egg Tray or Pocket": "Yes/1",
		"Gross Volume": "193 Ltr.",
		"Gross Weight": "50± 2Kg",
		"Handle (Recessed/ Grip)": "Recessed",
		"Height/mm": "1230",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "144/109/49",
		"Lock": "Yes",
		"Medicine Box": "Yes/1",
		"Net Volume": "175 Ltr.",
		"Net Weight": "45± 2Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz/ watt": "220-240 V/ 50 Hz/ 72W",
		"Recommended voltage stabilizer capacity": "No Voltage Stabilizer",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "No",
		"Shelf (Material/No.)": "Glass/3",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Uniform Flow": "Yes",
		"Vegetable Crisper": "Yes/1 (17L)",
		"Vegetable Crisper Cover": "Yes (Glass)",
		"Width/mm": "538",
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
