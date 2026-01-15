package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MazdaBt50Seeder seeds specifications for Mazda BT-50
type MazdaBt50Seeder struct {
	BaseSeeder
}

// NewMazdaBt50Seeder creates a new Mazda BT-50 seeder
func NewMazdaBt50Seeder() *MazdaBt50Seeder {
	return &MazdaBt50Seeder{
		BaseSeeder: BaseSeeder{name: "Mazda BT-50 Specifications"},
	}
}

func (mb50s *MazdaBt50Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.0L":                         "৩.০ লিটার",
		"2nd":                          "২য় প্রজন্ম",
		"Pickup Truck":                 "পিকআপ ট্রাক",
		"2016":                         "২০১৬",
		"Pickup":                       "পিকআপ",
		"Arctic White":                 "আর্কটিক হোয়াইট",
		"Inline":                       "ইনলাইন",
		"2999":                         "২৯৯৯ সিসি",
		"4":                            "৪ (সিলিন্ডার)",
		"190 hp":                       "১৯০ হর্স পাওয়ার",
		"190 hp @ 3000 rpm":            "৩০০০ আরপিএম এ ১৯০ হর্স পাওয়ার",
		"450 Nm @ 1600 rpm":            "১৬০০ আরপিএম এ ৪৫০ এনএম",
		"Diesel":                       "ডিজেল",
		"Common Rail Direct Injection": "কমন রেল সরাসরি ইনজেকশন",
		"10.5 seconds":                 "১০.৫ সেকেন্ড",
		"175 km/h":                     "১৭৫ কিমি/ঘণ্টা",
		"12 km/L":                      "১২ কিমি/লিটার",
		"14 km/L":                      "১৪ কিমি/লিটার",
		"13 km/L":                      "১৩ কিমি/লিটার",
		"Manual (Transmission)":        "ম্যানুয়াল (ট্রান্সমিশন)",
		"6-Speed Manual":               "৬-স্পিড ম্যানুয়াল",
		"Rear-Wheel Drive":             "পিছনের চাকা চালিত",
		"Single Clutch":                "একক ক্লাচ",
		"5355 mm":                      "৫৩৫৫ মিমি",
		"1860 mm":                      "১৮৬০ মিমি",
		"1780 mm":                      "১৭৮০ মিমি",
		"3125 mm":                      "৩১২৫ মিমি",
		"232 mm":                       "২৩২ মিমি",
		"1835 kg":                      "১৮৩৫ কেজি",
		"1450L":                        "১৪৫০ লিটার",
		"80L":                          "৮০ লিটার",
		"255/70 R16":                   "২৫৫/৭০ আর১৬",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"Halogen":                      "হ্যালোজেন",
		"LED":                          "এলইডি",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"Power":                        "পাওয়ার",
		"Climate Control":              "জলবায়ু নিয়ন্ত্রণ",
		"Manual (Air Conditioning)":    "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"Touchscreen":                  "টাচস্ক্রীন",
		"7 inch":                       "৭ ইঞ্চি",
		"Premium":                      "প্রিমিয়াম",
		"6":                            "৬",
		"Dual":                         "ডুয়াল",
		"Front & Rear":                 "সামনে এবং পিছনে",
		"Speed Sensitive":              "গতি সংবেদনশীল",
		"Electric":                     "বৈদ্যুতিক",
		"Liquid Cooled":                "লিকুইড কুলড",
		"Electronic":                   "ইলেকট্রনিক",
		"DOHC":                         "ডিওএইচসি",
		"104 hp/ton":                   "১০৪ হর্স পাওয়ার/টন",
		"3 Years":                      "৩ বছর",
		"5 Years":                      "৫ বছর",
		"8 Years":                      "৮ বছর",
		"Double Wishbone":              "ডাবল উইশবোন",
		"Leaf Spring":                  "লিফ স্প্রিং",
		"Rack and Pinion":              "র্যাক এবং পিনিয়ন",
		"Tilt & Telescopic":            "টিল্ট এবং টেলিস্কোপিক",
		"Disc (Front) / Drum (Rear)":   "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 5":                       "ইউরো ৫",
		"16.0:1":                       "১৬.০:১",
		"Turbocharged":                 "টার্বোচার্জড",
	}
}

// Seed inserts specification records for Mazda BT-50
func (mb50s *MazdaBt50Seeder) Seed(db *gorm.DB) error {
	productSlug := "mazda-bt-50"
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
		"Variant":                        "3.0L",
		"Generation":                     "2nd",
		"Segment":                        "Pickup Truck",
		"Launch Year":                    "2016",
		"Body Type":                      "Pickup",
		"Color":                          "Arctic White",
		"Engine Type":                    "3.0L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "2999",
		"Number of Cylinders":            "4",
		"Horsepower":                     "190 hp",
		"Max Power":                      "190 hp @ 3000 rpm",
		"Max Torque":                     "450 Nm @ 1600 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Diesel",
		"Fuel System":                    "Common Rail Direct Injection",
		"Acceleration 0-100 km/h":        "10.5 seconds",
		"Top Speed":                      "175 km/h",
		"Mileage":                        "12 km/L",
		"Mileage City (km/L)":            "12 km/L",
		"Mileage Highway (km/L)":         "14 km/L",
		"Mileage Combined (km/L)":        "13 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "6-Speed Manual",
		"Number of Gears":                "6",
		"Drive Type":                     "Rear-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "5355 mm",
		"Width":                          "1860 mm",
		"Height":                         "1780 mm",
		"Wheelbase":                      "3125 mm",
		"Ground Clearance":               "232 mm",
		"Kerb Weight":                    "1835 kg",
		"Boot Space":                     "1450L",
		"Fuel Tank Capacity":             "80L",
		"Fuel Capacity":                  "80L",
		"Tyre Size":                      "255/70 R16",
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
		"Suspension Type":                "Double Wishbone / Leaf Spring",
		"Emission Standard":              "Euro 5",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "16.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "104 hp/ton",
		"Front Suspension":               "Double Wishbone",
		"Rear Suspension":                "Leaf Spring",
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
		if banglaValue, exists := mb50s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mazda BT-50 specifications seeded successfully")
	return nil
}
