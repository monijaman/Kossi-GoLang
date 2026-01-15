package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MitsubishiASXSeeder seeds specifications for Mitsubishi ASX
type MitsubishiASXSeeder struct {
	BaseSeeder
}

// NewMitsubishiASXSeeder creates a new Mitsubishi ASX seeder
func NewMitsubishiASXSeeder() *MitsubishiASXSeeder {
	return &MitsubishiASXSeeder{
		BaseSeeder: BaseSeeder{name: "Mitsubishi ASX Specifications"},
	}
}

func (mass *MitsubishiASXSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.6L":                       "১.৬ লিটার",
		"2nd":                        "২য় প্রজন্ম",
		"B-Segment":                  "বি-সেগমেন্ট",
		"2010":                       "২০১০",
		"SUV":                        "এসইউভি",
		"White Pearl":                "হোয়াইট পার্ল",
		"Inline":                     "ইনলাইন",
		"1590":                       "১৫৯০ সিসি",
		"4 (Cylinders)":              "৪ (সিলিন্ডার)",
		"117 hp":                     "১১৭ হর্স পাওয়ার",
		"117 hp @ 6100 rpm":          "৬১০০ আরপিএম এ ১১৭ হর্স পাওয়ার",
		"154 Nm @ 4000 rpm":          "৪০০০ আরপিএম এ ১৫৪ এনএম",
		"Petrol":                     "পেট্রোল",
		"Multi-Point Injection":      "মাল্টি-পয়েন্ট ইনজেকশন",
		"11.4 seconds":               "১১.৪ সেকেন্ড",
		"180 km/h":                   "১৮০ কিমি/ঘণ্টা",
		"14 km/L":                    "১৪ কিমি/লিটার",
		"16 km/L":                    "১৬ কিমি/লিটার",
		"15 km/L":                    "১৫ কিমি/লিটার",
		"Manual (Transmission)":      "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":             "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":          "সামনের চাকা চালিত",
		"Single Clutch":              "একক ক্লাচ",
		"4295 mm":                    "৪২৯৫ মিমি",
		"1770 mm":                    "১৭৭০ মিমি",
		"1615 mm":                    "১৬১৫ মিমি",
		"2600 mm":                    "২৬০০ মিমি",
		"195 mm":                     "১৯৫ মিমি",
		"1240 kg":                    "১২৪০ কেজি",
		"384L":                       "৩৮৪ লিটার",
		"63L":                        "৬৩ লিটার",
		"205/60 R16":                 "২০৫/৬০ আর১৬",
		"Radial Tubeless":            "রেডিয়াল টিউবলেস",
		"Alloy":                      "অ্যালয়",
		"Halogen":                    "হ্যালোজেন",
		"Yes":                        "হ্যাঁ",
		"No":                         "না",
		"Power":                      "পাওয়ার",
		"Climate Control":            "জলবায়ু নিয়ন্ত্রণ",
		"Manual (Air Conditioning)":  "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"Touchscreen":                "টাচস্ক্রীন",
		"6.1 inch":                   "৬.১ ইঞ্চি",
		"Premium":                    "প্রিমিয়াম",
		"4 (Speakers)":               "৪ (স্পিকার)",
		"Dual":                       "ডুয়াল",
		"Front & Rear":               "সামনে এবং পিছনে",
		"Speed Sensitive":            "গতি সংবেদনশীল",
		"Electric":                   "বৈদ্যুতিক",
		"Liquid Cooled":              "লিকুইড কুলড",
		"Electronic":                 "ইলেকট্রনিক",
		"DOHC":                       "ডিওএইচসি",
		"Open":                       "ওপেন",
		"94 hp/ton":                  "৯৪ হর্স পাওয়ার/টন",
		"3 Years":                    "৩ বছর",
		"5 Years":                    "৫ বছর",
		"8 Years":                    "৮ বছর",
		"MacPherson Strut":           "ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":               "টর্শন বিম",
		"Rack and Pinion":            "র্যাক এবং পিনিয়ন",
		"Tilt":                       "টিল্ট",
		"Disc (Front) / Drum (Rear)": "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 4":                     "ইউরো ৪",
		"10.8:1":                     "১০.৮:১",
		"Naturally Aspirated":        "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mitsubishi ASX
func (mass *MitsubishiASXSeeder) Seed(db *gorm.DB) error {
	productSlug := "mitsubishi-asx"
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
		"Variant":                        "1.6L",
		"Generation":                     "2nd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2010",
		"Body Type":                      "SUV",
		"Color":                          "White Pearl",
		"Engine Type":                    "1.6L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1590",
		"Number of Cylinders":            "4",
		"Horsepower":                     "117 hp",
		"Max Power":                      "117 hp @ 6100 rpm",
		"Max Torque":                     "154 Nm @ 4000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Injection",
		"Acceleration 0-100 km/h":        "11.4 seconds",
		"Top Speed":                      "180 km/h",
		"Mileage":                        "14 km/L",
		"Mileage City (km/L)":            "14 km/L",
		"Mileage Highway (km/L)":         "16 km/L",
		"Mileage Combined (km/L)":        "15 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "5-Speed Manual",
		"Number of Gears":                "5",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4295 mm",
		"Width":                          "1770 mm",
		"Height":                         "1615 mm",
		"Wheelbase":                      "2600 mm",
		"Ground Clearance":               "195 mm",
		"Kerb Weight":                    "1240 kg",
		"Boot Space":                     "384L",
		"Fuel Tank Capacity":             "63L",
		"Fuel Capacity":                  "63L",
		"Tyre Size":                      "205/60 R16",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "Halogen",
		"DRL":                            "No",
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
		"Infotainment Screen (inch)":     "6.1 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "4",
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
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.8:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "94 hp/ton",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt",
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
		if banglaValue, exists := mass.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mitsubishi ASX specifications seeded successfully")
	return nil
}
