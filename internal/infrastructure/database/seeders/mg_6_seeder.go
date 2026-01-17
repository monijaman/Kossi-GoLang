package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MG6Seeder seeds specifications for MG 6
type MG6Seeder struct {
	BaseSeeder
}

// NewMG6Seeder creates a new MG 6 seeder
func NewMG6Seeder() *MG6Seeder {
	return &MG6Seeder{
		BaseSeeder: BaseSeeder{name: "MG 6 Specifications"},
	}
}

func (ms *MG6Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.8L Turbo":                    "১.৮ লিটার টার্বো",
		"Mid-size Sedan":                "মিড-সাইজ সেডান",
		"2018":                          "২০১৮",
		"Sedan":                         "সেডান",
		"Arctic White":                  "আর্টিক হোয়াইট",
		"Inline":                        "ইনলাইন",
		"1796":                          "১৭৯৬ সিসি",
		"4":                             "৪ (সিলিন্ডার)",
		"166 hp":                        "১৬৬ হর্স পাওয়ার",
		"166 hp @ 5500 rpm":             "৫৫০০ আরপিএম এ ১৬৬ হর্স পাওয়ার",
		"215 Nm @ 2500-4500 rpm":        "২৫০০-৪৫০০ আরপিএম এ ২১৫ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"8.4 seconds":                   "৮.৪ সেকেন্ড",
		"210 km/h":                      "২১০ কিমি/ঘণ্টা",
		"13 km/L":                       "১৩ কিমি/লিটার",
		"10 km/L":                       "১০ কিমি/লিটার",
		"16 km/L":                       "১৬ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"7-Speed DCT":                   "৭-স্পিড ডিসিটি",
		"Front-Wheel Drive":             "ফ্রন্ট-হুইল ড্রাইভ",
		"Dual Clutch":                   "ডুয়াল ক্লাচ",
		"4695 mm":                       "৪৬৯৫ মিমি",
		"1848 mm":                       "১৮৪৮ মিমি",
		"1450 mm":                       "১৪৫০ মিমি",
		"2715 mm":                       "২৭১৫ মিমি",
		"5":                             "৫ (সিট)",
		"4 Doors":                       "৪ ডোর",
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
		"16-inch Alloy Wheels":          "১৬-ইঞ্চি অ্যালয় হুইল",
		"17-inch Alloy Wheels":          "১৭-ইঞ্চি অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"205/55 R16":                    "২০৫/৫৫ আর১৬",
		"225/45 R17":                    "২২৫/৪৫ আর১৭",
		"225/40 R18":                    "২২৫/৪০ আর১৮",
		"Manual Adjustment":             "ম্যানুয়াল অ্যাডজাস্টমেন্ট",
		"Electric Adjustment":           "ইলেকট্রিক অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"6-way Manual":                  "৬-ওয়ে ম্যানুয়াল",
		"8-way Electric":                "৮-ওয়ে ইলেকট্রিক",
		"10-way Electric":               "১০-ওয়ে ইলেকট্রিক",
		"Fabric":                        "ফ্যাব্রিক",
		"Leather":                       "লেদার",
		"Fabric Seats":                  "ফ্যাব্রিক সিট",
		"Leather Seats":                 "লেদার সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Manual AC":                     "ম্যানুয়াল এসি",
		"Auto AC":                       "অটো এসি",
		"Single Zone":                   "সিঙ্গেল জোন",
		"Dual Zone":                     "ডুয়াল জোন",
		"Rear AC Vents":                 "রিয়ার এসি ভেন্ট",
		"8-inch Touchscreen":            "৮-ইঞ্চি টাচস্ক্রিন",
		"10-inch Touchscreen":           "১০-ইঞ্চি টাচস্ক্রিন",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"4 Speakers":                    "৪ স্পিকার",
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
		"Front Parking Sensors":         "ফ্রন্ট পার্কিং সেন্সর",
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
		"Sunroof":                       "সানরুফ",
		"Panoramic Sunroof":             "প্যানোরামিক সানরুফ",
		"Chrome Accents":                "ক্রোম অ্যাকসেন্ট",
		"Body Colored Bumpers":          "বডি কালার্ড বাম্পার",
		"Rear Spoiler":                  "রিয়ার স্পয়লার",
		"LED Tail Lights":               "এলইডি টেইল লাইট",
		"Fog Lights":                    "ফগ লাইট",
		"LED Fog Lights":                "এলইডি ফগ লাইট",
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
		"Cigarette Lighter":             "সিগারেট লাইটার",
		"Boot Lamp":                     "বুট ল্যাম্প",
		"Vanity Mirrors":                "ভ্যানিটি মিরর",
		"Adjustable Steering":           "অ্যাডজাস্টেবল স্টিয়ারিং",
		"Tilt Steering":                 "টিল্ট স্টিয়ারিং",
		"Telescopic Steering":           "টেলিস্কোপিক স্টিয়ারিং",
		"Height Adjustable Steering":    "হাইট অ্যাডজাস্টেবল স্টিয়ারিং",
		"Leather Steering Wheel":        "লেদার স্টিয়ারিং হুইল",
		"Multi-function Steering":       "মাল্টি-ফাংশন স্টিয়ারিং",
		"Audio Controls":                "অডিও কন্ট্রোল",
		"Voice Commands":                "ভয়েস কমান্ড",
		"Cruise Control":                "ক্রুজ কন্ট্রোল",
		"Wireless Charger":              "ওয়্যারলেস চার্জার",
		"Ambient Lighting":              "অ্যাম্বিয়েন্ট লাইটিং",
		"Reading Lamps":                 "রিডিং ল্যাম্প",
		"Glove Box Lamp":                "গ্লোভ বক্স ল্যাম্প",
		"Trunk Lamp":                    "ট্রাঙ্ক ল্যাম্প",
		"Welcome Function":              "ওয়েলকাম ফাংশন",
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
		"464 L":                         "৪৬৪ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"62 L":                          "৬২ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"148 mm":                        "১৪৮ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1395 kg":                       "১৩৯৫ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1895 kg":                       "১৮৯৫ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.9 m":                         "৫.৯ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"210 km/h":                      "২১০ কিমি/ঘণ্টা",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"8.4 seconds":                   "৮.৪ সেকেন্ড",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"1.8L Turbo":                    "১.৮লি টার্বো",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1796 cc":                       "১৭৯৬ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"166 hp @ 5500 rpm":             "১৬৬ হর্স পাওয়ার @ ৫৫০০ আরপিএম",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"215 Nm @ 2500-4500 rpm":        "২১৫ এনএম @ ২৫০০-৪৫০০ আরপিএম",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"4":                             "৪",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"4":                             "৪",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"80.0 x 89.6 mm":                "৮০.০ x ৮৯.৬ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"9.3:1":                         "৯.৩:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
		"7-Speed DCT":                   "৭-স্পিড ডিসিটি",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"FWD":                           "এফডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"Dual Clutch":                   "ডুয়াল ক্লাচ",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"13 km/L":                       "১৩ কিমি/লিটার",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"10 km/L":                       "১০ কিমি/লিটার",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"16 km/L":                       "১৬ কিমি/লিটার",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"4695 mm":                       "৪৬৯৫ মিমি",
		"Width":                         "প্রস্থ",
		"1848 mm":                       "১৮৪৮ মিমি",
		"Height":                        "উচ্চতা",
		"1450 mm":                       "১৪৫০ মিমি",
		"Wheelbase":                     "হুইলবেস",
		"2715 mm":                       "২৭১৫ মিমি",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1570 mm":                       "১৫৭০ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1573 mm":                       "১৫৭৩ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"5":                             "৫",
		"Door Count":                    "ডোর কাউন্ট",
		"4":                             "৪",
		"Boot Space":                    "বুট স্পেস",
		"464 L":                         "৪৬৪ লিটার",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"62 L":                          "৬২ লিটার",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"148 mm":                        "১৪৮ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1395 kg":                       "১৩৯৫ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1895 kg":                       "১৮৯৫ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.9 m":                         "৫.৯ মিটার",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"MacPherson Strut":              "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Multi-link":                    "মাল্টি-লিঙ্ক",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Tyre Size":                     "টায়ার সাইজ",
		"225/45 R17":                    "২২৫/৪৫ আর১৭",
		"Wheel Size":                    "হুইল সাইজ",
		"17 inches":                     "১৭ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
		"225/45 R17":                    "২২৫/৪৫ আর১৭",
	}
}

