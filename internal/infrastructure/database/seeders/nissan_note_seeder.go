package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// NissanNoteSeeder seeds specifications for Nissan Note
type NissanNoteSeeder struct {
	BaseSeeder
}

// NewNissanNoteSeeder creates a new Nissan Note seeder
func NewNissanNoteSeeder() *NissanNoteSeeder {
	return &NissanNoteSeeder{
		BaseSeeder: BaseSeeder{name: "Nissan Note Specifications"},
	}
}

func (nns *NissanNoteSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Note 1.2":                "নোট ১.২",
		"2nd Generation (E12)":    "২য় প্রজন্ম (ই১২)",
		"Compact MPV":             "কমপ্যাক্ট এমপিভি",
		"2019":                    "২০১৯",
		"Inline-3":                "ইনলাইন-৩",
		"1198":                    "১১৯৮ সিসি",
		"4":                       "৪ (ভালভ প্রতি সিলিন্ডার)",
		"Naturally Aspirated":     "ন্যাচারালি অ্যাসপিরেটেড",
		"Open":                    "ওপেন",
		"150 hp/ton":              "১৫০ হর্স পাওয়ার/টন",
		"13 km/L":                 "১৩ কিমি/লিটার",
		"19 km/L":                 "১৯ কিমি/লিটার",
		"16 km/L":                 "১৬ কিমি/লিটার",
		"MacPherson Strut":        "ম্যাকফারসন স্ট্রাট",
		"Torsion Beam":            "টর্শন বিম",
		"Electric Power Steering": "ইলেকট্রিক পাওয়ার স্টিয়ারিং",
		"Tilt":                    "টিল্ট",
		"15":                      "১৫",
		"Full Size":               "ফুল সাইজ",
		"Halogen":                 "হ্যালোজেন",
		"Yes":                     "হ্যাঁ",
		"No":                      "না",
		"Electric":                "ইলেকট্রিক",
		"Manual":                  "ম্যানুয়াল",
		"5":                       "৫",
		"Automatic":               "অটোমেটিক",
		"Single Zone":             "সিঙ্গেল জোন",
		"NissanConnect":           "নিসান কানেক্ট",
		"7 inch":                  "৭ ইঞ্চি",
		"Standard":                "স্ট্যান্ডার্ড",
		"Disc":                    "ডিস্ক",
		"Independent":             "ইন্ডিপেন্ডেন্ট",
		"Euro 6d":                 "ইউরো ৬ডি",
		"Push Button":             "পুশ বাটন",
		"Water Cooled":            "ওয়াটার কুলড",
		"Electronic":              "ইলেকট্রনিক",
		"11.0:1":                  "১১.০:১",
		"DOHC":                    "ডিওএইচসি",
	}
}

// Seed implements the Seeder interface for Nissan Note
func (nns *NissanNoteSeeder) Seed(db *gorm.DB) error {
	productSlug := "nissan-note"
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
		"Brand":                       "Nissan",
		"Model Name":                  "Note",
		"Variant":                        "Note 1.2",
		"Generation":                     "2nd Generation (E12)",
		"Segment":                        "Compact MPV",
		"Launch Year":                    "2019",
		"Engine Configuration":           "Inline-3",
		"Displacement (cc)":              "1198",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "150 hp/ton",
		"Mileage City (km/L)":            "13 km/L",
		"Mileage Highway (km/L)":         "19 km/L",
		"Mileage Combined (km/L)":        "16 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Electric Power Steering",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "15",
		"Spare Wheel Type":               "Full Size",
		"DRL":                            "Halogen",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Electric",
		"Wiper Type":                     "Manual",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Single Zone",
		"Infotainment System":            "NissanConnect",
		"Infotainment Screen (inch)":     "7 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Standard",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "6",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Euro 6d",
		"Starting System":                "Push Button",
		"Cooling System":                 "Water Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "11.0:1",
		"Valve Configuration":            "DOHC",
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
		if banglaValue, exists := nns.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Nissan Note specifications seeded successfully")
	return nil
}
