package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LexusESSeeder seeds specifications for Lexus ES
type LexusESSeeder struct {
	BaseSeeder
}

// NewLexusESSeeder creates a new Lexus ES seeder
func NewLexusESSeeder() *LexusESSeeder {
	return &LexusESSeeder{
		BaseSeeder: BaseSeeder{name: "Lexus ES Specifications"},
	}
}

func (less *LexusESSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.5L V6":                       "৩.৫ লিটার ভি৬",
		"Executive Sedan":               "এক্সিকিউটিভ সেডান",
		"2023":                          "২০২৩",
		"Sedan":                         "সেডান",
		"Obsidian Black":                "অবসিডিয়ান ব্ল্যাক",
		"V":                             "ভি",
		"3456":                          "৩৪৫৬ সিসি",
		"6":                             "৬ (সিলিন্ডার)",
		"302 hp":                        "৩০২ হর্স পাওয়ার",
		"302 hp @ 6600 rpm":             "৬৬০০ আরপিএম এ ৩০২ হর্স পাওয়ার",
		"367 Nm @ 4700 rpm":             "৪৭০০ আরপিএম এ ৩৬৭ এনএম",
		"Petrol":                        "পেট্রোল",
		"Naturally Aspirated":           "ন্যাচারালি অ্যাসপিরেটেড",
		"6.8 seconds":                   "৬.৮ সেকেন্ড",
		"210 km/h":                      "২১০ কিমি/ঘণ্টা",
		"12 km/L":                       "১২ কিমি/লিটার",
		"8 km/L":                        "৮ কিমি/লিটার",
		"16 km/L":                       "১৬ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"8-Speed Automatic":             "৮-স্পিড অটোমেটিক",
		"Front-Wheel Drive":             "ফ্রন্ট-হুইল ড্রাইভ",
		"All-Wheel Drive":               "অল-হুইল ড্রাইভ",
		"4976 mm":                       "৪৯৭৬ মিমি",
		"1864 mm":                       "১৮৬৪ মিমি",
		"1445 mm":                       "১৪৪৫ মিমি",
		"2870 mm":                       "২৮৭০ মিমি",
		"5":                             "৫ (সিট)",
		"4 Doors":                       "৪ ডোর",
		"5 Seats":                       "৫ সিট",
		"Black":                         "ব্ল্যাক",
		"Gray":                          "গ্রে",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Brown":                         "ব্রাউন",
		"Beige":                         "বেইজ",
		"LED":                           "এলইডি",
		"LED Headlights":                "এলইডি হেডলাইট",
		"LED Daytime Running Lights":    "এলইডি ডে টাইম রানিং লাইট",
		"No Roof Rails":                 "রুফ রেল নেই",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"17-inch Alloy Wheels":          "১৭-ইঞ্চি অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"235/50 R17":                    "২৩৫/৫০ আর১৭",
		"Power Adjustment":              "পাওয়ার অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"10-way Power":                  "১০-ওয়ে পাওয়ার",
		"Leather":                       "লেদার",
		"Semi-Aniline Leather":          "সেমি-অ্যানিলাইন লেদার",
		"Leather Seats":                 "লেদার সিট",
		"Semi-Aniline Leather Seats":    "সেমি-অ্যানিলাইন লেদার সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Auto AC":                       "অটো এসি",
		"Dual Zone":                     "ডুয়াল জোন",
		"12.3-inch Touchscreen":         "১২.৩-ইঞ্চি টাচস্ক্রিন",
		"Lexus Link":                    "লেক্সাস লিঙ্ক",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"10 Speakers":                   "১০ স্পিকার",
		"17 Speakers":                   "১৭ স্পিকার",
		"ABS":                           "এবিএস",
		"EBD":                           "ইবিডি",
		"Brake Assist":                  "ব্রেক অ্যাসিস্ট",
		"ESP":                           "ইএসপি",
		"Traction Control":              "ট্র্যাকশন কন্ট্রোল",
		"Hill Hold Control":             "হিল হোল্ড কন্ট্রোল",
		"Front Airbags":                 "ফ্রন্ট এয়ারব্যাগ",
		"Side Airbags":                  "সাইড এয়ারব্যাগ",
		"Curtain Airbags":               "কার্টেন এয়ারব্যাগ",
		"Knee Airbags":                  "নী এয়ারব্যাগ",
		"Rear Parking Sensors":          "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":           "রিয়ার পার্কিং ক্যামেরা",
		"360-degree Camera":             "৩৬০-ডিগ্রি ক্যামেরা",
		"ISOFIX":                        "আইএসওফিক্স",
		"Child Safety Locks":            "চাইল্ড সেফটি লক",
		"Speed Sensing Door Locks":      "স্পিড সেন্সিং ডোর লক",
		"Central Locking":               "সেন্ট্রাল লকিং",
		"Keyless Entry":                 "কীলেস এন্ট্রি",
		"Push Button Start":             "পুশ বাটন স্টার্ট",
		"Engine Immobilizer":            "ইঞ্জিন ইমোবিলাইজার",
		"Anti-theft Alarm":              "অ্যান্টি-থেফট অ্যালার্ম",
		"Tyre Pressure Monitoring":      "টায়ার প্রেশার মনিটরিং",
		"Follow-me-home Headlamps":      "ফলো-মি-হোম হেডল্যাম্প",
		"Rain Sensing Wipers":           "রেইন সেন্সিং ওয়াইপার",
		"Auto Headlamps":                "অটো হেডল্যাম্প",
		"Chrome Accents":                "ক্রোম অ্যাকসেন্ট",
		"Body Colored Bumpers":          "বডি কালার্ড বাম্পার",
		"Rear Spoiler":                  "রিয়ার স্পয়লার",
		"LED Tail Lights":               "এলইডি টেইল লাইট",
		"Fog Lights":                    "ফগ লাইট",
		"Projector Headlamps":           "প্রজেক্টর হেডল্যাম্প",
		"DRLs":                          "ডিআরএল",
		"LED DRLs":                      "এলইডি ডিআরএল",
		"Digital Instrument Cluster":    "ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার",
		"Analog Instrument Cluster":     "অ্যানালগ ইন্সট্রুমেন্ট ক্লাস্টার",
		"Trip Computer":                 "ট্রিপ কম্পিউটার",
		"Tachometer":                    "ট্যাকোমিটার",
		"Gear Shift Indicator":          "গিয়ার শিফ্ট ইন্ডিকেটর",
		"Door Ajar Warning":             "ডোর অজার ওয়ার্নিং",
		"Low Fuel Warning":              "লো ফুয়েল ওয়ার্নিং",
		"12V Power Outlet":              "১২ভি পাওয়ার আউটলেট",
		"Boot Lamp":                     "বুট ল্যাম্প",
		"Vanity Mirrors":                "ভ্যানিটি মিরর",
		"Adjustable Steering":           "অ্যাডজাস্টেবল স্টিয়ারিং",
		"Tilt & Telescopic":             "টিল্ট & টেলিস্কোপিক",
		"Height Adjustable Steering":    "হাইট অ্যাডজাস্টেবল স্টিয়ারিং",
		"Leather Steering Wheel":        "লেদার স্টিয়ারিং হুইল",
		"Multi-function Steering":       "মাল্টি-ফাংশন স্টিয়ারিং",
		"Audio Controls":                "অডিও কন্ট্রোল",
		"Cruise Control":                "ক্রুজ কন্ট্রোল",
		"Reading Lamps":                 "রিডিং ল্যাম্প",
		"Glove Box Lamp":                "গ্লোভ বক্স ল্যাম্প",
		"Trunk Lamp":                    "ট্রাঙ্ক ল্যাম্প",
		"ORVMs with Turn Indicators":    "ওআরভিএম উইথ টার্ন ইন্ডিকেটর",
		"Power Windows":                 "পাওয়ার উইন্ডো",
		"One Touch Up/Down":             "ওয়ান টাচ আপ/ডাউন",
		"Anti-pinch":                    "অ্যান্টি-পিঞ্চ",
		"Boot Release":                  "বুট রিলিজ",
		"Remote Boot Release":           "রিমোট বুট রিলিজ",
		"Rear Defogger":                 "রিয়ার ডিফগার",
		"Rear Wiper":                    "রিয়ার ওয়াইপার",
		"Heated ORVMs":                  "হিটেড ওআরভিএম",
		"Electrically Adjustable ORVMs": "ইলেকট্রিক্যালি অ্যাডজাস্টেবল ওআরভিএম",
		"Boot Space":                    "বুট স্পেস",
		"454 L":                         "৪৫৪ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"65 L":                          "৬৫ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"150 mm":                        "১৫০ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1680 kg":                       "১৬৮০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2100 kg":                       "২১০০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.9 m":                         "৫.৯ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"Engine Type":                   "ইঞ্জিন টাইপ",
	}
}

