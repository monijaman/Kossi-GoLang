package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoCAMON40Pro5G seeds specifications/options for product 'tecno-camon-40-pro-5g'
type SpecificationSeederMobileTecnoCAMON40Pro5G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCAMON40Pro5G creates a new seeder instance
func NewSpecificationSeederMobileTecnoCAMON40Pro5G() *SpecificationSeederMobileTecnoCAMON40Pro5G {
	return &SpecificationSeederMobileTecnoCAMON40Pro5G{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 40 Pro 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCAMON40Pro5G) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Corning Gorilla Glass 5":                 "কর্নিং গরিলা গ্লাস ৫",
		"Mediatek Dimensity 8200 Ultimate (4 nm)": "মিডিয়াটেক ডাইমেনসিটি ৮২০০ আলটিমেট (৪ এনএম)",
		"Dimensity 8200 Ultimate":                 "ডাইমেনসিটি ৮২০০ আলটিমেট",
		"Mali-G610 MC6":                           "মালি-জি৬১০ এমসি৬",
		"Glass front, glass back":                 "গ্লাস ফ্রন্ট, গ্লাস ব্যাক",
		"Iceland Basalt, Alps Snowy Silver":       "আইসল্যান্ড ব্যাসল্ট, আল্পস স্নোয়ি সিলভার",
		"IP54":                                    "আইপি৫৪",
		"50 MP + 50 MP + 2 MP":                    "৫০ এমপি + ৫০ এমপি + ২ এমপি",
		"50 MP":                                   "৫০ এমপি",
		"6.8 inches":                              "৬.৮ ইঞ্চি",
		"12 GB":                                   "১২ জিবি",
		"200 g":                                   "২০০ গ্রাম",
		"164 x 74.5 x 7.7 mm":                     "১৬৪ × ৭৪.৫ × ৭.৭ মিমি",
		"February 2024":                           "ফেব্রুয়ারি ২০২৪",
		"Android 14, HIOS 14":                     "অ্যান্ড্রয়েড ১৪, হাইওএস ১৪",
		"AMOLED, 120Hz":                           "অ্যামোলেড, ১২০ হার্জ",
		"1080 x 2436 pixels":                      "১০৮০ × ২৪৩৬ পিক্সেল",
		"5,000 mAh":                               "৫,০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"512 GB":                                  "৫১২ জিবি",
		"Octa-core":                               "অক্টা-কোর",
		"Available":                               "উপলব্ধ",
		"5G":                                      "৫জি",
		// memory & storage
		"8 GB":                "৮ জিবি",
		"MicroSD, up to 1 TB": "মাইক্রোএসডি, সর্বোচ্চ ১ টেরাবাইট",
		// charging & connectivity
		"65W":        "৬৫ ওয়াট",
		"No":         "না",
		"Yes":        "হ্যাঁ",
		"USB Type-C": "ইউএসবি টাইপ-সি",
		"5.3":        "৫.৩",
		// sensors & positioning
		"GPS, A-GPS, GLONASS": "জিপিএস, এ-জিপিএস, গ্লোনাস",
		"Fingerprint (side-mounted), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (সাইড-মাউন্টেড), অ্যাক্সেলারোমিটার, প্রোক্সিমিটি সেন্সর, কম্পাস",
		// camera/video
		"4K@30fps": "৪কে @ ৩০ ফ্রেম/সেকেন্ড",
		"AI":       "এআই",
	}
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}
	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-camon-40-pro-5g'
func (s *SpecificationSeederMobileTecnoCAMON40Pro5G) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-40-pro-5g"
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
		"device_status":      "Available",
		"dimensions":         "164 x 74.5 x 7.7 mm",
		"gpu_type":           "Mali-G610 MC6",
		"screen_protection":  "Corning Gorilla Glass 5",
		"announcement_date":  "February 2024",
		"chipset":            "Mediatek Dimensity 8200 Ultimate (4 nm)",
		"cpu_type":           "Octa-core",
		"storage":            "512 GB",
		"network_technology": "5G",
		"battery":            "5,000 mAh",
		"display_size":       "6.8 inches",
		"ram":                "12 GB",
		"resolution":         "1080 x 2436 pixels",
		"refresh_rate":       "120Hz",
		"build_material":     "Glass front, glass back",
		"weight":             "200 g",
		"water_resistance":   "IP54",
		"rear_camera":        "50 MP + 50 MP + 2 MP",
		"processor":          "Dimensity 8200 Ultimate",
		"display_type":       "AMOLED, 120Hz",
		"front_camera":       "50 MP",
		"operating_system":   "Android 14, HIOS 14",
		"available_colors":   "Iceland Basalt, Alps Snowy Silver",
		// additional
		"internal_memory_capacity": "512 GB",
		"card_slot_type":           "MicroSD, up to 1 TB",
		"fast_charging":            "65W",
		"charging_speed":           "65W",
		"wireless_charging":        "No",
		"audio_jack":               "No",
		"sensors":                  "Fingerprint (side-mounted), Accelerometer, Proximity, Compass",
		"positioning_system":       "GPS, A-GPS, GLONASS",
		"usb_type":                 "USB Type-C",
		"bluetooth_version":        "5.3",
		"nfc_support":              "No",
		"camera_video_resolution":  "4K@30fps",
		"camera_features":          "AI",
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
