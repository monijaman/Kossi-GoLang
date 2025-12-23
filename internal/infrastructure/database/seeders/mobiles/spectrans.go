package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleBajajPlatina110ES seeds specifications/options for product 'bajaj-platina-110-es'
type SpecificationSeederMotorcycleBajajPlatina110ES struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleBajajPlatina110ES creates a new seeder instance
func NewSpecificationSeederMotorcycleBajajPlatina110ES() *SpecificationSeederMotorcycleBajajPlatina110ES {
	return &SpecificationSeederMotorcycleBajajPlatina110ES{BaseSeeder: BaseSeeder{name: "Specifications for Bajaj Platina 110 ES"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMotorcycleBajajPlatina110ES) getBanglaTranslations() map[string]string {
	return map[string]string{
		"No":                                "না",
		"115.45 cc":                         "১১৫.৪৫ সিসি",
		"Wet multi-plate":                   "ওয়েট মাল্টি-প্লেট",
		"Air-cooled":                        "এয়ার-কুলড",
		"Chain":                             "চেইন",
		"Single-cylinder, 4-stroke, SOHC":   "সিঙ্গেল-সিলিন্ডার, ৪-স্ট্রোক, SOHC",
		"11 L":                              "১১ লিটার",
		"Petrol":                            "পেট্রোল",
		"Manual":                            "ম্যানুয়াল",
		"180 mm":                            "১৮০ মিমি",
		"Halogen":                           "হ্যালোজেন",
		"8.4 HP":                            "৮.৪ এইচপি",
		"CDI":                               "সিডিআই",
		"119 kg":                            "১১৯ কেজি",
		"2035 mm":                           "২০৩৫ মিমি",
		"8.4 HP @ 7000 rpm":                 "৮.৪ এইচপি @ ৭০০০ আরপিএম",
		"9.7 Nm @ 5000 rpm":                 "৯.৭ এনএম @ ৫০০০ আরপিএম",
		"70-76 km/l":                        "৭০-৭৬ কিমি/লিটার",
		"4":                                 "৪",
		"2":                                 "২",
		"Electric Start":                    "ইলেকট্রিক স্টার্ট",
		"Telescopic front, Twin Shock rear": "টেলিস্কোপিক ফ্রন্ট, টুইন শক রিয়ার",
		"90 km/h":                           "৯০ কিমি/ঘণ্টা",
		"9.7 Nm":                            "৯.৭ এনএম",
		"4-speed":                           "৪-স্পিড",
		"2.75-18 (front), 3.00-18 (rear)":   "২.৭৫-১৮ (সামনে), ৩.০০-১৮ (পেছনে)",
		"Tube-type":                         "টিউব-টাইপ",
		"SOHC 2-valve":                      "SOHC ২-ভালভ",
		"1255 mm":                           "১২৫৫ মিমি",
		"710 mm":                            "৭১০ মিমি",
		"Front Drum / Rear Drum":            "ফ্রন্ট ড্রাম / রিয়ার ড্রাম",
		"Drum":                              "ড্রাম",
		"Spoke":                             "স্পোক",
		"18 inch":                           "১৮ ইঞ্চি",
		"805 mm":                            "৮০৫ মিমি",
		"12V 2.5Ah":                         "১২ভি ২.৫এএইচ",
		"Carburetor":                        "কার্বুরেটর",
		"BS6":                               "বিএস৬",
		"9.5:1":                             "৯.৫:১",
		"Analog":                            "অ্যানালগ",
		"12V 35W":                           "১২ভি ৩৫ডব্লিউ",
		"57 mm":                             "৫৭ মিমি",
		"45 mm":                             "৪৫ মিমি",
		"Wet sump":                          "ওয়েট সাম্প",
		"135 mm":                            "১৩৫ মিমি",
		"130 mm front, 110 mm rear":         "১৩০ মিমি সামনে, ১১০ মিমি পেছনে",
		"Telescopic":                        "টেলিস্কোপিক",
		"Twin Shock":                        "টুইন শক",
		"Yes":                               "হ্যাঁ",
		"Tubular Double Cradle":             "টিউবুলার ডাবল ক্র্যাডল",
		"2.75-18":                           "২.৭৫-১৮",
		"3.00-18":                           "৩.০০-১৮",
		"Ebony Black, Charcoal Black":       "ইবোনি ব্ল্যাক, চারকোল ব্ল্যাক",
		"LED":                               "এলইডি",
		"2.5 Ah":                            "২.৫ এএইচ",
		"Combi Brake System":                "কম্বি ব্রেক সিস্টেম",
	}
}

// Seed inserts specification records for the product identified by slug 'bajaj-platina-110-es'
func (s *SpecificationSeederMotorcycleBajajPlatina110ES) Seed(db *gorm.DB) error {
	productSlug := "bajaj-platina-110-es"

	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	productID := prod.ID

	specs := DefaultMotorcycleSpecs()
	banglaTranslations := s.getBanglaTranslations()

	// Override model-specific values for Bajaj Platina 110 ES
	specs["ABS (Anti-lock Braking System)"] = "No"
	specs["Brake System"] = "Combi Brake System"
	specs["CC (Cubic Capacity)"] = "115.45 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "115.45 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, SOHC"
	specs["Fuel Capacity"] = "11 L"
	specs["Fuel Tank Capacity"] = "11 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "180 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "8.4 HP"
	specs["Ignition Type"] = "CDI"
	specs["Kerb Weight"] = "119 kg"
	specs["Length"] = "2035 mm"
	specs["Max Power"] = "8.4 HP @ 7000 rpm"
	specs["Max Torque"] = "9.7 Nm @ 5000 rpm"
	specs["Mileage"] = "70-76 km/l"
	specs["Number of Gears"] = "4"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric Start"
	specs["Suspension Type"] = "Telescopic front, Twin Shock rear"
	specs["Top Speed"] = "90 km/h"
	specs["Torque"] = "9.7 Nm"
	specs["Transmission"] = "4-speed"
	specs["Tyre Size"] = "2.75-18 (front), 3.00-18 (rear)"
	specs["Tyre Type"] = "Tube-type"
	specs["Valve Configuration"] = "SOHC 2-valve"
	specs["Wheelbase"] = "1255 mm"
	specs["Width"] = "710 mm"
	specs["Brake Type"] = "Front Drum / Rear Drum"
	specs["Front Brake Type"] = "Drum"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Spoke"
	specs["Rim Size"] = "18 inch"
	specs["Seat Height"] = "805 mm"
	specs["Battery Capacity"] = "12V 2.5Ah"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "9.5:1"
	specs["Instrument Cluster"] = "Analog"
	specs["Electrical System"] = "12V 35W"
	specs["Bore"] = "57 mm"
	specs["Stroke"] = "45 mm"
	specs["Lubrication System"] = "Wet sump"
	specs["Front Suspension Travel"] = "135 mm"
	specs["Frame Type"] = "Tubular Double Cradle"
	specs["Front Tyre Size"] = "2.75-18"
	specs["Rear Tyre Size"] = "3.00-18"
	specs["Color Options"] = "Ebony Black, Charcoal Black"
	specs["Tail Light Type"] = "LED"
	specs["Brake Diameter"] = "130 mm front, 110 mm rear"
	specs["Battery Type"] = "2.5 Ah"

	for key, value := range specs {
		sk, err := CreateOrFindSpecificationKey(db, key)
		if err != nil {
			return err
		}

		var existing models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", productID, sk.ID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sModel := &models.SpecificationModel{
					ProductID:          productID,
					SpecificationKeyID: sk.ID,
					Value:              value,
					Status:             1,
				}
				if err := db.Create(sModel).Error; err != nil {
					return err
				}

				// Create Bangla translation for the specification
				banglaValue, exists := banglaTranslations[value]
				if exists && banglaValue != "" {
					var existingTranslation models.SpecificationTranslationModel
					if err := db.Where("specification_id = ? AND locale = ?", sModel.ID, "bn").First(&existingTranslation).Error; err != nil {
						if err == gorm.ErrRecordNotFound {
							translation := &models.SpecificationTranslationModel{
								SpecificationID: sModel.ID,
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
			} else {
				return err
			}
		} else {
			// If specification already exists, check and create Bangla translation if missing
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
