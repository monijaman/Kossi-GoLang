package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// RangeRoverAutobiographySeeder seeds specifications for Range Rover Autobiography
type RangeRoverAutobiographySeeder struct {
	BaseSeeder
}

// NewRangeRoverAutobiographySeeder creates a new Range Rover Autobiography seeder
func NewRangeRoverAutobiographySeeder() *RangeRoverAutobiographySeeder {
	return &RangeRoverAutobiographySeeder{
		BaseSeeder: BaseSeeder{name: "Range Rover Autobiography Specifications"},
	}
}

func (rras *RangeRoverAutobiographySeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.0L":                         "৩.০ লিটার",
		"5th":                          "৫ম প্রজন্ম",
		"Full-size SUV":                "ফুল-সাইজ এসইউভি",
		"2012":                         "২০১২",
		"SUV":                          "এসইউভি",
		"Carpathian Grey":              "কারপ্যাথিয়ান গ্রে",
		"V6":                           "ভি৬",
		"2993":                         "২৯৯৩ সিসি",
		"4":                            "৪ (সিলিন্ডার)",
		"375 hp":                       "৩৭৫ হর্স পাওয়ার",
		"375 hp @ 6500 rpm":            "৬৫০০ আরপিএম এ ৩৭৫ হর্স পাওয়ার",
		"525 Nm @ 3500 rpm":            "৩৫০০ আরপিএম এ ৫২৫ এনএম",
		"Diesel":                       "ডিজেল",
		"Turbocharged":                 "টার্বোচার্জড",
		"7.9 seconds":                  "৭.৯ সেকেন্ড",
		"218 km/h":                     "২১৮ কিমি/ঘণ্টা",
		"8 km/L":                       "৮ কিমি/লিটার",
		"12 km/L":                      "১২ কিমি/লিটার",
		"10 km/L":                      "১০ কিমি/লিটার",
		"Automatic (Transmission)":     "অটোমেটিক (ট্রান্সমিশন)",
		"8-Speed Automatic":            "৮-স্পিড অটোমেটিক",
		"All-Wheel Drive":              "সমস্ত চাকা চালিত",
		"Single Plate Dry Clutch":      "সিঙ্গেল প্লেট ড্রাই ক্লাচ",
		"4999 mm":                      "৪৯৯৯ মিমি",
		"2073 mm":                      "২০৭৩ মিমি",
		"1835 mm":                      "১৮৩৫ মিমি",
		"2922 mm":                      "২৯২২ মিমি",
		"220 mm":                       "২২০ মিমি",
		"2330 kg":                      "২৩৩০ কেজি",
		"900 L":                        "৯০০ লিটার",
		"275/50 R20":                   "২৭৫/৫০ আর২০",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"LED (Headlamps)":              "এলইডি (হেডল্যাম্পস)",
		"Yes (DRL)":                    "হ্যাঁ (ডিআরএল)",
		"LED (Fog Lamp)":               "এলইডি (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":               "হ্যাঁ (ফগ ল্যাম্প)",
		"Yes (Sunroof)":                "হ্যাঁ (সানরুফ)",
		"Yes (Roof Rails)":             "হ্যাঁ (রুফ রেলস)",
		"Power (ORVM)":                 "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":            "অটোমেটিক (ওয়াইপার)",
		"5 (Seating)":                  "৫ (সিটিং)",
		"5 (Seats)":                    "৫ (সিটস)",
		"Power (Driver Seat)":          "পাওয়ার (ড্রাইভার সিট)",
		"Yes (Ventilated Seats)":       "হ্যাঁ (ভেন্টিলেটেড সিটস)",
		"Automatic (Air Conditioning)": "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Yes (Climate Control)":        "হ্যাঁ (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":   "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"12.3 inch (Screen)":           "১২.৩ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":          "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":           "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Range Rover (Sound System)":   "রেঞ্জ রোভার (সাউন্ড সিস্টেম)",
		"29 (Speakers)":                "২৯ (স্পিকার)",
		"Yes (Ambient Lighting)":       "হ্যাঁ (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                    "হ্যাঁ (এবিএস)",
		"Driver + Passenger + Side + Curtain (Airbags)": "ড্রাইভার + প্যাসেঞ্জার + সাইড + কার্টেন (এয়ারব্যাগস)",
		"Yes (EBD)":                           "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Disc (Rear)":          "ডিস্ক (সামনে) / ডিস্ক (পিছনে)",
		"Yes (Traction Control)":              "হ্যাঁ (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                           "হ্যাঁ (ইএসসি)",
		"Yes (Hill Assist)":                   "হ্যাঁ (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":                "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":                   "হ্যাঁ (টাচস্ক্রিন)",
		"Yes (Rear Camera)":                   "হ্যাঁ (রিয়ার ক্যামেরা)",
		"Keyless Entry":                       "কীলেস এন্ট্রি",
		"Yes (Push Button Start)":             "হ্যাঁ (পুশ বাটন স্টার্ট)",
		"Yes (Height Adjustable Driver Seat)": "হ্যাঁ (হাইট অ্যাডজাস্টেবল ড্রাইভার সিট)",
		"Yes (Rear AC Vents)":                 "হ্যাঁ (রিয়ার এসি ভেন্টস)",
		"Yes (Steering Mounted Controls)":     "হ্যাঁ (স্টিয়ারিং মাউন্টেড কন্ট্রোলস)",
		"Yes (Cruise Control)":                "হ্যাঁ (ক্রুজ কন্ট্রোল)",
		"Yes (Front Parking Sensors)":         "হ্যাঁ (সামনে পার্কিং সেন্সরস)",
		"Yes (Rear Parking Sensors)":          "হ্যাঁ (রিয়ার পার্কিং সেন্সরস)",
		"Yes (Auto Headlamps)":                "হ্যাঁ (অটো হেডল্যাম্পস)",
		"Yes (Rain Sensing Wipers)":           "হ্যাঁ (রেইন সেন্সিং ওয়াইপারস)",
		"Yes (Follow Me Home Headlamps)":      "হ্যাঁ (ফলো মি হোম হেডল্যাম্পস)",
		"Yes (Alloy Wheels)":                  "হ্যাঁ (অ্যালয় হুইলস)",
		"20 inch":                             "২০ ইঞ্চি",
		"Double Wishbone":                     "ডাবল উইশবোন",
		"Multi-link":                          "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                     "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt & Telescopic":                   "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"20":                                  "২০",
	}
}

// Seed implements the Seeder interface for Range Rover Autobiography
func (rras *RangeRoverAutobiographySeeder) Seed(db *gorm.DB) error {
	productSlug := "range-rover-autobiography"
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
		"Brand":                             310,
		"Model Name":                        316,
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
		"Brand":                       "Range Rover",
		"Model Name":                  "Autobiography",
		"Variant":                        "3.0 SDV6 Autobiography",
		"Generation":                     "5th",
		"Segment":                        "Full-size SUV",
		"Launch Year":                    "2012",
		"Engine Configuration":           "V6",
		"Displacement (cc)":              "2993",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "161 hp/ton",
		"Mileage City (km/L)":            "8 km/L",
		"Mileage Highway (km/L)":         "12 km/L",
		"Mileage Combined (km/L)":        "10 km/L",
		"Front Suspension":               "Double Wishbone",
		"Rear Suspension":                "Multi-link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "20",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Yes",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "12.3 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Range Rover",
		"Number of Speakers":             "29",
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
		"Compression Ratio":              "16.1:1",
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
		if banglaValue, exists := rras.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Range Rover Autobiography specifications seeded successfully")
	return nil
}
