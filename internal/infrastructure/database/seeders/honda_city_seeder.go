package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HondaCitySeeder seeds specifications for Honda City
type HondaCitySeeder struct {
	BaseSeeder
}

// NewHondaCitySeeder creates a new Honda City seeder
func NewHondaCitySeeder() *HondaCitySeeder {
	return &HondaCitySeeder{
		BaseSeeder: BaseSeeder{name: "Honda City Specifications"},
	}
}

func (hcs *HondaCitySeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L":                         "১.৫ লিটার",
		"5th":                          "৫ম প্রজন্ম",
		"B-Segment":                    "বি-সেগমেন্ট",
		"2022":                         "২০২২",
		"Sedan":                        "সেডান",
		"Lunar Silver Metallic":        "লুনার সিলভার মেটালিক",
		"Inline":                       "ইনলাইন",
		"1498":                         "১৪৯৮ সিসি",
		"4 (Cylinder)":                 "৪ (সিলিন্ডার)",
		"121 hp":                       "১২১ হর্স পাওয়ার",
		"121 hp @ 6600 rpm":            "৬৬০০ আরপিএম এ ১২১ হর্স পাওয়ার",
		"145 Nm @ 4300 rpm":            "৪৩০০ আরপিএম এ ১৪৫ এনএম",
		"Petrol":                       "পেট্রোল",
		"Multi-Point Injection":        "মাল্টি-পয়েন্ট ইনজেকশন",
		"10.2 seconds":                 "১০.২ সেকেন্ড",
		"195 km/h":                     "১৯৫ কিমি/ঘণ্টা",
		"18 km/L":                      "১৮ কিমি/লিটার",
		"20 km/L":                      "২০ কিমি/লিটার",
		"19 km/L":                      "১৯ কিমি/লিটার",
		"CVT":                          "সিভিটি",
		"Manual":                       "ম্যানুয়াল",
		"6-Speed Manual":               "৬-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":            "সামনের চাকা চালিত",
		"Single Clutch":                "একক ক্লাচ",
		"4545 mm":                      "৪৫৪৫ মিমি",
		"1748 mm":                      "১৭৪৮ মিমি",
		"1489 mm":                      "১৪৮৯ মিমি",
		"2600 mm":                      "২৬০০ মিমি",
		"165 mm":                       "১৬৫ মিমি",
		"1107 kg":                      "১১০৭ কেজি",
		"506L":                         "৫০৬ লিটার",
		"40L":                          "৪০ লিটার",
		"185/60 R15":                   "১৮৫/৬০ আর১৫",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"Steel":                        "স্টিল",
		"LED":                          "এলইডি",
		"Halogen":                      "হ্যালোজেন",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"Power":                        "পাওয়ার",
		"Automatic":                    "স্বয়ংক্রিয়",
		"Climate Control":              "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                    "ডুয়াল জোন",
		"Touchscreen":                  "টাচস্ক্রীন",
		"8 inch":                       "৮ ইঞ্চি",
		"Premium":                      "প্রিমিয়াম",
		"4 (Speaker)":                  "৪ (স্পিকার)",
		"6":                            "৬",
		"Dual":                         "ডুয়াল",
		"Front & Rear":                 "সামনে এবং পিছনে",
		"Speed Sensitive":              "গতি সংবেদনশীল",
		"Electric":                     "বৈদ্যুতিক",
		"Liquid Cooled":                "লিকুইড কুলড",
		"Electronic":                   "ইলেকট্রনিক",
		"SOHC":                         "এসওএইচসি",
		"Open":                         "ওপেন",
		"109 hp/ton":                   "১০৯ হর্স পাওয়ার/টন",
		"3 Years":                      "৩ বছর",
		"5 Years":                      "৫ বছর",
		"Independent MacPherson Strut": "স্বাধীন ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":                 "টর্শন বিম",
		"Rack and Pinion":              "র্যাক এবং পিনিয়ন",
		"Collapsible":                  "সংকোচনযোগ্য",
		"Disc (Front) / Drum (Rear)":   "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 6":                       "ইউরো ৬",
		"10.3:1":                       "১০.৩:১",
		"Naturally Aspirated":          "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Honda City
func (hcs *HondaCitySeeder) Seed(db *gorm.DB) error {
	productSlug := "honda-city"
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
		"Variant":                        "1.5L",
		"Generation":                     "5th",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2022",
		"Body Type":                      "Sedan",
		"Color":                          "Lunar Silver Metallic",
		"Engine Type":                    "1.5L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1498",
		"Number of Cylinders":            "4 (Cylinder)",
		"Horsepower":                     "121 hp",
		"Max Power":                      "121 hp @ 6600 rpm",
		"Max Torque":                     "145 Nm @ 4300 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Injection",
		"Acceleration 0-100 km/h":        "10.2 seconds",
		"Top Speed":                      "195 km/h",
		"Mileage":                        "18 km/L",
		"Mileage City (km/L)":            "18 km/L",
		"Mileage Highway (km/L)":         "20 km/L",
		"Mileage Combined (km/L)":        "19 km/L",
		"Transmission":                   "CVT",
		"Gearbox":                        "CVT",
		"Number of Gears":                "CVT",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4545 mm",
		"Width":                          "1748 mm",
		"Height":                         "1489 mm",
		"Wheelbase":                      "2600 mm",
		"Ground Clearance":               "165 mm",
		"Kerb Weight":                    "1107 kg",
		"Boot Space":                     "506L",
		"Fuel Tank Capacity":             "40L",
		"Fuel Capacity":                  "40L",
		"Tyre Size":                      "185/60 R15",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Steel",
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
		"Air Conditioning":               "Climate Control",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "8 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "4 (Speaker)",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Hold":                      "Yes",
		"ISOFIX Mounts":                  "Yes",
		"Camera Type":                    "Rear Camera",
		"Adaptive Cruise Control":        "No",
		"Lane Keep Assist":               "No",
		"Automatic Emergency Braking":    "No",
		"Blind Spot Monitor":             "No",
		"Parking Sensors":                "Rear",
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
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Naturally Aspirated",
		"Compression Ratio":              "10.3:1",
		"Valve Configuration":            "SOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "109 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := hcs.getBanglaTranslations()

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
