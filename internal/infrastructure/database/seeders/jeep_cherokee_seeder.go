package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JeepCherokeeSeeder seeds specifications for Jeep Cherokee
type JeepCherokeeSeeder struct {
	BaseSeeder
}

// NewJeepCherokeeSeeder creates a new Jeep Cherokee seeder
func NewJeepCherokeeSeeder() *JeepCherokeeSeeder {
	return &JeepCherokeeSeeder{
		BaseSeeder: BaseSeeder{name: "Jeep Cherokee Specifications"},
	}
}

func (jchs *JeepCherokeeSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.4L I4":                       "২.৪ লিটার আই৪",
		"Mid-size SUV":                  "মিড-সাইজ এসইউভি",
		"2023":                          "২০২৩",
		"SUV":                           "এসইউভি",
		"Bright White":                  "ব্রাইট হোয়াইট",
		"Inline":                        "ইনলাইন",
		"2360 cc":                       "২৩৬০ সিসি",
		"180 hp @ 6400 rpm":             "১৮০ হর্স পাওয়ার @ ৬৪০০ আরপিএম",
		"229 Nm @ 3900 rpm":             "২২৯ এনএম @ ৩৯০০ আরপিএম",
		"Petrol":                        "পেট্রোল",
		"Naturally Aspirated":           "ন্যাচারালি অ্যাসপিরেটেড",
		"10.2 seconds":                  "১০.২ সেকেন্ড",
		"196 km/h":                      "১৯৬ কিমি/ঘণ্টা",
		"12 km/L":                       "১২ কিমি/লিটার",
		"9 km/L":                        "৯ কিমি/লিটার",
		"15 km/L":                       "১৫ কিমি/লিটার",
		"Automatic":                     "অটোমেটিক",
		"9-Speed":                       "৯-স্পিড",
		"Front-Wheel Drive":             "ফ্রন্ট-হুইল ড্রাইভ",
		"All-Wheel Drive":               "অল-হুইল ড্রাইভ",
		"4624 mm":                       "৪৬২৪ মিমি",
		"1859 mm":                       "১৮৫৯ মিমি",
		"1686 mm":                       "১৬৮৬ মিমি",
		"2707 mm":                       "২৭০৭ মিমি",
		"1585 mm":                       "১৫৮৫ মিমি",
		"5":                             "৫",
		"4":                             "৪",
		"570 L":                         "৫৭০ লিটার",
		"60 L":                          "৬০ লিটার",
		"193 mm":                        "১৯৩ মিমি",
		"1680 kg":                       "১৬৮০ কেজি",
		"2250 kg":                       "২২৫০ কেজি",
		"5.8 m":                         "৫.৮ মিটার",
		"Black":                         "ব্ল্যাক",
		"Gray":                          "গ্রে",
		"White":                         "হোয়াইট",
		"Silver":                        "সিলভার",
		"Blue":                          "ব্লু",
		"Red":                           "রেড",
		"Green":                         "গ্রিন",
		"Orange":                        "অরেঞ্জ",
		"Halogen":                       "হ্যালোজেন",
		"Halogen Headlights":            "হ্যালোজেন হেডলাইট",
		"LED Daytime Running Lights":    "এলইডি ডে টাইম রানিং লাইট",
		"Roof Rails":                    "রুফ রেল",
		"Alloy Wheels":                  "অ্যালয় হুইল",
		"17-inch Alloy Wheels":          "১৭-ইঞ্চি অ্যালয় হুইল",
		"225/60 R17":                    "২২৫/৬০ আর১৭",
		"Manual Adjustment":             "ম্যানুয়াল অ্যাডজাস্টমেন্ট",
		"6-way Manual":                  "৬-ওয়ে ম্যানুয়াল",
		"Cloth":                         "ক্লথ",
		"Cloth Seats":                   "ক্লথ সিট",
		"60:40 Split":                   "৬০:৪০ স্প্লিট",
		"Folding Rear Seats":            "ফোল্ডিং রিয়ার সিট",
		"Manual AC":                     "ম্যানুয়াল এসি",
		"7-inch Touchscreen":            "৭-ইঞ্চি টাচস্ক্রিন",
		"Uconnect":                      "ইউকানেক্ট",
		"Android Auto":                  "অ্যান্ড্রয়েড অটো",
		"Apple CarPlay":                 "অ্যাপল কারপ্লে",
		"Bluetooth":                     "ব্লুটুথ",
		"6 Speakers":                    "৬ স্পিকার",
		"ABS":                           "এবিএস",
		"EBD":                           "ইবিডি",
		"Brake Assist":                  "ব্রেক অ্যাসিস্ট",
		"ESP":                           "ইএসপি",
		"Traction Control":              "ট্র্যাকশন কন্ট্রোল",
		"Hill Hold Control":             "হিল হোল্ড কন্ট্রোল",
		"Front Airbags":                 "ফ্রন্ট এয়ারব্যাগ",
		"Side Airbags":                  "সাইড এয়ারব্যাগ",
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
		"Chrome Accents":                "ক্রোম অ্যাকসেন্ট",
		"Body Colored Bumpers":          "বডি কালার্ড বাম্পার",
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
		"Urethane Steering Wheel":       "ইউরেথেন স্টিয়ারিং হুইল",
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
		"Fuel Tank":                     "ফুয়েল ট্যাঙ্ক",
		"Ground Clearance":              "গ্রাউন্ড ক্লিয়ারেন্স",
		"Kerb Weight":                   "কার্ব ওয়েট",
		"Gross Weight":                  "গ্রস ওয়েট",
		"Turning Radius":                "টার্নিং রেডিয়াস",
		"Top Speed":                     "টপ স্পিড",
		"Acceleration 0-100 km/h":       "০-১০০ কিমি/ঘণ্টা অ্যাকসেলারেশন",
		"Engine Type":                   "ইঞ্জিন টাইপ",
		"Displacement":                  "ডিসপ্লেসমেন্ট",
		"Max Power":                     "ম্যাক্স পাওয়ার",
		"Max Torque":                    "ম্যাক্স টর্ক",
		"No. of Cylinders":              "সিলিন্ডারের সংখ্যা",
		"Valves per Cylinder":           "প্রতি সিলিন্ডার ভালভ",
		"Fuel Supply System":            "ফুয়েল সাপ্লাই সিস্টেম",
		"Multi-Point Fuel Injection":    "মাল্টি-পয়েন্ট ফুয়েল ইনজেকশন",
		"Bore x Stroke":                 "বোর x স্ট্রোক",
		"88.0 x 97.0 mm":                "৮৮.০ x ৯৭.০ মিমি",
		"Compression Ratio":             "কম্প্রেশন রেশিও",
		"10.0:1":                        "১০.০:১",
		"Turbo Charger":                 "টার্বো চার্জার",
		"No":                            "না",
		"Super Charger":                 "সুপার চার্জার",
		"Transmission Type":             "ট্রান্সমিশন টাইপ",
		"Gear Box":                      "গিয়ার বক্স",
		"Drive Type":                    "ড্রাইভ টাইপ",
		"FWD":                           "এফডব্লিউডি",
		"Clutch Type":                   "ক্লাচ টাইপ",
		"N/A":                           "এন/এ",
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
		"Rear Tread":                    "রিয়ার ট্রেড",
		"Seating Capacity":              "সিটিং ক্যাপাসিটি",
		"Door Count":                    "ডোর কাউন্ট",
		"Fuel Tank Capacity":            "ফুয়েল ট্যাঙ্ক ক্যাপাসিটি",
		"Ground Clearance Unladen":      "গ্রাউন্ড ক্লিয়ারেন্স আনলোডেড",
		"Front Suspension":              "ফ্রন্ট সাসপেনশন",
		"McPherson Strut":               "ম্যাকফারসন স্ট্রাট",
		"Rear Suspension":               "রিয়ার সাসপেনশন",
		"Multi-Link":                    "মাল্টি-লিঙ্ক",
		"Front Brake Type":              "ফ্রন্ট ব্রেক টাইপ",
		"Disc":                          "ডিস্ক",
		"Rear Brake Type":               "রিয়ার ব্রেক টাইপ",
		"Tyre Size":                     "টায়ার সাইজ",
		"Wheel Size":                    "হুইল সাইজ",
		"17 inches":                     "১৭ ইঞ্চি",
		"Spare Tyre Size":               "স্পেয়ার টায়ার সাইজ",
	}
}

