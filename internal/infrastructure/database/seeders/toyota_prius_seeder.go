package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaPriusSeeder seeds specifications for Toyota Prius
type ToyotaPriusSeeder struct {
	BaseSeeder
}

// NewToyotaPriusSeeder creates a new Toyota Prius seeder
func NewToyotaPriusSeeder() *ToyotaPriusSeeder {
	return &ToyotaPriusSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Prius Specifications"},
	}
}

func (tps *ToyotaPriusSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.8L Hybrid":                  "১.৮ লিটার হাইব্রিড",
		"5th":                          "৫ম প্রজন্ম",
		"C-Segment":                    "সি-সেগমেন্ট",
		"2024":                         "২০২৪",
		"Hatchback":                    "হ্যাচব্যাক",
		"Pearl White":                  "পার্ল সাদা",
		"Inline":                       "ইনলাইন",
		"1798":                         "১৭৯৮ সিসি",
		"4":                            "৪",
		"172 hp":                       "১৭২ হর্স পাওয়ার",
		"98 hp @ 5200 rpm":             "৫২০০ আরপিএম এ ৯৮ হর্স পাওয়ার",
		"142 Nm @ 3600 rpm":            "৩৬০০ আরপিএম এ ১৪২ এনএম",
		"Hybrid":                       "হাইব্রিড",
		"Direct Injection":             "সরাসরি ইনজেকশন",
		"8.3 seconds":                  "৮.৩ সেকেন্ড",
		"180 km/h":                     "১৮০ কিমি/ঘণ্টা",
		"25 km/L":                      "২৫ কিমি/লিটার",
		"27 km/L":                      "২৭ কিমি/লিটার",
		"26 km/L":                      "২৬ কিমি/লিটার",
		"CVT":                          "সিভিটি",
		"e-CVT":                        "ই-সিভিটি",
		"CVT (Gearbox)":                "সিভিটি (গিয়ারবক্স)",
		"Front-Wheel Drive":            "সামনের চাকা চালিত",
		"Single Clutch":                "একক ক্লাচ",
		"4575 mm":                      "৪৫৭৫ মিমি",
		"1760 mm":                      "১৭৬০ মিমি",
		"1470 mm":                      "১৪৭০ মিমি",
		"2750 mm":                      "২৭৫০ মিমি",
		"135 mm":                       "১৩৫ মিমি",
		"1420 kg":                      "১৪২০ কেজি",
		"502L":                         "৫০২ লিটার",
		"43L":                          "৪৩ লিটার",
		"195/65 R15":                   "১৯৫/৬৫ আর১৫",
		"Radial Tubeless":              "রেডিয়াল টিউবলেস",
		"Alloy":                        "অ্যালয়",
		"LED":                          "এলইডি",
		"Halogen":                      "হ্যালোজেন",
		"Yes":                          "হ্যাঁ",
		"No":                           "না",
		"Power":                        "পাওয়ার",
		"Panoramic":                    "প্যানোরামিক",
		"Climate Control":              "জলবায়ু নিয়ন্ত্রণ",
		"Dual Zone":                    "ডুয়াল জোন",
		"Touchscreen":                  "টাচস্ক্রীন",
		"12.3 inch":                    "১২.৩ ইঞ্চি",
		"Premium":                      "প্রিমিয়াম",
		"10":                           "১০",
		"Dual":                         "ডুয়াল",
		"Front & Rear":                 "সামনে এবং পিছনে",
		"Eco, Normal, Sport, EV":       "ইকো, নরমাল, স্পোর্ট, ইভি",
		"Speed Sensitive":              "গতি সংবেদনশীল",
		"Electric":                     "বৈদ্যুতিক",
		"Liquid Cooled":                "লিকুইড কুলড",
		"Electronic":                   "ইলেকট্রনিক",
		"DOHC":                         "ডিওএইচসি",
		"Open":                         "ওপেন",
		"121 hp/ton":                   "১২১ হর্স পাওয়ার/টন",
		"3 Years":                      "৩ বছর",
		"8 Years":                      "৮ বছর",
		"Independent MacPherson Strut": "স্বাধীন ম্যাকফার্সন স্ট্রাট",
		"Torsion Beam":                 "টর্শন বিম",
		"Rack and Pinion":              "র্যাক এবং পিনিয়ন",
		"Collapsible":                  "সংকোচনযোগ্য",
		"Ventilated Disc (Front) / Solid Disc (Rear)": "ভেন্টিলেটেড ডিস্ক (সামনে) / সলিড ডিস্ক (পিছনে)",
		"Euro 5":              "ইউরো ৫",
		"13.0:1":              "১৩.০:১",
		"Naturally Aspirated": "ন্যাচারালি অ্যাসপিরেটেড",
		"0.75 kWh":            "০.৭৫ কিলোওয়াট আওয়ার",
		"53 kW":               "৫৩ কিলোওয়াট",
		"163 Nm":              "১৬৩ এনএম",
		"AC Charging":         "এসি চার্জিং",
	}
}

// Seed inserts specification records for Toyota Prius
func (tps *ToyotaPriusSeeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-prius"
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
		"Battery Capacity (kWh)":         741,
		"Motor Power (kW)":               742,
		"Motor Torque (Nm)":              743,
		"Charging Type":                  744,
	}

	specs := map[string]string{
		"Variant":                        "1.8L Hybrid",
		"Generation":                     "5th",
		"Segment":                        "C-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "Hatchback",
		"Color":                          "Pearl White",
		"Engine Type":                    "1.8L Hybrid",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1798",
		"Number of Cylinders":            "4",
		"Horsepower":                     "172 hp",
		"Max Power":                      "98 hp @ 5200 rpm",
		"Max Torque":                     "142 Nm @ 3600 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Hybrid",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "8.3 seconds",
		"Top Speed":                      "180 km/h",
		"Mileage":                        "25 km/L",
		"Mileage City (km/L)":            "25 km/L",
		"Mileage Highway (km/L)":         "27 km/L",
		"Mileage Combined (km/L)":        "26 km/L",
		"Transmission":                   "CVT",
		"Gearbox":                        "e-CVT",
		"Number of Gears":                "CVT",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "4575 mm",
		"Width":                          "1760 mm",
		"Height":                         "1470 mm",
		"Wheelbase":                      "2750 mm",
		"Ground Clearance":               "135 mm",
		"Kerb Weight":                    "1420 kg",
		"Boot Space":                     "502L",
		"Fuel Tank Capacity":             "43L",
		"Fuel Capacity":                  "43L",
		"Tyre Size":                      "195/65 R15",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Panoramic",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Climate Control",
		"Climate Control":                "Dual Zone",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "12.3 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "10",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Ventilated Disc (Front) / Solid Disc (Rear)",
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
		"Drive Modes":                    "Eco, Normal, Sport, EV",
		"Connected Car Features":         "Yes",
		"OTA Updates":                    "Yes",
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
		"Emission Standard":              "Euro 5",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Naturally Aspirated",
		"Compression Ratio":              "13.0:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "121 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "8 Years",
		"Battery Capacity (kWh)":         "0.75 kWh",
		"Motor Power (kW)":               "53 kW",
		"Motor Torque (Nm)":              "163 Nm",
		"Charging Type":                  "AC Charging",
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
