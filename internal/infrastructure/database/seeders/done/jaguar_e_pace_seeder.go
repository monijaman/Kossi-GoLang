package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JaguarEPACESeeder seeds specifications for Jaguar E-PACE
type JaguarEPACESeeder struct {
	BaseSeeder
}

// NewJaguarEPACESeeder creates a new Jaguar E-PACE seeder
func NewJaguarEPACESeeder() *JaguarEPACESeeder {
	return &JaguarEPACESeeder{
		BaseSeeder: BaseSeeder{name: "Jaguar E-PACE Specifications"},
	}
}

func (jeps *JaguarEPACESeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":          "২.০ লিটার টার্বো আই৪",
		"1997 cc":                "১৯৯৭ সিসি",
		"4":                      "৪",
		"Inline":                 "ইনলাইন",
		"Petrol":                 "পেট্রোল",
		"200 hp @ 5500 rpm":      "২০০ হর্স পাওয়ার @ ৫৫০০ আরপিএম",
		"320 Nm @ 1200-4000 rpm": "৩২০ এনএম @ ১২০০-৪০০০ আরপিএম",
		"Direct Injection":       "ডাইরেক্ট ইনজেকশন",
		"83.0 x 92.3 mm":         "৮৩.০ x ৯২.৩ মিমি",
		"10.5:1":                 "১০.৫:১",
		"Yes":                    "হ্যাঁ",
		"No":                     "না",
		"Automatic":              "অটোমেটিক",
		"9-Speed":                "৯-স্পিড",
		"FWD":                    "এফডব্লিউডি",
		"200 km/h":               "২০০ কিমি/ঘণ্টা",
		"8.0 seconds":            "৮.০ সেকেন্ড",
		"14 km/L":                "১৪ কিমি/লিটার",
		"10 km/L":                "১০ কিমি/লিটার",
		"18 km/L":                "১৮ কিমি/লিটার",
		"BS VI":                  "বিএস ভি",
		"4395 mm":                "৪৩৯৫ মিমি",
		"1984 mm":                "১৯৮৪ মিমি",
		"1649 mm":                "১৬৪৯ মিমি",
		"2681 mm":                "২৬৮১ মিমি",
		"1621 mm":                "১৬২১ মিমি",
		"204 mm":                 "২০৪ মিমি",
		"577 L":                  "৫৭৭ লিটার",
		"56 L":                   "৫৬ লিটার",
		"5.7 m":                  "৫.৭ মিটার",
		"1665 kg":                "১৬৬৫ কেজি",
		"2200 kg":                "২২০০ কেজি",
		"5":                      "৫",
		"MacPherson Strut":       "ম্যাকফারসন স্ট্রাট",
		"Integral Link":          "ইন্টিগ্রাল লিঙ্ক",
		"Disc":                   "ডিস্ক",
		"235/60 R18":             "২৩৫/৬০ আর১৮",
		"18 inches":              "১৮ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red, Green": "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড, গ্রিন",
		"LED":                  "এলইডি",
		"LED Matrix":           "এলইডি ম্যাট্রিক্স",
		"18-inch Alloy Wheels": "১৮-ইঞ্চি অ্যালয় হুইল",
		"Leather":              "লেদার",
		"10-way Power":         "১০-ওয়ে পাওয়ার",
		"Auto AC":              "অটো এসি",
		"10-inch Touchscreen":  "১০-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth, Navigation": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ, নেভিগেশন",
		"6 Speakers":                    "৬ স্পিকার",
		"Front, Side & Curtain Airbags": "ফ্রন্ট, সাইড এবং কার্টেন এয়ারব্যাগ",
		"Front & Rear Parking Sensors":  "ফ্রন্ট এবং রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":           "রিয়ার পার্কিং ক্যামেরা",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Compact Luxury SUV":            "কমপ্যাক্ট লাক্সারি এসইউভি",
	}
}

func (jeps *JaguarEPACESeeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Jaguar brand
	var brand models.BrandModel
	brandSlug := "jaguar"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Jaguar",
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
	productSlug := "jaguar-e-pace"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jaguar E-PACE",
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
		"brand":                       "Jaguar",
		"model":                       "E-PACE",
		"engine_type":                 "2.0L Turbo I4",
		"engine_displacement":         "1997 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "200 hp @ 5500 rpm",
		"engine_max_torque":           "320 Nm @ 1200-4000 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "83.0 x 92.3 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "9-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Automatic",
		"performance_top_speed":       "200 km/h",
		"performance_acceleration":    "8.0 seconds",
		"performance_mileage_arai":    "14 km/L",
		"performance_mileage_city":    "10 km/L",
		"performance_mileage_highway": "18 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4395 mm",
		"dimensions_width":            "1984 mm",
		"dimensions_height":           "1649 mm",
		"dimensions_wheelbase":        "2681 mm",
		"dimensions_front_tread":      "1621 mm",
		"dimensions_rear_tread":       "1621 mm",
		"dimensions_ground_clearance": "204 mm",
		"dimensions_boot_space":       "577 L",
		"dimensions_fuel_tank":        "56 L",
		"dimensions_turning_radius":   "5.7 m",
		"weight_kerb_weight":          "1665 kg",
		"weight_gross_weight":         "2200 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Integral Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "235/60 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "235/60 R18",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red, Green",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED Matrix",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "10-way Power",
		"interior_ac":                 "Auto AC",
		"infotainment_screen_size":    "10-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth, Navigation",
		"infotainment_speakers":       "6 Speakers",
		"safety_airbags":              "Front, Side & Curtain Airbags",
		"safety_abs":                  "Yes",
		"safety_ebd":                  "Yes",
		"safety_brake_assist":         "Yes",
		"safety_esp":                  "Yes",
		"safety_parking_sensors":      "Front & Rear Parking Sensors",
		"safety_camera":               "Rear Parking Camera",
		"features_central_locking":    "Yes",
		"features_keyless_entry":      "Yes",
		"features_push_start":         "Yes",
		"features_cruise_control":     "Yes",
		"features_sunroof":            "Yes",
		"year":                        "2023",
		"category":                    "SUV",
		"subcategory":                 "Compact Luxury SUV",
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
				Value:           jeps.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jaguar E-PACE")
	return nil
}
