package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleSuzukiGSXR150ABS seeds specifications/options for product 'suzuki-gsx-r150-abs'
type SpecificationSeederMotorcycleSuzukiGSXR150ABS struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleSuzukiGSXR150ABS creates a new seeder instance
func NewSpecificationSeederMotorcycleSuzukiGSXR150ABS() *SpecificationSeederMotorcycleSuzukiGSXR150ABS {
	return &SpecificationSeederMotorcycleSuzukiGSXR150ABS{BaseSeeder: BaseSeeder{name: "Specifications for Suzuki GSX-R150 ABS"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMotorcycleSuzukiGSXR150ABS) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Yes (Dual Channel)":                 "হ্যাঁ (ডুয়াল চ্যানেল)",
		"147.3 cc":                           "১৪৭.৩ সিসি",
		"Wet multi-plate":                    "ওয়েট মাল্টি-প্লেট",
		"Liquid-cooled":                      "লিকুইড-কুলড",
		"Chain":                              "চেইন",
		"Single-cylinder, 4-stroke, DOHC":    "সিঙ্গেল-সিলিন্ডার, ৪-স্ট্রোক, DOHC",
		"11 L":                               "১১ লিটার",
		"Petrol":                             "পেট্রোল",
		"Manual":                             "ম্যানুয়াল",
		"160 mm":                             "১৬০ মিমি",
		"LED":                                "LED",
		"19.2 HP":                            "১৯.২ এইচপি",
		"Digital CDI":                        "ডিজিটাল CDI",
		"133 kg":                             "১৩৩ কেজি",
		"2020 mm":                            "২০২০ মিমি",
		"19.2 HP @ 10500 rpm":                "১৯.২ এইচপি @ ১০৫০০ আরপিএম",
		"14.0 Nm @ 9000 rpm":                 "১৪.০ এনএম @ ৯০০০ আরপিএম",
		"40 km/l (approx)":                   "৪০ কিমি/লিটার (আনুমানিক)",
		"6":                                  "৬",
		"2":                                  "২",
		"Electric Start":                     "ইলেকট্রিক স্টার্ট",
		"Telescopic front, Mono shock rear":  "টেলিস্কোপিক ফ্রন্ট, মনো শক রিয়ার",
		"145 km/h":                           "১৪৫ কিমি/ঘণ্টা",
		"14.0 Nm":                            "১৪.০ এনএম",
		"6-speed":                            "৬-স্পিড",
		"90/80-17 (front), 130/70-17 (rear)": "৯০/৮০-১৭ (সামনে), ১৩০/৭০-১৭ (পেছনে)",
		"Tubeless":                           "টিউবলেস",
		"DOHC":                               "DOHC",
		"1300 mm":                            "১৩০০ মিমি",
		"685 mm":                             "৬৮৫ মিমি",
		"Front Disc / Rear Disc":             "ফ্রন্ট ডিস্ক / রিয়ার ডিস্ক",
		"Disc":                               "ডিস্ক",
		"Alloy":                              "অ্যালয়",
		"17 inch":                            "১৭ ইঞ্চি",
		"785 mm":                             "৭৮৫ মিমি",
		"12V":                                "১২ভি",
		"Fuel Injection":                     "ফুয়েল ইঞ্জেকশন",
		"BS6":                                "বিএস৬",
		"11.5:1":                             "১১.৫:১",
		"Full LCD":                           "ফুল LCD",
		"Yes":                                "হ্যাঁ",
		"Iridium":                            "ইরিডিয়াম",
		"12V DC":                             "১২ভি ডিসি",
		"62.0 mm":                            "৬২.০ মিমি",
		"48.8 mm":                            "৪৮.৮ মিমি",
		"Wet sump":                           "ওয়েট সাম্প",
		"Gear type":                          "গিয়ার টাইপ",
		"115 mm":                             "১১৫ মিমি",
		"1.2 L":                              "১.২ লিটার",
		"Twin Spar":                          "টুইন স্পার",
		"2.5 L":                              "২.৫ লিটার",
		"90/80-17":                           "৯০/৮০-১৭",
		"130/70-17":                          "১৩০/৭০-১৭",
		"Metallic Triton Blue, Matte Black":  "মেটালিক ট্রাইটন নীল, ম্যাট কালো",
		"Digital":                            "ডিজিটাল",
		"8 Ah":                               "৮ এএইচ",
		"290 mm front, 187 mm rear":          "২৯০ মিমি সামনে, ১৮৭ মিমি পেছনে",
		"Telescopic SFF-BP":                  "টেলিস্কোপিক SFF-BP",
		"Radiator":                           "রেডিয়েটর",
	}
}

