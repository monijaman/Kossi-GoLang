package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MazdaMazda6Seeder seeds specifications for Mazda Mazda6
type MazdaMazda6Seeder struct {
	BaseSeeder
}

// NewMazdaMazda6Seeder creates a new Mazda Mazda6 seeder
func NewMazdaMazda6Seeder() *MazdaMazda6Seeder {
	return &MazdaMazda6Seeder{
		BaseSeeder: BaseSeeder{name: "Mazda Mazda6 Specifications"},
	}
}

func (mm6s *MazdaMazda6Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L":                         "২.০ লিটার",
		"3rd":                          "৩য় প্রজন্ম",
		"D-Segment":                    "ডি-সেগমেন্ট",
		"2018":                         "২০১৮",
		"Sedan":                        "সেডান",
		"Soul Red Crystal":             "সোল রেড ক্রিস্টাল",
		"Inline":                       "ইনলাইন",
		"1998":                         "১৯৯৮ সিসি",
		"4":                            "৪ (সিলিন্ডার)",
		"165 hp":                       "১৬৫ হর্স পাওয়ার",
		"165 hp @ 6000 rpm":            "৬০০০ আরপিএম এ ১৬৫ হর্স পাওয়ার",
		"213 Nm @ 4000 rpm":            "৪০০০ আরপিএম এ ২১৩ এনএম",
		"Petrol":                       "পেট্রোল",
		"Direct Injection":             "সরাসরি ইনজেকশন",
		"8.2 seconds":                  "৮.২ সেকেন্ড",
		"210 km/h":                     "২১০ কিমি/ঘণ্টা",
		"14 km/L":                      "১৪ কিমি/লিটার",
		"16 km/L":                      "১৬ কিমি/লিটার",
		"15 km/L":                      "১৫ কিমি/লিটার",
		"Automatic (Transmission)":     "অটোমেটিক (ট্রান্সমিশন)",
		"6-Speed Automatic":            "৬-স্পিড অটোমেটিক",
		"Front-Wheel Drive":            "সামনের চাকা চালিত",
		"Single Clutch":                "একক ক্লাচ",
		"4870 mm":                      "৪৮৭০ মিমি",
		"1840 mm":                      "১৮৪০ মিমি",
		"1450 mm":                      "১৪৫০ মিমি",
		"2830 mm":                      "২৮৩০ মিমি",
		"165 mm":                       "১৬৫ মিমি",
		"1420 kg":                      "১৪২০ কেজি",
		"489L":                         "৪৮৯ লিটার",
		"62L":                          "৬২ লিটার",
		"225/55 R17":                   "২২৫/৫৫ আর১৭",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"Halogen":                      "হ্যালোজেন",
		"LED":                          "এলইডি",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"Power":                        "পাওয়ার",
		"Climate Control":              "জলবায়ু নিয়ন্ত্রণ",
		"Automatic (Air Conditioning)": "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Touchscreen":                  "টাচস্ক্রীন",
		"8 inch":                       "৮ ইঞ্চি",
		"Bose":                         "বোস",
		"11":                           "১১",
		"Dual":                         "ডুয়াল",
		"Front & Rear":                 "সামনে এবং পিছনে",
		"Speed Sensitive":              "গতি সংবেদনশীল",
		"Electric":                     "বৈদ্যুতিক",
		"Liquid Cooled":                "লিকুইড কুলড",
		"Electronic":                   "ইলেকট্রনিক",
		"DOHC":                         "ডিওএইচসি",
		"Open":                         "ওপেন",
		"116 hp/ton":                   "১১৬ হর্স পাওয়ার/টন",
		"3 Years":                      "৩ বছর",
		"5 Years":                      "৫ বছর",
		"8 Years":                      "৮ বছর",
		"MacPherson Strut":             "ম্যাকফার্সন স্ট্রাট",
		"Multi-Link":                   "মাল্টি-লিঙ্ক",
		"Rack and Pinion":              "র্যাক এবং পিনিয়ন",
		"Tilt & Telescopic":            "টিল্ট এবং টেলিস্কোপিক",
		"Disc (Front & Rear)":          "ডিস্ক (সামনে এবং পিছনে)",
		"Euro 6":                       "ইউরো ৬",
		"13.0:1":                       "১৩.০:১",
		"Naturally Aspirated":          "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mazda Mazda6
func (mm6s *MazdaMazda6Seeder) Seed(db *gorm.DB) error {
	productSlug := "mazda-mazda6"
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
		"Model Name":                  "Mazda6",
		"Variant":                        "2.0L",
		"Generation":                     "3rd",
		"Segment":                        "D-Segment",
		"Launch Year":                    "2018",
		"Body Type":                      "Sedan",
		"Color":                          "Soul Red Crystal",
		"Engine Type":                    "2.0L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1998",
		"Number of Cylinders":            "4",
		"Horsepower":                     "165 hp",
		"Max Power":                      "165 hp @ 6000 rpm",
		"Max Torque":                     "213 Nm @ 4000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "8.2 seconds",
		"Top Speed":                      "210 km/h",
		"Mileage":                        "14 km/L",
		"Mileage City (km/L)":            "14 km/L",
		"Mileage Highway (km/L)":         "16 km/L",
		"Mileage Combined (km/L)":        "15 km/L",
		"Transmission":                   "Automatic",
		"Gearbox":                        "6-Speed Automatic",
		"Number of Gears":                "6",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4870 mm",
		"Width":                          "1840 mm",
		"Height":                         "1450 mm",
		"Wheelbase":                      "2830 mm",
		"Ground Clearance":               "165 mm",
		"Kerb Weight":                    "1420 kg",
		"Boot Space":                     "489L",
		"Fuel Tank Capacity":             "62L",
		"Fuel Capacity":                  "62L",
		"Tyre Size":                      "225/55 R17",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Bose",
		"Number of Speakers":             "11",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front & Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Multi-Link",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "13.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "116 hp/ton",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Multi-Link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "17",
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
		if banglaValue, exists := mm6s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mazda Mazda6 specifications seeded successfully")
	return nil
}
