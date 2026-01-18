package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// IsuzuDMaxSeeder seeds specifications for Isuzu D-Max
type IsuzuDMaxSeeder struct {
	BaseSeeder
}

// NewIsuzuDMaxSeeder creates a new Isuzu D-Max seeder
func NewIsuzuDMaxSeeder() *IsuzuDMaxSeeder {
	return &IsuzuDMaxSeeder{
		BaseSeeder: BaseSeeder{name: "Isuzu D-Max Specifications"},
	}
}

func (ids *IsuzuDMaxSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.9L Turbo Diesel I4":                   "১.৯ লিটার টার্বো ডিজেল আই৪",
		"1898 cc":                                "১৮৯৮ সিসি",
		"4":                                      "৪",
		"Inline":                                 "ইনলাইন",
		"Diesel":                                 "ডিজেল",
		"161 hp @ 3600 rpm":                      "১৬১ হর্স পাওয়ার @ ৩৬০০ আরপিএম",
		"360 Nm @ 2000-2500 rpm":                 "৩৬০ এনএম @ ২০০০-২৫০০ আরপিএম",
		"Common Rail Direct Injection":           "কমন রেল ডাইরেক্ট ইনজেকশন",
		"80.0 x 94.4 mm":                         "৮০.০ x ৯৪.৪ মিমি",
		"16.5:1":                                 "১৬.৫:১",
		"Yes":                                    "হ্যাঁ",
		"No":                                     "না",
		"Automatic":                              "অটোমেটিক",
		"6-Speed":                                "৬-স্পিড",
		"4WD":                                    "৪ডব্লিউডি",
		"175 km/h":                               "১৭৫ কিমি/ঘণ্টা",
		"12.5 seconds":                           "১২.৫ সেকেন্ড",
		"14 km/L":                                "১৪ কিমি/লিটার",
		"10 km/L":                                "১০ কিমি/লিটার",
		"18 km/L":                                "১৮ কিমি/লিটার",
		"BS VI":                                  "বিএস ভি",
		"5265 mm":                                "৫২৬৫ মিমি",
		"1870 mm":                                "১৮৭০ মিমি",
		"1790 mm":                                "১৭৯০ মিমি",
		"3095 mm":                                "৩০৯৫ মিমি",
		"1570 mm":                                "১৫৭০ মিমি",
		"220 mm":                                 "২২০ মিমি",
		"1490 L":                                 "১৪৯০ লিটার",
		"76 L":                                   "৭৬ লিটার",
		"6.4 m":                                  "৬.৪ মিটার",
		"1835 kg":                                "১৮৩৫ কেজি",
		"2850 kg":                                "২৮৫০ কেজি",
		"5":                                      "৫",
		"Independent Double Wishbone":            "ইন্ডিপেন্ডেন্ট ডাবল উইশবোন",
		"Leaf Spring":                            "লিফ স্প্রিং",
		"Disc":                                   "ডিস্ক",
		"Drum":                                   "ড্রাম",
		"255/70 R16":                             "২৫৫/৭০ আর১৬",
		"16 inches":                              "১৬ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red":  "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড",
		"Halogen":                                "হ্যালোজেন",
		"LED DRLs":                               "এলইডি ডিআরএল",
		"16-inch Alloy Wheels":                   "১৬-ইঞ্চি অ্যালয় হুইল",
		"Fabric":                                 "ফ্যাব্রিক",
		"6-way Manual":                           "৬-ওয়ে ম্যানুয়াল",
		"Manual AC":                              "ম্যানুয়াল এসি",
		"7-inch Touchscreen":                     "৭-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"6 Speakers":                             "৬ স্পিকার",
		"Front Airbags":                          "ফ্রন্ট এয়ারব্যাগ",
		"Rear Parking Sensors":                   "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":                    "রিয়ার পার্কিং ক্যামেরা",
		"2023":                                   "২০২৩",
		"Pickup Truck":                           "পিকআপ ট্রাক",
		"Mid-size Pickup":                        "মিড-সাইজ পিকআপ",
	}
}

func (ids *IsuzuDMaxSeeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Isuzu brand
	var brand models.BrandModel
	brandSlug := "isuzu"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Isuzu",
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
	productSlug := "isuzu-d-max"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Isuzu D-Max",
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
		"engine_type":                 "1.9L Turbo Diesel I4",
		"engine_displacement":         "1898 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Diesel",
		"engine_max_power":            "161 hp @ 3600 rpm",
		"engine_max_torque":           "360 Nm @ 2000-2500 rpm",
		"engine_fuel_supply_system":   "Common Rail Direct Injection",
		"engine_bore_stroke":          "80.0 x 94.4 mm",
		"engine_compression_ratio":    "16.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "6-Speed",
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Automatic",
		"performance_top_speed":       "175 km/h",
		"performance_acceleration":    "12.5 seconds",
		"performance_mileage_arai":    "14 km/L",
		"performance_mileage_city":    "10 km/L",
		"performance_mileage_highway": "18 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "5265 mm",
		"dimensions_width":            "1870 mm",
		"dimensions_height":           "1790 mm",
		"dimensions_wheelbase":        "3095 mm",
		"dimensions_front_tread":      "1570 mm",
		"dimensions_rear_tread":       "1570 mm",
		"dimensions_ground_clearance": "220 mm",
		"dimensions_boot_space":       "1490 L",
		"dimensions_fuel_tank":        "76 L",
		"dimensions_turning_radius":   "6.4 m",
		"weight_kerb_weight":          "1835 kg",
		"weight_gross_weight":         "2850 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "Independent Double Wishbone",
		"suspension_rear":             "Leaf Spring",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "255/70 R16",
		"tyres_wheel_size":            "16 inches",
		"tyres_spare_size":            "255/70 R16",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "16-inch Alloy Wheels",
		"interior_seats_material":     "Fabric",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "7-inch Touchscreen",
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
		"category":                    "Pickup Truck",
		"subcategory":                 "Mid-size Pickup",
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
				Value:           ids.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Isuzu D-Max")
	return nil
}
