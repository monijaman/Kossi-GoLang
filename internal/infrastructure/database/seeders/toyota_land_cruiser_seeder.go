package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ToyotaLandCruiserSeeder seeds specifications for Toyota Land Cruiser
type ToyotaLandCruiserSeeder struct {
	BaseSeeder
}

// NewToyotaLandCruiserSeeder creates a new Toyota Land Cruiser seeder
func NewToyotaLandCruiserSeeder() *ToyotaLandCruiserSeeder {
	return &ToyotaLandCruiserSeeder{
		BaseSeeder: BaseSeeder{name: "Toyota Land Cruiser Specifications"},
	}
}

func (tlcs *ToyotaLandCruiserSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"3.5L Diesel V8":                 "৩.৫ লিটার ডিজেল ভি৮",
		"300 Series":                     "৩০০ সিরিজ",
		"E-Segment":                      "ই-সেগমেন্ট",
		"2024":                           "২০২৪",
		"SUV":                            "এসইউভি",
		"Premium Black":                  "প্রিমিয়াম কালো",
		"V8":                             "ভি৮",
		"3445":                           "৩৪৪৫ সিসি",
		"4":                              "৪",
		"282 hp":                         "২৮২ হর্স পাওয়ার",
		"282 hp @ 3000 rpm":              "৩০০০ আরপিএম এ ২৮২ হর্স পাওয়ার",
		"700 Nm @ 1600 rpm":              "১৬০০ আরপিএম এ ৭০০ এনএম",
		"Diesel":                         "ডিজেল",
		"Common Rail Direct Injection":   "কমন রেল সরাসরি ইনজেকশন",
		"7.8 seconds":                    "৭.৮ সেকেন্ড",
		"210 km/h":                       "২১০ কিমি/ঘণ্টা",
		"4WD: 10 km/L | AWD: 12 km/L":    "৪হুইল ড্রাইভ: ১০ কিমি/লিটার | অল-হুইল ড্রাইভ: ১২ কিমি/লিটার",
		"10 km/L":                        "১০ কিমি/লিটার",
		"12 km/L":                        "১২ কিমি/লিটার",
		"11 km/L":                        "১১ কিমি/লিটার",
		"10-Speed Automatic":             "১০-স্পীড অটোমেটিক",
		"Automatic":                      "অটোমেটিক",
		"10":                             "১০",
		"Four-Wheel Drive":               "ফোর-হুইল ড্রাইভ",
		"Multi-Plate":                    "মাল্টি-প্লেট",
		"4980 mm":                        "৪৯৮০ মিমি",
		"1980 mm":                        "১৯৮০ মিমি",
		"1940 mm":                        "১৯৪০ মিমি",
		"2850 mm":                        "২৮৫০ মিমি",
		"235 mm":                         "২৩৫ মিমি",
		"2700 kg":                        "২৭০০ কেজি",
		"850L":                           "৮৫০ লিটার",
		"110L":                           "১১০ লিটার",
		"275/50 R20":                     "২৭৫/৫০ আর২০",
		"Radial Tubeless":                "রেডিয়াল টিউবলেস",
		"Alloy":                          "অ্যালয়",
		"LED":                            "এলইডি",
		"Halogen":                        "হ্যালোজেন",
		"Yes":                            "হ্যাঁ",
		"No":                             "না",
		"Power":                          "পাওয়ার",
		"Panoramic":                      "প্যানোরামিক",
		"Climate Control":                "জলবায়ু নিয়ন্ত্রণ",
		"Four Zone":                      "ফোর জোন",
		"Touchscreen":                    "টাচস্ক্রীন",
		"12.3 inch":                      "১২.৩ ইঞ্চি",
		"Premium":                        "প্রিমিয়াম",
		"14":                             "১৪",
		"Dual":                           "ডুয়াল",
		"Front & Rear":                   "সামনে এবং পিছনে",
		"Eco, Normal, Sport, Sand, Rock": "ইকো, নরমাল, স্পোর্ট, স্যান্ড, রক",
		"Speed Sensitive":                "গতি সংবেদনশীল",
		"Electric":                       "বৈদ্যুতিক",
		"Liquid Cooled":                  "লিকুইড কুলড",
		"Electronic":                     "ইলেকট্রনিক",
		"DOHC":                           "ডিওএইচসি",
		"Open":                           "ওপেন",
		"104 hp/ton":                     "১০৪ হর্স পাওয়ার/টন",
		"3 Years":                        "৩ বছর",
		"5 Years":                        "৫ বছর",
		"Independent Double Wishbone":    "স্বাধীন ডাবল উইশবোন",
		"Double Wishbone":                "ডাবল উইশবোন",
		"Rack and Pinion":                "র্যাক এবং পিনিয়ন",
		"Collapsible":                    "সংকোচনযোগ্য",
		"Ventilated Disc (All Around)":   "ভেন্টিলেটেড ডিস্ক (সবদিকে)",
		"Euro 6":                         "ইউরো ৬",
		"15.4:1":                         "১৫.৪:১",
		"Turbocharged":                   "টার্বোচার্জড",
	}
}

