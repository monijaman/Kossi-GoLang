package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoCAMON40 seeds specifications/options for product 'tecno-camon-40'
type SpecificationSeederMobileTecnoCAMON40 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoCAMON40 creates a new seeder instance
func NewSpecificationSeederMobileTecnoCAMON40() *SpecificationSeederMobileTecnoCAMON40 {
	return &SpecificationSeederMobileTecnoCAMON40{BaseSeeder: BaseSeeder{name: "Specifications for Tecno CAMON 40"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoCAMON40) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Mediatek Helio G99 (6 nm)":     "মিডিয়াটেক হেলিও জি৯৯ (৬ এনএম)",
		"Helio G99":                     "হেলিও জি৯৯",
		"Mali-G57 MC2":                  "মালি-জি৫৭ এমসি২",
		"Glass front, plastic back":     "গ্লাস ফ্রন্ট, প্লাস্টিক ব্যাক",
		"1080 x 2436 pixels":            "১০৮০ x ২৪৩৬ পিক্সেল",
		"64 MP + 2 MP":                  "৬৪ এমপি + ২ এমপি",
		"32 MP":                         "৩২ এমপি",
		"186 g":                         "১৮৬ গ্রাম",
		"Corning Gorilla Glass":         "কর্নিং গরিলা গ্লাস",
		"Black, Blue":                   "কালো, নীল",
		"6.67 inches":                   "৬.৬৭ ইঞ্চি",
		"IP53":                          "আইপি৫৩",
		"Octa-core":                     "অক্টা-কোর",
		"256 GB":                        "২৫৬ জিবি",
		"120Hz":                         "১২০ হার্টজ",
		"4G":                            "৪জি",
		"5,000 mAh":                     "৫,০০০ এমএএইচ",
		"February 2024":                 "ফেব্রুয়ারি ২০২৪",
		"Available":                     "অ্যাভেইলেবল",
		"8 GB":                          "৮ জিবি",
		"AMOLED, 120Hz":                 "অ্যামোলেড, ১২০ হার্টজ",
		"163.4 x 75.9 x 7.8 mm":         "১৬৩.৪ x ৭৫.৯ x ৭.৮ মিমি",
		"Android 14, HIOS 14":           "অ্যান্ড্রয়েড ১৪, হাইওএস ১৪",
	}
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}
	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-camon-40'
func (s *SpecificationSeederMobileTecnoCAMON40) Seed(db *gorm.DB) error {
	productSlug := "tecno-camon-40"
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
		"screen_protection": "Corning Gorilla Glass",
		"rear_camera":       "64 MP + 2 MP",
		"front_camera":      "32 MP",
		"available_colors":  "Black, Blue",
		"display_size":      "6.67 inches",
		"chipset":           "Mediatek Helio G99 (6 nm)",
		"gpu_type":          "Mali-G57 MC2",
		"resolution":        "1080 x 2436 pixels",
		"build_material":    "Glass front, plastic back",
		"weight":            "186 g",
		"water_resistance":  "IP53",
		"processor":         "Helio G99",
		"cpu_type":          "Octa-core",
		"storage":           "256 GB",
		"refresh_rate":      "120Hz",
		"network_technology": "4G",
		"battery":           "5,000 mAh",
		"announcement_date": "February 2024",
		"device_status":     "Available",
		"ram":               "8 GB",
		"display_type":      "AMOLED, 120Hz",
		"dimensions":        "163.4 x 75.9 x 7.8 mm",
		"operating_system":  "Android 14, HIOS 14",
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
