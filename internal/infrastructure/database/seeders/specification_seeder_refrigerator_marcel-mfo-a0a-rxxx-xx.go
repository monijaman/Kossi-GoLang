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
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
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
		"Type":                              456,
		"Lock":                              361,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Brand":               "Marcel",
		"Model Name":          "MFO-A0A-RXXX-XX",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "T",
		"Compressor": "RSCR",
		"Cooling Efect": "Inside temp. 0C to 5CPreservation of Fresh food",
		"Cooling Effect": "Inside temp. -2C to +3C​ Short time preservation of Frozen food",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "525",
		"Gross Volume": "101 Ltr.",
		"Gross Weight": "28.6± 2 Kg",
		"Height (mm)": "840",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "264/ 176/ 88",
		"Lock": "Yes",
		"Net Volume": "93 Ltr.",
		"Net Weight": "25.55± 2 Kg",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (HCFC Free) Green Technology]",
		"Rated Voltage/ Hz/ watt": "V 0101-220~240/ 50/ 69V 0201-220~240/ 50/ 67",
		"Recommended voltage stabilizer capacity": "600VA or More",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Width": "490",
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
