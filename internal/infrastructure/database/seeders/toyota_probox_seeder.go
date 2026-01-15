package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaProboxSeeder seeds specifications for Toyota Probox
type ToyotaProboxSeeder struct {
	BaseSeeder
}

// NewToyotaProboxSeeder creates a new Toyota Probox seeder
func NewToyotaProboxSeeder() *ToyotaProboxSeeder {
	return &ToyotaProboxSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Probox Specifications"},
	}
}

func (tps *ToyotaProboxSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.5L Petrol":                    "১.৫ লিটার পেট্রল",
		"2nd":                            "২য় প্রজন্ম",
		"B-Segment":                      "বি-সেগমেন্ট",
		"2024":                           "২০২৪",
		"Wagon":                          "ওয়াগন",
		"Commercial White":               "কমার্শিয়াল সাদা",
		"Inline":                         "ইনলাইন",
		"1496":                           "১৪৯৬ সিসি",
		"4":                              "৪",
		"109 hp":                         "১০৯ হর্স পাওয়ার",
		"109 hp @ 6000 rpm":              "৬০০০ আরপিএম এ ১০৯ হর্স পাওয়ার",
		"140 Nm @ 4200 rpm":              "৪২০০ আরপিএম এ ১৪০ এনএম",
		"Petrol":                         "পেট্রল",
		"Multi-Point Fuel Injection":     "মাল্টি-পয়েন্ট ফুয়েল ইনজেকশন",
		"11.5 seconds":                   "১১.৫ সেকেন্ড",
		"170 km/h":                       "১৭০ কিমি/ঘণ্টা",
		"Manual: 17 km/L | CVT: 19 km/L": "ম্যানুয়াল: ১৭ কিমি/লিটার | সিভিটি: ১৯ কিমি/লিটার",
		"17 km/L":                        "১৭ কিমি/লিটার",
		"19 km/L":                        "১৯ কিমি/লিটার",
		"18 km/L":                        "১৮ কিমি/লিটার",
		"5-Speed Manual":                 "৫-স্পীড ম্যানুয়াল",
		"Manual":                         "ম্যানুয়াল",
		"5":                              "৫",
		"Front-Wheel Drive":              "ফ্রন্ট-হুইল ড্রাইভ",
		"Single Clutch":                  "একক ক্লাচ",
		"4240 mm":                        "৪২৪০ মিমি",
		"1690 mm":                        "১৬৯০ মিমি",
		"1510 mm":                        "১৫১০ মিমি",
		"2550 mm":                        "২৫৫০ মিমি",
		"150 mm":                         "১৫০ মিমি",
		"1200 kg":                        "১২০০ কেজি",
		"620L":                           "৬২০ লিটার",
		"42L":                            "৪২ লিটার",
		"165/70 R14":                     "১৬৫/৭০ আর১৪",
		"Radial Tubeless":                "রেডিয়াল টিউবলেস",
		"Steel":                          "স্টিল",
		"Halogen":                        "হ্যালোজেন",
		"No (DRL)":                       "না (ডিআরএল)",
		"Yes":                            "হ্যাঁ",
		"Power":                          "পাওয়ার",
		"No (Fog Lamp)":                  "না (ফগ ল্যাম্প)",
		"Manual (Wiper)":                 "ম্যানুয়াল (ওয়াইপার)",
		"Touchscreen":                    "টাচস্ক্রীন",
		"7 inch":                         "৭ ইঞ্চি",
		"Commercial":                     "কমার্শিয়াল",
		"4 (Speakers)":                   "৪ (স্পিকার)",
		"No (Ambient Lighting)":          "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Front & Rear":                   "সামনে এবং পিছনে",
		"Normal":                         "নরমাল",
		"Speed Sensitive":                "গতি সংবেদনশীল",
		"Electric":                       "বৈদ্যুতিক",
		"Water Cooled":                   "ওয়াটার কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"DOHC":                           "ডিওএইচসি",
		"Open":                           "ওপেন",
		"91 hp/ton":                      "৯১ হর্স পাওয়ার/টন",
		"3 Years":                        "৩ বছর",
		"5 Years":                        "৫ বছর",
		"MacPherson Strut":               "ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":                   "টর্শন বিম",
		"Rack and Pinion":                "র্যাক এবং পিনিয়ন",
		"Collapsible":                    "সংকোচনযোগ্য",
		"Disc (Front) & Drum (Rear)":     "ডিস্ক (সামনে) এবং ড্রাম (পিছনে)",
		"Euro 4":                         "ইউরো ৪",
		"10.5:1":                         "১০.৫:১",
		"Naturally Aspirated":            "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Toyota Probox
func (tps *ToyotaProboxSeeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-probox"
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
		"Variant":                        "1.5L Petrol",
		"Generation":                     "2nd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "Wagon",
		"Color":                          "Commercial White",
		"Engine Type":                    "1.5L Petrol",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1496",
		"Number of Cylinders":            "4",
		"Horsepower":                     "109 hp",
		"Max Power":                      "109 hp @ 6000 rpm",
		"Max Torque":                     "140 Nm @ 4200 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Fuel Injection",
		"Acceleration 0-100 km/h":        "11.5 seconds",
		"Top Speed":                      "170 km/h",
		"Mileage":                        "Manual: 17 km/L | CVT: 19 km/L",
		"Mileage City (km/L)":            "17 km/L",
		"Mileage Highway (km/L)":         "19 km/L",
		"Mileage Combined (km/L)":        "18 km/L",
		"Transmission":                   "5-Speed Manual",
		"Gearbox":                        "Manual",
		"Number of Gears":                "5",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4240 mm",
		"Width":                          "1690 mm",
		"Height":                         "1510 mm",
		"Wheelbase":                      "2550 mm",
		"Ground Clearance":               "150 mm",
		"Kerb Weight":                    "1200 kg",
		"Boot Space":                     "620L",
		"Fuel Tank Capacity":             "42L",
		"Fuel Capacity":                  "42L",
		"Tyre Size":                      "165/70 R14",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Steel",
		"Headlight Type":                 "Halogen",
		"DRL":                            "No",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Manual",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "7 inch",
		"Apple CarPlay":                  "No",
		"Android Auto":                   "No",
		"Sound System Brand":             "Commercial",
		"Number of Speakers":             "4",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) & Drum (Rear)",
		"Traction Control":               "No",
		"ESC":                            "No",
		"Hill Hold":                      "Yes",
		"ISOFIX Mounts":                  "Yes",
		"Camera Type":                    "Rear Camera",
		"Adaptive Cruise Control":        "No",
		"Lane Keep Assist":               "No",
		"Automatic Emergency Braking":    "No",
		"Blind Spot Monitor":             "No",
		"Parking Sensors":                "Front & Rear",
		"Keyless Entry":                  "No",
		"Push Button Start":              "No",
		"Digital Instrument Cluster":     "No",
		"Heads Up Display":               "No",
		"Drive Modes":                    "Normal",
		"Connected Car Features":         "No",
		"OTA Updates":                    "No",
		"Bluetooth Connectivity":         "Yes",
		"Rear Camera":                    "Yes",
		"Cruise Control":                 "No",
		"Power Steering":                 "Speed Sensitive",
		"Touchscreen":                    "Yes",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Torsion Beam",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Semi-independent",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Water Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Naturally Aspirated",
		"Compression Ratio":              "10.5:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "91 hp/ton",
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
