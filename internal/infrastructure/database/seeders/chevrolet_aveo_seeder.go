+package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ChevroletAveoSeeder seeds specifications for Chevrolet Aveo
type ChevroletAveoSeeder struct {
	BaseSeeder
}

// NewChevroletAveoSeeder creates a new Chevrolet Aveo seeder
func NewChevroletAveoSeeder() *ChevroletAveoSeeder {
	return &ChevroletAveoSeeder{
		BaseSeeder: BaseSeeder{name: "Chevrolet Aveo Specifications"},
	}
}

func (cas *ChevroletAveoSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LS":                       "এলএস",
		"2nd Generation":           "২য় প্রজন্ম",
		"Subcompact Sedan":         "সাবকমপ্যাক্ট সেডান",
		"2011":                     "২০১১",
		"Inline-4":                 "ইনলাইন-৪",
		"1400":                     "১৪০০ সিসি",
		"4":                        "৪ (ভালভ প্রতি সিলিন্ডার)",
		"Naturally Aspirated":      "ন্যাচারালি অ্যাসপিরেটেড",
		"Open":                     "ওপেন",
		"95 hp/ton":                "৯৫ হর্স পাওয়ার/টন",
		"14 km/L":                  "১৪ কিমি/লিটার",
		"20 km/L":                  "২০ কিমি/লিটার",
		"17 km/L":                  "১৭ কিমি/লিটার",
		"MacPherson Strut":         "ম্যাকফারসন স্ট্রাট",
		"Torsion Beam":             "টর্শন বিম",
		"Hydraulic Power Steering": "হাইড্রোলিক পাওয়ার স্টিয়ারিং",
		"Tilt":                     "টিল্ট",
		"15":                       "১৫",
		"Full Size":                "ফুল সাইজ",
		"No":                       "না",
		"Manual":                   "ম্যানুয়াল",
		"Variable Speed":           "ভেরিয়েবল স্পিড",
		"5":                        "৫",
		"Single Zone":              "সিঙ্গেল জোন",
		"Basic Radio":              "বেসিক রেডিও",
		"5 inch":                   "৫ ইঞ্চি",
		"Chevrolet":                "শেভ্রোলেট",
		"4":                        "৪",
		"Disc":                     "ডিস্ক",
		"Semi-Independent":         "সেমি-ইন্ডিপেন্ডেন্ট",
		"Euro 4":                   "ইউরো ৪",
		"Key":                      "কী",
		"Water Cooled":             "ওয়াটার কুলড",
		"Electronic":               "ইলেকট্রনিক",
		"10.5:1":                   "১০.৫:১",
		"DOHC":                     "ডিওএইচসি",
	}
}

// Seed implements the Seeder interface for Chevrolet Aveo
func (cas *ChevroletAveoSeeder) Seed(db *gorm.DB) error {
	productSlug := "chevrolet-aveo"
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
		"Variant":                        "LS",
		"Generation":                     "2nd Generation",
		"Segment":                        "Subcompact Sedan",
		"Launch Year":                    "2011",
		"Engine Configuration":           "Inline-4",
		"Displacement (cc)":              "1400",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "95 hp/ton",
		"Mileage City (km/L)":            "14 km/L",
		"Mileage Highway (km/L)":         "20 km/L",
		"Mileage Combined (km/L)":        "17 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Hydraulic Power Steering",
		"Steering Adjustment":            "Tilt",
		"Wheel Size (inch)":              "15",
		"Spare Wheel Type":               "Full Size",
		"DRL":                            "Halogen",
		"Fog Lamp Type":                  "No",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Manual",
		"Wiper Type":                     "Variable Speed",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "Single Zone",
		"Infotainment System":            "Basic Radio",
		"Infotainment Screen (inch)":     "5 inch",
		"Apple CarPlay":                  "No",
		"Android Auto":                   "No",
		"Sound System Brand":             "Chevrolet",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "4",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc",
		"Traction Control":               "No",
		"ESC":                            "No",
		"Hill Assist":                    "No",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "No",
		"Suspension Type":                "Semi-Independent",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Key",
		"Cooling System":                 "Water Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.5:1",
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
		if banglaValue, exists := cas.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Chevrolet Aveo specifications seeded successfully")
	return nil
}
