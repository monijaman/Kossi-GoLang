package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// PorscheMacanSeeder seeds specifications for Porsche Macan
type PorscheMacanSeeder struct {
	BaseSeeder
}

// NewPorscheMacanSeeder creates a new Porsche Macan seeder
func NewPorscheMacanSeeder() *PorscheMacanSeeder {
	return &PorscheMacanSeeder{
		BaseSeeder: BaseSeeder{name: "Porsche Macan Specifications"},
	}
}

func (pms *PorscheMacanSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Turbo I4":                 "২.০ লিটার টার্বো আই৪",
		"Compact Luxury SUV":            "কমপ্যাক্ট লাক্সারি এসইউভি",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Jet Black":                     "জেট ব্ল্যাক",
		"Inline":                        "ইনলাইন",
		"1984":                          "১৯৮৪ সিসি",
		"4":                             "৪ (সিলিন্ডার)",
		"265 hp":                        "২৬৫ হর্স পাওয়ার",
		"265 hp @ 5000 rpm":             "৫০০০ আরপিএম এ ২৬৫ হর্স পাওয়ার",
		"400 Nm @ 1800 rpm":             "১৮০০ আরপিএম এ ৪০০ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"6.7 seconds":                   "৬.৭ সেকেন্ড",
		"232 km/h":                      "২৩২ কিমি/ঘণ্টা",
		"12 km/L":                       "১২ কিমি/লিটার",
		"9 km/L":                        "৯ কিমি/লিটার",
		"15 km/L":                       "১৫ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"7-Speed Dual Clutch":           "৭-স্পিড ডুয়াল ক্লাচ",
		"All-Wheel Drive":               "অল-হুইল ড্রাইভ",
		"PDK":                           "পিডিকে",
		"4696 mm":                       "৪৬৯৬ মিমি",
		"1923 mm":                       "১৯২৩ মিমি",
		"1621 mm":                       "১৬২১ মিমি",
		"2807 mm":                       "২৮০৭ মিমি",
		"5":                             "৫ (সিট)",
		"5 Doors":                       "৫ ডোর",
		"5 Seats":                       "৫ সিট",
		"Black":                         "ব্ল্যাক",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Gray":                          "গ্রে",
		"Green":                         "গ্রিন",
		"Orange":                        "অরেঞ্জ",
		"LED":                           "এলইডি",
		"Full LED":                      "ফুল এলইডি",
		"LED Headlights":                "এলইডি হেডলাইট",
		"Porsche Dynamic Light System":  "পর্শে ডায়নামিক লাইট সিস্টেম",
		"Roof Rails":                    "রুফ রেল",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"19-inch Alloy Wheels":          "১৯-ইঞ্চি অ্যালয় হুইল",
		"20-inch Alloy Wheels":          "২০-ইঞ্চি অ্যালয় হুইল",
		"235/60 R18":                    "২৩৫/৬০ আর১৮",
		"Electric Adjustment":           "ইলেকট্রিক অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"8-way Electric":                "৮-ওয়ে ইলেকট্রিক",
		"Leather":                       "লেদার",
		"Premium Leather":               "প্রিমিয়াম লেদার",
		"Alcantara":                     "অ্যালকান্টারা",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Auto AC":                       "অটো এসি",
		"2-Zone":                        "২-জোন",
		"10.9-inch Touchscreen":         "১০.৯-ইঞ্চি টাচস্ক্রিন",
		"PCM":                           "পিসিএম",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"WiFi":                          "ওয়াইফাই",
		"8 Speakers":                    "৮ স্পিকার",
		"Bose Surround Sound":           "বোজ সারাউন্ড সাউন্ড",
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
		"10.9-inch Digital Cluster":     "১০.৯-ইঞ্চি ডিজিটাল ক্লাস্টার",
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
		"Heated Steering Wheel":         "হিটেড স্টিয়ারিং হুইল",
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
		"488 L":                         "৪৮৮ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"65 L":                          "৬৫ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"188 mm":                        "১৮৮ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1770 kg":                       "১৭৭০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2400 kg":                       "২৪০০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.9 m":                         "৫.৯ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"232 km/h":                      "২৩২ কিমি/ঘণ্টা",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"6.7 seconds":                   "৬.৭ সেকেন্ড",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"2.0L Turbo I4":                 "২.০লি টার্বো আই৪",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1984 cc":                       "১৯৮৪ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"265 hp @ 5000 rpm":             "২৬৫ হর্স পাওয়ার @ ৫০০০ আরপিএম",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"400 Nm @ 1800 rpm":             "৪০০ এনএম @ ১৮০০ আরপিএম",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"4":                             "৪",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"4":                             "৪",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"82.5 x 92.8 mm":                "৮২.৫ x ৯২.৮ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"9.6:1":                         "৯.৬:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
		"7-Speed":                       "৭-স্পিড",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"AWD":                           "এডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"PDK":                           "পিডিকে",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"12 km/L":                       "১২ কিমি/লিটার",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"9 km/L":                        "৯ কিমি/লিটার",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"15 km/L":                       "১৫ কিমি/লিটার",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"4696 mm":                       "৪৬৯৬ মিমি",
		"Width":                         "প্রস্থ",
		"1923 mm":                       "১৯২৩ মিমি",
		"Height":                        "উচ্চতা",
		"1621 mm":                       "১৬২১ মিমি",
		"Wheelbase":                     "হুইলবেস",
		"2807 mm":                       "২৮০৭ মিমি",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1655 mm":                       "১৬৫৫ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1654 mm":                       "১৬৫৪ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"5":                             "৫",
		"Door Count":                    "ডোর কাউন্ট",
		"5":                             "৫",
		"Boot Space":                    "বুট স্পেস",
		"488 L":                         "৪৮৮ লিটার",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"65 L":                          "৬৫ লিটার",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"188 mm":                        "১৮৮ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1770 kg":                       "১৭৭০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"2400 kg":                       "২৪০০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.9 m":                         "৫.৯ মিটার",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"Double Wishbone":               "ডাবল উইশবোন",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Multi-link":                    "মাল্টি-লিঙ্ক",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Tyre Size":                     "টায়ার সাইজ",
		"235/60 R18":                    "২৩৫/৬০ আর১৮",
		"Wheel Size":                    "হুইল সাইজ",
		"18 inches":                     "১৮ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
		"235/60 R18":                    "২৩৫/৬০ আর১৮",
	}
}

