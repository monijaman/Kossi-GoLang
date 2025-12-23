package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonNEXGN75 seeds specifications/options for product 'walton-nexg-n75'
type SpecificationSeederMobileWaltonNEXGN75 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonNEXGN75 creates a new seeder instance
func NewSpecificationSeederMobileWaltonNEXGN75() *SpecificationSeederMobileWaltonNEXGN75 {
	return &SpecificationSeederMobileWaltonNEXGN75{BaseSeeder: BaseSeeder{name: "Specifications for Walton NEXG N75"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonNEXGN75) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"Fingerprint (rear-mounted), accelerometer, proximity": "ফিঙ্গারপ্রিন্ট (রিয়ার), অ্যাক্সেলেরোমিটার, প্রক্সিমিটি",
		"GPS, GLONASS, GALILEO, BDS":                           "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"5.1":                                                  "৫.১",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band":                   "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি, ডুয়াল-ব্যান্ড",
		"USB Type-C 2.0":                                       "ইউএসবি টাইপ-সি ২.০",
		"Dual SIM (Nano-SIM, dual stand-by)":                   "ডুয়াল সিম (ন্যানো-সিম, ডুয়াল স্ট্যান্ড-বাই)",
		"Yes":                                                  "হ্যাঁ",
		"Single speaker":                                       "সিঙ্গেল স্পিকার",
		"24-bit/192kHz audio":                                  "২৪-বিট/১৯২কেএইচজেড অডিও",
		"Li-Po":                                                "লি-পো",
		"10W wired":                                            "১০ডব্লিউ ওয়্যার্ড",
		"microSDXC (dedicated slot)":                           "মাইক্রোএসডিএক্সসি (ডেডিকেটেড স্লট)",
		"128 GB":                                               "১২৮ জিবি",
		"1.6 GHz":                                              "১.৬ গিগাহার্টজ",
		"LED flash, HDR, panorama":                             "এলইডি ফ্ল্যাশ, এইচডিআর, প্যানোরামা",
		"50 MP, f/1.8, (wide), 1/2.76\", 0.64µm, PDAF\n2 MP, f/2.4, (depth)": "৫০ এমপি, এফ/১.৮, (ওয়াইড), ১/২.৭৬\", ০.৬৪µম, পিডিএএফ\n২ এমপি, এফ/২.৪, (ডেপথ)",
		"1080p@30fps":                       "১০৮০পি@৩০এফপিএস",
		"None":                              "কোনটি না",
		"No":                                "না",
		"0.43 W/kg (head) 1.29 W/kg (body)": "০.৪৩ ডব্লিউ/কেজি (হেড) ১.২৯ ডব্লিউ/কেজি (বডি)",
		"0.32 W/kg (head) 1.50 W/kg (body)": "০.৩২ ডব্লিউ/কেজি (হেড) ১.৫০ ডব্লিউ/কেজি (বডি)",
		"NEXG N75 (8GB RAM + 128GB)":        "এনইএক্সজি এন৭৫ (৮জিবি র্যাম + ১২৮জিবি)",
		"Face unlock, FM radio":             "ফেস আনলক, এফএম রেডিও",
		"GSM 850 / 900 / 1800 / 1900":       "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 2100":            "এইচএসডিপিএ ৮৫০ / ৯০০ / ২১০০",
		"LTE":                               "এলটিই",
		"4G":                                "4G",
		"February 2024":                     "ফেব্রুয়ারি 2024",
		"6.6 inches":                        "6.6 ইঞ্চি",
		"IPS LCD, 90Hz":                     "IPS LCD 90Hz",
		"90Hz":                              "90Hz",
		"50 MP + 2 MP":                      "50+2 মেগা",
		"Android 13":                        "অ্যান্ড্রয়েড 13",
		"Unisoc T606":                       "ইউনিসক T606",
		"Octa-core":                         "অক্টা-কোর",
		"8 MP":                              "8 মেগা",
		"Mali-G57 MP1":                      "মালি-G57 MP1",
		"720 x 1612 pixels":                 "720x1612",
		"Plastic body":                      "প্লাস্টিক বডি",
		"190 g":                             "190 গ্রাম",
		"5,000 mAh":                         "5000 mAh",
		"Golden, Black":                     "গোল্ডেন, ব্ল্যাক",
		"Available":                         "উপলব্ধ",
		"Unisoc T606 (12 nm)":               "ইউনিসক T606",
		"8 GB":                              "8 জিবি",
		"Glass front":                       "গ্লাস ফ্রন্ট",
		"164.2 x 76 x 8.5 mm":               "164.2x76x8.5 মিমি",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'walton-nexg-n75'
func (s *SpecificationSeederMobileWaltonNEXGN75) Seed(db *gorm.DB) error {
	productSlug := "walton-nexg-n75"
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
		"processor":                     "Unisoc T606",
		"ram":                           "8 GB",
		"resolution":                    "720 x 1612 pixels",
		"screen_protection":             "Glass front",
		"dimensions":                    "164.2 x 76 x 8.5 mm",
		"available_colors":              "Golden, Black",
		"device_status":                 "Available",
		"gpu_type":                      "Mali-G57 MP1",
		"display_type":                  "IPS LCD, 90Hz",
		"refresh_rate":                  "90Hz",
		"build_material":                "Plastic body",
		"battery":                       "5,000 mAh",
		"announcement_date":             "February 2024",
		"cpu_type":                      "Octa-core",
		"water_resistance":              "No",
		"network_technology":            "4G",
		"rear_camera":                   "50 MP + 2 MP",
		"front_camera":                  "8 MP",
		"display_size":                  "6.6 inches",
		"chipset":                       "Unisoc T606 (12 nm)",
		"storage":                       "128 GB",
		"weight":                        "190 g",
		"operating_system":              "Android 13",
		"sensors":                       "Fingerprint (rear-mounted), accelerometer, proximity",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"bluetooth_version":             "5.1",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac, dual-band",
		"usb_type":                      "USB Type-C 2.0",
		"sim_card_type":                 "Dual SIM (Nano-SIM, dual stand-by)",
		"audio_jack":                    "Yes",
		"loudspeaker_quality":           "Single speaker",
		"audio_quality":                 "24-bit/192kHz audio",
		"battery_type":                  "Li-Po",
		"charging_speed":                "10W wired",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"internal_memory_capacity":      "128 GB",
		"processor_speed":               "1.6 GHz",
		"camera_features":               "LED flash, HDR, panorama",
		"quad_camera_setup":             "50 MP, f/1.8, (wide), 1/2.76\", 0.64µm, PDAF\n2 MP, f/2.4, (depth)",
		"camera_video_resolution":       "1080p@30fps",
		"optical_zoom":                  "None",
		"front_camera_video_resolution": "1080p@30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.43 W/kg (head) 1.29 W/kg (body)",
		"sar_rating_eu":                 "0.32 W/kg (head) 1.50 W/kg (body)",
		"model_variants":                "NEXG N75 (8GB RAM + 128GB)",
		"special_features":              "Face unlock, FM radio",
		"fast_charging":                 "Yes",
		"wireless_charging":             "No",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                      "HSDPA 850 / 900 / 2100",
		"4g_bands":                      "LTE",
		"5g_bands":                      "No",
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
