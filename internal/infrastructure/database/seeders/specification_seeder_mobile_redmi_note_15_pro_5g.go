package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote15Pro5g seeds specifications/options for product 'redmi-note-15-pro-5g'
type SpecificationSeederMobileRedmiNote15Pro5g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote15Pro5g creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote15Pro5g() *SpecificationSeederMobileRedmiNote15Pro5g {
	return &SpecificationSeederMobileRedmiNote15Pro5g{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 15 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote15Pro5g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ পিক্সেল",
		"16 MP": "১৬ মেগাপিক্সেল",
		"161.2 x 74.3 x 8 mm": "১৬১.২ x ৭৪.৩ x ৮ মিমি",
		"187 g": "১৮৭ g",
		"200 MP + 8 MP + 2 MP": "২০০ মেগাপিক্সেল + ৮ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"256 GB / 512 GB": "২৫৬ জিবি / ৫১২ GB",
		"5,100 mAh": "৫,১০০ এমএএইচ",
		"5G": "৫G",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"8 GB / 12 GB": "৮ জিবি / ১২ GB",
		"AMOLED, 68B colors, 120Hz, Dolby Vision, 1800 nits": "অ্যামোলেড, ৬৮B রঙ, ১২০Hz, ডলবি ভিশন, ১৮০০ nits",
		"Adreno 710": "অ্যাড্রেনো ৭১০",
		"Android 15, HyperOS": "অ্যান্ড্রয়েড ১৫, হাইপার ওএস",
		"Glass front/back, plastic frame": "গ্লাস সামনে/পেছনে, প্লাস্টিক ফ্রেম",
		"IP68": "IP৬৮",
		"Midnight Black, Aurora Purple, Ocean Teal": "মিডনাইট কালো, অরোরা বেগুনি, ওশান টিল",
		"Qualcomm Snapdragon 7s Gen 3 (4 nm)": "কোয়ালকম স্ন্যাপড্রাগন ৭s জেন ৩ (৪ ন্যানোমিটার)",
		"September 2025": "সেপ্টেম্বর ২০২৫",
		"Snapdragon 7s Gen 3": "স্ন্যাপড্রাগন ৭s জেন ৩",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-15-pro-5g'
func (s *SpecificationSeederMobileRedmiNote15Pro5g) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-15-pro-5g"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMobileSpecs()
	banglaTranslations := s.getBanglaTranslations()

	// Override model-specific values for Redmi Note 15 Pro 5G
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Snapdragon 7s Gen 3"
	specs["Chipset"] = "Qualcomm Snapdragon 7s Gen 3 (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Adreno 710"
	specs["Ram"] = "8 GB / 12 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "AMOLED, 68B colors, 120Hz, Dolby Vision, 1800 nits"
	specs["Resolution"] = "1220 x 2712 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front/back, plastic frame"
	specs["Weight"] = "187 g"
	specs["Dimensions"] = "161.2 x 74.3 x 8 mm"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "200 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5,100 mAh"
	specs["Operating System"] = "Android 15, HyperOS"
	specs["Available Colors"] = "Midnight Black, Aurora Purple, Ocean Teal"
	specs["Announcement Date"] = "September 2025"
	specs["Device Status"] = "Rumored"

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
