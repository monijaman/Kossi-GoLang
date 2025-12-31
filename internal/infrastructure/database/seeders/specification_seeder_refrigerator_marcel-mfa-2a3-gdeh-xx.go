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
        "Brand":               "Marcel",
        "Model Name":          "MFA-2A3-GDEH-XX",
        "Cooling Technology":  "Direct Cool",
        "Gross Volume":        "177 Ltr.",
        "Net Volume":          "175 Ltr.",
        "Weight":              "50 ± 2 Kg",
        "Refrigerant":         "R600a",
        "Temperature Control": "Mechanical",
        "Voltage":             "220 ~ 240",
        "Dimensions":          "555 x 630 x 1410 mm",
        "Packing Dimensions":  "580 x 645 x 1455 mm",
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