func (ms *MG6Seeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "MG 6").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:        "MG 6",
				Brand:       "MG",
				Category:    "Sedan",
				Subcategory: "Mid-size Sedan",
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
		"engine_type":                 "1.8L Turbo",
		"engine_displacement":         "1796 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "166 hp @ 5500 rpm",
		"engine_max_torque":           "215 Nm @ 2500-4500 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "80.0 x 89.6 mm",
		"engine_compression_ratio":    "9.3:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "7-Speed DCT",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Dual Clutch",
		"performance_top_speed":       "210 km/h",
		"performance_acceleration":    "8.4 seconds",
		"performance_mileage_arai":    "13 km/L",
		"performance_mileage_city":    "10 km/L",
		"performance_mileage_highway": "16 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4695 mm",
		"dimensions_width":            "1848 mm",
		"dimensions_height":           "1450 mm",
		"dimensions_wheelbase":        "2715 mm",
		"dimensions_front_tread":      "1570 mm",
		"dimensions_rear_tread":       "1573 mm",
		"dimensions_ground_clearance": "148 mm",
		"dimensions_boot_space":       "464 L",
		"dimensions_fuel_tank":        "62 L",
		"dimensions_turning_radius":   "5.9 m",
		"weight_kerb_weight":          "1395 kg",
		"weight_gross_weight":         "1895 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "225/45 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "225/45 R17",
		"exterior_colors":             "Arctic White, Black, Gray, Silver, Blue, Red, Orange",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "No",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "8-way Electric",
		"interior_ac":                 "Dual Zone Auto AC",
		"infotainment_screen_size":    "10-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "6 Speakers",
		"safety_airbags":              "Front, Side, Curtain Airbags",
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
		"features_sunroof":            "Panoramic Sunroof",
		"year":                        "2018",
		"category":                    "Sedan",
		"subcategory":                 "Mid-size Sedan",
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
				Value:           ms.getBanglaTranslations()[value],
				Status:          1,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for MG 6")
	return nil
}
