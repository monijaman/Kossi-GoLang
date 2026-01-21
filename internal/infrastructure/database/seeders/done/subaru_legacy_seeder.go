package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SubaruLegacySeeder seeds specifications for Subaru Legacy
type SubaruLegacySeeder struct {
	BaseSeeder
}

// NewSubaruLegacySeeder creates a new Subaru Legacy seeder
func NewSubaruLegacySeeder() *SubaruLegacySeeder {
	return &SubaruLegacySeeder{
		BaseSeeder: BaseSeeder{name: "Subaru Legacy Specifications"},
	}
}

func (sls *SubaruLegacySeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"2.5L":                                "২.৫ লিটার",
		"7th":                                 "৭ম প্রজন্ম",
		"D-Segment":                           "ডি-সেগমেন্ট",
		"2020":                                "২০২০",
		"Sedan":                               "সেডান",
		"Crystal White Pearl":                 "ক্রিস্টাল হোয়াইট পার্ল",
		"Boxer":                               "বক্সার",
		"2498":                                "২৪৯৮ সিসি",
		"4":                                   "৪ (সিলিন্ডার)",
		"182 hp":                              "১৮২ হর্স পাওয়ার",
		"182 hp @ 5800 rpm":                   "৫৮০০ আরপিএম এ ১৮২ হর্স পাওয়ার",
		"239 Nm @ 4400 rpm":                   "৪৪০০ আরপিএম এ ২৩৯ এনএম",
		"Petrol":                              "পেট্রোল",
		"Naturally Aspirated":                 "ন্যাচারালি অ্যাসপিরেটেড",
		"Direct Injection":                    "সরাসরি ইনজেকশন",
		"9.2 seconds":                         "৯.২ সেকেন্ড",
		"210 km/h":                            "২১০ কিমি/ঘণ্টা",
		"12 km/L":                             "১২ কিমি/লিটার",
		"15 km/L":                             "১৫ কিমি/লিটার",
		"13 km/L":                             "১৩ কিমি/লিটার",
		"CVT (Transmission)":                  "সিভিটি (ট্রান্সমিশন)",
		"Lineartronic CVT":                    "লিনিয়ারট্রনিক সিভিটি",
		"All-Wheel Drive":                     "সমস্ত চাকা চালিত",
		"Multi-plate Clutch":                  "মাল্টি-প্লেট ক্লাচ",
		"4875 mm":                             "৪৮৭৫ মিমি",
		"1840 mm":                             "১৮৪০ মিমি",
		"1500 mm":                             "১৫০০ মিমি",
		"2825 mm":                             "২৮২৫ মিমি",
		"1555 mm":                             "১৫৫৫ মিমি",
		"1590 kg":                             "১৫৯০ কেজি",
		"60 L":                                "৬০ লিটার",
		"225/50 R18":                          "২২৫/৫০ আর১৮",
		"Radial Tubeless":                     "রেডিয়াল টিউবলেস",
		"Alloy":                               "অ্যালয়",
		"LED (Headlamps)":                     "এলইডি (হেডল্যাম্পস)",
		"Yes (DRL)":                           "হ্যাঁ (ডিআরএল)",
		"LED (Fog Lamp)":                      "এলইডি (ফগ ল্যাম্প)",
		"Yes (Fog Lamp)":                      "হ্যাঁ (ফগ ল্যাম্প)",
		"No (Sunroof)":                        "না (সানরুফ)",
		"No (Roof Rails)":                     "না (রুফ রেলস)",
		"Power (ORVM)":                        "পাওয়ার (ওআরভিএম)",
		"Automatic (Wiper)":                   "অটোমেটিক (ওয়াইপার)",
		"5 (Seating)":                         "৫ (সিটিং)",
		"5 (Seats)":                           "৫ (সিটস)",
		"Electric (Driver Seat)":              "ইলেকট্রিক (ড্রাইভার সিট)",
		"No (Ventilated Seats)":               "না (ভেন্টিলেটেড সিটস)",
		"Automatic (Air Conditioning)":        "অটোমেটিক (এয়ার কন্ডিশনিং)",
		"Yes (Climate Control)":               "হ্যাঁ (ক্লাইমেট কন্ট্রোল)",
		"Touchscreen (Infotainment)":          "টাচস্ক্রিন (ইনফোটেইনমেন্ট)",
		"11.6 inch (Screen)":                  "১১.৬ ইঞ্চি (স্ক্রিন)",
		"Yes (Apple CarPlay)":                 "হ্যাঁ (অ্যাপল কারপ্লে)",
		"Yes (Android Auto)":                  "হ্যাঁ (অ্যান্ড্রয়েড অটো)",
		"Harman Kardon (Sound System)":        "হারম্যান কার্ডন (সাউন্ড সিস্টেম)",
		"12 (Speakers)":                       "১২ (স্পিকার)",
		"Yes (Ambient Lighting)":              "হ্যাঁ (অ্যাম্বিয়েন্ট লাইটিং)",
		"Yes (ABS)":                           "হ্যাঁ (এবিএস)",
		"Driver + Passenger + Side (Airbags)": "ড্রাইভার + প্যাসেঞ্জার + সাইড (এয়ারব্যাগস)",
		"Yes (EBD)":                           "হ্যাঁ (ইবিডি)",
		"Disc (Front) / Disc (Rear)":          "ডিস্ক (সামনে) / ডিস্ক (পিছনে)",
		"Yes (Traction Control)":              "হ্যাঁ (ট্র্যাকশন কন্ট্রোল)",
		"Yes (ESC)":                           "হ্যাঁ (ইএসসি)",
		"Yes (Hill Assist)":                   "হ্যাঁ (হিল অ্যাসিস্ট)",
		"Yes (Power Steering)":                "হ্যাঁ (পাওয়ার স্টিয়ারিং)",
		"Yes (Touchscreen)":                   "হ্যাঁ (টাচস্ক্রিন)",
		"Yes (Rear Camera)":                   "হ্যাঁ (রিয়ার ক্যামেরা)",
		"Keyless Entry":                       "কীলেস এন্ট্রি",
		"Push Button Start":                   "পুশ বাটন স্টার্ট",
		"Height Adjustable":                   "হাইট অ্যাডজাস্টেবল",
		"Rear AC Vents":                       "রিয়ার এসি ভেন্টস",
		"Steering Mounted Controls":           "স্টিয়ারিং মাউন্টেড কন্ট্রোলস",
		"Cruise Control":                      "ক্রুজ কন্ট্রোল",
		"Parking Sensors":                     "পার্কিং সেন্সরস",
		"Camera":                              "ক্যামেরা",
		"360 Degree Camera":                   "৩৬০ ডিগ্রি ক্যামেরা",
		"Auto Headlamps":                      "অটো হেডল্যাম্পস",
		"Rain Sensing Wipers":                 "রেইন সেন্সিং ওয়াইপারস",
		"Follow Me Home Headlamps":            "ফলো মি হোম হেডল্যাম্পস",
		"Yes (Sunroof)":                       "হ্যাঁ (সানরুফ)",
		"Panoramic Sunroof":                   "প্যানোরামিক সানরুফ",
		"Yes (Roof Rails)":                    "হ্যাঁ (রুফ রেলস)",
		"No (Alloy Wheels)":                   "না (অ্যালয় হুইলস)",
		"Yes (Alloy Wheels)":                  "হ্যাঁ (অ্যালয় হুইলস)",
		"18 inch":                             "১৮ ইঞ্চি",
		"MacPherson Strut":                    "ম্যাকফারসন স্ট্রাট",
		"Double Wishbone":                     "ডাবল উইশবোন",
		"Rack and Pinion":                     "র্যাক অ্যান্ড পিনিয়ন",
		"Tilt & Telescopic":                   "টিল্ট অ্যান্ড টেলিস্কোপিক",
		"18":                                  "১৮",
	}
}

