package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileRedmiNote15Pro seeds specifications/options for product 'redmi-note-15-pro'
type SpecificationSeederMobileRedmiNote15Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileRedmiNote15Pro creates a new seeder instance
func NewSpecificationSeederMobileRedmiNote15Pro() *SpecificationSeederMobileRedmiNote15Pro {
	return &SpecificationSeederMobileRedmiNote15Pro{BaseSeeder: BaseSeeder{name: "Specifications for Redmi Note 15 Pro+"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileRedmiNote15Pro) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"120Hz": "১২০Hz",
		"120W wired": "১২০W তারযুক্ত",
		"1220 x 2712 pixels": "১২২০ x ২৭১২ pixels",
		"16 MP": "১৬ MP",
		"200 MP + 8 MP + 2 MP": "২০০ MP + ৮ MP + ২ MP",
		"205 g": "২০৫ g",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.67 inches": "৬.৬৭ ইঞ্চি",
		"AMOLED, 120Hz, Dolby Vision, HDR10+": "AMOLED, ১২০Hz, Dolby Vision, HDR১০+",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Dimensity 8300 Ultra": "Dimensity ৮৩০০ Ultra",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"Glass front (Gorilla Glass Victus), glass back, aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"IP68": "IP৬৮",
		"Mali-G615-MC6": "Mali-G৬১৫-MC৬",
		"Mediatek Dimensity 8300 Ultra (4 nm)": "Mediatek Dimensity ৮৩০০ Ultra (৪ nm)",
		"Midnight Black, Moonlight White, Aurora Purple": "Midnight কালো, Moonlight সাদা, Aurora বেগুনি",
		"September 2025": "September ২০২৫",
	}
}

// Seed inserts specification records for the product identified by slug 'redmi-note-15-pro'
func (s *SpecificationSeederMobileRedmiNote15Pro) Seed(db *gorm.DB) error {
	productSlug := "redmi-note-15-pro"

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

	// Override model-specific values for Redmi Note 15 Pro+
	specs["Display Size"] = "6.67 inches"
	specs["Processor"] = "Dimensity 8300 Ultra"
	specs["Chipset"] = "Mediatek Dimensity 8300 Ultra (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G615-MC6"
	specs["Ram"] = "12 GB / 16 GB"
	specs["Storage"] = "256 GB / 512 GB"
	specs["Display Type"] = "AMOLED, 120Hz, Dolby Vision, HDR10+"
	specs["Resolution"] = "1220 x 2712 pixels"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front (Gorilla Glass Victus), glass back, aluminum frame"
	specs["Weight"] = "205 g"
	specs["Water Resistance"] = "IP68"
	specs["Network Technology"] = "GSM / HSPA / LTE / 5G"
	specs["Rear Camera"] = "200 MP + 8 MP + 2 MP"
	specs["Front Camera"] = "16 MP"
	specs["Battery"] = "5000 mAh"
	specs["Fast Charging"] = "120W wired"
	specs["Operating System"] = "Android 15, HyperOS"
	specs["Available Colors"] = "Midnight Black, Moonlight White, Aurora Purple"
	specs["Announcement Date"] = "September 2025"
	specs["Device Status"] = "Upcoming"

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
