package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaHiaceSeeder seeds specifications for Toyota Hiace
type ToyotaHiaceSeeder struct {
	BaseSeeder
}

// NewToyotaHiaceSeeder creates a new Toyota Hiace seeder
func NewToyotaHiaceSeeder() *ToyotaHiaceSeeder {
	return &ToyotaHiaceSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Hiace Specifications"},
	}
}

func (ths *ToyotaHiaceSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.8L Diesel":                  "২.৮ লিটার ডিজেল",
		"4th":                          "৪র্থ প্রজন্ম",
		"M-Segment":                    "এম-সেগমেন্ট",
		"2024":                         "২০২৪",
		"MPV":                          "এমপিভি",
		"Commercial White":             "কমার্শিয়াল সাদা",
		"Inline":                       "ইনলাইন",
		"2755":                         "২৭৫৫ সিসি",
		"4":                            "৪",
		"177 hp":                       "১৭৭ হর্স পাওয়ার",
		"177 hp @ 3400 rpm":            "৩৪০০ আরপিএম এ ১৭৭ হর্স পাওয়ার",
		"450 Nm @ 1600 rpm":            "১৬০০ আরপিএম এ ৪৫০ এনএম",
		"Diesel":                       "ডিজেল",
		"Common Rail Direct Injection": "কমন রেল সরাসরি ইনজেকশন",
		"12.5 seconds":                 "১২.৫ সেকেন্ড",
		"160 km/h":                     "১৬০ কিমি/ঘণ্টা",
		"RWD: 10 km/L | 4WD: 12 km/L":  "রিয়ার-হুইল ড্রাইভ: ১০ কিমি/লিটার | ৪হুইল ড্রাইভ: ১২ কিমি/লিটার",
		"10 km/L":                      "১০ কিমি/লিটার",
		"12 km/L":                      "১২ কিমি/লিটার",
		"11 km/L":                      "১১ কিমি/লিটার",
		"6-Speed Manual":               "৬-স্পীড ম্যানুয়াল",
		"Manual (Transmission)":        "ম্যানুয়াল (ট্রান্সমিশন)",
		"6":                            "৬",
		"Rear-Wheel Drive":             "রিয়ার-হুইল ড্রাইভ",
		"Single Clutch":                "একক ক্লাচ",
		"5265 mm":                      "৫২৬৫ মিমি",
		"1950 mm":                      "১৯৫০ মিমি",
		"1990 mm":                      "১৯৯০ মিমি",
		"3210 mm":                      "৩২১০ মিমি",
		"185 mm":                       "১৮৫ মিমি",
		"2000 kg":                      "২০০০ কেজি",
		"4000L":                        "৪০০০ লিটার",
		"70L":                          "৭০ লিটার",
		"195/70 R15":                   "১৯৫/৭০ আর১৫",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Steel":                        "স্টিল",
		"Halogen":                      "হ্যালোজেন",
		"No (DRL)":                     "না (ডিআরএল)",
		"Yes":                          "হ্যাঁ",
		"Power":                        "পাওয়ার",
		"No (Fog Lamp)":                "না (ফগ ল্যাম্প)",
		"Manual (Wiper)":               "ম্যানুয়াল (ওয়াইপার)",
		"Touchscreen":                  "টাচস্ক্রীন",
		"7 inch":                       "৭ ইঞ্চি",
		"Commercial":                   "কমার্শিয়াল",
		"4 (Speakers)":                 "৪ (স্পিকার)",
		"No (Ambient Lighting)":        "না (অ্যাম্বিয়েন্ট লাইটিং)",
		"Front & Rear":                 "সামনে এবং পিছনে",
		"Normal":                       "নরমাল",
		"Speed Sensitive":              "গতি সংবেদনশীল",
		"Electric":                     "বৈদ্যুতিক",
		"Liquid Cooled":                "লিকুইড কুলড",
		"Electronic":                   "ইলেকট্রনিক",
		"DOHC":                         "ডিওএইচসি",
		"Open":                         "ওপেন",
		"89 hp/ton":                    "৮৯ হর্স পাওয়ার/টন",
		"3 Years":                      "৩ বছর",
		"5 Years":                      "৫ বছর",
		"Leaf Spring":                  "লিফ স্প্রিং",
		"Leaf Spring (Rear)":           "লিফ স্প্রিং (পিছনে)",
		"Rack and Pinion":              "র্যাক এবং পিনিয়ন",
		"Collapsible":                  "সংকোচনযোগ্য",
		"Disc (Front) & Drum (Rear)":   "ডিস্ক (সামনে) এবং ড্রাম (পিছনে)",
		"Euro 4":                       "ইউরো ৪",
		"16.0:1":                       "১৬.০:১",
		"Turbocharged":                 "টার্বোচার্জড",
	}
}

// Seed inserts specification records for Toyota Hiace
func (ths *ToyotaHiaceSeeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-hiace"
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
		"Model Name":                  "Hiace",
		"Variant":                        "2.8L Diesel",
		"Generation":                     "4th",
		"Segment":                        "M-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "MPV",
		"Color":                          "Commercial White",
		"Engine Type":                    "2.8L Diesel",
		"Engine Configuration":           "Inline",
		"Displacement":                   "2755",
		"Number of Cylinders":            "4",
		"Horsepower":                     "177 hp",
		"Max Power":                      "177 hp @ 3400 rpm",
		"Max Torque":                     "450 Nm @ 1600 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Diesel",
		"Fuel System":                    "Common Rail Direct Injection",
		"Acceleration 0-100 km/h":        "12.5 seconds",
		"Top Speed":                      "160 km/h",
		"Mileage":                        "RWD: 10 km/L | 4WD: 12 km/L",
		"Mileage City (km/L)":            "10 km/L",
		"Mileage Highway (km/L)":         "12 km/L",
		"Mileage Combined (km/L)":        "11 km/L",
		"Transmission":                   "6-Speed Manual",
		"Gearbox":                        "Manual",
		"Number of Gears":                "6",
		"Drive Type":                     "Rear-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "5265 mm",
		"Width":                          "1950 mm",
		"Height":                         "1990 mm",
		"Wheelbase":                      "3210 mm",
		"Ground Clearance":               "185 mm",
		"Kerb Weight":                    "2000 kg",
		"Boot Space":                     "4000L",
		"Fuel Tank Capacity":             "70L",
		"Fuel Capacity":                  "70L",
		"Tyre Size":                      "195/70 R15",
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
		"Seating Capacity":               "9",
		"Number of Seats":                "9",
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
		"Front Suspension":               "Leaf Spring",
		"Rear Suspension":                "Leaf Spring",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Rigid",
		"Emission Standard":              "Euro 4",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Turbocharged",
		"Compression Ratio":              "16.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "89 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := ths.getBanglaTranslations()

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
