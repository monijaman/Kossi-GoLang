package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MahindraTharSeeder seeds specifications for Mahindra Thar
type MahindraTharSeeder struct {
	BaseSeeder
}

// NewMahindraTharSeeder creates a new Mahindra Thar seeder
func NewMahindraTharSeeder() *MahindraTharSeeder {
	return &MahindraTharSeeder{
		BaseSeeder: BaseSeeder{name: "Mahindra Thar Specifications"},
	}
}

func (mts *MahindraTharSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.2L mHawk Diesel":              "২.২ লিটার এমহক ডিজেল",
		"Off-road SUV":                   "অফ-রোড এসইউভি",
		"2023":                           "২০২৩",
		"SUV":                            "এসইউভি",
		"Red Rage":                       "রেড রেজ",
		"Inline":                         "ইনলাইন",
		"2179":                           "২১৭৯ সিসি",
		"4":                              "৪ (সিলিন্ডার)",
		"130 hp":                         "১৩০ হর্স পাওয়ার",
		"130 hp @ 3750 rpm":              "৩৭৫০ আরপিএম এ ১৩০ হর্স পাওয়ার",
		"300 Nm @ 1600-2800 rpm":         "১৬০০-২৮০০ আরপিএম এ ৩০০ এনএম",
		"Diesel":                         "ডিজেল",
		"Turbocharged":                   "টার্বোচার্জড",
		"14.3 seconds":                   "১৪.৩ সেকেন্ড",
		"155 km/h":                       "১৫৫ কিমি/ঘণ্টা",
		"13 km/L":                        "১৩ কিমি/লিটার",
		"10 km/L":                        "১০ কিমি/লিটার",
		"15 km/L":                        "১৫ কিমি/লিটার",
		"Manual":                         "ম্যানুয়াল",
		"Automatic":                      "অটোমেটিক",
		"6-Speed Manual":                 "৬-স্পিড ম্যানুয়াল",
		"6-Speed Automatic":              "৬-স্পিড অটোমেটিক",
		"Four-Wheel Drive":               "ফোর-হুইল ড্রাইভ",
		"All-Wheel Drive":                "অল-হুইল ড্রাইভ",
		"3985 mm":                        "৩৯৮৫ মিমি",
		"1820 mm":                        "১৮২০ মিমি",
		"1844 mm":                        "১৮৪৪ মিমি",
		"2450 mm":                        "২৪৫০ মিমি",
		"4":                              "৪ (সিট)",
		"3 Doors":                        "৩ ডোর",
		"4 Seats":                        "৪ সিট",
		"Black":                          "ব্ল্যাক",
		"Gray":                           "গ্রে",
		"White":                          "হোয়াইট",
		"Silver":                         "সিলভার",
		"Blue":                           "ব্লু",
		"Red":                            "রেড",
		"Brown":                          "ব্রাউন",
		"Orange":                         "অরেঞ্জ",
		"Halogen":                        "হ্যালোজেন",
		"Halogen Headlights":             "হ্যালোজেন হেডলাইট",
		"Halogen Daytime Running Lights": "হ্যালোজেন ডে টাইম রানিং লাইট",
		"Roof Rails":                     "রুফ রেল",
		"Steel Wheels":                   "স্টিল হুইল",
		"Alloy Wheels":                   "অ্যালয় হুইল",
		"16-inch Alloy Wheels":           "১৬-ইঞ্চি অ্যালয় হুইল",
		"255/65 R18":                     "২৫৫/৬৫ আর১৮",
		"Manual Adjustment":              "ম্যানুয়াল অ্যাডজাস্টমেন্ট",
		"Height Adjustable":              "হাইট অ্যাডজাস্টেবল",
		"6-way Manual":                   "৬-ওয়ে ম্যানুয়াল",
		"Fabric":                         "ফ্যাব্রিক",
		"Fabric Seats":                   "ফ্যাব্রিক সিট",
		"50:50 Split":                    "৫০:৫০ স্প্লিট",
		"Folding Rear Seats":             "ফোল্ডিং রিয়ার সিট",
		"Manual AC":                      "ম্যানুয়াল এসি",
		"Auto AC":                        "অটো এসি",
		"Dual Zone":                      "ডুয়াল জোন",
		"7-inch Touchscreen":             "৭-ইঞ্চি টাচস্ক্রিন",
		"8-inch Touchscreen":             "৮-ইঞ্চি টাচস্ক্রিন",
		"Android Auto":                   "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                  "অ্যাপল কারপ্লে",
		"Bluetooth":                      "ব্লুটুথ",
		"USB":                            "ইউএসবি",
		"AUX":                            "অক্সিলারি",
		"4 Speakers":                     "৪ স্পিকার",
		"6 Speakers":                     "৬ স্পিকার",
		"ABS":                            "এবিএস",
		"EBD":                            "ইবিডি",
		"Brake Assist":                   "ব্রেক অ্যাসিস্ট",
		"ESP":                            "ইএসপি",
		"Traction Control":               "ট্র্যাকশন কন্ট্রোল",
		"Hill Hold Control":              "হিল হোল্ড কন্ট্রোল",
		"Front Airbags":                  "ফ্রন্ট এয়ারব্যাগ",
		"Side Airbags":                   "সাইড এয়ারব্যাগ",
		"Curtain Airbags":                "কার্টেন এয়ারব্যাগ",
		"Rear Parking Sensors":           "রিয়ার পার্কিং সেন্সর",
		"Rear Parking Camera":            "রিয়ার পার্কিং ক্যামেরা",
		"360-degree Camera":              "৩৬০-ডিগ্রি ক্যামেরা",
		"ISOFIX":                         "আইএসওফিক্স",
		"Child Safety Locks":             "চাইল্ড সেফটি লক",
		"Speed Sensing Door Locks":       "স্পিড সেন্সিং ডোর লক",
		"Central Locking":                "সেন্ট্রাল লকিং",
		"Keyless Entry":                  "কীলেস এন্ট্রি",
		"Push Button Start":              "পুশ বাটন স্টার্ট",
		"Engine Immobilizer":             "ইঞ্জিন ইমোবিলাইজার",
		"Anti-theft Alarm":               "অ্যান্টি-থেফট অ্যালার্ম",
		"Tyre Pressure Monitoring":       "টায়ার প্রেশার মনিটরিং",
		"Follow-me-home Headlamps":       "ফলো-মি-হোম হেডল্যাম্প",
		"Rain Sensing Wipers":            "রেইন সেন্সিং ওয়াইপার",
		"Auto Headlamps":                 "অটো হেডল্যাম্প",
		"Chrome Accents":                 "ক্রোম অ্যাকসেন্ট",
		"Body Colored Bumpers":           "বডি কালার্ড বাম্পার",
		"Rear Spoiler":                   "রিয়ার স্পয়লার",
		"Halogen Tail Lights":            "হ্যালোজেন টেইল লাইট",
		"Fog Lights":                     "ফগ লাইট",
		"Projector Headlamps":            "প্রজেক্টর হেডল্যাম্প",
		"DRLs":                           "ডিআরএল",
		"Halogen DRLs":                   "হ্যালোজেন ডিআরএল",
		"Digital Instrument Cluster":     "ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার",
		"Analog Instrument Cluster":      "অ্যানালগ ইন্সট্রুমেন্ট ক্লাস্টার",
		"Trip Computer":                  "ট্রিপ কম্পিউটার",
		"Tachometer":                     "ট্যাকোমিটার",
		"Gear Shift Indicator":           "গিয়ার শিফ্ট ইন্ডিকেটর",
		"Door Ajar Warning":              "ডোর অজার ওয়ার্নিং",
		"Low Fuel Warning":               "লো ফুয়েল ওয়ার্নিং",
		"12V Power Outlet":               "১২ভি পাওয়ার আউটলেট",
		"Boot Lamp":                      "বুট ল্যাম্প",
		"Vanity Mirrors":                 "ভ্যানিটি মিরর",
		"Adjustable Steering":            "অ্যাডজাস্টেবল স্টিয়ারিং",
		"Tilt & Telescopic":              "টিল্ট & টেলিস্কোপিক",
		"Height Adjustable Steering":     "হাইট অ্যাডজাস্টেবল স্টিয়ারিং",
		"Urethane Steering Wheel":        "ইউরেথেন স্টিয়ারিং হুইল",
		"Multi-function Steering":        "মাল্টি-ফাংশন স্টিয়ারিং",
		"Audio Controls":                 "অডিও কন্ট্রোল",
		"Cruise Control":                 "ক্রুজ কন্ট্রোল",
		"Reading Lamps":                  "রিডিং ল্যাম্প",
		"Glove Box Lamp":                 "গ্লোভ বক্স ল্যাম্প",
		"Trunk Lamp":                     "ট্রাঙ্ক ল্যাম্প",
		"ORVMs with Turn Indicators":     "ওআরভিএম উইথ টার্ন ইন্ডিকেটর",
		"Power Windows":                  "পাওয়ার উইন্ডো",
		"One Touch Up/Down":              "ওয়ান টাচ আপ/ডাউন",
		"Anti-pinch":                     "অ্যান্টি-পিঞ্চ",
		"Boot Release":                   "বুট রিলিজ",
		"Remote Boot Release":            "রিমোট বুট রিলিজ",
		"Rear Defogger":                  "রিয়ার ডিফগার",
		"Rear Wiper":                     "রিয়ার ওয়াইপার",
		"Heated ORVMs":                   "হিটেড ওআরভিএম",
		"Electrically Adjustable ORVMs":  "ইলেকট্রিক্যালি অ্যাডজাস্টেবল ওআরভিএম",
		"Boot Space":                     "বুট স্পেস",
		"300 L":                          "৩০০ লিটার",
		"Fuel Tank":                      "ফুয়েল ট্যাঙ্ক",
		"57 L":                           "৫৭ লিটার",
		"Ground Clearance":               "গ্রাউন্ড ক্লিয়ারেন্স",
		"226 mm":                         "২২৬ মিমি",
		"Kerb Weight":                    "কার্ব ওয়েট",
		"1752 kg":                        "১৭৫২ কেজি",
		"Gross Weight":                   "গ্রস ওয়েট",
		"2450 kg":                        "২৪৫০ কেজি",
		"Turning Radius":                 "টার্নিং রেডিয়াস",
		"5.25 m":                         "৫.২৫ মিটার",
		"Top Speed":                      "টপ স্পিড",
		"155 km/h":                       "১৫৫ কিমি/ঘণ্টা",
		"Acceleration 0-100 km/h":        "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"14.3 seconds":                   "১৪.৩ সেকেন্ড",
		"Engine Type":                    "ইঞ্জিন টাইপ",
		"2.2L mHawk Diesel":              "২.২লি এমহক ডিজেল",
		"Displacement":                   "ডিসপ্লেসমেন্ট",
		"2179 cc":                        "২১৭৯ সিসি",
		"Max Power":                      "ম্যাক্স পাওয়ার",
		"130 hp @ 3750 rpm":              "১৩০ হর্স পাওয়ার @ ৩৭৫০ আরপিএম",
		"Max Torque":                     "ম্যাক্স টর্ক",
		"300 Nm @ 1600-2800 rpm":         "৩০০ এনএম @ ১৬০০-২৮০০ আরপিএম",
		"No. of Cylinders":               "সিলিন্ডারের সংখ্যা",
		"4":                              "৪",
		"Valves per Cylinder":            "প্রতি সিলিন্ডার ভালভ",
		"4":                              "৪",
		"Fuel Supply System":             "ফুয়েল সাপ্লাই সিস্টেম",
		"CRDi":                           "সিআরডিআই",
		"Bore x Stroke":                  "বোর x স্ট্রোক",
		"85 x 96 mm":                     "৮৫ x ৯৬ মিমি",
		"Compression Ratio":              "কম্প্রেশন রেশিও",
		"16.5:1":                         "১৬.৫:১",
		"Turbo Charger":                  "টার্বো চার্জার",
		"Yes":                            "হ্যাঁ",
		"Super Charger":                  "সুপার চার্জার",
		"No":                             "না",
		"Transmission Type":              "ট্রান্সমিশন টাইপ",
		"Manual":                         "ম্যানুয়াল",
		"Gear Box":                       "গিয়ার বক্স",
		"6-Speed Manual":                 "৬-স্পিড ম্যানুয়াল",
		"Drive Type":                     "ড্রাইভ টাইপ",
		"4WD":                            "৪ডব্লিউডি",
		"Clutch Type":                    "ক্লাচ টাইপ",
		"Single Plate Dry Clutch":        "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"Mileage (ARAI)":                 "মাইলেজ (এআরএআই)",
		"13 km/L":                        "১৩ কিমি/লিটার",
		"Mileage (City)":                 "মাইলেজ (সিটি)",
		"10 km/L":                        "১০ কিমি/লিটার",
		"Mileage (Highway)":              "মাইলেজ (হাইওয়ে)",
		"15 km/L":                        "১৫ কিমি/লিটার",
		"Emission Norm Compliance":       "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                          "বিএস ভি",
		"Length":                         "দৈর্ঘ্য",
		"3985 mm":                        "৩৯৮৫ মিমি",
		"Width":                          "প্রস্থ",
		"1820 mm":                        "১৮২০ মিমি",
		"Height":                         "উচ্চতা",
		"1844 mm":                        "১৮৪৪ মিমি",
		"Wheelbase":                      "হুইলবেস",
		"2450 mm":                        "২৪৫০ মিমি",
		"Front Tread":                    "ফ্রন্ট ট্রেড",
		"1440 mm":                        "১৪৪০ মিমি",
		"Rear Tread":                     "রিয়ার ট্রেড",
		"1440 mm":                        "১৪৪০ মিমি",
		"Seating Capacity":               "সিটিং ক্যাপাসিটি",
		"4":                              "৪",
		"Door Count":                     "ডোর কাউন্ট",
		"3":                              "৩",
		"Boot Space":                     "বুট স্পেস",
		"300 L":                          "৩০০ লিটার",
		"Fuel Tank Capacity":             "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"57 L":                           "৫৭ লিটার",
		"Ground Clearance Unladen":       "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"226 mm":                         "২২৬ মিমি",
		"Kerb Weight":                    "কার্ব ওয়েট",
		"1752 kg":                        "১৭৫২ কেজি",
		"Gross Weight":                   "গ্রস ওয়েট",
		"2450 kg":                        "২৪৫০ কেজি",
		"Turning Radius":                 "টার্নিং রেডিয়াস",
		"5.25 m":                         "৫.২৫ মিটার",
		"Front Suspension":               "ফ্রন্ট সাসপেনশন",
		"Independent Double Wishbone":    "ইন্ডিপেন্ডেন্ট ডাবল উইশবোন",
		"Rear Suspension":                "রিয়ার সাসপেনশন",
		"Non-independent Multi-link":     "নন-ইন্ডিপেন্ডেন্ট মাল্টি-লিঙ্ক",
		"Front Brake Type":               "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                           "ডিস্ক",
		"Rear Brake Type":                "রিয়ার ব্রেক টাইপ",
		"Drum":                           "ড্রাম",
		"Tyre Size":                      "টায়ার সাইজ",
		"255/65 R18":                     "২৫৫/৬৫ আর১৮",
		"Wheel Size":                     "হুইল সাইজ",
		"18 inches":                      "১৮ ইঞ্চি",
		"Spare Tyre Size":                "স্পেয়ার টায়ার সাইজ",
		"255/65 R18":                     "২৫৫/৬৫ আর১৮",
	}
}

