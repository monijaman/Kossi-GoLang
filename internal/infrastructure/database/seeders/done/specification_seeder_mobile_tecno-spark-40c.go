package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSPARK40C seeds specifications/options for product 'tecno-spark-40c'
type SpecificationSeederMobileTecnoSPARK40C struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSPARK40C creates a new seeder instance
func NewSpecificationSeederMobileTecnoSPARK40C() *SpecificationSeederMobileTecnoSPARK40C {
	return &SpecificationSeederMobileTecnoSPARK40C{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 40C"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSPARK40C) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Unisoc T606":                  "ইউনিসক টি৬০৬",
		"Unisoc T606 (12 nm)":          "ইউনিসক টি৬০৬ (১২ এনএম)",
		"Mali-G57 MP1":                 "মালি-জি৫৭ এমপি১",
		"Gravity Black, Mystery White": "গ্র্যাভিটি ব্ল্যাক, মিস্ট্রি হোয়াইট",
		"November 2023":                "নভেম্বর ২০২৩",
		"185 g":                        "১৮৫ গ্রাম",
		"50 MP + AI":                   "৫০ এমপি + এআই",
		"163.7 x 75.6 x 8.6 mm":        "১৬৩.৭ × ৭৫.৬ × ৮.৬ মিমি",
		"4 GB / 8 GB":                  "৪ জিবি / ৮ জিবি",
		"IPS LCD, 90Hz":                "আইপিএস এলসিডি, ৯০ হার্জ",
		"720 x 1612 pixels":            "৭২০ × ১৬১২ পিক্সেল",
		"8 MP":                         "৮ এমপি",
		"Plastic body":                 "প্লাস্টিক বডি",
		"Android 13":                   "অ্যান্ড্রয়েড ১৩",
		"Glass front":                  "গ্লাস ফ্রন্ট",
		"No":                           "না",
		"6.6 inches":                   "৬.৬ ইঞ্চি",
		"90Hz":                         "৯০ হার্জ",
		// storage & memory
		"4 GB":                "৪ জিবি",
		"8 GB":                "৮ জিবি",
		"128 GB":              "১২৮ জিবি",
		"MicroSD, up to 1 TB": "মাইক্রোএসডি, সর্বোচ্চ ১ টেরাবাইট",
		// charging & battery
		"5,000 mAh": "৫,০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"65W":       "৬৫ ওয়াট",

		// connectivity & sensors
		"USB Type-C":          "ইউএসবি টাইপ-সি",
		"GPS, A-GPS, GLONASS": "জিপিএস, এ-জিপিএস, গ্লোনাস",
		"Fingerprint (side-mounted), Accelerometer, Proximity, Compass": "ফিঙ্গারপ্রিন্ট (সাইড-মাউন্টেড), অ্যাক্সেলারোমিটার, প্রোক্সিমিটি সেন্সর, কম্পাস",
		"5.3": "৫.৩",
		// camera/video/features
		"4K@30fps": "৪কে @ ৩০ ফ্রেম/সেকেন্ড",
		"AI":       "এআই",
	}
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}
	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-40c'
func (s *SpecificationSeederMobileTecnoSPARK40C) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-40c"
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
		"weight":                   "185 g",
		"rear_camera":              "50 MP + AI",
		"front_camera":             "8 MP",
		"announcement_date":        "November 2023",
		"processor":                "Unisoc T606",
		"available_colors":         "Gravity Black, Mystery White",
		"dimensions":               "163.7 x 75.6 x 8.6 mm",
		"cpu_type":                 "Octa-core",
		"ram":                      "4 GB / 8 GB",
		"internal_memory_capacity": "128 GB",
		"card_slot_type":           "MicroSD, up to 1 TB",
		"display_type":             "IPS LCD, 90Hz",
		"refresh_rate":             "90Hz",
		"build_material":           "Plastic body",
		"battery":                  "5,000 mAh",
		"fast_charging":            "65W",
		"charging_speed":           "65W",
		"wireless_charging":        "No",
		"operating_system":         "Android 13",
		"chipset":                  "Unisoc T606 (12 nm)",
		"gpu_type":                 "Mali-G57 MP1",
		"resolution":               "720 x 1612 pixels",
		"screen_protection":        "Glass front",
		"water_resistance":         "No",
		"network_technology":       "4G",
		"device_status":            "Available",
		"display_size":             "6.6 inches",
		"storage":                  "128 GB",
		"usb_type":                 "USB Type-C",
		"positioning_system":       "GPS, A-GPS, GLONASS",
		"sensors":                  "Fingerprint (side-mounted), Accelerometer, Proximity, Compass",
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
