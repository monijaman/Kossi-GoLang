package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HavalH6Seeder seeds specifications for Haval H6
type HavalH6Seeder struct {
	BaseSeeder
}

// NewHavalH6Seeder creates a new Haval H6 seeder
func NewHavalH6Seeder() *HavalH6Seeder {
	return &HavalH6Seeder{
		BaseSeeder: BaseSeeder{name: "Haval H6 Specifications"},
	}
}

func (hhs *HavalH6Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":          "২.০ লিটার টার্বো আই৪",
		"1981 cc":                "১৯৮১ সিসি",
		"4":                      "৪",
		"Inline":                 "ইনলাইন",
		"Petrol":                 "পেট্রোল",
		"201 hp @ 5500 rpm":      "২০১ হর্স পাওয়ার @ ৫৫০০ আরপিএম",
		"320 Nm @ 1500-4000 rpm": "৩২০ এনএম @ ১৫০০-৪০০০ আরপিএম",
		"Direct Injection":       "ডাইরেক্ট ইনজেকশন",
		"82.0 x 94.4 mm":         "৮২.০ x ৯৪.৪ মিমি",
		"10.5:1":                 "১০.৫:১",
		"Yes":                    "হ্যাঁ",
		"No":                     "না",
		"Automatic":              "অটোমেটিক",
		"6-Speed":                "৬-স্পিড",
		"FWD":                    "এফডব্লিউডি",
		"Dual Clutch":            "ডুয়াল ক্লাচ",
		"190 km/h":               "১৯০ কিমি/ঘণ্টা",
		"9.7 seconds":            "৯.৭ সেকেন্ড",
		"12 km/L":                "১২ কিমি/লিটার",
		"8 km/L":                 "৮ কিমি/লিটার",
		"16 km/L":                "১৬ কিমি/লিটার",
		"BS VI":                  "বিএস ভি",
		"4649 mm":                "৪৬৪৯ মিমি",
		"1860 mm":                "১৮৬০ মিমি",
		"1690 mm":                "১৬৯০ মিমি",
		"2680 mm":                "২৬৮০ মিমি",
		"1580 mm":                "১৫৮০ মিমি",
		"155 mm":                 "১৫৫ মিমি",
		"808 L":                  "৮০৮ লিটার",
		"58 L":                   "৫৮ লিটার",
		"5.7 m":                  "৫.৭ মিটার",
		"1655 kg":                "১৬৫৫ কেজি",
		"2155 kg":                "২১৫৫ কেজি",
		"5":                      "৫",
		"MacPherson Strut":       "ম্যাকফারসন স্ট্রাট",
		"Multi-Link":             "মাল্টি-লিঙ্ক",
		"Disc":                   "ডিস্ক",
		"Drum":                   "ড্রাম",
		"225/55 R18":             "২২৫/৫৫ আর১৮",
		"18 inches":              "১৮ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red, Brown": "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড, ব্রাউন",
		"Halogen":                                "হ্যালোজেন",
		"Halogen DRLs":                           "হ্যালোজেন ডিআরএল",
		"18-inch Alloy Wheels":                   "১৮-ইঞ্চি অ্যালয় হুইল",
		"Fabric":                                 "ফ্যাব্রিক",
		"6-way Manual":                           "৬-ওয়ে ম্যানুয়াল",
		"Manual AC":                              "ম্যানুয়াল এসি",
		"8-inch Touchscreen":                     "৮-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"6 Speakers":                             "৬ স্পিকার",
		"Front Airbags":                          "ফ্রন্ট এয়ারব্যাগ",
		"Rear Parking Sensors":                   "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":                    "রিয়ার পার্কিং ক্যামেরা",
		"2023":                                   "২০২৩",
		"SUV":                                    "এসইউভি",
		"Mid-size SUV":                           "মিড-সাইজ এসইউভি",
	}
}

func (hhs *HavalH6Seeder) Seed(db *gorm.DB) error {
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
	productSlug := "haval-h6"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Haval H6",
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
		"model":                       "H6",
		"engine_type":                 "2.0L Turbo I4",
		"engine_displacement":         "1981 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "201 hp @ 5500 rpm",
		"engine_max_torque":           "320 Nm @ 1500-4000 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "82.0 x 94.4 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "6-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Dual Clutch",
		"performance_top_speed":       "190 km/h",
		"performance_acceleration":    "9.7 seconds",
		"performance_mileage_arai":    "12 km/L",
		"performance_mileage_city":    "8 km/L",
		"performance_mileage_highway": "16 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4649 mm",
		"dimensions_width":            "1860 mm",
		"dimensions_height":           "1690 mm",
		"dimensions_wheelbase":        "2680 mm",
		"dimensions_front_tread":      "1580 mm",
		"dimensions_rear_tread":       "1580 mm",
		"dimensions_ground_clearance": "155 mm",
		"dimensions_boot_space":       "808 L",
		"dimensions_fuel_tank":        "58 L",
		"dimensions_turning_radius":   "5.7 m",
		"weight_kerb_weight":          "1655 kg",
		"weight_gross_weight":         "2155 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Multi-Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "225/55 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "225/55 R18",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red, Brown",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "Halogen DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Fabric",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "8-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "6 Speakers",
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
		"subcategory":                 "Mid-size SUV",
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

	log.Printf("Seeded specifications for Haval H6")
	return nil
}
