package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple seeds specifications/options for product 'vision-vsn-gd-refrigerator-re-240l-mirror-purple'
type SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple creates a new seeder instance
func NewSpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple() *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple {
	return &SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple{
		BaseSeeder: BaseSeeder{name: "Specifications for vision-vsn-gd-refrigerator-re-240l-mirror-purple"},
	}
}

func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Vision": "ভিশন",
		"VSN GD Refrigerator RE-240L Mirror Purple": "ভিএসএন জিডি রেফ্রিজারেটর আরই-২৪০এল মিরর পার্পল",
		"Single Door":         "একক দরজা",
		"240 Liters":          "২৪০ লিটার",
		"118 Liters":          "১১৮ লিটার",
		"122 Liters":          "১২২ লিটার",
		"5 Star":              "৫ তারা",
		"5":                   "৫",
		"295 kWh":             "২৯৫ কিলোওয়াট ঘণ্টা",
		"577 x 590 x 1572 mm": "৫৭৭ x ৫৯০ x ১৫৭২ মিমি",
		"62.40 kg":            "৬২.৪০ কেজি",
		"Mirror Purple":       "মিরর পার্পল",
		"Direct Cool":         "সরাসরি কুলিং",
		"Manual":              "ম্যানুয়াল",
		"Mechanical":          "মেকানিক্যাল",
		"Wire":                "ওয়্যার",
		"3":                   "৩",
		"2":                   "২",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"40 dB":               "৪০ ডেসিবেল",
		"160 ~ 260V":          "১৬০ ~ ২৬০ ভোল্ট",
		"50":                  "৫০",
		"5 Years":             "৫ বছর",
		"10":                  "১০",
		"R600a":               "আর৬০০এ",
		"Attractive glass door design, Royal modern color, High efficient cooling system, Freezer box free, Huge inside space, Big freezer compartment, Low noise compressor, Lock and key facility, Interior LED light, 100% copper condenser": "আকর্ষণীয় গ্লাস দরজা ডিজাইন, রয়েল মডার্ন কালার, উচ্চ কার্যকর কুলিং সিস্টেম, ফ্রিজার বক্স ফ্রি, বিশাল ইনসাইড স্পেস, বড় ফ্রিজার কম্পার্টমেন্ট, লো নয়েজ কম্প্রেসর, লক এবং কী ফ্যাসিলিটি, ইন্টেরিয়র এলইডি লাইট, ১০০% কপার কন্ডেনসার",
		"Special Features": "বিশেষ বৈশিষ্ট্য",
		"Refrigerant":      "রেফ্রিজারেন্ট",
		"Gross Volume":     "মোট ভলিউম",
		"Net Volume":       "নেট ভলিউম",
		"BLDC Inverter":    "বিএলডিসি ইনভার্টার",
	}
}

// Seed inserts specification records for the product identified by slug 'vision-vsn-gd-refrigerator-re-240l-mirror-purple'
func (s *SpecificationSeederRefrigeratorVisionVSNGDRefrigeratorRE240LMirrorPurple) Seed(db *gorm.DB) error {
	productSlug := "vision-vsn-gd-refrigerator-re-240l-mirror-purple"
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
		"Model Name":                  "VSN GD Refrigerator RE-240L Mirror Purple",
		"Door Type":                   "Single Door",
		"Capacity":                    "240 Liters",
		"Refrigerator Capacity":       "122 Liters",
		"Freezer Capacity":            "118 Liters",
		"Energy Efficiency Rating":    "5 Star",
		"Energy Star Rating":          "5",
		"Annual Energy Consumption":   "295 kWh",
		"Dimensions":                  "577 x 590 x 1572 mm",
		"Weight":                      "62.40 kg",
		"Color":                       "Mirror Purple",
		"Compressor Type":             "Normal",
		"Cooling Technology":          "Direct Cool",
		"Defrost Type":                "Manual",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "Wire",
		"Number of Shelves":           "3",
		"Door Bins":                   "2",
		"Crisper Drawers":             "4",
		"Ice Maker":                   "No",
		"Water Dispenser":             "No",
		"Noise Level":                 "40 dB",
		"Voltage":                     "160 ~ 260V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "5 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Gross Volume":                "240 Liters",
		"Net Volume":                  "240 Liters",
		"Special Features":            "Attractive glass door design, Royal modern color, High efficient cooling system, Freezer box free, Huge inside space, Big freezer compartment, Low noise compressor, Lock and key facility, Interior LED light, 100% copper condenser",
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
