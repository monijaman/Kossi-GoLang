package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// IsuzuMUXSeeder seeds specifications for Isuzu MU-X
type IsuzuMUXSeeder struct {
	BaseSeeder
}

// NewIsuzuMUXSeeder creates a new Isuzu MU-X seeder
func NewIsuzuMUXSeeder() *IsuzuMUXSeeder {
	return &IsuzuMUXSeeder{
		BaseSeeder: BaseSeeder{name: "Isuzu MU-X Specifications"},
	}
}

func (ims *IsuzuMUXSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.0L Turbo Diesel I4":         "৩.০ লিটার টার্বো ডিজেল আই৪",
		"2999 cc":                      "২৯৯৯ সিসি",
		"4":                            "৪",
		"Inline":                       "ইনলাইন",
		"Diesel":                       "ডিজেল",
		"177 hp @ 3600 rpm":            "১৭৭ হর্স পাওয়ার @ ৩৬০০ আরপিএম",
		"430 Nm @ 2000-2500 rpm":       "৪৩০ এনএম @ ২০০০-২৫০০ আরপিএম",
		"Common Rail Direct Injection": "কমন রেল ডাইরেক্ট ইনজেকশন",
		"95.4 x 104.9 mm":              "৯৫.৪ x ১০৪.৯ মিমি",
		"16.5:1":                       "১৬.৫:১",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"Automatic":                    "অটোমেটিক",
		"6-Speed":                      "৬-স্পিড",
		"4WD":                          "৪ডব্লিউডি",
		"180 km/h":                     "১৮০ কিমি/ঘণ্টা",
		"11.5 seconds":                 "১১.৫ সেকেন্ড",
		"12 km/L":                      "১২ কিমি/লিটার",
		"9 km/L":                       "৯ কিমি/লিটার",
		"15 km/L":                      "১৫ কিমি/লিটার",
		"BS VI":                        "বিএস ভি",
		"4825 mm":                      "৪৮২৫ মিমি",
		"1860 mm":                      "১৮৬০ মিমি",
		"1830 mm":                      "১৮৩০ মিমি",
		"2845 mm":                      "২৮৪৫ মিমি",
		"1570 mm":                      "১৫৭০ মিমি",
		"230 mm":                       "২৩০ মিমি",
		"235 L":                        "২৩৫ লিটার",
		"65 L":                         "৬৫ লিটার",
		"5.8 m":                        "৫.৮ মিটার",
		"1985 kg":                      "১৯৮৫ কেজি",
		"2600 kg":                      "২৬০০ কেজি",
		"7":                            "৭",
		"Independent Double Wishbone":  "ইন্ডিপেন্ডেন্ট ডাবল উইশবোন",
		"5-Link Rigid":                 "৫-লিঙ্ক রিজিড",
		"Disc":                         "ডিস্ক",
		"255/60 R18":                   "২৫৫/৬০ আর১৮",
		"18 inches":                    "১৮ ইঞ্চি",
		"White, Black, Gray, Silver, Blue, Red, Brown": "হোয়াইট, ব্ল্যাক, গ্রে, সিলভার, ব্লু, রেড, ব্রাউন",
		"LED":                  "এলইডি",
		"LED DRLs":             "এলইডি ডিআরএল",
		"18-inch Alloy Wheels": "১৮-ইঞ্চি অ্যালয় হুইল",
		"Leather":              "লেদার",
		"8-way Power":          "৮-ওয়ে পাওয়ার",
		"Auto AC":              "অটো এসি",
		"9-inch Touchscreen":   "৯-ইঞ্চি টাচস্ক্রিন",
		"Android Auto, Apple CarPlay, Bluetooth, Navigation": "অ্যান্ড্রয়েড অটো, অ্যাপল কারপ্লে, ব্লুটুথ, নেভিগেশন",
		"8 Speakers":                    "৮ স্পিকার",
		"Front, Side & Curtain Airbags": "ফ্রন্ট, সাইড এবং কার্টেন এয়ারব্যাগ",
		"Rear Parking Sensors":          "রিয়ার পার্কিং সেন্সর",
		"360-degree Camera":             "৩৬০-ডিগ্রি ক্যামেরা",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Mid-size SUV":                  "মিড-সাইজ এসইউভি",
	}
}

func (ims *IsuzuMUXSeeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Isuzu brand
	var brand models.BrandModel
	brandSlug := "isuzu"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Isuzu",
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
	productSlug := "isuzu-mu-x"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Isuzu MU-X",
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
		"brand":                       "Isuzu",
		"model":                       "MU-X",
		"engine_type":                 "3.0L Turbo Diesel I4",
		"engine_displacement":         "2999 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Diesel",
		"engine_max_power":            "177 hp @ 3600 rpm",
		"engine_max_torque":           "430 Nm @ 2000-2500 rpm",
		"engine_fuel_supply_system":   "Common Rail Direct Injection",
		"engine_bore_stroke":          "95.4 x 104.9 mm",
		"engine_compression_ratio":    "16.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "6-Speed",
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Automatic",
		"performance_top_speed":       "180 km/h",
		"performance_acceleration":    "11.5 seconds",
		"performance_mileage_arai":    "12 km/L",
		"performance_mileage_city":    "9 km/L",
		"performance_mileage_highway": "15 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4825 mm",
		"dimensions_width":            "1860 mm",
		"dimensions_height":           "1830 mm",
		"dimensions_wheelbase":        "2845 mm",
		"dimensions_front_tread":      "1570 mm",
		"dimensions_rear_tread":       "1570 mm",
		"dimensions_ground_clearance": "230 mm",
		"dimensions_boot_space":       "235 L",
		"dimensions_fuel_tank":        "65 L",
		"dimensions_turning_radius":   "5.8 m",
		"weight_kerb_weight":          "1985 kg",
		"weight_gross_weight":         "2600 kg",
		"capacity_seating_capacity":   "7",
		"capacity_doors":              "5",
		"suspension_front":            "Independent Double Wishbone",
		"suspension_rear":             "5-Link Rigid",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "255/60 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "255/60 R18",
		"exterior_colors":             "White, Black, Gray, Silver, Blue, Red, Brown",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "8-way Power",
		"interior_ac":                 "Auto AC",
		"infotainment_screen_size":    "9-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth, Navigation",
		"infotainment_speakers":       "8 Speakers",
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
				Value:           ims.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Isuzu MU-X")
	return nil
}
