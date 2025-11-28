package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileOneplus13 seeds specifications/options for product 'oneplus-13'
type SpecificationSeederMobileOneplus13 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileOneplus13 creates a new seeder instance
func NewSpecificationSeederMobileOneplus13() *SpecificationSeederMobileOneplus13 {
	return &SpecificationSeederMobileOneplus13{BaseSeeder: BaseSeeder{name: "Specifications for OnePlus 13"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileOneplus13) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1-120Hz (LTPO)": "১-১২০Hz (LTPO)",
		"100W wired (100% in 36 min), 50W wireless (100% in 56 min), 10W reverse wireless": "১০০W তারযুক্ত (১০০% in ৩৬ মিনিট), ৫০W বেতার (১০০% in ৫৬ মিনিট), ১০W reverse বেতার",
		"100W wired, 50W wireless, 10W reverse wireless": "১০০W তারযুক্ত, ৫০W বেতার, ১০W reverse বেতার",
		"12 GB / 16 GB / 24 GB": "১২ GB / ১৬ GB / ২৪ GB",
		"1440 x 3168 pixels (~510 ppi density)": "১৪৪০ x ৩১৬৮ pixels (~৫১০ ppi density)",
		"162.9 x 76.5 x 8.5 mm (glass) / 8.9 mm (leather)": "১৬২.৯ x ৭৬.৫ x ৮.৫ মিমি (glass) / ৮.৯ মিমি (leather)",
		"213 g (glass) / 210 g (leather) (7.51 oz)": "২১৩ g (glass) / ২১০ g (leather) (৭.৫১ oz)",
		"24-bit/192kHz audio, Hi-Res Audio": "২৪-bit/১৯২kHz audio, Hi-Res Audio",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"2x4.32 GHz + 6x3.53 GHz": "২x৪.৩২ GHz + ৬x৩.৫৩ GHz",
		"32 MP, f/2.4, 21mm (wide), Auto-HDR": "৩২ MP, f/২.৪, ২১mm (wide), Auto-HDR",
		"3x Optical Zoom": "৩x অপটিক্যাল জুম",
		"4K@30/60fps, 1080p@30/60fps": "৪K@৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5.4, A2DP, LE, aptX HD": "৫.৪, A২DP, LE, aptX HD",
		"50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP",
		"50MP (wide, f/1.6, OIS) + 50MP (periscope telephoto 3x, f/2.6, OIS) + 50MP (ultrawide, f/2.0)": "৫০MP (wide, f/১.৬, OIS) + ৫০MP (periscope telephoto ৩x, f/২.৬, OIS) + ৫০MP (ultrawide, f/২.০)",
		"6,000 mAh": "৬,০০০ এমএএইচ",
		"6.82 inches": "৬.৮২ ইঞ্চি",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/240fps, Auto HDR, gyro-EIS, Dolby Vision": "৮K@২৪fps, ৪K@৩০/৬০fps, ১০৮০p@৩০/৬০/২৪০fps, Auto HDR, gyro-EIS, Dolby Vision",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, OxygenOS 15 (Global) / ColorOS 15 (China)": "Android ১৫, OxygenOS ১৫ (Global) / ColorOS ১৫ (China)",
		"Available. Released January 07, 2025": "উপলব্ধ. Released জানুয়ারি ০৭, ২০২৫",
		"CPH2649, PJZ110": "CPH২৬৪৯, PJZ১১০",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, optical), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার",
		"GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a+B2b), GALILEO (E1+E5a+E5b), QZSS (L1+L5), NavIC (L5)": "GPS (L১+L৫), GLONASS (G১), BDS (B১I+B১c+B২a+B২b), GALILEO (E১+E৫a+E৫b), QZSS (L১+L৫), NavIC (L৫)",
		"GSM / CDMA / HSPA / LTE / 5G": "GSM / CDMA / HSPA / LTE / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ MHz",
		"Glass front (Ceramic Guard), glass back or eco leather back, Aluminum frame": "গ্লাস সামনে (Ceramic Guard), গ্লাস পেছনে or eco leather back, অ্যালুমিনিয়াম ফ্রেম",
		"HSDPA 800 / 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "HSDPA ৮০০ / ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০ MHz",
		"IP68/IP69 dust tight and water resistant": "IP৬৮/IP৬৯ dust tight and water resistant",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 34, 38, 39, 40, 41, 42, 46, 48, 66, 71": "LTE Band ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৫, ২৬, ২৮, ৩২, ৩৪, ৩৮, ৩৯, ৪০, ৪১, ৪২, ৪৬, ৪৮, ৬৬, ৭১",
		"LTPO AMOLED, 1B colors, 120Hz, Dolby Vision, HDR10+, 4500 nits (peak)": "LTPO AMOLED, ১B colors, ১২০Hz, Dolby Vision, HDR১০+, ৪৫০০ nits (peak)",
		"Octa-core (2x4.32 GHz Oryon V2 Phoenix L + 6x3.53 GHz Oryon V2 Phoenix M)": "Octa-core (২x৪.৩২ GHz Oryon V২ Phoenix L + ৬x৩.৫৩ GHz Oryon V২ Phoenix M)",
		"October 31, 2024": "অক্টোবর ৩১, ২০২৪",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০ Snapdragon ৮ Elite (৩ nm)",
		"Si/C 6000 mAh (non-removable)": "Si/C ৬০০০ এমএএইচ (অপসারণযোগ্য নয়)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"USB Type-C 3.2, DisplayPort 1.4, OTG": "USB Type-C ৩.২, DisplayPort ১.৪, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-band, Wi-Fi Direct",
		"Yes - 50W AIRVOOC wireless": "Yes - ৫০W AIRVOOC wireless",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩৮, n৪০, n৪১, n৭৭, n৭৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'oneplus-13'
func (s *SpecificationSeederMobileOneplus13) Seed(db *gorm.DB) error {
	productSlug := "oneplus-13"

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
