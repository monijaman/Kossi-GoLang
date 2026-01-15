package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// FordEcoSportSeeder seeds specifications for Ford EcoSport
type FordEcoSportSeeder struct {
	BaseSeeder
}

// NewFordEcoSportSeeder creates a new Ford EcoSport seeder
func NewFordEcoSportSeeder() *FordEcoSportSeeder {
	return &FordEcoSportSeeder{
		BaseSeeder: BaseSeeder{name: "Ford EcoSport Specifications"},
	}
}

func (fess *FordEcoSportSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L":                                "১.৫ লিটার",
		"1st":                                 "১ম প্রজন্ম",
		"Compact SUV":                         "কমপ্যাক্ট এসইউভি",
		"2012":                                "২০১২",
		"SUV":                                 "এসইউভি",
		"Carbon Black":                        "কার্বন ব্ল্যাক",
		"In-line":                             "ইন-লাইন",
		"1499":                                "১৪৯৯ সিসি",
		"4":                                   "৪ (সিলিন্ডার)",
		"110 hp":                              "১১০ হর্স পাওয়ার",
		"110 hp @ 6300 rpm":                   "৬৩০০ আরপিএম এ ১১০ হর্স পাওয়ার",
		"140 Nm @ 4400 rpm":                   "৪৪০০ আরপিএম এ ১৪০ এনএম",
		"Petrol":                              "পেট্রোল",
		"Naturally Aspirated":                 "ন্যাচারালি অ্যাসপিরেটেড",
		"12.5 seconds":                        "১২.৫ সেকেন্ড",
		"182 km/h":                            "১৮২ কিমি/ঘণ্টা",
		"13 km/L":                             "১৩ কিমি/লিটার",
		"17 km/L":                             "১৭ কিমি/লিটার",
		"15 km/L":                             "১৫ কিমি/লিটার",
		"Manual (Transmission)":               "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":                      "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":                   "সামনে চাকা চালিত",
		"Single Plate Dry Clutch":             "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"3999 mm":                             "৩৯৯৯ মিমি",
		"1765 mm":                             "১৭৬৫ মিমি",
		"1708 mm":                             "১৭০৮ মিমি",
		"2520 mm":                             "২৫২০ মিমি",
		"200 mm":                              "২০০ মিমি",
		"1200 kg":                             "১২০০ কেজি",
		"352 L":                               "৩৫২ লিটার",
		"205/60 R16":                          "২০৫/৬০ আর১৬",
		"Radial Tubeless":                     "রেডিয়াল টিউবলেস",
		"Alloy":                               "অ্যালয়",
		"Halogen (Headlamps)":                 "হ্যালোজেন (হেডল্যাম্পস)",
		"Yes (DRL)":                           "হ্যাঁ (ডিআরএল)",
		"Halogen (Fog Lamp)":                  "হ্যালোজেন (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":                      "হ্যাঁ (ফগ ল্যাম্প)",
		"No (Sunroof)":                        "না (সানরুফ)",
		"Yes (Roof Rails)":                    "হ্যাঁ (রুফ রেলস)",
		"Power (ORVM)":                        "পাওয়ার (ওআরভিএম)",
		"Manual (Wiper)":                      "ম্যানুয়াল (ওয়াইপার)",
		"5 (Seating)":                         "৫ (সিটিং)",
		"5 (Seats)":                           "৫ (সিটস)",
		"Manual (Driver Seat)":                "ম্যানুয়াল (ড্রাইভার সিট)",
		"No (Ventilated Seats)":               "না (ভেন্টিলেটেড সিটস)",
		"Manual (Air Conditioning)":           "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"No (Climate Control)":                "না (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":          "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"6.5 inch (Screen)":                   "৬.৫ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":                 "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":                  "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Sony (Sound System)":                 "সনি (সাউন্ড সিস্টেম)",
		"4 (Speakers)":                        "৪ (স্পিকার)",
		"No (Ambient Lighting)":               "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                           "হ্যাঁ (এবিএস)",
		"Driver + Passenger + Side (Airbags)": "ড্রাইভার + প্যাসেঞ্জার + সাইড (এয়ারব্যাগস)",
		"Yes (EBD)":                           "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Drum (Rear)":          "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"No (Traction Control)":               "না (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                           "হ্যাঁ (ইএসসি)",
		"No (Hill Assist)":                    "না (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":                "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":                   "হ্যাঁ (টাচস্ক্রিন)",
		"No (Rear Camera)":                    "না (রিয়ার ক্যামেরা)",
		"Keyless Entry":                       "কীলেস এন্ট্রি",
		"No (Push Button Start)":              "না (পুশ বাটন স্টার্ট)",
		"No (Height Adjustable Driver Seat)":  "না (হাইট অ্যাডজাস্টেবল ড্রাইভার সিট)",
		"No (Rear AC Vents)":                  "না (রিয়ার এসি ভেন্টস)",
		"No (Steering Mounted Controls)":      "না (স্টিয়ারিং মাউন্টেড কন্ট্রোলস)",
		"No (Cruise Control)":                 "না (ক্রুজ কন্ট্রোল)",
		"No (Front Parking Sensors)":          "না (সামনে পার্কিং সেন্সরস)",
		"No (Rear Parking Sensors)":           "না (রিয়ার পার্কিং সেন্সরস)",
		"No (Auto Headlamps)":                 "না (অটো হেডল্যাম্পস)",
		"No (Rain Sensing Wipers)":            "না (রেইন সেন্সিং ওয়াইপারস)",
		"No (Follow Me Home Headlamps)":       "না (ফলো মি হোম হেডল্যাম্পস)",
		"Yes (Alloy Wheels)":                  "হ্যাঁ (অ্যালয় হুইলস)",
		"16 inch":                             "১৬ ইঞ্চি",
		"Independent McPherson Strut":         "ইন্ডিপেন্ডেন্ট ম্যাকফারসন স্ট্রাট",
		"Twist Beam":                          "টুইস্ট বিম",
		"Rack and Pinion":                     "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt":                                "টিল্ট",
		"16":                                  "১৬",
	}
}

// Seed implements the Seeder interface for Ford EcoSport
func (fess *FordEcoSportSeeder) Seed(db *gorm.DB) error {
	productSlug := "ford-ecosport"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}
		return err
	}
	productID := prod.ID

	existingKeyMapping := map[string]uint{
		"Variant":                        728,
		"Generation":                     729,
		"Segment":                        730,
		"Launch Year":                    731,
		"Engine Configuration":           732,
		"Valves Per Cylinder":            733,
		"Engine Aspiration":              734,
		"Differential Type":              735,
		"Power to Weight (HP/ton)":       737,
		"Mileage City (km/L)":            738,
		"Mileage Highway (km/L)":         739,
		"Mileage Combined (km/L)":        740,
		"Front Suspension":               745,
		"Rear Suspension":                746,
		"Steering Type":                  747,
		"Steering Adjustment":            748,
		"Wheel Size (inch)":              749,
		"Spare Wheel Type":               750,
		"DRL":                            753,
		"Fog Lamp Type":                  754,
		"Alloy Wheels":                   755,
		"Sunroof Type":                   756,
		"Roof Rails":                     757,
		"ORVM Type":                      758,
		"Wiper Type":                     759,
		"Seating Capacity":               760,
		"Number of Seats":                761,
		"Driver Seat Adjustment":         762,
		"Ventilated Seats":               763,
		"Air Conditioning":               764,
		"Climate Control":                765,
		"Infotainment System":            766,
		"Infotainment Screen (inch)":     767,
		"Apple CarPlay":                  768,
		"Android Auto":                   769,
		"Sound System Brand":             770,
		"Number of Speakers":             771,
		"Ambient Lighting":               772,
		"ABS (Anti-lock Braking System)": 773,
		"Airbags":                        774,
		"EBD":                            775,
		"Brake Type":                     776,
		"Traction Control":               777,
		"ESC":                            778,
		"Hill Assist":                    779,
		"Power Steering":                 780,
		"Touchscreen":                    781,
		"Suspension Type":                782,
		"Emission Standard":              783,
		"Starting System":                784,
		"Cooling System":                 785,
		"Ignition Type":                  786,
		"Compression Ratio":              787,
		"Valve Configuration":            788,
		"Valve Per Cylinder":             789,
		"Displacement (cc)":              790,
	}

	specs := map[string]string{
		"Variant":                        "Titanium 1.5L Ti-VCT",
		"Generation":                     "1st",
		"Segment":                        "Compact SUV",
		"Launch Year":                    "2012",
		"Engine Configuration":           "In-line",
		"Displacement (cc)":              "1499",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "92 hp/ton",
		"Mileage City (km/L)":            "13 km/L",
		"Mileage Highway (km/L)":         "17 km/L",
		"Mileage Combined (km/L)":        "15 km/L",
		"Front Suspension":               "Independent McPherson Strut",
		"Rear Suspension":                "Twist Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "16",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Manual",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "6.5 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Sony",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver + Passenger + Side",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "No",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "Independent McPherson Strut / Twist Beam",
		"Emission Standard":              "BS6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "11.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
	}

	// Create specifications
	for key, value := range specs {
		keyID, exists := existingKeyMapping[key]
		if !exists {
			log.Printf("⚠️  Specification key not found: %s", key)
			continue
		}

		spec := models.SpecificationModel{
			ProductID:          productID,
			SpecificationKeyID: keyID,
			Value:              value,
		}

		// Check if specification already exists
		var existingSpec models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, keyID).First(&existingSpec).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new specification
				if err := db.Create(&spec).Error; err != nil {
					log.Printf("❌ Failed to create specification for key %s: %v", key, err)
					continue
				}
			} else {
				log.Printf("❌ Error checking existing specification for key %s: %v", key, err)
				continue
			}
		} else {
			log.Printf("ℹ️  Specification already exists for key %s, skipping", key)
			continue
		}

		// Create translation if Bangla translation exists
		if banglaValue, exists := fess.getBanglaTranslations()[value]; exists {
			translation := models.SpecificationTranslationModel{
				SpecificationID: spec.ID,
				Locale:          "bn",
				Value:           banglaValue,
			}

			if err := db.Create(&translation).Error; err != nil {
				log.Printf("❌ Failed to create translation for key %s: %v", key, err)
			}
		}
	}

	log.Printf("✅ Ford EcoSport specifications seeded successfully")
	return nil
}
