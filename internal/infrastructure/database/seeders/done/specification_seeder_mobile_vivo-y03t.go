package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY03t seeds specifications/options for product 'vivo-y03t'
type SpecificationSeederMobileVivoY03t struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY03t creates a new seeder instance
func NewSpecificationSeederMobileVivoY03t() *SpecificationSeederMobileVivoY03t {
	return &SpecificationSeederMobileVivoY03t{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y03t"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY03t) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"6.56 inches": "6.56 ইঞ্চি",
		"Glass front, plastic back, plastic frame": "গ্লাস ফ্রন্ট",
		"13 MP + 0.08 MP":                          "13+0.08 মেগা",
		"Space Black, Gem Green":                   "স্পেস ব্ল্যাক, গেম গ্রিন",
		"185 g":                                    "185 গ্রাম",
		"GSM / HSPA / LTE":                         "GSM/LTE",
		"Unisoc T612":                              "Unisoc T612",
		"Mali-G57":                                 "Mali-G57",
		"4 GB":                                     "4 জিবি",
		"720 x 1612 pixels":                        "720x1612",
		"Android 14, Funtouch 14":                  "অ্যান্ড্রয়েড 14",
		"Available":                                "উপলব্ধ",
		"Unisoc T612 (12 nm)":                      "Unisoc T612",
		"IP54":                                     "IP54",
		"15W wired":                                "15W",
		"Octa-core":                                "অক্টা-কোর",
		"64 GB / 128 GB":                           "64/128 জিবি",
		"IPS LCD, 90Hz":                            "IPS LCD 90Hz",
		"90Hz":                                     "90Hz",
		"5 MP":                                     "5 মেগা",
		"5000 mAh":                                 "5000 mAh",
		"August 2024":                              "আগস্ট 2024",
		"164.4 x 75.6 x 8.3 mm":                    "164.4x75.6x8.3 মিমি",
		"Fingerprint (side-mounted), accelerometer, proximity": "ফিঙ্গারপ্রিন্ট (সাইড), অ্যাক্সেলেরোমিটার, প্রক্সিমিটি",
		"GPS, GLONASS, GALILEO, BDS":                           "GPS, GLONASS, GALILEO, BDS",
		"5.1":                                                  "5.1",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band":                   "ওয়াই-ফাই 802.11 a/b/g/n/ac, ডুয়াল-ব্যান্ড",
		"USB Type-C 2.0":                                       "USB Type-C 2.0",
		"Dual SIM (Nano-SIM, dual stand-by)":                   "ডুয়াল SIM (ন্যানো-SIM, ডুয়াল স্ট্যান্ড-বাই)",
		"Yes":                                                  "হ্যাঁ",
		"24-bit/192kHz audio":                                  "24-বিট/192kHz অডিও",
		"Li-Po":                                                "Li-Po",
		"15W":                                                  "15W",
		"microSDXC (dedicated slot)":                           "মাইক্রোSDXC (ডেডিকেটেড স্লট)",
		"64/128 GB":                                            "64/128 GB",
		"1.8 GHz":                                              "1.8 GHz",
		"LED flash, HDR, panorama":                             "LED ফ্ল্যাশ, HDR, প্যানোরামা",
		"13 MP, f/2.2, (wide), 1/3.06\", 1.12µm, PDAF\n0.08 MP, (depth)": "13 MP, f/2.2, (ওয়াইড), PDAF\n0.08 MP, (ডেপ্থ)",
		"1080p@30fps":                       "1080p@30fps",
		"No":                                "না",
		"720p@30fps":                        "720p@30fps",
		"0.99 W/kg (head) 0.99 W/kg (body)": "0.99 W/kg (হেড) 0.99 W/kg (বডি)",
		"0.34 W/kg (head) 1.59 W/kg (body)": "0.34 W/kg (হেড) 1.59 W/kg (বডি)",
		"V2307 (4GB RAM)":                   "V2307 (4GB RAM)",
		"Face unlock":                       "ফেস আনলক",
	}
	for k, v := range s.GetCommonBanglaTranslations() {
		if _, exists := translations[k]; !exists {
			translations[k] = v
		}
	}
	return translations
}

// Seed inserts specification records for the product identified by slug 'vivo-y03t'
func (s *SpecificationSeederMobileVivoY03t) Seed(db *gorm.DB) error {
	productSlug := "vivo-y03t"
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
		"display_size":                  "6.56 inches",
		"build_material":                "Glass front, plastic back, plastic frame",
		"rear_camera":                   "13 MP + 0.08 MP",
		"available_colors":              "Space Black, Gem Green",
		"weight":                        "185 g",
		"network_technology":            "GSM / HSPA / LTE",
		"processor":                     "Unisoc T612",
		"gpu_type":                      "Mali-G57",
		"ram":                           "4 GB",
		"resolution":                    "720 x 1612 pixels",
		"operating_system":              "Android 14, Funtouch 14",
		"device_status":                 "Available",
		"chipset":                       "Unisoc T612 (12 nm)",
		"water_resistance":              "IP54",
		"fast_charging":                 "15W wired",
		"cpu_type":                      "Octa-core",
		"storage":                       "64 GB / 128 GB",
		"display_type":                  "IPS LCD, 90Hz",
		"refresh_rate":                  "90Hz",
		"front_camera":                  "5 MP",
		"battery":                       "5000 mAh",
		"announcement_date":             "August 2024",
		"dimensions":                    "164.4 x 75.6 x 8.3 mm",
		"sensors":                       "Fingerprint (side-mounted), accelerometer, proximity",
		"positioning_system":            "GPS, GLONASS, GALILEO, BDS",
		"bluetooth_version":             "5.1",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac, dual-band",
		"usb_type":                      "USB Type-C 2.0",
		"sim_card_type":                 "Dual SIM (Nano-SIM, dual stand-by)",
		"audio_jack":                    "Yes",
		"loudspeaker_quality":           "Yes",
		"audio_quality":                 "24-bit/192kHz audio",
		"battery_type":                  "Li-Po",
		"charging_speed":                "15W",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"internal_memory_capacity":      "64/128 GB",
		"processor_speed":               "1.8 GHz",
		"camera_features":               "LED flash, HDR, panorama",
		"quad_camera_setup":             "13 MP, f/2.2, (wide), 1/3.06\", 1.12µm, PDAF\n0.08 MP, (depth)",
		"camera_video_resolution":       "1080p@30fps",
		"optical_zoom":                  "No",
		"front_camera_video_resolution": "720p@30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.99 W/kg (head) 0.99 W/kg (body)",
		"sar_rating_eu":                 "0.34 W/kg (head) 1.59 W/kg (body)",
		"model_variants":                "V2307 (4GB RAM)",
		"special_features":              "Face unlock",
	}
	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
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
