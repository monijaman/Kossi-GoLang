package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SeedMobileFormGenerator seeds form generator configurations for mobile products
func SeedMobileFormGenerator(db *gorm.DB) error {
	fmt.Println("\n🔄 Seeding Mobile Form Generator...")

	// The mobile category ID used across seeders
	const mobileCategoryID uint = 79

	// Desired mobile specification keys (snake_case) matching keys seeded in SeedMobileSpecifications
	keyNames := []string{
		// Core display
		"display_size",
		"display_type",
		"resolution",
		"refresh_rate",
		"screen_protection",

		// Performance
		"processor",
		"chipset",
		"cpu_type",
		"gpu_type",
		"processor_speed",

		// Memory & storage
		"ram",
		"storage",
		"internal_memory_capacity",
		"card_slot_type",

		// Cameras
		"rear_camera",
		"camera_features",
		"quad_camera_setup",
		"camera_video_resolution",
		"optical_zoom",
		"front_camera",
		"front_camera_video_resolution",

		// Battery & charging
		"battery",
		"battery_type",
		"fast_charging",
		"charging_speed",
		"wireless_charging",

		// Software
		"operating_system",

		// Build & body
		"build_material",
		"weight",
		"dimensions",
		"water_resistance",

		// Network & connectivity
		"network_technology",
		"2g_bands",
		"3g_bands",
		"4g_bands",
		"5g_bands",
		"wifi_support",
		"bluetooth_version",
		"nfc_support",
		"usb_type",

		// Positioning & sensors
		"positioning_system",
		"sensors",
		"special_features",

		// SIM & compliance
		"sim_card_type",
		"sar_rating",
		"sar_rating_eu",

		// Commercial info
		"available_colors",
		"model_variants",
		"announcement_date",
		"device_status",
	}

	// Resolve keys to their numeric IDs
	var keyIDs []uint
	for _, name := range keyNames {
		sk, err := CreateOrFindSpecificationKey(db, name)
		if err != nil {
			return err
		}
		keyIDs = append(keyIDs, sk.ID)
	}

	specJSON, err := json.Marshal(keyIDs)
	if err != nil {
		return err
	}

	// Clean any previous form generator rows for mobiles (avoid duplicates of per-key rows)
	if err := db.Where("category_id = ?", mobileCategoryID).Delete(&models.FormGeneratorModel{}).Error; err != nil {
		log.Printf("❌ Error cleaning existing form generator rows for category %d: %v", mobileCategoryID, err)
		return err
	}

	// Insert a single consolidated row with JSON array of spec key IDs
	fg := &models.FormGeneratorModel{
		CategoryID:      mobileCategoryID,
		SpecificationID: string(specJSON),
		Status:          "active",
	}

	if err := db.Create(fg).Error; err != nil {
		log.Printf("❌ Error creating mobile form generator: %v", err)
		return err
	}

	fmt.Printf("✅ Successfully seeded form generator for category %d with %d spec keys\n", mobileCategoryID, len(keyIDs))
	return nil
}
