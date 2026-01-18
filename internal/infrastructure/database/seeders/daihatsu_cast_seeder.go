package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// DaihatsuCastSeeder seeds specifications for Daihatsu Cast
type DaihatsuCastSeeder struct {
	BaseSeeder
}

// NewDaihatsuCastSeeder creates a new Daihatsu Cast seeder
func NewDaihatsuCastSeeder() *DaihatsuCastSeeder {
	return &DaihatsuCastSeeder{
		BaseSeeder: BaseSeeder{name: "Daihatsu Cast Specifications"},
	}
}

func (dcs *DaihatsuCastSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.66L":                      "০.৬৬ লিটার",
		"2nd":                        "২য় প্রজন্ম",
		"A-Segment":                  "এ-সেগমেন্ট",
		"2015":                       "২০১৫",
		"Hatchback":                  "হ্যাচব্যাক",
		"White Pearl":                "হোয়াইট পার্ল",
		"Inline":                     "ইনলাইন",
		"658":                        "৬৫৮ সিসি",
		"3":                          "৩ (সিলিন্ডার)",
		"52 hp":                      "৫২ হর্স পাওয়ার",
		"52 hp @ 7200 rpm":           "৭২০০ আরপিএম এ ৫২ হর্স পাওয়ার",
		"60 Nm @ 4000 rpm":           "৪০০০ আরপিএম এ ৬০ এনএম",
		"Petrol":                     "পেট্রোল",
		"Naturally Aspirated":        "ন্যাচারালি অ্যাসপিরেটেড",
		"Multipoint Injection":       "মাল্টিপয়েন্ট ইনজেকশন",
		"14.5 seconds":               "১৪.৫ সেকেন্ড",
		"130 km/h":                   "১৩০ কিমি/ঘণ্টা",
		"19 km/L":                    "১৯ কিমি/লিটার",
		"24 km/L":                    "২৪ কিমি/লিটার",
		"21 km/L":                    "২১ কিমি/লিটার",
		"Automatic (Transmission)":   "অটোমেটিক (ট্রান্সমিশন)",
		"CVT":                        "সিভিটি",
		"Front-Wheel Drive":          "সামনের চাকা চালিত",
		"Single Clutch":              "একক ক্লাচ",
		"3395 mm":                    "৩৩৯৫ মিমি",
		"1475 mm":                    "১৪৭৫ মিমি",
		"1480 mm":                    "১৪৮০ মিমি",
		"2430 mm":                    "২৪৩০ মিমি",
		"1520 mm":                    "১৫২০ মিমি",
		"820 kg":                     "৮২০ কেজি",
		"30 L":                       "৩০ লিটার",
		"155/65 R14":                 "১৫৫/৬৫ আর১৪",
		"Radial Tubeless":            "রেডিয়াল টিউবলেস",
		"Steel":                      "স্টিল",
		"Halogen (Headlamps)":        "হ্যালোজেন (হেডল্যাম্পস)",
		"Yes (DRL)":                  "হ্যাঁ (ডিআরএল)",
		"Halogen (Fog Lamp)":         "হ্যালোজেন (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":             "হ্যাঁ (ফগ ল্যাম্প)",
		"Yes (Roof Rails)":           "হ্যাঁ (রুফ রেলস)",
		"Manual (ORVM)":              "ম্যানুয়াল (ওআরভিএম)",
		"Automatic (Wiper)":          "অটোমেটিক (ওয়াইপার)",
		"4 (Seating)":                "৪ (সিটিং)",
		"4 (Seats)":                  "৪ (সিটস)",
		"Manual (Driver Seat)":       "ম্যানুয়াল (ড্রাইভার সিট)",
		"No (Ventilated Seats)":      "না (ভেন্টিলেটেড সিটস)",
		"Manual (Air Conditioning)":  "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"No (Climate Control)":       "না (ক্লাইমেট কন্ট্রোল)",
		"Basic (Infotainment)":       "বেসিক (ইনফোটেইনমেন্ট)",
		"6.1 inch (Screen)":          "৬.১ ইঞ্চি (স্ক্রিন)",
		"No (Apple CarPlay)":         "না (অ্যাপল কারপ্লে)",
		"No (Android Auto)":          "না (অ্যান্ড্রয়েড অটো)",
		"Standard (Sound System)":    "স্ট্যান্ডার্ড (সাউন্ড সিস্টেম)",
		"2 (Speakers)":               "২ (স্পিকার)",
		"No (Ambient Lighting)":      "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                  "হ্যাঁ (এবিএস)",
		"Driver (Airbags)":           "ড্রাইভার (এয়ারব্যাগস)",
		"Yes (EBD)":                  "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Drum (Rear)": "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Yes (Traction Control)":     "হ্যাঁ (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                  "হ্যাঁ (ইএসসি)",
		"Yes (Hill Assist)":          "হ্যাঁ (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":       "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"No (Touchscreen)":           "না (টাচস্ক্রিন)",
		"Keyless Entry":              "কীলেস এন্ট্রি",
		"No (Height Adjustable)":     "না (হাইট অ্যাডজাস্টেবল)",
		"Rear AC Vents":              "রিয়ার এসি ভেন্টস",
		"Steering Mounted Controls":  "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"No (Parking Sensors)":       "না (পার্কিং সেন্সরস)",
		"No (Rear Camera)":           "না (রিয়ার ক্যামেরা)",
		"No (Rain Sensing Wipers)":   "না (রেইন সেন্সিং ওয়াইপারস)",
		"No (Follow Me Home)":        "না (ফলো মি হোম)",
		"No (Sunroof)":               "না (সানরুফ)",
		"No (Roof Rails)":            "না (রুফ রেলস)",
		"No (Alloy Wheels)":          "না (অ্যালয় হুইলস)",
		"14 inch (Wheel)":            "১৪ ইঞ্চি (হুইল)",
		"MacPherson Strut":           "ম্যাকফারসন স্ট্রাট",
		"Torsion Beam":               "টর্শন বিম",
		"Rack and Pinion":            "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt (Steering)":            "টিল্ট (স্টিয়ারিং)",
		"14 (Wheel Size)":            "১৪ (হুইল সাইজ)",
	}
}

// Seed implements the Seeder interface for Daihatsu Cast
func (dcs *DaihatsuCastSeeder) Seed(db *gorm.DB) error {
	productSlug := "daihatsu-cast"
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
		"Brand":                       "Daihatsu",
		"Model Name":                  "Cast",
		"Variant":                        "0.66L Style",
		"Generation":                     "2nd",
		"Segment":                        "A-Segment",
		"Launch Year":                    "2015",
		"Engine Configuration":           "Inline",
		"Displacement (cc)":              "658",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "63 hp/ton",
		"Mileage City (km/L)":            "19 km/L",
		"Mileage Highway (km/L)":         "24 km/L",
		"Mileage Combined (km/L)":        "21 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "14",
		"Spare Wheel Type":               "Steel",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Manual",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "4",
		"Number of Seats":                "4",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Basic",
		"Infotainment Screen (inch)":     "6.1 inch",
		"Apple CarPlay":                  "No",
		"Android Auto":                   "No",
		"Sound System Brand":             "Standard",
		"Number of Speakers":             "2",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "No",
		"Suspension Type":                "MacPherson Strut / Torsion Beam",
		"Emission Standard":              "Euro 5",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "11.5:1",
		"Valve Configuration":            "SOHC",
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
		if banglaValue, exists := dcs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Daihatsu Cast specifications seeded successfully")
	return nil
}
