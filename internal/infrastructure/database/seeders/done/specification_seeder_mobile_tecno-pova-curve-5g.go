package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoPOVACurve5G seeds specifications/options for product 'tecno-pova-curve-5g'
type SpecificationSeederMobileTecnoPOVACurve5G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPOVACurve5G creates a new seeder instance
func NewSpecificationSeederMobileTecnoPOVACurve5G() *SpecificationSeederMobileTecnoPOVACurve5G {
	return &SpecificationSeederMobileTecnoPOVACurve5G{BaseSeeder: BaseSeeder{name: "Specifications for Tecno POVA Curve 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPOVACurve5G) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Dimensity 7050":                 "ডাইমেনসিটি ৭০৫০",
		"Mediatek Dimensity 7050 (6 nm)": "মিডিয়াটেক ডাইমেনসিটি ৭০৫০ (৬ এনএম)",
		"Mali-G68 MC4":                   "মালি-জি৬৮ এমসি৪",
		"AMOLED, 120Hz, Curved":          "অ্যামোলেড, ১২০ হার্টজ, কার্ভড",
		"50 MP + 2 MP":                   "৫০ এমপি + ২ এমপি",
		"16 MP":                          "১৬ এমপি",
		"Black, Blue":                    "কালো, নীল",
		"162 x 75 x 7.9 mm":              "১৬২ x ৭৫ x ৭.৯ মিমি",
		"185 g":                          "১৮৫ গ্রাম",
		"Glass front":                    "গ্লাস ফ্রন্ট",
		"Plastic body":                   "প্লাস্টিক বডি",
		"1080 x 2400 pixels":             "১০৮০ x ২৪০০ পিক্সেল",
		"6.67 inches":                    "৬.৬৭ ইঞ্চি",
		"5,000 mAh":                      "৫,০০০ এমএএইচ",
		"256 GB":                         "২৫৬ জিবি",
		"8 GB":                           "৮ জিবি",
		"120Hz":                          "১২০ হার্টজ",
		"No":                             "না",
		"Available":                      "অ্যাভেইলেবল",
		"5G":                             "৫জি",
		"2024":                           "২০২৪",
		"Octa-core":                      "অক্টা-কোর",
		"Android 14":                     "অ্যান্ড্রয়েড ১৪",
	}
	for key, value := range tecnoTranslations {
		commonTranslations[key] = value
	}
	return commonTranslations
}

// Seed inserts specification records for the product identified by slug 'tecno-pova-curve-5g'
func (s *SpecificationSeederMobileTecnoPOVACurve5G) Seed(db *gorm.DB) error {
	productSlug := "tecno-pova-curve-5g"
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
		"processor":          "Dimensity 7050",
		"chipset":            "Mediatek Dimensity 7050 (6 nm)",
		"ram":                "8 GB",
		"screen_protection":  "Glass front",
		"rear_camera":        "50 MP + 2 MP",
		"device_status":      "Available",
		"resolution":         "1080 x 2400 pixels",
		"weight":             "185 g",
		"dimensions":         "162 x 75 x 7.9 mm",
		"water_resistance":   "No",
		"cpu_type":           "Octa-core",
		"gpu_type":           "Mali-G68 MC4",
		"storage":            "256 GB",
		"display_type":       "AMOLED, 120Hz, Curved",
		"build_material":     "Plastic body",
		"network_technology": "5G",
		"battery":            "5,000 mAh",
		"available_colors":   "Black, Blue",
		"display_size":       "6.67 inches",
		"refresh_rate":       "120Hz",
		"front_camera":       "16 MP",
		"operating_system":   "Android 14",
		"announcement_date":  "2024",
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
