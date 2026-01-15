package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// Bmw7SeriesSeeder seeds specifications for BMW 7 Series
type Bmw7SeriesSeeder struct {
	BaseSeeder
}

// NewBmw7SeriesSeeder creates a new BMW 7 Series seeder
func NewBmw7SeriesSeeder() *Bmw7SeriesSeeder {
	return &Bmw7SeriesSeeder{
		BaseSeeder: BaseSeeder{name: "BMW 7 Series Specifications"},
	}
}

func (b7ss *Bmw7SeriesSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.0L":                            "৩.০ লিটার",
		"6th":                             "৬ষ্ঠ প্রজন্ম",
		"F-Segment":                       "এফ-সেগমেন্ট",
		"2016":                            "২০১৬",
		"Sedan":                           "সেডান",
		"Carbon Black":                    "কার্বন ব্ল্যাক",
		"Inline":                          "ইনলাইন",
		"2998":                            "২৯৯৮ সিসি",
		"6":                               "৬ (সিলিন্ডার)",
		"320 hp":                          "৩২০ হর্স পাওয়ার",
		"320 hp @ 5500 rpm":               "৫৫০০ আরপিএম এ ৩২০ হর্স পাওয়ার",
		"450 Nm @ 1380-5200 rpm":          "১৩৮০-৫২০০ আরপিএম এ ৪৫০ এনএম",
		"Petrol":                          "পেট্রোল",
		"Turbocharged":                    "টার্বোচার্জড",
		"TwinPower Turbo":                 "টুইনপাওয়ার টার্বো",
		"5.6 seconds":                     "৫.৬ সেকেন্ড",
		"250 km/h":                        "২৫০ কিমি/ঘণ্টা",
		"9 km/L":                          "৯ কিমি/লিটার",
		"13 km/L":                         "১৩ কিমি/লিটার",
		"11 km/L":                         "১১ কিমি/লিটার",
		"Automatic (Transmission)":        "অটোমেটিক (ট্রান্সমিশন)",
		"8-Speed Automatic":               "৮-স্পিড অটোমেটিক",
		"Rear-Wheel Drive":                "পিছনের চাকা চালিত",
		"Single Plate Dry Clutch":         "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"5238 mm":                         "৫২৩৮ মিমি",
		"1902 mm":                         "১৯০২ মিমি",
		"1481 mm":                         "১৪৮১ মিমি",
		"3210 mm":                         "৩২১০ মিমি",
		"120 mm":                          "১২০ মিমি",
		"1940 kg":                         "১৯৪০ কেজি",
		"515 L":                           "৫১৫ লিটার",
		"245/50 R18":                      "২৪৫/৫০ আর১৮",
		"Radial Tubeless":                 "রেডিয়াল টিউবলেস",
		"Alloy":                           "অ্যালয়",
		"LED (Headlamps)":                 "এলইডি (হেডল্যাম্পস)",
		"Yes (DRL)":                       "হ্যাঁ (ডিআরএল)",
		"LED (Fog Lamp)":                  "এলইডি (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":                  "হ্যাঁ (ফগ ল্যাম্প)",
		"Yes (Sunroof)":                   "হ্যাঁ (সানরুফ)",
		"No (Roof Rails)":                 "না (রুফ রেলস)",
		"Power (ORVM)":                    "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":               "অটোমেটিক (ওয়াইপার)",
		"5 (Seating)":                     "৫ (সিটিং)",
		"5 (Seats)":                       "৫ (সিটস)",
		"Power (Driver Seat)":             "পাওয়ার (ড্রাইভার সিট)",
		"Yes (Ventilated Seats)":          "হ্যাঁ (ভেন্টিলেটেড সিটস)",
		"Automatic (Air Conditioning)":    "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Yes (Climate Control)":           "হ্যাঁ (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":      "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"10.25 inch (Screen)":             "১০.২৫ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":             "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":              "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Bowers & Wilkins (Sound System)": "বাওয়ার্স অ্যান্ড উইলকিন্স (সাউন্ড সিস্টেম)",
		"16 (Speakers)":                   "১৬ (স্পিকার)",
		"Yes (Ambient Lighting)":          "হ্যাঁ (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                       "হ্যাঁ (এবিএস)",
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
		"18 inch":                         "১৮ ইঞ্চি",
		"Double Wishbone":                 "ডাবল উইশবোন",
		"Multi-link":                      "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                 "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt & Telescopic":               "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"18":                              "১৮",
	}
}

// Seed implements the Seeder interface for BMW 7 Series
func (b7ss *Bmw7SeriesSeeder) Seed(db *gorm.DB) error {
	productSlug := "bmw-7-series"
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
		"Variant":                        "730Ld",
		"Generation":                     "6th",
		"Segment":                        "F-Segment",
		"Launch Year":                    "2016",
		"Engine Configuration":           "Inline",
		"Displacement (cc)":              "2998",
		"Valves Per Cylinder":            "6",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "165 hp/ton",
		"Mileage City (km/L)":            "9 km/L",
		"Mileage Highway (km/L)":         "13 km/L",
		"Mileage Combined (km/L)":        "11 km/L",
		"Front Suspension":               "Double Wishbone",
		"Rear Suspension":                "Multi-link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "18",
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
		"Infotainment Screen (inch)":     "10.25 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Bowers & Wilkins",
		"Number of Speakers":             "16",
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
		"Suspension Type":                "Double Wishbone / Multi-link",
		"Emission Standard":              "BS6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "16.5:1",
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
		if banglaValue, exists := b7ss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ BMW 7 Series specifications seeded successfully")
	return nil
}