// Seed inserts specification records for the product identified by slug 'suzuki-gsx-r150-abs'
func (s *SpecificationSeederMotorcycleSuzukiGSXR150ABS) Seed(db *gorm.DB) error {
	productSlug := "suzuki-gsx-r150-abs"

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

	// Override model-specific values for Suzuki GSX-R150 ABS
	specs["ABS (Anti-lock Braking System)"] = "Yes (Dual Channel)"
	specs["CC (Cubic Capacity)"] = "147.3 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Liquid-cooled"
	specs["Displacement"] = "147.3 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, DOHC"
	specs["Fuel Capacity"] = "11 L"
	specs["Fuel Tank Capacity"] = "11 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "160 mm"
	specs["Headlight Type"] = "LED"
	specs["Horsepower"] = "19.2 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "133 kg"
	specs["Length"] = "2020 mm"
	specs["Max Power"] = "19.2 HP @ 10500 rpm"
	specs["Max Torque"] = "14.0 Nm @ 9000 rpm"
	specs["Mileage"] = "40 km/l (approx)"
	specs["Number of Gears"] = "6"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric Start"
	specs["Suspension Type"] = "Telescopic front, Mono shock rear"
	specs["Top Speed"] = "145 km/h"
	specs["Torque"] = "14.0 Nm"
	specs["Transmission"] = "6-speed"
	specs["Tyre Size"] = "90/80-17 (front), 130/70-17 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "DOHC"
	specs["Wheelbase"] = "1300 mm"
	specs["Width"] = "685 mm"
	specs["Brake Type"] = "Front Disc / Rear Disc"
	specs["Front Brake Type"] = "Disc"
	specs["Rear Brake Type"] = "Disc"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "785 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Fuel Injection"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "11.5:1"
	specs["Instrument Cluster"] = "Full LCD"
	specs["DRL (Daytime Running Light)"] = "Yes"
	specs["Spark Plug Type"] = "Iridium"
	specs["Electrical System"] = "12V DC"
	specs["Bore"] = "62.0 mm"
	specs["Stroke"] = "48.8 mm"
	specs["Lubrication System"] = "Wet sump"
	specs["Oil Pump Type"] = "Gear type"
	specs["Front Suspension Travel"] = "115 mm"
	specs["Rear Suspension Travel"] = "115 mm"
	specs["Oil Capacity"] = "1.2 L"
	specs["Frame Type"] = "Twin Spar"
	specs["Reserve Fuel Capacity"] = "2.5 L"
	specs["Side Stand Engine Cutoff"] = "Yes"
	specs["Front Tyre Size"] = "90/80-17"
	specs["Rear Tyre Size"] = "130/70-17"
	specs["Color Options"] = "Metallic Triton Blue, Matte Black"
	specs["Speedometer Type"] = "Digital"
	specs["Battery Type"] = "8 Ah"
	specs["Brake Diameter"] = "290 mm front, 187 mm rear"
	specs["Suspension Material"] = "Telescopic SFF-BP"
	specs["Wheel Type"] = "Alloy"
	specs["Cooling Type"] = "Radiator"

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
