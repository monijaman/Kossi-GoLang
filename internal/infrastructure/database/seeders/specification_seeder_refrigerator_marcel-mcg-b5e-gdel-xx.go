package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcgB5eGdelXx seeds specifications/options for product 'marcel-mcg-b5e-gdel-xx'
type SpecificationSeederRefrigeratorMarcelMcgB5eGdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcgB5eGdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcgB5eGdelXx() *SpecificationSeederRefrigeratorMarcelMcgB5eGdelXx {
	return &SpecificationSeederRefrigeratorMarcelMcgB5eGdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcg-b5e-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcgB5eGdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                  "মার্সেল",
		"marcel-mcg-b5e-gdel-xx":  "মার্সেল-mcg-b5e-gdel-xx",
		"MCG-B5E-GDEL-XX":         "MCG-B5E-GDEL-XX",
		"Direct Cool":             "ডাইরেক্ট কুল",
		"255 Ltr.":                "২৫৫ লিটার",
		"52±2 Kg":                 "৫২±২ কেজি",
		"220-240V/ 50Hz":          "২২০-২৪০ভি/ ৫০হার্টজ",
		"50Hz":                    "৫০হার্টজ",
		"RSIR, RSCR":              "আরএসআইআর, আরএসসিআর",
		"Mechanical, Electronics": "মেকানিক্যাল, ইলেকট্রনিক্স",
		"Manual":                  "ম্যানুয়াল",
		"R600a":                   "আর৬০০এ",
		"Wire":                    "ওয়্যার",
		"1":                       "১",
		"No":                      "না",
		"1085 x 685 x 865 mm":     "১০৮৫ x ৬৮৫ x ৮৬৫ মিমি",
		"Replacement Guarantee: 1 Year (Conditions Apply), Main Parts (Compressor): 12 Years, Door: 3 Years*, Spare Parts: 4 Years*, After Sales Service: 5 Years*": "রিপ্লেসমেন্ট গ্যারান্টি: ১ বছর (শর্ত প্রযোজ্য), মূল অংশ (কম্প্রেসার): ১২ বছর, দরজা: ৩ বছর*, স্পেয়ার পার্টস: ৪ বছর*, আফটার সেলস সার্ভিস: ৫ বছর*",
		"12": "১২",
		"Lock: Yes, Interior Lamp: Yes, Handle: Yes, Condenser: Steel, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Exterior Material: Painted Steel (PCM), Interior Material: Embossed Aluminium Sheet, Shelf: Wire/1, Basket: Plastic, Loading quantity: 32/64/96, Climatic Type: V.0301 N~ST, V.0302 N~ST, V.0401 N~T, Wide voltage range: V 03.01, Cooling Effect: Freezer Cabinet Less than -18℃, Recommended voltage stabilizer capacity: Not Required for 140V and above (most variants), Warranty Note: This warranty does not cover the following cases: 1. Any damage due to accident, electrical fault, natural causes, negligence or improper installation. 2. Any damage or failure caused by unauthorized modification or alteration. 3. Products with original serial numbers that have been removed, distorted or cannot be readily recognized.": "লক: হ্যাঁ, ইন্টেরিয়র ল্যাম্প: হ্যাঁ, হ্যান্ডেল: হ্যাঁ, কনডেনসার: স্টিল, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], এক্সটেরিয়র ম্যাটেরিয়াল: পেইন্টেড স্টিল (পিসিএম), ইন্টেরিয়র ম্যাটেরিয়াল: এমবসড অ্যালুমিনিয়াম শীট, শেল্ফ: ওয়্যার/১, বাস্কেট: প্লাস্টিক, লোডিং কোয়ান্টিটি: ৩২/৬৪/৯৬, ক্লাইমেটিক টাইপ: ভি.০৩০১ এন~এসটি, ভি.০৩০২ এন~এসটি, ভি.০৪০১ এন~টি, ওয়াইড ভোল্টেজ রেঞ্জ: ভি ০৩.০১, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮℃, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ১৪০ভি থেকে উপরে প্রয়োজন নেই (বেশিরভাগ ভেরিয়েন্ট), ওয়ারেন্টি নোট: এই ওয়ারেন্টি নিম্নলিখিত ক্ষেত্রে কভার করে না: ১. দুর্ঘটনা, বৈদ্যুতিক ত্রুটি, প্রাকৃতিক কারণ, অবহেলা বা অনুপযুক্ত ইনস্টলেশনের কারণে কোনো ক্ষতি। ২. অননুমোদিত পরিবর্তন বা পরিবর্ধনের কারণে কোনো ক্ষতি বা ব্যর্থতা। ৩. মূল সিরিয়াল নম্বর যেগুলি সরানো, বিকৃত বা সহজেই চেনা যায় না এমন পণ্য।",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcg-b5e-gdel-xx'
func (s *SpecificationSeederRefrigeratorMarcelMcgB5eGdelXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcg-b5e-gdel-xx"
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
        "Model Name":          "MCG-B5E-GDEL-XX",
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
