package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SpecificationSeederMobileSamsungGalaxyS25Ultra seeds specifications/options for product 'samsung-galaxy-s25-ultra'
type SpecificationSeederMobileSamsungGalaxyS25Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS25Ultra creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS25Ultra() *SpecificationSeederMobileSamsungGalaxyS25Ultra {
	return &SpecificationSeederMobileSamsungGalaxyS25Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S25 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS25Ultra) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.25 W/kg (head), 1.42 W/kg (body)": "১.২৫ W/kg (head), ১.৪২ W/kg (body)",
		"1.26 W/kg (head), 0.64 W/kg (body)": "১.২৬ W/kg (head), ০.৬৪ W/kg (body)",
		"12 GB / 16 GB": "১২ GB / ১৬ GB",
		"12 MP, f/2.2, 26mm (wide), 1.12µm, dual pixel PDAF": "১২ MP, f/২.২, ২৬mm (wide), ১.১২µm, dual pixel PDAF",
		"120Hz": "১২০Hz",
		"1440 x 3120 pixels, 19.5:9 ratio (~498 ppi density)": "১৪৪০ x ৩১২০ pixels, ১৯.৫:৯ ratio (~৪৯৮ ppi density)",
		"162.8 x 77.6 x 8.2 mm (6.41 x 3.06 x 0.32 in)": "১৬২.৮ x ৭৭.৬ x ৮.২ মিমি (৬.৪১ x ৩.০৬ x ০.৩২ in)",
		"200 MP + 10 MP + 50 MP + 50 MP": "২০০ MP + ১০ MP + ৫০ MP + ৫০ MP",
		"200MP (wide, f/1.7) + 10MP (telephoto 3x, f/2.4) + 50MP (periscope 5x, f/3.4) + 50MP (ultrawide, f/1.9)": "২০০MP (wide, f/১.৭) + ১০MP (telephoto ৩x, f/২.৪) + ৫০MP (periscope ৫x, f/৩.৪) + ৫০MP (ultrawide, f/১.৯)",
		"218 g (7.69 oz)": "২১৮ g (৭.৬৯ oz)",
		"256 GB / 512 GB / 1 TB": "২৫৬ GB / ৫১২ GB / ১ TB",
		"2x4.47 GHz (Performance) + 6x3.53 GHz (Efficiency)": "২x৪.৪৭ GHz (Performance) + ৬x৩.৫৩ GHz (Efficiency)",
		"3x + 5x Optical Zoom": "৩x + ৫x অপটিক্যাল জুম",
		"45W wired (PD3.0), 65% in 30 min": "৪৫W তারযুক্ত (PD৩.০), ৬৫% in ৩০ মিনিট",
		"45W wired (PD3.0), 65% in 30 min; 15W wireless (Qi2 Ready); 4.5W reverse wireless": "৪৫W তারযুক্ত (PD৩.০), ৬৫% in ৩০ মিনিট; ১৫W বেতার (Qi২ Ready); ৪.৫W reverse বেতার",
		"4K@30/60fps, 1080p@30fps": "৪K@৩০/৬০fps, ১০৮০p@৩০fps",
		"5.4, A2DP, LE": "৫.৪, A২DP, LE",
		"5000 mAh": "৫০০০ এমএএইচ",
		"6.9 inches": "৬.৯ ইঞ্চি",
		"8K@24/30fps, 4K@30/60/120fps, 1080p@30/60/120/240fps, 10-bit HDR, HDR10+, stereo sound rec., gyro-EIS": "৮K@২৪/৩০fps, ৪K@৩০/৬০/১২০fps, ১০৮০p@৩০/৬০/১২০/২৪০fps, ১০-bit HDR, HDR১০+, stereo sound rec., gyro-EIS",
		"Adreno 830 (1200 MHz)": "Adreno ৮৩০ (১২০০ MHz)",
		"Android 15, One UI 7": "Android ১৫, One UI ৭",
		"Available. Released February 03, 2025": "উপলব্ধ. Released ফেব্রুয়ারি ০৩, ২০২৫",
		"Corning Gorilla Armor 2, Mohs level 6; DX anti-reflective coating": "Corning Gorilla Armor ২, Mohs level ৬; DX anti-reflective coating",
		"Dynamic LTPO AMOLED 2X, 120Hz, HDR10+, 2600 nits (peak)": "Dynamic LTPO AMOLED ২X, ১২০Hz, HDR১০+, ২৬০০ nits (peak)",
		"Fingerprint (under display, ultrasonic), accelerometer, gyro, proximity, compass, barometer; UWB support": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, ultrasonic), অ্যাক্সিলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার; UWB support",
		"GSM / CDMA / HSPA / EVDO / LTE / 5G": "GSM / CDMA / HSPA / EVDO / LTE / ৫G",
		"GSM 850 / 900 / 1800 / 1900 MHz": "GSM ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ MHz",
		"Glass front (Gorilla Armor 2), glass back (Victus 2), titanium frame (grade 5)": "গ্লাস সামনে (Gorilla Armor ২), গ্লাস পেছনে (Victus ২), টাইটানিয়াম ফ্রেম (grade ৫)",
		"HSPA 850 / 900 / 1900 / 2100 MHz; EVDO Rev. A 800 / 1900 / 2100 MHz": "HSPA ৮৫০ / ৯০০ / ১৯০০ / ২১০০ MHz; EVDO Rev. A ৮০০ / ১৯০০ / ২১০০ MHz",
		"IP68 dust tight and water resistant (immersible up to 1.5m for 30 min)": "IP৬৮ dust tight and water resistant (immersible up to ১.৫m for ৩০ min)",
		"January 22, 2025": "জানুয়ারি ২২, ২০২৫",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 14, 17, 20, 32, 66": "LTE Band ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৪, ১৭, ২০, ৩২, ৬৬",
		"Li-Ion 5000 mAh": "Li-Ion ৫০০০ mAh",
		"Nano-SIM + Nano-SIM + eSIM + eSIM (max 2 at a time)": "ন্যানো-সিম + ন্যানো-সিম + ই-সিম + ই-সিম (max ২ at a time)",
		"Octa-core (2x4.47 GHz Oryon V2 Phoenix L + 6x3.53 GHz Oryon V2 Phoenix M)": "Octa-core (২x৪.৪৭ GHz Oryon V২ Phoenix L + ৬x৩.৫৩ GHz Oryon V২ Phoenix M)",
		"Qualcomm SM8750-AC Snapdragon 8 Elite (3 nm)": "Qualcomm SM৮৭৫০-AC Snapdragon ৮ Elite (৩ nm)",
		"SM-S938B, SM-S938B/DS, SM-S938U, SM-S938U1, SM-S938W, SM-S938N, SM-S9380, SM-S938E, SM-S938E/DS": "SM-S৯৩৮B, SM-S৯৩৮B/DS, SM-S৯৩৮U, SM-S৯৩৮U১, SM-S৯৩৮W, SM-S৯৩৮N, SM-S৯৩৮০, SM-S৯৩৮E, SM-S৯৩৮E/DS",
		"Samsung DeX, Samsung Wireless DeX, Ultra Wideband (UWB) support": "Samsung DeX, Samsung বেতার DeX, Ultra Wideband (UWB) support",
		"Snapdragon 8 Elite": "Snapdragon ৮ Elite",
		"Titanium Silver Blue, Titanium Black, Titanium White Silver, Titanium Gray, Titanium Jade Green, Titanium Jet Black, Titanium Pink Gold": "Titanium রূপালী নীল, Titanium কালো, Titanium সাদা রূপালী, Titanium ধূসর, Titanium Jade সবুজ, Titanium Jet কালো, Titanium গোলাপী সোনালী",
		"UFS 4.0": "UFS ৪.০",
		"USB Type-C 3.2, DisplayPort 1.2, OTG": "USB Type-C ৩.২, DisplayPort ১.২, OTG",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct": "Wi-Fi ৮০২.১১ a/b/g/n/ac/৬e/৭, tri-band, Wi-Fi Direct",
		"Yes - 15W wireless (Qi2 Ready); 4.5W reverse wireless": "হ্যাঁ - ১৫W বেতার (Qi২ Ready); ৪.৫W reverse বেতার",
		"Yes, stereo speakers": "হ্যাঁ, স্টেরিও স্পিকার",
		"n1, n3, n5, n7, n8, n20, n28, n32, n38, n40, n41, n48, n77, n78, n79": "n১, n৩, n৫, n৭, n৮, n২০, n২৮, n৩২, n৩৮, n৪০, n৪১, n৪৮, n৭৭, n৭৮, n৭৯",
	}
}

// Seed inserts specification_translations for existing specifications for product 'samsung-galaxy-s25-ultra'
func (s *SpecificationSeederMobileSamsungGalaxyS25Ultra) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s25-ultra"

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
