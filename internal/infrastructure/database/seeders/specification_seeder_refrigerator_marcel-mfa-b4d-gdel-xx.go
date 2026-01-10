package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfaB4dGdelXx seeds specifications/options for product 'marcel-mfa-b4d-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMfaB4dGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfaB4dGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfaB4dGdelXx() *SpecificationSeederRefrigeratorMarcelMfaB4dGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMfaB4dGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-b4d-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfaB4dGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-b4d-gdel-xx": "মার্সেল-mfa-b4d-gdel-xx",
		"MFA-B4D-GDEL-XX":        "MFA-B4D-GDEL-XX",
		"244 Ltr.":               "244 লিটার",
		"220 Ltr.":               "220 লিটার",
		"545 x 605 x 1760 mm":    "545 x 605 x 1760 মিমি",
		"580 x 645 x 1770 mm":    "580 x 645 x 1770 মিমি",
		"51/58":                  "51/58",
		"RSCR":                   "RSCR",
		"Manual":                 "ম্যানুয়াল",
		"Mechanical":             "মেকানিক্যাল",
		"R134a":                  "R134a",
		"R600a":                  "R600a",
		"Copper":                 "তামা",
		"Cyclopentene":           "সাইক্লোপেন্টেন",
		"Wire/3":                 "তার/3",
		"Wire/2":                 "তার/2",
		"GPPS/4":                 "GPPS/4",
		"5 Star":                 "5 তারা",
		"V 0501-119":             "V 0501-119",
		"V 0601-108.6":           "V 0601-108.6",
		"V 0701-108.6":           "V 0701-108.6",
		"1000VA":                 "1000VA",
		"98/ 72/ 36":             "98/ 72/ 36",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-b4d-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfaB4dGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-b4d-gdel-xx"
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
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 0501-119V 0601-108.6V 0701 108.6",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Deodorizer": "No",
		"Depth/mm": "605",
		"Door Basket": "No",
		"Door Pocket": "GPPS/4",
		"Drawer": "No",
		"Egg Case": "Yes/1",
		"Energy Rating": "5 star (BDS 1850:2012)",
		"Gross Volume(Outer Dimension, Manufacturer declared)": "244 Ltr.",
		"Handle (Recessed/ Grip)": "Recessed",
		"Height/mm": "1760",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "98/ 72/ 36",
		"Lock": "Yes",
		"Net Volume": "220 Ltr.",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Recommended voltage stabilizer capacity": "V05.01: 1000VAV06.01: No Need to use voltage stabilizer.If out of voltage range(145V-260V), then suggested voltage stabilizer capacity is 1000VA.",
		"Refrigerant": "V 0501-R134aV 0601-R600aV 0701-R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Shelf (Material/No.)": "Wire/3",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Box": "Yes/1",
		"Vegetable Box Cover": "Yes/1",
		"Weight/Kg - Net/Packing (±2 KG)": "51/58",
		"Width/mm": "545",
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
