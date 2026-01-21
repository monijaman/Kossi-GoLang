package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JeepGrandCherokeeSeeder seeds specifications for Jeep Grand Cherokee
type JeepGrandCherokeeSeeder struct {
	BaseSeeder
}

// NewJeepGrandCherokeeSeeder creates a new Jeep Grand Cherokee seeder
func NewJeepGrandCherokeeSeeder() *JeepGrandCherokeeSeeder {
	return &JeepGrandCherokeeSeeder{
		BaseSeeder: BaseSeeder{name: "Jeep Grand Cherokee Specifications"},
	}
}

func (jgcs *JeepGrandCherokeeSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.6L V6":                    "৩.৬ লিটার ভি৬",
		"3604 cc":                    "৩৬০৪ সিসি",
		"6":                          "৬",
		"4":                          "৪",
		"V":                          "ভি",
		"Petrol":                     "পেট্রোল",
		"295 hp @ 6400 rpm":          "২৯৫ হর্স পাওয়ার @ ৬৪০০ আরপিএম",
		"353 Nm @ 4000 rpm":          "৩৫৩ এনএম @ ৪০০০ আরপিএম",
		"Multi-Point Fuel Injection": "মাল্টি-পয়েন্ট ফুয়েল ইনজেকশন",
		"96.0 x 83.0 mm":             "৯৬.০ x ৮৩.০ মিমি",
		"10.2:1":                     "১০.২:১",
		"No":                         "না",
		"Automatic":                  "অটোমেটিক",
		"8-Speed":                    "৮-স্পিড",
		"4WD":                        "৪ডব্লিউডি",
		"Quadra-Trac II":             "কোয়াড্রা-ট্র্যাক II",
		"206 km/h":                   "২০৬ কিমি/ঘণ্টা",
		"8.2 seconds":                "৮.২ সেকেন্ড",
		"10 km/L":                    "১০ কিমি/লিটার",
		"7 km/L":                     "৭ কিমি/লিটার",
		"13 km/L":                    "১৩ কিমি/লিটার",
		"BS VI":                      "বিএস ভি",
		"4914 mm":                    "৪৯১৪ মিমি",
		"1943 mm":                    "১৯৪৩ মিমি",
		"1802 mm":                    "১৮০২ মিমি",
		"2915 mm":                    "২৯১৫ মিমি",
		"1643 mm":                    "১৬৪৩ মিমি",
		"1642 mm":                    "১৬৪২ মিমি",
		"218 mm":                     "২১৮ মিমি",
		"1064 L":                     "১০৬৪ লিটার",
		"93 L":                       "৯৩ লিটার",
		"5.7 m":                      "৫.৭ মিটার",
		"2177 kg":                    "২১৭৭ কেজি",
		"2948 kg":                    "২৯৪৮ কেজি",
		"5":                          "৫",
		"Independent Multi-Link":     "ইন্ডিপেন্ডেন্ট মাল্টি-লিঙ্ক",
		"Multi-Link":                 "মাল্টি-লিঙ্ক",
		"Disc":                       "ডিস্ক",
		"265/60 R18":                 "২৬৫/৬০ আর১৮",
		"18 inches":                  "১৮ ইঞ্চি",
		"Diamond Black, Black, Gray, White, Silver, Blue, Red, Green, Brown": "ডায়মন্ড ব্ল্যাক, ব্ল্যাক, গ্রে, হোয়াইট, সিলভার, ব্লু, রেড, গ্রিন, ব্রাউন",
		"LED":                  "এলইডি",
		"LED DRLs":             "এলইডি ডিআরএল",
		"Yes":                  "হ্যাঁ",
		"18-inch Alloy Wheels": "১৮-ইঞ্চি অ্যালয় হুইল",
		"Nappa Leather":        "নাপ্পা লেদার",
		"8-way Power":          "৮-ওয়ে পাওয়ার",
		"Dual Zone Auto AC":    "ডুয়াল জোন অটো এসি",
		"8.4-inch Touchscreen": "৮.৪-ইঞ্চি টাচস্ক্রিন",
		"Uconnect, Android Auto, Apple CarPlay, Bluetooth": "ইউকানেক্ট, অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"9 Speakers":                    "৯ স্পিকার",
		"Front, Side & Curtain Airbags": "ফ্রন্ট, সাইড & কার্টেন এয়ারব্যাগ",
		"Rear Parking Sensors":          "রিয়ার পার্কিং সেন্সর",
		"360-degree Camera":             "৩৬০-ডিগ্রি ক্যামেরা",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Luxury SUV":                    "লাক্সারি এসইউভি",
	}
}

func (jgcs *JeepGrandCherokeeSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "jeep-grand-cherokee"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jeep Grand Cherokee",
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
		"model":                       "Grand Cherokee",
		"engine_type":                 "3.6L V6",
		"engine_displacement":         "3604 cc",
		"engine_cylinders":            "6",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "V",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "295 hp @ 6400 rpm",
		"engine_max_torque":           "353 Nm @ 4000 rpm",
		"engine_fuel_supply_system":   "Multi-Point Fuel Injection",
		"engine_bore_stroke":          "96.0 x 83.0 mm",
		"engine_compression_ratio":    "10.2:1",
		"engine_turbo_charger":        "No",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Quadra-Trac II",
		"performance_top_speed":       "206 km/h",
		"performance_acceleration":    "8.2 seconds",
		"performance_mileage_arai":    "10 km/L",
		"performance_mileage_city":    "7 km/L",
		"performance_mileage_highway": "13 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4914 mm",
		"dimensions_width":            "1943 mm",
		"dimensions_height":           "1802 mm",
		"dimensions_wheelbase":        "2915 mm",
		"dimensions_front_tread":      "1643 mm",
		"dimensions_rear_tread":       "1642 mm",
		"dimensions_ground_clearance": "218 mm",
		"dimensions_boot_space":       "1064 L",
		"dimensions_fuel_tank":        "93 L",
		"dimensions_turning_radius":   "5.7 m",
		"weight_kerb_weight":          "2177 kg",
		"weight_gross_weight":         "2948 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "Independent Multi-Link",
		"suspension_rear":             "Multi-Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "265/60 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "265/60 R18",
		"exterior_colors":             "Diamond Black, Black, Gray, White, Silver, Blue, Red, Green, Brown",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Nappa Leather",
		"interior_seats_adjustment":   "8-way Power",
		"interior_ac":                 "Dual Zone Auto AC",
		"infotainment_screen_size":    "8.4-inch Touchscreen",
		"infotainment_connectivity":   "Uconnect, Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "9 Speakers",
		"safety_airbags":              "Front, Side & Curtain Airbags",
		"safety_abs":                  "Yes",
		"safety_ebd":                  "Yes",
		"safety_brake_assist":         "Yes",
		"safety_esp":                  "Yes",
		"safety_parking_sensors":      "Rear Parking Sensors",
		"safety_camera":               "360-degree Camera",
		"features_central_locking":    "Yes",
		"features_keyless_entry":      "Yes",
		"features_push_start":         "Yes",
		"features_cruise_control":     "Yes",
		"features_sunroof":            "Yes",
		"year":                        "2023",
		"category":                    "SUV",
		"subcategory":                 "Luxury SUV",
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
				Value:           jgcs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jeep Grand Cherokee")
	return nil
}
