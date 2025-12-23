package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederMobileXiaomi15Ultra seeds specifications/options for product 'xiaomi-15-ultra'
type SpecificationSeederMobileXiaomi15Ultra struct {
	BaseSeeder
}

// NewSpecificationSeederMobileXiaomi15Ultra creates a new seeder instance
func NewSpecificationSeederMobileXiaomi15Ultra() *SpecificationSeederMobileXiaomi15Ultra {
	return &SpecificationSeederMobileXiaomi15Ultra{BaseSeeder: BaseSeeder{name: "Specifications for Xiaomi 15 Ultra"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMobileXiaomi15Ultra) getBanglaTranslations() map[string]string {
	specific := map[string]string{
		"90W wired, 80W wireless, 10W reverse wireless":                                               "৯০ডব্লিউ ওয়্যার্ড, ৮০ডব্লিউ ওয়্যারলেস, ১০ডব্লিউ রিভার্স ওয়্যারলেস",
		"GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a+B2b), GALILEO (E1+E5a+E5b), QZSS (L1+L5), NavIC": "জিপিএস (এল১+এল৫), গ্লোনাস (জি১), বিডিএস (বি১আই+বি১সি+বি২এ+বি২বি), গ্যালিলিও (ই১+ই৫এ+ই৫বি), কিউজেডএসএস (এল১+এল৫), ন্যাভআইসি",
		"Black, Titanium, White":                        "ব্ল্যাক, টাইটানিয়াম, হোয়াইট",
		"Qualcomm SM8750 Snapdragon 8 Elite (3 nm)":     "কোয়ালকম এসএম৮৭৫০ স্ন্যাপড্রাগন ৮ এলিট (৩ ন্যানোমিটার)",
		"GSM 850 / 900 / 1800 / 1900 MHz":               "জিএসএম ৮৫০ / ৯০০ / ১৮০০ / ১৯০০ মেগাহার্টজ",
		"HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100 MHz": "এইচএসডিপিএ ৮৫০ / ৯০০ / ১৭০০(এডব্লিউএস) / ১৯০০ / ২১০০ মেগাহার্টজ",
		"Yes": "হ্যাঁ",
		"8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision": "৮কে@২৪এফপিএস, ৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০/১২০/২৪০এফপিএস, জাইরো-ইআইএস, এইচডিআর১০+, ডলবি ভিশন",
		"Octa-core (2x4.32 GHz Oryon V2 + 6x3.53 GHz Oryon V2)":                         "অক্টা-কোর (২x৪.৩২ গিগাহার্টজ ওরিয়ন ভি২ + ৬x৩.৫৩ গিগাহার্টজ ওরিয়ন ভি২)",
		"16 GB":                                 "১৬ জিবি",
		"UFS 4.0":                               "ইউএফএস ৪.০",
		"No microSD card slot":                  "কোন মাইক্রোএসডি কার্ড স্লট নেই",
		"USB Type-C 3.2, DisplayPort 1.4, OTG":  "ইউএসবি টাইপ-সি ৩.২, ডিসপ্লেপোর্ট ১.৪, ওটিজি",
		"3.2x + 5x Optical Zoom":                "৩.২x + ৫x অপটিকাল জুম",
		"1440 x 3200 pixels (~522 ppi density)": "১৪৪০ x ৩২০০ পিক্সেল (~৫২২ পিপিআই ঘনত্ব)",
		"IP68 dust tight and water resistant (up to 1.5m for 30 min)":                                              "আইপি৬৮ ধুলো প্রতিরোধী এবং জল প্রতিরোধী (৩০ মিনিটের জন্য ১.৫মি পর্যন্ত)",
		"n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78":                                                    "এন১, এন৩, এন৫, এন৭, এন৮, এন২০, এন২৮, এন৩৮, এন৪০, এন৪১, এন৭৭, এন৭৮",
		"90W wired (100% in 32 min), 80W wireless (100% in 38 min), 10W reverse wireless":                          "৯০ডব্লিউ ওয়্যার্ড (৩২ মিনিটে ১০০%), ৮০ডব্লিউ ওয়্যারলেস (৩৮ মিনিটে ১০০%), ১০ডব্লিউ রিভার্স ওয়্যারলেস",
		"Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum": "ফিঙ্গারপ্রিন্ট (আন্ডার ডিসপ্লে, অপটিকাল), অ্যাক্সেলেরোমিটার, জাইরো, প্রক্সিমিটি, কম্পাস, ব্যারোমিটার, কালার স্পেকট্রাম",
		"6.8 inches":                           "৬.৮ ইঞ্চি",
		"Snapdragon 8 Elite":                   "স্ন্যাপড্রাগন ৮ এলিট",
		"Glass front and back, Aluminum frame": "গ্লাস ফ্রন্ট এবং ব্যাক, অ্যালুমিনিয়াম ফ্রেম",
		"50MP (wide, f/1.6) + 50MP (periscope telephoto 5x, f/2.5) + 50MP (telephoto 3.2x, f/2.0) + 50MP (ultrawide, f/1.8)": "৫০এমপি (ওয়াইড, এফ/১.৬) + ৫০এমপি (পেরিস্কোপ টেলিফটো ৫x, এফ/২.৫) + ৫০এমপি (টেলিফটো ৩.২x, এফ/২.০) + ৫০এমপি (আল্ট্রাওয়াইড, এফ/১.৮)",
		"Stereo speakers, Dolby Atmos": "স্টেরিও স্পিকার, ডলবি অ্যাটমস",
		"Adreno 830":                   "অ্যাড্রেনো ৮৩০",
		"LTPO AMOLED, 1B colors, 120Hz, Dolby Vision, HDR10+, 3200 nits (peak)": "এলটিপিও এএমওএলইডি, ১বি কালার, ১২০হার্টজ, ডলবি ভিশন, এইচডিআর১০+, ৩২০০ নিটস (পিক)",
		"Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct":                  "ওয়াই-ফাই ৮০২.১১ এ/বি/জি/এন/এসি/৬ই/৭, ট্রাই-ব্যান্ড, ওয়াই-ফাই ডাইরেক্ট",
		"4K@30/60fps, 1080p@30/60fps":                                           "৪কে@৩০/৬০এফপিএস, ১০৮০পি@৩০/৬০এফপিএস",
		"Android 15, HyperOS":                                                   "অ্যান্ড্রয়েড ১৫, হাইপারওএস",
		"Dual Nano-SIM":                                                         "ডুয়াল ন্যানো-সিম",
		"24116PN76C":                                                            "২৪১১৬পিএন৭৬সি",
		"February 2025":                                                         "ফেব্রুয়ারি ২০২৫",
		"162.8 x 74.5 x 8.3 mm (6.41 x 2.93 x 0.33 in)":                         "১৬২.৮ x ৭৪.৫ x ৮.৩ মিমি (৬.৪১ x ২.৯৩ x ০.৩৩ ইঞ্চি)",
		"GSM / CDMA / HSPA / LTE / 5G":                                          "জিএসএম / সিডিএমএ / এইচএসপিএ / এলটিই / ৫জি",
		"Li-Po (non-removable)":                                                 "লি-পো (অপসারণযোগ্য নয়)",
		"Xiaomi HyperOS, AI image processing, Leica camera mode":                "শাওমি হাইপারওএস, এআই ইমেজ প্রসেসিং, লাইকা ক্যামেরা মোড",
		"24-bit/192kHz audio, Hi-Res Audio certification":                       "২৪-বিট/১৯২কেএইচজেড অডিও, হাই-রেস অডিও সার্টিফিকেশন",
		"No":                                "না",
		"Available. Released February 2025": "উপলব্ধ। ফেব্রুয়ারি ২০২৫-এ রিলিজ",
		"2x4.32 GHz + 6x3.53 GHz":           "২x৪.৩২ গিগাহার্টজ + ৬x৩.৫৩ গিগাহার্টজ",
		"Xiaomi Dragon Crystal Glass":       "শাওমি ড্রাগন ক্রিস্টাল গ্লাস",
		"120Hz":                             "১২০হার্টজ",
		"220 g (7.76 oz)":                   "২২০ গ্রাম (৭.৭৬ আউন্স)",
		"50 MP + 50 MP + 50 MP + 50 MP":     "৫০ এমপি + ৫০ এমপি + ৫০ এমপি + ৫০ এমপি",
		"5,300 mAh":                         "৫,৩০০ এমএএইচ",
		"Yes - 80W wireless":                "হ্যাঁ - ৮০ডব্লিউ ওয়্যারলেস",
		"512 GB / 1 TB":                     "৫১২ জিবি / ১ টিবি",
		"LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 26, 28, 32, 38, 40, 41, 42, 66": "এলটিই ব্যান্ড ১, ২, ৩, ৪, ৫, ৭, ৮, ১২, ১৩, ১৭, ১৮, ১৯, ২০, ২৬, ২৮, ৩২, ৩৮, ৪০, ৪১, ৪২, ৬৬",
		"5.4, A2DP, LE, aptX HD":                                "৫.৪, এ২ডিপি, এলই, অ্যাপটিএক্স এইচডি",
		"Leica optics, Dual-LED dual-tone flash, HDR, panorama": "লাইকা অপটিক্স, ডুয়াল-এলইডি ডুয়াল-টোন ফ্ল্যাশ, এইচডিআর, প্যানোরামা",
		"32 MP, f/2.0, 25mm (wide), HDR":                        "৩২ এমপি, এফ/২.০, ২৫মিমি (ওয়াইড), এইচডিআর",
	}
	common := s.GetCommonBanglaTranslations()
	for k, v := range specific {
		common[k] = v
	}
	return common
}

// Seed inserts specification records for the product identified by slug 'xiaomi-15-ultra'
func (s *SpecificationSeederMobileXiaomi15Ultra) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-15-ultra"
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
		"fast_charging":                 "90W wired, 80W wireless, 10W reverse wireless",
		"positioning_system":            "GPS (L1+L5), GLONASS (G1), BDS (B1I+B1c+B2a+B2b), GALILEO (E1+E5a+E5b), QZSS (L1+L5), NavIC",
		"available_colors":              "Black, Titanium, White",
		"chipset":                       "Qualcomm SM8750 Snapdragon 8 Elite (3 nm)",
		"2g_bands":                      "GSM 850 / 900 / 1800 / 1900 MHz",
		"3g_bands":                      "HSDPA 850 / 900 / 1700(AWS) / 1900 / 2100 MHz",
		"nfc_support":                   "Yes",
		"camera_video_resolution":       "8K@24fps, 4K@30/60fps, 1080p@30/60/120/240fps, gyro-EIS, HDR10+, Dolby Vision",
		"5g_support":                    "Yes",
		"cpu_type":                      "Octa-core (2x4.32 GHz Oryon V2 + 6x3.53 GHz Oryon V2)",
		"ram":                           "16 GB",
		"internal_memory_capacity":      "UFS 4.0",
		"card_slot_type":                "No microSD card slot",
		"usb_type":                      "USB Type-C 3.2, DisplayPort 1.4, OTG",
		"optical_zoom":                  "3.2x + 5x Optical Zoom",
		"resolution":                    "1440 x 3200 pixels (~522 ppi density)",
		"water_resistance":              "IP68 dust tight and water resistant (up to 1.5m for 30 min)",
		"5g_bands":                      "n1, n3, n5, n7, n8, n20, n28, n38, n40, n41, n77, n78",
		"charging_speed":                "90W wired (100% in 32 min), 80W wireless (100% in 38 min), 10W reverse wireless",
		"sensors":                       "Fingerprint (under display, optical), accelerometer, gyro, proximity, compass, barometer, color spectrum",
		"display_size":                  "6.8 inches",
		"processor":                     "Snapdragon 8 Elite",
		"build_material":                "Glass front and back, Aluminum frame",
		"quad_camera_setup":             "50MP (wide, f/1.6) + 50MP (periscope telephoto 5x, f/2.5) + 50MP (telephoto 3.2x, f/2.0) + 50MP (ultrawide, f/1.8)",
		"loudspeaker_quality":           "Stereo speakers, Dolby Atmos",
		"gpu_type":                      "Adreno 830",
		"display_type":                  "LTPO AMOLED, 1B colors, 120Hz, Dolby Vision, HDR10+, 3200 nits (peak)",
		"wifi_support":                  "Wi-Fi 802.11 a/b/g/n/ac/6e/7, tri-band, Wi-Fi Direct",
		"front_camera_video_resolution": "4K@30/60fps, 1080p@30/60fps",
		"operating_system":              "Android 15, HyperOS",
		"sim_card_type":                 "Dual Nano-SIM",
		"model_variants":                "24116PN76C",
		"announcement_date":             "February 2025",
		"dimensions":                    "162.8 x 74.5 x 8.3 mm (6.41 x 2.93 x 0.33 in)",
		"network_technology":            "GSM / CDMA / HSPA / LTE / 5G",
		"battery_type":                  "Li-Po (non-removable)",
		"special_features":              "Xiaomi HyperOS, AI image processing, Leica camera mode",
		"audio_quality":                 "24-bit/192kHz audio, Hi-Res Audio certification",
		"audio_jack":                    "No",
		"device_status":                 "Available. Released February 2025",
		"processor_speed":               "2x4.32 GHz + 6x3.53 GHz",
		"screen_protection":             "Xiaomi Dragon Crystal Glass",
		"refresh_rate":                  "120Hz",
		"weight":                        "220 g (7.76 oz)",
		"rear_camera":                   "50 MP + 50 MP + 50 MP + 50 MP",
		"battery":                       "5,300 mAh",
		"wireless_charging":             "Yes - 80W wireless",
		"storage":                       "512 GB / 1 TB",
		"4g_bands":                      "LTE Band 1, 2, 3, 4, 5, 7, 8, 12, 13, 17, 18, 19, 20, 26, 28, 32, 38, 40, 41, 42, 66",
		"bluetooth_version":             "5.4, A2DP, LE, aptX HD",
		"camera_features":               "Leica optics, Dual-LED dual-tone flash, HDR, panorama",
		"front_camera":                  "32 MP, f/2.0, 25mm (wide), HDR",
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
