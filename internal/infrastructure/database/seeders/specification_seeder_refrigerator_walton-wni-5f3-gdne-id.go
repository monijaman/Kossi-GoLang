package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWni5f3GdneId seeds specifications/options for product 'walton-wni-5f3-gdne-id'
type SpecificationSeederRefrigeratorWaltonWni5f3GdneId struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWni5f3GdneId creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWni5f3GdneId() *SpecificationSeederRefrigeratorWaltonWni5f3GdneId {
	return &SpecificationSeederRefrigeratorWaltonWni5f3GdneId{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wni-5f3-gdne-id"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWni5f3GdneId) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WCF-1B5-GDEL-XX":     "ডব্লিউসিএফ-১বি৫-জিডিইএল-এক্সএক্স",
		"Single Door":         "একটি দরজা",
		"150 Liters":          "১৫০ লিটার",
		"125 Liters":          "১২৫ লিটার",
		"25 Liters":           "২৫ লিটার",
		"4 Star":              "৪ তারা",
		"4":                   "৪",
		"200 kWh":             "২০০ কিলোওয়াট ঘণ্টা",
		"500 x 520 x 1200 mm": "৫০০ x ৫২০ x ১২০০ মিমি",
		"38 kg":               "৩৮ কেজি",
		"Silver":              "সিলভার",
		"Reciprocating":       "রিসিপ্রোকেটিং",
		"Direct Cool":         "ডাইরেক্ট কুল",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Toughened Glass":     "টাফেন্ড গ্লাস",
		"2":                   "২",
		"1":                   "১",
		"No":                  "না",
		"38 dB":               "৩৮ ডেসিবেল",
		"220V":                "২২০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"Large Vegetable Crisper, Egg Tray, Ice Cube Tray": "বড় শাকসবজি ক্রিস্পার, ডিমের ট্রে, বরফের কিউব ট্রে",
		"Refrigerant":         "রেফ্রিজারেন্ট",
		"Gross Volume":        "মোট ভলিউম",
		"Net Volume":          "নেট ভলিউম",
		"563 Liters":          "৫৬৩ লিটার",
		"No Frost":            "নো ফ্রস্ট",
		"865 x 725 x 1700 mm": "৮৬৫ x ৭২৫ x ১৭০০ মিমি",
		"103 kg":              "১০৩ কেজি",
		"BLDC Inverter":       "বিএলডিসি ইনভার্টার",
		"Automatic":           "অটোমেটিক",
		"Electronic":          "ইলেকট্রনিক",
		"Glass":               "গ্লাস",
		"5":                   "৫",
		"Yes":                 "হ্যাঁ",
		"220~240V":            "২২০~২৪০ভোল্ট",
		"5 Years":             "৫ বছর",
		"12":                  "১২",
		"R600a":               "আর৬০০এ",
		"501 Liters":          "৫০১ লিটার",
		"N/A":                 "এন/এ",
		"Digital Display, Real Tempered Glass Door, Wide Voltage Design, Intelligent Inverter Compressor, Dual Control": "ডিজিটাল ডিসপ্লে, রিয়েল টেম্পার্ড গ্লাস ডোর, ওয়াইড ভোল্টেজ ডিজাইন, ইন্টেলিজেন্ট ইনভার্টার কম্প্রেসার, ডুয়াল কন্ট্রোল",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wni-5f3-gdne-id'
func (s *SpecificationSeederRefrigeratorWaltonWni5f3GdneId) Seed(db *gorm.DB) error {
	productSlug := "walton-wni-5f3-gdne-id"
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
		"Model Name":                  "WNI-5F3-GDNE-ID",
		"Door Type":                   "Single Door",
		"Capacity":                    "563 Liters",
		"Refrigerator Capacity":       "N/A",
		"Freezer Capacity":            "N/A",
		"Energy Efficiency Rating":    "N/A",
		"Energy Star Rating":          "N/A",
		"Annual Energy Consumption":   "N/A",
		"Dimensions":                  "865 x 725 x 1700 mm",
		"Weight":                      "103 kg",
		"Color":                       "Silver",
		"Compressor Type":             "BLDC Inverter",
		"Cooling Technology":          "No Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Electronic",
		"Shelf Material":              "Glass",
		"Number of Shelves":           "5",
		"Door Bins":                   "4",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "N/A",
		"Voltage":                     "220~240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "12",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "563 Liters",
		"Net Volume":                  "501 Liters",
		"Special Features":            "Digital Display, Real Tempered Glass Door, Wide Voltage Design, Intelligent Inverter Compressor, Dual Control",
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
