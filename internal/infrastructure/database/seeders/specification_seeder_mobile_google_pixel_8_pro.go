package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"os"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel8Pro seeds specifications/options for product 'google-pixel-8-pro'
type SpecificationSeederMobileGooglePixel8Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8Pro creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8Pro() *SpecificationSeederMobileGooglePixel8Pro {
	return &SpecificationSeederMobileGooglePixel8Pro{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz":                             "১-১২০Hz",
		"10.5 MP":                             "১০.৫ মেগাপিক্সেল",
		"12 GB":                               "১২ GB",
		"128 GB / 256 GB / 512 GB / 1 TB":     "১২৮ জিবি / ২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"1344 x 2992 pixels":                  "১৩৪৪ x ২৯৯২ পিক্সেল",
		"162.6 x 76.5 x 8.8 mm":               "১৬২.৬ x ৭৬.৫ x ৮.৮ মিমি",
		"213 g":                               "২১৩ g",
		"5,050 mAh":                           "৫,০৫০ এমএএইচ",
		"50 MP + 48 MP + 48 MP":               "৫০ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল + ৪৮ মেগাপিক্সেল",
		"5G":                                  "৫G",
		"6.7 inches":                          "৬.৭ ইঞ্চি",
		"Android 14":                          "অ্যান্ড্রয়েড ১৪",
		"Corning Gorilla Glass Victus 2":      "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"Glass front/back, aluminum frame":    "গ্লাস সামনে/পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3":                    "গুগল টেনসর G৩",
		"Google Tensor G3 (4 nm)":             "গুগল টেনসর G৩ (৪ ন্যানোমিটার)",
		"IP68":                                "IP৬৮",
		"Immortalis-G715s MC10":               "Immবাtalis-G৭১৫s MC১০",
		"LTPO OLED, 120Hz, HDR10+, 2400 nits": "এলটিপিও ওএলইডি, ১২০Hz, এইচডিআর১০+, ২৪০০ nits",
		"October 2023":                        "অক্টোবর ২০২৩",
	}
}

// Seed inserts specification records for the product identified by slug 'google-pixel-8-pro'
func (s *SpecificationSeederMobileGooglePixel8Pro) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-8-pro"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("[DEBUG] Product not found: %s\n", productSlug)
			return nil
		}
		return err
	}
	productID := prod.ID
	fmt.Printf("uuuuuuuuuuuuuuuuuu [DEBUG] Found product: %s (ID: %d)\n", productSlug, productID)

	specs := DefaultMobileSpecs()
	banglaTranslations := s.getBanglaTranslations()
	println(banglaTranslations)
	// Override model-specific values for Google Pixel 8 Pro
	specs["Display Size"] = "6.7 inches"
	specs["Processor"] = "Google Tensor G3"
	specs["Chipset"] = "Google Tensor G3 (4 nm)"
	specs["Cpu Type"] = "Nona-core"
	specs["Gpu Type"] = "Immortalis-G715s MC10"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "128 GB / 256 GB / 512 GB / 1 TB"
	specs["Display Type"] = "LTPO OLED, 120Hz, HDR10+, 2400 nits"
	specs["Resolution"] = "1344 x 2992 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus 2"
	specs["Refresh Rate"] = "1-120Hz"
	specs["Build Material"] = "Glass front/back, aluminum frame"
	specs["Weight"] = "213 g"
	specs["Dimensions"] = "162.6 x 76.5 x 8.8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 48 MP + 48 MP"
	specs["Front Camera"] = "10.5 MP"
	specs["Battery"] = "5,050 mAh"
	specs["Operating System"] = "Android 14"
	specs["Available Colors"] = "Obsidian, Porcelain, Bay, Mint"
	specs["Announcement Date"] = "October 2023"
	specs["Device Status"] = "Available"

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}
		fmt.Printf("EEEEEEEEEEEEEEEEEEEEEEEEEEEE")
		fmt.Println(sk)
		os.Exit(1)
		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {

			fmt.Printf(" 77777777777777777777777777777 ")

			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {

					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							translation := &models.SpecificationTranslationModel{
								SpecificationID: sModel.ID,
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
			} else {
				return err
			}
			fmt.Printf(" ------------------------------------------- ")
			fmt.Println(existing)
			return nil
		} else {
			// fmt.Printf(" 8888888888888888888 ")

			// os.Exit(1)
			// Update the specification value if it's different
			if existing.Value != value {
				existing.Value = value
				if err := db.Save(&existing).Error; err != nil {
					return err
				}
			}

			// If specification already exists, check and update/create Bangla translation
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				fmt.Printf(" ++++++++++++++++++++ ")
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
						fmt.Printf("[DEBUG] Created translation for EXISTING SpecID: %d, Locale: bn, Value: %s\n", existing.ID, banglaValue)
					} else {
						return err
					}
				} else {
					fmt.Printf(" SpecificationTranslationModel ")
					fmt.Printf("Existing Translation: %+v\n", existingTranslation)
					// Update translation if it's different
					if existingTranslation.Value != banglaValue {
						existingTranslation.Value = banglaValue
						if err := db.Save(&existingTranslation).Error; err != nil {
							return err
						}
						fmt.Printf("[DEBUG] Updated translation for SpecID: %d with new value: %s\n", existing.ID, banglaValue)
					} else {
						fmt.Printf("[DEBUG] Translation already exists and is up-to-date for existing SpecID: %d\n", existing.ID)
					}
				}
			}
		}

	}

	return nil
}
