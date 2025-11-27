package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileGooglePixel8 seeds specifications for 'google-pixel-8'
type SpecificationSeederMobileGooglePixel8 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel8 creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel8() *SpecificationSeederMobileGooglePixel8 {
	return &SpecificationSeederMobileGooglePixel8{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 8"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel8) getBanglaTranslations() map[string]string {
	return map[string]string{
		"10.5 MP": "১০.৫ মেগাপিক্সেল",
		"10.5 Mayগাpixels": "১০.৫ মেগাপিক্সেল",
		"1080 x 2400 pixels": "১০৮০ x ২৪০০ পিক্সেল",
		"128 GB / 256 GB": "১২৮ গিগাবাইট / ২৫৬ গিগাবাইট",
		"150.5 x 70.8 x 8.9 mm": "১৫০.৫ x ৭০.৮ x ৮.৯ মিমি",
		"187 g": "১৮৭ গ্রাম",
		"4,575 mAh": "৪,৫৭৫ এমএএইচ",
		"50 MP + 12 MP": "৫০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"50 Mayগাpixels + 12 Mayগাpixels": "৫০ মেগাপিক্সেল + ১২ মেগাপিক্সেল",
		"5G": "৫G",
		"6.2 inches": "৬.২ ইঞ্চি",
		"60-120Hz": "৬০-১২০হার্টজ",
		"8 GB": "৮ গিগাবাইট",
		"Android 14": "অ্যান্ড্রয়েড ১৪",
		"Available": "উপলব্ধ",
		"Corning Gorilla Glass Victus": "Corning Gorilla গ্লাস Victus",
		"Glass front/back, aluminum frame": "গ্লাস সামনে/পিছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G3": "Google Tensor G৩",
		"Google Tensor G3 (4 nm)": "Google Tensor G৩ (৪ nm)",
		"IP68": "IP৬৮",
		"Immortalis-G715s MC10": "Immortalis-G৭১৫s MC১০",
		"Nona-core": "Nona-core",
		"OLED, 120Hz, HDR10+, 2000 nits": "ওলেড, ১২০হার্টজ, HDR১০+, ২০০০ নিট",
		"Obsidian, Hazel, Rose, Mint": "Obsidian, Hazel, Rose, Mint",
		"October 2023": "অক্টোবর ২০২৩",
}
}

// Seed inserts specification records for the device
func (s *SpecificationSeederMobileGooglePixel8) Seed(db *gorm.DB) error {
	deviceSlug := "google-pixel-8"

	var prod models.ProductModel
	if err := db.Where("slug = ?", deviceSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := map[string]string{
		"announcement_date": "October 2023",
		"available_colors": "Obsidian, Hazel, Rose, Mint",
		"battery": "4,575 mAh",
		"build_material": "Glass front/back, aluminum frame",
		"chipset": "Google Tensor G3 (4 nm)",
		"cpu_type": "Nona-core",
		"device_status": "Available",
		"dimensions": "150.5 x 70.8 x 8.9 mm",
		"display_size": "6.2 inches",
		"display_type": "OLED, 120Hz, HDR10+, 2000 nits",
		"front_camera": "10.5 Mayগাpixels",
		"gpu_type": "Immortalis-G715s MC10",
		"network_technology": "5G",
		"operating_system": "Android 14",
		"processor": "Google Tensor G3",
		"ram": "8 GB",
		"rear_camera": "50 Mayগাpixels + 12 Mayগাpixels",
		"refresh_rate": "60-120Hz",
		"resolution": "1080 x 2400 pixels",
		"screen_protection": "Corning Gorilla Glass Victus",
		"storage": "128 GB / 256 GB",
		"water_resistance": "IP68",
		"weight": "187 g",
	
}
	banglaTranslations := s.getBanglaTranslations()

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
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
		} else {
			// If specification already exists, check and create Bangla translation if missing
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
