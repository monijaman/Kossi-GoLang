package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// Peugeot208Seeder seeds specifications for Peugeot 208
type Peugeot208Seeder struct {
	BaseSeeder
}

// NewPeugeot208Seeder creates a new Peugeot 208 seeder
func NewPeugeot208Seeder() *Peugeot208Seeder {
	return &Peugeot208Seeder{
		BaseSeeder: BaseSeeder{name: "Peugeot 208 Specifications"},
	}
}

func (p2s *Peugeot208Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.2L Turbo":                    "১.২ লিটার টার্বো",
		"Subcompact Hatchback":          "সাবকমপ্যাক্ট হ্যাচব্যাক",
		"2022":                          "২০২২",
		"Hatchback":                     "হ্যাচব্যাক",
		"Nera Black":                    "নেরা ব্ল্যাক",
		"Inline":                        "ইনলাইন",
		"1199":                          "১১৯৯ সিসি",
		"3":                             "৩ (সিলিন্ডার)",
		"100 hp":                        "১০০ হর্স পাওয়ার",
		"100 hp @ 5500 rpm":             "৫৫০০ আরপিএম এ ১০০ হর্স পাওয়ার",
		"205 Nm @ 1750 rpm":             "১৭৫০ আরপিএম এ ২০৫ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"10.8 seconds":                  "১০.৮ সেকেন্ড",
		"180 km/h":                      "১৮০ কিমি/ঘণ্টা",
		"19 km/L":                       "১৯ কিমি/লিটার",
		"15 km/L":                       "১৫ কিমি/লিটার",
		"23 km/L":                       "২৩ কিমি/লিটার",
		"Manual (Transmission)":         "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":                "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":             "ফ্রন্ট-হুইল ড্রাইভ",
		"Single Plate Dry Clutch":       "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"3973 mm":                       "৩৯৭৩ মিমি",
		"1745 mm":                       "১৭৪৫ মিমি",
		"1430 mm":                       "১৪৩০ মিমি",
		"2540 mm":                       "২৫৪০ মিমি",
		"5":                             "৫ (সিট)",
		"5 Doors":                       "৫ ডোর",
		"5 Seats":                       "৫ সিট",
		"Black":                         "ব্ল্যাক",
		"Gray":                          "গ্রে",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Orange":                        "অরেঞ্জ",
		"LED":                           "এলইডি",
		"Halogen":                       "হ্যালোজেন",
		"LED Headlights":                "এলইডি হেডলাইট",
		"Halogen Headlights":            "হ্যালোজেন হেডলাইট",
		"LED Daytime Running Lights":    "এলইডি ডে টাইম রানিং লাইট",
		"No Roof Rails":                 "রুফ রেল নেই",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"15-inch Alloy Wheels":          "১৫-ইঞ্চি অ্যালয় হুইল",
		"16-inch Alloy Wheels":          "১৬-ইঞ্চি অ্যালয় হুইল",
		"195/55 R16":                    "১৯৫/৫৫ আর১৬",
		"Manual Adjustment":             "ম্যানুয়াল অ্যাডজাস্টমেন্ট",
		"Electric Adjustment":           "ইলেকট্রিক অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"6-way Manual":                  "৬-ওয়ে ম্যানুয়াল",
		"Fabric":                        "ফ্যাব্রিক",
		"Leather":                       "লেদার",
		"Fabric Seats":                  "ফ্যাব্রিক সিট",
		"Leather Seats":                 "লেদার সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Manual AC":                     "ম্যানুয়াল এসি",
		"Auto AC":                       "অটো এসি",
		"Single Zone":                   "সিঙ্গেল জোন",
		"7-inch Touchscreen":            "৭-ইঞ্চি টাচস্ক্রিন",
		"10-inch Touchscreen":           "১০-ইঞ্চি টাচস্ক্রিন",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"4 Speakers":                    "৪ স্পিকার",
		"6 Speakers":                    "৬ স্পিকার",
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
		"Rear Parking Camera":           "রিয়ার পার্কিং ক্যামেরা",
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
		"Tilt Steering":                 "টিল্ট স্টিয়ারিং",
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
		"285 L":                         "২৮৫ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"44 L":                          "৪৪ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"165 mm":                        "১৬৫ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"980 kg":                        "৯৮০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1480 kg":                       "১৪৮০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.4 m":                         "৫.৪ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"180 km/h":                      "১৮০ কিমি/ঘণ্টা",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"10.8 seconds":                  "১০.৮ সেকেন্ড",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"1.2L Turbo":                    "১.২লি টার্বো",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1199 cc":                       "১১৯৯ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"100 hp @ 5500 rpm":             "১০০ হর্স পাওয়ার @ ৫৫০০ আরপিএম",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"205 Nm @ 1750 rpm":             "২০৫ এনএম @ ১৭৫০ আরপিএম",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"3":                             "৩",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"4":                             "৪",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"75.0 x 90.5 mm":                "৭৫.০ x ৯০.৫ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"10.5:1":                        "১০.৫:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Manual":                        "ম্যানুয়াল",
		"Gear Box":                      "গিয়ার বক্স",
		"5-Speed":                       "৫-স্পিড",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"FWD":                           "এফডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"Single Plate Dry Clutch":       "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"19 km/L":                       "১৯ কিমি/লিটার",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"15 km/L":                       "১৫ কিমি/লিটার",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"23 km/L":                       "২৩ কিমি/লিটার",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"3973 mm":                       "৩৯৭৩ মিমি",
		"Width":                         "প্রস্থ",
		"1745 mm":                       "১৭৪৫ মিমি",
		"Height":                        "উচ্চতা",
		"1430 mm":                       "১৪৩০ মিমি",
		"Wheelbase":                     "হুইলবেস",
		"2540 mm":                       "২৫৪০ মিমি",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1500 mm":                       "১৫০০ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1500 mm":                       "১৫০০ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"5":                             "৫",
		"Door Count":                    "ডোর কাউন্ট",
		"5":                             "৫",
		"Boot Space":                    "বুট স্পেস",
		"285 L":                         "২৮৫ লিটার",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"44 L":                          "৪৪ লিটার",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"165 mm":                        "১৬৫ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"980 kg":                        "৯৮০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1480 kg":                       "১৪৮০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.4 m":                         "৫.৪ মিটার",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"MacPherson Strut":              "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Torsion Beam":                  "টর্শন বিম",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Drum":                          "ড্রাম",
		"Tyre Size":                     "টায়ার সাইজ",
		"195/55 R16":                    "১৯৫/৫৫ আর১৬",
		"Wheel Size":                    "হুইল সাইজ",
		"16 inches":                     "১৬ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
		"195/55 R16":                    "১৯৫/৫৫ আর১৬",
	}
}

