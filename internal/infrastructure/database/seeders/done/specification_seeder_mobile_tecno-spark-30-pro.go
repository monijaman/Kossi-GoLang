package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoSPARK30Pro seeds specifications/options for product 'tecno-spark-30-pro'
type SpecificationSeederMobileTecnoSPARK30Pro struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoSPARK30Pro creates a new seeder instance
func NewSpecificationSeederMobileTecnoSPARK30Pro() *SpecificationSeederMobileTecnoSPARK30Pro {
	return &SpecificationSeederMobileTecnoSPARK30Pro{BaseSeeder: BaseSeeder{name: "Specifications for Tecno SPARK 30 Pro"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoSPARK30Pro) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Mediatek Helio G99 (6 nm)": "মিডিয়াটেক হেলিও জি৯৯ (৬ এনএম)",
		"Helio G99":                 "হেলিও জি৯৯",
		"108 MP + AI":               "১০৮ এমপি + এআই",
		"Black, White":              "কালো, সাদা",
		"September 2024":            "সেপ্টেম্বর ২০২৪",
		"Mali-G57 MC2":              "মালি-জি৫৭ এমসি২",
		"32 MP":                     "৩২ এমপি",
		"5,000 mAh":                 "৫,০০০ এমএএইচ",
		"6.78 inches":               "৬.৭৮ ইঞ্চি",
		"1080 x 2460 pixels":        "১০৮০ x ২৪৬০ পিক্সেল",
		"120Hz":                     "১২০ হার্টজ",
		"AMOLED, 120Hz":             "অ্যামোলেড, ১২০ হার্টজ",
		"Android 14":                "অ্যান্ড্রয়েড ১৪",
		"Available":                 "অ্যাভেইলেবল",
		"4G":                        "৪জি",
		"8 GB":                      "৮ জিবি",
		"256 GB":                    "২৫৬ জিবি",
		"Octa-core":                 "অক্টা-কোর",
		"Glass front":               "গ্লাস ফ্রন্ট",
		"Plastic body":              "প্লাস্টিক বডি",
		"IP53":                      "আইপি৫৩",
		"190 g":                     "১৯০ গ্রাম",
		"168 x 76 x 8.5 mm":         "১৬৮ x ৭৬ x ৮.৫ মিমি",
	}
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}
	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-spark-30-pro'
func (s *SpecificationSeederMobileTecnoSPARK30Pro) Seed(db *gorm.DB) error {
	productSlug := "tecno-spark-30-pro"
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
		"chipset":            "Mediatek Helio G99 (6 nm)",
		"ram":                "8 GB",
		"refresh_rate":       "120Hz",
		"weight":             "190 g",
		"dimensions":         "168 x 76 x 8.5 mm",
		"network_technology": "4G",
		"rear_camera":        "108 MP + AI",
		"front_camera":       "32 MP",
		"display_size":       "6.78 inches",
		"processor":          "Helio G99",
		"storage":            "256 GB",
		"battery":            "5,000 mAh",
		"operating_system":   "Android 14",
		"available_colors":   "Black, White",
		"announcement_date":  "September 2024",
		"device_status":      "Available",
		"screen_protection":  "Glass front",
		"cpu_type":           "Octa-core",
		"gpu_type":           "Mali-G57 MC2",
		"display_type":       "AMOLED, 120Hz",
		"resolution":         "1080 x 2460 pixels",
		"build_material":     "Plastic body",
		"water_resistance":   "IP53",
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
