package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MercedesBenzCClassSeeder seeds specifications for Mercedes-Benz C-Class
type MercedesBenzCClassSeeder struct {
	BaseSeeder
}

// NewMercedesBenzCClassSeeder creates a new Mercedes-Benz C-Class seeder
func NewMercedesBenzCClassSeeder() *MercedesBenzCClassSeeder {
	return &MercedesBenzCClassSeeder{
		BaseSeeder: BaseSeeder{name: "Mercedes-Benz C-Class Specifications"},
	}
}

func (mbccs *MercedesBenzCClassSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"C 300":                          "সি ৩০০",
		"5th Generation (W206)":          "৫ম প্রজন্ম (ডব্লিউ২০৬)",
		"Compact Executive Sedan":        "কমপ্যাক্ট এক্সিকিউটিভ সেডান",
		"2021":                           "২০২১",
		"Inline-4":                       "ইনলাইন-৪",
		"1991":                           "১৯৯১ সিসি",
		"4":                              "৪ (ভালভ প্রতি সিলিন্ডার)",
		"Turbocharged":                   "টার্বোচার্জড",
		"Open":                           "ওপেন",
		"135 hp/ton":                     "১৩৫ হর্স পাওয়ার/টন",
		"10 km/L":                        "১০ কিমি/লিটার",
		"16 km/L":                        "১৬ কিমি/লিটার",
		"13 km/L":                        "১৩ কিমি/লিটার",
		"Multi-Link":                     "মাল্টি-লিঙ্ক",
		"Electric Power Steering":        "ইলেকট্রিক পাওয়ার স্টিয়ারিং",
		"Tilt & Telescopic":              "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"18":                             "১৮",
		"Tire Repair Kit":                "টায়ার রিপেয়ার কিট",
		"LED":                            "এলইডি",
		"Yes":                            "হ্যাঁ",
		"Panoramic":                      "প্যানোরামিক",
		"No":                             "না",
		"Electric, Heated, Auto-Dimming": "ইলেকট্রিক, হিটেড, অটো-ডিমিং",
		"Rain Sensing":                   "রেইন সেন্সিং",
		"5":                              "৫",
		"Electric":                       "ইলেকট্রিক",
		"Automatic":                      "অটোমেটিক",
		"Dual Zone":                      "ডুয়াল জোন",
		"MBUX":                           "এমবিইউএক্স",
		"11.9 inch":                      "১১.৯ ইঞ্চি",
		"Burmester":                      "বার্মেস্টার",
		"13":                             "১৩",
		"Disc":                           "ডিস্ক",
		"Independent":                    "ইন্ডিপেন্ডেন্ট",
		"Euro 6d":                        "ইউরো ৬ডি",
		"Push Button":                    "পুশ বাটন",
		"Water Cooled":                   "ওয়াটার কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"9.8:1":                          "৯.৮:১",
		"DOHC":                           "ডিওএইচসি",
	}
}

// Seed implements the Seeder interface for Mercedes-Benz C-Class
func (mbccs *MercedesBenzCClassSeeder) Seed(db *gorm.DB) error {
	productSlug := "mercedes-benz-c-class"
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
		"Displacement (cc)":              790,
	}

	specs := map[string]string{
		"Brand":                       "Mercedes-Benz",
		"Model Name":                  "C-Class",
		"Variant":                        "C 300",
		"Generation":                     "5th Generation (W206)",
		"Segment":                        "Compact Executive Sedan",
		"Launch Year":                    "2021",
		"Engine Configuration":           "Inline-4",
		"Displacement (cc)":              "1991",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Turbocharged",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "135 hp/ton",
		"Mileage City (km/L)":            "10 km/L",
		"Mileage Highway (km/L)":         "16 km/L",
		"Mileage Combined (km/L)":        "13 km/L",
		"Front Suspension":               "Multi-Link",
		"Rear Suspension":                "Multi-Link",
		"Steering Type":                  "Electric Power Steering",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "18",
		"Spare Wheel Type":               "Tire Repair Kit",
		"DRL":                            "LED",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Panoramic",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Electric, Heated, Auto-Dimming",
		"Wiper Type":                     "Rain Sensing",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Electric",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "MBUX",
		"Infotainment Screen (inch)":     "11.9 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Burmester",
		"Number of Speakers":             "13",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "7",
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
		"Compression Ratio":              "9.8:1",
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
		if banglaValue, exists := mbccs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mercedes-Benz C-Class specifications seeded successfully")
	return nil
}
