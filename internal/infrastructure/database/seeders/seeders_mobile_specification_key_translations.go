package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SeedMobileSpecificationKeyTranslations seeds Bengali translations for specification keys
func SeedMobileSpecificationKeyTranslations(db *gorm.DB) error {
	fmt.Println("\n🔄 Seeding Mobile Specification Key Translations (Bangla)...")

	// Get all specification keys
	var specKeys []models.SpecificationKeyModel
	db.Find(&specKeys)

	if len(specKeys) == 0 {
		log.Println("⚠️  No specification keys found for translation")
		return nil
	}

	translations := []models.SpecificationKeyTranslationModel{}

	// Translation map from English specification keys to Bangla
	translationMap := map[string]string{
		// Original 15 keys
		"display_size":      "ডিসপ্লে সাইজ",
		"processor":         "প্রসেসর",
		"ram":               "র্যাম",
		"storage":           "স্টোরেজ",
		"rear_camera":       "রিয়ার ক্যামেরা",
		"front_camera":      "ফ্রন্ট ক্যামেরা",
		"battery":           "ব্যাটারি",
		"fast_charging":     "দ্রুত চার্জিং",
		"operating_system":  "অপারেটিং সিস্টেম",
		"weight":            "ওজন",
		"dimensions":        "মাত্রা",
		"water_resistance":  "জলরোধী রেটিং",
		"5g_support":        "৫জি সাপোর্ট",
		"bluetooth_version": "ব্লুটুথ সংস্করণ",
		"refresh_rate":      "রিফ্রেশ রেট",

		// Processor & Performance
		"chipset":         "চিপসেট",
		"cpu_type":        "সিপিইউ টাইপ",
		"gpu_type":        "জিপিইউ টাইপ",
		"processor_speed": "প্রসেসর গতি",

		// Display Details
		"display_type":      "ডিসপ্লে প্রকার",
		"resolution":        "রেজোলিউশন",
		"screen_protection": "স্ক্রিন সুরক্ষা",

		// Physical Build
		"build_material": "নির্মাণ সামগ্রী",

		// Network & Connectivity
		"network_technology": "নেটওয়ার্ক প্রযুক্তি",
		"2g_bands":           "২জি ব্যান্ড",
		"3g_bands":           "৩জি ব্যান্ড",
		"4g_bands":           "৪জি ব্যান্ড",
		"5g_bands":           "৫জি ব্যান্ড",
		"wifi_support":       "ওয়াই-ফাই সাপোর্ট",
		"nfc_support":        "এনএফসি সাপোর্ট",
		"usb_type":           "ইউএসবি টাইপ",

		// Storage & Memory
		"internal_memory_capacity": "অভ্যন্তরীণ মেমোরি ক্ষমতা",
		"card_slot_type":           "কার্ড স্লট প্রকার",

		// Camera Details
		"camera_features":               "ক্যামেরা বৈশিষ্ট্য",
		"quad_camera_setup":             "চতুর্মুখী ক্যামেরা সেটআপ",
		"camera_video_resolution":       "ক্যামেরা ভিডিও রেজোলিউশন",
		"optical_zoom":                  "অপটিক্যাল জুম",
		"front_camera_video_resolution": "ফ্রন্ট ক্যামেরা ভিডিও রেজোলিউশন",

		// Audio & Sound
		"loudspeaker_quality": "লাউডস্পিকার গুণমান",
		"audio_quality":       "অডিও গুণমান",
		"audio_jack":          "অডিও জ্যাক",

		// Battery & Charging
		"battery_type":      "ব্যাটারি প্রকার",
		"charging_speed":    "চার্জিং গতি",
		"wireless_charging": "ওয়্যারলেস চার্জিং",

		// Positioning & Sensors
		"positioning_system": "অবস্থান নির্ধারণ সিস্টেম",
		"sensors":            "সেন্সর",
		"special_features":   "বিশেষ বৈশিষ্ট্য",

		// SIM & Connectivity
		"sim_card_type": "সিম কার্ড প্রকার",
		"sar_rating":    "এসএআর রেটিং",
		"sar_rating_eu": "এসএআর রেটিং (ইউরোপীয়)",

		// Product Information
		"available_colors":  "উপলব্ধ রং",
		"model_variants":    "মডেল ভেরিয়েন্ট",
		"announcement_date": "ঘোষণার তারিখ",
		"device_status":     "ডিভাইস অবস্থা",
	}

	// Create translations for each specification key
	for _, specKey := range specKeys {
		if bengaliName, exists := translationMap[specKey.SpecificationKey]; exists {
			translation := models.SpecificationKeyTranslationModel{
				SpecificationKeyID: specKey.ID,
				Locale:             "bn",
				SpecificationKey:   bengaliName,
			}
			translations = append(translations, translation)
		}
	}

	if len(translations) == 0 {
		log.Println("⚠️  No translations found to seed")
		return nil
	}

	// Batch create translations
	result := db.CreateInBatches(translations, 50)
	if result.Error != nil {
		log.Printf("❌ Error seeding specification key translations: %v", result.Error)
		return result.Error
	}

	fmt.Printf("✅ Successfully seeded %d specification key translations (Bangla)\n", result.RowsAffected)
	return nil
}
