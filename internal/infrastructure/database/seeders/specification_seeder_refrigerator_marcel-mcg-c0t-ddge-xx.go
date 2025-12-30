package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx seeds specifications/options for product 'marcel-mcg-c0t-ddge-xx'
type SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx() *SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx {
	return &SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcg-c0t-ddge-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mcg-c0t-ddge-xx": "মার্সেল-mcg-c0t-ddge-xx",
		"MCG-C0T-DDGE-XX":        "MCG-C0T-DDGE-XX",
		"Direct Cool":            "ডাইরেক্ট কুল",
		"300 Ltr.":               "৩০০ লিটার",
		"220-240V/ 50Hz":         "২২০-২৪০ভি/ ৫০হার্টজ",
		"50Hz":                   "৫০হার্টজ",
		"RSCR, CSIR":             "আরএসসিআর, সিএসআইআর",
		"Mechanical":             "মেকানিক্যাল",
		"Manual":                 "ম্যানুয়াল",
		"R600a":                  "আর৬০০এ",
		"Wire":                   "ওয়্যার",
		"2":                      "২",
		"1":                      "১",
		"No":                     "না",
		"1210 x 675 x 845 mm":    "১২১০ x ৬৭৫ x ৮৪৫ মিমি",
		"53.7 ± 2 Kg":            "৫৩.৭ ± ২ কেজি",
		"Replacement Guarantee: 1 Year (Condition Apply), Main Parts (Compressor): 12 Years, Door: 3 Years, Spare Parts: 4 Years, After Sales Service: 5 Years": "রিপ্লেসমেন্ট গ্যারান্টি: ১ বছর (শর্ত প্রযোজ্য), মূল অংশ (কম্প্রেসার): ১২ বছর, দরজা: ৩ বছর, স্পেয়ার পার্টস: ৪ বছর, আফটার সেলস সার্ভিস: ৫ বছর",
		"12": "১২",
		"Lock: No, Interior Lamp: No, Handle: Yes, Condenser: Steel, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Exterior Material: Painted Steel (PCM), Interior Material: Embossed Aluminium Sheet, Shelf: Wire/2, Basket: Wire/1, Sliding Glass: Yes, Loading quantity: 81/54/26 (40HQ/40Ft/20Ft), Climatic Type: N~ST, Wide voltage range: V 06.01, Cooling Effect: Freezer Cabinet Less than -18°C, Recommended voltage stabilizer capacity: 2000VA or More, Warranty Note: Does not cover damage due to accident, electrical fault, natural causes, negligence, improper installation, unauthorized modification, or unreadable serial numbers. Warranty period may be changed without notice. Covers only manufacturing defects subject to verification.": "লক: না, ইন্টেরিয়র ল্যাম্প: না, হ্যান্ডেল: হ্যাঁ, কনডেনসার: স্টিল, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], এক্সটেরিয়র ম্যাটেরিয়াল: পেইন্টেড স্টিল (পিসিএম), ইন্টেরিয়র ম্যাটেরিয়াল: এমবসড অ্যালুমিনিয়াম শীট, শেল্ফ: ওয়্যার/২, বাস্কেট: ওয়্যার/১, স্লাইডিং গ্লাস: হ্যাঁ, লোডিং কোয়ান্টিটি: ৮১/৫৪/২৬ (৪০এইচকিউ/৪০এফটি/২০এফটি), ক্লাইমেটিক টাইপ: এন~এসটি, ওয়াইড ভোল্টেজ রেঞ্জ: ভি ০৬.০১, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮°C, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ২০০০ভিএ বা মোর, ওয়ারেন্টি নোট: দুর্ঘটনা, বৈদ্যুতিক ত্রুটি, প্রাকৃতিক কারণ, অবহেলা, অনুপযুক্ত ইনস্টলেশন, অননুমোদিত পরিবর্তন, বা অপঠনযোগ্য সিরিয়াল নম্বরের কারণে ক্ষতি কভার করে না। ওয়ারেন্টি সময়কাল বিজ্ঞপ্তি ছাড়াই পরিবর্তিত হতে পারে। যাচাইকরণের সাপেক্ষে শুধুমাত্র উৎপাদন ত্রুটি কভার করে।",
	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcg-c0t-ddge-xx'
func (s *SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcg-c0t-ddge-xx"
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
		"Brand":                       "Marcel",
		"Model Name":                  "MCG-C0T-DDGE-XX",
		"Door Type":                   "Single Door",
		"Capacity":                    "300 Ltr.",
		"Freezer Capacity":            "300 Ltr.",
		"Gross Volume":                "300 Ltr.",
		"Net Volume":                  "300 Ltr.",
		"Dimensions":                  "1210 x 675 x 845 mm",
		"Weight":                      "53.7 ± 2 Kg",
		"Compressor Type":             "RSCR, CSIR",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Number of Shelves":           "2",
		"Voltage":                     "220-240V/ 50Hz",
		"Frequency (Hz)":              "50Hz",
		"Warranty":                    "Replacement Guarantee: 1 Year (Condition Apply), Main Parts (Compressor): 12 Years, Door: 3 Years, Spare Parts: 4 Years, After Sales Service: 5 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R600a",
		"Special Features":            "Lock: No, Interior Lamp: No, Handle: Yes, Condenser: Steel, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Exterior Material: Painted Steel (PCM), Interior Material: Embossed Aluminium Sheet, Shelf: Wire/2, Basket: Wire/1, Sliding Glass: Yes, Loading quantity: 81/54/26 (40HQ/40Ft/20Ft), Climatic Type: N~ST, Wide voltage range: V 06.01, Cooling Effect: Freezer Cabinet Less than -18°C, Recommended voltage stabilizer capacity: 2000VA or More, Warranty Note: Does not cover damage due to accident, electrical fault, natural causes, negligence, improper installation, unauthorized modification, or unreadable serial numbers. Warranty period may be changed without notice. Covers only manufacturing defects subject to verification.",
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
