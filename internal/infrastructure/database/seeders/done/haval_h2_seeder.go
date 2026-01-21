package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HavalH2Seeder seeds specifications for Haval H2
type HavalH2Seeder struct {
	BaseSeeder
}

// NewHavalH2Seeder creates a new Haval H2 seeder
func NewHavalH2Seeder() *HavalH2Seeder {
	return &HavalH2Seeder{
		BaseSeeder: BaseSeeder{name: "Haval H2 Specifications"},
	}
}

func (hhs *HavalH2Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Turbo I4":          "১.৫ লিটার টার্বো আই৪",
		"1497 cc":                "১৪৯৭ সিসি",
		"4":                      "৪",
		"Inline":                 "ইনলাইন",
		"Petrol":                 "পেট্রোল",
		"150 hp @ 5600 rpm":      "১৫০ হর্স পাওয়ার @ ৫৬০০ আরপিএম",
		"220 Nm @ 2000-4500 rpm": "২২০ এনএম @ ২০০০-৪৫০০ আরপিএম",
		"Direct Injection":       "ডাইরেক্ট ইনজেকশন",
		"75.0 x 84.8 mm":         "৭৫.০ x ৮৪.৮ মিমি",
		"10.5:1":                 "১০.৫:১",
		"Yes":                    "হ্যাঁ",
		"No":                     "না",
		"Manual":                 "ম্যানুয়াল",
		"6-Speed":                "৬-স্পিড",
		"FWD":                    "এফডব্লিউডি",
		"Single Clutch":          "সিঙ্গেল ক্লাচ",
		"175 km/h":               "১৭৫ কিমি/ঘণ্টা",
		"10.0 seconds":           "১০.০ সেকেন্ড",
		"14 km/L":                "১৪ কিমি/লিটার",
		"11 km/L":                "১১ কিমি/লিটার",
		"17 km/L":                "১৭ কিমি/লিটার",
		"BS VI":                  "বিএস ভি",
		"4335 mm":                "৪৩৩৫ মিমি",
		"1814 mm":                "১৮১৪ মিমি",
		"1695 mm":                "১৬৯৫ মিমি",
		"2560 mm":                "২৫৬০ মিমি",
		"1540 mm":                "১৫৪০ মিমি",
		"180 mm":                 "১৮০ মিমি",
		"405 L":                  "৪০৫ লিটার",
		"55 L":                   "৫৫ লিটার",
		"5.6 m":                  "৫.৬ মিটার",
		"1380 kg":                "১৩৮০ কেজি",
		"1880 kg":                "১৮৮০ কেজি",
		"5":                      "৫",
		"MacPherson Strut":       "ম্যাকফারসন স্ট্রাট",
		"Multi-Link":             "মাল্টি-লিঙ্ক",
		"Disc":                   "ডিস্ক",
		"Drum":                   "ড্রাম",
		"215/60 R17":             "২১৫/৬০ আর১৭",
		"17 inches":              "১৭ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red, Orange": "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড, অরেঞ্জ",
		"Halogen":                                "হ্যালোজেন",
		"Halogen DRLs":                           "হ্যালোজেন ডিআরএল",
		"17-inch Alloy Wheels":                   "১৭-ইঞ্চি অ্যালয় হুইল",
		"Fabric":                                 "ফ্যাব্রিক",
		"6-way Manual":                           "৬-ওয়ে ম্যানুয়াল",
		"Manual AC":                              "ম্যানুয়াল এসি",
		"8-inch Touchscreen":                     "৮-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"4 Speakers":                             "৪ স্পিকার",
		"Front Airbags":                          "ফ্রন্ট এয়ারব্যাগ",
		"Rear Parking Sensors":                   "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":                    "রিয়ার পার্কিং ক্যামেরা",
		"2023":                                   "২০২৩",
		"SUV":                                    "এসইউভি",
		"Compact SUV":                            "কমপ্যাক্ট এসইউভি",
	}
}

func (hhs *HavalH2Seeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Great Wall brand
	var brand models.BrandModel
	brandSlug := "great-wall"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Great Wall",
				Slug:   brandSlug,
				Status: 1,
			}
			if err := db.Create(&brand).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// First, find or create the product
	var product models.ProductModel
	productSlug := "haval-h2"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Haval H2",
				Slug:       productSlug,
				CategoryID: &carCategory.ID,
				BrandID:    &brand.ID,
				Status:     1,
			}
			if err := db.Create(&product).Error; err != nil {
				return err
			}
			log.Printf("Created product: %s", product.Name)
		} else {
			return err
		}
	}

	// Get all specification keys
	var specKeys []models.SpecificationKeyModel
	if err := db.Find(&specKeys).Error; err != nil {
		return err
	}

	// Create a map for quick lookup
	specKeyMap := make(map[string]uint)
	for _, key := range specKeys {
		specKeyMap[key.SpecificationKey] = key.ID
	}

	// Define specifications
	specifications := map[string]string{
		"brand":                       "Haval",
		"model":                       "H2",
		"engine_type":                 "1.5L Turbo I4",
		"engine_displacement":         "1497 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "150 hp @ 5600 rpm",
		"engine_max_torque":           "220 Nm @ 2000-4500 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "75.0 x 84.8 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Manual",
		"transmission_gearbox":        "6-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Single Clutch",
		"performance_top_speed":       "175 km/h",
		"performance_acceleration":    "10.0 seconds",
		"performance_mileage_arai":    "14 km/L",
		"performance_mileage_city":    "11 km/L",
		"performance_mileage_highway": "17 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4335 mm",
		"dimensions_width":            "1814 mm",
		"dimensions_height":           "1695 mm",
		"dimensions_wheelbase":        "2560 mm",
		"dimensions_front_tread":      "1540 mm",
		"dimensions_rear_tread":       "1540 mm",
		"dimensions_ground_clearance": "180 mm",
		"dimensions_boot_space":       "405 L",
		"dimensions_fuel_tank":        "55 L",
		"dimensions_turning_radius":   "5.6 m",
		"weight_kerb_weight":          "1380 kg",
		"weight_gross_weight":         "1880 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Multi-Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "215/60 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "215/60 R17",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red, Orange",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "Halogen DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Fabric",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "8-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "4 Speakers",
		"safety_airbags":              "Front Airbags",
		"safety_abs":                  "Yes",
		"safety_ebd":                  "Yes",
		"safety_brake_assist":         "Yes",
		"safety_esp":                  "Yes",
		"safety_parking_sensors":      "Rear Parking Sensors",
		"safety_camera":               "Rear Parking Camera",
		"features_central_locking":    "Yes",
		"features_keyless_entry":      "Yes",
		"features_push_start":         "No",
		"features_cruise_control":     "Yes",
		"features_sunroof":            "No",
		"year":                        "2023",
		"category":                    "SUV",
		"subcategory":                 "Compact SUV",
	}

	// Create specifications
	for key, value := range specifications {
		if specKeyID, exists := specKeyMap[key]; exists {
			spec := models.SpecificationModel{
				ProductID:          product.ID,
				SpecificationKeyID: specKeyID,
				Value:              value,
				Status:             1,
			}
			if err := db.Create(&spec).Error; err != nil {
				log.Printf("Error creating specification for key %s: %v", key, err)
				continue
			}

			// Create translation
			translation := models.SpecificationTranslationModel{
				SpecificationID: spec.ID,
				Locale:          "bn",
				Value:           hhs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Haval H2")
	return nil
}
