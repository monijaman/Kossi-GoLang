package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LexusNXSeeder seeds specifications for Lexus NX
type LexusNXSeeder struct {
	BaseSeeder
}

// NewLexusNXSeeder creates a new Lexus NX seeder
func NewLexusNXSeeder() *LexusNXSeeder {
	return &LexusNXSeeder{
		BaseSeeder: BaseSeeder{name: "Lexus NX Specifications"},
	}
}

func (lnxs *LexusNXSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":                 "২.০ লিটার টার্বো আই৪",
		"Compact Luxury SUV":            "কমপ্যাক্ট লাক্সারি এসইউভি",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Radiant Red":                   "রেডিয়েন্ট রেড",
		"Inline":                        "ইনলাইন",
		"1998":                          "১৯৯৮ সিসি",
		"4":                             "৪ (সিলিন্ডার)",
		"275 hp":                        "২৭৫ হর্স পাওয়ার",
		"275 hp @ 6000 rpm":             "৬০০০ আরপিএম এ ২৭৫ হর্স পাওয়ার",
		"380 Nm @ 1650 rpm":             "১৬৫০ আরপিএম এ ৩৮০ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"6.3 seconds":                   "৬.৩ সেকেন্ড",
		"200 km/h":                      "২০০ কিমি/ঘণ্টা",
		"13 km/L":                       "১৩ কিমি/লিটার",
		"9 km/L":                        "৯ কিমি/লিটার",
		"17 km/L":                       "১৭ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"CVT":                           "সিভিটি",
		"All-Wheel Drive":               "অল-হুইল ড্রাইভ",
		"E-Four":                        "ই-ফোর",
		"4640 mm":                       "৪৬৪০ মিমি",
		"1845 mm":                       "১৮৪৫ মিমি",
		"1645 mm":                       "১৬৪৫ মিমি",
		"2660 mm":                       "২৬৬০ মিমি",
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
		"Roof Rails":                    "রুফ রেল",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"17-inch Alloy Wheels":          "১৭-ইঞ্চি অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"225/65 R17":                    "২২৫/৬৫ আর১৭",
		"Power Adjustment":              "পাওয়ার অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"8-way Power":                   "৮-ওয়ে পাওয়ার",
		"Leather":                       "লেদার",
		"Leather Seats":                 "লেদার সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Auto AC":                       "অটো এসি",
		"Dual Zone":                     "ডুয়াল জোন",
		"10.3-inch Touchscreen":         "১০.৩-ইঞ্চি টাচস্ক্রিন",
		"Lexus Link":                    "লেক্সাস লিঙ্ক",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"10 Speakers":                   "১০ স্পিকার",
		"14 Speakers":                   "১৪ স্পিকার",
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
		"525 L":                         "৫২৫ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"60 L":                          "৬০ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"185 mm":                        "১৮৫ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1780 kg":                       "১৭৮০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2260 kg":                       "২২৬০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.8 m":                         "৫.৮ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"200 km/h":                      "২০০ কিমি/ঘণ্টা",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"6.3 seconds":                   "৬.৩ সেকেন্ড",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"2.0L Turbo I4":                 "২.০লি টার্বো আই৪",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1998 cc":                       "১৯৯৮ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"275 hp @ 6000 rpm":             "২৭৫ হর্স পাওয়ার @ ৬০০০ আরপিএম",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"380 Nm @ 1650 rpm":             "৩৮০ এনএম @ ১৬৫০ আরপিএম",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"4":                             "৪",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"4":                             "৪",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"86.0 x 86.0 mm":                "৮৬.০ x ৮৬.০ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"10.0:1":                        "১০.০:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
		"CVT":                           "সিভিটি",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"AWD":                           "এডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"E-Four":                        "ই-ফোর",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"13 km/L":                       "১৩ কিমি/লিটার",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"9 km/L":                        "৯ কিমি/লিটার",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"17 km/L":                       "১৭ কিমি/লিটার",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"4640 mm":                       "৪৬৪০ মিমি",
		"Width":                         "প্রস্থ",
		"1845 mm":                       "১৮৪৫ মিমি",
		"Height":                        "উচ্চতা",
		"1645 mm":                       "১৬৪৫ মিমি",
		"Wheelbase":                     "হুইলবেস",
		"2660 mm":                       "২৬৬০ মিমি",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1580 mm":                       "১৫৮০ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1585 mm":                       "১৫৮৫ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"5":                             "৫",
		"Door Count":                    "ডোর কাউন্ট",
		"4":                             "৪",
		"Boot Space":                    "বুট স্পেস",
		"525 L":                         "৫২৫ লিটার",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"60 L":                          "৬০ লিটার",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"185 mm":                        "১৮৫ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1780 kg":                       "১৭৮০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2260 kg":                       "২২৬০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.8 m":                         "৫.৮ মিটার",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"MacPherson Strut":              "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Double Wishbone":               "ডাবল উইশবোন",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Tyre Size":                     "টায়ার সাইজ",
		"225/65 R17":                    "২২৫/৬৫ আর১৭",
		"Wheel Size":                    "হুইল সাইজ",
		"17 inches":                     "১৭ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
		"225/65 R17":                    "২২৫/৬৫ আর১৭",
	}
}

func (lnxs *LexusNXSeeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Lexus NX").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:        "Lexus NX",
				Brand:       "Lexus",
				Category:    "SUV",
				Subcategory: "Compact Luxury SUV",
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
		"engine_type":                 "2.0L Turbo I4",
		"engine_displacement":         "1998 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "275 hp @ 6000 rpm",
		"engine_max_torque":           "380 Nm @ 1650 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "86.0 x 86.0 mm",
		"engine_compression_ratio":    "10.0:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "CVT",
		"transmission_drive_type":     "AWD",
		"transmission_clutch_type":    "E-Four",
		"performance_top_speed":       "200 km/h",
		"performance_acceleration":    "6.3 seconds",
		"performance_mileage_arai":    "13 km/L",
		"performance_mileage_city":    "9 km/L",
		"performance_mileage_highway": "17 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4640 mm",
		"dimensions_width":            "1845 mm",
		"dimensions_height":           "1645 mm",
		"dimensions_wheelbase":        "2660 mm",
		"dimensions_front_tread":      "1580 mm",
		"dimensions_rear_tread":       "1585 mm",
		"dimensions_ground_clearance": "185 mm",
		"dimensions_boot_space":       "525 L",
		"dimensions_fuel_tank":        "60 L",
		"dimensions_turning_radius":   "5.8 m",
		"weight_kerb_weight":          "1780 kg",
		"weight_gross_weight":         "2260 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Double Wishbone",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "225/65 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "225/65 R17",
		"exterior_colors":             "Radiant Red, Black, Gray, White, Silver, Blue, Red, Brown, Beige",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Leather",
		"interior_seats_adjustment":   "8-way Power",
		"interior_ac":                 "Dual Zone Auto AC",
		"infotainment_screen_size":    "10.3-inch Touchscreen",
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
		"category":                    "SUV",
		"subcategory":                 "Compact Luxury SUV",
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
				Value:           lnxs.getBanglaTranslations()[value],
				Status:          1,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Lexus NX")
	return nil
}
