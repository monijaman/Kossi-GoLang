package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileSamsungGalaxyA07 seeds specifications/options for product 'samsung-galaxy-a07'
type SpecificationSeederMobileSamsungGalaxyA07 struct {
	BaseSeeder
}

// NewSpecificationSeederMobileSamsungGalaxyA07 creates a new seeder instance
func NewSpecificationSeederMobileSamsungGalaxyA07() *SpecificationSeederMobileSamsungGalaxyA07 {
	return &SpecificationSeederMobileSamsungGalaxyA07{BaseSeeder: BaseSeeder{name: "Specifications for Samsung Galaxy A07"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileSamsungGalaxyA07) getBanglaTranslations() map[string]string {
	common := s.GetCommonBanglaTranslations()
	specific := map[string]string{
		"Android 14, One UI Core 6":            "অ্যান্ড্রয়েড ১৪, ওয়ান ইউআই কোর ৬",
		"Samsung Exynos 850 (8 nm)":            "স্যামসাং এক্সিনস ৮৫০ (৮ ন্যানোমিটার)",
		"720 x 1600 pixels (~270 ppi density)": "৭২০ × ১৬০০ পিক্সেল (~২৭০ পিপিআই ঘনত্ব)",
		"5 MP":                                 "৫ মেগাপিক্সেল",
		"June 2025":                            "জুন ২০২৫",
		"6.5 inches":                           "৬.৫ ইঞ্চি",
		"4 GB / 6 GB":                          "৪ জিবি / ৬ জিবি",
		"64 GB / 128 GB":                       "৬৪ জিবি / ১২৮ জিবি",
		"60Hz":                                 "৬০ হার্জ",
		"Plastic frame, plastic back":          "প্লাস্টিক ফ্রেম, প্লাস্টিক ব্যাক",
		"186 g (6.56 oz)":                      "১৮৬ গ্রাম (৬.৫৬ আউন্স)",
		"164.2 x 75.9 x 9.1 mm":                "১৬৪.২ × ৭৫.৯ × ৯.১ মিলিমিটার",
		"5000 mAh":                             "৫০০০ মিলিঅ্যাম্পিয়ার আওয়ার",
		"Li-Ion (non-removable)":               "লি-আয়ন (অপসারণযোগ্য নয়)",
		"15W wired":                            "১৫ ওয়াট তারযুক্ত",
		"Exynos 850":                           "এক্সিনস ৮৫০",
		"Octa-core (4x2.0 GHz Cortex-A55 & 4x2.0 GHz Cortex-A55)": "অক্টা-কোর (৪×২.০ গিগাহার্জ কোর্টেক্স-এ৫৫ এবং ৪×২.০ গিগাহার্জ কোর্টেক্স-এ৫৫)",
		"Mali-G52":              "মালি জি৫২",
		"PLS LCD":               "পিএলএস এলসিডি",
		"Black, Green, White":   "কালো, সবুজ, সাদা",
		"Upcoming":              "আসছে",
		"No":                    "না",
		"GSM / HSPA / LTE":      "জিএসএম / এইচএসপিএ / এলটিই",
		"13 MP + 2 MP":          "১৩ মেগাপিক্সেল + ২ মেগাপিক্সেল",
		"LED flash, HDR":        "এলইডি ফ্ল্যাশ, এইচডিআর",
		"1080p @ 30fps":         "১০৮০পি @ ৩০ এফপিএস",
		"None":                  "কোনটি নয়",
		"Wi-Fi 802.11 b/g/n":    "ওয়াই-ফাই ৮০২.১১ বি/জি/এন",
		"5.1":                   "৫.১",
		"USB-C":                 "ইউএসবি-সি",
		"3.5 mm":                "৩.৫ মিলিমিটার",
		"Mono":                  "মোনো",
		"GPS, GLONASS, GALILEO": "জিপিএস, গ্লোনাস, গ্যালিলিও",
		"Fingerprint (rear), Accelerometer, Proximity": "ফিঙ্গারপ্রিন্ট (পিছনে), অ্যাকসেলারোমিটার, প্রক্সিমিটি",
		"Nano-SIM + Nano-SIM":                          "ন্যানো-সিম + ন্যানো-সিম",
		"microSD":                                      "মাইক্রোএসডি",
		"4/64GB, 6/128GB":                              "৪/৬৪ জিবি, ৬/১২৮ জিবি",
		"Face unlock":                                  "ফেস আনলক",
	}
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'samsung-galaxy-a07'
func (s *SpecificationSeederMobileSamsungGalaxyA07) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-a07"
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
		"operating_system":              "Android 14, One UI Core 6",
		"chipset":                       "Samsung Exynos 850 (8 nm)",
		"resolution":                    "720 x 1600 pixels (~270 ppi density)",
		"front_camera":                  "5 MP",
		"announcement_date":             "June 2025",
		"display_size":                  "6.5 inches",
		"ram":                           "4 GB / 6 GB",
		"storage":                       "64 GB / 128 GB",
		"refresh_rate":                  "60Hz",
		"build_material":                "Plastic frame, plastic back",
		"weight":                        "186 g (6.56 oz)",
		"dimensions":                    "164.2 x 75.9 x 9.1 mm",
		"battery":                       "5000 mAh",
		"battery_type":                  "Li-Ion (non-removable)",
		"charging_speed":                "15W wired",
		"processor":                     "Exynos 850",
		"cpu_type":                      "Octa-core (4x2.0 GHz Cortex-A55 & 4x2.0 GHz Cortex-A55)",
		"gpu_type":                      "Mali-G52",
		"display_type":                  "PLS LCD",
		"available_colors":              "Black, Green, White",
		"device_status":                 "Upcoming",
		"water_resistance":              "No",
		"network_technology":            "GSM / HSPA / LTE",
		"rear_camera":                   "13 MP + 2 MP",
		"camera_features":               "LED flash, HDR",
		"camera_video_resolution":       "1080p @ 30fps",
		"front_camera_video_resolution": "1080p @ 30fps",
		"optical_zoom":                  "None",
		"fast_charging":                 "15W wired",
		"wifi_support":                  "Wi-Fi 802.11 b/g/n",
		"bluetooth_version":             "5.1",
		"nfc_support":                   "No",
		"usb_type":                      "USB-C",
		"audio_jack":                    "3.5 mm",
		"loudspeaker_quality":           "Mono",
		"positioning_system":            "GPS, GLONASS, GALILEO",
		"sensors":                       "Fingerprint (rear), Accelerometer, Proximity",
		"sim_card_type":                 "Nano-SIM + Nano-SIM",
		"card_slot_type":                "microSD",
		"model_variants":                "4/64GB, 6/128GB",
		"special_features":              "Face unlock",
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
