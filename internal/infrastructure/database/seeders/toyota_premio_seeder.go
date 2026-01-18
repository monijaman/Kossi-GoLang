package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaPremioSeeder seeds specifications for Toyota Premio
type ToyotaPremioSeeder struct {
	BaseSeeder
}

// NewToyotaPremioSeeder creates a new Toyota Premio seeder
func NewToyotaPremioSeeder() *ToyotaPremioSeeder {
	return &ToyotaPremioSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Premio Specifications"},
	}
}

func (tps *ToyotaPremioSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.8L Petrol":                    "১.৮ লিটার পেট্রল",
		"3rd":                            "৩য় প্রজন্ম",
		"D-Segment":                      "ডি-সেগমেন্ট",
		"2024":                           "২০২৪",
		"Sedan":                          "সেডান",
		"Premium Black":                  "প্রিমিয়াম কালো",
		"Inline":                         "ইনলাইন",
		"1798":                           "১৭৯৮ সিসি",
		"4":                              "৪",
		"140 hp":                         "১৪০ হর্স পাওয়ার",
		"140 hp @ 6200 rpm":              "৬২০০ আরপিএম এ ১৪০ হর্স পাওয়ার",
		"173 Nm @ 4000 rpm":              "৪০০০ আরপিএম এ ১৭৩ এনএম",
		"Petrol":                         "পেট্রল",
		"Direct Injection":               "সরাসরি ইনজেকশন",
		"9.8 seconds":                    "৯.৮ সেকেন্ড",
		"190 km/h":                       "১৯০ কিমি/ঘণ্টা",
		"Manual: 14 km/L | CVT: 16 km/L": "ম্যানুয়াল: ১৪ কিমি/লিটার | সিভিটি: ১৬ কিমি/লিটার",
		"14 km/L":                        "১৪ কিমি/লিটার",
		"18 km/L":                        "১৮ কিমি/লিটার",
		"15 km/L":                        "১৫ কিমি/লিটার",
		"Manual / CVT":                   "ম্যানুয়াল / সিভিটি",
		"CVT":                            "সিভিটি",
		"CVT (Gearbox)":                  "সিভিটি (গিয়ারবক্স)",
		"Front-Wheel Drive":              "সামনের চাকা চালিত",
		"Single Clutch":                  "একক ক্লাচ",
		"4600 mm":                        "৪৬০০ মিমি",
		"1695 mm":                        "১৬৯৫ মিমি",
		"1460 mm":                        "১৪৬০ মিমি",
		"2700 mm":                        "২৭০০ মিমি",
		"135 mm":                         "১৩৫ মিমি",
		"1350 kg":                        "১৩৫০ কেজি",
		"520L":                           "৫২০ লিটার",
		"60L":                            "৬০ লিটার",
		"195/65 R15":                     "১৯৫/৬৫ আর১৫",
		"Radial Tubeless":                "রেডিয়াল টিউবলেস",
		"Alloy":                          "অ্যালয়",
		"LED":                            "এলইডি",
		"Halogen":                        "হ্যালোজেন",
		"Yes":                            "হ্যাঁ",
		"No":                             "না",
		"Power":                          "পাওয়ার",
		"Climate Control":                "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                      "ডুয়াল জোন",
		"Touchscreen":                    "টাচস্ক্রীন",
		"8 inch":                         "৮ ইঞ্চি",
		"Premium":                        "প্রিমিয়াম",
		"6":                              "৬",
		"Dual":                           "ডুয়াল",
		"Front & Rear":                   "সামনে এবং পিছনে",
		"Multiple Modes":                 "একাধিক মোড",
		"Speed Sensitive":                "গতি সংবেদনশীল",
		"Electric":                       "বৈদ্যুতিক",
		"Liquid Cooled":                  "লিকুইড কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"DOHC":                           "ডিওএইচসি",
		"Open":                           "ওপেন",
		"104 hp/ton":                     "১০৪ হর্স পাওয়ার/টন",
		"3 Years":                        "৩ বছর",
		"5 Years":                        "৫ বছর",
		"Independent MacPherson Strut":   "স্বাধীন ম্যাকফার্সন স্ট্রাট",
		"Double Wishbone":                "ডাবল উইশবোন",
		"Rack and Pinion":                "র্যাক এবং পিনিয়ন",
		"Collapsible":                    "সংকোচনযোগ্য",
		"Ventilated Disc (Front) / Solid Disc (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / সলিড ডিস্ক (পিছনে)",
		"Euro 4":              "ইউরো ৪",
		"10.0:1":              "১০.০:১",
		"Naturally Aspirated": "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Toyota Premio
func (tps *ToyotaPremioSeeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-premio"
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
		"Driver Seat Adjustment":         760,
		"Ventilated Seats":               761,
		"Infotainment Screen (inch)":     762,
		"Apple CarPlay":                  763,
		"Android Auto":                   764,
		"Sound System Brand":             769,
		"Number of Speakers":             770,
		"Ambient Lighting":               771,
		"EBD":                            772,
		"Traction Control":               773,
		"ESC":                            774,
		"Hill Hold":                      775,
		"ISOFIX Mounts":                  776,
		"Camera Type":                    777,
		"Adaptive Cruise Control":        778,
		"Lane Keep Assist":               779,
		"Automatic Emergency Braking":    780,
		"Blind Spot Monitor":             781,
		"Keyless Entry":                  782,
		"Push Button Start":              783,
		"Digital Instrument Cluster":     784,
		"Heads Up Display":               785,
		"Drive Modes":                    786,
		"Connected Car Features":         787,
		"OTA Updates":                    788,
		"Vehicle Warranty (Years)":       789,
		"Engine Warranty (Years)":        790,
		"Body Type":                      89,
		"Color":                          311,
		"Engine Type":                    101,
		"Displacement":                   98,
		"Number of Cylinders":            117,
		"Horsepower":                     109,
		"Max Power":                      114,
		"Max Torque":                     115,
		"Fuel Type":                      105,
		"Fuel System":                    103,
		"Acceleration 0-100 km/h":        83,
		"Top Speed":                      127,
		"Mileage":                        116,
		"Transmission":                   130,
		"Gearbox":                        106,
		"Number of Gears":                118,
		"Drive Type":                     99,
		"Clutch Type":                    94,
		"Length":                         113,
		"Width":                          136,
		"Height":                         25,
		"Wheelbase":                      135,
		"Ground Clearance":               107,
		"Kerb Weight":                    112,
		"Boot Space":                     90,
		"Fuel Tank Capacity":             104,
		"Fuel Capacity":                  102,
		"Tyre Size":                      131,
		"Tyre Type":                      132,
		"Headlight Type":                 108,
		"Seating Capacity":               123,
		"Number of Seats":                119,
		"Air Conditioning":               86,
		"Climate Control":                93,
		"Infotainment System":            111,
		"ABS (Anti-lock Braking System)": 84,
		"Airbags":                        85,
		"Brake Type":                     91,
		"Parking Sensors":                120,
		"Bluetooth Connectivity":         88,
		"Rear Camera":                    122,
		"Cruise Control":                 96,
		"Power Steering":                 121,
		"Touchscreen":                    129,
		"Suspension Type":                126,
		"Emission Standard":              100,
		"Starting System":                124,
		"Cooling System":                 95,
		"Ignition Type":                  110,
		"Compression Ratio":              329,
		"Valve Configuration":            133,
		"Valve Per Cylinder":             134,
	}

	specs := map[string]string{
		"Brand":                       "Toyota",
		"Model Name":                  "Premio",
		"Variant":                        "1.8L Petrol",
		"Generation":                     "3rd",
		"Segment":                        "D-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "Sedan",
		"Color":                          "Premium Black",
		"Engine Type":                    "1.8L Petrol",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1798",
		"Number of Cylinders":            "4",
		"Horsepower":                     "140 hp",
		"Max Power":                      "140 hp @ 6200 rpm",
		"Max Torque":                     "173 Nm @ 4000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "9.8 seconds",
		"Top Speed":                      "190 km/h",
		"Mileage":                        "Manual: 14 km/L | CVT: 16 km/L",
		"Mileage City (km/L)":            "14 km/L",
		"Mileage Highway (km/L)":         "18 km/L",
		"Mileage Combined (km/L)":        "15 km/L",
		"Transmission":                   "Manual / CVT",
		"Gearbox":                        "CVT",
		"Number of Gears":                "CVT",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4600 mm",
		"Width":                          "1695 mm",
		"Height":                         "1460 mm",
		"Wheelbase":                      "2700 mm",
		"Ground Clearance":               "135 mm",
		"Kerb Weight":                    "1350 kg",
		"Boot Space":                     "520L",
		"Fuel Tank Capacity":             "60L",
		"Fuel Capacity":                  "60L",
		"Tyre Size":                      "195/65 R15",
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
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Climate Control",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Ventilated Disc (Front) / Solid Disc (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Hold":                      "Yes",
		"ISOFIX Mounts":                  "Yes",
		"Camera Type":                    "Rear Camera",
		"Adaptive Cruise Control":        "No",
		"Lane Keep Assist":               "No",
		"Automatic Emergency Braking":    "No",
		"Blind Spot Monitor":             "No",
		"Parking Sensors":                "Front & Rear",
		"Keyless Entry":                  "Yes",
		"Push Button Start":              "Yes",
		"Digital Instrument Cluster":     "Yes",
		"Heads Up Display":               "No",
		"Drive Modes":                    "Multiple Modes",
		"Connected Car Features":         "Yes",
		"OTA Updates":                    "No",
		"Bluetooth Connectivity":         "Yes",
		"Rear Camera":                    "Yes",
		"Cruise Control":                 "Yes",
		"Power Steering":                 "Speed Sensitive",
		"Touchscreen":                    "Yes",
		"Front Suspension":               "Independent MacPherson Strut",
		"Rear Suspension":                "Double Wishbone",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Naturally Aspirated",
		"Compression Ratio":              "10.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "104 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := tps.getBanglaTranslations()

	for key, value := range specs {
		specKeyID, exists := existingKeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingKeyMapping: '%s' (used in product: %s)", key, productSlug)
			continue
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, specKeyID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: specKeyID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}
				if sModel.ID == 0 {
					if err := db.Where("product_id = ? AND specification_key_id = ? AND value = ?", productID, specKeyID, value).First(sModel).Error; err != nil {
						return err
					}
				}
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					translation := &models.SpecificationTranslationModel{
						SpecificationID: sModel.ID,
						Locale:          "bn",
						Value:           banglaValue,
					}
					if err := db.Create(translation).Error; err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			banglaValue, exists := banglaTranslations[value]
			if exists && banglaValue != "" {
				var existingTranslation models.SpecificationTranslationModel
				if err := db.Where("specification_id = ? AND locale = ?", existing.ID, "bn").First(&existingTranslation).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						translation := &models.SpecificationTranslationModel{
							SpecificationID: existing.ID,
							Locale:          "bn",
							Value:           banglaValue,
						}
						if err := db.Create(translation).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			}
		}
	}

	return nil
}
