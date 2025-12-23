package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyS23FE seeds specifications/options for product 'samsung-galaxy-s23-fe'
type SpecificationSeederMobileSamsungGalaxyS23FE struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyS23FE creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyS23FE() *SpecificationSeederMobileSamsungGalaxyS23FE {
	return &SpecificationSeederMobileSamsungGalaxyS23FE{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy S23 FE"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyS23FE) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Exynos 2200 / Snapdragon 8 Gen 1": "এক্সিনস ২২০০ / স্ন্যাপড্রাগন ৮ জেন ১",
		"Xclipse 920":                      "এক্সক্লিপস ৯২০",
		"Octa-core":                        "অক্টা-কোর",
		"Exynos 2200 (4 nm)":               "এক্সিনস ২২০০ (৪ ন্যানোমিটার)",
		"Dynamic AMOLED 2X, 120Hz, HDR10+": "ডায়নামিক অ্যামোলেড ২এক্স, ১২০হার্টজ, এইচডিআর১০+",
		"1080 x 2340 pixels":               "১০৮০ × ২৩৪০ পিক্সেল",
		"6.4 inches":                       "৬.৪ ইঞ্চি",
		"50 MP + 8 MP + 12 MP":             "৫০ এমপি + ৮ এমপি + ১২ এমপি",
		"10 MP":                            "১০ এমপি",
		"Mint, Cream, Graphite, Purple, Indigo, Tangerine": "মিন্ট, ক্রিম, গ্রাফাইট, পার্পল, ইন্ডিগো, ট্যাঞ্জারিন",
		"158 x 76.5 x 8.2 mm":                              "১৫৮ × ৭৬.৫ × ৮.২ মিমি",
		"209 g":                                            "২০৯ গ্রাম",
		"4,500 mAh":                                        "৪,৫০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Android 13, upgradable":                           "অ্যান্ড্রয়েড ১৩, আপগ্রেডযোগ্য",
		"5G":                                               "৫জি",
		"128 GB / 256 GB":                                  "১২৮ জিবি / ২৫৬ জিবি",
		"8 GB":                                             "৮ জিবি",
		"120Hz":                                            "১২০হার্টজ",
		"Aluminum frame, glass front/back":                 "অ্যালুমিনিয়াম ফ্রেম, কাচের সামনের/পিছন",
		"Corning Gorilla Glass 5":                          "কর্নিং গরিলা গ্লাস ৫",
		"IP68":                                             "আইপি৬৮",
		"October 2023":                                     "অক্টোবর ২০২৩",
		"GSM 850 / 900 / 1800 / 1900":                      "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100":                                                                       "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০",
		"LTE bands 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 38, 39, 40, 41, 66":                       "এলটিই ব্যান্ড ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৫, ২৬, ২৮, ৩২, ৩৮, ৩৯, ৪০, ৪১, ৬৬",
		"SA/NSA/Sub6/mmWave - 5G bands 2, 3, 5, 7, 8, 12, 13, 14, 20, 25, 26, 28, 29, 30, 38, 40, 41, 48, 66, 71, 77, 78": "এসএ/এনএসএ/সাব৬/এমএমওয়েভ - ৫জি ব্যান্ড ২, ৩, ৫, ৭, ৮, ১২, ১৩, ১৪, ২০, ২৫, ২৬, ২৮, ২৯, ৩০, ৩৮, ৪০, ৪১, ৪৮, ৬৬, ৭১, ৭৭, ৭৮",
		"25W":                          "২৫ওয়াট",
		"Up to 2.8 GHz":                "২.৮ গিগাহার্টজ পর্যন্ত",
		"128GB 8GB RAM, 256GB 8GB RAM": "১২৮জিবি ৮জিবি র্যাম, ২৫৬জিবি ৮জিবি র্যাম",
		"No":                           "না",
		"Available":                    "উপলব্ধ",
		"50 MP, f/1.8, 24mm (wide), 1/1.56\", 1.0µm, Dual Pixel PDAF, OIS\n8 MP, f/2.4, 76mm (telephoto), 1/4.5\", 1.0µm, PDAF, OIS, 3x optical zoom\n12 MP, f/2.2, 13mm, 123˚ (ultrawide), 1/2.55\", 1.4µm, Dual Pixel PDAF, Super Steady video": "৫০ এমপি, এফ/১.৮, ২৪মিমি (ওয়াইড), ১/১.৫৬\", ১.০মাইক্রোমিটার, ডুয়াল পিক্সেল পিডিএএফ, ওআইএস\n৮ এমপি, এফ/২.৪, ৭৬মিমি (টেলিফোটো), ১/৪.৫\", ১.০মাইক্রোমিটার, পিডিএএফ, ওআইএস, ৩এক্স অপটিক্যাল জুম\n১২ এমপি, এফ/২.২, ১৩মিমি, ১২৩ডিগ্রি (আল্ট্রাওয়াইড), ১/২.৫৫\", ১.৪মাইক্রোমিটার, ডুয়াল পিক্সেল পিডিএএফ, সুপার স্টেডি ভিডিও",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-s23-fe'
func (s *SpecificationSeederMobileSamsungGalaxyS23FE) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-s23-fe"
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
		"dimensions":               "158 x 76.5 x 8.2 mm",
		"water_resistance":         "IP68",
		"rear_camera":              "50 MP + 8 MP + 12 MP",
		"operating_system":         "Android 13, upgradable",
		"ram":                      "8 GB",
		"resolution":               "1080 x 2340 pixels",
		"announcement_date":        "October 2023",
		"display_size":             "6.4 inches",
		"cpu_type":                 "Octa-core",
		"gpu_type":                 "Xclipse 920",
		"display_type":             "Dynamic AMOLED 2X, 120Hz, HDR10+",
		"weight":                   "209 g",
		"front_camera":             "10 MP",
		"device_status":            "Available",
		"processor":                "Exynos 2200 / Snapdragon 8 Gen 1",
		"build_material":           "Aluminum frame, glass front/back",
		"network_technology":       "5G",
		"battery":                  "4,500 mAh",
		"available_colors":         "Mint, Cream, Graphite, Purple, Indigo, Tangerine",
		"chipset":                  "Exynos 2200 (4 nm)",
		"storage":                  "128 GB / 256 GB",
		"screen_protection":        "Corning Gorilla Glass 5",
		"refresh_rate":             "120Hz",
		"2g_bands":                 "GSM 850 / 900 / 1800 / 1900",
		"3g_bands":                 "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100",
		"4g_bands":                 "LTE bands 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 25, 26, 28, 32, 38, 39, 40, 41, 66",
		"5g_bands":                 "SA/NSA/Sub6/mmWave - 5G bands 2, 3, 5, 7, 8, 12, 13, 14, 20, 25, 26, 28, 29, 30, 38, 40, 41, 48, 66, 71, 77, 78",
		"charging_speed":           "25W",
		"processor_speed":          "Up to 2.8 GHz",
		"internal_memory_capacity": "128GB 8GB RAM, 256GB 8GB RAM",
		"card_slot_type":           "No",
		"quad_camera_setup":        "50 MP, f/1.8, 24mm (wide), 1/1.56\", 1.0µm, Dual Pixel PDAF, OIS\n8 MP, f/2.4, 76mm (telephoto), 1/4.5\", 1.0µm, PDAF, OIS, 3x optical zoom\n12 MP, f/2.2, 13mm, 123˚ (ultrawide), 1/2.55\", 1.4µm, Dual Pixel PDAF, Super Steady video",
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
