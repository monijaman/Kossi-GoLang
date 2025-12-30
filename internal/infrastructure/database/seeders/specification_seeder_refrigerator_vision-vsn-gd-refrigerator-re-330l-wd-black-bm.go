package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM seeds specifications/options for product 'vision-vsn-gd-refrigerator-re-330l-wd-black-bm'
type SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM() *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM {
	return &SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vsn-gd-refrigerator-re-330l-wd-black-bm"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VSN GD Refrigerator RE-330L WD Black-BM": "ভিএসএন জিডি রেফ্রিজারেটর আরই-৩৩০এল ডব্লিউডি ব্ল্যাক-বিএম",
		"Single Door":         "একক দরজা",
		"330 Liters":          "৩৩০ লিটার",
		"0 Liters":            "০ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"295 kWh":             "২৯৫ কিলোওয়াট ঘণ্টা",
		"660 x 604 x 1898 mm": "৬৬০ x ৬০৪ x ১৮৯৮ মিমি",
		"76.8 kg":             "৭৬.৮ কেজি",
		"WD Black-BM":         "ডব্লিউডি ব্ল্যাক-বিএম",
		"Direct Cool":         "সরাসরি কুলিং",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Wire":                "ওয়্যার",
		"2":                   "২",
		"3":                   "৩",
		"4":                   "৪",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ ভোল্ট",
		"50":                  "৫০",
		"5 Years":             "৫ বছর",
		"10":                  "১০",
		"R600a":               "আর৬০০এ",
		"Water dispenser with Refrigerator, Fabulous ultra-Modern Glass door Colour, Very fast cooling speed, hygienic clean air, Anti-Bacterial gasket, Vast of storage capacity, Virgin & food grade Plastic liner, Low noise compressor, with lock and key, Interior LED light, 100% copper condenser": "রেফ্রিজারেটরের সাথে ওয়াটার ডিসপেন্সার, ফ্যাবুলাস আল্ট্রা-মডার্ন গ্লাস দরজা কালার, খুব দ্রুত কুলিং গতি, হাইজেনিক ক্লিন এয়ার, অ্যান্টি-ব্যাকটেরিয়াল গ্যাসকেট, বিশাল স্টোরেজ ক্যাপাসিটি, ভার্জিন এবং ফুড গ্রেড প্লাস্টিক লাইনার, লো নয়েজ কম্প্রেসর, লক এবং কী সহ, ইন্টেরিয়র এলইডি লাইট, ১০০% কপার কন্ডেনসার",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"BLDC Inverter":    "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vsn-gd-refrigerator-re-330l-wd-black-bm'
func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE330LWDBlackBM) Seed(db *gorm.DB) error {
	productSlug := "vision-vsn-gd-refrigerator-re-330l-wd-black-bm"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
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
		"Brand":                       "Vision",
		"Model Name":                  "VSN GD Refrigerator RE-330L WD Black-BM",
		"Door Type":                   "Single Door",
		"Capacity":                    "330 Liters",
		"Refrigerator Capacity":       "330 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "660 x 604 x 1898 mm",
		"Weight":                      "76.8 kg",
		"Color":                       "WD Black-BM",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "3",
		"Door Bins":                   "2",
		"Crisper Drawers":             "4",
		"Ice Maker":                   "No",
		"Water Dispenser":             "Yes",
		"Noise Level":                 "40 dB",
		"Voltage":                     "160 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "345 Liters",
		"Net Volume":                  "330 Liters",
		"Special Features":            "Water dispenser with Refrigerator, Fabulous ultra-Modern Glass door Colour, Very fast cooling speed, hygienic clean air, Anti-Bacterial gasket, Vast of storage capacity, Virgin & food grade Plastic liner, Low noise compressor, with lock and key, Interior LED light, 100% copper condenser",
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
				if err := db.Create(sModel).Last(&sModel).Error; err != nil {
					return err
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
