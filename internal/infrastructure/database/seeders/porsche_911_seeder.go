package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// Porsche911Seeder seeds specifications for Porsche 911
type Porsche911Seeder struct {
	BaseSeeder
}

// NewPorsche911Seeder creates a new Porsche 911 seeder
func NewPorsche911Seeder() *Porsche911Seeder {
	return &Porsche911Seeder{
		BaseSeeder: BaseSeeder{name: "Porsche 911 Specifications"},
	}
}

func (p9s *Porsche911Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.0L Turbo Flat-6":             "৩.০ লিটার টার্বো ফ্ল্যাট-৬",
		"Sports Car":                    "স্পোর্টস কার",
		"2023":                          "২০২৩",
		"Coupe":                         "কুপে",
		"Guards Red":                    "গার্ডস রেড",
		"Flat":                          "ফ্ল্যাট",
		"2981":                          "২৯৮১ সিসি",
		"6":                             "৬ (সিলিন্ডার)",
		"450 hp":                        "৪৫০ হর্স পাওয়ার",
		"450 hp @ 6500 rpm":             "৬৫০০ আরপিএম এ ৪৫০ হর্স পাওয়ার",
		"530 Nm @ 2300 rpm":             "২৩০০ আরপিএম এ ৫৩০ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"3.5 seconds":                   "৩.৫ সেকেন্ড",
		"308 km/h":                      "৩০৮ কিমি/ঘণ্টা",
		"9 km/L":                        "৯ কিমি/লিটার",
		"6 km/L":                        "৬ কিমি/লিটার",
		"12 km/L":                       "১২ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"8-Speed Dual Clutch":           "৮-স্পিড ডুয়াল ক্লাচ",
		"All-Wheel Drive":               "অল-হুইল ড্রাইভ",
		"PDK":                           "পিডিকে",
		"4519 mm":                       "৪৫১৯ মিমি",
		"1852 mm":                       "১৮৫২ মিমি",
		"1300 mm":                       "১৩০০ মিমি",
		"2450 mm":                       "২৪৫০ মিমি",
		"4":                             "৪ (সিট)",
		"2 Doors":                       "২ ডোর",
		"4 Seats":                       "৪ সিট",
		"Black":                         "ব্ল্যাক",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Gray":                          "গ্রে",
		"Green":                         "গ্রিন",
		"Yellow":                        "ইয়েলো",
		"LED":                           "এলইডি",
		"Full LED":                      "ফুল এলইডি",
		"LED Headlights":                "এলইডি হেডলাইট",
		"Porsche Dynamic Light System":  "পর্শে ডায়নামিক লাইট সিস্টেম",
		"No Roof Rails":                 "রুফ রেল নেই",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"20-inch Alloy Wheels":          "২০-ইঞ্চি অ্যালয় হুইল",
		"21-inch Alloy Wheels":          "২১-ইঞ্চি অ্যালয় হুইল",
		"265/35 R20":                    "২৬৫/৩৫ আর২০",
		"Electric Adjustment":           "ইলেকট্রিক অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"18-way Electric":               "১৮-ওয়ে ইলেকট্রিক",
		"Leather":                       "লেদার",
		"Premium Leather":               "প্রিমিয়াম লেদার",
		"Alcantara":                     "অ্যালকান্টারা",
		"Fixed Seats":                   "ফিক্সড সিট",
		"No Rear Seats":                 "রিয়ার সিট নেই",
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
		"Knee Airbags":                  "নি এয়ারব্যাগ",
		"Rear Parking Sensors":          "রিয়ার পার্কিং সেন্সর",
		"360-degree Camera":             "৩৬০-ডিগ্রি ক্যামেরা",
		"Night Vision":                  "নাইট ভিশন",
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
		"Glass Roof":                    "গ্লাস রুফ",
		"Chrome Accents":                "ক্রোম অ্যাকসেন্ট",
		"Body Colored Bumpers":          "বডি কালার্ড বাম্পার",
		"Rear Spoiler":                  "রিয়ার স্পয়লার",
		"LED Tail Lights":               "এলইডি টেইল লাইট",
		"Fog Lights":                    "ফগ লাইট",
		"Projector Headlamps":           "প্রজেক্টর হেডল্যাম্প",
		"DRLs":                          "ডিআরএল",
		"LED DRLs":                      "এলইডি ডিআরএল",
		"Digital Instrument Cluster":    "ডিজিটাল ইন্সট্রুমেন্ট ক্লাস্টার",
		"12.6-inch Digital Cluster":     "১২.৬-ইঞ্চি ডিজিটাল ক্লাস্টার",
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
		"132 L":                         "১৩২ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"64 L":                          "৬৪ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"130 mm":                        "১৩০ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1505 kg":                       "১৫০৫ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1920 kg":                       "১৯২০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.6 m":                         "৫.৬ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"2981 cc":                       "২৯৮১ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"91.0 x 76.4 mm":                "৯১.০ x ৭৬.৪ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"10.2:1":                        "১০.২:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
		"8-Speed":                       "৮-স্পিড",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"AWD":                           "এডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"Width":                         "প্রস্থ",
		"Height":                        "উচ্চতা",
		"Wheelbase":                     "হুইলবেস",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1550 mm":                       "১৫৫০ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1545 mm":                       "১৫৪৫ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"Door Count":                    "ডোর কাউন্ট",
		"2":                             "২",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",

		"Front Suspension": "ফ্রন্ট সাসপেনশন",
		"Double Wishbone":  "ডাবল উইশবোন",
		"Rear Suspension":  "রিয়ার সাসপেনশন",
		"Multi-link":       "মাল্টি-লিঙ্ক",
		"Front Brake Type": "ফ্রন্ট ব্রেক টাইপ",
		"Disc":             "ডিস্ক",
		"Rear Brake Type":  "রিয়ার ব্রেক টাইপ",
		"Tyre Size":        "টায়ার সাইজ",
		"Wheel Size":       "হুইল সাইজ",
		"20 inches":        "২০ ইঞ্চি",
		"Spare Tyre Size":  "স্পেয়ার টায়ার সাইজ",
	}
}

