package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// Peugeot5008Seeder seeds specifications for Peugeot 5008
type Peugeot5008Seeder struct {
	BaseSeeder
}

// NewPeugeot5008Seeder creates a new Peugeot 5008 seeder
func NewPeugeot5008Seeder() *Peugeot5008Seeder {
	return &Peugeot5008Seeder{
		BaseSeeder: BaseSeeder{name: "Peugeot 5008 Specifications"},
	}
}

func (p5s *Peugeot5008Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.6L Turbo Hybrid":             "১.৬ লিটার টার্বো হাইব্রিড",
		"Mid-size SUV":                  "মিড-সাইজ এসইউভি",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Pearl White":                   "পার্ল হোয়াইট",
		"Inline":                        "ইনলাইন",
		"1598":                          "১৫৯৮ সিসি",
		"4":                             "৪ (সিলিন্ডার)",
		"180 hp":                        "১৮০ হর্স পাওয়ার",
		"180 hp @ 5500 rpm":             "৫৫০০ আরপিএম এ ১৮০ হর্স পাওয়ার",
		"300 Nm @ 3000 rpm":             "৩০০০ আরপিএম এ ৩০০ এনএম",
		"Hybrid":                        "হাইব্রিড",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"8.3 seconds":                   "৮.৩ সেকেন্ড",
		"205 km/h":                      "২০৫ কিমি/ঘণ্টা",
		"18 km/L":                       "১৮ কিমি/লিটার",
		"14 km/L":                       "১৪ কিমি/লিটার",
		"22 km/L":                       "২২ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"8-Speed Automatic":             "৮-স্পিড অটোমেটিক",
		"Front-Wheel Drive":             "ফ্রন্ট-হুইল ড্রাইভ",
		"All-Wheel Drive":               "অল-হুইল ড্রাইভ",
		"Dual Clutch":                   "ডুয়াল ক্লাচ",
		"4670 mm":                       "৪৬৭০ মিমি",
		"1844 mm":                       "১৮৪৪ মিমি",
		"1640 mm":                       "১৬৪০ মিমি",
		"2840 mm":                       "২৮৪০ মিমি",
		"7":                             "৭ (সিট)",
		"5 Doors":                       "৫ ডোর",
		"7 Seats":                       "৭ সিট",
		"Black":                         "ব্ল্যাক",
		"Gray":                          "গ্রে",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Brown":                         "ব্রাউন",
		"Green":                         "গ্রিন",
		"LED":                           "এলইডি",
		"Full LED":                      "ফুল এলইডি",
		"LED Headlights":                "এলইডি হেডলাইট",
		"LED Daytime Running Lights":    "এলইডি ডে টাইম রানিং লাইট",
		"Roof Rails":                    "রুফ রেল",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"19-inch Alloy Wheels":          "১৯-ইঞ্চি অ্যালয় হুইল",
		"225/55 R18":                    "২২৫/৫৫ আর১৮",
		"Electric Adjustment":           "ইলেকট্রিক অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"8-way Electric":                "৮-ওয়ে ইলেকট্রিক",
		"Leather":                       "লেদার",
		"Leather Seats":                 "লেদার সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Third Row Seats":               "থার্ড রো সিট",
		"Auto AC":                       "অটো এসি",
		"Dual Zone":                     "ডুয়াল জোন",
		"10-inch Touchscreen":           "১০-ইঞ্চি টাচস্ক্রিন",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"6 Speakers":                    "৬ স্পিকার",
		"8 Speakers":                    "৮ স্পিকার",
		"ABS":                           "এবিএস",
		"EBD":                           "ইবিডি",
		"Brake Assist":                  "ব্রেক অ্যাসিস্ট",
		"ESP":                           "ইএসপি",
		"Traction Control":              "ট্র্যাকশন কন্ট্রোল",
		"Hill Hold Control":             "হিল হোল্ড কন্ট্রোল",
		"Front Airbags":                 "ফ্রন্ট এয়ারব্যাগ",
		"Side Airbags":                  "সাইড এয়ারব্যাগ",
		"Curtain Airbags":               "কার্টেন এয়ারব্যাগ",
		"Rear Parking Sensors":          "রিয়ার পার্কিং সেন্সর",
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
		"Sunroof":                       "সানরুফ",
		"Panoramic Sunroof":             "প্যানোরামিক সানরুফ",
		"Chrome Accents":                "ক্রোম অ্যাকসেন্ট",
		"Body Colored Bumpers":          "বডি কালার্ড বাম্পার",
		"Rear Spoiler":                  "রিয়ার স্পয়লার",
		"LED Tail Lights":               "এলইডি টেইল লাইট",
		"Fog Lights":                    "ফগ লাইট",
		"Projector Headlamps":           "প্রজেক্টর হেডল্যাম্প",
		"DRLs":                          "ডিআরএল",
		"LED DRLs":                      "এলইডি ডিআরএল",
		"Digital Instrument Cluster":    "ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার",
		"12.3-inch Digital Cluster":     "১২.৩-ইঞ্চি ডিজিটাল ক্লাস্টার",
		"Trip Computer":                 "ট্রিপ কম্পিউটার",
		"Tachometer":                    "ট্যাকোমিটার",
		"Gear Shift Indicator":          "গিয়ার শিফ্ট ইন্ডিকেটর",
		"Door Ajar Warning":             "ডোর অজার ওয়ার্নিং",
		"Low Fuel Warning":              "লো ফুয়েল ওয়ার্নিং",
		"12V Power Outlet":              "১২ভি পাওয়ার আউটলেট",
		"Boot Lamp":                     "বুট ল্যাম্প",
		"Vanity Mirrors":                "ভ্যানিটি মিরর",
		"Adjustable Steering":           "অ্যাডজাস্টেবল স্টিয়ারিং",
		"Tilt & Telescopic":             "টিল্ট এবং টেলিস্কোপিক",
		"Leather Steering Wheel":        "লেদার স্টিয়ারিং হুইল",
		"Multi-function Steering":       "মাল্টি-ফাংশন স্টিয়ারিং",
		"Audio Controls":                "অডিও কন্ট্রোল",
		"Cruise Control":                "ক্রুজ কন্ট্রোল",
		"Adaptive Cruise Control":       "অ্যাডাপটিভ ক্রুজ কন্ট্রোল",
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
		"780 L":                         "৭৮০ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"56 L":                          "৫৬ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"236 mm":                        "২৩৬ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1490 kg":                       "১৪৯০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2130 kg":                       "২১৩০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.2 m":                         "৫.২ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1598 cc":                       "১৫৯৮ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"77.0 x 85.8 mm":                "৭৭.০ x ৮৫.৮ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"10.5:1":                        "১০.৫:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
		"8-Speed":                       "৮-স্পিড",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"FWD":                           "এফডব্লিউডি",
		"AWD":                           "এডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"18 km/L":                       "১৮ কিমি/লিটার",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"14 km/L":                       "১৪ কিমি/লিটার",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"22 km/L":                       "২২ কিমি/লিটার",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"4670 mm":                       "৪৬৭০ মিমি",
		"Width":                         "প্রস্থ",
		"1844 mm":                       "১৮৪৪ মিমি",
		"Height":                        "উচ্চতা",
		"1640 mm":                       "১৬৪০ মিমি",
		"Wheelbase":                     "হুইলবেস",
		"2840 mm":                       "২৮৪০ মিমি",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1608 mm":                       "১৬০৮ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1617 mm":                       "১৬১৭ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"7":                             "৭",
		"Door Count":                    "ডোর কাউন্ট",
		"5":                             "৫",
		"Boot Space":                    "বুট স্পেস",
		"780 L":                         "৭৮০ লিটার",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"56 L":                          "৫৬ লিটার",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"236 mm":                        "২৩৬ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1490 kg":                       "১৪৯০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2130 kg":                       "২১৩০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.2 m":                         "৫.২ মিটার",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"MacPherson Strut":              "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Multi-link":                    "মাল্টি-লিঙ্ক",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Tyre Size":                     "টায়ার সাইজ",
		"225/55 R18":                    "২২৫/৫৫ আর১৮",
		"Wheel Size":                    "হুইল সাইজ",
		"18 inches":                     "১৮ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
		"225/55 R18":                    "২২৫/৫৫ আর১৮",
	}
}

