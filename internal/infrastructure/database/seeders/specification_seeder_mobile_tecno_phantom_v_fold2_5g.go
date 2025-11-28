package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoPhantomVFold25g seeds specifications/options for product 'tecno-phantom-v-fold2-5g'
type SpecificationSeederMobileTecnoPhantomVFold25g struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPhantomVFold25g creates a new seeder instance
func NewSpecificationSeederMobileTecnoPhantomVFold25g() *SpecificationSeederMobileTecnoPhantomVFold25g {
	return &SpecificationSeederMobileTecnoPhantomVFold25g{BaseSeeder: BaseSeeder{name: "Specifications for Tecno PHANTOM V Fold2 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPhantomVFold25g) getBanglaTranslations() map[string]string {
	return map[string]string{
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"16 MP (Cover) + 32 MP (Inner)": "১৬ মেগাপিক্সেল (Cover) + ৩২ মেগাপিক্সেল (Inner)",
		"2000 x 2296 pixels": "২০০০ x ২২৯৬ পিক্সেল",
		"299 g": "২৯৯ g",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"50 MP + 50 MP + 13 MP": "৫০ মেগাপিক্সেল + ৫০ মেগাপিক্সেল + ১৩ মেগাপিক্সেল",
		"512 GB": "৫১২ GB",
		"5G": "৫G",
		"7.85 inches (Folded: 6.42 inches)": "৭.৮৫ ইঞ্চি (Folded: ৬.৪২ ইঞ্চি)",
		"Android 14, HIOS 14 Fold": "অ্যান্ড্রয়েড ১৪, HIOS ১৪ Fold",
		"Black, White": "কালো, সাদা",
		"Dimensity 9000+": "ডাইমেনসিটি ৯০০০+",
		"Foldable LTPO AMOLED, 120Hz": "Foldable এলটিপিও অ্যামোলেড, ১২০Hz",
		"Glass front, glass back, aluminum frame": "গ্লাস সামনে, গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Mali-G710 MC10": "মালি-G৭১০ MC১০",
		"Mediatek Dimensity 9000+ (4 nm)": "মিডিয়াটেক ডাইমেনসিটি ৯০০০+ (৪ ন্যানোমিটার)",
		"September 2024": "সেপ্টেম্বর ২০২৪",
		"Unfolded: 159.4 x 140.4 x 6.9 mm": "Unfolded: ১৫৯.৪ x ১৪০.৪ x ৬.৯ মিমি",
	}
}

// Seed inserts specification records for the product identified by slug 'tecno-phantom-v-fold2-5g'
func (s *SpecificationSeederMobileTecnoPhantomVFold25g) Seed(db *gorm.DB) error {
	productSlug := "tecno-phantom-v-fold2-5g"

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

	// Override model-specific values for Tecno PHANTOM V Fold2 5G
	specs["Display Size"] = "7.85 inches (Folded: 6.42 inches)"
	specs["Processor"] = "Dimensity 9000+"
	specs["Chipset"] = "Mediatek Dimensity 9000+ (4 nm)"
	specs["Cpu Type"] = "Octa-core"
	specs["Gpu Type"] = "Mali-G710 MC10"
	specs["Ram"] = "12 GB"
	specs["Storage"] = "512 GB"
	specs["Display Type"] = "Foldable LTPO AMOLED, 120Hz"
	specs["Resolution"] = "2000 x 2296 pixels"
	specs["Screen Protection"] = "Corning Gorilla Glass Victus (Cover)"
	specs["Refresh Rate"] = "120Hz"
	specs["Build Material"] = "Glass front, glass back, aluminum frame"
	specs["Weight"] = "299 g"
	specs["Dimensions"] = "Unfolded: 159.4 x 140.4 x 6.9 mm"
	specs["Water Resistance"] = "No"
	specs["Network Technology"] = "5G"
	specs["Rear Camera"] = "50 MP + 50 MP + 13 MP"
	specs["Front Camera"] = "16 MP (Cover) + 32 MP (Inner)"
	specs["Battery"] = "5,000 mAh"
	specs["Operating System"] = "Android 14, HIOS 14 Fold"
	specs["Available Colors"] = "Black, White"
	specs["Announcement Date"] = "September 2024"
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
