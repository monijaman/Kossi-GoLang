package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MahindraTharSeeder seeds specifications for Mahindra Thar
type MahindraTharSeeder struct {
	BaseSeeder
}

// NewMahindraTharSeeder creates a new Mahindra Thar seeder
func NewMahindraTharSeeder() *MahindraTharSeeder {
	return &MahindraTharSeeder{
		BaseSeeder: BaseSeeder{name: "Mahindra Thar Specifications"},
	}
}

func (mts *MahindraTharSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.2L mHawk Diesel":           "২.২ লিটার এমহক ডিজেল",
		"2179 cc":                     "২১৭৯ সিসি",
		"4":                           "৪",
		"Inline":                      "ইনলাইন",
		"Diesel":                      "ডিজেল",
		"130 hp @ 3750 rpm":           "১৩০ হর্স পাওয়ার @ ৩৭৫০ আরপিএম",
		"300 Nm @ 1600-2800 rpm":      "৩০০ এনএম @ ১৬০০-২৮০০ আরপিএম",
		"CRDi":                        "সিআরডিআই",
		"85 x 96 mm":                  "৮৫ x ৯৬ মিমি",
		"16.5:1":                      "১৬.৫:১",
		"Yes":                         "হ্যাঁ",
		"No":                          "না",
		"Manual":                      "ম্যানুয়াল",
		"6-Speed Manual":              "৬-স্পিড ম্যানুয়াল",
		"4WD":                         "৪ডব্লিউডি",
		"Single Plate Dry Clutch":     "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"155 km/h":                    "১৫৫ কিমি/ঘণ্টা",
		"14.3 seconds":                "১৪.৩ সেকেন্ড",
		"13 km/L":                     "১৩ কিমি/লিটার",
		"10 km/L":                     "১০ কিমি/লিটার",
		"15 km/L":                     "১৫ কিমি/লিটার",
		"BS VI":                       "বিএস ভি",
		"3985 mm":                     "৩৯৮৫ মিমি",
		"1820 mm":                     "১৮২০ মিমি",
		"1844 mm":                     "১৮৪৪ মিমি",
		"2450 mm":                     "২৪৫০ মিমি",
		"1440 mm":                     "১৪৪০ মিমি",
		"226 mm":                      "২২৬ মিমি",
		"300 L":                       "৩০০ লিটার",
		"57 L":                        "৫৭ লিটার",
		"5.25 m":                      "৫.২৫ মিটার",
		"1752 kg":                     "১৭৫২ কেজি",
		"2450 kg":                     "২৪৫০ কেজি",
		"3":                           "৩",
		"Independent Double Wishbone": "ইন্ডিপেন্ডেন্ট ডাবল উইশবোন",
		"Non-independent Multi-link":  "নন-ইন্ডিপেন্ডেন্ট মাল্টি-লিঙ্ক",
		"Disc":                        "ডিস্ক",
		"Drum":                        "ড্রাম",
		"255/65 R18":                  "২৫৫/৬৫ আর১৮",
		"18 inches":                   "১৮ ইঞ্চি",
		"Red Rage, Black, Gray, White, Silver, Blue, Red, Brown, Orange": "রেড রেজ, ব্ল্যাক, গ্রে, হোয়াইট, সিলভার, ব্লু, রেড, ব্রাউন, অরেঞ্জ",
		"Halogen":                                "হ্যালোজেন",
		"Halogen DRLs":                           "হ্যালোজেন ডিআরএল",
		"16-inch Alloy Wheels":                   "১৬-ইঞ্চি অ্যালয় হুইল",
		"Fabric":                                 "ফ্যাব্রিক",
		"6-way Manual":                           "৬-ওয়ে ম্যানুয়াল",
		"Manual AC":                              "ম্যানুয়াল এসি",
		"7-inch Touchscreen":                     "৭-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"4 Speakers":                             "৪ স্পিকার",
		"Front Airbags":                          "ফ্রন্ট এয়ারব্যাগ",
		"Rear Parking Sensors":                   "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":                    "রিয়ার পার্কিং ক্যামেরা",
		"2023":                                   "২০২৩",
		"SUV":                                    "এসইউভি",
		"Off-road SUV":                           "অফ-রোড এসইউভি",
	}
}

func (mts *MahindraTharSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "mahindra-thar"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Mahindra Thar",
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
		"brand":                       "Mahindra",
		"model":                       "Thar",
		"engine_type":                 "2.2L mHawk Diesel",
		"engine_displacement":         "2179 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Diesel",
		"engine_max_power":            "130 hp @ 3750 rpm",
		"engine_max_torque":           "300 Nm @ 1600-2800 rpm",
		"engine_fuel_supply_system":   "CRDi",
		"engine_bore_stroke":          "85 x 96 mm",
		"engine_compression_ratio":    "16.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Manual",
		"transmission_gearbox":        "6-Speed Manual",
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Single Plate Dry Clutch",
		"performance_top_speed":       "155 km/h",
		"performance_acceleration":    "14.3 seconds",
		"performance_mileage_arai":    "13 km/L",
		"performance_mileage_city":    "10 km/L",
		"performance_mileage_highway": "15 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "3985 mm",
		"dimensions_width":            "1820 mm",
		"dimensions_height":           "1844 mm",
		"dimensions_wheelbase":        "2450 mm",
		"dimensions_front_tread":      "1440 mm",
		"dimensions_rear_tread":       "1440 mm",
		"dimensions_ground_clearance": "226 mm",
		"dimensions_boot_space":       "300 L",
		"dimensions_fuel_tank":        "57 L",
		"dimensions_turning_radius":   "5.25 m",
		"weight_kerb_weight":          "1752 kg",
		"weight_gross_weight":         "2450 kg",
		"capacity_seating_capacity":   "4",
		"capacity_doors":              "3",
		"suspension_front":            "Independent Double Wishbone",
		"suspension_rear":             "Non-independent Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "255/65 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "255/65 R18",
		"exterior_colors":             "Red Rage, Black, Gray, White, Silver, Blue, Red, Brown, Orange",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "Halogen DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "16-inch Alloy Wheels",
		"interior_seats_material":     "Fabric",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "7-inch Touchscreen",
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
				Value:           mts.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Mahindra Thar")
	return nil
}