func (mts *MahindraTharSeeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Mahindra Thar").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:        "Mahindra Thar",
				Brand:       "Mahindra",
				Category:    "SUV",
				Subcategory: "Off-road SUV",
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
		"transmission_drive_type":     "4WD",
		"transmission_clutch_type":    "Single Plate Dry Clutch",
		"performance_top_speed":       "155 km/h",
		"performance_acceleration":    "14.3 seconds",
		"performance_mileage_arai":    "13 km/L",
		"performance_mileage_city":    "10 km/L",
		"performance_mileage_highway": "15 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "3985 mm",
		"dimensions_width":            "1820 mm",
		"dimensions_height":           "1844 mm",
		"dimensions_wheelbase":        "2450 mm",
		"dimensions_front_tread":      "1440 mm",
		"dimensions_rear_tread":       "1440 mm",
		"dimensions_ground_clearance": "226 mm",
		"dimensions_boot_space":       "300 L",
		"dimensions_fuel_tank":        "57 L",
		"dimensions_turning_radius":   "5.25 m",
		"weight_kerb_weight":          "1752 kg",
		"weight_gross_weight":         "2450 kg",
		"capacity_seating_capacity":   "4",
		"capacity_doors":              "3",
		"suspension_front":            "Independent Double Wishbone",
		"suspension_rear":             "Non-independent Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Drum",
		"tyres_size":                  "255/65 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "255/65 R18",
		"exterior_colors":             "Red Rage, Black, Gray, White, Silver, Blue, Red, Brown, Orange",
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
		"subcategory":                 "Off-road SUV",
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
				Value:           mts.getBanglaTranslations()[value],
				Status:          1,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Mahindra Thar")
	return nil
}
