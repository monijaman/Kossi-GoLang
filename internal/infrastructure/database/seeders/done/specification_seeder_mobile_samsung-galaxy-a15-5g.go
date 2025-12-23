package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA155G seeds specifications/options for product 'samsung-galaxy-a15-5g'
type SpecificationSeederMobileSamsungGalaxyA155G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA155G creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA155G() *SpecificationSeederMobileSamsungGalaxyA155G {
	return &SpecificationSeederMobileSamsungGalaxyA155G{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A15 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA155G) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"6.5 inches":                      "৬.৫ ইঞ্চি",
		"Super AMOLED":                    "সুপার অ্যামোলেড",
		"1080 x 2340 pixels":              "১০৮০ × ২৩৪০ পিক্সেল",
		"90Hz":                            "৯০ হার্জ",
		"Corning Gorilla Glass 5":         "কর্নিং গরিলা গ্লাস ৫",
		"MediaTek Dimensity 6100+":        "মিডিয়াটেক ডাইমেনসিটি ৬১০০+",
		"MediaTek Dimensity 6100+ (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৬১০০+ (৬ ন্যানোমিটার)",
		"Octa-core":                       "অক্টা-কোর",
		"Mali-G57 MC2":                    "মালি জি৫৭ এমসি২",
		"2.2 GHz":                         "২.২ গিগাহার্জ",
		"4 GB / 6 GB / 8 GB":              "৪ জিবি / ৬ জিবি / ৮ জিবি",
		"128 GB / 256 GB":                 "১২৮ জিবি / ২৫৬ জিবি",
		"microSDXC (dedicated slot)":      "মাইক্রোএসডি (নির্দিষ্ট স্লট)",
		"50 MP + 5 MP + 2 MP":             "৫০ মেগাপিক্সেল + ৫ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"Autofocus, LED flash":            "অটোফোকাস, এলইডি ফ্ল্যাশ",
		"50 MP (wide) + 5 MP (ultrawide) + 2 MP (macro)": "৫০ মেগাপিক্সেল (ওয়াইড) + ৫ মেগাপিক্সেল (আল্ট্রাওয়াইড) + ২ মেগাপিক্সেল (ম্যাক্রো)",
		"1080p @ 30fps":          "১০৮০পি @ ৩০ এফপিএস",
		"None":                   "কোনোটিই নেই",
		"13 MP":                  "১৩ মেগাপিক্সেল",
		"5000 mAh":               "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)": "লি-আয়ন (অপসারণযোগ্য নয়)",
		"25W wired":              "২৫ ওয়াট তারযুক্ত",
		"No":                     "না",
		"Android 14, One UI 6":   "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই ৬",
		"Stereo speakers":        "স্টেরিও স্পিকার",
		"Mono speaker":           "মনো স্পিকার",
		"3.5 mm headphone jack":  "৩.৫ মিলিমিটার হেডফোন জ্যাক",
		"Glass front, plastic back, plastic frame": "গ্লাস ফ্রন্ট, প্লাস্টিক ব্যাক, প্লাস্টিক ফ্রেম",
		"200 g":                                     "২০০ গ্রাম",
		"160.1 x 76.8 x 8.4 mm":                     "১৬০.১ × ৭৬.৮ × ৮.৪ মিলিমিটার",
		"GSM / HSPA / LTE / 5G":                     "জিএসএম / এইচএসপিএ / এলটিই / ৫জি",
		"GSM 850 / 900 / 1800 / 1900":               "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০ (এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)": "এলটিই ব্যান্ড ১(২১০০), ২(১৯০০), ৩(১৮০০), ４(১৭০০/২১০০), ৫(৮৫০), ৭(২৬০০), ৮(৯০০), ১২(৭০০), ১৩(৭০০), ১৭(৭০০), ২০(৮০০), ২৬(৮৫০), ২৮(৭০০), ৩২(১৫০০), ৩৮(২৬০০), ৪০(২৩০০), ৪১(২৫০০), ৬৬(১৭০০/২১০০)",
		"SA/NSA/Sub6":                "এসএ/এনএসএ/সাব৬",
		"Yes":                        "হ্যাঁ",
		"Wi-Fi 802.11 a/b/g/n/ac":    "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি",
		"5.3":                        "৫.৩",
		"Yes (market dependent)":     "হ্যাঁ (বাজার-নির্ভর)",
		"USB-C 2.0":                  "ইউএসবি-সি ২.০",
		"GPS, GLONASS, GALILEO, BDS": "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"Fingerprint (side-mounted), accelerometer, gyro, proximity, compass": "ফিঙ্গারপ্রিন্ট (সাইড-মাউন্টেড), অ্যাকসেলারোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস",
		"Face unlock":                        "ফেস আনলক",
		"Dual SIM (Nano-SIM, dual stand-by)": "ডুয়াল সিম (ন্যানো-সিম, ডুয়াল স্ট্যান্ড-বাই)",
		"0.41 W/kg (head) 1.35 W/kg (body)":  "০.৪১ ওয়াট/কেজি (হেড) ১.৩৫ ওয়াট/কেজি (বডি)",
		"Brave Black, Optimistic Blue, Magical Blue, Personality Yellow": "ব্রেভ ব্ল্যাক, অপটিমিস্টিক ব্লু, ম্যাজিকাল ব্লু, পার্সোনালিটি ইয়েলো",
		"SM-A156B, SM-A156E, SM-A156U1":                                  "SM-A156B, SM-A156E, SM-A156U1",
		"December 2023":                                                  "ডিসেম্বর ২০২৩",
		"Available":                                                      "উপলব্ধ",
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

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a15-5g'
func (s *SpecificationSeederMobileSamsungGalaxyA155G) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a15-5g"
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
		"display_size":      "6.5 inches",
		"display_type":      "Super AMOLED",
		"resolution":        "1080 x 2340 pixels",
		"refresh_rate":      "90Hz",
		"screen_protection": "Corning Gorilla Glass 5",
		// Performance
		"processor":       "MediaTek Dimensity 6100+",
		"chipset":         "MediaTek Dimensity 6100+ (6 nm)",
		"cpu_type":        "Octa-core",
		"gpu_type":        "Mali-G57 MC2",
		"processor_speed": "2.2 GHz",
		// Memory & storage
		"ram":                      "4 GB / 6 GB / 8 GB",
		"storage":                  "128 GB / 256 GB",
		"internal_memory_capacity": "128 GB / 256 GB",
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
		"fast_charging":     "25W wired",
		"charging_speed":    "25W wired",
		"wireless_charging": "No",
		// Software
		"operating_system": "Android 14, One UI 6",
		// Audio
		"audio_quality":       "Stereo speakers",
		"loudspeaker_quality": "Mono speaker",
		"audio_jack":          "3.5 mm headphone jack",
		// Build & body
		"build_material":   "Glass front, plastic back, plastic frame",
		"weight":           "200 g",
		"dimensions":       "160.1 x 76.8 x 8.4 mm",
		"water_resistance": "No",
		// Network & connectivity
		"network_technology": "GSM / HSPA / LTE / 5G",
		"2g_bands":           "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":           "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":           "LTE band 1(2100), 2(1900), 3(1800), 4(1700/2100), 5(850), 7(2600), 8(900), 12(700), 13(700), 17(700), 20(800), 26(850), 28(700), 32(1500), 38(2600), 40(2300), 41(2500), 66(1700/2100)",
		"5g_bands":           "SA/NSA/Sub6",
		"5g_support":         "Yes",
		"wifi_support":       "Wi-Fi 802.11 a/b/g/n/ac",
		"bluetooth_version":  "5.3",
		"nfc_support":        "Yes (market dependent)",
		"usb_type":           "USB-C 2.0",
		// Positioning & sensors
		"positioning_system": "GPS, GLONASS, GALILEO, BDS",
		"sensors":            "Fingerprint (side-mounted), accelerometer, gyro, proximity, compass",
		"special_features":   "Face unlock",
		// SIM & compliance
		"sim_card_type": "Dual SIM (Nano-SIM, dual stand-by)",
		"sar_rating":    "0.41 W/kg (head) 1.35 W/kg (body)",
		"sar_rating_eu": "0.41 W/kg (head) 1.35 W/kg (body)",
		// Commercial info
		"available_colors":  "Brave Black, Optimistic Blue, Magical Blue, Personality Yellow",
		"model_variants":    "SM-A156B, SM-A156E, SM-A156U1",
		"announcement_date": "December 2023",
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
