package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSymphonyS72 seeds specifications/options for product 'symphony-s72'
type SpecificationSeederMobileSymphonyS72 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSymphonyS72 creates a new seeder instance
func NewSpecificationSeederMobileSymphonyS72() *SpecificationSeederMobileSymphonyS72 {
	return &SpecificationSeederMobileSymphonyS72{BaseSeeder: BaseSeeder{name: "Specifications for Symphony S72"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSymphonyS72) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.08 MP": "০.০৮ MP",
		"1000 contacts": "১০০০ contacts",
		"128 × 160 px": "১২৮ × ১৬০ px",
		"1500 mAh": "১৫০০ এমএএইচ",
		"2.4 inches": "২.৪ ইঞ্চি",
		"2G GSM": "২G GSM",
		"Bluetooth, microUSB, 3.5mm jack": "Bluetooth, microUSB, ৩.৫mm jack",
		"December 2024": "December ২০২৪",
		"Li-Ion (removable)": "লি-আয়ন (removable)",
		"Straw Green, Dew Green, Ocean Blue, Ink Black, Caramel Gold": "Straw সবুজ, Dew সবুজ, Ocean নীল, Ink কালো, Caramel সোনালী",
		"microSD up to 16 GB": "microSD up to ১৬ GB",
	}
}

// Seed inserts specification records for the product identified by slug 'symphony-s72'
func (s *SpecificationSeederMobileSymphonyS72) Seed(db *gorm.DB) error {
	productSlug := "symphony-s72"

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

	// Override model-specific values for Symphony S72
	specs["Device Type"] = "Feature Phone"
	specs["Display Size"] = "2.4 inches"
	specs["Display Type"] = "TFT"
	specs["Resolution"] = "128 × 160 px"
	specs["Battery"] = "1500 mAh"
	specs["Battery Type"] = "Li-Ion (removable)"
	specs["Network Technology"] = "2G GSM"
	specs["Sim Card Type"] = "Dual SIM"
	specs["Memory Storage"] = "microSD up to 16 GB"
	specs["Phonebook Capacity"] = "1000 contacts"
	specs["Rear Camera"] = "0.08 MP"
	specs["Camera Video Resolution"] = "Basic video recording"
	specs["Audio Features"] = "Loudspeaker, FM Radio"
	specs["Connectivity"] = "Bluetooth, microUSB, 3.5mm jack"
	specs["Special Features"] = "Auto call recording, call blocking"
	specs["Dimensions"] = "Standard feature phone size"
	specs["Build Material"] = "Plastic"
	specs["Available Colors"] = "Straw Green, Dew Green, Ocean Blue, Ink Black, Caramel Gold"
	specs["Announcement Date"] = "December 2024"
	specs["Device Status"] = "Available"

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
