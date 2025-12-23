package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileTecnoPHANTOMVFold25G seeds specifications/options for product 'tecno-phantom-v-fold2-5g'
type SpecificationSeederMobileTecnoPHANTOMVFold25G struct {
	BaseSeeder
}

// NewSpecificationSeederMobileTecnoPHANTOMVFold25G creates a new seeder instance
func NewSpecificationSeederMobileTecnoPHANTOMVFold25G() *SpecificationSeederMobileTecnoPHANTOMVFold25G {
	return &SpecificationSeederMobileTecnoPHANTOMVFold25G{BaseSeeder: BaseSeeder{name: "Specifications for Tecno PHANTOM V Fold2 5G"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileTecnoPHANTOMVFold25G) getBanglaTranslations() map[string]string {
	commonTranslations := s.GetCommonBanglaTranslations()
	tecnoTranslations := map[string]string{
		"Foldable LTPO AMOLED, 120Hz":             "ফোল্ডেবল এলটিপিও অ্যামোলেড, ১২০ হার্জ",
		"7.85 inches (Folded: 6.42 inches)":       "৭.৮৫ ইঞ্চি (ফোল্ডেড: ৬.৪২ ইঞ্চি)",
		"Unfolded: 159.4 x 140.4 x 6.9 mm":        "আনফোল্ডেড: ১৫৯.৪ × ১৪০.৪ × ৬.৯ মিমি",
		"Glass front, glass back, aluminum frame": "গ্লাস ফ্রন্ট, গ্লাস ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"Mediatek Dimensity 9000+ (4 nm)":         "মিডিয়াটেক ডাইমেনসিটি ৯০০০+ (৪ এনএম)",
		"Dimensity 9000+":                         "ডাইমেনসিটি ৯০০০+",
		"Mali-G710 MC10":                          "মালি-জি৭১০ এমসি১০",
		"50 MP + 50 MP + 13 MP":                   "৫০ এমপি + ৫০ এমপি + ১৩ এমপি",
		"16 MP (Cover) + 32 MP (Inner)":           "১৬ এমপি (কভার) + ৩২ এমপি (ইনার)",
		"2000 x 2296 pixels":                      "২০০০ × ২২৯৬ পিক্সেল",
		"299 g":                                   "২৯৯ গ্রাম",
		"September 2024":                          "সেপ্টেম্বর ২০২৪",
		"Android 14, HIOS 14 Fold":                "অ্যান্ড্রয়েড ১৪, হাইওএস ১৪ ফোল্ড",
		"Corning Gorilla Glass Victus (Cover)":    "কর্নিং গরিলা গ্লাস ভিক্টাস (কভার)",
		"5G":                                      "৫জি",
		"Black, White":                            "কালো, সাদা",
		"Available":                               "উপলব্ধ",
		"512 GB":                                  "৫১২ জিবি",
		"120Hz":                                   "১২০ হার্জ",
		"5,000 mAh":                               "৫,০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Octa-core":                               "অক্টা-কোর",
		"12 GB":                                   "১২ জিবি",
		"No":                                      "না",
		// storage & memory
		"MicroSD, up to 1 TB": "মাইক্রোএসডি, সর্বোচ্চ ১ টেরাবাইট",
		// charging & connectivity
		"65W":                 "৬৫ ওয়াট",
		"USB Type-C":          "ইউএসবি টাইপ-সি",
		"5.3":                 "৫.৩",
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

// Seed inserts specification records for the product identified by slug 'tecno-phantom-v-fold2-5g'
func (s *SpecificationSeederMobileTecnoPHANTOMVFold25G) Seed(db *gorm.DB) error {
	productSlug := "tecno-phantom-v-fold2-5g"
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
		"screen_protection":        "Corning Gorilla Glass Victus (Cover)",
		"dimensions":               "Unfolded: 159.4 x 140.4 x 6.9 mm",
		"network_technology":       "5G",
		"available_colors":         "Black, White",
		"device_status":            "Available",
		"processor":                "Dimensity 9000+",
		"chipset":                  "Mediatek Dimensity 9000+ (4 nm)",
		"gpu_type":                 "Mali-G710 MC10",
		"storage":                  "512 GB",
		"internal_memory_capacity": "512 GB",
		"card_slot_type":           "MicroSD, up to 1 TB",
		"refresh_rate":             "120Hz",
		"build_material":           "Glass front, glass back, aluminum frame",
		"weight":                   "299 g",
		"rear_camera":              "50 MP + 50 MP + 13 MP",
		"display_size":             "7.85 inches (Folded: 6.42 inches)",
		"display_type":             "Foldable LTPO AMOLED, 120Hz",
		"battery":                  "5,000 mAh",
		"fast_charging":            "65W",
		"charging_speed":           "65W",
		"wireless_charging":        "No",
		"audio_jack":               "No",
		"announcement_date":        "September 2024",
		"cpu_type":                 "Octa-core",
		"ram":                      "12 GB",
		"resolution":               "2000 x 2296 pixels",
		"water_resistance":         "No",
		"front_camera":             "16 MP (Cover) + 32 MP (Inner)",
		"camera_video_resolution":  "4K@30fps",
		"camera_features":          "AI",
		"operating_system":         "Android 14, HIOS 14 Fold",
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