// Seed inserts specification records for Toyota Land Cruiser
func (tlcs *ToyotaLandCruiserSeeder) Seed(db *gorm.DB) error {
	productSlug := "toyota-land-cruiser"
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
		"Variant":                        "3.5L Diesel V8",
		"Generation":                     "300 Series",
		"Segment":                        "E-Segment",
		"Launch Year":                    "2024",
		"Body Type":                      "SUV",
		"Color":                          "Premium Black",
		"Engine Type":                    "3.5L Diesel V8",
		"Engine Configuration":           "V8",
		"Displacement":                   "3445",
		"Number of Cylinders":            "8",
		"Horsepower":                     "282 hp",
		"Max Power":                      "282 hp @ 3000 rpm",
		"Max Torque":                     "700 Nm @ 1600 rpm",
		"Valves Per Cylinder":            "4",
		"Fuel Type":                      "Diesel",
		"Fuel System":                    "Common Rail Direct Injection",
		"Acceleration 0-100 km/h":        "7.8 seconds",
		"Top Speed":                      "210 km/h",
		"Mileage":                        "4WD: 10 km/L | AWD: 12 km/L",
		"Mileage City (km/L)":            "10 km/L",
		"Mileage Highway (km/L)":         "12 km/L",
		"Mileage Combined (km/L)":        "11 km/L",
		"Transmission":                   "10-Speed Automatic",
		"Gearbox":                        "Automatic",
		"Number of Gears":                "10",
		"Drive Type":                     "Four-Wheel Drive",
		"Clutch Type":                    "Multi-Plate",
		"Length":                         "4980 mm",
		"Width":                          "1980 mm",
		"Height":                         "1940 mm",
		"Wheelbase":                      "2850 mm",
		"Ground Clearance":               "235 mm",
		"Kerb Weight":                    "2700 kg",
		"Boot Space":                     "850L",
		"Fuel Tank Capacity":             "110L",
		"Fuel Capacity":                  "110L",
		"Tyre Size":                      "275/50 R20",
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
		"Seating Capacity":               "8",
		"Number of Seats":                "8",
		"Driver Seat Adjustment":         "Power",
		"Ventilated Seats":               "Yes",
		"Air Conditioning":               "Climate Control",
		"Climate Control":                "Four Zone",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "12.3 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Premium",
		"Number of Speakers":             "14",
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
		"Drive Modes":                    "Eco, Normal, Sport, Sand, Rock",
		"Connected Car Features":         "Yes",
		"OTA Updates":                    "Yes",
		"Bluetooth Connectivity":         "Yes",
		"Rear Camera":                    "Yes",
		"Cruise Control":                 "Yes",
		"Power Steering":                 "Speed Sensitive",
		"Touchscreen":                    "Yes",
		"Front Suspension":               "Independent Double Wishbone",
		"Rear Suspension":                "Double Wishbone",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Collapsible",
		"Suspension Type":                "Independent",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Engine Aspiration":              "Turbocharged",
		"Compression Ratio":              "15.4:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
		"Differential Type":              "Open",
		"Power to Weight (HP/ton)":       "104 hp/ton",
		"Vehicle Warranty (Years)":       "3 Years",
		"Engine Warranty (Years)":        "5 Years",
	}

	banglaTranslations := tlcs.getBanglaTranslations()

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
