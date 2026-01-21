package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MitsubishiLancerSeeder seeds specifications for Mitsubishi Lancer
type MitsubishiLancerSeeder struct {
	BaseSeeder
}

// NewMitsubishiLancerSeeder creates a new Mitsubishi Lancer seeder
func NewMitsubishiLancerSeeder() *MitsubishiLancerSeeder {
	return &MitsubishiLancerSeeder{
		BaseSeeder: BaseSeeder{name: "Mitsubishi Lancer Specifications"},
	}
}

func (mls *MitsubishiLancerSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L":                                  "২.০ লিটার",
		"10th":                                  "১০ম প্রজন্ম",
		"C-Segment":                             "সি-সেগমেন্ট",
		"2008":                                  "২০০৮",
		"Sedan":                                 "সেডান",
		"White Pearl":                           "হোয়াইট পার্ল",
		"Inline":                                "ইনলাইন",
		"1998":                                  "১৯৯৮ সিসি",
		"4 (Cylinders)":                         "৪ (সিলিন্ডার)",
		"152 hp":                                "১৫২ হর্স পাওয়ার",
		"152 hp @ 6000 rpm":                     "৬০০০ আরপিএম এ ১৫২ হর্স পাওয়ার",
		"198 Nm @ 4250 rpm":                     "৪২৫০ আরপিএম এ ১৯৮ এনএম",
		"Petrol":                                "পেট্রোল",
		"Multi-Point Injection":                 "মাল্টি-পয়েন্ট ইনজেকশন",
		"9.8 seconds":                           "৯.৮ সেকেন্ড",
		"200 km/h":                              "২০০ কিমি/ঘণ্টা",
		"12 km/L":                               "১২ কিমি/লিটার",
		"14 km/L":                               "১৪ কিমি/লিটার",
		"13 km/L":                               "১৩ কিমি/লিটার",
		"Manual (Transmission)":                 "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":                        "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":                     "সামনের চাকা চালিত",
		"Single Clutch":                         "একক ক্লাচ",
		"4570 mm":                               "৪৫৭০ মিমি",
		"1760 mm":                               "১৭৬০ মিমি",
		"1490 mm":                               "১৪৯০ মিমি",
		"2635 mm":                               "২৬৩৫ মিমি",
		"165 mm":                                "১৬৫ মিমি",
		"1290 kg":                               "১২৯০ কেজি",
		"420L":                                  "৪২০ লিটার",
		"59L":                                   "৫৯ লিটার",
		"205/60 R16":                            "২০৫/৬০ আর১৬",
		"Radial Tubeless":                       "রেডিয়াল টিউবলেস",
		"Alloy":                                 "অ্যালয়",
		"Halogen":                               "হ্যালোজেন",
		"Yes":                                   "হ্যাঁ",
		"No":                                    "না",
		"Power":                                 "পাওয়ার",
		"Climate Control":                       "জলবায়ু নিয়ন্ত্রণ",
		"Manual (Air Conditioning)":             "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"Touchscreen":                           "টাচস্ক্রীন",
		"6.1 inch":                              "৬.১ ইঞ্চি",
		"Premium":                               "প্রিমিয়াম",
		"4 (Speakers)":                          "৪ (স্পিকার)",
		"Dual":                                  "ডুয়াল",
		"Front & Rear":                          "সামনে এবং পিছনে",
		"Speed Sensitive":                       "গতি সংবেদনশীল",
		"Electric":                              "বৈদ্যুতিক",
		"Liquid Cooled":                         "লিকুইড কুলড",
		"Electronic":                            "ইলেকট্রনিক",
		"DOHC":                                  "ডিওএইচসি",
		"Open":                                  "ওপেন",
		"118 hp/ton":                            "১১৮ হর্স পাওয়ার/টন",
		"3 Years":                               "৩ বছর",
		"5 Years":                               "৫ বছর",
		"8 Years":                               "৮ বছর",
		"MacPherson Strut":                      "ম্যাকফার্সন স্ট্রাট",
		"Multi-Link":                            "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                       "র্যাক এবং পিনিয়ন",
		"Tilt & Telescopic":                     "টিল্ট এবং টেলিস্কোপিক",
		"Ventilated Disc (Front) / Drum (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 4":                                "ইউরো ৪",
		"10.8:1":                                "১০.৮:১",
		"Naturally Aspirated":                   "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mitsubishi Lancer
func (mls *MitsubishiLancerSeeder) Seed(db *gorm.DB) error {
	productSlug := "mitsubishi-lancer"
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
		"Brand":                       "Mitsubishi",
		"Model Name":                  "Lancer",
		"Variant":                        "2.0L",
		"Generation":                     "10th",
		"Segment":                        "C-Segment",
		"Launch Year":                    "2008",
		"Body Type":                      "Sedan",
		"Color":                          "White Pearl",
		"Engine Type":                    "2.0L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1998",
		"Number of Cylinders":            "4",
		"Horsepower":                     "152 hp",
		"Max Power":                      "152 hp @ 6000 rpm",
		"Max Torque":                     "198 Nm @ 4250 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Injection",
		"Acceleration 0-100 km/h":        "9.8 seconds",
		"Top Speed":                      "200 km/h",
		"Mileage":                        "12 km/L",
		"Mileage City (km/L)":            "12 km/L",
		"Mileage Highway (km/L)":         "14 km/L",
		"Mileage Combined (km/L)":        "13 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "5-Speed Manual",
		"Number of Gears":                "5",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4570 mm",
		"Width":                          "1760 mm",
		"Height":                         "1490 mm",
		"Wheelbase":                      "2635 mm",
		"Ground Clearance":               "165 mm",
		"Kerb Weight":                    "1290 kg",
		"Boot Space":                     "420L",
		"Fuel Tank Capacity":             "59L",
		"Fuel Capacity":                  "59L",
		"Tyre Size":                      "205/60 R16",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "Halogen",
		"DRL":                            "No",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
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
		"Brake Type":                     "Ventilated Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Multi-Link",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.8:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "118 hp/ton",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Multi-Link",
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
		if banglaValue, exists := mls.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mitsubishi Lancer specifications seeded successfully")
	return nil
}
