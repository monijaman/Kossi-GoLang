package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MitsubishiMirageSeeder seeds specifications for Mitsubishi Mirage
type MitsubishiMirageSeeder struct {
	BaseSeeder
}

// NewMitsubishiMirageSeeder creates a new Mitsubishi Mirage seeder
func NewMitsubishiMirageSeeder() *MitsubishiMirageSeeder {
	return &MitsubishiMirageSeeder{
		BaseSeeder: BaseSeeder{name: "Mitsubishi Mirage Specifications"},
	}
}

func (mms *MitsubishiMirageSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.2L":                       "১.২ লিটার",
		"5th":                        "৫ম প্রজন্ম",
		"A-Segment":                  "এ-সেগমেন্ট",
		"2012":                       "২০১২",
		"Hatchback":                  "হ্যাচব্যাক",
		"White Pearl":                "হোয়াইট পার্ল",
		"Inline":                     "ইনলাইন",
		"1193":                       "১১৯৩ সিসি",
		"3":                          "৩ (সিলিন্ডার)",
		"78 hp":                      "৭৮ হর্স পাওয়ার",
		"78 hp @ 6000 rpm":           "৬০০০ আরপিএম এ ৭৮ হর্স পাওয়ার",
		"100 Nm @ 4000 rpm":          "৪০০০ আরপিএম এ ১০০ এনএম",
		"Petrol":                     "পেট্রোল",
		"Multi-Point Injection":      "মাল্টি-পয়েন্ট ইনজেকশন",
		"14.2 seconds":               "১৪.২ সেকেন্ড",
		"160 km/h":                   "১৬০ কিমি/ঘণ্টা",
		"18 km/L":                    "১৮ কিমি/লিটার",
		"20 km/L":                    "২০ কিমি/লিটার",
		"19 km/L":                    "১৯ কিমি/লিটার",
		"Manual (Transmission)":      "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":             "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":          "সামনের চাকা চালিত",
		"Single Clutch":              "একক ক্লাচ",
		"3780 mm":                    "৩৭৮০ মিমি",
		"1665 mm":                    "১৬৬৫ মিমি",
		"1515 mm":                    "১৫১৫ মিমি",
		"2450 mm":                    "২৪৫০ মিমি",
		"170 mm":                     "১৭০ মিমি",
		"850 kg":                     "৮৫০ কেজি",
		"235L":                       "২৩৫ লিটার",
		"35L":                        "৩৫ লিটার",
		"165/65 R14":                 "১৬৫/৬৫ আর১৪",
		"Radial Tubeless":            "রেডিয়াল টিউবলেস",
		"Steel":                      "স্টিল",
		"Halogen":                    "হ্যালোজেন",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"Manual (Air Conditioning)":  "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"Touchscreen":                "টাচস্ক্রীন",
		"6.1 inch":                   "৬.১ ইঞ্চি",
		"Standard":                   "স্ট্যান্ডার্ড",
		"4":                          "৪ (স্পিকার)",
		"Dual":                       "ডুয়াল",
		"Front":                      "সামনে",
		"Speed Sensitive":            "গতি সংবেদনশীল",
		"Electric":                   "বৈদ্যুতিক",
		"Liquid Cooled":              "লিকুইড কুলড",
		"Electronic":                 "ইলেকট্রনিক",
		"DOHC":                       "ডিওএইচসি",
		"Open":                       "ওপেন",
		"92 hp/ton":                  "৯২ হর্স পাওয়ার/টন",
		"3 Years":                    "৩ বছর",
		"5 Years":                    "৫ বছর",
		"8 Years":                    "৮ বছর",
		"MacPherson Strut":           "ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":               "টর্শন বিম",
		"Rack and Pinion":            "র্যাক এবং পিনিয়ন",
		"Tilt":                       "টিল্ট",
		"Disc (Front) / Drum (Rear)": "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 4":                     "ইউরো ৪",
		"11.1:1":                     "১১.১:১",
		"Naturally Aspirated":        "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mitsubishi Mirage
func (mms *MitsubishiMirageSeeder) Seed(db *gorm.DB) error {
	productSlug := "mitsubishi-mirage"
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
		"Body Type":                      790,
	}

	specs := map[string]string{
		"Brand":                       "Mitsubishi",
		"Model Name":                  "Mirage",
		"Variant":                        "1.2L",
		"Generation":                     "5th",
		"Segment":                        "A-Segment",
		"Launch Year":                    "2012",
		"Body Type":                      "Hatchback",
		"Color":                          "White Pearl",
		"Engine Type":                    "1.2L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1193",
		"Number of Cylinders":            "3",
		"Horsepower":                     "78 hp",
		"Max Power":                      "78 hp @ 6000 rpm",
		"Max Torque":                     "100 Nm @ 4000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Injection",
		"Acceleration 0-100 km/h":        "14.2 seconds",
		"Top Speed":                      "160 km/h",
		"Mileage":                        "18 km/L",
		"Mileage City (km/L)":            "18 km/L",
		"Mileage Highway (km/L)":         "20 km/L",
		"Mileage Combined (km/L)":        "19 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "5-Speed Manual",
		"Number of Gears":                "5",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "3780 mm",
		"Width":                          "1665 mm",
		"Height":                         "1515 mm",
		"Wheelbase":                      "2450 mm",
		"Ground Clearance":               "170 mm",
		"Kerb Weight":                    "850 kg",
		"Boot Space":                     "235L",
		"Fuel Tank Capacity":             "35L",
		"Fuel Capacity":                  "35L",
		"Tyre Size":                      "165/65 R14",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Steel",
		"Headlight Type":                 "Halogen",
		"DRL":                            "No",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Manual",
		"Wiper Type":                     "Manual",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "6.1 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Standard",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "No",
		"ESC":                            "No",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Torsion Beam",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "11.1:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "92 hp/ton",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "14",
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
		if banglaValue, exists := mms.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mitsubishi Mirage specifications seeded successfully")
	return nil
}
