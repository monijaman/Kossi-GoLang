package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HavalJolionSeeder seeds specifications for Haval Jolion
type HavalJolionSeeder struct {
	BaseSeeder
}

// NewHavalJolionSeeder creates a new Haval Jolion seeder
func NewHavalJolionSeeder() *HavalJolionSeeder {
	return &HavalJolionSeeder{
		BaseSeeder: BaseSeeder{name: "Haval Jolion Specifications"},
	}
}

func (hjs *HavalJolionSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Turbo I4":                          "১.৫ লিটার টার্বো আই৪",
		"1497 cc":                                "১৪৯৭ সিসি",
		"4":                                      "৪",
		"Inline":                                 "ইনলাইন",
		"Petrol":                                 "পেট্রোল",
		"143 hp @ 5600 rpm":                      "১৪৩ হর্স পাওয়ার @ ৫৬০০ আরপিএম",
		"210 Nm @ 2200-4500 rpm":                 "২১০ এনএম @ ২২০০-৪৫০০ আরপিএম",
		"Direct Injection":                       "ডাইরেক্ট ইনজেকশন",
		"75.0 x 84.8 mm":                         "৭৫.০ x ৮৪.৮ মিমি",
		"11.0:1":                                 "১১.০:১",
		"Yes":                                    "হ্যাঁ",
		"No":                                     "না",
		"Manual":                                 "ম্যানুয়াল",
		"6-Speed":                                "৬-স্পিড",
		"FWD":                                    "এফডব্লিউডি",
		"Single Clutch":                          "সিঙ্গেল ক্লাচ",
		"170 km/h":                               "১৭০ কিমি/ঘণ্টা",
		"10.5 seconds":                           "১০.৫ সেকেন্ড",
		"15 km/L":                                "১৫ কিমি/লিটার",
		"12 km/L":                                "১২ কিমি/লিটার",
		"18 km/L":                                "১৮ কিমি/লিটার",
		"BS VI":                                  "বিএস ভি",
		"4105 mm":                                "৪১০৫ মিমি",
		"1810 mm":                                "১৮১০ মিমি",
		"1574 mm":                                "১৫৭৪ মিমি",
		"2560 mm":                                "২৫৬০ মিমি",
		"1530 mm":                                "১৫৩০ মিমি",
		"170 mm":                                 "১৭০ মিমি",
		"320 L":                                  "৩২০ লিটার",
		"55 L":                                   "৫৫ লিটার",
		"5.4 m":                                  "৫.৪ মিটার",
		"1320 kg":                                "১৩২০ কেজি",
		"1820 kg":                                "১৮২০ কেজি",
		"5":                                      "৫",
		"MacPherson Strut":                       "ম্যাকফারসন স্ট্রাট",
		"Torsion Beam":                           "টর্শন বিম",
		"Disc":                                   "ডিস্ক",
		"Drum":                                   "ড্রাম",
		"205/60 R16":                             "২০৫/৬০ আর১৬",
		"16 inches":                              "১৬ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red":  "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড",
		"Halogen":                                "হ্যালোজেন",
		"Halogen DRLs":                           "হ্যালোজেন ডিআরএল",
		"16-inch Alloy Wheels":                   "১৬-ইঞ্চি অ্যালয় হুইল",
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

func (hjs *HavalJolionSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "haval-jolion"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Haval Jolion",
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
		"engine_type":                 "1.5L Turbo I4",
		"engine_displacement":         "1497 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "143 hp @ 5600 rpm",
		"engine_max_torque":           "210 Nm @ 2200-4500 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "75.0 x 84.8 mm",
		"engine_compression_ratio":    "11.0:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Manual",
		"transmission_gearbox":        "6-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Single Clutch",
		"performance_top_speed":       "170 km/h",
		"performance_acceleration":    "10.5 seconds",
		"performance_mileage_arai":    "15 km/L",
		"performance_mileage_city":    "12 km/L",
		"performance_mileage_highway": "18 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4105 mm",
		"dimensions_width":            "1810 mm",
		"dimensions_height":           "1574 mm",
		"dimensions_wheelbase":        "2560 mm",
		"dimensions_front_tread":      "1530 mm",
		"dimensions_rear_tread":       "1530 mm",
		"dimensions_ground_clearance": "170 mm",
		"dimensions_boot_space":       "320 L",
		"dimensions_fuel_tank":        "55 L",
		"dimensions_turning_radius":   "5.4 m",
		"weight_kerb_weight":          "1320 kg",
		"weight_gross_weight":         "1820 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Torsion Beam",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "205/60 R16",
		"tyres_wheel_size":            "16 inches",
		"tyres_spare_size":            "205/60 R16",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "Halogen DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "16-inch Alloy Wheels",
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
				Value:           hjs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Haval Jolion")
	return nil
}
