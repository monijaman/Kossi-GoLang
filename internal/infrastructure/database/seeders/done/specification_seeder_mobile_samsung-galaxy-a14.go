package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA14 seeds specifications/options for product 'samsung-galaxy-a14'
type SpecificationSeederMobileSamsungGalaxyA14 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA14 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA14() *SpecificationSeederMobileSamsungGalaxyA14 {
	return &SpecificationSeederMobileSamsungGalaxyA14{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A14"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA14) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.6 inches":              "৬.৬ ইঞ্চি",
		"PLS LCD":                 "পিএলএস এলসিডি",
		"2408 × 1080 px (FHD+)":   "২৪০৮ × ১০৮০ পিক্সেল (এফএইচডি+)",
		"60 Hz":                   "৬০ হার্জ",
		"Corning Gorilla Glass 5": "কর্নিং গরিলা গ্লাস ৫",
		"MediaTek Helio G80 (4G) / Dimensity 700 (5G)": "মিডিয়াটেক হেলিও জি৮০ (৪জি) / ডাইমেনসিটি ৭০০ (৫জি)",
		"Helio G80 (12 nm) / Dimensity 700 (7 nm)":     "হেলিও জি৮০ (১২ ন্যানোমিটার) / ডাইমেনসিটি ৭০০ (৭ ন্যানোমিটার)",
		"Octa-core":                     "অক্টা-কোর",
		"Mali-G52 (4G) / Mali-G57 (5G)": "মালি জি৫২ (৪জি) / মালি জি৫৭ (৫জি)",
		"2.0 GHz":                       "২.০ গিগাহার্জ",
		"4 / 6 GB":                      "৪ / ৬ জিবি",
		"64 / 128 GB + microSD support": "৬৪ / ১২৮ জিবি + মাইক্রোএসডি সমর্থন",
		"64 / 128 GB":                   "৬৪ / ১২৮ জিবি",
		"microSDXC (dedicated slot)":    "মাইক্রোএসডি (নির্দিষ্ট স্লট)",
		"50 MP + 5 MP + 2 MP":           "৫০ মেগাপিক্সেল + ৫ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"Autofocus, LED flash":          "অটোফোকাস, এলইডি ফ্ল্যাশ",
		"50 MP (wide) + 5 MP (ultrawide) + 2 MP (macro)": "৫০ মেগাপিক্সেল (ওয়াইড) + ৫ মেগাপিক্সেল (আল্ট্রাওয়াইড) + ২ মেগাপিক্সেল (ম্যাক্রো)",
		"1080p @ 30fps":                         "১০৮০পি @ ৩০ এফপিএস",
		"None":                                  "কোনোটিই নেই",
		"13 MP":                                 "১৩ মেগাপিক্সেল",
		"5000 mAh":                              "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)":                "লি-আয়ন (অপসারণযোগ্য নয়)",
		"15 W wired":                            "১৫ ওয়াট তারযুক্ত",
		"15W wired":                             "১৫ ওয়াট তারযুক্ত",
		"No":                                    "না",
		"Android 13 (upgradable to Android 14)": "অ্যান্ড্রয়েড ১৩ (অ্যান্ড্রয়েড ১৪-এ আপগ্রেডযোগ্য)",
		"Stereo speakers":                       "স্টেরিও স্পিকার",
		"Mono speaker":                          "মনো স্পিকার",
		"3.5 mm headphone jack":                 "৩.৫ মিলিমিটার হেডফোন জ্যাক",
		"Glass front, plastic back & frame":     "গ্লাস ফ্রন্ট, প্লাস্টিক ব্যাক এবং ফ্রেম",
		"201 g":                                 "২০১ গ্রাম",
		"167.7 × 78.0 × 9.1 mm":                 "১৬৭.৭ × ৭৮.০ × ৯.১ মিলিমিটার",
		"GSM / HSPA / LTE (4G) / 5G (A14 5G variant)": "জিএসএম / এইচএসপিএ / এলটিই (৪জি) / ৫জি (A১৪ ৫জি ভ্যারিয়েন্ট)",
		"GSM 850 / 900 / 1800 / 1900":                 "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100":   "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০ (এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)": "এলটিই ব্যান্ড ১(২১০০), ২(১৯০০), ৩(১৮০০), ৪(১৭০০/২১০০), ৫(৮৫০), ৭(২৬০০), ৮(৯০০), ১২(৭০০), ১৩(৭০০), ১৭(৭০০), ২০(৮০০), ২৬(৮৫০), ২৮(৭০০), ৩২(১৫০০), ৩৮(২৬০০), ৪০(২৩০০), ৪১(২৫০০), ৬৬(১৭০০/২১০০)",
		"SA/NSA/Sub6 (A14 5G variant)": "এসএ/এনএসএ/সাব৬ (A১৪ ৫জি ভ্যারিয়েন্ট)",
		"Yes (only A14 5G variant)":    "হ্যাঁ (শুধুমাত্র A১৪ ৫জি ভ্যারিয়েন্ট)",
		"Wi-Fi 802.11 a/b/g/n/ac":      "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি",
		"5.2 (5G) / 5.0 (4G)":          "৫.২ (৫জি) / ৫.০ (৪জি)",
		"Yes (market dependent)":       "হ্যাঁ (বাজার-নির্ভর)",
		"USB-C 2.0":                    "ইউএসবি-সি ২.০",
		"GPS, GLONASS, GALILEO, BDS":   "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Side fingerprint, Accelerometer, Gyro, Proximity, Compass": "সাইড ফিঙ্গারপ্রিন্ট, অ্যাকসেলারোমিটার, জাইরো, প্রোক্সিমিটি, কম্পাস",
		"Virtual RAM, Eye Comfort Shield":                           "ভার্চুয়াল র‍্যাম, আই কমফোর্ট শিল্ড",
		"Nano-SIM or Dual SIM":                                      "ন্যানো-সিম বা ডুয়াল সিম",
		"0.41 W/kg (head) 1.35 W/kg (body)":                         "০.৪১ ওয়াট/কেজি (হেড) ১.৩৫ ওয়াট/কেজি (বডি)",
		"Black, Silver, Green, Red":                                 "কালো, সিলভার, সবুজ, লাল",
		"SM-A145F (4G), SM-A145M (4G), SM-A145P (5G)":               "SM-A145F (৪জি), SM-A145M (৪জি), SM-A145P (৫জি)",
		"January 2023":                                              "জানুয়ারি ২০২৩",
		"Available":                                                 "উপলব্ধ",
	}
	result := make(map[string]string)
	for k, v := range common {
		result[k] = v
	}
	for k, v := range specific {
		result[k] = v
	}
	return result
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a14'
func (s *SpecificationSeederMobileSamsungGalaxyA14) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a14"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID
	existingkeyMapping := map[string]uint{
		// Core display
		"display_size":      66, // Screen Size
		"display_type":      27, // Display Type
		"resolution":        62, // Resolution
		"refresh_rate":      61, // Refresh Rate
		"screen_protection": 65, // Screen Protection
		// Performance
		"processor":       535, // processor (lowercase)
		"chipset":         19,  // Chipset
		"cpu_type":        21,  // CPU Type
		"gpu_type":        33,  // GPU Type
		"processor_speed": 57,  // Processor Speed
		// Memory & storage
		"ram":                      60, // RAM
		"storage":                  71, // Storage Capacity
		"internal_memory_capacity": 36, // Internal Memory Capacity
		"card_slot_type":           17, // Card Slot Type
		// Cameras
		"rear_camera":                   122, // Rear Camera
		"camera_features":               16,  // Camera Features
		"quad_camera_setup":             58,  // Quad Camera Setup
		"camera_video_resolution":       40,  // Main Camera Video Resolution
		"optical_zoom":                  193, // Optical Zoom
		"front_camera":                  32,  // Front Camera
		"front_camera_video_resolution": 32,  // Front Camera Video Resolution
		// Battery & charging
		"battery":           11,  // Battery Capacity
		"battery_type":      12,  // Battery Type
		"fast_charging":     553, // fast_charging (lowercase)
		"charging_speed":    18,  // Charging Speed
		"wireless_charging": 82,  // Wireless Charging
		// Software
		"operating_system": 49, // Operating System
		// Audio
		"audio_quality":       9,  // Audio Quality
		"loudspeaker_quality": 38, // Loudspeaker Quality
		"audio_jack":          2,  // 3.5mm Audio Jack
		// Build & body
		"build_material":   14, // Build Material
		"weight":           80, // Weight
		"dimensions":       25, // Dimensions
		"water_resistance": 78, // Water Resistance
		// Network & connectivity
		"network_technology": 653, // Network Technology
		"2g_bands":           1,   // 2G Bands
		"3g_bands":           3,   // 3G Bands
		"4g_bands":           4,   // 4G Bands
		"5g_bands":           5,   // 5G Bands
		"5g_support":         5,   // 5G Support (maps to 5G Bands)
		"wifi_support":       81,  // WiFi
		"bluetooth_version":  13,  // Bluetooth Version
		"nfc_support":        46,  // NFC Support
		"usb_type":           76,  // USB Type
		// Positioning & sensors
		"positioning_system": 54, // Positioning System
		"sensors":            67, // Sensors
		"special_features":   69, // Special Features
		// SIM & compliance
		"sim_card_type": 68, // SIM Card Type
		"sar_rating":    63, // SAR (Specific Absorption Rate)
		"sar_rating_eu": 64, // SAR (Specific Absorption Rate) EU
		// Commercial info
		"available_colors":  10, // Available Colors
		"model_variants":    42, // Model Variants
		"announcement_date": 8,  // Announcement Date
		"device_status":     23, // Device Status
	}
	specs := map[string]string{
		// Core display
		"display_size":      "6.6 inches",
		"display_type":      "PLS LCD",
		"resolution":        "2408 × 1080 px (FHD+)",
		"refresh_rate":      "60 Hz",
		"screen_protection": "Corning Gorilla Glass 5",
		// Performance
		"processor":       "MediaTek Helio G80 (4G) / Dimensity 700 (5G)",
		"chipset":         "Helio G80 (12 nm) / Dimensity 700 (7 nm)",
		"cpu_type":        "Octa-core",
		"gpu_type":        "Mali-G52 (4G) / Mali-G57 (5G)",
		"processor_speed": "2.0 GHz",
		// Memory & storage
		"ram":                      "4 / 6 GB",
		"storage":                  "64 / 128 GB + microSD support",
		"internal_memory_capacity": "64 / 128 GB",
		"card_slot_type":           "microSDXC (dedicated slot)",
		// Cameras
		"rear_camera":                   "50 MP + 5 MP + 2 MP",
		"camera_features":               "Autofocus, LED flash",
		"quad_camera_setup":             "50 MP (wide) + 5 MP (ultrawide) + 2 MP (macro)",
		"camera_video_resolution":       "1080p @ 30fps",
		"optical_zoom":                  "None",
		"front_camera":                  "13 MP",
		"front_camera_video_resolution": "1080p @ 30fps",
		// Battery & charging
		"battery":           "5000 mAh",
		"battery_type":      "Li-Ion (non-removable)",
		"fast_charging":     "15 W wired",
		"charging_speed":    "15W wired",
		"wireless_charging": "No",
		// Software
		"operating_system": "Android 13 (upgradable to Android 14)",
		// Audio
		"audio_quality":       "Stereo speakers",
		"loudspeaker_quality": "Mono speaker",
		"audio_jack":          "3.5 mm headphone jack",
		// Build & body
		"build_material":   "Glass front, plastic back & frame",
		"weight":           "201 g",
		"dimensions":       "167.7 × 78.0 × 9.1 mm",
		"water_resistance": "None",
		// Network & connectivity
		"network_technology": "GSM / HSPA / LTE (4G) / 5G (A14 5G variant)",
		"2g_bands":           "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":           "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":           "LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)",
		"5g_bands":           "SA/NSA/Sub6 (A14 5G variant)",
		"5g_support":         "Yes (only A14 5G variant)",
		"wifi_support":       "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":  "5.2 (5G) / 5.0 (4G)",
		"nfc_support":        "Yes (market dependent)",
		"usb_type":           "USB-C 2.0",
		// Positioning & sensors
		"positioning_system": "GPS, GLONASS, GALILEO, BDS",
		"sensors":            "Side fingerprint, Accelerometer, Gyro, Proximity, Compass",
		"special_features":   "Virtual RAM, Eye Comfort Shield",
		// SIM & compliance
		"sim_card_type": "Nano-SIM or Dual SIM",
		"sar_rating":    "0.41 W/kg (head) 1.35 W/kg (body)",
		"sar_rating_eu": "0.41 W/kg (head) 1.35 W/kg (body)",
		// Commercial info
		"available_colors":  "Black, Silver, Green, Red",
		"model_variants":    "SM-A145F (4G), SM-A145M (4G), SM-A145P (5G)",
		"announcement_date": "January 2023",
		"device_status":     "Available",
	}
	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%!s(MISSING)' (used in product: %!s(MISSING))", key, productSlug)
			continue // Skip keys that don't exist in mapping
		}
		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
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
