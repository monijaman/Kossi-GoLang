package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS23 seeds specifications/options for product 'samsung-galaxy-s23'
type SpecificationSeederMobileSamsungGalaxyS23 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23() *SpecificationSeederMobileSamsungGalaxyS23 {
	return &SpecificationSeederMobileSamsungGalaxyS23{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Snapdragon 8 Gen 2": "স্ন্যাপড্রাগন ৮ জেন ২",
		"Adreno 740":         "অ্যাড্রেনো ৭৪০",
		"Octa-core":          "অক্টা-কোর",
		"Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)": "কোয়ালকম এসএম৮৫৫০-এসি স্ন্যাপড্রাগন ৮ জেন ২ (৪ ন্যানোমিটার)",
		"Dynamic AMOLED 2X, 120Hz, HDR10+":             "ডায়নামিক অ্যামোলেড ২এক্স, ১২০হার্টজ, এইচডিআর১০+",
		"1080 x 2340 pixels":                           "১০৮০ × ২৩৪০ পিক্সেল",
		"6.1 inches":                                   "৬.১ ইঞ্চি",
		"50 MP + 10 MP + 12 MP":                        "৫০ এমপি + ১০ এমপি + ১২ এমপি",
		"12 MP":                                        "১২ এমপি",
		"Phantom Black, Green, Cream, Lavender":        "ফ্যান্টম ব্ল্যাক, গ্রিন, ক্রিম, ল্যাভেন্ডার",
		"146.3 x 70.9 x 7.6 mm":                        "১৪৬.৩ × ৭০.৯ × ৭.৬ মিমি",
		"168 g":                                        "১৬৮ গ্রাম",
		"3,900 mAh":                                    "৩,৯০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Android 13, upgradable":                       "অ্যান্ড্রয়েড ১৩, আপগ্রেডযোগ্য",
		"5G":                                           "৫জি",
		"128 GB / 256 GB / 512 GB":                     "১২৮ জিবি / ২৫৬ জিবি / ৫১২ জিবি",
		"8 GB":                                         "৮ জিবি",
		"120Hz":                                        "১২০হার্টজ",
		"Aluminum frame, glass front/back":             "অ্যালুমিনিয়াম ফ্রেম, কাচের সামনের/পিছন",
		"Corning Gorilla Glass Victus 2":               "কর্নিং গরিলা গ্লাস ভিক্টাস ২",
		"IP68":                                         "আইপি৬৮",
		"February 2023":                                "ফেব্রুয়ারি ২০২৩",
		"GSM 850 / 900 / 1800 / 1900":                  "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100":                                                                       "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE bands 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 38, 39, 40, 41, 66":                       "এলটিই ব্যান্ড ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৫, ২৬, ২৮, ৩২, ৩৮, ৩৯, ৪০, ৪১, ৬৬",
		"SA/NSA/Sub6/mmWave - 5G bands 2, 3, 5, 7, 8, 12, 13, 14, 20, 25, 26, 28, 29, 30, 38, 40, 41, 48, 66, 71, 77, 78": "এসএ/এনএসএ/সাব৬/এমএমওয়েভ - ৫জি ব্যান্ড ২, ৩, ৫, ৭, ৮, ১২, ১৩, ১৪, ২০, ২৫, ২৬, ২৮, ২৯, ৩০, ৩৮, ৪০, ৪১, ৪৮, ৬৬, ৭১, ৭৭, ৭৮",
		"25W":            "২৫ওয়াট",
		"Up to 3.36 GHz": "৩.৩৬ গিগাহার্টজ পর্যন্ত",
		"128GB 8GB RAM, 256GB 8GB RAM, 512GB 8GB RAM": "১২৮জিবি ৮জিবি র্যাম, ২৫৬জিবি ৮জিবি র্যাম, ৫১২জিবি ৮জিবি র্যাম",
		"No":        "না",
		"Available": "উপলব্ধ",
		"50 MP, f/1.8, 24mm (wide), 1/1.56\", 1.0µm, Dual Pixel PDAF, OIS\n10 MP, f/2.4, 70mm (telephoto), 1/3.94\", 1.0µm, Dual Pixel PDAF, OIS, 3x optical zoom\n12 MP, f/2.2, 13mm, 120˚ (ultrawide), 1/2.55\", 1.4µm, Dual Pixel PDAF, Super Steady video": "৫০ এমপি, এফ/১.৮, ২৪মিমি (ওয়াইড), ১/১.৫৬\", ১.০মাইক্রোমিটার, ডুয়াল পিক্সেল পিডিএএফ, ওআইএস\n১০ এমপি, এফ/২.৪, ৭০মিমি (টেলিফোটো), ১/৩.৯৪\", ১.০মাইক্রোমিটার, ডুয়াল পিক্সেল পিডিএএফ, ওআইএস, ৩এক্স অপটিক্যাল জুম\n১২ এমপি, এফ/২.২, ১৩মিমি, ১২০ডিগ্রি (আল্ট্রাওয়াইড), ১/২.৫৫\", ১.৪মাইক্রোমিটার, ডুয়াল পিক্সেল পিডিএএফ, সুপার স্টেডি ভিডিও",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s23'
func (s *SpecificationSeederMobileSamsungGalaxyS23) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23"
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
		"front_camera":             "12 MP",
		"processor":                "Snapdragon 8 Gen 2",
		"build_material":           "Aluminum frame, glass front/back",
		"dimensions":               "146.3 x 70.9 x 7.6 mm",
		"operating_system":         "Android 13, upgradable",
		"available_colors":         "Phantom Black, Green, Cream, Lavender",
		"announcement_date":        "February 2023",
		"device_status":            "Available",
		"chipset":                  "Qualcomm SM8550-AC Snapdragon 8 Gen 2 (4 nm)",
		"cpu_type":                 "Octa-core",
		"ram":                      "8 GB",
		"storage":                  "128 GB / 256 GB / 512 GB",
		"battery":                  "3,900 mAh",
		"gpu_type":                 "Adreno 740",
		"screen_protection":        "Corning Gorilla Glass Victus 2",
		"refresh_rate":             "120Hz",
		"weight":                   "168 g",
		"water_resistance":         "IP68",
		"display_size":             "6.1 inches",
		"display_type":             "Dynamic AMOLED 2X, 120Hz, HDR10+",
		"resolution":               "1080 x 2340 pixels",
		"network_technology":       "5G",
		"rear_camera":              "50 MP + 10 MP + 12 MP",
		"2g_bands":                 "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                 "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                 "LTE bands 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 38, 39, 40, 41, 66",
		"5g_bands":                 "SA/NSA/Sub6/mmWave - 5G bands 2, 3, 5, 7, 8, 12, 13, 14, 20, 25, 26, 28, 29, 30, 38, 40, 41, 48, 66, 71, 77, 78",
		"charging_speed":           "25W",
		"processor_speed":          "Up to 3.36 GHz",
		"internal_memory_capacity": "128GB 8GB RAM, 256GB 8GB RAM, 512GB 8GB RAM",
		"card_slot_type":           "No",
		"quad_camera_setup":        "50 MP, f/1.8, 24mm (wide), 1/1.56\", 1.0µm, Dual Pixel PDAF, OIS\n10 MP, f/2.4, 70mm (telephoto), 1/3.94\", 1.0µm, Dual Pixel PDAF, OIS, 3x optical zoom\n12 MP, f/2.2, 13mm, 120˚ (ultrawide), 1/2.55\", 1.4µm, Dual Pixel PDAF, Super Steady video",
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
