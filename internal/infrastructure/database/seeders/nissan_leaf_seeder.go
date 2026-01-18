package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// NissanLeafSeeder seeds specifications for Nissan Leaf
type NissanLeafSeeder struct {
	BaseSeeder
}

// NewNissanLeafSeeder creates a new Nissan Leaf seeder
func NewNissanLeafSeeder() *NissanLeafSeeder {
	return &NissanLeafSeeder{
		BaseSeeder: BaseSeeder{name: "Nissan Leaf Specifications"},
	}
}

func (nls *NissanLeafSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Leaf e+":                    "লিফ ই+",
		"2nd Generation (ZE1)":       "২য় প্রজন্ম (জেডই১)",
		"Compact Hatchback Electric": "কমপ্যাক্ট হ্যাচব্যাক ইলেকট্রিক",
		"2021":                       "২০২১",
		"Electric Motor":             "ইলেকট্রিক মোটর",
		"N/A":                        "এন/এ",

		"MacPherson Strut":               "ম্যাকফারসন স্ট্রাট",
		"Torsion Beam":                   "টর্শন বিম",
		"Electric Power Steering":        "ইলেকট্রিক পাওয়ার স্টিয়ারিং",
		"Tilt & Telescopic":              "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"16":                             "১৬",
		"Full Size":                      "ফুল সাইজ",
		"LED":                            "এলইডি",
		"Yes":                            "হ্যাঁ",
		"No":                             "না",
		"Electric, Heated, Auto-Dimming": "ইলেকট্রিক, হিটেড, অটো-ডিমিং",
		"Rain Sensing":                   "রেইন সেন্সিং",
		"5":                              "৫",
		"Electric":                       "ইলেকট্রিক",
		"Automatic":                      "অটোমেটিক",
		"Single Zone":                    "সিঙ্গেল জোন",
		"NissanConnect":                  "নিসান কানেক্ট",
		"8 inch":                         "৮ ইঞ্চি",
		"Bose":                           "বোজ",
		"6":                              "৬",
		"Regenerative":                   "রিজেনারেটিভ",
		"Independent":                    "ইন্ডিপেন্ডেন্ট",
		"Zero Emission":                  "জিরো এমিশন",
		"Push Button":                    "পুশ বাটন",
	}
}

// Seed implements the Seeder interface for Nissan Leaf
func (nls *NissanLeafSeeder) Seed(db *gorm.DB) error {
	productSlug := "nissan-leaf"
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
		"Variant":                        "Leaf e+",
		"Generation":                     "2nd Generation (ZE1)",
		"Segment":                        "Compact Hatchback Electric",
		"Launch Year":                    "2021",
		"Engine Configuration":           "Electric Motor",
		"Displacement (cc)":              "N/A",
		"Valves Per Cylinder":            "N/A",
		"Engine Aspiration":              "N/A",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "N/A",
		"Mileage City (km/L)":            "N/A",
		"Mileage Highway (km/L)":         "N/A",
		"Mileage Combined (km/L)":        "N/A",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Electric Power Steering",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "16",
		"Spare Wheel Type":               "Full Size",
		"DRL":                            "LED",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Electric, Heated, Auto-Dimming",
		"Wiper Type":                     "Rain Sensing",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Electric",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Single Zone",
		"Infotainment System":            "NissanConnect",
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Bose",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "6",
		"EBD":                            "Yes",
		"Brake Type":                     "Regenerative",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Zero Emission",
		"Starting System":                "Push Button",
		"Cooling System":                 "N/A",
		"Ignition Type":                  "N/A",
		"Compression Ratio":              "N/A",
		"Valve Configuration":            "N/A",
		"Valve Per Cylinder":             "N/A",
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
		if banglaValue, exists := nls.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Nissan Leaf specifications seeded successfully")
	return nil
}