func (less *LexusESSeeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Lexus ES").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Lexus ES",
				BrandID:    func() *uint { id := uint(1); return &id }(),
				CategoryID: func() *uint { id := uint(18); return &id }(),
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
		"brand":                       "Lexus",
		"model":                       "ES",
		"engine_type":                 "3.5L V6",
		"engine_displacement":         "3456 cc",
		"engine_cylinders":            "6",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "V",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "302 hp @ 6600 rpm",
		"engine_max_torque":           "367 Nm @ 4700 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "94.0 x 83.0 mm",
		"engine_compression_ratio":    "11.8:1",
		"engine_turbo_charger":        "No",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "N/A",
		"performance_top_speed":       "210 km/h",
		"performance_acceleration":    "6.8 seconds",
		"performance_mileage_arai":    "12 km/L",
		"performance_mileage_city":    "8 km/L",
		"performance_mileage_highway": "16 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4976 mm",
		"dimensions_width":            "1864 mm",
		"dimensions_height":           "1445 mm",
		"dimensions_wheelbase":        "2870 mm",
		"dimensions_front_tread":      "1600 mm",
		"dimensions_rear_tread":       "1605 mm",
		"dimensions_ground_clearance": "150 mm",
		"dimensions_boot_space":       "454 L",
		"dimensions_fuel_tank":        "65 L",
		"dimensions_turning_radius":   "5.9 m",
		"weight_kerb_weight":          "1680 kg",
		"weight_gross_weight":         "2100 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Multi-Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "235/50 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "235/50 R17",
		"exterior_colors":             "Obsidian Black, Black, Gray, White, Silver, Blue, Red, Brown, Beige",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "No",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Semi-Aniline Leather",
		"interior_seats_adjustment":   "10-way Power",
		"interior_ac":                 "Dual Zone Auto AC",
		"infotainment_screen_size":    "12.3-inch Touchscreen",
		"infotainment_connectivity":   "Lexus Link, Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "10 Speakers",
		"safety_airbags":              "Front, Side, Curtain & Knee Airbags",
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
				Value:           less.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Lexus ES")
	return nil
}
