package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter seeds specifications/options for product 'marcel-mfb-b5d-gdel-sc-inverter'
type SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter() *SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter {
	return &SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b5d-gdel-sc-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mfb-b5d-gdel-sc-inverter": "মার্সেল-mfb-b5d-gdel-sc-inverter",
		"MFB-B5D-GDEL-SC-INVERTER":        "MFB-B5D-GDEL-SC-INVERTER",

		"Direct Cool":         "ডিরেক্ট কুল",
		"177 Ltr.":            "১৭৭ লিটার",
		"175 Ltr.":            "১৭৫ লিটার",
		"50 ± 2 Kg":           "৫০ ± ২ কেজি",
		"R600a":               "R600a",
		"Mechanical":          "ম্যানুয়াল/মেকানিক্যাল",
		"Manual":              "ম্যানুয়াল",
		"Recessed/ Grip":      "রিসেসড/গ্রিপ",
		"Copper":              "কপার",
		"Cyclopentene":        "সাইক্লোপেন্টেন",
		"220 ~ 240":           "২২0 ~ ২৪0",
		"Hermetic":            "হার্মেটিক",
		"555 x 630 x 1410 mm": "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm": "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"Compressor: V 01.01-RSCR; V 01.02-RSCR; Compressor Input Power (Watt): V 01.01-88; V 01.02-88; Voltage/Hz: 220 ~ 240/ 50; Stabilizer: 5 Ampere": "কমপ্রেসার: V 01.01-RSCR; V 01.02-RSCR; কমপ্রেসার ইনপুট পাওয়ার (ওয়াট): V 01.01-88; V 01.02-88; ভোল্টেজ/হার্জ: 220 ~ 240/50; স্ট্যাবিলাইজার: ৫ অ্যাম্পিয়ার",
		"1845 mm": "১৮৪৫ মিমি",
		"For V 05.03 - Wide Voltage Range (75Vac - 264Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "For V ০৫.০৩ - Wide Voltage Range (৭৫Vac - ২৬৪Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended.",
		"67 ± 2 Kg": "৬৭ ± ২ কেজি",
		"97/ 72/ 36": "৯৭/ ৭২/ ৩৬",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"V 05.03-BLDC": "V ০৫.০৩-BLDC",
		"645 mm": "৬৪৫ মিমি",
		"Electronic": "ইলেকট্রনিক",
		"268 Ltr.": "২৬৮ লিটার",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"HIPS/5": "HIPS/৫",
		"5 star (BDS 1850:2012)": "৫ star (BDS ১৮৫০:২০১২)",
		"Wire/3": "Wire/৩",
		"RoHS Certified": "RoHS Certified",
		"No": "না",
		"N~ST": "N~ST",
		"254 Ltr.": "২৫৪ লিটার",
		"580 mm": "৫৮০ মিমি",
		"V 05.03-33~126": "V ০৫.০৩-৩৩~১২৬",
		"62 ± 2 Kg": "৬২ ± ২ কেজি",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b5d-gdel-sc-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMfbB5dGdelScInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b5d-gdel-sc-inverter"
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
		"Type": "Direct Cool",
		"Gross Volume": "268 Ltr.",
		"Net Volume": "254 Ltr.",
		"Net Weight": "62 ± 2 Kg",
		"Gross Weight": "67 ± 2 Kg",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Compressor Input Power (Watt)": "V 05.03-33~126",
		"Compressor Type": "V 05.03-BLDC",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Energy Rating": "5 star (BDS 1850:2012)",
		"Temperature Control (Electronic/ Mechanical)": "Electronic",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "No",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "Yes",
		"Refrigerant": "R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "For V 05.03 - Wide Voltage Range (75Vac - 264Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.",
		"Shelf (Material/ No.)": "Wire/3",
		"Door Pocket": "No",
		"Interior Lamp": "No",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray": "Yes",
		"Drawer": "HIPS/5",
		"Width/mm": "580 mm",
		"Depth/mm": "645 mm",
		"Height/mm": "1845 mm",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "97/ 72/ 36",
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
