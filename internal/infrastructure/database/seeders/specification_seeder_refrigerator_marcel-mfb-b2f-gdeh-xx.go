package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx seeds specifications/options for product 'marcel-mfb-b2f-gdeh-xx'
type SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbB2fGdehXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbB2fGdehXx() *SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx {
	return &SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-b2f-gdeh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-b2f-gdeh-xx": "মার্সেল-mfb-b2f-gdeh-xx",
		"MFB-B2F-GDEH-XX":        "MFB-B2F-GDEH-XX",

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
		"66 ± 2 Kg": "৬৬ ± ২ কেজি",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Yes": "হ্যাঁ",
		"645 mm": "৬৪৫ মিমি",
		"RSCR": "RSCR",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"HIPS/5": "HIPS/৫",
		"Wire/3": "Wire/৩",
		"RoHS Certified": "RoHS Certified",
		"1765 mm": "১৭৬৫ মিমি",
		"N~ST": "N~ST",
		"V 06.01-124 V 07.01-109 V 07.02-109 V 07.03-109 V 08.01-127": "V ০৬.০১-১২৪ V ০৭.০১-১০৯ V ০৭.০২-১০৯ V ০৭.০৩-১০৯ V ০৮.০১-১২৭",
		"N/A": "N/A",
		"580 mm": "৫৮০ মিমি",
		"V 06.01-Low Voltage(130~260V) For V 06.01- Wide Voltage Range (130Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 07.01,V 07.02,V 07.03-Low Voltage(140~260V) For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended V 08.01-Low Voltage(150~260V) For V 08.01 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "V ০৬.০১-Low Voltage(১৩০~২৬০V) For V ০৬.০১- Wide Voltage Range (১৩০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended. V ০৭.০১,V ০৭.০২,V ০৭.০৩-Low Voltage(১৪০~২৬০V) For V০৭.০১, V ০৭.০২,V ০৭.০৩ - Wide Voltage Range (১৪০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended V ০৮.০১-Low Voltage(১৫০~২৬০V) For V ০৮.০১ - Wide Voltage Range (১৫০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is reco মিমিended.",
		"238 Ltr.": "২৩৮ লিটার",
		"Wire/2": "Wire/২",
		"Recessed/Grip": "Recessed/Grip",
		"252 Ltr.": "২৫২ লিটার",
		"GPPS/4": "GPPS/৪",
		"59 ± 2 Kg": "৫৯ ± ২ কেজি",
		"98/ 72/ 36": "৯৮/ ৭২/ ৩৬",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-b2f-gdeh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbB2fGdehXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-b2f-gdeh-xx"
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
		"Gross Volume": "252 Ltr.",
		"Net Volume": "238 Ltr.",
		"Net Weight": "59 ± 2 Kg",
		"Gross Weight": "66 ± 2 Kg",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Compressor Input Power (Watt)": "V 06.01-124 V 07.01-109 V 07.02-109 V 07.03-109 V 08.01-127",
		"Compressor Type": "RSCR",
		"Energy Rating": "N/A",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Defrosting (Automatic/ Manual):": "Manual",
		"Handle (Recessed/ Grip)": "Recessed/Grip",
		"Lock": "Yes",
		"Refrigerant": "R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "V 06.01-Low Voltage(130~260V) For V 06.01- Wide Voltage Range (130Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 07.01,V 07.02,V 07.03-Low Voltage(140~260V) For V07.01, V 07.02,V 07.03 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended V 08.01-Low Voltage(150~260V) For V 08.01 - Wide Voltage Range (150Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.",
		"Shelf (Material/No.)": "Wire/2",
		"Door Basket": "GPPS/4",
		"Interior Lamp": "Yes",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray": "Yes",
		"Shelf (Material/ No.)": "Wire/3",
		"Drawer": "HIPS/5",
		"Width/mm": "580 mm",
		"Depth/mm": "645 mm",
		"Height/mm": "1765 mm",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "98/ 72/ 36",
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