func (p2s *Peugeot208Seeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Peugeot 208").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:        "Peugeot 208",
				Brand:       "Peugeot",
				Category:    "Hatchback",
				Subcategory: "Subcompact Hatchback",
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
		"engine_type":                 "1.2L Turbo",
		"engine_displacement":         "1199 cc",
		"engine_cylinders":            "3",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "100 hp @ 5500 rpm",
		"engine_max_torque":           "205 Nm @ 1750 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "75.0 x 90.5 mm",
		"engine_compression_ratio":    "10.5:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Manual",
		"transmission_gearbox":        "5-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Single Plate Dry Clutch",
		"performance_top_speed":       "180 km/h",
		"performance_acceleration":    "10.8 seconds",
		"performance_mileage_arai":    "19 km/L",
		"performance_mileage_city":    "15 km/L",
		"performance_mileage_highway": "23 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "3973 mm",
		"dimensions_width":            "1745 mm",
		"dimensions_height":           "1430 mm",
		"dimensions_wheelbase":        "2540 mm",
		"dimensions_front_tread":      "1500 mm",
		"dimensions_rear_tread":       "1500 mm",
		"dimensions_ground_clearance": "165 mm",
		"dimensions_boot_space":       "285 L",
		"dimensions_fuel_tank":        "44 L",
		"dimensions_turning_radius":   "5.4 m",
		"weight_kerb_weight":          "980 kg",
		"weight_gross_weight":         "1480 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Torsion Beam",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "195/55 R16",
		"tyres_wheel_size":            "16 inches",
		"tyres_spare_size":            "195/55 R16",
		"exterior_colors":             "Nera Black, Black, Gray, White, Silver, Blue, Red, Orange",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "No",
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
		"features_push_start":         "Yes",
		"features_cruise_control":     "No",
		"features_sunroof":            "No",
		"year":                        "2022",
		"category":                    "Hatchback",
		"subcategory":                 "Subcompact Hatchback",
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
				Value:           p2s.getBanglaTranslations()[value],
				Status:          1,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Peugeot 208")
	return nil
}