func (p5s *Peugeot5008Seeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Peugeot 5008").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:        "Peugeot 5008",
				Brand:       "Peugeot",
				Category:    "SUV",
				Subcategory: "Mid-size SUV",
				Status:      1,
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
		specKeyMap[key.Key] = key.ID
	}

	// Define specifications
	specifications := map[string]string{
		"engine_type":                 "1.6L Turbo Hybrid",
		"engine_displacement":         "1598 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Hybrid",
		"engine_max_power":            "180 hp @ 5500 rpm",
		"engine_max_torque":           "300 Nm @ 3000 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "77.0 x 85.8 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Dual Clutch",
		"performance_top_speed":       "205 km/h",
		"performance_acceleration":    "8.3 seconds",
		"performance_mileage_arai":    "18 km/L",
		"performance_mileage_city":    "14 km/L",
		"performance_mileage_highway": "22 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4670 mm",
		"dimensions_width":            "1844 mm",
		"dimensions_height":           "1640 mm",
		"dimensions_wheelbase":        "2840 mm",
		"dimensions_front_tread":      "1608 mm",
		"dimensions_rear_tread":       "1617 mm",
		"dimensions_ground_clearance": "236 mm",
		"dimensions_boot_space":       "780 L",
		"dimensions_fuel_tank":        "56 L",
		"dimensions_turning_radius":   "5.2 m",
		"weight_kerb_weight":          "1490 kg",
		"weight_gross_weight":         "2130 kg",
		"capacity_seating_capacity":   "7",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "225/55 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "225/55 R18",
		"exterior_colors":             "Pearl White, Black, Gray, White, Silver, Blue, Red, Brown, Green",
		"exterior_headlights":         "Full LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "8-way Electric",
		"interior_ac":                 "Dual Zone Auto AC",
		"infotainment_screen_size":    "10-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "6 Speakers",
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
		"features_cruise_control":     "Adaptive Cruise Control",
		"features_sunroof":            "Panoramic Sunroof",
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
				LanguageCode:    "bn",
				Value:           p5s.getBanglaTranslations()[value],
				Status:          1,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Peugeot 5008")
	return nil
}
