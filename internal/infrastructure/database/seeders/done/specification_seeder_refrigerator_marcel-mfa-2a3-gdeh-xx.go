package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfa2a3GdehXx seeds specifications/options for product 'marcel-mfa-2a3-gdeh-xx'
type SpecificationSeederRefrigeratorMarcelMfa2a3GdehXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfa2a3GdehXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfa2a3GdehXx() *SpecificationSeederRefrigeratorMarcelMfa2a3GdehXx {
	return &SpecificationSeederRefrigeratorMarcelMfa2a3GdehXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfa-2a3-gdeh-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfa2a3GdehXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfa-2a3-gdeh-xx": "মার্সেল-mfa-2a3-gdeh-xx",
		"MFA-2A3-GDEH-XX":        "MFA-2A3-GDEH-XX",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"176 Ltr.":               "১৭৬ লিটার",
		"213 Ltr.":               "২১৩ লিটার",
		"220-240V/ 50Hz":         "২২০-২৪০ভি/ ৫০হার্টজ",
		"50Hz":                   "৫০হার্টজ",
		"RSCR":                   "আরএসসিআর",
		"Mechanical":             "মেকানিক্যাল",
		"Manual":                 "ম্যানুয়াল",
		"R600a":                  "আর৬০০এ",
		"Wire":                   "ওয়্যার",
		"2":                      "২",
		"3":                      "৩",
		"1":                      "১",
		"Yes":                    "হ্যাঁ",
		"557 x 645 x 1510 mm":    "৫৫৭ x ৬৪৫ x ১৫১০ মিমি",
		"45.5 ± 2 Kg":            "৪৫.৫ ± ২ কেজি",
		"5 Star":                 "৫ স্টার",
		"Replacement Guarantee: 1 Year, Main Parts (Compressor): 12 Years, Door: 3 Years, Spare Parts: 4 Years, After Sales Service: 5 Years": "রিপ্লেসমেন্ট গ্যারান্টি: ১ বছর, মূল অংশ (কম্প্রেসার): ১২ বছর, দরজা: ৩ বছর, স্পেয়ার পার্টস: ৪ বছর, আফটার সেলস সার্ভিস: ৫ বছর",
		"12": "১২",
		"Lock: Yes, Interior Lamp: Yes, Handle: Recessed/Grip, Capillary: Copper, Polyurethane foam blowing agent CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology], Refrigerator Compartment Shelves: Wire/2, Refrigerator Door Baskets: GPPS/3, Vegetable Crisper: Yes/1, Egg Tray: Yes/1-2, Freezer Compartment Shelves: Wire/2, Loading quantity: 102/102/50 (40HQ/40Ft/20Ft), Climatic Type: N~ST, Cooling Effect: Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C, Energy Rating: 5 Star, Recommended voltage stabilizer capacity: No need if within 145V-260V; 1000VA if outside range.": "লক: হ্যাঁ, ইন্টেরিয়র ল্যাম্প: হ্যাঁ, হ্যান্ডেল: রিসেসড/গ্রিপ, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], রেফ্রিজারেটর কম্পার্টমেন্ট শেল্ফস: ওয়্যার/২, রেফ্রিজারেটর দরজা বাস্কেটস: জিপিপিএস/৩, ভেজিটেবল ক্রিসপার: হ্যাঁ/১, এগ ট্রে: হ্যাঁ/১-২, ফ্রিজার কম্পার্টমেন্ট শেল্ফস: ওয়্যার/২, লোডিং কোয়ান্টিটি: ১০২/১০২/৫০ (৪০এইচকিউ/৪০এফটি/২০এফটি), ক্লাইমেটিক টাইপ: এন~এসটি, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮°C, রেফ্রিজারেটর ক্যাবিনেট ০°C টু +৫°C, এনার্জি রেটিং: ৫ স্টার, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ১৪৫ভি-২৬০ভি এর মধ্যে থাকলে প্রয়োজন নেই; বাইরে থাকলে ১০০০ভিএ।",
		"50 ± 2 Kg": "৫০ ± ২ কেজি",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Recessed/ Grip": "Recessed/ Grip",
		"220 ~ 240/ 50": "২২০ ~ ২৪০/ ৫০",
		"5 Star (BDS 1850:2012)": "৫ Star (BDS ১৮৫০:২০১২)",
		"1530": "১৫৩০",
		"Yes/1": "Yes/১",
		"Copper": "কপার",
		"RoHS Certified": "RoHS Certified",
		"645": "৬৪৫",
		"No": "না",
		"V 1101 - R600a V 1102 - R600a V 1301 - R600a V 1302 - R600a V 1401 - R600a V 1501 - R600a V 1601 - R600a": "V ১১০১ - R৬০০a V ১১০২ - R৬০০a V ১৩০১ - R৬০০a V ১৩০২ - R৬০০a V ১৪০১ - R৬০০a V ১৫০১ - R৬০০a V ১৬০১ - R৬০০a",
		"N~ST": "N~ST",
		"V 1101 - 102 V 1102 - 102 V 1301 - 108.6 V 1302 - 108.6 V 1401 - 102 V 1501 - 99.4 V 1601 - 108.6": "V ১১০১ - ১০২ V ১১০২ - ১০২ V ১৩০১ - ১০৮.৬ V ১৩০২ - ১০৮.৬ V ১৪০১ - ১০২ V ১৫০১ - ৯৯.৪ V ১৬০১ - ১০৮.৬",
		"V1101/ V1102/ V1301/ V1302/ V1401/ V1501/ V1601: No Need to use voltage stabilizer. If out of voltage range(145V-260V), then suggested voltage stabilizer capacity is 1000VA.": "V১১০১/ V১১০২/ V১৩০১/ V১৩০২/ V১৪০১/ V১৫০১/ V১৬০১: No Need to use voltage stabilizer. If out of voltage range(১৪৫V-২৬০V), then suggested voltage stabilizer capacity is ১০০০VA.",
		"Yes/1-2": "Yes/১-২",
		"CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "CycloPentane [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Wire/2": "Wire/২",
		"580": "৫৮০",
		"102/ 102 /50": "১০২/ ১০২ /৫০",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfa-2a3-gdeh-xx'
func (s *SpecificationSeederRefrigeratorMarcelMfa2a3GdehXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfa-2a3-gdeh-xx"
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
		"Gross Volume (Outer Dimension, Manufacturer declared)": "213 Ltr.",
		"Net Volume": "176 Ltr.",
		"Net Weight": "45.5 ± 2 Kg",
		"Gross Weight": "50 ± 2 Kg",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Rated Voltage/ Hz": "220 ~ 240/ 50",
		"Compressor Input Power (Watt)": "V 1101 - 102 V 1102 - 102 V 1301 - 108.6 V 1302 - 108.6 V 1401 - 102 V 1501 - 99.4 V 1601 - 108.6",
		"Compressor Type": "RSCR",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Energy rating": "5 Star (BDS 1850:2012)",
		"Temperature Control (Electronic/ Mechanical)": "Mechanical",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Reversible Door": "No",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "Yes",
		"Refrigerant": "V 1101 - R600a V 1102 - R600a V 1301 - R600a V 1302 - R600a V 1401 - R600a V 1501 - R600a V 1601 - R600a",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "CycloPentane [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "V1101/ V1102/ V1301/ V1302/ V1401/ V1501/ V1601: No Need to use voltage stabilizer. If out of voltage range(145V-260V), then suggested voltage stabilizer capacity is 1000VA.",
		"Shelf (Material/ No.)": "Wire/2",
		"Door Basket": "No",
		"Interior Lamp": "No",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray or Pocket": "Yes/1-2",
		"Can Storage Dispenser": "No",
		"Deodorizer": "No",
		"Shelf (Material/ No)": "Wire/2",
		"Drawer": "No",
		"Width/mm": "580",
		"Depth/mm": "645",
		"Height/mm": "1530",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "102/ 102 /50",
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
