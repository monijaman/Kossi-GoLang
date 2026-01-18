package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MahindraXUV500Seeder seeds specifications for Mahindra XUV500
type MahindraXUV500Seeder struct {
	BaseSeeder
}

// NewMahindraXUV500Seeder creates a new Mahindra XUV500 seeder
func NewMahindraXUV500Seeder() *MahindraXUV500Seeder {
	return &MahindraXUV500Seeder{
		BaseSeeder: BaseSeeder{name: "Mahindra XUV500 Specifications"},
	}
}

func (mx5s *MahindraXUV500Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.2L mHawk Diesel":           "২.২ লিটার এমহক ডিজেল",
		"2179 cc":                     "২১৭৯ সিসি",
		"4":                           "৪",
		"Inline":                      "ইনলাইন",
		"Diesel":                      "ডিজেল",
		"155 hp @ 3750 rpm":           "১৫৫ হর্স পাওয়ার @ ৩৭৫০ আরপিএম",
		"360 Nm @ 1750-2800 rpm":      "৩৬০ এনএম @ ১৭৫০-২৮০০ আরপিএম",
		"CRDi":                        "সিআরডিআই",
		"85 x 96 mm":                  "৮৫ x ৯৬ মিমি",
		"16.5:1":                      "১৬.৫:১",
		"Yes":                         "হ্যাঁ",
		"No":                          "না",
		"Manual":                      "ম্যানুয়াল",
		"6-Speed Manual":              "৬-স্পিড ম্যানুয়াল",
		"FWD":                         "এফডব্লিউডি",
		"Single Plate Dry Clutch":     "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"175 km/h":                    "১৭৫ কিমি/ঘণ্টা",
		"10.5 seconds":                "১০.৫ সেকেন্ড",
		"16 km/L":                     "১৬ কিমি/লিটার",
		"13 km/L":                     "১৩ কিমি/লিটার",
		"18 km/L":                     "১৮ কিমি/লিটার",
		"BS VI":                       "বিএস ভি",
		"4585 mm":                     "৪৫৮৫ মিমি",
		"1890 mm":                     "১৮৯০ মিমি",
		"1785 mm":                     "১৭৮৫ মিমি",
		"2700 mm":                     "২৭০০ মিমি",
		"1600 mm":                     "১৬০০ মিমি",
		"200 mm":                      "২০০ মিমি",
		"720 L":                       "৭২০ লিটার",
		"70 L":                        "৭০ লিটার",
		"5.6 m":                       "৫.৬ মিটার",
		"1780 kg":                     "১৭৮০ কেজি",
		"2450 kg":                     "২৪৫০ কেজি",
		"7":                           "৭",
		"5":                           "৫",
		"Independent Double Wishbone": "ইন্ডিপেন্ডেন্ট ডাবল উইশবোন",
		"Multi-link":                  "মাল্টি-লিঙ্ক",
		"Disc":                        "ডিস্ক",
		"Drum":                        "ড্রাম",
		"235/65 R17":                  "২৩৫/৬৫ আর১৭",
		"17 inches":                   "১৭ ইঞ্চি",
		"Napoli Black, Black, Gray, White, Silver, Blue, Red, Brown, Orange": "নাপোলি ব্ল্যাক, ব্ল্যাক, গ্রে, হোয়াইট, সিলভার, ব্লু, রেড, ব্রাউন, অরেঞ্জ",
		"Halogen":                                "হ্যালোজেন",
		"Halogen DRLs":                           "হ্যালোজেন ডিআরএল",
		"17-inch Alloy Wheels":                   "১৭-ইঞ্চি অ্যালয় হুইল",
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

func (mx5s *MahindraXUV500Seeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Mahindra brand
	var brand models.BrandModel
	brandSlug := "mahindra"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Mahindra",
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
	productSlug := "mahindra-xuv500"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Mahindra XUV500",
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
		"engine_type":                 "2.2L mHawk Diesel",
		"engine_displacement":         "2179 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Diesel",
		"engine_max_power":            "155 hp @ 3750 rpm",
		"engine_max_torque":           "360 Nm @ 1750-2800 rpm",
		"engine_fuel_supply_system":   "CRDi",
		"engine_bore_stroke":          "85 x 96 mm",
		"engine_compression_ratio":    "16.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Manual",
		"transmission_gearbox":        "6-Speed Manual",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Single Plate Dry Clutch",
		"performance_top_speed":       "175 km/h",
		"performance_acceleration":    "10.5 seconds",
		"performance_mileage_arai":    "16 km/L",
		"performance_mileage_city":    "13 km/L",
		"performance_mileage_highway": "18 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4585 mm",
		"dimensions_width":            "1890 mm",
		"dimensions_height":           "1785 mm",
		"dimensions_wheelbase":        "2700 mm",
		"dimensions_front_tread":      "1600 mm",
		"dimensions_rear_tread":       "1600 mm",
		"dimensions_ground_clearance": "200 mm",
		"dimensions_boot_space":       "720 L",
		"dimensions_fuel_tank":        "70 L",
		"dimensions_turning_radius":   "5.6 m",
		"weight_kerb_weight":          "1780 kg",
		"weight_gross_weight":         "2450 kg",
		"capacity_seating_capacity":   "7",
		"capacity_doors":              "5",
		"suspension_front":            "Independent Double Wishbone",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "235/65 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "235/65 R17",
		"exterior_colors":             "Napoli Black, Black, Gray, White, Silver, Blue, Red, Brown, Orange",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "Halogen DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "17-inch Alloy Wheels",
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
				Value:           mx5s.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Mahindra XUV500")
	return nil
}
