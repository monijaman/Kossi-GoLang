package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MitsubishiMonteroSeeder seeds specifications for Mitsubishi Montero
type MitsubishiMonteroSeeder struct {
	BaseSeeder
}

// NewMitsubishiMonteroSeeder creates a new Mitsubishi Montero seeder
func NewMitsubishiMonteroSeeder() *MitsubishiMonteroSeeder {
	return &MitsubishiMonteroSeeder{
		BaseSeeder: BaseSeeder{name: "Mitsubishi Montero Specifications"},
	}
}

func (mms *MitsubishiMonteroSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.5L":                                  "৩.৫ লিটার",
		"4th":                                   "৪র্থ প্রজন্ম",
		"B-Segment":                             "বি-সেগমেন্ট",
		"2006":                                  "২০০৬",
		"SUV":                                   "এসইউভি",
		"White Pearl":                           "হোয়াইট পার্ল",
		"V6":                                    "ভি৬",
		"3497":                                  "৩৪৯৭ সিসি",
		"6":                                     "৬ (সিলিন্ডার)",
		"200 hp":                                "২০০ হর্স পাওয়ার",
		"200 hp @ 5000 rpm":                     "৫০০০ আরপিএম এ ২০০ হর্স পাওয়ার",
		"323 Nm @ 4000 rpm":                     "৪০০০ আরপিএম এ ৩২৩ এনএম",
		"Petrol":                                "পেট্রোল",
		"Multi-Point Injection":                 "মাল্টি-পয়েন্ট ইনজেকশন",
		"10.8 seconds":                          "১০.৮ সেকেন্ড",
		"180 km/h":                              "১৮০ কিমি/ঘণ্টা",
		"10 km/L":                               "১০ কিমি/লিটার",
		"12 km/L":                               "১২ কিমি/লিটার",
		"11 km/L":                               "১১ কিমি/লিটার",
		"Automatic":                             "স্বয়ংক্রিয় (ট্রান্সমিশন)",
		"5-Speed Automatic":                     "৫-স্পিড স্বয়ংক্রিয়",
		"Four-Wheel Drive":                      "চার চাকা চালিত",
		"Single Clutch":                         "একক ক্লাচ",
		"4905 mm":                               "৪৯০৫ মিমি",
		"1875 mm":                               "১৮৭৫ মিমি",
		"1895 mm":                               "১৮৯৫ মিমি",
		"2780 mm":                               "২৭৮০ মিমি",
		"235 mm":                                "২৩৫ মিমি",
		"2350 kg":                               "২৩৫০ কেজি",
		"604L":                                  "৬০৪ লিটার",
		"90L":                                   "৯০ লিটার",
		"265/70 R16":                            "২৬৫/৭০ আর১৬",
		"Radial Tubeless":                       "রেডিয়াল টিউবলেস",
		"Alloy":                                 "অ্যালয়",
		"Halogen":                               "হ্যালোজেন",
		"Yes":                                   "হ্যাঁ",
		"No":                                    "না",
		"Power":                                 "পাওয়ার",
		"Panoramic":                             "প্যানোরামিক",
		"Climate Control":                       "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                             "ডুয়াল জোন",
		"Touchscreen":                           "টাচস্ক্রীন",
		"8 inch":                                "৮ ইঞ্চি",
		"Premium":                               "প্রিমিয়াম",
		"7":                                     "৭",
		"Dual":                                  "ডুয়াল",
		"Front & Rear":                          "সামনে এবং পিছনে",
		"Multiple Modes":                        "একাধিক মোড",
		"Speed Sensitive":                       "গতি সংবেদনশীল",
		"Electric":                              "বৈদ্যুতিক",
		"Liquid Cooled":                         "লিকুইড কুলড",
		"Electronic":                            "ইলেকট্রনিক",
		"SOHC":                                  "এসওএইচসি",
		"Open":                                  "ওপেন",
		"85 hp/ton":                             "৮৫ হর্স পাওয়ার/টন",
		"3 Years":                               "৩ বছর",
		"5 Years":                               "৫ বছর",
		"8 Years":                               "৮ বছর",
		"Independent Double Wishbone":           "স্বাধীন ডাবল উইশবোন",
		"Multi-Link":                            "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                       "র্যাক এবং পিনিয়ন",
		"Collapsible":                           "সংকোচনযোগ্য",
		"Ventilated Disc (Front) / Drum (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 4":                                "ইউরো ৪",
		"9.5:1":                                 "৯.৫:১",
		"Naturally Aspirated":                   "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Mitsubishi Montero
func (mms *MitsubishiMonteroSeeder) Seed(db *gorm.DB) error {
	productSlug := "mitsubishi-montero"
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
		"Model Name":                  "Montero",
		"Variant":                        "3.5L",
		"Generation":                     "4th",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2006",
		"Body Type":                      "SUV",
		"Color":                          "White Pearl",
		"Engine Type":                    "3.5L",
		"Engine Configuration":           "V6",
		"Displacement":                   "3497",
		"Number of Cylinders":            "6",
		"Horsepower":                     "200 hp",
		"Max Power":                      "200 hp @ 5000 rpm",
		"Max Torque":                     "323 Nm @ 4000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Injection",
		"Acceleration 0-100 km/h":        "10.8 seconds",
		"Top Speed":                      "180 km/h",
		"Mileage":                        "10 km/L",
		"Mileage City (km/L)":            "10 km/L",
		"Mileage Highway (km/L)":         "12 km/L",
		"Mileage Combined (km/L)":        "11 km/L",
		"Transmission":                   "Automatic",
		"Gearbox":                        "5-Speed Automatic",
		"Number of Gears":                "5",
		"Drive Type":                     "Four-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4905 mm",
		"Width":                          "1875 mm",
		"Height":                         "1895 mm",
		"Wheelbase":                      "2780 mm",
		"Ground Clearance":               "235 mm",
		"Kerb Weight":                    "2350 kg",
		"Boot Space":                     "604L",
		"Fuel Tank Capacity":             "90L",
		"Fuel Capacity":                  "90L",
		"Tyre Size":                      "265/70 R16",
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
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "7",
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
		"Compression Ratio":              "9.5:1",
		"Valve Configuration":            "SOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "85 hp/ton",
		"Front Suspension":               "Independent Double Wishbone",
		"Rear Suspension":                "Multi-Link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
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

	log.Printf("✅ Mitsubishi Montero specifications seeded successfully")
	return nil
}
