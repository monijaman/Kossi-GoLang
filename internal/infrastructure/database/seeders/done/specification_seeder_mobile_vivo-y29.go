package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileVivoY29 seeds specifications/options for product 'vivo-y29'
type SpecificationSeederMobileVivoY29 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileVivoY29 creates a new seeder instance
func NewSpecificationSeederMobileVivoY29() *SpecificationSeederMobileVivoY29 {
	return &SpecificationSeederMobileVivoY29{BaseSeeder: BaseSeeder{name: "Specifications for Vivo Y29"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileVivoY29) getBanglaTranslations() map[string]string {
	translations := map[string]string{
		"6 GB":                      "6 GB",
		"6000 mAh":                  "6000 mAh",
		"October 2024":              "অক্টোবর 2024",
		"6.68 inches":               "6.68 ইঞ্চি",
		"Helio G85":                 "হেলিও G85",
		"Mediatek Helio G85 (12nm)": "মিডিয়াটেক হেলিও G85",
		"Octa-core":                 "অক্টা-কোর",
		"Glass front, plastic back, plastic frame": "গ্লাস, প্লাস্টিক",
		"199 g":                      "199 গ্রাম",
		"44W wired":                  "44W ওয়্যার্ড",
		"Android 14, Funtouch 14":    "অ্যান্ড্রয়েড 14",
		"Mali-G52 MC2":               "মালি-G52 MC2",
		"128 GB":                     "128 GB",
		"720 x 1608 pixels":          "720x1608 পিক্সেল",
		"IP64":                       "IP64",
		"GSM / HSPA / LTE":           "GSM/HSPA/LTE",
		"50 MP + 2 MP":               "50+2 MP",
		"8 MP":                       "8 MP",
		"Deep Sea Blue, Pearl White": "ডিপ ব্লু, পার্ল হোয়াইট",
		"IPS LCD, 90Hz":              "IPS LCD, 90Hz",
		"90Hz":                       "90Hz",
		"Available":                  "উপলব্ধ",
		"164.4 x 75.6 x 8.3 mm":      "164.4x75.6x8.3 মিমি",
		"Fingerprint (side-mounted), accelerometer, proximity": "ফিঙ্গারপ্রিন্ট (সাইড), অ্যাক্সেলেরোমিটার, প্রক্সিমিটি",
		"GPS, GLONASS, GALILEO, BDS":                           "GPS, GLONASS, GALILEO, BDS",
		"5.1":                                                  "5.1",
		"Wi-Fi 802.11 a/b/g/n/ac, dual-band":                   "ওয়াই-ফাই 802.11 a/b/g/n/ac, ডুয়াল-ব্যান্ড",
		"USB Type-C 2.0":                                       "USB Type-C 2.0",
		"Dual SIM (Nano-SIM, dual stand-by)":                   "ডুয়াল SIM (ন্যানো-SIM, ডুয়াল স্ট্যান্ড-বাই)",
		"Yes":                                                  "হ্যাঁ",
		"24-bit/192kHz audio":                                  "24-বিট/192kHz অডিও",
		"Li-Po":                                                "Li-Po",
		"44W":                                                  "44W",
		"microSDXC (dedicated slot)":                           "মাইক্রোSDXC (ডেডিকেটেড স্লট)",
		"2.0 GHz":                                              "2.0 GHz",
		"LED flash, HDR, panorama":                             "LED ফ্ল্যাশ, HDR, প্যানোরামা",
		"50 MP, f/1.8, (wide), PDAF\n2 MP, f/2.4, (depth)": "50 MP, f/1.8, (ওয়াইড), PDAF\n2 MP, f/2.4, (ডেপ্থ)",
		"1080p@30fps":                       "1080p@30fps",
		"No":                                "না",
		"720p@30fps":                        "720p@30fps",
		"0.99 W/kg (head) 0.99 W/kg (body)": "0.99 W/kg (হেড) 0.99 W/kg (বডি)",
		"0.34 W/kg (head) 1.59 W/kg (body)": "0.34 W/kg (হেড) 1.59 W/kg (বডি)",
		"V2338 (6GB RAM)":                   "V2338 (6GB RAM)",
		"Face unlock":                       "ফেস আনলক",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range translations {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'vivo-y29'
func (s *SpecificationSeederMobileVivoY29) Seed(db *gorm.DB) error {
	productSlug := "vivo-y29"
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
		"ram":                           "6 GB",
		"battery":                       "6000 mAh",
		"announcement_date":             "October 2024",
		"display_size":                  "6.68 inches",
		"processor":                     "Helio G85",
		"chipset":                       "Mediatek Helio G85 (12nm)",
		"cpu_type":                      "Octa-core",
		"build_material":                "Glass front, plastic back, plastic frame",
		"weight":                        "199 g",
		"fast_charging":                 "44W wired",
		"operating_system":              "Android 14, Funtouch 14",
		"gpu_type":                      "Mali-G52 MC2",
		"storage":                       "128 GB",
		"resolution":                    "720 x 1608 pixels",
		"water_resistance":              "IP64",
		"network_technology":            "GSM / HSPA / LTE",
		"rear_camera":                   "50 MP + 2 MP",
		"front_camera":                  "8 MP",
		"available_colors":              "Deep Sea Blue, Pearl White",
		"display_type":                  "IPS LCD, 90Hz",
		"refresh_rate":                  "90Hz",
		"device_status":                 "Available",
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
		"charging_speed":                "44W",
		"card_slot_type":                "microSDXC (dedicated slot)",
		"internal_memory_capacity":      "128 GB",
		"processor_speed":               "2.0 GHz",
		"camera_features":               "LED flash, HDR, panorama",
		"quad_camera_setup":             "50 MP, f/1.8, (wide), PDAF\n2 MP, f/2.4, (depth)",
		"camera_video_resolution":       "1080p@30fps",
		"optical_zoom":                  "No",
		"front_camera_video_resolution": "720p@30fps",
		"nfc_support":                   "No",
		"sar_rating":                    "0.99 W/kg (head) 0.99 W/kg (body)",
		"sar_rating_eu":                 "0.34 W/kg (head) 1.59 W/kg (body)",
		"model_variants":                "V2338 (6GB RAM)",
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
