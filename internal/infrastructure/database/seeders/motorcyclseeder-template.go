package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationSeederMotorcycleBajajDiscover seeds specifications/options for product 'bajaj-discover'
type SpecificationSeederMotorcycleBajajDiscover struct {
	BaseSeeder
}

// NewSpecificationSeederMotorcycleBajajDiscover creates a new seeder instance
func NewSpecificationSeederMotorcycleBajajDiscover() *SpecificationSeederMotorcycleBajajDiscover {
	return &SpecificationSeederMotorcycleBajajDiscover{BaseSeeder: BaseSeeder{name: "Specifications for Bajaj Discover"}}
}

// getBanglaTranslations returns a map of English specification values to their Bangla translations
func (s *SpecificationSeederMotorcycleBajajDiscover) getBanglaTranslations() map[string]string {
	return map[string]string{
		"No (CBS)":                            "না (CBS)",
		"125 cc":                              "১২৫ সিসি",
		"Wet multi-plate":                     "ওয়েট মাল্টি-প্লেট",
		"Air-cooled":                          "এয়ার-কুলড",
		"124.6 cc":                            "১২৪.৬ সিসি",
		"Chain":                               "চেইন",
		"Single-cylinder, 4-stroke, DTS-i":    "সিঙ্গেল-সিলিন্ডার, ৪-স্ট্রোক, DTS-i",
		"11 L":                                "১১ লিটার",
		"Petrol":                              "পেট্রোল",
		"Manual":                              "ম্যানুয়াল",
		"165 mm":                              "১৬৫ মিমি",
		"Halogen":                             "হ্যালোজেন",
		"11 HP":                               "১১ এইচপি",
		"Digital CDI":                         "ডিজিটাল CDI",
		"135 kg":                              "১৩৫ কেজি",
		"2020 mm":                             "২০২০ মিমি",
		"11 HP @ 8000 rpm":                    "১১ এইচপি @ ৮০০০ আরপিএম",
		"11 Nm @ 6000 rpm":                    "১১ এনএম @ ৬০০০ আরপিএম",
		"55 km/l (approx)":                    "৫৫ কিমি/লিটার (আনুমানিক)",
		"5":                                   "৫",
		"2":                                   "২",
		"Electric & Kick":                     "ইলেকট্রিক ও কিক",
		"Telescopic front, Twin shock rear":   "টেলিস্কোপিক ফ্রন্ট, টুইন শক রিয়ার",
		"100 km/h":                            "১০০ কিমি/ঘণ্টা",
		"11 Nm":                               "১১ এনএম",
		"5-speed":                             "৫-স্পিড",
		"80/100-17 (front), 100/90-17 (rear)": "৮০/১০০-১৭ (সামনে), ১০০/৯০-১৭ (পেছনে)",
		"Tubeless":                            "টিউবলেস",
		"SOHC":                                "SOHC",
		"1305 mm":                             "১৩০৫ মিমি",
		"760 mm":                              "৭৬০ মিমি",
		"Front Disc / Rear Drum":              "ফ্রন্ট ডিস্ক / রিয়ার ড্রাম",
		"Disc":                                "ডিস্ক",
		"Drum":                                "ড্রাম",
		"Alloy":                               "অ্যালয়",
		"17 inch":                             "১৭ ইঞ্চি",
		"785 mm":                              "৭৮৫ মিমি",
		"12V":                                 "১২ভি",
		"Carburetor":                          "কার্বুরেটর",
		"BS6":                                 "বিএস৬",
		"9.5:1":                               "৯.৫:১",
		"Analog-Digital":                      "অ্যানালগ-ডিজিটাল",
		"No":                                  "না",
		"Standard":                            "স্ট্যান্ডার্ড",
		"12V DC":                              "১২ভি ডিসি",
		"56.5 mm":                             "৫৬.৫ মিমি",
		"49.8 mm":                             "৪৯.৮ মিমি",
		"Air":                                 "এয়ার",
		"Piston type":                         "পিস্টন টাইপ",
		"240/1260 mm":                         "২৪০/১২৬০ মিমি",
		"3.5 L / 2.7 bar":                     "৩.৫ লিটার / ২.৭ বার",
		"Pressed steel":                       "প্রেসড স্টিল",
		"8 L":                                 "৮ লিটার",
		"Yes":                                 "হ্যাঁ",
		"100/90-17":                           "১০০/৯০-১৭",
		"80/100-17":                           "৮০/১০০-১৭",
		"Standard (Black)":                    "স্ট্যান্ডার্ড (কালো)",
		"Analogue":                            "অ্যানালগ",
		"7 Ah":                                "৭ এএইচ",
		"130 mm front, 110 mm rear":           "১৩০ মিমি সামনে, ১১০ মিমি পেছনে",
		"Coil spring with hydraulic dampers":  "কয়েল স্প্রিং হাইড্রোলিক ড্যাম্পার সহ",
		"Spoke":                               "স্পোক",
		"Forced air (Natural air flow)":       "ফোর্সড এয়ার (প্রাকৃতিক বায়ু প্রবাহ)",
	}
}

