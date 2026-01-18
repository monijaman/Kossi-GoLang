package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JaguarXJSeeder seeds specifications for Jaguar XJ
type JaguarXJSeeder struct {
	BaseSeeder
}

// NewJaguarXJSeeder creates a new Jaguar XJ seeder
func NewJaguarXJSeeder() *JaguarXJSeeder {
	return &JaguarXJSeeder{
		BaseSeeder: BaseSeeder{name: "Jaguar XJ Specifications"},
	}
}

func (jjs *JaguarXJSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.0L Supercharged V6":   "৩.০ লিটার সুপারচার্জড ভি৬",
		"2995 cc":                "২৯৯৫ সিসি",
		"6":                      "৬",
		"V":                      "ভি",
		"Petrol":                 "পেট্রোল",
		"340 hp @ 6500 rpm":      "৩৪০ হর্স পাওয়ার @ ৬৫০০ আরপিএম",
		"450 Nm @ 3500-5000 rpm": "৪৫০ এনএম @ ৩৫০০-৫০০০ আরপিএম",
		"Direct Injection":       "ডাইরেক্ট ইনজেকশন",
		"84.5 x 89.0 mm":         "৮৪.৫ x ৮৯.০ মিমি",
		"10.5:1":                 "১০.৫:১",
		"No":                     "না",
		"Yes":                    "হ্যাঁ",
		"Automatic":              "অটোমেটিক",
		"8-Speed":                "৮-স্পিড",
		"RWD":                    "আরডব্লিউডি",
		"250 km/h":               "২৫০ কিমি/ঘণ্টা",
		"6.4 seconds":            "৬.৪ সেকেন্ড",
		"10 km/L":                "১০ কিমি/লিটার",
		"7 km/L":                 "৭ কিমি/লিটার",
		"14 km/L":                "১৪ কিমি/লিটার",
		"BS VI":                  "বিএস ভি",
		"5255 mm":                "৫২৫৫ মিমি",
		"1900 mm":                "১৯০০ মিমি",
		"1460 mm":                "১৪৬০ মিমি",
		"3157 mm":                "৩১৫৭ মিমি",
		"1626 mm":                "১৬২৬ মিমি",
		"1604 mm":                "১৬০৪ মিমি",
		"105 mm":                 "১০৫ মিমি",
		"520 L":                  "৫২০ লিটার",
		"82 L":                   "৮২ লিটার",
		"6.2 m":                  "৬.২ মিটার",
		"1870 kg":                "১৮৭০ কেজি",
		"2400 kg":                "২৪০০ কেজি",
		"5":                      "৫",
		"Air Suspension":         "এয়ার সাসপেনশন",
		"Disc":                   "ডিস্ক",
		"245/45 R19":             "২৪৫/৪৫ আর১৯",
		"19 inches":              "১৯ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red, Green, Gold": "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড, গ্রিন, গোল্ড",
		"LED":                   "এলইডি",
		"LED Matrix":            "এলইডি ম্যাট্রিক্স",
		"19-inch Alloy Wheels":  "১৯-ইঞ্চি অ্যালয় হুইল",
		"Leather":               "লেদার",
		"18-way Power":          "১৮-ওয়ে পাওয়ার",
		"Auto AC":               "অটো এসি",
		"12.3-inch Touchscreen": "১২.৩-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth, Navigation": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ, নেভিগেশন",
		"20 Speakers":                   "২০ স্পিকার",
		"Front, Side & Curtain Airbags": "ফ্রন্ট, সাইড এবং কার্টেন এয়ারব্যাগ",
		"Front & Rear Parking Sensors":  "ফ্রন্ট এবং রিয়ার পার্কিং সেন্সর",
		"360-degree Camera":             "৩৬০-ডিগ্রি ক্যামেরা",
		"2023":                          "২০২৩",
		"Sedan":                         "সেডান",
		"Luxury Sedan":                  "লাক্সারি সেডান",
	}
}

func (jjs *JaguarXJSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "jaguar-xj"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jaguar XJ",
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
		"model":                       "XJ",
		"engine_type":                 "3.0L Supercharged V6",
		"engine_displacement":         "2995 cc",
		"engine_cylinders":            "6",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "V",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "340 hp @ 6500 rpm",
		"engine_max_torque":           "450 Nm @ 3500-5000 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "84.5 x 89.0 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "No",
		"engine_super_charger":        "Yes",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "RWD",
		"transmission_clutch_type":    "Automatic",
		"performance_top_speed":       "250 km/h",
		"performance_acceleration":    "6.4 seconds",
		"performance_mileage_arai":    "10 km/L",
		"performance_mileage_city":    "7 km/L",
		"performance_mileage_highway": "14 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "5255 mm",
		"dimensions_width":            "1900 mm",
		"dimensions_height":           "1460 mm",
		"dimensions_wheelbase":        "3157 mm",
		"dimensions_front_tread":      "1626 mm",
		"dimensions_rear_tread":       "1604 mm",
		"dimensions_ground_clearance": "105 mm",
		"dimensions_boot_space":       "520 L",
		"dimensions_fuel_tank":        "82 L",
		"dimensions_turning_radius":   "6.2 m",
		"weight_kerb_weight":          "1870 kg",
		"weight_gross_weight":         "2400 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "Air Suspension",
		"suspension_rear":             "Air Suspension",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "245/45 R19",
		"tyres_wheel_size":            "19 inches",
		"tyres_spare_size":            "245/45 R19",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red, Green, Gold",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED Matrix",
		"exterior_roof_rails":         "No",
		"exterior_wheels":             "19-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "18-way Power",
		"interior_ac":                 "Auto AC",
		"infotainment_screen_size":    "12.3-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth, Navigation",
		"infotainment_speakers":       "20 Speakers",
		"safety_airbags":              "Front, Side & Curtain Airbags",
		"safety_abs":                  "Yes",
		"safety_ebd":                  "Yes",
		"safety_brake_assist":         "Yes",
		"safety_esp":                  "Yes",
		"safety_parking_sensors":      "Front & Rear Parking Sensors",
		"safety_camera":               "360-degree Camera",
		"features_central_locking":    "Yes",
		"features_keyless_entry":      "Yes",
		"features_push_start":         "Yes",
		"features_cruise_control":     "Yes",
		"features_sunroof":            "Yes",
		"year":                        "2023",
		"category":                    "Sedan",
		"subcategory":                 "Luxury Sedan",
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
				Value:           jjs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jaguar XJ")
	return nil
}
