package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// GMCSierraSeeder seeds specifications for GMC Sierra
type GMCSierraSeeder struct {
	BaseSeeder
}

// NewGMCSierraSeeder creates a new GMC Sierra seeder
func NewGMCSierraSeeder() *GMCSierraSeeder {
	return &GMCSierraSeeder{
		BaseSeeder: BaseSeeder{name: "GMC Sierra Specifications"},
	}
}

func (gss *GMCSierraSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Denali":                         "ডেনালি",
		"4th Generation":                 "৪র্থ প্রজন্ম",
		"Full-size Pickup":               "ফুল-সাইজ পিকআপ",
		"2019":                           "২০১৯",
		"V8":                             "ভি৮",
		"6200":                           "৬২০০ সিসি",
		"2":                              "২ (ভালভ প্রতি সিলিন্ডার)",
		"Naturally Aspirated":            "ন্যাচারালি অ্যাসপিরেটেড",
		"Locking":                        "লকিং",
		"120 hp/ton":                     "১২০ হর্স পাওয়ার/টন",
		"6 km/L":                         "৬ কিমি/লিটার",
		"10 km/L":                        "১০ কিমি/লিটার",
		"8 km/L":                         "৮ কিমি/লিটার",
		"Independent":                    "ইন্ডিপেন্ডেন্ট",
		"Solid Axle":                     "সলিড অ্যাক্সেল",
		"Electric Power Steering":        "ইলেকট্রিক পাওয়ার স্টিয়ারিং",
		"Tilt & Telescopic":              "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"20":                             "২০",
		"Full Size":                      "ফুল সাইজ",
		"LED":                            "এলইডি",
		"Yes":                            "হ্যাঁ",
		"No":                             "না",
		"Electric, Heated, Auto-Dimming": "ইলেকট্রিক, হিটেড, অটো-ডিমিং",
		"Variable Speed":                 "ভেরিয়েবল স্পিড",
		"5":                              "৫",
		"Electric":                       "ইলেকট্রিক",
		"Automatic":                      "অটোমেটিক",
		"Dual Zone":                      "ডুয়াল জোন",
		"GMC Infotainment":               "জিএমসি ইনফোটেইনমেন্ট",
		"8 inch":                         "৮ ইঞ্চি",
		"Bose":                           "বোজ",
		"7":                              "৭",
		"Disc":                           "ডিস্ক",
		"Independent Front, Solid Rear":  "ইন্ডিপেন্ডেন্ট ফ্রন্ট, সলিড রিয়ার",
		"Euro 6":                         "ইউরো ৬",
		"Push Button":                    "পুশ বাটন",
		"Water Cooled":                   "ওয়াটার কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"11.5:1":                         "১১.৫:১",
		"OHV":                            "ওএইচভি",
	}
}

// Seed implements the Seeder interface for GMC Sierra
func (gss *GMCSierraSeeder) Seed(db *gorm.DB) error {
	productSlug := "gmc-sierra"
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
		"Brand":                       "GMC",
		"Model Name":                  "Sierra",
		"Variant":                        "Denali",
		"Generation":                     "4th Generation",
		"Segment":                        "Full-size Pickup",
		"Launch Year":                    "2019",
		"Engine Configuration":           "V8",
		"Displacement (cc)":              "6200",
		"Valves Per Cylinder":            "2",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Locking",
		"Power to Weight (HP/ton)":       "120 hp/ton",
		"Mileage City (km/L)":            "6 km/L",
		"Mileage Highway (km/L)":         "10 km/L",
		"Mileage Combined (km/L)":        "8 km/L",
		"Front Suspension":               "Independent",
		"Rear Suspension":                "Solid Axle",
		"Steering Type":                  "Electric Power Steering",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "20",
		"Spare Wheel Type":               "Full Size",
		"DRL":                            "LED",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Electric, Heated, Auto-Dimming",
		"Wiper Type":                     "Variable Speed",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Electric",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "GMC Infotainment",
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Bose",
		"Number of Speakers":             "7",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "6",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "Independent Front, Solid Rear",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Push Button",
		"Cooling System":                 "Water Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "11.5:1",
		"Valve Configuration":            "OHV",
		"Valve Per Cylinder":             "2",
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
		if banglaValue, exists := gss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ GMC Sierra specifications seeded successfully")
	return nil
}
