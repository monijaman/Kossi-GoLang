package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// BmwI3Seeder seeds specifications for BMW i3
type BmwI3Seeder struct {
	BaseSeeder
}

// NewBmwI3Seeder creates a new BMW i3 seeder
func NewBmwI3Seeder() *BmwI3Seeder {
	return &BmwI3Seeder{
		BaseSeeder: BaseSeeder{name: "BMW i3 Specifications"},
	}
}

func (bi3s *BmwI3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Electric":                     "ইলেকট্রিক",
		"1st":                          "১ম প্রজন্ম",
		"A-Hatchback":                  "এ-হ্যাচব্যাক",
		"2014":                         "২০১৪",
		"Hatchback":                    "হ্যাচব্যাক",
		"White":                        "সাদা",
		"Single Motor":                 "সিঙ্গেল মোটর",
		"42 kWh":                       "৪২ কিলোওয়াট-ঘণ্টা",
		"170 hp":                       "১৭০ হর্স পাওয়ার",
		"250 Nm":                       "২৫০ এনএম",
		"7.3 seconds":                  "৭.৩ সেকেন্ড",
		"150 km/h":                     "১৫০ কিমি/ঘণ্টা",
		"470 km":                       "৪৭০ কিমি",
		"Single Speed":                 "সিঙ্গেল স্পিড",
		"Rear-Wheel Drive":             "পিছনের চাকা চালিত",
		"3979 mm":                      "৩৯৭৯ মিমি",
		"1775 mm":                      "১৭৭৫ মিমি",
		"1578 mm":                      "১৫৭৮ মিমি",
		"2570 mm":                      "২৫৭০ মিমি",
		"157 mm":                       "১৫৭ মিমি",
		"1195 kg":                      "১১৯৫ কেজি",
		"260 L":                        "২৬০ লিটার",
		"155/70 R19":                   "১৫৫/৭০ আর১৯",
		"Alloy":                        "অ্যালয়",
		"LED (Headlamps)":              "এলইডি (হেডল্যাম্পস)",
		"Yes (DRL)":                    "হ্যাঁ (ডিআরএল)",
		"No (Fog Lamp)":                "না (ফগ ল্যাম্প)",
		"No (Sunroof)":                 "না (সানরুফ)",
		"No (Roof Rails)":              "না (রুফ রেলস)",
		"Power (ORVM)":                 "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":            "অটোমেটিক (ওয়াইপার)",
		"4 (Seating)":                  "৪ (সিটিং)",
		"4 (Seats)":                    "৪ (সিটস)",
		"Power (Driver Seat)":          "পাওয়ার (ড্রাইভার সিট)",
		"No (Ventilated Seats)":        "না (ভেন্টিলেটেড সিটস)",
		"Automatic (Air Conditioning)": "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Yes (Climate Control)":        "হ্যাঁ (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":   "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"10.2 inch (Screen)":           "১০.২ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":          "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":           "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Harman Kardon (Sound System)": "হারম্যান কার্ডন (সাউন্ড সিস্টেম)",
		"4 (Speakers)":                 "৪ (স্পিকার)",
		"Yes (Ambient Lighting)":       "হ্যাঁ (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                    "হ্যাঁ (এবিএস)",
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
		"No (Rear AC Vents)":              "না (রিয়ার এসি ভেন্টস)",
		"Steering Mounted Controls":       "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"Cruise Control":                  "ক্রুজ কন্ট্রোল",
		"Front + Rear (Parking Sensors)":  "সামনে + পিছনে (পার্কিং সেন্সরস)",
		"Rear Camera":                     "রিয়ার ক্যামেরা",
		"Auto Headlamps":                  "অটো হেডল্যাম্পস",
		"Rain Sensing Wipers":             "রেইন সেন্সিং ওয়াইপারস",
		"Follow Me Home Headlamps":        "ফলো মি হোম হেডল্যাম্পস",
		"Yes (Alloy Wheels)":              "হ্যাঁ (অ্যালয় হুইলস)",
		"19 inch":                         "১৯ ইঞ্চি",
		"MacPherson Strut":                "ম্যাকফারসন স্ট্রাট",
		"Multi-link":                      "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                 "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt & Telescopic":               "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"19":                              "১৯",
	}
}

// Seed implements the Seeder interface for BMW i3
func (bi3s *BmwI3Seeder) Seed(db *gorm.DB) error {
	productSlug := "bmw-i3"
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
		"Brand":                       "BMW",
		"Model Name":                  "i3",
		"Variant":                        "120 Ah",
		"Generation":                     "1st",
		"Segment":                        "A-Hatchback",
		"Launch Year":                    "2014",
		"Engine Configuration":           "Single Motor",
		"Displacement (cc)":              "Electric",
		"Valves Per Cylinder":            "N/A",
		"Engine Aspiration":              "N/A",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "142 hp/ton",
		"Mileage City (km/L)":            "N/A",
		"Mileage Highway (km/L)":         "N/A",
		"Mileage Combined (km/L)":        "N/A",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Multi-link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "19",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "No",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "4",
		"Number of Seats":                "4",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "10.2 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Harman Kardon",
		"Number of Speakers":             "4",
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
		"Suspension Type":                "MacPherson Strut / Multi-link",
		"Emission Standard":              "Zero Emission",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "N/A",
		"Compression Ratio":              "N/A",
		"Valve Configuration":            "N/A",
		"Valve Per Cylinder":             "N/A",
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
		if banglaValue, exists := bi3s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ BMW i3 specifications seeded successfully")
	return nil
}
