package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JeepCompassSeeder seeds specifications for Jeep Compass
type JeepCompassSeeder struct {
	BaseSeeder
}

// NewJeepCompassSeeder creates a new Jeep Compass seeder
func NewJeepCompassSeeder() *JeepCompassSeeder {
	return &JeepCompassSeeder{
		BaseSeeder: BaseSeeder{name: "Jeep Compass Specifications"},
	}
}

func (jcs *JeepCompassSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":                "২.০ লিটার টার্বো আই৪",
		"1995 cc":                      "১৯৯৫ সিসি",
		"4":                            "৪",
		"Inline":                       "ইনলাইন",
		"Diesel":                       "ডিজেল",
		"170 hp @ 3750 rpm":            "১৭০ হর্স পাওয়ার @ ৩৭৫০ আরপিএম",
		"350 Nm @ 1750 rpm":            "৩৫০ এনএম @ ১৭৫০ আরপিএম",
		"Common Rail Direct Injection": "কমন রেল ডাইরেক্ট ইনজেকশন",
		"83.0 x 90.4 mm":               "৮৩.০ x ৯০.৪ মিমি",
		"16.5:1":                       "১৬.৫:১",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"Automatic":                    "অটোমেটিক",
		"9-Speed":                      "৯-স্পিড",
		"4WD":                          "৪ডব্লিউডি",
		"Active Drive":                 "অ্যাকটিভ ড্রাইভ",
		"186 km/h":                     "১৮৬ কিমি/ঘণ্টা",
		"10.8 seconds":                 "১০.৮ সেকেন্ড",
		"16 km/L":                      "১৬ কিমি/লিটার",
		"12 km/L":                      "১২ কিমি/লিটার",
		"20 km/L":                      "২০ কিমি/লিটার",
		"BS VI":                        "বিএস ভি",
		"4404 mm":                      "৪৪০৪ মিমি",
		"1819 mm":                      "১৮১৯ মিমি",
		"1640 mm":                      "১৬৪০ মিমি",
		"2636 mm":                      "২৬৩৬ মিমি",
		"1546 mm":                      "১৫৪৬ মিমি",
		"1541 mm":                      "১৫৪১ মিমি",
		"206 mm":                       "২০৬ মিমি",
		"438 L":                        "৪৩৮ লিটার",
		"60 L":                         "৬০ লিটার",
		"5.7 m":                        "৫.৭ মিটার",
		"1536 kg":                      "১৫৩৬ কেজি",
		"2087 kg":                      "২০৮৭ কেজি",
		"5":                            "৫",
		"McPherson Strut":              "ম্যাকফারসন স্ট্রাট",
		"Multi-Link":                   "মাল্টি-লিঙ্ক",
		"Disc":                         "ডিস্ক",
		"225/60 R16":                   "২২৫/৬০ আর১৬",
		"16 inches":                    "১৬ ইঞ্চি",
		"Minimal Gray, Bright White, Black, Silver, Red, Blue": "মিনিমাল গ্রে, ব্রাইট হোয়াইট, ব্ল্যাক, সিলভার, রেড, ব্লু",
		"Projector Headlamps":  "প্রজেক্টর হেডল্যাম্প",
		"LED DRLs":             "এলইডি ডিআরএল",
		"16-inch Alloy Wheels": "১৬-ইঞ্চি অ্যালয় হুইল",
		"Cloth":                "ক্লথ",
		"Manual":               "ম্যানুয়াল",
		"Manual AC":            "ম্যানুয়াল এসি",
		"7-inch Touchscreen":   "৭-ইঞ্চি টাচস্ক্রিন",
		"Uconnect, Android Auto, Apple CarPlay, Bluetooth": "ইউকানেক্ট, অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ",
		"6 Speakers":           "৬ স্পিকার",
		"Front Airbags":        "ফ্রন্ট এয়ারব্যাগ",
		"Rear Parking Sensors": "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":  "রিয়ার পার্কিং ক্যামেরা",
		"2023":                 "২০২৩",
		"SUV":                  "এসইউভি",
		"Compact SUV":          "কমপ্যাক্ট এসইউভি",
	}
}

func (jcs *JeepCompassSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "jeep-compass"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jeep Compass",
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
		"model":                       "Compass",
		"engine_type":                 "2.0L Turbo I4",
		"engine_displacement":         "1995 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Diesel",
		"engine_max_power":            "170 hp @ 3750 rpm",
		"engine_max_torque":           "350 Nm @ 1750 rpm",
		"engine_fuel_supply_system":   "Common Rail Direct Injection",
		"engine_bore_stroke":          "83.0 x 90.4 mm",
		"engine_compression_ratio":    "16.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "9-Speed",
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Active Drive",
		"performance_top_speed":       "186 km/h",
		"performance_acceleration":    "10.8 seconds",
		"performance_mileage_arai":    "16 km/L",
		"performance_mileage_city":    "12 km/L",
		"performance_mileage_highway": "20 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4404 mm",
		"dimensions_width":            "1819 mm",
		"dimensions_height":           "1640 mm",
		"dimensions_wheelbase":        "2636 mm",
		"dimensions_front_tread":      "1546 mm",
		"dimensions_rear_tread":       "1541 mm",
		"dimensions_ground_clearance": "206 mm",
		"dimensions_boot_space":       "438 L",
		"dimensions_fuel_tank":        "60 L",
		"dimensions_turning_radius":   "5.7 m",
		"weight_kerb_weight":          "1536 kg",
		"weight_gross_weight":         "2087 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "McPherson Strut",
		"suspension_rear":             "Multi-Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "225/60 R16",
		"tyres_wheel_size":            "16 inches",
		"tyres_spare_size":            "225/60 R16",
		"exterior_colors":             "Minimal Gray, Black, Gray, White, Silver, Blue, Red, Orange, Green",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "16-inch Alloy Wheels",
		"interior_seats_material":     "Cloth",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "7-inch Touchscreen",
		"infotainment_connectivity":   "Uconnect, Android Auto, Apple CarPlay, Bluetooth",
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
		"features_push_start":         "Yes",
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
				Value:           jcs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jeep Compass")
	return nil
}
