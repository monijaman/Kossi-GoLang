package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMniE6cGdneDd seeds specifications/options for product 'marcel-mni-e6c-gdne-dd'
type SpecificationSeederRefrigeratorMarcelMniE6cGdneDd struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMniE6cGdneDd creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMniE6cGdneDd() *SpecificationSeederRefrigeratorMarcelMniE6cGdneDd {
	return &SpecificationSeederRefrigeratorMarcelMniE6cGdneDd{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mni-e6c-gdne-dd"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMniE6cGdneDd) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mni-e6c-gdne-dd": "মার্সেল-mni-e6c-gdne-dd",
		"MNI-E6C-GDNE-DD":        "MNI-E6C-GDNE-DD",
		"No-Frost":               "নো ফ্রস্ট",
		"No Frost":               "নো ফ্রস্ট",
		"BLDC Inverter":          "বিএলডিসি ইনভার্টার",
		"Electronic":             "ইলেকট্রনিক",
		"Automatic":              "স্বয়ংক্রিয়",
		"Steel":                  "স্টীল",
		"Copper":                 "কপার",
		"Built In":               "বিল্ট ইন",
		"Cyclopentene":           "সাইকলোপেন্টেন",
		"Glass":                  "গ্লাস",
		"GPPS/4":                 "GPPS/4",
		"GPPS/5":                 "GPPS/5",
		"Yes":                    "হ্যাঁ",
		"No":                     "না",
		"Freezer < -18℃; Refrigerator 0℃ to +5℃": "ফ্রিজার: < -18℃; ফ্রিজ: 0℃ থেকে +5℃",
		"563 Ltr.":   "৫৬৩ লিটার",
		"501 Ltr.":   "৫০১ লিটার",
		"103 ± 2 Kg": "১০৩ ± ২ কেজি",
		"113 ± 2 Kg": "১১৩ ± ২ কেজি",
		"865 mm":     "৮৬৫ মিমি",
		"725 mm":     "৭২৫ মিমি",
		"1700 mm":    "১৭০০ মিমি",
		"900 mm":     "৯০০ মিমি",
		"775 mm":     "৭৭৫ মিমি",
		"1815 mm":    "১৮১৫ মিমি",
		"5":          "৫",
		"1":          "১",
		"39/39/19":   "৩৯/৩৯/১৯",
		"220~240 V/ 50 Hz": "২২০~২৪০ V/ ৫০ Hz",
		"Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃": "Freezer Cabinet Less than -১৮℃ Refrigerator Cabinet ০℃ to +৫℃",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"Yes/1": "Yes/১",
		"45.4 - 197": "৪৫.৪ - ১৯৭",
		"39/39/19 (Vertical Loading)": "৩৯/৩৯/১৯ (Vertical Loading)",
		"Glass/5": "Glass/৫",
		"Yes/ 1": "Yes/ ১",
		"N~T": "N~T",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mni-e6c-gdne-dd'
func (s *SpecificationSeederRefrigeratorMarcelMniE6cGdneDd) Seed(db *gorm.DB) error {
	productSlug := "marcel-mni-e6c-gdne-dd"
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
		"Type": "No-Frost",
		"Gross Volume": "563 Ltr.",
		"Net Volume": "501 Ltr.",
		"Net Weight": "103 ± 2 Kg",
		"Gross Weight": "113 ± 2 Kg",
		"Climatic Type (SN, N, ST, T)": "N~T",
		"Rated Voltage/ Hz/ watt": "220~240 V/ 50 Hz",
		"Compressor Input Power (watt)": "45.4 - 197",
		"Compressor": "BLDC Inverter",
		"Temperature Control (Electronic/ Mechanical)": "Electronic",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Condenser": "Steel",
		"Capillary": "Copper",
		"Handle (Recessed/ Grip)": "Built In",
		"Lock": "No",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Shelf (Material/No.)": "Glass/5",
		"Bottle Pocket": "GPPS/4",
		"Interior Lamp": "Yes",
		"Egg Tray or Pocket": "Yes/ 1",
		"Vegetable Crisper": "Yes/1",
		"Water Dispenser": "No",
		"Ice Case": "Yes/1",
		"Ice Box": "Yes/1",
		"Shelf (Material/ No.)": "GPPS/5",
		"Interior LED Lamp": "Yes",
		"Drawer": "Yes/1",
		"Cooling Effect": "Freezer Cabinet Less than -18℃ Refrigerator Cabinet 0℃ to +5℃",
		"Width/mm": "900 mm",
		"Depth/mm": "775 mm",
		"Height/mm": "1815 mm",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "39/39/19 (Vertical Loading)",
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
