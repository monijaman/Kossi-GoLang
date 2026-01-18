package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MINICountrymanSeeder seeds specifications for MINI Countryman
type MINICountrymanSeeder struct {
	BaseSeeder
}

// NewMINICountrymanSeeder creates a new MINI Countryman seeder
func NewMINICountrymanSeeder() *MINICountrymanSeeder {
	return &MINICountrymanSeeder{
		BaseSeeder: BaseSeeder{name: "MINI Countryman Specifications"},
	}
}

func (mcs *MINICountrymanSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Turbo":                    "১.৫ লিটার টার্বো",
		"Subcompact SUV":                "সাবকমপ্যাক্ট এসইউভি",
		"2020":                          "২০২০",
		"SUV":                           "এসইউভি",
		"British Racing Green":          "ব্রিটিশ রেসিং গ্রিন",
		"Inline":                        "ইনলাইন",
		"1499":                          "১৪৯৯ সিসি",
		"3":                             "৩ (সিলিন্ডার)",
		"136 hp":                        "১৩৬ হর্স পাওয়ার",
		"136 hp @ 4400 rpm":             "৪৪০০ আরপিএম এ ১৩৬ হর্স পাওয়ার",
		"220 Nm @ 1250 rpm":             "১২৫০ আরপিএম এ ২২০ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"9.7 seconds":                   "৯.৭ সেকেন্ড",
		"190 km/h":                      "১৯০ কিমি/ঘণ্টা",
		"14 km/L":                       "১৪ কিমি/লিটার",
		"11 km/L":                       "১১ কিমি/লিটার",
		"17 km/L":                       "১৭ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"7-Speed DCT":                   "৭-স্পিড ডিসিটি",
		"All-Wheel Drive":               "সমস্ত চাকা চালিত",
		"Dual Clutch":                   "ডুয়াল ক্লাচ",
		"4297 mm":                       "৪২৯৭ মিমি",
		"1822 mm":                       "১৮২২ মিমি",
		"1557 mm":                       "১৫৫৭ মিমি",
		"2595 mm":                       "২৫৯৫ মিমি",
		"5":                             "৫ (সিট)",
		"5 Doors":                       "৫ ডোর",
		"5 Seats":                       "৫ সিট",
		"Black":                         "ব্ল্যাক",
		"Gray":                          "গ্রে",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Green":                         "গ্রিন",
		"Orange":                        "অরেঞ্জ",
		"LED":                           "এলইডি",
		"Halogen":                       "হ্যালোজেন",
		"LED Headlights":                "এলইডি হেডলাইট",
		"Halogen Headlights":            "হ্যালোজেন হেডলাইট",
		"LED Daytime Running Lights":    "এলইডি ডে টাইম রানিং লাইট",
		"Roof Rails":                    "রুফ রেল",
		"No Roof Rails":                 "রুফ রেল নেই",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"16-inch Alloy Wheels":          "১৬-ইঞ্চি অ্যালয় হুইল",
		"17-inch Alloy Wheels":          "১৭-ইঞ্চি অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"205/55 R17":                    "২০৫/৫৫ আর১৭",
		"225/50 R18":                    "২২৫/৫০ আর১৮",
		"Manual Adjustment":             "ম্যানুয়াল অ্যাডজাস্টমেন্ট",
		"Electric Adjustment":           "ইলেকট্রিক অ্যাডজাস্টমেন্ট",
		"Height Adjustable":             "হাইট অ্যাডজাস্টেবল",
		"6-way Manual":                  "৬-ওয়ে ম্যানুয়াল",
		"8-way Electric":                "৮-ওয়ে ইলেকট্রিক",
		"Fabric":                        "ফ্যাব্রিক",
		"Leather":                       "লেদার",
		"Leatherette":                   "লেদারেট",
		"Fabric Seats":                  "ফ্যাব্রিক সিট",
		"Leather Seats":                 "লেদার সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Manual AC":                     "ম্যানুয়াল এসি",
		"Auto AC":                       "অটো এসি",
		"Single Zone":                   "সিঙ্গেল জোন",
		"Dual Zone":                     "ডুয়াল জোন",
		"Rear AC Vents":                 "রিয়ার এসি ভেন্ট",
		"6.5-inch Touchscreen":          "৬.৫-ইঞ্চি টাচস্ক্রিন",
		"8.8-inch Touchscreen":          "৮.৮-ইঞ্চি টাচস্ক্রিন",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"6 Speakers":                    "৬ স্পিকার",
		"8 Speakers":                    "৮ স্পিকার",
		"Harman Kardon":                 "হারম্যান কার্ডন",
		"ABS":                           "এবিএস",
		"EBD":                           "ইবিডি",
		"Brake Assist":                  "ব্রেক অ্যাসিস্ট",
		"ESP":                           "ইএসপি",
		"Traction Control":              "ট্র্যাকশন কন্ট্রোল",
		"Hill Hold Control":             "হিল হোল্ড কন্ট্রোল",
		"Front Airbags":                 "ফ্রন্ট এয়ারব্যাগ",
		"Side Airbags":                  "সাইড এয়ারব্যাগ",
		"Knee Airbags":                  "নী এয়ারব্যাগ",
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
		"350 L":                         "৩৫০ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"51 L":                          "৫১ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"165 mm":                        "১৬৫ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1410 kg":                       "১৪১০ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1910 kg":                       "১৯১০ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"6.0 m":                         "৬.০ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1499 cc":                       "১৪৯৯ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"4":                             "৪",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"82.0 x 94.6 mm":                "৮২.০ x ৯৪.৬ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"11.0:1":                        "১১.০:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
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
		"1561 mm":                       "১৫৬১ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1564 mm":                       "১৫৬৪ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"Door Count":                    "ডোর কাউন্ট",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"MacPherson Strut":              "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Multi-link":                    "মাল্টি-লিঙ্ক",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Tyre Size":                     "টায়ার সাইজ",
		"Wheel Size":                    "হুইল সাইজ",
		"17 inches":                     "১৭ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
	}
}

