package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HondaCRVSeeder seeds specifications for Honda CR-V
type HondaCRVSeeder struct {
	BaseSeeder
}

// NewHondaCRVSeeder creates a new Honda CR-V seeder
func NewHondaCRVSeeder() *HondaCRVSeeder {
	return &HondaCRVSeeder{
		BaseSeeder: BaseSeeder{name: "Honda CR-V Specifications"},
	}
}

func (hcvs *HondaCRVSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L":                                  "২.০ লিটার",
		"5th":                                   "৫ম প্রজন্ম",
		"C-Segment":                             "সি-সেগমেন্ট",
		"2022":                                  "২০২২",
		"SUV":                                   "এসইউভি",
		"Modern Steel Metallic":                 "মডার্ন স্টিল মেটালিক",
		"Inline":                                "ইনলাইন",
		"1997":                                  "১৯৯৭ সিসি",
		"4":                                     "৪",
		"154 hp":                                "১৫৪ হর্স পাওয়ার",
		"154 hp @ 6500 rpm":                     "৬৫০০ আরপিএম এ ১৫৪ হর্স পাওয়ার",
		"189 Nm @ 4300 rpm":                     "৪৩০০ আরপিএম এ ১৮৯ এনএম",
		"Petrol":                                "পেট্রোল",
		"Direct Injection":                      "সরাসরি ইনজেকশন",
		"10.5 seconds":                          "১০.৫ সেকেন্ড",
		"190 km/h":                              "১৯০ কিমি/ঘণ্টা",
		"14 km/L":                               "১৪ কিমি/লিটার",
		"16 km/L":                               "১৬ কিমি/লিটার",
		"15 km/L":                               "১৫ কিমি/লিটার",
		"CVT":                                   "সিভিটি",
		"Front-Wheel Drive":                     "সামনের চাকা চালিত",
		"All-Wheel Drive":                       "সমস্ত চাকা চালিত",
		"Single Clutch":                         "একক ক্লাচ",
		"4587 mm":                               "৪৫৮৭ মিমি",
		"1855 mm":                               "১৮৫৫ মিমি",
		"1679 mm":                               "১৬৭৯ মিমি",
		"2662 mm":                               "২৬৬২ মিমি",
		"198 mm":                                "১৯৮ মিমি",
		"1545 kg":                               "১৫৪৫ কেজি",
		"522L":                                  "৫২২ লিটার",
		"57L":                                   "৫৭ লিটার",
		"235/60 R18":                            "২৩৫/৬০ আর১৮",
		"Radial Tubeless":                       "রেডিয়াল টিউবলেস",
		"Alloy":                                 "অ্যালয়",
		"LED":                                   "এলইডি",
		"Halogen":                               "হ্যালোজেন",
		"Yes":                                   "হ্যাঁ",
		"No":                                    "না",
		"Power":                                 "পাওয়ার",
		"Panoramic":                             "প্যানোরামিক",
		"Automatic":                             "স্বয়ংক্রিয়",
		"Climate Control":                       "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                             "ডুয়াল জোন",
		"Touchscreen":                           "টাচস্ক্রীন",
		"7 inch":                                "৭ ইঞ্চি",
		"9 inch":                                "৯ ইঞ্চি",
		"Premium":                               "প্রিমিয়াম",
		"6":                                     "৬",
		"8":                                     "৮",
		"Dual":                                  "ডুয়াল",
		"Front & Rear":                          "সামনে এবং পিছনে",
		"Multiple Modes":                        "একাধিক মোড",
		"Speed Sensitive":                       "গতি সংবেদনশীল",
		"Electric":                              "বৈদ্যুতিক",
		"Liquid Cooled":                         "লিকুইড কুলড",
		"Electronic":                            "ইলেকট্রনিক",
		"DOHC":                                  "ডিওএইচসি",
		"Open":                                  "ওপেন",
		"100 hp/ton":                            "১০০ হর্স পাওয়ার/টন",
		"3 Years":                               "৩ বছর",
		"5 Years":                               "৫ বছর",
		"Independent MacPherson Strut":          "স্বাধীন ম্যাকফার্সন স্ট্রাট",
		"Multi-Link":                            "মাল্টি-লিঙ্ক",
		"Rack and Pinion":                       "র্যাক এবং পিনিয়ন",
		"Collapsible":                           "সংকোচনযোগ্য",
		"Ventilated Disc (Front) / Drum (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"Euro 6":                                "ইউরো ৬",
		"10.3:1":                                "১০.৩:১",
		"Naturally Aspirated":                   "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Honda CR-V
func (hcvs *HondaCRVSeeder) Seed(db *gorm.DB) error {
	productSlug := "honda-cr-v"
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
		"Variant":                        "2.0L",
		"Generation":                     "5th",
		"Segment":                        "C-Segment",
		"Launch Year":                    "2022",
		"Body Type":                      "SUV",
		"Color":                          "Modern Steel Metallic",
		"Engine Type":                    "2.0L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1997",
		"Number of Cylinders":            "4",
		"Horsepower":                     "154 hp",
		"Max Power":                      "154 hp @ 6500 rpm",
		"Max Torque":                     "189 Nm @ 4300 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "10.5 seconds",
		"Top Speed":                      "190 km/h",
		"Mileage":                        "14 km/L",
		"Mileage City (km/L)":            "14 km/L",
		"Mileage Highway (km/L)":         "16 km/L",
		"Mileage Combined (km/L)":        "15 km/L",
		"Transmission":                   "CVT",
		"Gearbox":                        "CVT",
		"Number of Gears":                "CVT",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4587 mm",
		"Width":                          "1855 mm",
		"Height":                         "1679 mm",
		"Wheelbase":                      "2662 mm",
		"Ground Clearance":               "198 mm",
		"Kerb Weight":                    "1545 kg",
		"Boot Space":                     "522L",
		"Fuel Tank Capacity":             "57L",
		"Fuel Capacity":                  "57L",
		"Tyre Size":                      "235/60 R18",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Panoramic",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Climate Control",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "9 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "8",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Ventilated Disc (Front) / Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Hold":                      "Yes",
		"ISOFIX Mounts":                  "Yes",
		"Camera Type":                    "360 Degree Camera",
		"Adaptive Cruise Control":        "Yes",
		"Lane Keep Assist":               "Yes",
		"Automatic Emergency Braking":    "Yes",
		"Blind Spot Monitor":             "Yes",
		"Parking Sensors":                "Front & Rear",
		"Keyless Entry":                  "Yes",
		"Push Button Start":              "Yes",
		"Digital Instrument Cluster":     "Yes",
		"Heads Up Display":               "No",
		"Drive Modes":                    "Multiple Modes",
		"Connected Car Features":         "Yes",
		"OTA Updates":                    "Yes",
		"Bluetooth Connectivity":         "Yes",
		"Rear Camera":                    "Yes",
		"Cruise Control":                 "Adaptive",
		"Power Steering":                 "Speed Sensitive",
		"Touchscreen":                    "Yes",
		"Front Suspension":               "Independent MacPherson Strut",
		"Rear Suspension":                "Multi-Link",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Naturally Aspirated",
		"Compression Ratio":              "10.3:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "100 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := hcvs.getBanglaTranslations()

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
