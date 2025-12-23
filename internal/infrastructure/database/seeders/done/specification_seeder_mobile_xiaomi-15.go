package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi15 seeds specifications/options for product 'xiaomi-15'
type SpecificationSeederMobileXiaomi15 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15 creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15() *SpecificationSeederMobileXiaomi15 {
	return &SpecificationSeederMobileXiaomi15{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"Xiaomi Dragon Crystal Glass": "শাওমি ড্রাগন ক্রিস্টাল গ্লাস",
		"5G":                          "৫জি",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)": "কোয়ালকম এসএম৮৭৫০ স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"Android 15, HyperOS":                       "অ্যান্ড্রয়েড ১৫, হাইপারওএস",
		"Black, White, Green, Purple, Silver":       "ব্ল্যাক, হোয়াইট, গ্রিন, পার্পল, সিলভার",
		"191 g":                                     "১৯১ গ্রাম",
		"152.8 x 71.5 x 8.2 mm":                     "১৫২.৮ x ৭১.৫ x ৮.২ মিমি",
		"6.36 inches":                               "৬.৩৬ ইঞ্চি",
		"Adreno 830":                                "অ্যাড্রেনো ৮৩০",
		"12 GB / 16 GB":                             "১২ জিবি / ১৬ জিবি",
		"LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3200 nits": "এলটিপিও ওএলইডি, ১২০হার্টজ, ডলবি ভিশন, এইচডিআর১০+, ৩২০০ নিটস",
		"IP68":                             "আইপি৬৮",
		"50 MP + 50 MP + 50 MP":            "৫০ এমপি + ৫০ এমপি + ৫০ এমপি",
		"Octa-core":                        "অক্টা-কোর",
		"120Hz":                            "১২০হার্টজ",
		"32 MP":                            "৩২ এমপি",
		"5,400 mAh":                        "৫,৪০০ এমএএইচ",
		"October 2024":                     "অক্টোবর ২০২৪",
		"Available":                        "উপলব্ধ",
		"Glass front/back, aluminum frame": "গ্লাস ফ্রন্ট/ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"Snapdragon 8 Elite":               "স্ন্যাপড্রাগন ৮ এলিট",
		"256 GB / 512 GB / 1 TB":           "২৫৬ জিবি / ৫১২ জিবি / ১ টিবি",
		"1200 x 2670 pixels":               "১২০০ x ২৬৭০ পিক্সেল",
		"Li-Po":                            "লি-পো",
		"90W wired, 30W wireless":          "৯০ডব্লিউ ওয়্যার্ড, ৩০ডব্লিউ ওয়্যারলেস",
		"Yes":                              "হ্যাঁ",
		"24-bit/192kHz":                    "২৪-বিট/১৯২কেএইচজেড",
		"Stereo":                           "স্টেরিও",
		"No":                               "না",
		"GSM 850 / 900 / 1800 / 1900":      "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE":                          "এলটিই",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7": "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই/৭",
		"5.4":                          "৫.৪",
		"USB Type-C 3.2":               "ইউএসবি টাইপ-সি ৩.২",
		"GPS, GLONASS, BDS, GALILEO":   "জিপিএস, গ্লোনাস, বিডিএস, গ্যালিলিও",
		"Fingerprint, accelerometer, gyro, proximity, compass, barometer": "ফিঙ্গারপ্রিন্ট, অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার",
		"AI image processing, Leica camera mode":                          "এআই ইমেজ প্রসেসিং, লাইকা ক্যামেরা মোড",
		"Dual Nano-SIM":                                                   "ডুয়াল ন্যানো-সিম",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'xiaomi-15'
func (s *SpecificationSeederMobileXiaomi15) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15"
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
		"screen_protection":   "Xiaomi Dragon Crystal Glass",
		"network_technology":  "5G",
		"chipset":             "Qualcomm SM8750 Snapdragon 8 Elite (3 nm)",
		"operating_system":    "Android 15, HyperOS",
		"available_colors":    "Black, White, Green, Purple, Silver",
		"weight":              "191 g",
		"dimensions":          "152.8 x 71.5 x 8.2 mm",
		"display_size":        "6.36 inches",
		"gpu_type":            "Adreno 830",
		"ram":                 "12 GB / 16 GB",
		"display_type":        "LTPO OLED, 120Hz, Dolby Vision, HDR10+, 3200 nits",
		"water_resistance":    "IP68",
		"rear_camera":         "50 MP + 50 MP + 50 MP",
		"cpu_type":            "Octa-core",
		"refresh_rate":        "120Hz",
		"front_camera":        "32 MP",
		"battery":             "5,400 mAh",
		"announcement_date":   "October 2024",
		"device_status":       "Available",
		"build_material":      "Glass front/back, aluminum frame",
		"processor":           "Snapdragon 8 Elite",
		"storage":             "256 GB / 512 GB / 1 TB",
		"resolution":          "1200 x 2670 pixels",
		"battery_type":        "Li-Po",
		"fast_charging":       "90W wired, 30W wireless",
		"charging_speed":      "90W wired, 30W wireless",
		"wireless_charging":   "Yes",
		"audio_quality":       "24-bit/192kHz",
		"loudspeaker_quality": "Stereo",
		"audio_jack":          "No",
		"2g_bands":            "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":            "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":            "LTE",
		"5g_bands":            "5G",
		"wifi_support":        "Wi-Fi 802.11 a/b/g/n/ac/6e/7",
		"bluetooth_version":   "5.4",
		"nfc_support":         "Yes",
		"usb_type":            "USB Type-C 3.2",
		"positioning_system":  "GPS, GLONASS, BDS, GALILEO",
		"sensors":             "Fingerprint, accelerometer, gyro, proximity, compass, barometer",
		"special_features":    "AI image processing, Leica camera mode",
		"sim_card_type":       "Dual Nano-SIM",
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
