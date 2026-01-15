package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaRav4Seeder seeds specifications for Toyota RAV4
type ToyotaRav4Seeder struct {
	BaseSeeder
}

// NewToyotaRav4Seeder creates a new Toyota RAV4 seeder
func NewToyotaRav4Seeder() *ToyotaRav4Seeder {
	return &ToyotaRav4Seeder{
		BaseSeeder: BaseSeeder{name: "Toyota RAV4 Specifications"},
	}
}

func (trs *ToyotaRav4Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.5L Petrol":                  "২.৫ লিটার পেট্রল",
		"6th":                          "৬ষ্ঠ প্রজন্ম",
		"D-Segment":                    "ডি-সেগমেন্ট",
		"2024":                         "২০২৪",
		"SUV":                          "এসইউভি",
		"Premium White":                "প্রিমিয়াম সাদা",
		"Inline":                       "ইনলাইন",
		"2487":                         "২৪৮৭ সিসি",
		"4":                            "৪",
		"203 hp":                       "২০৩ হর্স পাওয়ার",
		"203 hp @ 6600 rpm":            "৬৬০০ আরপিএম এ ২০৩ হর্স পাওয়ার",
		"250 Nm @ 5000 rpm":            "৫০০০ আরপিএম এ ২৫০ এনএম",
		"Petrol":                       "পেট্রল",
		"Direct Injection":             "সরাসরি ইনজেকশন",
		"8.5 seconds":                  "৮.৫ সেকেন্ড",
		"200 km/h":                     "২০০ কিমি/ঘণ্টা",
		"AWD: 23 km/L | FWD: 25 km/L":  "অল-হুইল ড্রাইভ: ২৩ কিমি/লিটার | ফ্রন্ট-হুইল ড্রাইভ: ২৫ কিমি/লিটার",
		"23 km/L":                      "২৩ কিমি/লিটার",
		"25 km/L":                      "২৫ কিমি/লিটার",
		"24 km/L":                      "২৪ কিমি/লিটার",
		"8-Speed Automatic":            "৮-স্পীড অটোমেটিক",
		"Automatic":                    "অটোমেটিক",
		"8":                            "৮",
		"All-Wheel Drive":              "অল-হুইল ড্রাইভ",
		"Multi-Plate":                  "মাল্টি-প্লেট",
		"4600 mm":                      "৪৬০০ মিমি",
		"1855 mm":                      "১৮৫৫ মিমি",
		"1680 mm":                      "১৬৮০ মিমি",
		"2690 mm":                      "২৬৯০ মিমি",
		"200 mm":                       "২০০ মিমি",
		"1745 kg":                      "১৭৪৫ কেজি",
		"580L":                         "৫৮০ লিটার",
		"55L":                          "৫৫ লিটার",
		"235/55 R19":                   "২৩৫/৫৫ আর১৯",
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
		"10.5 inch":                    "১০.৫ ইঞ্চি",
		"Premium":                      "প্রিমিয়াম",
		"6":                            "৬",
		"Dual":                         "ডুয়াল",
		"Front & Rear":                 "সামনে এবং পিছনে",
		"Eco, Normal, Sport":           "ইকো, নরমাল, স্পোর্ট",
		"Speed Sensitive":              "গতি সংবেদনশীল",
		"Electric":                     "বৈদ্যুতিক",
		"Liquid Cooled":                "লিকুইড কুলড",
		"Electronic":                   "ইলেকট্রনিক",
		"DOHC":                         "ডিওএইচসি",
		"Open":                         "ওপেন",
		"116 hp/ton":                   "১১৬ হর্স পাওয়ার/টন",
		"3 Years":                      "৩ বছর",
		"5 Years":                      "৫ বছর",
		"Independent MacPherson Strut": "স্বাধীন ম্যাকফার্সন স্ট্রাট",
		"Double Wishbone":              "ডাবল উইশবোন",
		"Rack and Pinion":              "র্যাক এবং পিনিয়ন",
		"Collapsible":                  "সংকোচনযোগ্য",
		"Ventilated Disc (All Around)": "ভেন্টিলেটেড ডিস্ক (সবদিকে)",
		"Euro 5":                       "ইউরো ৫",
		"10.4:1":                       "১০.৪:১",
		"Naturally Aspirated":          "ন্যাচারালি অ্যাসপিরেটেড",
	}
}

// Seed inserts specification records for Toyota RAV4
func (trs *ToyotaRav4Seeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-rav4"
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
		"Variant":                        "2.5L Petrol",
		"Generation":                     "6th",
		"Segment":                        "D-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "SUV",
		"Color":                          "Premium White",
		"Engine Type":                    "2.5L Petrol",
		"Engine Configuration":           "Inline",
		"Displacement":                   "2487",
		"Number of Cylinders":            "4",
		"Horsepower":                     "203 hp",
		"Max Power":                      "203 hp @ 6600 rpm",
		"Max Torque":                     "250 Nm @ 5000 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Petrol",
		"Fuel System":                    "Direct Injection",
		"Acceleration 0-100 km/h":        "8.5 seconds",
		"Top Speed":                      "200 km/h",
		"Mileage":                        "AWD: 23 km/L | FWD: 25 km/L",
		"Mileage City (km/L)":            "23 km/L",
		"Mileage Highway (km/L)":         "25 km/L",
		"Mileage Combined (km/L)":        "24 km/L",
		"Transmission":                   "8-Speed Automatic",
		"Gearbox":                        "Automatic",
		"Number of Gears":                "8",
		"Drive Type":                     "All-Wheel Drive",
		"Clutch Type":                    "Multi-Plate",
		"Length":                         "4600 mm",
		"Width":                          "1855 mm",
		"Height":                         "1680 mm",
		"Wheelbase":                      "2690 mm",
		"Ground Clearance":               "200 mm",
		"Kerb Weight":                    "1745 kg",
		"Boot Space":                     "580L",
		"Fuel Tank Capacity":             "55L",
		"Fuel Capacity":                  "55L",
		"Tyre Size":                      "235/55 R19",
		"Tyre Type":                      "Radial Tubeless",
		"Spare Wheel Type":               "Alloy",
		"Headlight Type":                 "LED",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "Halogen",
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
		"Infotainment Screen (inch)":     "10.5 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "6",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Dual",
		"EBD":                            "Yes",
		"Brake Type":                     "Ventilated Disc (All Around)",
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
		"Front Suspension":               "Independent MacPherson Strut",
		"Rear Suspension":                "Double Wishbone",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Euro 5",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Naturally Aspirated",
		"Compression Ratio":              "10.4:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "116 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := trs.getBanglaTranslations()

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
