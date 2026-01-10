package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMniF1nGdneWd seeds specifications/options for product 'marcel-mni-f1n-gdne-wd'
type SpecificationSeederRefrigeratorMarcelMniF1nGdneWd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMniF1nGdneWd creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMniF1nGdneWd() *SpecificationSeederRefrigeratorMarcelMniF1nGdneWd {
	return &SpecificationSeederRefrigeratorMarcelMniF1nGdneWd{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mni-f1n-gdne-wd"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMniF1nGdneWd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mni-f1n-gdne-wd": "মার্সেল-mni-f1n-gdne-wd",
		"MNI-F1N-GDNE-WD":        "MNI-F1N-GDNE-WD",
		// Add more translations as needed
		"775 mm": "৭৭৫ মিমি",
		"GLASS/5": "GLASS/৫",
		"Steel": "স্টীল",
		"Yes": "হ্যাঁ",
		"220~240 V/ 50 Hz": "২২০~২৪০ V/ ৫০ Hz",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"1885 mm": "১৮৮৫ মিমি",
		"Electronic": "ইলেকট্রনিক",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"125 ± 2 Kg": "১২৫ ± ২ কেজি",
		"BLDC Inverter": "বিএলডিসি ইনভার্টার",
		"Yes/1": "Yes/১",
		"900 mm": "৯০০ মিমি",
		"Copper": "কপার",
		"No": "না",
		"No-Frost": "নো ফ্রস্ট",
		"591": "৫৯১",
		"39/39/18 (Vertical Loading)": "৩৯/৩৯/১৮ (Vertical Loading)",
		"Built In": "বিল্ট ইন",
		"45.4 - 197": "৪৫.৪ - ১৯৭",
		"Glass/5": "Glass/৫",
		"Yes/ 1": "Yes/ ১",
		"Automatic": "স্বয়ংক্রিয়",
		"115 ± 2 Kg": "১১৫ ± ২ কেজি",
		"N~T": "N~T",
		"GPPS/4": "GPPS/৪",
		"548": "৫৪৮",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mni-f1n-gdne-wd'
func (s *SpecificationSeederRefrigeratorMarcelMniF1nGdneWd) Seed(db *gorm.DB) error {
	productSlug := "marcel-mni-f1n-gdne-wd"
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
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Bottle Pocket": "GPPS/4",
		"Capillary": "Copper",
		"Climatic Type (SN, N, ST, T)": "N~T",
		"Compressor": "BLDC Inverter",
		"Compressor Input Power (watt)": "45.4 - 197",
		"Condenser": "Steel",
		"Cooling Effect": "Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Depth/mm": "725 mm",
		"Drawer": "Yes/1",
		"Egg Tray or Pocket": "Yes/ 1",
		"Gross Volume": "591",
		"Gross Weight": "125 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Built In",
		"Height/mm": "1780 mm",
		"Ice Box": "Yes/1",
		"Ice Case": "Yes/1",
		"Interior LED Lamp": "Yes",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/40Ft/20Ft": "39/39/18 (Vertical Loading)",
		"Lock": "No",
		"Net Volume": "548",
		"Net Weight": "115 ± 2 Kg",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Rated Voltage/ Hz": "220~240 V/ 50 Hz",
		"Shelf (Material/ No.)": "GLASS/5",
		"Shelf (Material/No.)": "Glass/5",
		"Temperature Control (Electronic/  Mechanical)": "Electronic",
		"Type": "No-Frost",
		"Vegetable Crisper": "Yes/1",
		"Water Dispenser": "Yes",
		"Width/mm": "865 mm",
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
