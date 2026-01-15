package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AudiA4Seeder seeds specifications for Audi A4
type AudiA4Seeder struct {
	BaseSeeder
}

// NewAudiA4Seeder creates a new Audi A4 seeder
func NewAudiA4Seeder() *AudiA4Seeder {
	return &AudiA4Seeder{
		BaseSeeder: BaseSeeder{name: "Audi A4 Specifications"},
	}
}

func (aas *AudiA4Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L":                          "২.০ লিটার",
		"5th":                           "৫ম প্রজন্ম",
		"D-Segment":                     "ডি-সেগমেন্ট",
		"2016":                          "২০১৬",
		"Sedan":                         "সেডান",
		"Carbon Black":                  "কার্বন ব্ল্যাক",
		"Inline":                        "ইনলাইন",
		"1984":                          "১৯৮৪ সিসি",
		"4":                             "৪ (সিলিন্ডার)",
		"190 hp":                        "১৯০ হর্স পাওয়ার",
		"190 hp @ 4200 rpm":             "৪২০০ আরপিএম এ ১৯০ হর্স পাওয়ার",
		"320 Nm @ 1450-4200 rpm":        "১৪৫০-৪২০০ আরপিএম এ ৩২০ এনএম",
		"Petrol":                        "পেট্রোল",
		"Turbocharged":                  "টার্বোচার্জড",
		"7.3 seconds":                   "৭.৩ সেকেন্ড",
		"240 km/h":                      "২৪০ কিমি/ঘণ্টা",
		"14 km/L":                       "১৪ কিমি/লিটার",
		"18 km/L":                       "১৮ কিমি/লিটার",
		"16 km/L":                       "১৬ কিমি/লিটার",
		"Automatic (Transmission)":      "অটোমেটিক (ট্রান্সমিশন)",
		"7-Speed Automatic":             "৭-স্পিড অটোমেটিক",
		"Front-Wheel Drive":             "সামনের চাকা চালিত",
		"Single Plate Dry Clutch":       "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"4726 mm":                       "৪৭২৬ মিমি",
		"1842 mm":                       "১৮৪২ মিমি",
		"1427 mm":                       "১৪২৭ মিমি",
		"2820 mm":                       "২৮২০ মিমি",
		"165 mm":                        "১৬৫ মিমি",
		"1500 kg":                       "১৫০০ কেজি",
		"480 L":                         "৪৮০ লিটার",
		"225/50 R17":                    "২২৫/৫০ আর১৭",
		"Radial Tubeless":               "রেডিয়াল টিউবলেস",
		"Alloy":                         "অ্যালয়",
		"LED (Headlamps)":               "এলইডি (হেডল্যাম্পস)",
		"Yes (DRL)":                     "হ্যাঁ (ডিআরএল)",
		"LED (Fog Lamp)":                "এলইডি (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":                "হ্যাঁ (ফগ ল্যাম্প)",
		"Yes (Sunroof)":                 "হ্যাঁ (সানরুফ)",
		"No (Roof Rails)":               "না (রুফ রেলস)",
		"Power (ORVM)":                  "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":             "অটোমেটিক (ওয়াইপার)",
		"5 (Seating)":                   "৫ (সিটিং)",
		"5 (Seats)":                     "৫ (সিটস)",
		"Power (Driver Seat)":           "পাওয়ার (ড্রাইভার সিট)",
		"Yes (Ventilated Seats)":        "হ্যাঁ (ভেন্টিলেটেড সিটস)",
		"Automatic (Air Conditioning)":  "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Yes (Climate Control)":         "হ্যাঁ (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":    "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"8.3 inch (Screen)":             "৮.৩ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":           "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":            "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Bang & Olufsen (Sound System)": "ব্যাং অ্যান্ড অলুফসেন (সাউন্ড সিস্টেম)",
		"12 (Speakers)":                 "১২ (স্পিকার)",
		"Yes (Ambient Lighting)":        "হ্যাঁ (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                     "হ্যাঁ (এবিএস)",
		"Driver + Passenger + Side + Curtain (Airbags)": "ড্রাইভার + প্যাসেঞ্জার + সাইড + কার্টেন (এয়ারব্যাগস)",
		"Yes (EBD)":                       "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Disc (Rear)":      "ডিস্ক (সামনে) / ডিস্ক (পিছনে)",
		"Yes (Traction Control)":          "হ্যাঁ (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                       "হ্যাঁ (ইএসসি)",
		"Yes (Hill Assist)":               "হ্যাঁ (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":            "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":               "হ্যাঁ (টাচস্ক্রিন)",
		"Yes (Rear Camera)":               "হ্যাঁ (রিয়ার ক্যামেরা)",
		"Keyless Entry":                   "কীলেস এন্ট্রি",
		"Push Button Start":               "পুশ বাটন স্টার্ট",
		"Height Adjustable (Driver Seat)": "হাইট অ্যাডজাস্টেবল (ড্রাইভার সিট)",
		"Rear AC Vents":                   "রিয়ার এসি ভেন্টস",
		"Steering Mounted Controls":       "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"Cruise Control":                  "ক্রুজ কন্ট্রোল",
		"Front + Rear (Parking Sensors)":  "সামনে + পিছনে (পার্কিং সেন্সরস)",
		"Rear Camera":                     "রিয়ার ক্যামেরা",
		"Auto Headlamps":                  "অটো হেডল্যাম্পস",
		"Rain Sensing Wipers":             "রেইন সেন্সিং ওয়াইপারস",
		"Follow Me Home Headlamps":        "ফলো মি হোম হেডল্যাম্পস",
		"Yes (Alloy Wheels)":              "হ্যাঁ (অ্যালয় হুইলস)",
		"17 inch":                         "১৭ ইঞ্চি",
		"Multi-link (Front)":              "মাল্টি-লিঙ্ক (সামনে)",
		"Multi-link (Rear)":               "মাল্টি-লিঙ্ক (পিছনে)",
		"Rack and Pinion":                 "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt & Telescopic":               "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"17":                              "১৭",
	}
}

// Seed implements the Seeder interface for Audi A4
func (aas *AudiA4Seeder) Seed(db *gorm.DB) error {
	productSlug := "audi-a4"
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
		"Variant":                        "40 TFSI",
		"Generation":                     "5th",
		"Segment":                        "D-Segment",
		"Launch Year":                    "2016",
		"Engine Configuration":           "Inline",
		"Displacement (cc)":              "1984",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "127 hp/ton",
		"Mileage City (km/L)":            "14 km/L",
		"Mileage Highway (km/L)":         "18 km/L",
		"Mileage Combined (km/L)":        "16 km/L",
		"Front Suspension":               "Multi-link",
		"Rear Suspension":                "Multi-link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "17",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Yes",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "8.3 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Bang & Olufsen",
		"Number of Speakers":             "12",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver + Passenger + Side + Curtain",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Disc (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "Multi-link / Multi-link",
		"Emission Standard":              "BS6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.0:1",
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
		if banglaValue, exists := aas.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Audi A4 specifications seeded successfully")
	return nil
}
