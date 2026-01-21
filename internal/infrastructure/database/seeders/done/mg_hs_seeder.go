package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MGHSSeeder seeds specifications for MG HS
type MGHSSeeder struct {
	BaseSeeder
}

// NewMGHSSeeder creates a new MG HS seeder
func NewMGHSSeeder() *MGHSSeeder {
	return &MGHSSeeder{
		BaseSeeder: BaseSeeder{name: "MG HS Specifications"},
	}
}

func (mhs *MGHSSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Turbo":                    "১.৫ লিটার টার্বো",
		"Compact SUV":                   "কমপ্যাক্ট এসইউভি",
		"2019":                          "২০১৯",
		"SUV":                           "এসইউভি",
		"Pearl White":                   "পার্ল হোয়াইট",
		"Inline":                        "ইনলাইন",
		"1498":                          "১৪৯৮ সিসি",
		"4":                             "৪ (সিলিন্ডার)",
		"162 hp":                        "১৬২ হর্স পাওয়ার",
		"162 hp @ 5600 rpm":             "৫৬০০ আরপিএম এ ১৬২ হর্স পাওয়ার",
		"250 Nm @ 1700-4400 rpm":        "১৭০০-৪৪০০ আরপিএম এ ২৫০ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"9.1 seconds":                   "৯.১ সেকেন্ড",
		"190 km/h":                      "১৯০ কিমি/ঘণ্টা",
		"15 km/L":                       "১৫ কিমি/লিটার",
		"12 km/L":                       "১২ কিমি/লিটার",
		"17 km/L":                       "১৭ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"7-Speed DCT":                   "৭-স্পিড ডিসিটি",
		"Front-Wheel Drive":             "ফ্রন্ট-হুইল ড্রাইভ",
		"Dual Clutch":                   "ডুয়াল ক্লাচ",
		"4486 mm":                       "৪৪৮৬ মিমি",
		"1876 mm":                       "১৮৭৬ মিমি",
		"1662 mm":                       "১৬৬২ মিমি",
		"2720 mm":                       "২৭২০ মিমি",
		"5":                             "৫ (সিট)",
		"5 Doors":                       "৫ ডোর",
		"4 Doors":                       "৪ ডোর",
		"1 Door":                        "১ ডোর",
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
		"Roof Rails":                    "রুফ রেল",
		"No Roof Rails":                 "রুফ রেল নেই",
		"Steel Wheels":                  "স্টিল হুইল",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"17-inch Alloy Wheels":          "১৭-ইঞ্চি অ্যালয় হুইল",
		"18-inch Alloy Wheels":          "১৮-ইঞ্চি অ্যালয় হুইল",
		"19-inch Alloy Wheels":          "১৯-ইঞ্চি অ্যালয় হুইল",
		"235/50 R18":                    "২৩৫/৫০ আর১৮",
		"235/45 R19":                    "২৩৫/৪৫ আর১৯",
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
		"12.3-inch Touchscreen":         "১২.৩-ইঞ্চি টাচস্ক্রিন",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"USB":                           "ইউএসবি",
		"AUX":                           "অক্সিলারি",
		"HDMI":                          "এইচডিএমআই",
		"4 Speakers":                    "৪ স্পিকার",
		"6 Speakers":                    "৬ স্পিকার",
		"8 Speakers":                    "৮ স্পিকার",
		"10 Speakers":                   "১০ স্পিকার",
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
		"Shark Fin Antenna":             "শার্ক ফিন অ্যান্টেনা",
		"LED Tail Lights":               "এলইডি টেইল লাইট",
		"Fog Lights":                    "ফগ লাইট",
		"LED Fog Lights":                "এলইডি ফগ লাইট",
		"Cornering Lights":              "কর্নারিং লাইট",
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
		"Adaptive Cruise Control":       "অ্যাডাপটিভ ক্রুজ কন্ট্রোল",
		"Wireless Charger":              "ওয়্যারলেস চার্জার",
		"Ambient Lighting":              "অ্যাম্বিয়েন্ট লাইটিং",
		"Footwell Lighting":             "ফুটওয়েল লাইটিং",
		"Reading Lamps":                 "রিডিং ল্যাম্প",
		"Glove Box Lamp":                "গ্লোভ বক্স ল্যাম্প",
		"Trunk Lamp":                    "ট্রাঙ্ক ল্যাম্প",
		"Welcome Function":              "ওয়েলকাম ফাংশন",
		"Auto Dimming IRVM":             "অটো ডিমিং আইআরভিএম",
		"ORVMs with Turn Indicators":    "ওআরভিএম উইথ টার্ন ইন্ডিকেটর",
		"Power Windows":                 "পাওয়ার উইন্ডো",
		"One Touch Up/Down":             "ওয়ান টাচ আপ/ডাউন",
		"Anti-pinch":                    "অ্যান্টি-পিঞ্চ",
		"Boot Release":                  "বুট রিলিজ",
		"Remote Boot Release":           "রিমোট বুট রিলিজ",
		"Rear Defogger":                 "রিয়ার ডিফগার",
		"Rear Wiper":                    "রিয়ার ওয়াইপার",
		"Intermittent Rear Wiper":       "ইন্টারমিটেন্ট রিয়ার ওয়াইপার",
		"Heated ORVMs":                  "হিটেড ওআরভিএম",
		"Electrically Adjustable ORVMs": "ইলেকট্রিক্যালি অ্যাডজাস্টেবল ওআরভিএম",
		"Power Folding ORVMs":           "পাওয়ার ফোল্ডিং ওআরভিএম",
		"Boot Space":                    "বুট স্পেস",
		"448 L":                         "৪৪৮ লিটার",
		"1380 L":                        "১৩৮০ লিটার",
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"55 L":                          "৫৫ লিটার",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"185 mm":                        "১৮৫ মিমি",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"1485 kg":                       "১৪৮৫ কেজি",
		"Gross Weight":                  "গ্রস ওয়েট",
		"1985 kg":                       "১৯৮৫ কেজি",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"5.8 m":                         "৫.৮ মিটার",
		"Top Speed":                     "টপ স্পিড",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"1498 cc":                       "১৪৯৮ সিসি",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Direct Injection":              "ডাইরেক্ট ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"74.0 x 86.6 mm":                "৭৪.০ x ৮৬.৬ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"10.3:1":                        "১০.৩:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"Yes":                           "হ্যাঁ",
		"Super Charger":                 "সুপার চার্জার",
		"No":                            "না",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Automatic":                     "অটোমেটিক",
		"Gear Box":                      "গিয়ার বক্স",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"FWD":                           "এফডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"Mileage (ARAI)":                "মাইলেজ (এআরএআই)",
		"Mileage (City)":                "মাইলেজ (সিটি)",
		"Mileage (Highway)":             "মাইলেজ (হাইওয়ে)",
		"18 km/L":                       "১৮ কিমি/লিটার",
		"Emission Norm Compliance":      "ইমিশন নর্ম কমপ্লায়েন্স",
		"BS VI":                         "বিএস ভি",
		"Length":                        "দৈর্ঘ্য",
		"Width":                         "প্রস্থ",
		"Height":                        "উচ্চতা",
		"Wheelbase":                     "হুইলবেস",
		"Front Tread":                   "ফ্রন্ট ট্রেড",
		"1580 mm":                       "১৫৮০ মিমি",
		"Rear Tread":                    "রিয়ার ট্রেড",
		"1583 mm":                       "১৫৮৩ মিমি",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"Door Count":                    "ডোর কাউন্ট",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",

		"Front Suspension": "ফ্রন্ট সাসপেনশন",
		"MacPherson Strut": "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":  "রিয়ার সাসপেনশন",
		"Torsion Beam":     "টর্শন বিম",
		"Front Brake Type": "ফ্রন্ট ব্রেক টাইপ",
		"Disc":             "ডিস্ক",
		"Rear Brake Type":  "রিয়ার ব্রেক টাইপ",
		"Tyre Size":        "টায়ার সাইজ",
		"Wheel Size":       "হুইল সাইজ",
		"18 inches":        "১৮ ইঞ্চি",
		"Spare Tyre Size":  "স্পেয়ার টায়ার সাইজ",
	}
}

