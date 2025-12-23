package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoCAMON30 seeds specifications/options for product 'tecno-camon-30'
type SpecificationSeederMobileTecnoCAMON30 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCAMON30 creates a new seeder instance
func NewSpecificationSeederMobileTecnoCAMON30() *SpecificationSeederMobileTecnoCAMON30 {
	return &SpecificationSeederMobileTecnoCAMON30{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 30"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCAMON30) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()

	// Add Tecno CAMON 30 specific translations
	tecnoTranslations := map[string]string{
		"8 GB / 12 GB":                       "৮ জিবি / ১২ জিবি",
		"256 GB":                             "২৫৬ জিবি",
		"120Hz":                              "১২০ হার্জ",
		"189 g":                              "১৮৯ গ্রাম",
		"IP53":                               "IP৫৩",
		"4G":                                 "৪জি",
		"50 MP + 2 MP + AI":                  "৫০ মেগাপিক্সেল + ২ মেগাপিক্সেল + এআই",
		"50 MP":                              "৫০ মেগাপিক্সেল",
		"6.78 inches":                        "৬.৭৮ ইঞ্চি",
		"Helio G99 Ultimate":                 "হেলিও G৯৯ আলটিমেট",
		"Mali-G57 MC2":                       "মালি-G৫৭ MC২",
		"1080 x 2436 pixels":                 "১০৮০ x ২৪৩৬ পিক্সেল",
		"Glass front":                        "গ্লাস ফ্রন্ট",
		"Glass front, glass back":            "গ্লাস ফ্রন্ট, গ্লাস ব্যাক",
		"165.3 x 75.3 x 7.7 mm":              "১৬৫.৩ x ৭৫.৩ x ৭.৭ মিলিমিটার",
		"5,000 mAh":                          "৫,০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Octa-core":                          "অক্টা-কোর",
		"AMOLED, 120Hz":                      "অ্যামোলেড, ১২০ হার্জ",
		"Iceland Basalt, Uyuni Salt White":   "আইসল্যান্ড ব্যাসাল্ট, উয়ুনি সল্ট হোয়াইট",
		"February 2024":                      "ফেব্রুয়ারি ২০২৪",
		"Mediatek Helio G99 Ultimate":        "মিডিয়াটেক হেলিও G৯৯ আলটিমেট",
		"Android 14, HIOS 14":                "অ্যান্ড্রয়েড ১৪, HIOS ১৪",
		"Available":                          "উপলব্ধ",
		"5.3, A2DP, LE":                      "৫.৩, A2DP, LE",
		"Yes":                                "হ্যাঁ",
		"USB Type-C 2.0, OTG":                "USB টাইপ-সি ২.০, OTG",
		"Dual SIM (Nano-SIM, dual stand-by)": "ডুয়াল সিম (ন্যানো-সিম, ডুয়াল স্ট্যান্ড-বাই)",
		"Fingerprint (under display, optical), accelerometer, gyro, compass": "ফিঙ্গারপ্রিন্ট (ডিসপ্লের নিচে, অপটিক্যাল), অ্যাক্সেলেরোমিটার, জাইরো, কম্পাস",
		"GPS, GNSS":                 "GPS, GNSS",
		"Yes, with stereo speakers": "হ্যাঁ, স্টেরিও স্পিকার সহ",
	}

	// Merge common translations with Tecno-specific ones
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}

	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-camon-30'
func (s *SpecificationSeederMobileTecnoCAMON30) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-30"
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
		"ram":                 "8 GB / 12 GB",
		"storage":             "256 GB",
		"refresh_rate":        "120Hz",
		"weight":              "189 g",
		"water_resistance":    "IP53",
		"network_technology":  "4G",
		"rear_camera":         "50 MP + 2 MP + AI",
		"front_camera":        "50 MP",
		"display_size":        "6.78 inches",
		"processor":           "Helio G99 Ultimate",
		"gpu_type":            "Mali-G57 MC2",
		"resolution":          "1080 x 2436 pixels",
		"screen_protection":   "Glass front",
		"build_material":      "Glass front, glass back",
		"dimensions":          "165.3 x 75.3 x 7.7 mm",
		"battery":             "5,000 mAh",
		"cpu_type":            "Octa-core",
		"display_type":        "AMOLED, 120Hz",
		"available_colors":    "Iceland Basalt, Uyuni Salt White",
		"announcement_date":   "February 2024",
		"chipset":             "Mediatek Helio G99 Ultimate",
		"operating_system":    "Android 14, HIOS 14",
		"device_status":       "Available",
		"bluetooth_version":   "5.3, A2DP, LE",
		"nfc_support":         "Yes",
		"usb_type":            "USB Type-C 2.0, OTG",
		"sim_card_type":       "Dual SIM (Nano-SIM, dual stand-by)",
		"sensors":             "Fingerprint (under display, optical), accelerometer, gyro, compass",
		"positioning_system":  "GPS, GNSS",
		"loudspeaker_quality": "Yes, with stereo speakers",
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
