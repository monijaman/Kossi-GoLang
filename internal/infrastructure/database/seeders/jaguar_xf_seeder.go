package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JaguarXFSeeder seeds specifications for Jaguar XF
type JaguarXFSeeder struct {
	BaseSeeder
}

// NewJaguarXFSeeder creates a new Jaguar XF seeder
func NewJaguarXFSeeder() *JaguarXFSeeder {
	return &JaguarXFSeeder{
		BaseSeeder: BaseSeeder{name: "Jaguar XF Specifications"},
	}
}

func (jfs *JaguarXFSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":          "২.০ লিটার টার্বো আই৪",
		"1997 cc":                "১৯৯৭ সিসি",
		"4":                      "৪",
		"Inline":                 "ইনলাইন",
		"Petrol":                 "পেট্রোল",
		"250 hp @ 5500 rpm":      "২৫০ হর্স পাওয়ার @ ৫৫০০ আরপিএম",
		"365 Nm @ 1300-4500 rpm": "৩৬৫ এনএম @ ১৩০০-৪৫০০ আরপিএম",
		"Direct Injection":       "ডাইরেক্ট ইনজেকশন",
		"83.0 x 92.3 mm":         "৮৩.০ x ৯২.৩ মিমি",
		"10.5:1":                 "১০.৫:১",
		"Yes":                    "হ্যাঁ",
		"No":                     "না",
		"Automatic":              "অটোমেটিক",
		"8-Speed":                "৮-স্পিড",
		"RWD":                    "আরডব্লিউডি",
		"250 km/h":               "২৫০ কিমি/ঘণ্টা",
		"6.5 seconds":            "৬.৫ সেকেন্ড",
		"12 km/L":                "১২ কিমি/লিটার",
		"8 km/L":                 "৮ কিমি/লিটার",
		"16 km/L":                "১৬ কিমি/লিটার",
		"BS VI":                  "বিএস ভি",
		"4954 mm":                "৪৯৫৪ মিমি",
		"1880 mm":                "১৮৮০ মিমি",
		"1457 mm":                "১৪৫৭ মিমি",
		"2960 mm":                "২৯৬০ মিমি",
		"1605 mm":                "১৬০৫ মিমি",
		"1595 mm":                "১৫৯৫ মিমি",
		"110 mm":                 "১১০ মিমি",
		"505 L":                  "৫০৫ লিটার",
		"74 L":                   "৭৪ লিটার",
		"5.8 m":                  "৫.৮ মিটার",
		"1720 kg":                "১৭২০ কেজি",
		"2250 kg":                "২২৫০ কেজি",
		"5":                      "৫",
		"Double Wishbone":        "ডাবল উইশবোন",
		"Integral Link":          "ইন্টিগ্রাল লিঙ্ক",
		"Disc":                   "ডিস্ক",
		"235/45 R18":             "২৩৫/৪৫ আর১৮",
		"18 inches":              "১৮ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red, Green": "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড, গ্রিন",
		"LED":                   "এলইডি",
		"LED Matrix":            "এলইডি ম্যাট্রিক্স",
		"18-inch Alloy Wheels":  "১৮-ইঞ্চি অ্যালয় হুইল",
		"Leather":               "লেদার",
		"12-way Power":          "১২-ওয়ে পাওয়ার",
		"Auto AC":               "অটো এসি",
		"11.4-inch Touchscreen": "১১.৪-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth, Navigation": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ, নেভিগেশন",
		"11 Speakers":                   "১১ স্পিকার",
		"Front, Side & Curtain Airbags": "ফ্রন্ট, সাইড এবং কার্টেন এয়ারব্যাগ",
		"Front & Rear Parking Sensors":  "ফ্রন্ট এবং রিয়ার পার্কিং সেন্সর",
		"360-degree Camera":             "৩৬০-ডিগ্রি ক্যামেরা",
		"2023":                          "২০২৩",
		"Sedan":                         "সেডান",
		"Executive Sedan":               "এক্সিকিউটিভ সেডান",
	}
}

func (jfs *JaguarXFSeeder) Seed(db *gorm.DB) error {
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
	productSlug := "jaguar-xf"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jaguar XF",
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
		"engine_type":                 "2.0L Turbo I4",
		"engine_displacement":         "1997 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "250 hp @ 5500 rpm",
		"engine_max_torque":           "365 Nm @ 1300-4500 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "83.0 x 92.3 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "RWD",
		"transmission_clutch_type":    "Automatic",
		"performance_top_speed":       "250 km/h",
		"performance_acceleration":    "6.5 seconds",
		"performance_mileage_arai":    "12 km/L",
		"performance_mileage_city":    "8 km/L",
		"performance_mileage_highway": "16 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4954 mm",
		"dimensions_width":            "1880 mm",
		"dimensions_height":           "1457 mm",
		"dimensions_wheelbase":        "2960 mm",
		"dimensions_front_tread":      "1605 mm",
		"dimensions_rear_tread":       "1595 mm",
		"dimensions_ground_clearance": "110 mm",
		"dimensions_boot_space":       "505 L",
		"dimensions_fuel_tank":        "74 L",
		"dimensions_turning_radius":   "5.8 m",
		"weight_kerb_weight":          "1720 kg",
		"weight_gross_weight":         "2250 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "Double Wishbone",
		"suspension_rear":             "Integral Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "235/45 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "235/45 R18",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red, Green",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED Matrix",
		"exterior_roof_rails":         "No",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "12-way Power",
		"interior_ac":                 "Auto AC",
		"infotainment_screen_size":    "11.4-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth, Navigation",
		"infotainment_speakers":       "11 Speakers",
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
		"subcategory":                 "Executive Sedan",
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
				Value:           jfs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jaguar XF")
	return nil
}
