package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorKonkaKrt200gb seeds specifications/options for product 'konka-krt-200gb'
type SpecificationSeederRefrigeratorKonkaKrt200gb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorKonkaKrt200gb creates a new seeder instance
func NewSpecificationSeederRefrigeratorKonkaKrt200gb() *SpecificationSeederRefrigeratorKonkaKrt200gb {
	return &SpecificationSeederRefrigeratorKonkaKrt200gb{
		BaseSeeder: BaseSeeder{name: "Specifications for konka-krt-200gb"},
	}
}

func (s *SpecificationSeederRefrigeratorKonkaKrt200gb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Konka":              "কনকা",
		"KRT-200GB":          "KRT-200GB",
		"China":              "চীন",
		"Direct Cool":        "ডিরেক্ট কুল",
		"No Frost":           "নো ফ্রস্ট",
		"Yes":                "হ্যাঁ",
		"No":                 "না",
		"Manual":             "ম্যানুয়াল",
		"Automatic":          "স্বয়ংক্রিয়",
		"Mechanical":         "মেকানিক্যাল",
		"Digital":            "ডিজিটাল",
		"R600a":              "R600a",
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Inverter":           "ইনভার্টার",
		"Conventional":       "প্রচলিত",
		"Top Mount":          "টপ মাউন্ট",
		"Bottom Mount":       "বটম মাউন্ট",
		"Single Door":        "সিঙ্গেল ডোর",
		"Double Door":        "ডাবল ডোর",
		"Side by Side":       "সাইড বাই সাইড",
		"Cross Door":         "ক্রস ডোর",
		"French Door":        "ফ্রেঞ্চ ডোর",
		"Multi-Door":         "মাল্টি-ডোর",
		"Tempered Glass":     "টেম্পার্ড গ্লাস",
		"Wire":               "ওয়্যার",
		"Glass":              "গ্লাস",
		"LED":                "LED",
		"1 Year":             "১ বছর",
		"2 Years":            "২ বছর",
		"5 Years":            "৫ বছর",
		"10 Years":           "১০ বছর",
		"Copper":             "তামা",
		"Smart":              "স্মার্ট",
		"WiFi":               "ওয়াইফাই",
		"N/A":                "N/A",
		"N~ST":               "N~ST",
		"RoHS Certified":     "RoHS সার্টিফাইড",
		"Cyclopentane":       "সাইক্লোপেন্টেন",
		"Recessed/ Grip":     "রিসেসড/গ্রিপ",
		"GPPS/3":             "GPPS/৩",
		"GPPS/4":             "GPPS/৪",
		"HIPS/2":             "HIPS/২",
		"HIPS/3":             "HIPS/৩",
		"Wire/2":             "Wire/২",
		"Wire/3":             "Wire/৩",
		"Wire/4":             "Wire/৪",
		"Tempered Glass/3":   "টেম্পার্ড গ্লাস/৩",
		"Tempered Glass/4":   "টেম্পার্ড গ্লাস/৪",
		"2":                  "২",
		"3":                  "৩",
		"4":                  "৪",
		"50":                 "৫০",
		"Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C":        "ফ্রিজার ক্যাবিনেট -১৮°C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০°C থেকে +৫°C",
		"Compressor: 10 Years, Spare Parts: 2 Years, After Sales Service: 5 Years": "কম্প্রেসর: ১০ বছর, খুচরা যন্ত্রাংশ: ২ বছর, বিক্রয়োত্তর সেবা: ৫ বছর",
		"200 Ltr.":    "২০০ Ltr.",
		"Black Daisy": "ব্ল্যাক ডেইজি",
		"BSTI 5 Star": "BSTI ৫ স্টার",
		"125V~265V":   "১২৫V~২৬৫V",
		"Supreme Freezing Capacity, BSTI 5 Star Energy Rating, Stylist Interior Decoration with LED Lighting, Wide Voltage range (125V~265V), Real Tempered Glass Door. Shiny, Colorful & Scratch Resistance, European Standard Anti-Fungal Door gasket, Large Moist Fresh Zone (Vegetable Box) with humidity adjust Control, Active Carbon Deodorizer Ensure odorless, fresh and pure air in the Refrigerator, Optimal humidity for lasting freshness, Garden Fresh Technology, Anti-microbial Technology (Nano- Sliver Technology), Interior parts are 100% Food Grade, Healthy, Fresh and Protect Vitamin, R600a Refrigerant Gas, 100% CFC and HCFC Free, Eco Friendly refrigerant, No destructive effect to the ozone sphere, The operating pressure of R600a is not high, so the noisy is very low, DEC System and Faster Cooling": "সুপ্রিম ফ্রিজিং ক্যাপাসিটি, BSTI ৫ স্টার এনার্জি রেটিং, স্টাইলিস্ট ইন্টেরিয়র ডেকোরেশন LED লাইটিং সহ, ওয়াইড ভোল্টেজ রেঞ্জ (১২৫V~২৬৫V), রিয়েল টেম্পার্ড গ্লাস ডোর। শাইনি, কালারফুল এবং স্ক্র্যাচ রেসিস্ট্যান্স, ইউরোপিয়ান স্ট্যান্ডার্ড অ্যান্টি-ফাঙ্গাল ডোর গ্যাসকেট, লার্জ ময়েস্ট ফ্রেশ জোন (ভেজিটেবল বক্স) হিউমিডিটি অ্যাডজাস্ট কন্ট্রোল সহ, অ্যাকটিভ কার্বন ডিওডোরাইজার রেফ্রিজারেটরে গন্ধহীন, তাজা এবং বিশুদ্ধ বায়ু নিশ্চিত করে, অপ্টিমাল হিউমিডিটি লাস্টিং ফ্রেশনেসের জন্য, গার্ডেন ফ্রেশ টেকনোলজি, অ্যান্টি-মাইক্রোবিয়াল টেকনোলজি (ন্যানো-স্লিভার টেকনোলজি), ইন্টেরিয়র পার্টস ১০০% ফুড গ্রেড, স্বাস্থ্যকর, তাজা এবং ভিটামিন রক্ষা করে, R600a রেফ্রিজারেন্ট গ্যাস, ১০০% CFC এবং HCFC ফ্রি, ইকো ফ্রেন্ডলি রেফ্রিজারেন্ট, ওজোন স্ফিয়ারে কোনো ধ্বংসাত্মক প্রভাব নেই, R600a এর অপারেটিং প্রেশার উচ্চ নয়, তাই নয়েজ খুব কম, DEC সিস্টেম এবং দ্রুত কুলিং",
	}
}

