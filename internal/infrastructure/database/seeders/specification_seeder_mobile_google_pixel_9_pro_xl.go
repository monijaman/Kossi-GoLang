package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileGooglePixel9ProXl seeds specifications/options for product 'google-pixel-9-pro-xl'
type SpecificationSeederMobileGooglePixel9ProXl struct {
	BaseSeeder
}

// NewSpecificationSeederMobileGooglePixel9ProXl creates a new seeder instance
func NewSpecificationSeederMobileGooglePixel9ProXl() *SpecificationSeederMobileGooglePixel9ProXl {
	return &SpecificationSeederMobileGooglePixel9ProXl{BaseSeeder: BaseSeeder{name: "Specifications for Google Pixel 9 Pro XL"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileGooglePixel9ProXl) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz (LTPO)": "১-১২০Hz (LTPO)",
		"1.15 W/kg (head), 1.36 W/kg (body)": "১.১৫ W/kg (head), ১.৩৬ W/kg (body)",
		"1.17 W/kg (head), 1.24 W/kg (body)": "১.১৭ W/kg (head), ১.২৪ W/kg (body)",
		"1344 x 2992 pixels (~486 ppi density)": "১৩৪৪ x ২৯৯২ pixels (~৪৮৬ ppi density)",
		"16 GB": "১৬ GB",
		"162.8 x 76.6 x 8.5 mm (6.41 x 3.02 x 0.33 in)": "১৬২.৮ x ৭৬.৬ x ৮.৫ মিমি (৬.৪১ x ৩.০২ x ০.৩৩ in)",
		"1x3.1 GHz + 3x2.6 GHz + 4x1.92 GHz": "১x৩.১ GHz + ৩x২.৬ GHz + ৪x১.৯২ GHz",
		"221 g (7.80 oz)": "২২১ g (৭.৮০ oz)",
		"24-bit/192kHz audio": "২৪-bit/১৯২kHz audio",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"37W wired (50% in 28 min), 23W wireless, 5W reverse wireless": "৩৭W তারযুক্ত (৫০% in ২৮ মিনিট), ২৩W বেতার, ৫W reverse বেতার",
		"37W wired, 23W wireless, 5W reverse wireless": "৩৭W তারযুক্ত, ২৩W বেতার, ৫W reverse বেতার",
		"42 MP, f/2.2, 17mm (ultrawide), PDAF": "৪২ MP, f/২.২, ১৭mm (ultrawide), PDAF",
		"4K@24/30/60fps, 1080p@24/30/60/120/240fps, gyro-EIS, OIS, 10-bit HDR": "৪K@২৪/৩০/৬০fps, ১০৮০p@২৪/৩০/৬০/১২০/২৪০fps, gyro-EIS, OIS, ১০-bit HDR",
		"4K@24/30/60fps, 1080p@30/60fps": "৪K@২৪/৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5,060 mAh": "৫,০৬০ এমএএইচ",
		"5.3, A2DP, LE, aptX HD": "৫.৩, A২DP, LE, aptX HD",
		"50 MP + 48 MP + 48 MP": "৫০ MP + ৪৮ MP + ৪৮ MP",
		"5x Optical Zoom": "৫x অপটিক্যাল জুম",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"Android 15, upgradable to Android 22": "Android ১৫, upgradable to Android ২২",
		"August 13, 2024": "আগস্ট ১৩, ২০২৪",
		"Available. Released August 22, 2024": "উপলব্ধ. Released আগস্ট ২২, ২০২৪",
		"Corning Gorilla Glass Victus 2": "Corning Gorilla Glass Victus ২",
		"Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer, thermometer (skin temperature)": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, ultrasonic), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার, thermometer (skin temperature)",
		"GA05844-US, GA05845-US, GA05846-US": "GA০৫৮৪৪-US, GA০৫৮৪৫-US, GA০৫৮৪৬-US",
		"GPS (L1+L5), GLONASS (G1), GALILEO (E1+E5a), BDS (B1I+B1c+B2a+B2b), QZSS (L1+L5)": "GPS (L১+L৫), GLONASS (G১), GALILEO (E১+E৫a), BDS (B১I+B১c+B২a+B২b), QZSS (L১+L৫)",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ MHz",
		"Gemini AI assistant, Circle to Search, Magic Eraser, Best Take, Audio Magic Eraser, Call Screen, Hold for Me": "Gemini AI assistant, Circle to Search, Magic Eraser, Best Take, Audio Magic Eraser, Call Screen, Hold জন্য Me",
		"Glass front (Gorilla Glass Victus 2), glass back, aluminum frame": "গ্লাস সামনে (Gorilla Glass Victus ২), গ্লাস পেছনে, অ্যালুমিনিয়াম ফ্রেম",
		"Google Tensor G4": "Google Tensor G৪",
		"Google Tensor G4 (4 nm)": "Google Tensor G৪ (৪ nm)",
		"HSDPA 800 / 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "HSDPA ৮০০ / ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০ MHz",
		"IP68 dust tight and water resistant (up to 1.5m for 30 min)": "IP৬৮ dust tight and water resistant (up to ১.৫m for ৩০ min)",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 14, 17, 18, 19, 20, 25, 26, 28, 29, 30, 32, 38, 39, 40, 41, 42, 46, 48, 66, 71": "LTE Band ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৪, ১৭, ১৮, ১৯, ২০, ২৫, ২৬, ২৮, ২৯, ৩০, ৩২, ৩৮, ৩৯, ৪০, ৪১, ৪২, ৪৬, ৪৮, ৬৬, ৭১",
		"LTPO OLED, 120Hz, HDR10+, 3000 nits (peak)": "LTPO OLED, ১২০Hz, HDR১০+, ৩০০০ nits (peak)",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Mali-G715 MC7": "Mali-G৭১৫ MC৭",
		"Octa-core (1x3.1 GHz Cortex-X4 & 3x2.6 GHz Cortex-A720 & 4x1.92 GHz Cortex-A520)": "Octa-core (১x৩.১ GHz Cortex-X৪ & ৩x২.৬ GHz Cortex-A৭২০ & ৪x১.৯২ GHz Cortex-A৫২০)",
		"UFS 3.1": "UFS ৩.১",
		"USB Type-C 3.2, DisplayPort 1.4, OTG": "USB Type-C ৩.২, DisplayPort ১.৪, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-band, Wi-Fi Direct",
		"Yes - 23W wireless": "Yes - ২৩W wireless",
		"n1, n2, n3, n5, n7, n8, n12, n14, n20, n25, n26, n28, n30, n38, n40, n41, n48, n66, n71, n75, n76, n77, n78": "n১, n২, n৩, n৫, n৭, n৮, n১২, n১৪, n২০, n২৫, n২৬, n২৮, n৩০, n৩৮, n৪০, n৪১, n৪৮, n৬৬, n৭১, n৭৫, n৭৬, n৭৭, n৭৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'google-pixel-9-pro-xl'
func (s *SpecificationSeederMobileGooglePixel9ProXl) Seed(db *gorm.DB) error {
	productSlug := "google-pixel-9-pro-xl"

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
