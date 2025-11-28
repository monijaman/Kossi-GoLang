package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS25 seeds specifications/options for product 'samsung-galaxy-s25'
type SpecificationSeederMobileSamsungGalaxyS25 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS25 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS25() *SpecificationSeederMobileSamsungGalaxyS25 {
	return &SpecificationSeederMobileSamsungGalaxyS25{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S25+"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS25) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.99 W/kg (head), 1.47 W/kg (body)": "০.৯৯ W/kg (head), ১.৪৭ W/kg (body)",
		"12 GB": "১২ GB",
		"12 MP, f/2.2, 26mm (wide), Dual Pixel PDAF": "১২ MP, f/২.২, ২৬mm (wide), Dual Pixel PDAF",
		"120Hz (Adaptive)": "১২০Hz (Adaptive)",
		"1440 x 3120 pixels (~516 ppi density)": "১৪৪০ x ৩১২০ pixels (~৫১৬ ppi density)",
		"162.3 x 75.6 x 7.7 mm (6.39 x 2.98 x 0.30 in)": "১৬২.৩ x ৭৫.৬ x ৭.৭ মিমি (৬.৩৯ x ২.৯৮ x ০.৩০ in)",
		"195 g (6.88 oz)": "১৯৫ g (৬.৮৮ oz)",
		"256 GB / 512 GB": "২৫৬ GB / ৫১২ GB",
		"3.4 GHz": "৩.৪ GHz",
		"32-bit/384kHz audio, Tuned by AKG, Dolby Atmos": "৩২-bit/৩৮৪kHz audio, Tuned by AKG, Dolby Atmos",
		"3x optical zoom, up to 30x Space Zoom": "৩x অপটিক্যাল জুম, পর্যন্ত ৩০x Space Zoom",
		"4,900 mAh": "৪,৯০০ এমএএইচ",
		"45W Super Fast Charging 2.0 (50% in 30 min); 15W Fast Wireless Charging 2.0; 4.5W reverse wireless": "৪৫W Super ফাস্ট চার্জিং ২.০ (৫০% in ৩০ মিনিট); ১৫W Fast বেতার Charging ২.০; ৪.৫W reverse বেতার",
		"45W wired, PD3.0; 15W wireless (Qi/PMA); 4.5W reverse wireless": "৪৫W তারযুক্ত, PD৩.০; ১৫W বেতার (Qi/PMA); ৪.৫W reverse বেতার",
		"4K@30/60fps, 1080p@30fps": "৪K@৩০/৬০fps, ১০৮০p@৩০fps",
		"5.4, A2DP, LE, aptX HD": "৫.৪, A২DP, LE, aptX HD",
		"50 MP, f/1.8, 24mm (wide), 1/1.56\", PDAF, OIS + 10 MP, f/2.4, 67mm (telephoto), PDAF, OIS, 3x optical zoom + 12 MP, f/2.2, 13mm (ultrawide), 120˚, 1/2.55\", Dual Pixel PDAF": "৫০ MP, f/১.৮, ২৪mm (wide), ১/১.৫৬\", PDAF, OIS + ১০ MP, f/২.৪, ৬৭mm (telephoto), PDAF, OIS, ৩x অপটিক্যাল জুম + ১২ MP, f/২.২, ১৩mm (ultrawide), ১২০˚, ১/২.৫৫\", Dual Pixel PDAF",
		"6.8 inches": "৬.৮ ইঞ্চি",
		"8K@30fps, 4K@30/60fps, 1080p@60/240fps, 720p@960fps, HDR10+": "৮K@৩০fps, ৪K@৩০/৬০fps, ১০৮০p@৬০/২৪০fps, ৭২০p@৯৬০fps, HDR১০+",
		"Adreno 830": "Adreno ৮৩০",
		"Android 15, One UI 7": "Android ১৫, One UI ৭",
		"Armor aluminum frame, Gorilla Glass Victus 3 front/back": "Armor অ্যালুমিনিয়াম ফ্রেম, Gorilla Glass Victus ৩ front/back",
		"Available. Released February 07, 2025": "উপলব্ধ. Released ফেব্রুয়ারি ০৭, ২০২৫",
		"Corning Gorilla Glass Victus 3": "Corning Gorilla Glass Victus ৩",
		"Dynamic AMOLED 2X, 120Hz, HDR10+, 2600 nits (peak)": "Dynamic AMOLED ২X, ১২০Hz, HDR১০+, ২৬০০ nits (peak)",
		"Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, ultrasonic), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার",
		"GPS (L1+L5), GLONASS, BDS, GALILEO (E1+E5a), QZSS (L1+L5)": "GPS (L১+L৫), GLONASS, BDS, GALILEO (E১+E৫a), QZSS (L১+L৫)",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"GSM 850 / 900 / 1800 / 1900": "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100": "HSDPA ৮৫০ / ৯০০ / ১৭০০(AWS) / ১৯০০ / ২১০০",
		"IP68 dust/water resistant (up to 1.5m for 30 min)": "IP৬৮ dust/water resistant (up to ১.৫m for ৩০ min)",
		"January 22, 2025": "জানুয়ারি ২২, ২০২৫",
		"LTE Band 1/2/3/4/5/7/8/12/13/17/18/19/20/25/26/28/32/38/39/40/41/66": "LTE Band ১/২/৩/৪/৫/৭/৮/১২/১৩/১৭/১৮/১৯/২০/২৫/২৬/২৮/৩২/৩৮/৩৯/৪০/৪১/৬৬",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"Navy, Silver Shadow, Icy Blue, Mint, Coral Red": "Navy, রূপালী Shadow, Icy নীল, Mint, Coral লাল",
		"Octa-core (2x3.4 GHz Cortex-X5 & 6x3.0 GHz Cortex-A730)": "Octa-core (২x৩.৪ GHz Cortex-X৫ & ৬x৩.০ GHz Cortex-A৭৩০)",
		"Qualcomm SM8750 Snapdragon 8 Gen 4 (3 nm)": "Qualcomm SM৮৭৫০ Snapdragon ৮ Gen ৪ (৩ nm)",
		"Qualcomm Snapdragon 8 Gen 4": "Qualcomm Snapdragon ৮ Gen ৪",
		"SA/NSA - n1/2/3/5/7/8/12/20/25/26/28/38/40/41/66/75/77/78": "SA/NSA - n১/২/৩/৫/৭/৮/১২/২০/২৫/২৬/২৮/৩৮/৪০/৪১/৬৬/৭৫/৭৭/৭৮",
		"SM-S936B, SM-S936B/DS, SM-S936U, SM-S936U1, SM-S936W, SM-S936N, SM-S9360": "SM-S৯৩৬B, SM-S৯৩৬B/DS, SM-S৯৩৬U, SM-S৯৩৬U১, SM-S৯৩৬W, SM-S৯৩৬N, SM-S৯৩৬০",
		"Triple Camera: 50MP Wide + 10MP Telephoto (3x) + 12MP Ultrawide": "Triple Camera: ৫০MP Wide + ১০MP Telephoto (৩x) + ১২MP Ultrawide",
		"UFS 4.0": "UFS ৪.০",
		"USB Type-C 3.2, DisplayPort, OTG": "USB Type-C ৩.২, DisplayPort, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-band, Wi-Fi Direct",
		"Yes - 15W wireless (Qi/PMA), 4.5W reverse wireless": "হ্যাঁ - ১৫W বেতার (Qi/PMA), ৪.৫W reverse বেতার",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s25'
func (s *SpecificationSeederMobileSamsungGalaxyS25) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s25"

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