func (mcs *MINICountrymanSeeder) Seed(db *gorm.DB) error {
	// Lookup brand by slug
	var brand models.BrandModel
	if err := db.Where("slug = ?", "mini").First(&brand).Error; err != nil {
		return fmt.Errorf("brand not found: %w", err)
	}

	// Lookup category by ID
	var category models.CategoryModel
	if err := db.Where("id = ?", 18).First(&category).Error; err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	// First, find or create the product
	var product models.ProductModel
	if err := db.Where("name = ?", "MINI Countryman").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "MINI Countryman",
				Slug:       "mini-countryman",
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
		"brand":                       "MINI",
		"model":                       "Countryman",
		"engine_type":                 "1.5L Turbo",
		"engine_displacement":         "1499 cc",
		"engine_cylinders":            "3",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "136 hp @ 4400 rpm",
		"engine_max_torque":           "220 Nm @ 1250 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "82.0 x 94.6 mm",
		"engine_compression_ratio":    "11.0:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "7-Speed DCT",
		"transmission_drive_type":     "AWD",
		"transmission_clutch_type":    "Dual Clutch",
		"performance_top_speed":       "190 km/h",
		"performance_acceleration":    "9.7 seconds",
		"performance_mileage_arai":    "14 km/L",
		"performance_mileage_city":    "11 km/L",
		"performance_mileage_highway": "17 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4297 mm",
		"dimensions_width":            "1822 mm",
		"dimensions_height":           "1557 mm",
		"dimensions_wheelbase":        "2595 mm",
		"dimensions_front_tread":      "1561 mm",
		"dimensions_rear_tread":       "1564 mm",
		"dimensions_ground_clearance": "165 mm",
		"dimensions_boot_space":       "350 L",
		"dimensions_fuel_tank":        "51 L",
		"dimensions_turning_radius":   "6.0 m",
		"weight_kerb_weight":          "1410 kg",
		"weight_gross_weight":         "1910 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Multi-link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "205/55 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "205/55 R17",
		"exterior_colors":             "British Racing Green, Black, Gray, White, Silver, Blue, Red, Green, Orange",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Leatherette",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Dual Zone Auto AC",
		"infotainment_screen_size":    "8.8-inch Touchscreen",
		"infotainment_connectivity":   "Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "6 Speakers",
		"safety_airbags":              "Front, Side, Knee, Curtain Airbags",
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
		"year":                        "2020",
		"category":                    "SUV",
		"subcategory":                 "Subcompact SUV",
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
				Value:           mcs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for MINI Countryman")
	return nil
}