// Seed implements the Seeder interface for Subaru Legacy
func (sls *SubaruLegacySeeder) Seed(db *gorm.DB) error {
	productSlug := "subaru-legacy"
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
		"Seating Capacity":               760,
		"Number of Seats":                761,
		"Driver Seat Adjustment":         762,
		"Ventilated Seats":               763,
		"Air Conditioning":               764,
		"Climate Control":                765,
		"Infotainment System":            766,
		"Infotainment Screen (inch)":     767,
		"Apple CarPlay":                  768,
		"Android Auto":                   769,
		"Sound System Brand":             770,
		"Number of Speakers":             771,
		"Ambient Lighting":               772,
		"ABS (Anti-lock Braking System)": 773,
		"Airbags":                        774,
		"EBD":                            775,
		"Brake Type":                     776,
		"Traction Control":               777,
		"ESC":                            778,
		"Hill Assist":                    779,
		"Power Steering":                 780,
		"Touchscreen":                    781,
		"Suspension Type":                782,
		"Emission Standard":              783,
		"Starting System":                784,
		"Cooling System":                 785,
		"Ignition Type":                  786,
		"Compression Ratio":              787,
		"Valve Configuration":            788,
		"Valve Per Cylinder":             789,
		"Displacement (cc)":              790,
	}

	specs := map[string]string{
		"Brand":                       "Subaru",
		"Model Name":                  "Legacy",
		"Variant":                        "2.5L Limited",
		"Generation":                     "7th",
		"Segment":                        "D-Segment",
		"Launch Year":                    "2020",
		"Engine Configuration":           "Boxer",
		"Displacement (cc)":              "2498",
		"Valves Per Cylinder":            "4",
		"Engine Aspiration":              "Naturally Aspirated",
		"Differential Type":              "Intercooled",
		"Power to Weight (HP/ton)":       "114 hp/ton",
		"Mileage City (km/L)":            "12 km/L",
		"Mileage Highway (km/L)":         "15 km/L",
		"Mileage Combined (km/L)":        "13 km/L",
		"Front Suspension":               "MacPherson Strut",
		"Rear Suspension":                "Double Wishbone",
		"Steering Type":                  "Rack and Pinion",
		"Steering Adjustment":            "Tilt & Telescopic",
		"Wheel Size (inch)":              "18",
		"Spare Wheel Type":               "Alloy",
		"DRL":                            "Yes",
		"Fog Lamp Type":                  "LED",
		"Alloy Wheels":                   "Yes",
		"Sunroof Type":                   "Panoramic Sunroof",
		"Roof Rails":                     "No",
		"ORVM Type":                      "Power",
		"Wiper Type":                     "Automatic",
		"Seating Capacity":               "5",
		"Number of Seats":                "5",
		"Driver Seat Adjustment":         "Electric",
		"Ventilated Seats":               "No",
		"Air Conditioning":               "Automatic",
		"Climate Control":                "Yes",
		"Infotainment System":            "Touchscreen",
		"Infotainment Screen (inch)":     "11.6 inch",
		"Apple CarPlay":                  "Yes",
		"Android Auto":                   "Yes",
		"Sound System Brand":             "Harman Kardon",
		"Number of Speakers":             "12",
		"Ambient Lighting":               "Yes",
		"ABS (Anti-lock Braking System)": "Yes",
		"Airbags":                        "Driver + Passenger + Side",
		"EBD":                            "Yes",
		"Brake Type":                     "Disc (Front) / Disc (Rear)",
		"Traction Control":               "Yes",
		"ESC":                            "Yes",
		"Hill Assist":                    "Yes",
		"Power Steering":                 "Yes",
		"Touchscreen":                    "Yes",
		"Suspension Type":                "MacPherson Strut / Double Wishbone",
		"Emission Standard":              "Euro 6",
		"Starting System":                "Electric",
		"Cooling System":                 "Liquid Cooled",
		"Ignition Type":                  "Electronic",
		"Compression Ratio":              "10.3:1",
		"Valve Configuration":            "DOHC",
		"Valve Per Cylinder":             "4",
	}

	// Create specifications
	for key, value := range specs {
		keyID, exists := existingKeyMapping[key]
		if !exists {
			log.Printf("⚠️  Specification key not found: %s", key)
			continue
		}

		spec := models.SpecificationModel{
			ProductID:          productID,
			SpecificationKeyID: keyID,
			Value:              value,
		}

		// Check if specification already exists
		var existingSpec models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, keyID).First(&existingSpec).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new specification
				if err := db.Create(&spec).Error; err != nil {
					log.Printf("❌ Failed to create specification for key %s: %v", key, err)
					continue
				}
			} else {
				log.Printf("❌ Error checking existing specification for key %s: %v", key, err)
				continue
			}
		} else {
			log.Printf("ℹ️  Specification already exists for key %s, skipping", key)
			continue
		}

		// Create translation if Bangla translation exists
		if banglaValue, exists := sls.getBanglaTranslations()[value]; exists {
			translation := models.SpecificationTranslationModel{
				SpecificationID: spec.ID,
				Locale:          "bn",
				Value:           banglaValue,
			}

			if err := db.Create(&translation).Error; err != nil {
				log.Printf("❌ Failed to create translation for key %s: %v", key, err)
			}
		}
	}

	log.Printf("✅ Subaru Legacy specifications seeded successfully")
	return nil
}
