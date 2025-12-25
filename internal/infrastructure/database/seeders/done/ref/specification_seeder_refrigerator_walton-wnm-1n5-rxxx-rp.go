package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp seeds specifications/options for product 'walton-wnm-1n5-rxxx-rp'
type SpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp creates a new seeder instance
func NewSpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp() *SpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp {
	return &SpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp{
		BaseSeeder: BaseSeeder{name: "Specifications for walton-wnm-1n5-rxxx-rp"},
	}
}

func (s *SpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":              "ওয়ালটন",
		"WNM-1N5-RXXX-RP":     "ডব্লিউএনএম-১এন৫-আরএক্সএক্সএক্স-আরপি",
		"No-Frost":            "নো ফ্রস্ট",
		"195 Liters":          "১৯৫ লিটার",
		"551 x 665 x 1373 mm": "৫৫১ x ৬৬৫ x ১৩৭৩ মিমি",
		"48 kg":               "৪৮ কেজি",
		"Silver":              "সিলভার",
		"RSCR":                "আরএসসিআর",
		"Automatic":           "অটোমেটিক",
		"Mechanical":          "মেকানিক্যাল",
		"GPPS, Wire":          "জিপিপিএস, ওয়াইয়ার",
		"2":                   "২",
		"1":                   "১",
		"Yes":                 "হ্যাঁ",
		"No":                  "না",
		"220-240V":            "২২০-২৪০ভোল্ট",
		"50":                  "৫০",
		"2 Years":             "২ বছর",
		"10":                  "১০",
		"R600a (CFC Free)":    "আর৬০০এ (সিএফসি ফ্রি)",
		"167 Liters":          "১৬৭ লিটার",
		"No need to use Voltage Stabilizer, Elegant door provides a contemporary look, Ion Anti-bacteria keeps your refrigerator germ-free, Clear back with built-in condenser, Power cooler, dynamic flow, ice tray, Reinforce glass shelves(wire shelves available), Transparent chilled room, Door panel & door liner foam together, Optimum noise level, Longer enduring cooling system, CFC free:R600a, HCFC free:Cyclopentane, Key & lock (optional)": "ভোল্টেজ স্ট্যাবিলাইজার ব্যবহারের প্রয়োজন নেই, এলিগ্যান্ট দরজা আধুনিক লুক প্রদান করে, আয়ন অ্যান্টি-ব্যাকটেরিয়া আপনার রেফ্রিজারেটরকে জার্ম-ফ্রি রাখে, বিল্ট-ইন কন্ডেন্সার সহ ক্লিয়ার ব্যাক, পাওয়ার কুলার, ডায়নামিক ফ্লো, আইস ট্রে, রিইনফোর্স গ্লাস শেল্ফস (ওয়াইয়ার শেল্ফস উপলব্ধ), ট্রান্সপারেন্ট চিলড রুম, দরজা প্যানেল এবং দরজা লাইনার ফোম একসাথে, অপটিমাম নয়েজ লেভেল, দীর্ঘস্থায়ী কুলিং সিস্টেম, সিএফসি ফ্রি: আর৬০০এ, এইচসিএফসি ফ্রি: সাইক্লোপেন্টেন, কী এবং লক (অপশনাল)",
		"Refrigerant":  "রেফ্রিজারেন্ট",
		"Gross Volume": "মোট ভলিউম",
		"Net Volume":   "নেট ভলিউম",
	}
}

// Seed inserts specification records for the product identified by slug 'walton-wnm-1n5-rxxx-rp'
func (s *SpecificationSeederRefrigeratorWaltonWnm1n5RxxxRp) Seed(db *gorm.DB) error {
	productSlug := "walton-wnm-1n5-rxxx-rp"
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
		"Model Name":                  "WNM-1N5-RXXX-RP",
		"Door Type":                   "No-Frost",
		"Capacity":                    "195 Liters",
		"Refrigerator Capacity":       "",
		"Freezer Capacity":            "",
		"Energy Efficiency Rating":    "",
		"Energy Star Rating":          "",
		"Annual Energy Consumption":   "",
		"Dimensions":                  "551 x 665 x 1373 mm",
		"Weight":                      "48 kg",
		"Color":                       "Silver",
		"Compressor Type":             "RSCR",
		"Cooling Technology":          "No-Frost",
		"Defrost Type":                "Automatic",
		"Temperature Control":         "Mechanical",
		"Shelf Material":              "GPPS, Wire",
		"Number of Shelves":           "2",
		"Door Bins":                   "2",
		"Crisper Drawers":             "1",
		"Ice Maker":                   "Yes",
		"Water Dispenser":             "No",
		"Noise Level":                 "",
		"Voltage":                     "220-240V",
		"Frequency (Hz)":              "50",
		"App Control":                 "No",
		"Voice Assistant Support":     "No",
		"Warranty":                    "2 Years",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a (CFC Free)",
		"Gross Volume":                "195 Liters",
		"Net Volume":                  "167 Liters",
		"Special Features":            "No need to use Voltage Stabilizer, Elegant door provides a contemporary look, Ion Anti-bacteria keeps your refrigerator germ-free, Clear back with built-in condenser, Power cooler, dynamic flow, ice tray, Reinforce glass shelves(wire shelves available), Transparent chilled room, Door panel & door liner foam together, Optimum noise level, Longer enduring cooling system, CFC free:R600a, HCFC free:Cyclopentane, Key & lock (optional)",
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
