package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MercedesBenzSClassSeeder seeds specifications for Mercedes-Benz S-Class
type MercedesBenzSClassSeeder struct {
	BaseSeeder
}

// NewMercedesBenzSClassSeeder creates a new Mercedes-Benz S-Class seeder
func NewMercedesBenzSClassSeeder() *MercedesBenzSClassSeeder {
	return &MercedesBenzSClassSeeder{
		BaseSeeder: BaseSeeder{name: "Mercedes-Benz S-Class Specifications"},
	}
}

func (mbscs *MercedesBenzSClassSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"S 500":                          "এস ৫০০",
		"7th Generation (W223)":          "৭ম প্রজন্ম (ডব্লিউ২২৩)",
		"Flagship Luxury Sedan":          "ফ্ল্যাগশিপ লাক্সারি সেডান",
		"2020":                           "২০২০",
		"V8":                             "ভি৮",
		"2999":                           "২৯৯৯ সিসি",
		"4":                              "৪ (ভালভ প্রতি সিলিন্ডার)",
		"Twin-Turbocharged":              "টুইন-টার্বোচার্জড",
		"Open":                           "ওপেন",
		"110 hp/ton":                     "১১০ হর্স পাওয়ার/টন",
		"8 km/L":                         "৮ কিমি/লিটার",
		"14 km/L":                        "১৪ কিমি/লিটার",
		"11 km/L":                        "১১ কিমি/লিটার",
		"Air Suspension":                 "এয়ার সাসপেনশন",
		"Electric Power Steering":        "ইলেকট্রিক পাওয়ার স্টিয়ারিং",
		"Tilt & Telescopic":              "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"19":                             "১৯",
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
		"Quad Zone":                      "কোয়াড জোন",
		"MBUX":                           "এমবিইউএক্স",
		"12.8 inch":                      "১২.৮ ইঞ্চি",
		"Burmester":                      "বার্মেস্টার",
		"30":                             "৩০",
		"Disc":                           "ডিস্ক",
		"Independent":                    "ইন্ডিপেন্ডেন্ট",
		"Euro 6d":                        "ইউরো ৬ডি",
		"Push Button":                    "পুশ বাটন",
		"Water Cooled":                   "ওয়াটার কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"10.5:1":                         "১০.৫:১",
		"DOHC":                           "ডিওএইচসি",
	}
}

// Seed implements the Seeder interface for Mercedes-Benz S-Class
func (mbscs *MercedesBenzSClassSeeder) Seed(db *gorm.DB) error {
	productSlug := "mercedes-benz-s-class"
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
		"Variant":                        "S 500",
		"Generation":                     "7th Generation (W223)",
		"Segment":                        "Flagship Luxury Sedan",
		"Launch Year":                    "2020",
		"Engine Configuration":           "V8",
		"Displacement (cc)":              "2999",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Twin-Turbocharged",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "110 hp/ton",
		"Mileage City (km/L)":            "8 km/L",
		"Mileage Highway (km/L)":         "14 km/L",
		"Mileage Combined (km/L)":        "11 km/L",
		"Front Suspension":               "Air Suspension",
		"Rear Suspension":                "Air Suspension",
		"Steering Type":                  "Electric Power Steering",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "19",
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
		"Climate Control":                "Quad Zone",
		"Infotainment System":            "MBUX",
		"Infotainment Screen (inch)":     "12.8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Burmester",
		"Number of Speakers":             "30",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "8",
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
		if banglaValue, exists := mbscs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mercedes-Benz S-Class specifications seeded successfully")
	return nil
}
