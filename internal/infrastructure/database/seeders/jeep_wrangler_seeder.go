package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JeepWranglerSeeder seeds specifications for Jeep Wrangler
type JeepWranglerSeeder struct {
	BaseSeeder
}

// NewJeepWranglerSeeder creates a new Jeep Wrangler seeder
func NewJeepWranglerSeeder() *JeepWranglerSeeder {
	return &JeepWranglerSeeder{
		BaseSeeder: BaseSeeder{name: "Jeep Wrangler Specifications"},
	}
}

func (jws *JeepWranglerSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":                                      "২.০ লিটার টার্বো আই৪",
		"1995 cc":                                            "১৯৯৫ সিসি",
		"4":                                                  "৪",
		"Inline":                                             "ইনলাইন",
		"Petrol":                                             "পেট্রোল",
		"268 hp @ 5250 rpm":                                  "২৬৮ হর্স পাওয়ার @ ৫২৫০ আরপিএম",
		"400 Nm @ 3000 rpm":                                  "৪০০ এনএম @ ৩০০০ আরপিএম",
		"Direct Injection":                                   "ডাইরেক্ট ইনজেকশন",
		"84.0 x 90.0 mm":                                     "৮৪.০ x ৯০.০ মিমি",
		"10.0:1":                                             "১০.০:১",
		"Yes":                                                "হ্যাঁ",
		"No":                                                 "না",
		"Automatic":                                          "অটোমেটিক",
		"8-Speed":                                            "৮-স্পিড",
		"4WD":                                                "৪ডব্লিউডি",
		"Command-Trac":                                       "কমান্ড-ট্র্যাক",
		"180 km/h":                                           "১৮০ কিমি/ঘণ্টা",
		"8.6 seconds":                                        "৮.৬ সেকেন্ড",
		"11 km/L":                                            "১১ কিমি/লিটার",
		"8 km/L":                                             "৮ কিমি/লিটার",
		"14 km/L":                                            "১৪ কিমি/লিটার",
		"BS VI":                                              "বিএস ভি",
		"4882 mm":                                            "৪৮৮২ মিমি",
		"1894 mm":                                            "১৮৯৪ মিমি",
		"1838 mm":                                            "১৮৩৮ মিমি",
		"3007 mm":                                            "৩০০৭ মিমি",
		"1626 mm":                                            "১৬২৬ মিমি",
		"257 mm":                                             "২৫৭ মিমি",
		"897 L":                                              "৮৯৭ লিটার",
		"81 L":                                               "৮১ লিটার",
		"6.3 m":                                              "৬.৩ মিটার",
		"2010 kg":                                            "২০১০ কেজি",
		"2495 kg":                                            "২৪৯৫ কেজি",
		"5":                                                  "৫",
		"Independent":                                        "ইন্ডিপেন্ডেন্ট",
		"Solid Axle":                                         "সলিড অ্যাক্সেল",
		"Disc":                                               "ডিস্ক",
		"Drum":                                               "ড্রাম",
		"255/75 R17":                                         "২৫৫/৭৫ আর১৭",
		"17 inches":                                          "১৭ ইঞ্চি",
		"Firecracker Red, Black, Gray, White, Silver, Blue, Red, Green, Orange": "ফায়ারক্র্যাকার রেড, ব্ল্যাক, গ্রে, হোয়াইট, সিলভার, ব্লু, রেড, গ্রিন, অরেঞ্জ",
		"Halogen":                                            "হ্যালোজেন",
		"LED DRLs":                                           "এলইডি ডিআরএল",
		"17-inch Alloy Wheels":                               "১৭-ইঞ্চি অ্যালয় হুইল",
		"Cloth":                                              "ক্লথ",
		"4-way Manual":                                       "৪-ওয়ে ম্যানুয়াল",
		"Manual AC":                                          "ম্যানুয়াল এসি",
		"7-inch Touchscreen":                                 "৭-ইঞ্চি টাচস্ক্রিন",
		"Uconnect, Android Auto, Apple CarPlay, Bluetooth":   "ইউকানেক্ট, অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"8 Speakers":                                         "৮ স্পিকার",
		"Front Airbags":                                      "ফ্রন্ট এয়ারব্যাগ",
		"Rear Parking Sensors":                               "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":                                "রিয়ার পার্কিং ক্যামেরা",
		"2023":                                               "২০২৩",
		"SUV":                                                "এসইউভি",
		"Off-road SUV":                                       "অফ-রোড এসইউভি",
	}
}

func (jws *JeepWranglerSeeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Jeep brand
	var brand models.BrandModel
	brandSlug := "jeep"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Jeep",
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
	productSlug := "jeep-wrangler"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jeep Wrangler",
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
		"brand":                       "Jeep",
		"model":                       "Wrangler",
		"engine_type":                 "2.0L Turbo I4",
		"engine_displacement":         "1995 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "268 hp @ 5250 rpm",
		"engine_max_torque":           "400 Nm @ 3000 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "84.0 x 90.0 mm",
		"engine_compression_ratio":    "10.0:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Command-Trac",
		"performance_top_speed":       "180 km/h",
		"performance_acceleration":    "8.6 seconds",
		"performance_mileage_arai":    "11 km/L",
		"performance_mileage_city":    "8 km/L",
		"performance_mileage_highway": "14 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4882 mm",
		"dimensions_width":            "1894 mm",
		"dimensions_height":           "1838 mm",
		"dimensions_wheelbase":        "3007 mm",
		"dimensions_front_tread":      "1626 mm",
		"dimensions_rear_tread":       "1626 mm",
		"dimensions_ground_clearance": "257 mm",
		"dimensions_boot_space":       "897 L",
		"dimensions_fuel_tank":        "81 L",
		"dimensions_turning_radius":   "6.3 m",
		"weight_kerb_weight":          "2010 kg",
		"weight_gross_weight":         "2495 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "Independent",
		"suspension_rear":             "Solid Axle",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "255/75 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "255/75 R17",
		"exterior_colors":             "Firecracker Red, Black, Gray, White, Silver, Blue, Red, Green, Orange",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "No",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Cloth",
		"interior_seats_adjustment":   "4-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "7-inch Touchscreen",
		"infotainment_connectivity":   "Uconnect, Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "8 Speakers",
		"safety_airbags":              "Front Airbags",
		"safety_abs":                  "Yes",
		"safety_ebd":                  "Yes",
		"safety_brake_assist":         "Yes",
		"safety_esp":                  "Yes",
		"safety_parking_sensors":      "Rear Parking Sensors",
		"safety_camera":               "Rear Parking Camera",
		"features_central_locking":    "Yes",
		"features_keyless_entry":      "Yes",
		"features_push_start":         "Yes",
		"features_cruise_control":     "Yes",
		"features_sunroof":            "No",
		"year":                        "2023",
		"category":                    "SUV",
		"subcategory":                 "Off-road SUV",
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
				Value:           jws.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jeep Wrangler")
	return nil
}