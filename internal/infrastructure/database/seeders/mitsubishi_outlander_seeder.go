package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MitsubishiOutlanderSeeder seeds specifications for Mitsubishi Outlander
type MitsubishiOutlanderSeeder struct {
	BaseSeeder
}

// NewMitsubishiOutlanderSeeder creates a new Mitsubishi Outlander seeder
func NewMitsubishiOutlanderSeeder() *MitsubishiOutlanderSeeder {
	return &MitsubishiOutlanderSeeder{
		BaseSeeder: BaseSeeder{name: "Mitsubishi Outlander Specifications"},
	}
}

func (mos *MitsubishiOutlanderSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.4L":                                  "২.৪ লিটার",
		"3rd":                                   "৩য় প্রজন্ম",
		"C-Segment":                             "সি-সেগমেন্ট",
		"2014":                                  "২০১৪",
		"SUV":                                   "এসইউভি",
		"White Pearl":                           "হোয়াইট পার্ল",
		"Inline":                                "ইনলাইন",
		"2360":                                  "২৩৬০ সিসি",
		"4":                                     "৪ (সিলিন্ডার)",
		"170 hp":                                "১৭০ হর্স পাওয়ার",
		"170 hp @ 6000 rpm":                     "৬০০০ আরপিএম এ ১৭০ হর্স পাওয়ার",
		"232 Nm @ 4100 rpm":                     "৪১০০ আরপিএম এ ২৩২ এনএম",
		"Petrol":                                "পেট্রোল",
		"Direct Injection":                      "সরাসরি ইনজেকশন",
		"10.2 seconds":                          "১০.২ সেকেন্ড",
		"190 km/h":                              "১৯০ কিমি/ঘণ্টা",
		"12 km/L":                               "১২ কিমি/লিটার",
		"14 km/L":                               "১৪ কিমি/লিটার",
		"13 km/L":                               "১৩ কিমি/লিটার",
		"CVT":                                   "সিভিটি",
		"Automatic":                             "স্বয়ংক্রিয়",
		"Front-Wheel Drive":                     "সামনের চাকা চালিত",
		"Single Clutch":                         "একক ক্লাচ",
		"4695 mm":                               "৪৬৯৫ মিমি",
		"1800 mm":                               "১৮০০ মিমি",
		"1680 mm":                               "১৬৮০ মিমি",
		"2670 mm":                               "২৬৭০ মিমি",
		"190 mm":                                "১৯০ মিমি",
		"1530 kg":                               "১৫৩০ কেজি",
		"591L":                                  "৫৯১ লিটার",
		"63L":                                   "৬৩ লিটার",
		"215/70 R16":                            "২১৫/৭০ আর১৬",
		"Radial Tubeless":                       "রেডিয়াল টিউবলেস",
		"Alloy":                                 "অ্যালয়",
		"LED":                                   "এলইডি",
		"Halogen":                               "হ্যালোজেন",
		"Yes":                                   "হ্যাঁ",
		"No":                                    "না",
		"Power":                                 "পাওয়ার",
		"Panoramic":                             "প্যানোরামিক",
		"Climate Control":                       "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                             "ডুয়াল জোন",
		"Touchscreen":                           "টাচস্ক্রীন",
		"6.1 inch":                              "৬.১ ইঞ্চি",
		"Premium":                               "প্রিমিয়াম",
		"6":                                     "৬",
		"Dual":                                  "ডুয়াল",
		"Front & Rear":                          "সামনে এবং পিছনে",
		"Multiple Modes":                        "একাধিক মোড",
		"Speed Sensitive":                       "গতি সংবেদনশীল",
		"Electric":                              "বৈদ্যুতিক",
		"Liquid Cooled":                         "লিকুইড কুলড",
		"Electronic":                            "ইলেকট্রনিক",
		"DOHC":                                  "ডিওএইচসি",
		"Open":                                  "ওপেন",
		"111 hp/ton":                            "১১১ হর্স পাওয়ার/টন",
		"3 Years":                               "৩ বছর",
		"5 Years":                               "৫ বছর",
		"8 Years":                               "৮ বছর",
		"MacPherson Strut":                      "ম্যাকফার্সন স্ট্রাট",
		"Multi-Link":                            "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                       "র্যাক এবং পিনিয়ন",
		"Tilt & Telescopic":                     "টিল্ট এবং টেলিস্কোপিক",
		"Ventilated Disc (Front) / Drum (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 5":                                "ইউরো ৫",
		"10.5:1":                                "১০.৫:১",
		"Naturally Aspirated":                   "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mitsubishi Outlander
func (mos *MitsubishiOutlanderSeeder) Seed(db *gorm.DB) error {
	productSlug := "mitsubishi-outlander"
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
		"Variant":                        "2.4L",
		"Generation":                     "3rd",
		"Segment":                        "C-Segment",
		"Launch Year":                    "2014",
		"Body Type":                      "SUV",
		"Color":                          "White Pearl",
		"Engine Type":                    "2.4L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "2360",
		"Number of Cylinders":            "4",
		"Horsepower":                     "170 hp",
		"Max Power":                      "170 hp @ 6000 rpm",
		"Max Torque":                     "232 Nm @ 4100 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "10.2 seconds",
		"Top Speed":                      "190 km/h",
		"Mileage":                        "12 km/L",
		"Mileage City (km/L)":            "12 km/L",
		"Mileage Highway (km/L)":         "14 km/L",
		"Mileage Combined (km/L)":        "13 km/L",
		"Transmission":                   "CVT",
		"Gearbox":                        "CVT",
		"Number of Gears":                "CVT",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4695 mm",
		"Width":                          "1800 mm",
		"Height":                         "1680 mm",
		"Wheelbase":                      "2670 mm",
		"Ground Clearance":               "190 mm",
		"Kerb Weight":                    "1530 kg",
		"Boot Space":                     "591L",
		"Fuel Tank Capacity":             "63L",
		"Fuel Capacity":                  "63L",
		"Tyre Size":                      "215/70 R16",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Panoramic",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "7",
		"Number of Seats":                "7",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Climate Control",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "6.1 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Ventilated Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Multi-Link",
		"Emission Standard":              "Euro 5",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.5:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "111 hp/ton",
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
		if banglaValue, exists := mos.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mitsubishi Outlander specifications seeded successfully")
	return nil
}
