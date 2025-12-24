package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWnm2a7GdelXx seeds specifications/options for product 'walton-wnm-2a7-gdel-xx'
type SpecificationSeederRefrigeratorWaltonWnm2a7GdelXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWnm2a7GdelXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWnm2a7GdelXx() *SpecificationSeederRefrigeratorWaltonWnm2a7GdelXx {
	return &SpecificationSeederRefrigeratorWaltonWnm2a7GdelXx{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wnm-2a7-gdel-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWnm2a7GdelXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WNM-2A7-GDEL-XX":     "ডব্লিউএনএম-২এ৭-জিডিইএল-এক্সএক্স",
		"No-Frost":            "নো ফ্রস্ট",
		"217 Liters":          "২১৭ লিটার",
		"183 Liters":          "১৮৩ লিটার",
		"555 x 651 x 1475 mm": "৫৫৫ x ৬৫১ x ১৪৭৫ মিমি",
		"52 kg":               "৫২ কেজি",
		"Silver":              "সিলভার",
		"BLDC Inverter":       "বিএলডিসি ইনভার্টার",
		"Automatic":           "অটোমেটিক",
		"Electronic":          "ইলেকট্রনিক",
		"GPPS, Wire, Glass":   "জিপিপিএস, ওয়াইয়ার, গ্লাস",
		"3":                   "৩",
		"4":                   "৪",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"220-240V":            "২২০-২৪০ভোল্ট",
		"50":                  "৫০",
		"1 Year":              "১ বছর",
		"12":                  "১২",
		"R600a (CFC Free)":    "আর৬০০এ (সিএফসি ফ্রি)",
		"Chilled Room, Egg Tray, Vegetable Crisper, Ice Tray, Ice Box": "চিলড রুম, ডিমের ট্রে, শাকসবজি ক্রিস্পার, বরফের ট্রে, বরফের বক্স",
		"Refrigerant":         "রেফ্রিজারেন্ট",
		"Gross Volume":        "মোট ভলিউম",
		"Net Volume":          "নেট ভলিউম",
		"T":                   "টি",
		"220-240/ 50 Hz/ 88W": "২২০-২৪০/ ৫০ হার্টজ/ ৮৮ওয়াট",
		"Steel":               "স্টিল",
		"Copper":              "কপার",
		"RoHS Certified":      "রোএইচএস সার্টিফাইড",
		"Cyclopentene":        "সাইক্লোপেন্টিন",
		"Eco-friendly (100% CFC & HCFC Free) Green Technology": "ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রীন টেকনোলজি",
		"GPPS/1, Wire/2": "জিপিপিএস/১, ওয়াইয়ার/২",
		"Glass/1":        "গ্লাস/১",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"600 x 703 x 1550 mm": "৬০০ x ৭০৩ x ১৫৫০ মিমি",
		"57 ± 2 Kg":           "৫৭ ± ২ কেজি",
		"83/ 83/ 39":          "৮৩/ ৮৩/ ৩৯",
		"0 Liters":            "০ লিটার",
		"Eco-friendly (100% CFC & HCFC Free) Green Technology, Chilled Room, Egg Tray, Vegetable Crisper, Ice Tray, Ice Box, Interior Lamp": "ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রীন টেকনোলজি, চিলড রুম, ডিমের ট্রে, শাকসবজি ক্রিস্পার, বরফের ট্রে, বরফের বক্স, ইন্টেরিয়র ল্যাম্প",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wnm-2a7-gdel-xx'
func (s *SpecificationSeederRefrigeratorWaltonWnm2a7GdelXx) Seed(db *gorm.DB) error {
	productSlug := "walton-wnm-2a7-gdel-xx"
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
		"Brand":                       "Walton",
		"Model Name":                  "WNM-2A7-GDEL-XX",
		"Door Type":                   "No-Frost",
		"Capacity":                    "183 Liters",
		"Refrigerator Capacity":       "183 Liters",
		"Freezer Capacity":            "0 Liters",
		"Energy Efficiency Rating":    "",
		"Energy Star Rating":          "",
		"Annual Energy Consumption":   "",
		"Dimensions":                  "555 x 651 x 1475 mm",
		"Weight":                      "52 kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "No-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "GPPS, Wire, Glass",
		"Number of Shelves":           "4",
		"Door Bins":                   "4",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "",
		"Voltage":                     "220-240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "1 Year",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R600a (CFC Free)",
		"Gross Volume":                "217 Liters",
		"Net Volume":                  "183 Liters",
		"Special Features":            "Eco-friendly (100% CFC & HCFC Free) Green Technology, Chilled Room, Egg Tray, Vegetable Crisper, Ice Tray, Ice Box, Interior Lamp",
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
