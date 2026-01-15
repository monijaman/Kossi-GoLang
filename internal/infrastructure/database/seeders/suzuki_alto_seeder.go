package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SuzukiAltoSeeder seeds specifications for Suzuki Alto
type SuzukiAltoSeeder struct {
	BaseSeeder
}

// NewSuzukiAltoSeeder creates a new Suzuki Alto seeder
func NewSuzukiAltoSeeder() *SuzukiAltoSeeder {
	return &SuzukiAltoSeeder{
		BaseSeeder: BaseSeeder{name: "Suzuki Alto Specifications"},
	}
}

func (sas *SuzukiAltoSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"0.8L":                       "০.৮ লিটার",
		"9th":                        "৯ম প্রজন্ম",
		"A-Segment":                  "এ-সেগমেন্ট",
		"2012":                       "২০১২",
		"Hatchback":                  "হ্যাচব্যাক",
		"Pearl Metallic":             "পার্ল মেটালিক",
		"Inline":                     "ইনলাইন",
		"796":                        "৭৯৬ সিসি",
		"3":                          "৩ (সিলিন্ডার)",
		"47 hp":                      "৪৭ হর্স পাওয়ার",
		"47 hp @ 6000 rpm":           "৬০০০ আরপিএম এ ৪৭ হর্স পাওয়ার",
		"69 Nm @ 3500 rpm":           "৩৫০০ আরপিএম এ ৬৯ এনএম",
		"Petrol":                     "পেট্রোল",
		"Naturally Aspirated":        "ন্যাচারালি অ্যাসপিরেটেড",
		"Multipoint Injection":       "মাল্টিপয়েন্ট ইনজেকশন",
		"18.5 seconds":               "১৮.৫ সেকেন্ড",
		"140 km/h":                   "১৪০ কিমি/ঘণ্টা",
		"22 km/L":                    "২২ কিমি/লিটার",
		"27 km/L":                    "২৭ কিমি/লিটার",
		"24 km/L":                    "২৪ কিমি/লিটার",
		"Manual (Transmission)":      "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":             "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":          "সামনের চাকা চালিত",
		"Single Clutch":              "একক ক্লাচ",
		"3395 mm":                    "৩৩৯৫ মিমি",
		"1475 mm":                    "১৪৭৫ মিমি",
		"1460 mm":                    "১৪৬০ মিমি",
		"2360 mm":                    "২৩৬০ মিমি",
		"1460 mm":                    "১৪৬০ মিমি",
		"730 kg":                     "৭৩০ কেজি",
		"35 L":                       "৩৫ লিটার",
		"145/80 R12":                 "১৪৫/৮০ আর১২",
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
		"Basic":                      "বেসিক",
		"6.1 inch":                   "৬.১ ইঞ্চি",
		"No":                         "না",
		"No":                         "না",
		"Standard":                   "স্ট্যান্ডার্ড",
		"2":                          "২ (স্পিকার)",
		"No":                         "না",
		"Yes":                        "হ্যাঁ",
		"Driver":                     "ড্রাইভার",
		"Yes":                        "হ্যাঁ",
		"Drum (Front) / Drum (Rear)": "ড্রাম (সামনে) / ড্রাম (পিছনে)",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"No":                         "না",
		"No":                         "না",
		"No":                         "না",
		"No":                         "না",
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
		"12 inch":                    "১২ ইঞ্চি",
		"MacPherson Strut":           "ম্যাকফারসন স্ট্রাট",
		"3-Link Rigid":               "৩-লিঙ্ক রিজিড",
		"Rack and Pinion":            "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt":                       "টিল্ট",
		"12":                         "১২",
	}
}

// Seed implements the Seeder interface for Suzuki Alto
func (sas *SuzukiAltoSeeder) Seed(db *gorm.DB) error {
	productSlug := "suzuki-alto"
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
		"Variant":                        "0.8L VXi",
		"Generation":                     "9th",
		"Segment":                        "A-Segment",
		"Launch Year":                    "2012",
		"Engine Configuration":           "Inline",
		"Displacement (cc)":              "796",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "64 hp/ton",
		"Mileage City (km/L)":            "22 km/L",
		"Mileage Highway (km/L)":         "27 km/L",
		"Mileage Combined (km/L)":        "24 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "3-Link Rigid",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "12",
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
		"Brake Type":                     "Drum (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "No",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "No",
		"Suspension Type":                "MacPherson Strut / 3-Link Rigid",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "11.0:1",
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
		if banglaValue, exists := sas.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Suzuki Alto specifications seeded successfully")
	return nil
}
