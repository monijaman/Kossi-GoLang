package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MahindraScorpioSeeder seeds specifications for Mahindra Scorpio
type MahindraScorpioSeeder struct {
	BaseSeeder
}

// NewMahindraScorpioSeeder creates a new Mahindra Scorpio seeder
func NewMahindraScorpioSeeder() *MahindraScorpioSeeder {
	return &MahindraScorpioSeeder{
		BaseSeeder: BaseSeeder{name: "Mahindra Scorpio Specifications"},
	}
}

func (mss *MahindraScorpioSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.2L mHawk Diesel":       "২.২ লিটার এমহক ডিজেল",
		"2179 cc":                 "২১৭৯ সিসি",
		"4":                       "৪",
		"Inline":                  "ইনলাইন",
		"Diesel":                  "ডিজেল",
		"130 hp @ 3750 rpm":       "১৩০ হর্স পাওয়ার @ ৩৭৫০ আরপিএম",
		"300 Nm @ 1600-2800 rpm":  "৩০০ এনএম @ ১৬০০-২৮০০ আরপিএম",
		"CRDi":                    "সিআরডিআই",
		"85 x 96 mm":              "৮৫ x ৯৬ মিমি",
		"16.5:1":                  "১৬.৫:১",
		"Yes":                     "হ্যাঁ",
		"No":                      "না",
		"Manual":                  "ম্যানুয়াল",
		"6-Speed Manual":          "৬-স্পিড ম্যানুয়াল",
		"RWD":                     "আরডব্লিউডি",
		"Single Plate Dry Clutch": "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"165 km/h":                "১৬৫ কিমি/ঘণ্টা",
		"14.5 seconds":            "১৪.৫ সেকেন্ড",
		"15 km/L":                 "১৫ কিমি/লিটার",
		"12 km/L":                 "১২ কিমি/লিটার",
		"17 km/L":                 "১৭ কিমি/লিটার",
		"BS VI":                   "বিএস ভি",
		"4456 mm":                 "৪৪৫৬ মিমি",
		"1820 mm":                 "১৮২০ মিমি",
		"1995 mm":                 "১৯৯৫ মিমি",
		"2680 mm":                 "২৬৮০ মিমি",
		"1520 mm":                 "১৫২০ মিমি",
		"1515 mm":                 "১৫১৫ মিমি",
		"180 mm":                  "১৮০ মিমি",
		"460 L":                   "৪৬০ লিটার",
		"60 L":                    "৬০ লিটার",
		"5.4 m":                   "৫.৪ মিটার",
		"1730 kg":                 "১৭৩০ কেজি",
		"2510 kg":                 "২৫১০ কেজি",
		"7":                       "৭",
		"5":                       "৫",
		"Double Wishbone":         "ডাবল উইশবোন",
		"Multi-link":              "মাল্টি-লিঙ্ক",
		"Disc":                    "ডিস্ক",
		"Drum":                    "ড্রাম",
		"235/65 R16":              "২৩৫/৬৫ আর১৬",
		"16 inches":               "১৬ ইঞ্চি",
		"Napoli Black, Black, Gray, White, Silver, Blue, Red, Brown, Orange": "নাপোলি ব্ল্যাক, ব্ল্যাক, গ্রে, হোয়াইট, সিলভার, ব্লু, রেড, ব্রাউন, অরেঞ্জ",
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
		"Mid-size SUV":                           "মিড-সাইজ এসইউভি",
	}
}

func (mss *MahindraScorpioSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "mahindra-scorpio"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Mahindra Scorpio",
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
		"model":                       "Scorpio",
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
		"transmission_drive_type":     "RWD",
		"transmission_clutch_type":    "Single Plate Dry Clutch",
		"performance_top_speed":       "165 km/h",
		"performance_acceleration":    "14.5 seconds",
		"performance_mileage_arai":    "15 km/L",
		"performance_mileage_city":    "12 km/L",
		"performance_mileage_highway": "17 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4456 mm",
		"dimensions_width":            "1820 mm",
		"dimensions_height":           "1995 mm",
		"dimensions_wheelbase":        "2680 mm",
		"dimensions_front_tread":      "1520 mm",
		"dimensions_rear_tread":       "1515 mm",
		"dimensions_ground_clearance": "180 mm",
		"dimensions_boot_space":       "460 L",
		"dimensions_fuel_tank":        "60 L",
		"dimensions_turning_radius":   "5.4 m",
		"weight_kerb_weight":          "1730 kg",
		"weight_gross_weight":         "2510 kg",
		"capacity_seating_capacity":   "7",
		"capacity_doors":              "5",
		"suspension_front":            "Double Wishbone",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "235/65 R16",
		"tyres_wheel_size":            "16 inches",
		"tyres_spare_size":            "235/65 R16",
		"exterior_colors":             "Napoli Black, Black, Gray, White, Silver, Blue, Red, Brown, Orange",
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
				Value:           mss.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Mahindra Scorpio")
	return nil
}
