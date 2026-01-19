package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SubaruXVSeeder seeds specifications for Subaru XV
type SubaruXVSeeder struct {
	BaseSeeder
}

// NewSubaruXVSeeder creates a new Subaru XV seeder
func NewSubaruXVSeeder() *SubaruXVSeeder {
	return &SubaruXVSeeder{
		BaseSeeder: BaseSeeder{name: "Subaru XV Specifications"},
	}
}

func (sxs *SubaruXVSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L":                         "২.০ লিটার",
		"3rd":                          "৩য় প্রজন্ম",
		"C-Segment":                    "সি-সেগমেন্ট",
		"2017":                         "২০১৭",
		"Crossover SUV":                "ক্রসওভার এসইউভি",
		"Crystal Black Silica":         "ক্রিস্টাল ব্ল্যাক সিলিকা",
		"Boxer":                        "বক্সার",
		"1995":                         "১৯৯৫ সিসি",
		"4":                            "৪ (সিলিন্ডার)",
		"150 hp":                       "১৫০ হর্স পাওয়ার",
		"150 hp @ 6200 rpm":            "৬২০০ আরপিএম এ ১৫০ হর্স পাওয়ার",
		"196 Nm @ 4200 rpm":            "৪২০০ আরপিএম এ ১৯৬ এনএম",
		"Petrol":                       "পেট্রোল",
		"Naturally Aspirated":          "ন্যাচারালি অ্যাসপিরেটেড",
		"Direct Injection":             "সরাসরি ইনজেকশন",
		"10.6 seconds":                 "১০.৬ সেকেন্ড",
		"187 km/h":                     "১৮৭ কিমি/ঘণ্টা",
		"13 km/L":                      "১৩ কিমি/লিটার",
		"15 km/L":                      "১৫ কিমি/লিটার",
		"14 km/L":                      "১৪ কিমি/লিটার",
		"CVT (Transmission)":           "সিভিটি (ট্রান্সমিশন)",
		"Lineartronic CVT":             "লিনিয়ারট্রনিক সিভিটি",
		"All-Wheel Drive":              "সমস্ত চাকা চালিত",
		"Multi-plate Clutch":           "মাল্টি-প্লেট ক্লাচ",
		"4465 mm":                      "৪৪৬৫ মিমি",
		"1800 mm":                      "১৮০০ মিমি",
		"1595 mm":                      "১৫৯৫ মিমি",
		"2665 mm":                      "২৬৬৫ মিমি",
		"1615 mm":                      "১৬১৫ মিমি",
		"1530 kg":                      "১৫৩০ কেজি",
		"50 L":                         "৫০ লিটার",
		"225/55 R17":                   "২২৫/৫৫ আর১৭",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"Halogen (Headlamps)":          "হ্যালোজেন (হেডল্যাম্পস)",
		"Yes (DRL)":                    "হ্যাঁ (ডিআরএল)",
		"Halogen (Fog Lamp)":           "হ্যালোজেন (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":               "হ্যাঁ (ফগ ল্যাম্প)",
		"No (Sunroof)":                 "না (সানরুফ)",
		"Yes (Roof Rails)":             "হ্যাঁ (রুফ রেলস)",
		"Power (ORVM)":                 "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":            "অটোমেটিক (ওয়াইপার)",
		"5 (Seating)":                  "৫ (সিটিং)",
		"5 (Seats)":                    "৫ (সিটস)",
		"Manual (Driver Seat)":         "ম্যানুয়াল (ড্রাইভার সিট)",
		"No (Ventilated Seats)":        "না (ভেন্টিলেটেড সিটস)",
		"Manual (Air Conditioning)":    "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"No (Climate Control)":         "না (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":   "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"6.2 inch (Screen)":            "৬.২ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":          "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":           "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Standard (Sound System)":      "স্ট্যান্ডার্ড (সাউন্ড সিস্টেম)",
		"4 (Speakers)":                 "৪ (স্পিকার)",
		"No (Ambient Lighting)":        "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                    "হ্যাঁ (এবিএস)",
		"Driver + Passenger (Airbags)": "ড্রাইভার + প্যাসেঞ্জার (এয়ারব্যাগস)",
		"Yes (EBD)":                    "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Drum (Rear)":   "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Yes (Traction Control)":       "হ্যাঁ (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                    "হ্যাঁ (ইএসসি)",
		"Yes (Hill Assist)":            "হ্যাঁ (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":         "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":            "হ্যাঁ (টাচস্ক্রিন)",
		"Yes (Rear Camera)":            "হ্যাঁ (রিয়ার ক্যামেরা)",
		"Keyless Entry":                "কীলেস এন্ট্রি",
		"Push Button Start":            "পুশ বাটন স্টার্ট",
		"Height Adjustable":            "হাইট অ্যাডজাস্টেবল",
		"Rear AC Vents":                "রিয়ার এসি ভেন্টস",
		"Steering Mounted Controls":    "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"Cruise Control":               "ক্রুজ কন্ট্রোল",
		"Parking Sensors":              "পার্কিং সেন্সরস",
		"Camera":                       "ক্যামেরা",
		"Rear Camera (Feature)":        "রিয়ার ক্যামেরা (ফিচার)",
		"Auto Headlamps":               "অটো হেডল্যাম্পস",
		"Rain Sensing Wipers":          "রেইন সেন্সিং ওয়াইপারস",
		"Follow Me Home Headlamps":     "ফলো মি হোম হেডল্যাম্পস",
		"Yes (Alloy Wheels)":           "হ্যাঁ (অ্যালয় হুইলস)",
		"17 inch":                      "১৭ ইঞ্চি",
		"MacPherson Strut":             "ম্যাকফারসন স্ট্রাট",
		"Double Wishbone":              "ডাবল উইশবোন",
		"Rack and Pinion":              "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt":                         "টিল্ট",
		"17":                           "১৭",
	}
}

// Seed implements the Seeder interface for Subaru XV
func (sxs *SubaruXVSeeder) Seed(db *gorm.DB) error {
	productSlug := "subaru-xv"
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
		"Brand":                       "Subaru",
		"Model Name":                  "XV",
		"Variant":                        "2.0L EyeSight",
		"Generation":                     "3rd",
		"Segment":                        "C-Segment",
		"Launch Year":                    "2017",
		"Engine Configuration":           "Boxer",
		"Displacement (cc)":              "1995",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "98 hp/ton",
		"Mileage City (km/L)":            "13 km/L",
		"Mileage Highway (km/L)":         "15 km/L",
		"Mileage Combined (km/L)":        "14 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Double Wishbone",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "17",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "6.2 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Standard",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver + Passenger",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Double Wishbone",
		"Emission Standard":              "Euro 5",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.5:1",
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
		if banglaValue, exists := sxs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Subaru XV specifications seeded successfully")
	return nil
}