// Seed inserts specification records for the product identified by slug 'konka-krt-200gb'
func (s *SpecificationSeederRefrigeratorKonkaKrt200gb) Seed(db *gorm.DB) error {
	productSlug := "konka-krt-200gb"
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
		"Origin":                            318,
		"Gross Weight":                      728,
		"Drawer":                            729,
		"Energy Rating":                     730,
		"Shelf (Material/No.)":              731,
	}

	specs := map[string]string{
		"Brand":                           "Konka",
		"Model Name":                      "KRT-200GB-Black Daisy-2-Door, Upper Freezer (200 Ltr)",
		"Origin":                          "China",
		"Type":                            "Top Mount",
		"Door Type":                       "Double Door",
		"Compressor Type":                 "Conventional",
		"Cooling Technology":              "Direct Cool",
		"Defrost Type":                    "Manual",
		"Defrosting (Automatic/ Manual)":  "Manual",
		"Temperature Control":             "Mechanical",
		"Shelf Material":                  "Wire",
		"Shelf (Material/No.)":            "Wire/2",
		"Number of Shelves":               "2",
		"Door Basket":                     "GPPS/3",
		"Vegetable Crisper":               "Yes",
		"Vegetable Crisper Cover":         "Yes",
		"Egg Tray or Pocket":              "Yes",
		"Interior Lamp":                   "Yes",
		"Ice Maker":                       "No",
		"Handle (Recessed/ Grip)":         "Recessed/ Grip",
		"Lock":                            "Yes",
		"Reversible Door":                 "Yes",
		"Capillary":                       "Copper",
		"Polyurethane Foam Blowing Agent": "Cyclopentane",
		"Climate Type (SN, N, ST, T)":     "N~ST",
		"Thermostat":                      "RoHS Certified",
		"Cooling Effect":                  "Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C",
		"Drawer":                          "HIPS/2",
		"Warranty":                        "Compressor: 10 Years, Spare Parts: 2 Years, After Sales Service: 5 Years",
		"Compressor Warranty (Years)":     "10 Years",
		"Gross Volume":                    "200 Ltr.",
		"Refrigerant":                     "R600a",
		"Voltage":                         "125V~265V",
		"Frequency (Hz)":                  "50",
		"Color":                           "Black Daisy",
		"Energy Efficiency Rating":        "BSTI 5 Star",
		"Special Features":                "Supreme Freezing Capacity, BSTI 5 Star Energy Rating, Stylist Interior Decoration with LED Lighting, Wide Voltage range (125V~265V), Real Tempered Glass Door. Shiny, Colorful & Scratch Resistance, European Standard Anti-Fungal Door gasket, Large Moist Fresh Zone (Vegetable Box) with humidity adjust Control, Active Carbon Deodorizer Ensure odorless, fresh and pure air in the Refrigerator, Optimal humidity for lasting freshness, Garden Fresh Technology, Anti-microbial Technology (Nano- Sliver Technology), Interior parts are 100% Food Grade, Healthy, Fresh and Protect Vitamin, R600a Refrigerant Gas, 100% CFC and HCFC Free, Eco Friendly refrigerant, No destructive effect to the ozone sphere, The operating pressure of R600a is not high, so the noisy is very low, DEC System and Faster Cooling",
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
