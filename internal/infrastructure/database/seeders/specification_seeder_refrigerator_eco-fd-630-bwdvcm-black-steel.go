package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorecofd630bwdvcmblacksteel seeds specifications/options for product 'eco-fd-630-bwdvcm-black-steel'
type SpecificationSeederRefrigeratorecofd630bwdvcmblacksteel struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorecofd630bwdvcmblacksteel creates a new seeder instance
func NewSpecificationSeederRefrigeratorecofd630bwdvcmblacksteel() *SpecificationSeederRefrigeratorecofd630bwdvcmblacksteel {
	return &SpecificationSeederRefrigeratorecofd630bwdvcmblacksteel{
		BaseSeeder: BaseSeeder{name: "Specifications for eco-fd-630-bwdvcm-black-steel"},
	}
}

func (s *SpecificationSeederRefrigeratorecofd630bwdvcmblacksteel) getBanglaTranslations() map[string]string {
	return map[string]string{
		"ECO+":                                  "ইকো+",
		"Eco+ FD-630-BWDVCM Black Steel":        "Eco+ FD-630-BWDVCM Black Steel",
		"Glass Door":                            "গ্লাস ডোর",
		"630":                                   "৬৩০",
		"414":                                   "৪১৪",
		"122":                                   "১২২",
		"536":                                   "৫৩৬",
		"1776*912*765 mm":                       "১৭৭৬*৯১২*৭৬৫ মিমি",
		"Black Steel":                           "ব্ল্যাক স্টিল",
		"R600a":                                 "R600a",
		"Auto":                                  "অটো",
		"LED display with touch button control": "LED ডিসপ্লে উইথ টাচ বাটন কন্ট্রোল",
		"3":                                     "৩",
		"2":                                     "২",
		"Multi Air Flow System, LED Display with Touch Button Control, External Water Dispenser, Folded coke rack, Easy-Slide Freezer Drawer, Movable twist ice-maker": "মাল্টি এয়ার ফ্লো সিস্টেম, LED ডিসপ্লে উইথ টাচ বাটন কন্ট্রোল, এক্সটার্নাল ওয়াটার ডিসপেন্সার, ফোল্ডেড কোক র্যাক, ইজি-স্লাইড ফ্রিজার ড্রয়ার, মুভেবল টুইস্ট আইস-মেকার",
		"Yes":     "হ্যাঁ",
		"No":      "না",
		"220~240": "২২০~২৪০",
		"50":      "৫০",
		// Add more translations as needed
	}
}

// Seed inserts specification records for the product identified by slug 'eco-fd-630-bwdvcm-black-steel'
func (s *SpecificationSeederRefrigeratorecofd630bwdvcmblacksteel) Seed(db *gorm.DB) error {
	productSlug := "eco-fd-630-bwdvcm-black-steel"
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
		"Brand":                       "ECO+",
		"Model Name":                  "Eco+ FD-630-BWDVCM Black Steel",
		"Door Type":                   "Glass Door",
		"Capacity":                    "630",
		"Refrigerator Capacity":       "414",
		"Freezer Capacity":            "122",
		"Gross Volume":                "630",
		"Net Volume":                  "536",
		"Dimensions":                  "1776*912*765 mm",
		"Weight":                      "120",
		"Color":                       "Black Steel",
		"Compressor Type":             "R600a",
		"Defrost Type":                "Auto",
		"Temperature Control":         "LED display with touch button control",
		"Number of Shelves":           "2",
		"Ice Maker":                   "Moveable Twist ice-maker",
		"Water Dispenser":             "External Water Dispenser",
		"Voltage":                     "220~240",
		"Frequency (Hz)":              "50",
		"Warranty":                    "10 Years Compressor Warranty, 2 Years Parts and Service Warranty",
		"Compressor Warranty (Years)": "10",
		"Refrigerant":                 "R600a",
		"Special Features":            "Multi Air Flow System, LED Display with Touch Button Control, External Water Dispenser, Folded coke rack, Easy-Slide Freezer Drawer, Movable twist ice-maker",
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
			// Update the value if different
			if existing.Value != value {
				existing.Value = value
				if err := db.Save(&existing).Error; err != nil {
					return err
				}
			}
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
				} else {
					// Update translation if different
					if existingTranslation.Value != banglaValue {
						existingTranslation.Value = banglaValue
						if err := db.Save(&existingTranslation).Error; err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
