package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HyundaiI20Seeder seeds specifications for Hyundai i20
type HyundaiI20Seeder struct {
	BaseSeeder
}

// NewHyundaiI20Seeder creates a new Hyundai i20 seeder
func NewHyundaiI20Seeder() *HyundaiI20Seeder {
	return &HyundaiI20Seeder{
		BaseSeeder: BaseSeeder{name: "Hyundai i20 Specifications"},
	}
}

func (his *HyundaiI20Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.2L":                            "১.২ লিটার",
		"3rd":                             "৩য় প্রজন্ম",
		"B-Segment":                       "বি-সেগমেন্ট",
		"2020":                            "২০২০",
		"Hatchback":                       "হ্যাচব্যাক",
		"Phantom Black":                   "ফ্যান্টম ব্ল্যাক",
		"Inline":                          "ইনলাইন",
		"1197":                            "১১৯৭ সিসি",
		"4":                               "৪ (সিলিন্ডার)",
		"83 hp":                           "৮৩ হর্স পাওয়ার",
		"83 hp @ 6000 rpm":                "৬০০০ আরপিএম এ ৮৩ হর্স পাওয়ার",
		"115 Nm @ 4200 rpm":               "৪২০০ আরপিএম এ ১১৫ এনএম",
		"Petrol":                          "পেট্রোল",
		"Naturally Aspirated":             "ন্যাচারালি অ্যাসপিরেটেড",
		"MPi":                             "এমপিআই",
		"13.2 seconds":                    "১৩.২ সেকেন্ড",
		"165 km/h":                        "১৬৫ কিমি/ঘণ্টা",
		"17 km/L":                         "১৭ কিমি/লিটার",
		"21 km/L":                         "২১ কিমি/লিটার",
		"19 km/L":                         "১৯ কিমি/লিটার",
		"Manual (Transmission)":           "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":                  "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":               "সামনের চাকা চালিত",
		"Single Plate Dry Clutch":         "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"3995 mm":                         "৩৯৯৫ মিমি",
		"1775 mm":                         "১৭৭৫ মিমি",
		"1505 mm":                         "১৫০৫ মিমি",
		"2580 mm":                         "২৫৮০ মিমি",
		"165 mm":                          "১৬৫ মিমি",
		"1050 kg":                         "১০৫০ কেজি",
		"295 L":                           "২৯৫ লিটার",
		"195/55 R16":                      "১৯৫/৫৫ আর১৬",
		"Radial Tubeless":                 "রেডিয়াল টিউবলেস",
		"Steel":                           "স্টিল",
		"Halogen (Headlamps)":             "হ্যালোজেন (হেডল্যাম্পস)",
		"No (DRL)":                        "না (ডিআরএল)",
		"Halogen (Fog Lamp)":              "হ্যালোজেন (ফগ ল্যাম্প)",
		"No (Fog Lamp)":                   "না (ফগ ল্যাম্প)",
		"No (Sunroof)":                    "না (সানরুফ)",
		"No (Roof Rails)":                 "না (রুফ রেলস)",
		"Manual (ORVM)":                   "ম্যানুয়াল (ওআরভিএম)",
		"Manual (Wiper)":                  "ম্যানুয়াল (ওয়াইপার)",
		"5 (Seating)":                     "৫ (সিটিং)",
		"5 (Seats)":                       "৫ (সিটস)",
		"Manual (Driver Seat)":            "ম্যানুয়াল (ড্রাইভার সিট)",
		"No (Ventilated Seats)":           "না (ভেন্টিলেটেড সিটস)",
		"Manual (Air Conditioning)":       "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"No (Climate Control)":            "না (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":      "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"8 inch (Screen)":                 "৮ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":             "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":              "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Standard (Sound System)":         "স্ট্যান্ডার্ড (সাউন্ড সিস্টেম)",
		"4 (Speakers)":                    "৪ (স্পিকার)",
		"No (Ambient Lighting)":           "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                       "হ্যাঁ (এবিএস)",
		"Driver + Passenger (Airbags)":    "ড্রাইভার + প্যাসেঞ্জার (এয়ারব্যাগস)",
		"Yes (EBD)":                       "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Drum (Rear)":      "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"No (Traction Control)":           "না (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                       "হ্যাঁ (ইএসসি)",
		"No (Hill Assist)":                "না (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":            "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":               "হ্যাঁ (টাচস্ক্রিন)",
		"No (Rear Camera)":                "না (রিয়ার ক্যামেরা)",
		"Keyless Entry":                   "কীলেস এন্ট্রি",
		"Push Button Start":               "পুশ বাটন স্টার্ট",
		"Height Adjustable (Driver Seat)": "হাইট অ্যাডজাস্টেবল (ড্রাইভার সিট)",
		"No (Rear AC Vents)":              "না (রিয়ার এসি ভেন্টস)",
		"Steering Mounted Controls":       "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"No (Cruise Control)":             "না (ক্রুজ কন্ট্রোল)",
		"Rear (Parking Sensors)":          "পিছনে (পার্কিং সেন্সরস)",
		"No (Auto Headlamps)":             "না (অটো হেডল্যাম্পস)",
		"No (Rain Sensing Wipers)":        "না (রেইন সেন্সিং ওয়াইপারস)",
		"No (Follow Me Home Headlamps)":   "না (ফলো মি হোম হেডল্যাম্পস)",
		"No (Alloy Wheels)":               "না (অ্যালয় হুইলস)",
		"16 inch":                         "১৬ ইঞ্চি",
		"MacPherson Strut":                "ম্যাকফারসন স্ট্রাট",
		"Coupled Torsion Beam":            "কাপলড টর্শন বিম",
		"Rack and Pinion":                 "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt":                            "টিল্ট",
		"16":                              "১৬",
	}
}

// Seed implements the Seeder interface for Hyundai i20
func (his *HyundaiI20Seeder) Seed(db *gorm.DB) error {
	productSlug := "hyundai-i20"
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
		"Brand":                       "Hyundai",
		"Model Name":                  "i20",
		"Variant":                        "1.2L MPi",
		"Generation":                     "3rd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2020",
		"Engine Configuration":           "Inline",
		"Displacement (cc)":              "1197",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "79 hp/ton",
		"Mileage City (km/L)":            "17 km/L",
		"Mileage Highway (km/L)":         "21 km/L",
		"Mileage Combined (km/L)":        "19 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Coupled Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "16",
		"Spare Wheel Type":               "Steel",
		"DRL":                            "No",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Manual",
		"Wiper Type":                     "Manual",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Standard",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver + Passenger",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "No",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Coupled Torsion Beam",
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
		if banglaValue, exists := his.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Hyundai i20 specifications seeded successfully")
	return nil
}
