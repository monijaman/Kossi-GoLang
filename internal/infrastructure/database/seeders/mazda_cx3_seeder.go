package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MazdaCx3Seeder seeds specifications for Mazda CX-3
type MazdaCx3Seeder struct {
	BaseSeeder
}

// NewMazdaCx3Seeder creates a new Mazda CX-3 seeder
func NewMazdaCx3Seeder() *MazdaCx3Seeder {
	return &MazdaCx3Seeder{
		BaseSeeder: BaseSeeder{name: "Mazda CX-3 Specifications"},
	}
}

func (mc3s *MazdaCx3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L":                       "১.৫ লিটার",
		"2nd":                        "২য় প্রজন্ম",
		"B-Segment":                  "বি-সেগমেন্ট",
		"2018":                       "২০১৮",
		"SUV":                        "এসইউভি",
		"Jet Black Mica":             "জেট ব্ল্যাক মাইকা",
		"Inline":                     "ইনলাইন",
		"1496":                       "১৪৯৬ সিসি",
		"4":                          "৪ (সিলিন্ডার)",
		"115 hp":                     "১১৫ হর্স পাওয়ার",
		"115 hp @ 6000 rpm":          "৬০০০ আরপিএম এ ১১৫ হর্স পাওয়ার",
		"152 Nm @ 4000 rpm":          "৪০০০ আরপিএম এ ১৫২ এনএম",
		"Petrol":                     "পেট্রোল",
		"Direct Injection":           "সরাসরি ইনজেকশন",
		"9.9 seconds":                "৯.৯ সেকেন্ড",
		"190 km/h":                   "১৯০ কিমি/ঘণ্টা",
		"16 km/L":                    "১৬ কিমি/লিটার",
		"18 km/L":                    "১৮ কিমি/লিটার",
		"17 km/L":                    "১৭ কিমি/লিটার",
		"Manual (Transmission)":      "ম্যানুয়াল (ট্রান্সমিশন)",
		"6-Speed Manual":             "৬-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":          "সামনের চাকা চালিত",
		"Single Clutch":              "একক ক্লাচ",
		"4275 mm":                    "৪২৭৫ মিমি",
		"1765 mm":                    "১৭৬৫ মিমি",
		"1550 mm":                    "১৫৫০ মিমি",
		"2570 mm":                    "২৫৭০ মিমি",
		"160 mm":                     "১৬০ মিমি",
		"1205 kg":                    "১২০৫ কেজি",
		"350L":                       "৩৫০ লিটার",
		"48L":                        "৪৮ লিটার",
		"215/60 R16":                 "২১৫/৬০ আর১৬",
		"Radial Tubeless":            "রেডিয়াল টিউবলেস",
		"Alloy":                      "অ্যালয়",
		"Halogen":                    "হ্যালোজেন",
		"LED":                        "এলইডি",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"Power":                      "পাওয়ার",
		"Climate Control":            "জলবায়ু নিয়ন্ত্রণ",
		"Manual (Air Conditioning)":  "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"Touchscreen":                "টাচস্ক্রীন",
		"7 inch":                     "৭ ইঞ্চি",
		"Premium":                    "প্রিমিয়াম",
		"6":                          "৬",
		"Dual":                       "ডুয়াল",
		"Front & Rear":               "সামনে এবং পিছনে",
		"Speed Sensitive":            "গতি সংবেদনশীল",
		"Electric":                   "বৈদ্যুতিক",
		"Liquid Cooled":              "লিকুইড কুলড",
		"Electronic":                 "ইলেকট্রনিক",
		"DOHC":                       "ডিওএইচসি",
		"Open":                       "ওপেন",
		"95 hp/ton":                  "৯৫ হর্স পাওয়ার/টন",
		"3 Years":                    "৩ বছর",
		"5 Years":                    "৫ বছর",
		"8 Years":                    "৮ বছর",
		"MacPherson Strut":           "ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":               "টর্শন বিম",
		"Rack and Pinion":            "র্যাক এবং পিনিয়ন",
		"Tilt & Telescopic":          "টিল্ট এবং টেলিস্কোপিক",
		"Disc (Front) / Drum (Rear)": "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 6":                     "ইউরো ৬",
		"13.0:1":                     "১৩.০:১",
		"Naturally Aspirated":        "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mazda CX-3
func (mc3s *MazdaCx3Seeder) Seed(db *gorm.DB) error {
	productSlug := "mazda-cx-3"
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
		"Body Type":                      790,
	}

	specs := map[string]string{
		"Brand":                       "Mazda",
		"Model Name":                  "CX-3",
		"Variant":                        "1.5L",
		"Generation":                     "2nd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2018",
		"Body Type":                      "SUV",
		"Color":                          "Jet Black Mica",
		"Engine Type":                    "1.5L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1496",
		"Number of Cylinders":            "4",
		"Horsepower":                     "115 hp",
		"Max Power":                      "115 hp @ 6000 rpm",
		"Max Torque":                     "152 Nm @ 4000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "9.9 seconds",
		"Top Speed":                      "190 km/h",
		"Mileage":                        "16 km/L",
		"Mileage City (km/L)":            "16 km/L",
		"Mileage Highway (km/L)":         "18 km/L",
		"Mileage Combined (km/L)":        "17 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "6-Speed Manual",
		"Number of Gears":                "6",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4275 mm",
		"Width":                          "1765 mm",
		"Height":                         "1550 mm",
		"Wheelbase":                      "2570 mm",
		"Ground Clearance":               "160 mm",
		"Kerb Weight":                    "1205 kg",
		"Boot Space":                     "350L",
		"Fuel Tank Capacity":             "48L",
		"Fuel Capacity":                  "48L",
		"Tyre Size":                      "215/60 R16",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
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
		"Infotainment Screen (inch)":     "7 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Torsion Beam",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "13.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "95 hp/ton",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "16",
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
		if banglaValue, exists := mc3s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mazda CX-3 specifications seeded successfully")
	return nil
}
