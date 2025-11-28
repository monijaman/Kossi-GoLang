package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileXiaomi15Ultra seeds specifications/options for product 'xiaomi-15-ultra'
type SpecificationSeederMobileXiaomi15Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15Ultra() *SpecificationSeederMobileXiaomi15Ultra {
	return &SpecificationSeederMobileXiaomi15Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"120Hz": "১২০Hz",
		"1440 x 3200 pixels (~522 ppi density)": "১৪৪০ x ৩২০০ pixels (~৫২২ ppi density)",
		"16 GB": "১৬ GB",
		"162.8 x 74.5 x 8.3 mm (6.41 x 2.93 x 0.33 in)": "১৬২.৮ x ৭৪.৫ x ৮.৩ মিমি (৬.৪১ x ২.৯৩ x ০.৩৩ in)",
		"220 g (7.76 oz)": "২২০ g (৭.৭৬ oz)",
		"24-bit/192kHz audio, Hi-Res Audio certification": "২৪-bit/১৯২kHz audio, Hi-Res Audio certification",
		"24116PN76C": "২৪১১৬PN৭৬C",
		"2x4.32 GHz + 6x3.53 GHz": "২x৪.৩২ GHz + ৬x৩.৫৩ GHz",
		"3.2x + 5x Optical Zoom": "৩.২x + ৫x অপটিক্যাল জুম",
		"32 MP, f/2.0, 25mm (wide), HDR": "৩২ MP, f/২.০, ২৫mm (wide), HDR",
		"4K@30/60fps, 1080p@30/60fps": "৪K@৩০/৬০fps, ১০৮০p@৩০/৬০fps",
		"5,300 mAh": "৫,৩০০ এমএএইচ",
		"5.4, A2DP, LE, aptX HD": "৫.৪, A২DP, LE, aptX HD",
		"50 MP + 50 MP + 50 MP + 50 MP": "৫০ MP + ৫০ MP + ৫০ MP + ৫০ MP",
		"50MP (wide, f/1.6) + 50MP (periscope telephoto 5x, f/2.5) + 50MP (telephoto 3.2x, f/2.0) + 50MP (ultrawide, f/1.8)": "৫০MP (wide, f/১.৬) + ৫০MP (periscope telephoto ৫x, f/২.৫) + ৫০MP (telephoto ৩.২x, f/২.০) + ৫০MP (ultrawide, f/১.৮)",
		"512 GB / 1 TB": "৫১২ GB / ১ TB",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision": "৮K@২৪fps, ৪K@৩০/৬০fps, ১০৮০p@৩০/৬০/১২০/২৪০fps, gyro-EIS, HDR১০+, Dolby Vision",
		"90W wired (100% in 32 min), 80W wireless (100% in 38 min), 10W reverse wireless": "৯০W তারযুক্ত (১০০% in ৩২ মিনিট), ৮০W বেতার (১০০% in ৩৮ মিনিট), ১০W reverse বেতার",
		"90W wired, 80W wireless, 10W reverse wireless": "৯০W তারযুক্ত, ৮০W বেতার, ১০W reverse বেতার",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, HyperOS": "Android ১৫, HyperOS",
		"Available. Released February 2025": "উপলব্ধ. Released ফেব্রুয়ারি ২০২৫",
		"Black, Titanium, White": "কালো, Titanium, সাদা",
		"February 2025": "February ২০২৫",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, optical), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার, color spectrum",
		"GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a+B2b), GALILEO (E1+E5a+E5b), QZSS (L1+L5), NavIC": "GPS (L১+L৫), GLONASS (G১), BDS (B১I+B১c+B২a+B২b), GALILEO (E১+E৫a+E৫b), QZSS (L১+L৫), NavIC",
		"GSM / CDMA / HSPA / LTE / 5G": "GSM / CDMA / HSPA / LTE / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ MHz",
		"Glass front and back, Aluminum frame": "গ্লাস সামনে and back, অ্যালুমিনিয়াম ফ্রেম",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "HSDPA ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০ MHz",
		"IP68 dust tight and water resistant (up to 1.5m for 30 min)": "IP৬৮ dust tight and water resistant (up to ১.৫m for ৩০ min)",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 26, 28, 32, 38, 40, 41, 42, 66": "LTE Band ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৬, ২৮, ৩২, ৩৮, ৪০, ৪১, ৪২, ৬৬",
		"LTPO AMOLED, 1B colors, 120Hz, Dolby Vision, HDR10+, 3200 nits (peak)": "LTPO AMOLED, ১B colors, ১২০Hz, Dolby Vision, HDR১০+, ৩২০০ nits (peak)",
		"Li-Po (non-removable)": "লি-পো (অপসারণযোগ্য নয়)",
		"Octa-core (2x4.32 GHz Oryon V2 + 6x3.53 GHz Oryon V2)": "Octa-core (২x৪.৩২ GHz Oryon V২ + ৬x৩.৫৩ GHz Oryon V২)",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০ Snapdragon ৮ Elite (৩ nm)",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"UFS 4.0": "UFS ৪.০",
		"USB Type-C 3.2, DisplayPort 1.4, OTG": "USB Type-C ৩.২, DisplayPort ১.৪, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-band, Wi-Fi Direct",
		"Yes - 80W wireless": "Yes - ৮০W wireless",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩৮, n৪০, n৪১, n৭৭, n৭৮",
	}
}

// Seed inserts specification_translations for existing specifications for product 'xiaomi-15-ultra'
func (s *SpecificationSeederMobileXiaomi15Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15-ultra"

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
