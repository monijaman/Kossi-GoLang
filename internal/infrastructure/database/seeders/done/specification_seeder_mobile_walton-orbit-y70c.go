package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileWaltonORBITY70C seeds specifications/options for product 'walton-orbit-y70c'
type SpecificationSeederMobileWaltonORBITY70C struct {
	BaseSeeder
}

// NewSpecificationSeederMobileWaltonORBITY70C creates a new seeder instance
func NewSpecificationSeederMobileWaltonORBITY70C() *SpecificationSeederMobileWaltonORBITY70C {
	return &SpecificationSeederMobileWaltonORBITY70C{BaseSeeder: BaseSeeder{name: "Specifications for Walton ORBIT Y70C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileWaltonORBITY70C) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"Accelerometer, proximity":           "অ্যাক্সেলেরোমিটার, প্রক্সিমিটি",
		"GPS, GLONASS, GALILEO, BDS":         "জিপিএস, গ্লোনাস, গ্যালিলিও, বিডিএস",
		"5.0":                                "৫.০",
		"Wi-Fi 802.11 a/b/g/n, dual-band":    "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন, ডুয়াল-ব্যান্ড",
		"USB Type-C 2.0":                     "ইউএসবি টাইপ-সি ২.০",
		"Dual SIM (Nano-SIM, dual stand-by)": "ডুয়াল সিম (ন্যানো-সিম, ডুয়াল স্ট্যান্ড-বাই)",
		"Yes":                                "হ্যাঁ",
		"Single speaker":                     "সিঙ্গেল স্পিকার",
		"24-bit/192kHz audio":                "২৪-বিট/১৯২কেএইচজেড অডিও",
		"Li-Po":                              "লি-পো",
		"10W wired":                          "১০ডব্লিউ ওয়্যার্ড",
		"microSDXC (dedicated slot)":         "মাইক্রোএসডিএক্সসি (ডেডিকেটেড স্লট)",
		"32 GB":                              "৩২ জিবি",
		"2.0 GHz":                            "২.০ গিগাহার্টজ",
		"LED flash, HDR":                     "এলইডি ফ্ল্যাশ, এইচডিআর",
		"8 MP, (wide)\n0.3 MP, (depth)":      "৮ এমপি, (ওয়াইড)\n০.৩ এমপি, (ডেপথ)",
		"1080p@30fps":                        "১০৮০পি@৩০এফপিএস",
		"None":                               "কোনটি না",
		"No":                                 "না",
		"0.43 W/kg (head) 1.29 W/kg (body)":  "০.৪৩ ডব্লিউ/কেজি (হেড) ১.২৯ ডব্লিউ/কেজি (বডি)",
		"0.32 W/kg (head) 1.50 W/kg (body)":  "০.৩২ ডব্লিউ/কেজি (হেড) ১.৫০ ডব্লিউ/কেজি (বডি)",
		"ORBIT Y70C (3GB RAM + 32GB)":        "অরবিট ওয়াই৭০সি (৩জিবি র্যাম + ৩২জিবি)",
		"Face unlock":                        "ফেস আনলক",
		"GSM 850 / 900 / 1800 / 1900":        "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 2100":             "এইচএসডিপিএ ৮৫০ / ৯০০ / ২১০০",
		"LTE":                                "এলটিই",
		"Glass front":                        "গ্লাস ফ্রন্ট",
		"8 MP + 0.3 MP":                      "৮ এমপি + ০.৩ এমপি",
		"5 MP":                               "৫ এমপি",
		"4,000 mAh":                          "৪,০০০ এমএএইচ",
		"60Hz":                               "৬০হার্টজ",
		"164.5 x 76 x 9 mm":                  "১৬৪.৫ x ৭৬ x ৯ মিমি",
		"4G":                                 "৪জি",
		"June 2023":                          "জুন ২০২৩",
		"3 GB":                               "৩ জিবি",
		"Plastic body":                       "প্লাস্টিক বডি",
		"188 g":                              "১৮৮ গ্রাম",
		"6.5 inches":                         "৬.৫ ইঞ্চি",
		"Helio A22":                          "হেলিও এ২২",
		"Mediatek Helio A22 (12 nm)":         "মিডিয়াটেক হেলিও এ২২ (১২ ন্যানোমিটার)",
		"Quad-core":                          "কোয়াড-কোর",
		"IPS LCD":                            "আইপিএস এলসিডি",
		"Android 12 Go Edition":              "অ্যান্ড্রয়েড ১২ গো এডিশন",
		"Blue, Green":                        "ব্লু, গ্রিন",
		"Available":                          "উপলব্ধ",
		"PowerVR GE8320":                     "পাওয়ারভিআর জিই৮৩২০",
		"720 x 1600 pixels":                  "৭২০ x ১৬০০ পিক্সেল",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'walton-orbit-y70c'
func (s *SpecificationSeederMobileWaltonORBITY70C) Seed(db *gorm.DB) error {
	productSlug := "walton-orbit-y70c"
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
		"screen_protection":             "Glass front",
		"water_resistance":              "No",
		"rear_camera":                   "8 MP + 0.3 MP",
		"front_camera":                  "5 MP",
		"battery":                       "4,000 mAh",
		"refresh_rate":                  "60Hz",
		"dimensions":                    "164.5 x 76 x 9 mm",
		"network_technology":            "4G",
		"announcement_date":             "June 2023",
		"ram":                           "3 GB",
		"build_material":                "Plastic body",
		"weight":                        "188 g",
		"display_size":                  "6.5 inches",
		"processor":                     "Helio A22",
		"chipset":                       "Mediatek Helio A22 (12 nm)",
		"cpu_type":                      "Quad-core",
		"display_type":                  "IPS LCD",
		"operating_system":              "Android 12 Go Edition",
		"available_colors":              "Blue, Green",
		"device_status":                 "Available",
		"gpu_type":                      "PowerVR GE8320",
		"storage":                       "32 GB",
		"resolution":                    "720 x 1600 pixels",
		"sensors":                       "Accelerometer, proximity",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"bluetooth_version":             "5.0",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n, dual-band",
		"usb_type":                      "USB Type-C 2.0",
		"sim_card_type":                 "Dual SIM (Nano-SIM, dual stand-by)",
		"audio_jack":                    "Yes",
		"loudspeaker_quality":           "Single speaker",
		"audio_quality":                 "24-bit/192kHz audio",
		"battery_type":                  "Li-Po",
		"charging_speed":                "10W wired",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"internal_memory_capacity":      "32 GB",
		"processor_speed":               "2.0 GHz",
		"camera_features":               "LED flash, HDR",
		"quad_camera_setup":             "8 MP, (wide)\n0.3 MP, (depth)",
		"camera_video_resolution":       "1080p@30fps",
		"optical_zoom":                  "None",
		"front_camera_video_resolution": "1080p@30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.43 W/kg (head) 1.29 W/kg (body)",
		"sar_rating_eu":                 "0.32 W/kg (head) 1.50 W/kg (body)",
		"model_variants":                "ORBIT Y70C (3GB RAM + 32GB)",
		"special_features":              "Face unlock",
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
