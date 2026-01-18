package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaVoxySeeder seeds specifications for Toyota Voxy
type ToyotaVoxySeeder struct {
	BaseSeeder
}

// NewToyotaVoxySeeder creates a new Toyota Voxy seeder
func NewToyotaVoxySeeder() *ToyotaVoxySeeder {
	return &ToyotaVoxySeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Voxy Specifications"},
	}
}

func (tvs *ToyotaVoxySeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.0L Petrol":                    "২.০ লিটার পেট্রল",
		"3rd":                            "৩য় প্রজন্ম",
		"B-Segment":                      "বি-সেগমেন্ট",
		"2024":                           "২০২৪",
		"MPV":                            "এমপিভি",
		"Premium White":                  "প্রিমিয়াম সাদা",
		"Inline":                         "ইনলাইন",
		"1987":                           "১৯৮৭ সিসি",
		"4":                              "৪",
		"152 hp":                         "১৫২ হর্স পাওয়ার",
		"152 hp @ 6100 rpm":              "৬১০০ আরপিএম এ ১৫২ হর্স পাওয়ার",
		"193 Nm @ 4400 rpm":              "৪৪০০ আরপিএম এ ১৯৩ এনএম",
		"Petrol":                         "পেট্রল",
		"Direct Injection":               "ডাইরেক্ট ইনজেকশন",
		"11.2 seconds":                   "১১.২ সেকেন্ড",
		"180 km/h":                       "১৮০ কিমি/ঘণ্টা",
		"Manual: 13 km/L | CVT: 15 km/L": "ম্যানুয়াল: ১৩ কিমি/লিটার | সিভিটি: ১৫ কিমি/লিটার",
		"13 km/L":                        "১৩ কিমি/লিটার",
		"15 km/L":                        "১৫ কিমি/লিটার",
		"14 km/L":                        "১৪ কিমি/লিটার",
		"CVT":                            "সিভিটি",
		"Automatic":                      "অটোমেটিক",
		"8-speed":                        "৮-স্পীড",
		"Front-Wheel Drive":              "ফ্রন্ট-হুইল ড্রাইভ",
		"Single Clutch":                  "একক ক্লাচ",
		"4695 mm":                        "৪৬৯৫ মিমি",
		"1730 mm":                        "১৭৩০ মিমি",
		"1825 mm":                        "১৮২৫ মিমি",
		"2850 mm":                        "২৮৫০ মিমি",
		"160 mm":                         "১৬০ মিমি",
		"1650 kg":                        "১৬৫০ কেজি",
		"900L":                           "৯০০ লিটার",
		"55L":                            "৫৫ লিটার",
		"205/60 R16":                     "২০৫/৬০ আর১৬",
		"Radial Tubeless":                "রেডিয়াল টিউবলেস",
		"Alloy":                          "অ্যালয়",
		"LED":                            "এলইডি",
		"Yes (DRL)":                      "হ্যাঁ (ডিআরএল)",
		"No":                             "না",
		"Power":                          "পাওয়ার",
		"Yes (Fog Lamp)":                 "হ্যাঁ (ফগ ল্যাম্প)",
		"Auto":                           "অটো",
		"Touchscreen":                    "টাচস্ক্রিন",
		"10.5 inch":                      "১০.৫ ইঞ্চি",
		"Premium":                        "প্রিমিয়াম",
		"6":                              "৬",
		"Yes (Ambient Lighting)":         "হ্যাঁ (অ্যাম্বিয়েন্ট লাইটিং)",
		"Front & Rear":                   "সামনে এবং পিছনে",
		"Eco, Normal, Sport":             "ইকো, নরমাল, স্পোর্ট",
		"Speed Sensitive":                "গতি সংবেদনশীল",
		"Electric":                       "বৈদ্যুতিক",
		"Water Cooled":                   "ওয়াটার কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"DOHC":                           "ডিওএইচসি",
		"Open":                           "ওপেন",
		"92 hp/ton":                      "৯২ হর্স পাওয়ার/টন",
		"3 Years":                        "৩ বছর",
		"5 Years":                        "৫ বছর",
		"MacPherson Strut":               "ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":                   "টর্শন বিম",
		"Rack and Pinion":                "র্যাক এবং পিনিয়ন",
		"Collapsible":                    "সংকোচনযোগ্য",
		"Disc (Front) & Drum (Rear)":     "ডিস্ক (সামনে) এবং ড্রাম (পিছনে)",
		"Euro 4":                         "ইউরো ৪",
		"12.5:1":                         "১২.৫:১",
		"Naturally Aspirated":            "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Toyota Voxy
func (tvs *ToyotaVoxySeeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-voxy"
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
		"Model Name":                  "Voxy",
		"Variant":                        "2.0L Petrol",
		"Generation":                     "3rd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "MPV",
		"Color":                          "Premium White",
		"Engine Type":                    "2.0L Petrol",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1987",
		"Number of Cylinders":            "4",
		"Horsepower":                     "152 hp",
		"Max Power":                      "152 hp @ 6100 rpm",
		"Max Torque":                     "193 Nm @ 4400 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "11.2 seconds",
		"Top Speed":                      "180 km/h",
		"Mileage":                        "Manual: 13 km/L | CVT: 15 km/L",
		"Mileage City (km/L)":            "13 km/L",
		"Mileage Highway (km/L)":         "15 km/L",
		"Mileage Combined (km/L)":        "14 km/L",
		"Transmission":                   "CVT",
		"Gearbox":                        "Automatic",
		"Number of Gears":                "8-speed",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4695 mm",
		"Width":                          "1730 mm",
		"Height":                         "1825 mm",
		"Wheelbase":                      "2850 mm",
		"Ground Clearance":               "160 mm",
		"Kerb Weight":                    "1650 kg",
		"Boot Space":                     "900L",
		"Fuel Tank Capacity":             "55L",
		"Fuel Capacity":                  "55L",
		"Tyre Size":                      "205/60 R16",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Yes",
		"Roof Rails":                     "Yes",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Auto",
		"Seating Capacity":               "8",
		"Number of Seats":                "8",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Dual-zone",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "10.5 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) & Drum (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Hold":                      "Yes",
		"ISOFIX Mounts":                  "Yes",
		"Camera Type":                    "360° Camera",
		"Adaptive Cruise Control":        "Yes",
		"Lane Keep Assist":               "Yes",
		"Automatic Emergency Braking":    "Yes",
		"Blind Spot Monitor":             "Yes",
		"Parking Sensors":                "Front & Rear",
		"Keyless Entry":                  "Yes",
		"Push Button Start":              "Yes",
		"Digital Instrument Cluster":     "Yes",
		"Heads Up Display":               "Yes",
		"Drive Modes":                    "Eco, Normal, Sport",
		"Connected Car Features":         "Yes",
		"OTA Updates":                    "Yes",
		"Bluetooth Connectivity":         "Yes",
		"Rear Camera":                    "Yes",
		"Cruise Control":                 "Yes",
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
		"Compression Ratio":              "12.5:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "92 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := tvs.getBanglaTranslations()

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