func (jchs *JeepCherokeeSeeder) Seed(db *gorm.DB) error {
	// Find the Car category (ID: 18)
	var carCategory models.CategoryModel
	if err := db.Where("id = ?", 18).First(&carCategory).Error; err != nil {
		return fmt.Errorf("category with id 18 not found: %w", err)
	}

	// Find or create the Jeep brand
	var brand models.BrandModel
	brandSlug := "jeep"
	if err := db.Where("slug = ?", brandSlug).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			brand = models.BrandModel{
				Name:   "Jeep",
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
	productSlug := "jeep-cherokee"
	if err := db.Where("slug = ?", productSlug).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			product = models.ProductModel{
				Name:       "Jeep Cherokee",
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
		"engine_type":                 "2.4L I4",
		"engine_displacement":         "2360 cc",
		"engine_cylinders":            "4",
		"engine_valves_per_cylinder":  "4",
		"engine_configuration":        "Inline",
		"engine_fuel_type":            "Petrol",
		"engine_max_power":            "180 hp @ 6400 rpm",
		"engine_max_torque":           "229 Nm @ 3900 rpm",
		"engine_fuel_supply_system":   "Multi-Point Fuel Injection",
		"engine_bore_stroke":          "88.0 x 97.0 mm",
		"engine_compression_ratio":    "10.0:1",
		"engine_turbo_charger":        "No",
		"engine_super_charger":        "No",
		"transmission_type":           "Automatic",
		"transmission_gearbox":        "9-Speed",
		"transmission_drive_type":     "FWD",
		"transmission_clutch_type":    "N/A",
		"performance_top_speed":       "196 km/h",
		"performance_acceleration":    "10.2 seconds",
		"performance_mileage_arai":    "12 km/L",
		"performance_mileage_city":    "9 km/L",
		"performance_mileage_highway": "15 km/L",
		"performance_emission_norm":   "BS VI",
		"dimensions_length":           "4624 mm",
		"dimensions_width":            "1859 mm",
		"dimensions_height":           "1686 mm",
		"dimensions_wheelbase":        "2707 mm",
		"dimensions_front_tread":      "1585 mm",
		"dimensions_rear_tread":       "1585 mm",
		"dimensions_ground_clearance": "193 mm",
		"dimensions_boot_space":       "570 L",
		"dimensions_fuel_tank":        "60 L",
		"dimensions_turning_radius":   "5.8 m",
		"weight_kerb_weight":          "1680 kg",
		"weight_gross_weight":         "2250 kg",
		"capacity_seating_capacity":   "5",
		"capacity_doors":              "4",
		"suspension_front":            "McPherson Strut",
		"suspension_rear":             "Multi-Link",
		"brakes_front":                "Disc",
		"brakes_rear":                 "Disc",
		"tyres_size":                  "225/60 R17",
		"tyres_wheel_size":            "17 inches",
		"tyres_spare_size":            "225/60 R17",
		"exterior_colors":             "Bright White, Black, Gray, White, Silver, Blue, Red, Green, Orange",
		"exterior_headlights":         "Halogen",
		"exterior_daytime_running":    "LED DRLs",
		"exterior_roof_rails":         "Yes",
		"exterior_wheels":             "17-inch Alloy Wheels",
		"interior_seats_material":     "Cloth",
		"interior_seats_adjustment":   "6-way Manual",
		"interior_ac":                 "Manual AC",
		"infotainment_screen_size":    "7-inch Touchscreen",
		"infotainment_connectivity":   "Uconnect, Android Auto, Apple CarPlay, Bluetooth",
		"infotainment_speakers":       "6 Speakers",
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
				Locale:          "bn",
				Value:           jchs.getBanglaTranslations()[value],
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("Error creating translation for specification %d: %v", spec.ID, err)
			}
		} else {
			log.Printf("Specification key not found: %s", key)
		}
	}

	log.Printf("Seeded specifications for Jeep Cherokee")
	return nil
}
