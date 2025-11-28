package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileWaltonXanonX1Ultra seeds specifications/options for product 'walton-xanon-x1-ultra'
type SpecificationSeederMobileWaltonXanonX1Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonXanonX1Ultra creates a new seeder instance
func NewSpecificationSeederMobileWaltonXanonX1Ultra() *SpecificationSeederMobileWaltonXanonX1Ultra {
	return &SpecificationSeederMobileWaltonXanonX1Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Walton XANON X1 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonXanonX1Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"108 MP + 50 MP + 20 MP": "১০৮ MP + ৫০ MP + ২০ MP",
		"1080 x 2412 pixels (~394 ppi density)": "১০৮০ x ২৪১২ pixels (~৩৯৪ ppi density)",
		"12 GB": "১২ GB",
		"120Hz": "১২০Hz",
		"161.8 x 73.9 x 8.4 mm (6.37 x 2.91 x 0.33 in)": "১৬১.৮ x ৭৩.৯ x ৮.৪ মিমি (৬.৩৭ x ২.৯১ x ০.৩৩ in)",
		"195 g (6.88 oz)": "১৯৫ g (৬.৮৮ oz)",
		"1x3.25 GHz + 3x2.85 GHz + 4x2.0 GHz": "১x৩.২৫ GHz + ৩x২.৮৫ GHz + ৪x২.০ GHz",
		"24-bit/192kHz audio, Hi-Res Audio": "২৪-bit/১৯২kHz audio, Hi-Res Audio",
		"256 GB": "২৫৬ GB",
		"32 MP, f/2.4, (wide), HDR": "৩২ MP, f/২.৪, (wide), HDR",
		"3x Optical Zoom": "৩x অপটিক্যাল জুম",
		"4K@30/60fps, 1080p@30/60/120fps, gyro-EIS": "৪K@৩০/৬০fps, ১০৮০p@৩০/৬০/১২০fps, জাইরো-EIS",
		"4K@30fps, 1080p@30/60fps": "৪K@৩০fps, ১০৮০p@৩০/৬০fps",
		"5,000 mAh": "৫,০০০ এমএএইচ",
		"5.3, A2DP, LE, aptX HD": "৫.৩, A২DP, LE, aptX HD",
		"6.7 inches": "৬.৭ ইঞ্চি",
		"66W wired": "৬৬W তারযুক্ত",
		"66W wired (100% in 45 min)": "৬৬W তারযুক্ত (১০০% in ৪৫ মিনিট)",
		"AMOLED, 120Hz, HDR10+, 1200 nits (peak)": "AMOLED, ১২০Hz, HDR১০+, ১২০০ nits (peak)",
		"Android 14, Walton WaltonOS": "Android ১৪, Walton WaltonOS",
		"December 2024": "December ২০২৪",
		"Dual stereo speakers, Hi-Res Audio": "Dual স্টেরিও স্পিকার, Hi-Res Audio",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, optical), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"GSM / HSPA / LTE / 5G": "GSM / HSPA / LTE / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ MHz",
		"Glass front (Gorilla Glass Victus), glass back, aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"HSDPA 850 / 900 / 1900 / 2100 MHz": "HSDPA ৮৫০ / ৯০০ / ১৯০০ / ২১০০ MHz",
		"IP67 dust tight and water resistant (up to 1m for 30 min)": "IP৬৭ dust tight and water resistant (up to ১m for ৩০ min)",
		"LTE Band 1, 3, 5, 7, 8, 20, 28, 38, 40, 41": "LTE Band ১, ৩, ৫, ৭, ৮, ২০, ২৮, ৩৮, ৪০, ৪১",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Mali-G720 Immortalis MC12": "Mali-G৭২০ Immortalis MC১২",
		"MediaTek Dimensity 9300": "MediaTek Dimensity ৯৩০০",
		"MediaTek Dimensity 9300 (4 nm)": "MediaTek Dimensity ৯৩০০ (৪ nm)",
		"মিডনাইট Black, Pearl White, Ocean Blue": "মিডনাইট কালো, Pearl সাদা, Ocean নীল",
		"Octa-core (1x3.25 GHz Cortex-X4 & 3x2.85 GHz Cortex-X4 & 4x2.0 GHz Cortex-A720)": "Octa-core (১x৩.২৫ GHz Cortex-X৪ & ৩x২.৮৫ GHz Cortex-X৪ & ৪x২.০ GHz Cortex-A৭২০)",
		"UFS 4.0": "UFS ৪.০",
		"USB Type-C 3.1, OTG": "USB Type-C ৩.১, OTG",
		"WX1U-001": "WX১U-০০১",
		"Wi-Fi 802.11 a/b/g/n/ac/6e, dual-band, Wi-Fi Direct": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e, dual-band, Wi-Fi Direct",
		"microSD card slot up to 1TB": "microSD card slot up to ১TB",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩৮, n৪০, n৪১, n৭৭, n৭৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'walton-xanon-x1-ultra'
func (s *SpecificationSeederMobileWaltonXanonX1Ultra) Seed(db *gorm.DB) error {
	productSlug := "walton-xanon-x1-ultra"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	productID := prod.ID
	banglaTranslations := s.getBanglaTranslations()

	// Get all existing specifications for this product
	var existingSpecs []models.SpecificationModel
	if err := db.Where("product_id = ?", productID).Find(&existingSpecs).Error; err != nil {
		return err
	}

	// Insert translations for all existing specifications
	for _, spec := range existingSpecs {
		banglaValue, exists := banglaTranslations[spec.Value]
		if exists && banglaValue != "" {
			translation := &models.SpecificationTranslationModel{
				SpecificationID: spec.ID,
				Locale:          "bn",
				Value:           banglaValue,
			}
			// Use OnConflict to ignore if translation already exists
			if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(translation).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
