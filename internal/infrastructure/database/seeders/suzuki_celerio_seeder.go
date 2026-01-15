package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SuzukiCelerioSeeder seeds specifications for Suzuki Celerio
type SuzukiCelerioSeeder struct {
	BaseSeeder
}

// NewSuzukiCelerioSeeder creates a new Suzuki Celerio seeder
func NewSuzukiCelerioSeeder() *SuzukiCelerioSeeder {
	return &SuzukiCelerioSeeder{
		BaseSeeder: BaseSeeder{name: "Suzuki Celerio Specifications"},
	}
}

func (scs *SuzukiCelerioSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.0L":                       "১.০ লিটার",
		"2nd":                        "২য় প্রজন্ম",
		"A-Segment":                  "এ-সেগমেন্ট",
		"2017":                       "২০১৭",
		"Hatchback":                  "হ্যাচব্যাক",
		"Silky Silver":               "সিল্কি সিলভার",
		"Inline":                     "ইনলাইন",
		"998":                        "৯৯৮ সিসি",
		"3":                          "৩ (সিলিন্ডার)",
		"68 hp":                      "৬৮ হর্স পাওয়ার",
		"68 hp @ 6000 rpm":           "৬০০০ আরপিএম এ ৬৮ হর্স পাওয়ার",
		"90 Nm @ 3500 rpm":           "৩৫০০ আরপিএম এ ৯০ এনএম",
		"Petrol":                     "পেট্রোল",
		"Naturally Aspirated":        "ন্যাচারালি অ্যাসপিরেটেড",
		"Multipoint Injection":       "মাল্টিপয়েন্ট ইনজেকশন",
		"15.2 seconds":               "১৫.২ সেকেন্ড",
		"150 km/h":                   "১৫০ কিমি/ঘণ্টা",
		"20 km/L":                    "২০ কিমি/লিটার",
		"25 km/L":                    "২৫ কিমি/লিটার",
		"22 km/L":                    "২২ কিমি/লিটার",
		"Manual (Transmission)":      "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":             "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":          "সামনের চাকা চালিত",
		"Single Clutch":              "একক ক্লাচ",
		"3600 mm":                    "৩৬০০ মিমি",
		"1600 mm":                    "১৬০০ মিমি",
		"1560 mm":                    "১৫৬০ মিমি",
		"2435 mm":                    "২৪৩৫ মিমি",
		"1500 mm":                    "১৫০০ মিমি",
		"815 kg":                     "৮১৫ কেজি",
		"35 L":                       "৩৫ লিটার",
		"155/80 R13":                 "১৫৫/৮০ আর১৩",
		"Radial Tubeless":            "রেডিয়াল টিউবলেস",
		"Steel":                      "স্টিল",
		"Halogen":                    "হ্যালোজেন",
		"Yes":                        "হ্যাঁ",
		"Halogen":                    "হ্যালোজেন",
		"No":                         "না",
		"No":                         "না",
		"No":                         "না",
		"Manual":                     "ম্যানুয়াল",
		"Automatic":                  "অটোমেটিক",
		"5":                          "৫",
		"5":                          "৫",
		"Manual":                     "ম্যানুয়াল",
		"No":                         "না",
		"Manual (Air Conditioning)":  "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"No":                         "না",
		"Touchscreen":                "টাচস্ক্রিন",
		"7 inch":                     "৭ ইঞ্চি",
		"Yes":                        "হ্যাঁ",
		"Yes":                        "হ্যাঁ",
		"Standard":                   "স্ট্যান্ডার্ড",
		"4":                          "৪ (স্পিকার)",
		"No":                         "না",
		"Yes":                        "হ্যাঁ",
		"Driver + Passenger":         "ড্রাইভার + প্যাসেঞ্জার",
		"Yes":                        "হ্যাঁ",
		"Disc (Front) / Drum (Rear)": "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Yes":                        "হ্যাঁ",
		"Yes":                        "হ্যাঁ",
		"Yes":                        "হ্যাঁ",
		"Yes":                        "হ্যাঁ",
		"Yes":                        "হ্যাঁ",
		"Yes":                        "হ্যাঁ",
		"Keyless Entry":              "কীলেস এন্ট্রি",
		"No":                         "না",
		"Height Adjustable":          "হাইট অ্যাডজাস্টেবল",
		"Rear AC Vents":              "রিয়ার এসি ভেন্টস",
		"Steering Mounted Controls":  "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"No":                         "না",
		"Parking Sensors":            "পার্কিং সেন্সরস",
		"No":                         "না",
		"Rear Camera":                "রিয়ার ক্যামেরা",
		"No":                         "না",
		"Rain Sensing Wipers":        "রেইন সেন্সিং ওয়াইপারস",
		"No":                         "না",
		"Follow Me Home Headlamps":   "ফলো মি হোম হেডল্যাম্পস",
		"No":                         "না",
		"Sunroof":                    "সানরুফ",
		"No":                         "না",
		"Roof Rails":                 "রুফ রেলস",
		"No":                         "না",
		"Alloy Wheels":               "অ্যালয় হুইলস",
		"No":                         "না",
		"13 inch":                    "১৩ ইঞ্চি",
		"MacPherson Strut":           "ম্যাকফারসন স্ট্রাট",
		"Torsion Beam":               "টর্শন বিম",
		"Rack and Pinion":            "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt":                       "টিল্ট",
		"13":                         "১৩",
	}
}

// Seed implements the Seeder interface for Suzuki Celerio
func (scs *SuzukiCelerioSeeder) Seed(db *gorm.DB) error {
	productSlug := "suzuki-celerio"
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
		"Variant":                        "1.0L VXi",
		"Generation":                     "2nd",
		"Segment":                        "A-Segment",
		"Launch Year":                    "2017",
		"Engine Configuration":           "Inline",
		"Displacement (cc)":              "998",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "83 hp/ton",
		"Mileage City (km/L)":            "20 km/L",
		"Mileage Highway (km/L)":         "25 km/L",
		"Mileage Combined (km/L)":        "22 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "13",
		"Spare Wheel Type":               "Steel",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Manual",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "7 inch",
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
		"Suspension Type":                "MacPherson Strut / Torsion Beam",
		"Emission Standard":              "Euro 5",
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
		if banglaValue, exists := scs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Suzuki Celerio specifications seeded successfully")
	return nil
}