func (pms *PorscheMacanSeeder) Seed(db *gorm.DB) error {
	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Porsche Macan").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:        "Porsche Macan",
				Brand:       "Porsche",
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
		"engine_displacement":         "1984 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "265 hp @ 5000 rpm",
		"engine_max_torque":           "400 Nm @ 1800 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "82.5 x 92.8 mm",
		"engine_compression_ratio":    "9.6:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "7-Speed",
		"transmission_drive_type":     "AWD",
		"transmission_clutch_type":    "PDK",
		"performance_top_speed":       "232 km/h",
		"performance_acceleration":    "6.7 seconds",
		"performance_mileage_arai":    "12 km/L",
		"performance_mileage_city":    "9 km/L",
		"performance_mileage_highway": "15 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4696 mm",
		"dimensions_width":            "1923 mm",
		"dimensions_height":           "1621 mm",
		"dimensions_wheelbase":        "2807 mm",
		"dimensions_front_tread":      "1655 mm",
		"dimensions_rear_tread":       "1654 mm",
		"dimensions_ground_clearance": "188 mm",
		"dimensions_boot_space":       "488 L",
		"dimensions_fuel_tank":        "65 L",
		"dimensions_turning_radius":   "5.9 m",
		"weight_kerb_weight":          "1770 kg",
		"weight_gross_weight":         "2400 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "Double Wishbone",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "235/60 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "235/60 R18",
		"exterior_colors":             "Jet Black, Black, White, Silver, Blue, Red, Gray, Green, Orange",
		"exterior_headlights":         "Full LED",
		"exterior_daytime_running":    "Porsche Dynamic Light System",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
		"interior_seats_material":     "Premium Leather",
		"interior_seats_adjustment":   "8-way Electric",
		"interior_ac":                 "2-Zone Auto AC",
		"infotainment_screen_size":    "10.9-inch Touchscreen",
		"infotainment_connectivity":   "PCM, Android Auto, Apple CarPlay, Bluetooth, WiFi",
		"infotainment_speakers":       "Bose Surround Sound",
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
				Value:           pms.getBanglaTranslations()[value],
				Status:          1,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Porsche Macan")
	return nil
}
