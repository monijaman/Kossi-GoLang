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

	// Target form generator ID (primary key)
	const formGeneratorID uint = 5
	const mobileCategoryID uint = 79

	// Map of form generator keys to existing specification_key IDs from database
	// These are ONLY existing keys from the database - NO new keys will be created
	keyMapping := map[string]uint{
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
		"rear_camera":             122, // Rear Camera
		"camera_features":         16,  // Camera Features
		"quad_camera_setup":       58,  // Quad Camera Setup
		"camera_video_resolution": 40,  // Main Camera Video Resolution
		"optical_zoom":            193, // Optical Zoom
		// NOTE: "front_camera" (ID 31) DELETED from database
		"front_camera_video_resolution": 32, // Front Camera Video Resolution

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

	// Collect all spec key IDs in order
	keyNames := []string{
		"display_size", "display_type", "resolution", "refresh_rate", "screen_protection",
		"processor", "chipset", "cpu_type", "gpu_type", "processor_speed",
		"ram", "storage", "internal_memory_capacity", "card_slot_type",
		"rear_camera", "camera_features", "quad_camera_setup", "camera_video_resolution", "optical_zoom",
		"front_camera_video_resolution",
		"battery", "battery_type", "fast_charging", "charging_speed", "wireless_charging",
		"operating_system",
		"audio_quality", "loudspeaker_quality", "audio_jack",
		"build_material", "weight", "dimensions", "water_resistance",
		"network_technology", "2g_bands", "3g_bands", "4g_bands", "5g_bands",
		"wifi_support", "bluetooth_version", "nfc_support", "usb_type",
		"positioning_system", "sensors", "special_features",
		"sim_card_type", "sar_rating", "sar_rating_eu",
		"available_colors", "model_variants", "announcement_date", "device_status",
	}

	var keyIDs []uint
	missingKeys := []string{}

	for _, keyName := range keyNames {
		if keyID, exists := keyMapping[keyName]; exists {
			keyIDs = append(keyIDs, keyID)
		} else {
			missingKeys = append(missingKeys, keyName)
		}
	}

	if len(missingKeys) > 0 {
		log.Printf("⚠️  WARNING: The following specification keys are NOT in the database: %v", missingKeys)
		log.Printf("⚠️  These keys will be SKIPPED. Update the form to remove or provide valid IDs.", missingKeys)
	}

	specJSON, err := json.Marshal(keyIDs)
	if err != nil {
		return err
	}

	// Update form generator ID 5 with new specification keys (if exists, update; if not, create)
	result := db.Model(&models.FormGeneratorModel{}).
		Where("id = ?", formGeneratorID).
		Updates(map[string]interface{}{
			"specification_id": string(specJSON),
			"category_id":      mobileCategoryID,
			"status":           "active",
		})

	if result.Error != nil {
		log.Printf("❌ Error updating form generator ID %d: %v", formGeneratorID, result.Error)
		return result.Error
	}

	// If no rows were updated, insert a new one
	if result.RowsAffected == 0 {
		fg := &models.FormGeneratorModel{
			ID:              formGeneratorID,
			CategoryID:      mobileCategoryID,
			SpecificationID: string(specJSON),
			Status:          "active",
		}

		if err := db.Create(fg).Error; err != nil {
			log.Printf("❌ Error creating mobile form generator: %v", err)
			return err
		}
	}

	actualCount := len(keyIDs)
	fmt.Printf("✅ Successfully updated form generator ID %d (category %d) with %d spec keys\n", formGeneratorID, mobileCategoryID, actualCount)
	if len(missingKeys) > 0 {
		fmt.Printf("⚠️  Skipped %d keys that don't exist in database: %v\n", len(missingKeys), missingKeys)
	}
	return nil
}