func (p9s *Porsche911Seeder) Seed(db *gorm.DB) error {
	// Lookup brand by slug
	var brand models.BrandModel
	if err := db.Where("slug = ?", "porsche").First(&brand).Error; err != nil {
		return fmt.Errorf("brand not found: %w", err)
	}

	// Lookup category by ID
	var category models.CategoryModel
	if err := db.Where("id = ?", 18).First(&category).Error; err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "Porsche 911").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Porsche 911",
				Slug:       "porsche-911",
				BrandID:    &brand.ID,
				CategoryID: &category.ID,
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
		"engine_type":                 "3.0L Turbo Flat-6",
		"engine_displacement":         "2981 cc",
		"engine_cylinders":            "6",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Flat",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "450 hp @ 6500 rpm",
		"engine_max_torque":           "530 Nm @ 2300 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "91.0 x 76.4 mm",
		"engine_compression_ratio":    "10.2:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "8-Speed",
		"transmission_drive_type":     "AWD",
		"transmission_clutch_type":    "PDK",
		"performance_top_speed":       "308 km/h",
		"performance_acceleration":    "3.5 seconds",
		"performance_mileage_arai":    "9 km/L",
		"performance_mileage_city":    "6 km/L",
		"performance_mileage_highway": "12 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4519 mm",
		"dimensions_width":            "1852 mm",
		"dimensions_height":           "1300 mm",
		"dimensions_wheelbase":        "2450 mm",
		"dimensions_front_tread":      "1550 mm",
		"dimensions_rear_tread":       "1545 mm",
		"dimensions_ground_clearance": "130 mm",
		"dimensions_boot_space":       "132 L",
		"dimensions_fuel_tank":        "64 L",
		"dimensions_turning_radius":   "5.6 m",
		"weight_kerb_weight":          "1505 kg",
		"weight_gross_weight":         "1920 kg",
		"capacity_seating_capacity":   "4",
		"capacity_doors":              "2",
		"suspension_front":            "Double Wishbone",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "265/35 R20",
		"tyres_wheel_size":            "20 inches",
		"tyres_spare_size":            "265/35 R20",
		"exterior_colors":             "Guards Red, Black, White, Silver, Blue, Red, Gray, Green, Yellow",
		"exterior_headlights":         "Full LED",
		"exterior_daytime_running":    "Porsche Dynamic Light System",
		"exterior_roof_rails":         "No",
		"exterior_wheels":             "20-inch Alloy Wheels",
		"interior_seats_material":     "Premium Leather",
		"interior_seats_adjustment":   "18-way Electric",
		"interior_ac":                 "2-Zone Auto AC",
		"infotainment_screen_size":    "10.9-inch Touchscreen",
		"infotainment_connectivity":   "PCM, Android Auto, Apple CarPlay, Bluetooth, WiFi",
		"infotainment_speakers":       "Bose Surround Sound",
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
		"features_cruise_control":     "Adaptive Cruise Control",
		"features_sunroof":            "Glass Roof",
		"year":                        "2023",
		"category":                    "Sports Car",
		"subcategory":                 "Sports Car",
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
				Value:           p9s.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Porsche 911")
	return nil
}