func (mhs *MGHSSeeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the MG brand
	var brand models.BrandModel
	brandSlug := "mg"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "MG",
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
	productSlug := "mg-hs"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "MG HS",
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
		"brand":                       "MG",
		"model":                       "HS",
		"engine_type":                 "1.5L Turbo",
		"engine_displacement":         "1498 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "162 hp @ 5600 rpm",
		"engine_max_torque":           "250 Nm @ 1700-4400 rpm",
		"engine_fuel_supply_system":   "Direct Injection",
		"engine_bore_stroke":          "74.0 x 86.6 mm",
		"engine_compression_ratio":    "10.3:1",
		"engine_turbo_charger":        "Yes",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "7-Speed DCT",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "Dual Clutch",
		"performance_top_speed":       "190 km/h",
		"performance_acceleration":    "9.1 seconds",
		"performance_mileage_arai":    "15 km/L",
		"performance_mileage_city":    "12 km/L",
		"performance_mileage_highway": "18 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4486 mm",
		"dimensions_width":            "1876 mm",
		"dimensions_height":           "1662 mm",
		"dimensions_wheelbase":        "2720 mm",
		"dimensions_front_tread":      "1580 mm",
		"dimensions_rear_tread":       "1583 mm",
		"dimensions_ground_clearance": "185 mm",
		"dimensions_boot_space":       "448 L",
		"dimensions_fuel_tank":        "55 L",
		"dimensions_turning_radius":   "5.8 m",
		"weight_kerb_weight":          "1485 kg",
		"weight_gross_weight":         "1985 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "5",
		"suspension_front":            "MacPherson Strut",
		"suspension_rear":             "Torsion Beam",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "235/50 R18",
		"tyres_wheel_size":            "18 inches",
		"tyres_spare_size":            "235/50 R18",
		"exterior_colors":             "Pearl White, Black, Gray, Silver, Blue, Red, Orange",
		"exterior_headlights":         "LED",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "18-inch Alloy Wheels",
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
		"features_cruise_control":     "Adaptive Cruise Control",
		"features_sunroof":            "Panoramic Sunroof",
		"year":                        "2019",
		"category":                    "SUV",
		"subcategory":                 "Compact SUV",
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
				Value:           mhs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for MG HS")
	return nil
}
