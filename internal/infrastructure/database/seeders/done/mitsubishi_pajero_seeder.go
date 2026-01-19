package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MitsubishiPajeroSeeder seeds specifications for Mitsubishi Pajero
type MitsubishiPajeroSeeder struct {
	BaseSeeder
}

// NewMitsubishiPajeroSeeder creates a new Mitsubishi Pajero seeder
func NewMitsubishiPajeroSeeder() *MitsubishiPajeroSeeder {
	return &MitsubishiPajeroSeeder{
		BaseSeeder: BaseSeeder{name: "Mitsubishi Pajero Specifications"},
	}
}

func (mps *MitsubishiPajeroSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.2L":                                  "৩.২ লিটার",
		"3rd":                                   "৩য় প্রজন্ম",
		"B-Segment":                             "বি-সেগমেন্ট",
		"2006":                                  "২০০৬",
		"SUV":                                   "এসইউভি",
		"Black Mica":                            "ব্ল্যাক মাইকা",
		"Inline":                                "ইনলাইন",
		"3200":                                  "৩২০০ সিসি",
		"4":                                     "৪ (সিলিন্ডার)",
		"165 hp":                                "১৬৫ হর্স পাওয়ার",
		"165 hp @ 4000 rpm":                     "৪০০০ আরপিএম এ ১৬৫ হর্স পাওয়ার",
		"380 Nm @ 2000 rpm":                     "২০০০ আরপিএম এ ৩৮০ এনএম",
		"Diesel":                                "ডিজেল",
		"Common Rail Direct Injection":          "কমন রেল ডাইরেক্ট ইনজেকশন",
		"12.5 seconds":                          "১২.৫ সেকেন্ড",
		"170 km/h":                              "১৭০ কিমি/ঘণ্টা",
		"12 km/L":                               "১২ কিমি/লিটার",
		"14 km/L":                               "১৪ কিমি/লিটার",
		"13 km/L":                               "১৩ কিমি/লিটার",
		"Manual":                                "ম্যানুয়াল (ট্রান্সমিশন)",
		"5-Speed Manual":                        "৫-স্পিড ম্যানুয়াল",
		"Four-Wheel Drive":                      "চার চাকা চালিত",
		"Single Clutch":                         "একক ক্লাচ",
		"4900 mm":                               "৪৯০০ মিমি",
		"1875 mm":                               "১৮৭৫ মিমি",
		"1900 mm":                               "১৯০০ মিমি",
		"2780 mm":                               "২৭৮০ মিমি",
		"225 mm":                                "২২৫ মিমি",
		"2100 kg":                               "২১০০ কেজি",
		"663L":                                  "৬৬৩ লিটার",
		"90L":                                   "৯০ লিটার",
		"265/65 R17":                            "২৬৫/৬৫ আর১৭",
		"Radial Tubeless":                       "রেডিয়াল টিউবলেস",
		"Alloy":                                 "অ্যালয়",
		"Halogen":                               "হ্যালোজেন",
		"Yes":                                   "হ্যাঁ",
		"No":                                    "না",
		"Power":                                 "পাওয়ার",
		"Panoramic":                             "প্যানোরামিক",
		"Automatic":                             "স্বয়ংক্রিয়",
		"Climate Control":                       "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                             "ডুয়াল জোন",
		"Touchscreen":                           "টাচস্ক্রীন",
		"7 inch":                                "৭ ইঞ্চি",
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
		"79 hp/ton":                             "৭৯ হর্স পাওয়ার/টন",
		"3 Years":                               "৩ বছর",
		"5 Years":                               "৫ বছর",
		"8 Years":                               "৮ বছর",
		"Independent Double Wishbone":           "স্বাধীন ডাবল উইশবোন",
		"Multi-Link":                            "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                       "র্যাক এবং পিনিয়ন",
		"Collapsible":                           "সংকোচনযোগ্য",
		"Ventilated Disc (Front) / Drum (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 4":                                "ইউরো ৪",
		"16.5:1":                                "১৬.৫:১",
		"Naturally Aspirated":                   "ন্যাচারালি অ্যাসপিরেটেড",
		"Turbocharged":                          "টার্বোচার্জড",
		"Intercooled":                           "ইন্টারকুলড",
	}
}

// Seed inserts specification records for Mitsubishi Pajero
func (mps *MitsubishiPajeroSeeder) Seed(db *gorm.DB) error {
	productSlug := "mitsubishi-pajero"
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
		"Model Name":                  "Pajero",
		"Variant":                        "3.2L",
		"Generation":                     "3rd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2006",
		"Body Type":                      "SUV",
		"Color":                          "Black Mica",
		"Engine Type":                    "3.2L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "3200",
		"Number of Cylinders":            "4",
		"Horsepower":                     "165 hp",
		"Max Power":                      "165 hp @ 4000 rpm",
		"Max Torque":                     "380 Nm @ 2000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Diesel",
		"Fuel System":                    "Common Rail Direct Injection",
		"Acceleration 0-100 km/h":        "12.5 seconds",
		"Top Speed":                      "170 km/h",
		"Mileage":                        "12 km/L",
		"Mileage City (km/L)":            "12 km/L",
		"Mileage Highway (km/L)":         "14 km/L",
		"Mileage Combined (km/L)":        "13 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "5-Speed Manual",
		"Number of Gears":                "5",
		"Drive Type":                     "Four-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4900 mm",
		"Width":                          "1875 mm",
		"Height":                         "1900 mm",
		"Wheelbase":                      "2780 mm",
		"Ground Clearance":               "225 mm",
		"Kerb Weight":                    "2100 kg",
		"Boot Space":                     "663L",
		"Fuel Tank Capacity":             "90L",
		"Fuel Capacity":                  "90L",
		"Tyre Size":                      "265/65 R17",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "Halogen",
		"DRL":                            "No",
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
		"Infotainment Screen (inch)":     "7 inch",
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
		"Suspension Type":                "Independent Double Wishbone / Multi-Link",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "16.5:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "79 hp/ton",
		"Front Suspension":               "Independent Double Wishbone",
		"Rear Suspension":                "Multi-Link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
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
		if banglaValue, exists := mps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mitsubishi Pajero specifications seeded successfully")
	return nil
}
