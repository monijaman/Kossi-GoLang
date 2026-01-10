package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5xGdshXx seeds specifications/options for product 'marcel-mfb-b5x-gdsh-xx'
type SpecificationSeederRefrigeratorMarcelMfbB5xGdshXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5xGdshXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5xGdshXx() *SpecificationSeederRefrigeratorMarcelMfbB5xGdshXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB5xGdshXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5x-gdsh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdshXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1675 mm":       "১৬৭৫ মিমি",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"250 Ltr.":      "২৫০ লিটার",
		"274 Ltr.":      "২৭৪ লিটার",
		"54.5 ± 2 Kg":   "৫৪.৫ ± ২ কেজি",
		"555 mm":        "৫৫৫ মিমি",
		"59.5 ± 2 Kg":   "৫৯.৫ ± ২ কেজি",
		"5 mm":          "৫ mm",
		"580 mm":        "৫৮০ মিমি",
		"630 mm":        "৬৩০ মিমি",
		"97/ 72/ 36":    "৯৭/ ৭২/ ৩৬",
		"Copper":        "কপার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Direct Cool": "ডাইরেক্ট কুল",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"GPPS/3":         "GPPS/3",
		"Mechanical":     "মেকানিক্যাল",
		"Manual":         "ম্যানুয়াল",
		"N~ST":           "N~ST",
		"No":             "না",
		"Recessed/ Grip": "রিসেসড/ গ্রিপ",
		"RoHS Certified": "RoHS Certified",
		"R600a":          "R600a",
		"RSCR":           "RSCR",
		"V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4": "V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4",
		"Wire/2": "ওয়্যার/2",
		"645 mm": "৬৪৫ মিমি",
		"Yes":    "হ্যাঁ",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b5x-gdsh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5xGdshXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5x-gdsh-xx"
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
		"Capillary":                               719,
		"Climatic Type (SN, N, ST, T)":            712,
		"Compressor Input Power (Watt)":           713,
		"Compressor Type":                         139,
		"Cooling Effect":                          714,
		"Defrosting (Automatic/ Manual)":          141,
		"Door Basket":                             700,
		"Drawer":                                  724,
		"Egg Tray":                                723,
		"Gross Volume":                            709,
		"Gross Weight":                            81,
		"Handle (Recessed/ Grip)":                 716,
		"Interior Lamp":                           722,
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":      725,
		"Lock":                                    717,
		"Net Depth/mm":                            25,
		"Net Height/mm":                           25,
		"Net Volume":                              710,
		"Net Weight":                              80,
		"Net Width/mm":                            25,
		"Packaging Depth/mm":                      25,
		"Packaging Height/mm":                     25,
		"Packaging Width/mm":                      25,
		"Polyurethane foam blowing agent":         720,
		"Rated Voltage/ Hz":                       160,
		"Recommended voltage stabilizer capacity": 721,
		"Refrigerant":                             708,
		"Shelf (Material/ No.)":                   699,
		"Temperature Control (Electronic/ Mechanical)": 715,
		"Thermostat":              718,
		"Type":                    3,
		"Vegetable Crisper":       701,
		"Vegetable Crisper Cover": 701,
	}

	specs := map[string]string{
		"Capillary":                               "Copper",
		"Climatic Type (SN, N, ST, T)":            "N~ST",
		"Compressor Input Power (Watt)":           "V 01.01-97.4 V 02.01-97.4 V 03.01-97.4 V 03.02-97.4",
		"Compressor Type":                         "RSCR",
		"Cooling Effect":                          "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)":          "Manual",
		"Door Basket":                             "GPPS/3",
		"Drawer":                                  "No",
		"Egg Tray":                                "Yes",
		"Gross Volume":                            "250 Ltr.",
		"Gross Weight":                            "59.5 ± 2 Kg",
		"Handle (Recessed/ Grip)":                 "Recessed/ Grip",
		"Interior Lamp":                           "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft":      "97/ 72/ 36",
		"Lock":                                    "Yes",
		"Net Depth/mm":                            "630 mm",
		"Net Height/mm":                           "1675 mm",
		"Net Volume":                              "274 Ltr.",
		"Net Weight":                              "54.5 ± 2 Kg",
		"Net Width/mm":                            "555 mm",
		"Packaging Depth/mm":                      "5 mm",
		"Packaging Height/mm":                     "1725 mm",
		"Packaging Width/mm":                      "580 mm",
		"Polyurethane foam blowing agent":         "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Voltage/ Hz":                       "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 01.01, V 02.01-Low Voltage(140~260V) For V01.01, V 02.01-Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 03.01, V 03.02-Low Voltage(150~260V) For V 03.01, V 03.02 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.",
		"Refrigerant":                             "R600a",
		"Shelf (Material/ No.)":                   "Wire/2",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Thermostat":              "RoHS Certified",
		"Type":                    "Direct Cool",
		"Vegetable Crisper":       "Yes",
		"Vegetable Crisper Cover": "Yes",
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
