package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HondaAmazeSeeder seeds specifications for Honda Amaze
type HondaAmazeSeeder struct {
	BaseSeeder
}

// NewHondaAmazeSeeder creates a new Honda Amaze seeder
func NewHondaAmazeSeeder() *HondaAmazeSeeder {
	return &HondaAmazeSeeder{
		BaseSeeder: BaseSeeder{name: "Honda Amaze Specifications"},
	}
}

func (has *HondaAmazeSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"1.2L":                            "১.২ লিটার",
		"2nd":                             "২য় প্রজন্ম",
		"B-Segment":                       "বি-সেগমেন্ট",
		"2021":                            "২০২১",
		"Sedan":                           "সেডান",
		"Golden Brown Metallic":           "গোল্ডেন ব্রাউন মেটালিক",
		"Inline":                          "ইনলাইন",
		"1199":                            "১১৯৯ সিসি",
		"4 (Cylinder)":                    "৪ (সিলিন্ডার)",
		"90 hp":                           "৯০ হর্স পাওয়ার",
		"90 hp @ 6000 rpm":                "৬০০০ আরপিএম এ ৯০ হর্স পাওয়ার",
		"110 Nm @ 4800 rpm":               "৪৮০০ আরপিএম এ ১১০ এনএম",
		"Petrol":                          "পেট্রোল",
		"Multi-Point Injection":           "মাল্টি-পয়েন্ট ইনজেকশন",
		"13.2 seconds":                    "১৩.২ সেকেন্ড",
		"160 km/h":                        "১৬০ কিমি/ঘণ্টা",
		"18 km/L":                         "১৮ কিমি/লিটার",
		"20 km/L":                         "২০ কিমি/লিটার",
		"19 km/L":                         "১৯ কিমি/লিটার",
		"Manual":                          "ম্যানুয়াল (ট্রান্সমিশন)",
		"Manual (ORVM)":                   "ম্যানুয়াল (ওআরভিএম)",
		"Manual (Wiper)":                  "ম্যানুয়াল (ওয়াইপার)",
		"Manual (Driver Seat Adjustment)": "ম্যানুয়াল (ড্রাইভার সিট অ্যাডজাস্টমেন্ট)",
		"Manual (Air Conditioning)":       "ম্যানুয়াল (এয়ার কন্ডিশনিং)",
		"5-Speed Manual":                  "৫-স্পিড ম্যানুয়াল",
		"Front-Wheel Drive":               "সামনের চাকা চালিত",
		"Single Clutch":                   "একক ক্লাচ",
		"3995 mm":                         "৩৯৯৫ মিমি",
		"1680 mm":                         "১৬৮০ মিমি",
		"1501 mm":                         "১৫০১ মিমি",
		"2470 mm":                         "২৪৭০ মিমি",
		"170 mm":                          "১৭০ মিমি",
		"905 kg":                          "৯০৫ কেজি",
		"420L":                            "৪২০ লিটার",
		"35L":                             "৩৫ লিটার",
		"175/65 R14":                      "১৭৫/৬৫ আর১৪",
		"Radial Tubeless":                 "রেডিয়াল টিউবলেস",
		"Steel":                           "স্টিল",
		"Halogen":                         "হ্যালোজেন",
		"Yes":                             "হ্যাঁ",
		"No":                              "না",
		"Touchscreen":                     "টাচস্ক্রীন",
		"7 inch":                          "৭ ইঞ্চি",
		"Standard":                        "স্ট্যান্ডার্ড",
		"4 (Speaker)":                     "৪ (স্পিকার)",
		"Dual":                            "ডুয়াল",
		"Disc (Front) / Drum (Rear)":      "ডিস্ক (সামনে) / ড্রাম (পিছনে)",
		"5":                               "৫",
	}
}

// Seed inserts specification records for Honda Amaze
func (has *HondaAmazeSeeder) Seed(db *gorm.DB) error {
	productSlug := "honda-amaze"
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
		"Brand":                             310,
		"Model Name":                        316,
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
		"Brand":                       "Honda",
		"Model Name":                  "Amaze",
		"Variant":                        "1.2L",
		"Generation":                     "2nd",
		"Segment":                        "B-Segment",
		"Launch Year":                    "2021",
		"Body Type":                      "Sedan",
		"Color":                          "Golden Brown Metallic",
		"Engine Type":                    "1.2L",
		"Engine Configuration":           "Inline",
		"Displacement":                   "1199",
		"Number of Cylinders":            "4 (Cylinder)",
		"Horsepower":                     "90 hp",
		"Max Power":                      "90 hp @ 6000 rpm",
		"Max Torque":                     "110 Nm @ 4800 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Multi-Point Injection",
		"Acceleration 0-100 km/h":        "13.2 seconds",
		"Top Speed":                      "160 km/h",
		"Mileage":                        "18 km/L",
		"Mileage City (km/L)":            "18 km/L",
		"Mileage Highway (km/L)":         "20 km/L",
		"Mileage Combined (km/L)":        "19 km/L",
		"Transmission":                   "Manual",
		"Gearbox":                        "5-Speed Manual",
		"Number of Gears":                "5",
		"Drive Type":                     "Front-Wheel Drive",
		"Clutch Type":                    "Single Clutch",
		"Length":                         "3995 mm",
		"Width":                          "1680 mm",
		"Height":                         "1501 mm",
		"Wheelbase":                      "2470 mm",
		"Ground Clearance":               "170 mm",
		"Kerb Weight":                    "905 kg",
		"Boot Space":                     "420L",
		"Fuel Tank Capacity":             "35L",
		"Fuel Capacity":                  "35L",
		"Tyre Size":                      "175/65 R14",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Steel",
		"Headlight Type":                 "Halogen",
		"DRL":                            "No",
		"Fog Lamp Type":                  "Halogen",
		"Alloy Wheels":                   "No",
		"Sunroof Type":                   "No",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Manual (ORVM)",
		"Wiper Type":                     "Manual (Wiper)",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Manual (Driver Seat Adjustment)",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Manual (Air Conditioning)",
		"Climate Control":                "No",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "7 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Standard",
		"Number of Speakers":             "4 (Speaker)",
		"Ambient Lighting":               "No",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Drum (Rear)",
		"Traction Control":               "No",
		"ESC":                            "No",
		"Hill Hold":                      "No",
		"ISOFIX Mounts":                  "Yes",
		"Camera Type":                    "Rear Camera",
		"Adaptive Cruise Control":        "No",
		"Lane Keep Assist":               "No",
		"Automatic Emergency Braking":    "No",
		"Blind Spot Monitor":             "No",
		"Parking Sensors":                "Rear",
		"Keyless Entry":                  "No",
		"Push Button Start":              "No",
		"Digital Instrument Cluster":     "No",
		"Heads Up Display":               "No",
		"Drive Modes":                    "No",
		"Connected Car Features":         "No",
		"OTA Updates":                    "No",
		"Bluetooth Connectivity":         "Yes",
		"Rear Camera":                    "Yes",
		"Cruise Control":                 "No",
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
		"Compression Ratio":              "10.0:1",
		"Valve Configuration":            "SOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "100 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := has.getBanglaTranslations()

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
