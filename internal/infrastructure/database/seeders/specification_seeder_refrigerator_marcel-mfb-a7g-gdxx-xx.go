package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx seeds specifications/options for product 'marcel-mfb-a7g-gdxx-xx'
type SpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx() *SpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfb-a7g-gdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfb-a7g-gdxx-xx": "মার্সেল-mfb-a7g-gdxx-xx",
		"MFB-A7G-GDXX-XX":        "MFB-A7G-GDXX-XX",

		// Common values used in this seeder
		"Direct Cool":                "ডিরেক্ট কুল",
		"177 Ltr.":                   "১৭৭ লিটার",
		"175 Ltr.":                   "১৭৫ লিটার",
		"N~ST":                       "N~ST",
		"220 ~ 240/ 50":              "২২০ ~ ২৪০/ ৫০",
		"V 01.01-88; V 01.02-88":     "V 01.01-88; V 01.02-88",
		"V 01.01-RSCR; V 01.02-RSCR": "V 01.01-RSCR; V 01.02-RSCR",
		"RSCR":                       "RSCR",
		"Mechanical":                 "মেকানিক্যাল",
		"Manual":                     "ম্যানুয়াল",
		"Recessed/ Grip":             "রিসেসড/গ্রিপ",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"Copper":                     "তামা",
		"Cyclopentene":               "সাইক্লোপেন্টেন",
		"R600a":                      "R600a",
		"50 ± 2 Kg":                  "৫০ ± ২ কেজি",
		"54 ± 2 Kg":                  "৫৪ ± ২ কেজি",
		"555 x 630 x 1410 mm":        "৫৫৫ x ৬৩০ x ১৪১০ মিমি",
		"580 x 645 x 1455 mm":        "৫৮০ x ৬৪৫ x ১৪৫৫ মিমি",
		"105/ 105/ 52":               "১০৫/ ১০৫/ ৫২",
		"GPPS/3":                     "জিপিপিএস/৩",
		"Lock: Yes, Interior Lamp: Yes, Handle: Recessed/Grip, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Refrigerator Compartment Shelves: Wire/2, Refrigerator Door Baskets: GPPS/3, Vegetable Crisper: Yes, Egg Tray or Pocket: Yes, Freezer Compartment Shelves: Wire/2, Drawer: HIPS/3, Loading quantity: 105/105/52 (40HQ/40Ft/20Ft), Climatic Type: N~ST, Cooling Effect: Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C, Recommended voltage stabilizer capacity: V 01.01-Low Voltage(140~260V) For V 01.01 - Wide Voltage Range (140Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended. V 01.02-Low Voltage(135~260V) For V01.02 - Wide Voltage Range (135Vac - 260Vac). Voltage stabilizer is not required. In case of voltages beyond this range, 1000VA is recommended.": "লক: হ্যাঁ, ইন্টেরিয়র ল্যাম্প: হ্যাঁ, হ্যান্ডেল: রিসেসড/গ্রিপ, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], রেফ্রিজারেটর কম্পার্টমেন্ট শেল্ফস: ওয়্যার/২, রেফ্রিজারেটর দরজা বাস্কেটস: জিপিপিএস/৩, ভেজিটেবল ক্রিসপার: হ্যাঁ, এগ ট্রে অর পকেট: হ্যাঁ, ফ্রিজার কম্পার্টমেন্ট শেল্ফস: ওয়্যার/২, ড্রয়ার: এইচআইপিএস/৩, লোডিং কোয়ান্টিটি: ১০৫/১০৫/৫২ (৪০এইচকিউ/৪০এফটি/২০এফটি), ক্লাইমেটিক টাইপ: এন~এসটি, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮°C, রেফ্রিজারেটর ক্যাবিনেট ০°C টু +৫°C, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: V ০১.০১-Low Voltage(১৪০~২৬০V) For V ০১.০১ - Wide Voltage Range (১৪০Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is recommended. V ০১.০২-Low Voltage(১৩৫~২৬০V) For V০১.০২ - Wide Voltage Range (১৩৫Vac - ২৬০Vac). Voltage stabilizer is not required. In case of voltages beyond this range, ১০০০VA is recommended.",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfb-a7g-gdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfbA7gGdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfb-a7g-gdxx-xx"
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
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 01.01-88V 01.02-88",
		"Compressor Type": "V 01.01-RSCRV 01.02-RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "630 mm",
		"Door Basket": "GPPS/3",
		"Drawer": "HIPS/3",
		"Egg Tray or Pocket": "Yes",
		"Gross Volume": "177 Ltr.",
		"Gross Weight": "54± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1410 mm",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "105/ 105/ 52",
		"Lock": "Yes",
		"Net Volume": "175 Ltr.",
		"Net Weight": "50± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Recommended voltage stabilizer capacity": "V 01.01-Low Voltage(140~260V)For V 01.01 - Wide Voltage Range (140Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended.V 01.02-Low Voltage(135~260V)For V01.02 - Wide Voltage Range (135Vac - 260Vac).Voltage stabilizer is not required.In case of voltages beyond this range, 1000VA is recommended.",
		"Refrigerant": "R600a",
		"Shelf (Material/ No.)": "Wire/2",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Width/mm": "555 mm",
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
