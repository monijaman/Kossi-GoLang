package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SpecificationKey struct {
	ID  int
	Key string
}

type SpecificationKeyTranslation struct {
	SpecificationKeyID int
	Language           string
	Translation        string
}

func runTranslateKeys() {
	// Database connection setup
	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Keys to translate
	specificationKeys := []struct {
		Key         string
		Translation string
	}{
		{"display_size", "ডিসপ্লে আকার"},
		{"processor", "প্রসেসর"},
		{"ram", "র‍্যাম"},
		{"storage", "স্টোরেজ"},
		{"display_type", "ডিসপ্লে টাইপ"},
		{"resolution", "রেজোলিউশন"},
		{"screen_protection", "স্ক্রিন প্রোটেকশন"},
		{"refresh_rate", "রিফ্রেশ রেট"},
		{"build_material", "বিল্ড মেটেরিয়াল"},
		{"weight", "ওজন"},
		{"dimensions", "মাত্রা"},
		{"water_resistance", "জল প্রতিরোধ"},
		{"network_technology", "নেটওয়ার্ক প্রযুক্তি"},
		{"wifi_support", "ওয়াইফাই সাপোর্ট"},
		{"bluetooth_version", "ব্লুটুথ সংস্করণ"},
		{"nfc_support", "এনএফসি সাপোর্ট"},
		{"usb_type", "ইউএসবি টাইপ"},
		{"rear_camera", "পিছনের ক্যামেরা"},
		{"camera_features", "ক্যামেরার বৈশিষ্ট্য"},
		{"camera_video_resolution", "ক্যামেরা ভিডিও রেজোলিউশন"},
		{"optical_zoom", "অপটিক্যাল জুম"},
		{"front_camera", "সামনের ক্যামেরা"},
		{"front_camera_video_resolution", "সামনের ক্যামেরা ভিডিও রেজোলিউশন"},
		{"operating_system", "অপারেটিং সিস্টেম"},
		{"battery", "ব্যাটারি"},
		{"battery_type", "ব্যাটারি টাইপ"},
		{"fast_charging", "দ্রুত চার্জিং"},
		{"wireless_charging", "ওয়্যারলেস চার্জিং"},
		{"5g_support", "৫জি সাপোর্ট"},
		{"positioning_system", "পজিশনিং সিস্টেম"},
		{"sensors", "সেন্সর"},
		{"special_features", "বিশেষ বৈশিষ্ট্য"},
		{"sim_card_type", "সিম কার্ড টাইপ"},
		{"loudspeaker_quality", "লাউডস্পিকার গুণমান"},
		{"audio_quality", "অডিও গুণমান"},
		{"audio_jack", "অডিও জ্যাক"},
		{"available_colors", "উপলব্ধ রং"},
		{"announcement_date", "ঘোষণার তারিখ"},
		{"device_status", "ডিভাইসের অবস্থা"},
	}

	for _, key := range specificationKeys {
		var existingKey SpecificationKey
		// Check if the key exists
		result := db.Where("key = ?", key.Key).First(&existingKey)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				log.Printf("Specification key not found: %s", key.Key)
				continue
			} else {
				log.Printf("Error querying specification key: %v", result.Error)
				continue
			}
		}

		// Insert the translation
		translation := SpecificationKeyTranslation{
			SpecificationKeyID: existingKey.ID,
			Language:           "bn",
			Translation:        key.Translation,
		}
		if err := db.Create(&translation).Error; err != nil {
			log.Printf("Failed to insert translation: %v", err)
		} else {
			fmt.Printf("Inserted translation for key: %s\n", key.Key)
		}
	}
}