// Seed inserts specification records for the product identified by slug 'bajaj-discover'
func (s *SpecificationSeederMotorcycleBajajDiscover) Seed(db *gorm.DB) error {
	productSlug := "bajaj-discover"

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

	// Override model-specific values for Bajaj Discover
	specs["ABS (Anti-lock Braking System)"] = "No (CBS)"
	specs["CC (Cubic Capacity)"] = "125 cc"
	specs["Clutch Type"] = "Wet multi-plate"
	specs["Cooling System"] = "Air-cooled"
	specs["Displacement"] = "124.6 cc"
	specs["Drive Type"] = "Chain"
	specs["Engine Type"] = "Single-cylinder, 4-stroke, DTS-i"
	specs["Fuel Capacity"] = "11 L"
	specs["Fuel Tank Capacity"] = "11 L"
	specs["Fuel Type"] = "Petrol"
	specs["Gearbox"] = "Manual"
	specs["Ground Clearance"] = "165 mm"
	specs["Headlight Type"] = "Halogen"
	specs["Horsepower"] = "11 HP"
	specs["Ignition Type"] = "Digital CDI"
	specs["Kerb Weight"] = "135 kg"
	specs["Length"] = "2020 mm"
	specs["Max Power"] = "11 HP @ 8000 rpm"
	specs["Max Torque"] = "11 Nm @ 6000 rpm"
	specs["Mileage"] = "55 km/l (approx)"
	specs["Number of Gears"] = "5"
	specs["Number of Seats"] = "2"
	specs["Seating Capacity"] = "2"
	specs["Starting System"] = "Electric & Kick"
	specs["Suspension Type"] = "Telescopic front, Twin shock rear"
	specs["Top Speed"] = "100 km/h"
	specs["Torque"] = "11 Nm"
	specs["Transmission"] = "5-speed"
	specs["Tyre Size"] = "80/100-17 (front), 100/90-17 (rear)"
	specs["Tyre Type"] = "Tubeless"
	specs["Valve Configuration"] = "SOHC"
	specs["Wheelbase"] = "1305 mm"
	specs["Width"] = "760 mm"
	specs["Brake Type"] = "Front Disc / Rear Drum"
	specs["Front Brake Type"] = "Disc"
	specs["Rear Brake Type"] = "Drum"
	specs["Rim Type"] = "Alloy"
	specs["Rim Size"] = "17 inch"
	specs["Seat Height"] = "785 mm"
	specs["Battery Capacity"] = "12V"
	specs["Fuel System"] = "Carburetor"
	specs["Emission Standard"] = "BS6"
	specs["Compression Ratio"] = "9.5:1"
	specs["Instrument Cluster"] = "Analog-Digital"
	specs["DRL (Daytime Running Light)"] = "No"
	specs["Spark Plug Type"] = "Standard"
	specs["Electrical System"] = "12V DC"
	specs["Bore"] = "56.5 mm"
	specs["Stroke"] = "49.8 mm"
	specs["Lubrication System"] = "Air"
	specs["Oil Pump Type"] = "Piston type"
	specs["Front Suspension Travel"] = "240/1260 mm"
	specs["Rear Suspension Travel"] = "240/1260 mm"
	specs["Oil Capacity"] = "3.5 L / 2.7 bar"
	specs["Frame Type"] = "Pressed steel"
	specs["Reserve Fuel Capacity"] = "8 L"
	specs["Side Stand Engine Cutoff"] = "Yes"
	specs["Front Tyre Size"] = "80/100-17"
	specs["Rear Tyre Size"] = "100/90-17"
	specs["Color Options"] = "Standard (Black)"
	specs["Speedometer Type"] = "Analogue"
	specs["Battery Type"] = "7 Ah"
	specs["Brake Diameter"] = "130 mm front, 110 mm rear"
	specs["Suspension Material"] = "Coil spring with hydraulic dampers"
	specs["Wheel Type"] = "Spoke"
	specs["Cooling Type"] = "Forced air (Natural air flow)"

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
