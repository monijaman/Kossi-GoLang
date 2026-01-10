package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMfsT9cC2srVb seeds specifications/options for product 'marcel-mfs-t9c-c2sr-vb'
type SpecificationSeederRefrigeratorMarcelMfsT9cC2srVb struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMfsT9cC2srVb creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMfsT9cC2srVb() *SpecificationSeederRefrigeratorMarcelMfsT9cC2srVb {
	return &SpecificationSeederRefrigeratorMarcelMfsT9cC2srVb{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mfs-t9c-c2sr-vb"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMfsT9cC2srVb) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                 "মার্সেল",
		"marcel-mfs-t9c-c2sr-vb": "মার্সেল-mfs-t9c-c2sr-vb",
		"MFS-T9C-C2SR-VB":        "MFS-T9C-C2SR-VB",

		// specs values -> Bangla
		"Direct Cool":      "ডাইরেক্ট কুল",
		"93 Ltr.":          "৯৩ লিটার",
		"90 Ltr.":          "৯০ লিটার",
		"23 ± 2 Kg":        "২৩ ± ২ কেজি",
		"26 ± 2 Kg":        "২৬ ± ২ কেজি",
		"SN~T":             "SN~T",
		"220~240 and 50Hz": "২২০~২৪০ এবং ৫০Hz",
		"D43WY1 / RSIR":    "D43WY1 / RSIR",
		"Mechanical":       "যান্ত্রিক",
		"Manual":           "ম্যানুয়াল",
		"Recessed":         "রিসেসড",
		"No":               "না",
		"R600a":            "R600a",
		"RoHS Certified":   "RoHS সার্টিফাইড",
		"Copper":           "কপার",
		"CycloPentene":     "সাইক্লোপেন্টেন",
		"Tempered Glass/2": "টেম্পার্ড গ্লাস/২",
		"Yes/3":            "হ্যাঁ/৩",
		"Yes":              "হ্যাঁ",
		"12 cubes":         "১২ পিস",
		"478 mm":           "৪৭৮ মিমি",
		"446 mm":           "৪৪৬ মিমি",
		"847 mm":           "৮৪৭ মিমি",
		"500 mm":           "৫০০ মিমি",
		"460 mm":           "৪৬০ মিমি",
		"880 mm":           "৮৮০ মিমি",
		"340/ 240/ 9":      "৩৪০/ ২৪০/ ৯",
	
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Recessed/ Grip": "Recessed/ Grip",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"V 0301- R600a": "V ০৩০১- R৬০০a",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mfs-t9c-c2sr-vb'
func (s *SpecificationSeederRefrigeratorMarcelMfsT9cC2srVb) Seed(db *gorm.DB) error {
	productSlug := "marcel-mfs-t9c-c2sr-vb"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}
		return err
	}
	productID := prod.ID

	existingkeyMapping := map[string]uint{
		"Brand":                             310,
		"Model Name":                        316,
		"Door Type":                         142,
		"Capacity":                          138,
		"Refrigerator Capacity":             156,
		"Freezer Capacity":                  146,
		"Energy Efficiency Rating":          143,
		"Energy Star Rating":                144,
		"Annual Energy Consumption":         137,
		"Dimensions":                        25,
		"Weight":                            80,
		"Color":                             311,
		"Compressor Type":                   139,
		"Cooling Technology":                698,
		"Defrost Type":                      141,
		"Temperature Control":               158,
		"Shelf Material":                    699,
		"Number of Shelves":                 154,
		"Door Bins":                         700,
		"Crisper Drawers":                   701,
		"Ice Maker":                         702,
		"Water Dispenser":                   703,
		"Noise Level":                       150,
		"Voltage":                           160,
		"Frequency (Hz)":                    704,
		"App Control":                       705,
		"Voice Assistant Support":           385,
		"Warranty":                          323,
		"Compressor Warranty (Years)":       707,
		"Refrigerant":                       708,
		"Gross Volume":                      709,
		"Net Volume":                        710,
		"Special Features":                  69,
		"Capillary":                         711,
		"Climate Type (SN, N, ST, T)":       712,
		"Compressor Input Power (Watt)":     713,
		"Cooling Effect":                    714,
		"Defrosting (Automatic/ Manual)":    715,
		"Deodorizer":                        716,
		"Depth (mm)":                        717,
		"Door Basket":                       718,
		"Egg Tray or Pocket":                719,
		"Handle (Recessed/ Grip)":           720,
		"Interior Lamp":                     721,
		"Polyurethane Foam Blowing Agent":   722,
		"Reversible Door":                   723,
		"Thermostat":                        724,
		"Vegetable Crisper":                 725,
		"Vegetable Crisper Cover":           726,
		"Loading Capacity (40HQ/40Ft/20Ft)": 727,
		"Type":                              456,
		"Lock":                              361,
		"Height (mm)":                       587,
		"Width":                             136,
		"Lock Type":                         299,
	}
	specs := map[string]string{
		"Brand":               "Marcel",
		"Model Name":          "MFS-T9C-C2SR-VB",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Type (SN, N, ST, T)": "SN~T",
		"Compressor Input Power (Watt)": "V.0301- 57",
		"Compressor Model & Type": "D43WY1 & RSIR",
		"Cooling": "Faster Metal Cooling",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth (mm)": "446 mm",
		"Door Bin": "Yes/3",
		"Egg Tray": "No",
		"Gross Volume": "93 Ltr.",
		"Gross Weight": "26 ± 2 Kg",
		"Handle (Recessed/ Grip)": "Recessed",
		"Height (mm)": "847 mm",
		"Ice Making Compartment Cover": "Flip Cover",
		"Ice Tray": "Yes ,12 cubes",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "340/ 240/ 96",
		"Lock": "No",
		"Net Volume": "90 Ltr.",
		"Net Weight": "23 ± 2 Kg",
		"Polyurethane foam blowing agent": "CycloPentene[Eco-friendly (HCFC Free) Green Technology]",
		"Rated Voltage/ Hz/ watt": "220~240 and 50Hz",
		"Recommended voltage stabilizer capacity": "No need for Version 0301, Range Cover:140 -260V~",
		"Refrigerant": "V 0301- R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Tempered Glass/2",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Vegetable Crisper": "Yes/1",
		"Vegetable Crisper Cover": "Yes/1",
		"Width": "478 mm",
	}

	banglaTranslations := s.getBanglaTranslations()
	for key, value := range specs {
		specKeyID, exists := existingkeyMapping[key]
		if !exists {
			log.Printf("⚠️  Key not found in existingkeyMapping: '%s' (used in product: %s)", key, productSlug)
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
				// Ensure the ID is set after creation
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
