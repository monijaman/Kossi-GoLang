package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ChanganCs55Seeder seeds specifications for Changan CS55
type ChanganCs55Seeder struct {
	BaseSeeder
}

// NewChanganCs55Seeder creates a new Changan CS55 seeder
func NewChanganCs55Seeder() *ChanganCs55Seeder {
	return &ChanganCs55Seeder{
		BaseSeeder: BaseSeeder{name: "Changan CS55 Specifications"},
	}
}

func (ccs *ChanganCs55Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L":                         "১.৫ লিটার",
		"1st":                          "১ম প্রজন্ম",
		"Mid-size SUV":                 "মিড-সাইজ এসইউভি",
		"2019":                         "২০১৯",
		"SUV":                          "এসইউভি",
		"White":                        "সাদা",
		"In-line":                      "ইন-লাইন",
		"1499":                         "১৪৯৯ সিসি",
		"4":                            "৪ (সিলিন্ডার)",
		"156 hp":                       "১৫৬ হর্স পাওয়ার",
		"156 hp @ 5500 rpm":            "৫৫০০ আরপিএম এ ১৫৬ হর্স পাওয়ার",
		"225 Nm @ 2000 rpm":            "২০০০ আরপিএম এ ২২৫ এনএম",
		"Petrol":                       "পেট্রোল",
		"Turbocharged":                 "টার্বোচার্জড",
		"10.2 seconds":                 "১০.২ সেকেন্ড",
		"180 km/h":                     "১৮০ কিমি/ঘণ্টা",
		"10 km/L":                      "১০ কিমি/লিটার",
		"14 km/L":                      "১৪ কিমি/লিটার",
		"12 km/L":                      "১২ কিমি/লিটার",
		"Automatic (Transmission)":     "অটোমেটিক (ট্রান্সমিশন)",
		"6-Speed Automatic":            "৬-স্পিড অটোমেটিক",
		"Front-Wheel Drive":            "সামনে চাকা চালিত",
		"Single Plate Dry Clutch":      "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"4515 mm":                      "৪৫১৫ মিমি",
		"1860 mm":                      "১৮৬০ মিমি",
		"1690 mm":                      "১৬৯০ মিমি",
		"2650 mm":                      "২৬৫০ মিমি",
		"190 mm":                       "১৯০ মিমি",
		"1500 kg":                      "১৫০০ কেজি",
		"378 L":                        "৩৭৮ লিটার",
		"225/60 R18":                   "২২৫/৬০ আর১৮",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"LED (Headlamps)":              "এলইডি (হেডল্যাম্পস)",
		"Yes (DRL)":                    "হ্যাঁ (ডিআরএল)",
		"LED (Fog Lamp)":               "এলইডি (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":               "হ্যাঁ (ফগ ল্যাম্প)",
		"No (Sunroof)":                 "না (সানরুফ)",
		"Yes (Roof Rails)":             "হ্যাঁ (রুফ রেলস)",
		"Power (ORVM)":                 "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":            "অটোমেটিক (ওয়াইপার)",
		"5 (Seating)":                  "৫ (সিটিং)",
		"5 (Seats)":                    "৫ (সিটস)",
		"Manual (Driver Seat)":         "ম্যানুয়াল (ড্রাইভার সিট)",
		"No (Ventilated Seats)":        "না (ভেন্টিলেটেড সিটস)",
		"Automatic (Air Conditioning)": "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Yes (Climate Control)":        "হ্যাঁ (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":   "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"10.25 inch (Screen)":          "১০.২৫ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":          "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":           "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Changan (Sound System)":       "চ্যাংগান (সাউন্ড সিস্টেম)",
		"6 (Speakers)":                 "৬ (স্পিকার)",
		"No (Ambient Lighting)":        "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                    "হ্যাঁ (এবিএস)",
		"Driver + Passenger + Side + Curtain (Airbags)": "ড্রাইভার + প্যাসেঞ্জার + সাইড + কার্টেন (এয়ারব্যাগস)",
		"Yes (EBD)":                          "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Disc (Rear)":         "ডিস্ক (সামনে) / ডিস্ক (পিছনে)",
		"No (Traction Control)":              "না (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                          "হ্যাঁ (ইএসসি)",
		"No (Hill Assist)":                   "না (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":               "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":                  "হ্যাঁ (টাচস্ক্রিন)",
		"Yes (Rear Camera)":                  "হ্যাঁ (রিয়ার ক্যামেরা)",
		"Yes (Keyless Entry)":                "হ্যাঁ (কীলেস এন্ট্রি)",
		"No (Push Button Start)":             "না (পুশ বাটন স্টার্ট)",
		"No (Height Adjustable Driver Seat)": "না (হাইট অ্যাডজাস্টেবল ড্রাইভার সিট)",
		"Yes (Rear AC Vents)":                "হ্যাঁ (রিয়ার এসি ভেন্টস)",
		"Yes (Steering Mounted Controls)":    "হ্যাঁ (স্টিয়ারিং মাউন্টেড কন্ট্রোলস)",
		"No (Cruise Control)":                "না (ক্রুজ কন্ট্রোল)",
		"No (Front Parking Sensors)":         "না (সামনে পার্কিং সেন্সরস)",
		"No (Rear Parking Sensors)":          "না (রিয়ার পার্কিং সেন্সরস)",
		"No (Auto Headlamps)":                "না (অটো হেডল্যাম্পস)",
		"No (Rain Sensing Wipers)":           "না (রেইন সেন্সিং ওয়াইপারস)",
		"No (Follow Me Home Headlamps)":      "না (ফলো মি হোম হেডল্যাম্পস)",
		"Yes (Alloy Wheels)":                 "হ্যাঁ (অ্যালয় হুইলস)",
		"18 inch":                            "১৮ ইঞ্চি",
		"Independent McPherson Strut":        "ইন্ডিপেন্ডেন্ট ম্যাকফারসন স্ট্রাট",
		"Multi-link":                         "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                    "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt & Telescopic":                  "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"18":                                 "১৮",
	}
}

// Seed implements the Seeder interface for Changan CS55
func (ccs *ChanganCs55Seeder) Seed(db *gorm.DB) error {
	productSlug := "changan-cs55"
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
		"Variant":                        "Plus 1.5T AT",
		"Generation":                     "1st",
		"Segment":                        "Mid-size SUV",
		"Launch Year":                    "2019",
		"Engine Configuration":           "In-line",
		"Displacement (cc)":              "1499",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "104 hp/ton",
		"Mileage City (km/L)":            "10 km/L",
		"Mileage Highway (km/L)":         "14 km/L",
		"Mileage Combined (km/L)":        "12 km/L",
		"Front Suspension":               "Independent McPherson Strut",
		"Rear Suspension":                "Multi-link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "18",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "10.25 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Changan",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver + Passenger + Side + Curtain",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Disc (Rear)",
		"Traction Control":               "No",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "Independent McPherson Strut / Multi-link",
		"Emission Standard":              "BS6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "9.5:1",
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
		if banglaValue, exists := ccs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Changan CS55 specifications seeded successfully")
	return nil
}
