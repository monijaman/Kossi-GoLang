package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx seeds specifications/options for product 'marcel-mbb-b6x-tdxx-xx'
type SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx() *SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mbb-b6x-tdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mbb-b6x-tdxx-xx":          "মার্সেল-এমবিবি-বি৬এক্স-টিডিএক্সএক্স-এক্সএক্স",
		"MBB-B6X-TDXX-XX":                 "এমবিবি-বি৬এক্স-টিডিএক্সএক্স-এক্সএক্স",
		"Standard Door":                   "স্ট্যান্ডার্ড দরজা",
		"254 Liters":                      "২৫৪ লিটার",
		"0 Liters":                        "০ লিটার",
		"Not Rated":                       "রেটেড নয়",
		"Not Specified":                   "নির্দিষ্ট নয়",
		"620 x 660 x 1873 mm (W x D x H)": "৬২০ x ৬৬০ x ১৮৭৩ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"78.5 kg":                         "৭৮.৫ কেজি",
		"CSIR":                            "সিএসআইআর",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Manual":                          "ম্যানুয়াল",
		"Mechanical":                      "মেকানিক্যাল",
		"Steel":                           "স্টিল",
		"No":                              "না",
		"220-240V":                        "২২০-২৪০ ভোল্ট",
		"50":                              "৫০",
		"Commercial Use: 4 Years on Compressor, 1 Year on Door, 2 Years on Spare Parts, 2 Years After Sales Service": "কমার্শিয়াল ব্যবহার: কম্প্রেসারে ৪ বছর, দরজায় ১ বছর, স্পেয়ার পার্টসে ২ বছর, আফটার সেলস সার্ভিসে ২ বছর",
		"4":          "৪",
		"R134a":      "আর১৩৪এ",
		"260 Liters": "২৬০ লিটার",
		"Lock, Reversible Door No, Handle Recessed/Grip, Skin Condenser 100% Copper, RoHS Certified Thermostat, Copper Capillary, Eco-friendly PU Foam, IP24 Rating, Recommended for Commercial Use": "লক, রিভার্সিবল দরজা না, হ্যান্ডেল রিসেসড/গ্রিপ, স্কিন কন্ডেন্সার ১০০% কপার, রোএইচএস সার্টিফাইড থার্মোস্ট্যাট, কপার ক্যাপিলারি, ইকো-ফ্রেন্ডলি পিইউ ফোম, আইপি২৪ রেটিং, কমার্শিয়াল ব্যবহারের জন্য সুপারিশকৃত",
		// Add more translations as needed
		"Continuously maintained from 0.0°C and 7.2°C with a cabinet average below 3.2°C.": "Continuously maintained from ০.০°C and ৭.২°C with a cabinet average below ৩.২°C.",
		"2200VA": "২২০০VA",
		"1873": "১৮৭৩",
		"254 Ltr": "২৫৪ Ltr",
		"V 0101- 178 V 0201- 178": "V ০১০১- ১৭৮ V ০২০১- ১৭৮",
		"Yes": "হ্যাঁ",
		"57/ 57/ 27": "৫৭/ ৫৭/ ২৭",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"Recessed/ Grip": "Recessed/ Grip",
		"Optional": "Optional",
		"V 0101- R134a V 0201- R134a": "V ০১০১- R১৩৪a V ০২০১- R১৩৪a",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"260 Ltr.": "২৬০ লিটার",
		"V 0101- Yes V 0201- No": "V ০১০১- Yes V ০২০১- No",
		"Copper": "কপার",
		"RoHS Certified": "RoHS Certified",
		"660": "৬৬০",
		"IP24": "IP২৪",
		"V 0101- CSIR V 0201- CSIR": "V ০১০১- CSIR V ০২০১- CSIR",
		"Skin condenser 100% Cu + External condenser Steel": "Skin condenser ১০০% Cu + External condenser Steel",
		"N ~ ST": "N ~ ST",
		"Beverage Cooler": "Beverage Cooler",
		"Automatic": "স্বয়ংক্রিয়",
		"78.5 / 85.15 ± 2 Kg": "৭৮.৫ / ৮৫.১৫ ± ২ কেজি",
		"620": "৬২০",
		"Recommended for Commercial use only": "Recommended for Commercial use only",

	
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"R600a": "R600a",
		"RSCR": "RSCR",
		"RSIR": "RSIR",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"580": "৫৮০",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mbb-b6x-tdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMbbB6xTdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mbb-b6x-tdxx-xx"
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
		"Model Name":          "MBB-B6X-TDXX-XX",
		"Application Type": "Recommended for Commercial use only",
		"CB/Safety Certificate (IEC 60335-1 & 60335-2-89)": "V 0101- YesV 0201- No",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Class": "N ~ ST",
		"Compressor Input Power (Watt)": "V 0101- 178V 0201- 178",
		"Compressor Type": "V 0101- CSIRV 0201- CSIR",
		"Condenser": "Skin condenser 100% Cu + External condenser Steel",
		"Cooling Efect": "Continuously maintained from 0.0°C and 7.2°C with a cabinet average below 3.2°C.",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Deodorzier": "Optional",
		"Depth (mm)": "580",
		"Door Basket": "No",
		"Egg Tray or Pocket": "No",
		"Gross Volume": "260 Ltr.",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height (mm)": "1837",
		"Interior Lamp": "Yes",
		"Loading Capacity (40HQ/40Ft/20Ft)": "57/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "254 Ltr",
		"Polyurethane Foam Blowing Agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Product IP Rating": "IP24",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Recommended voltage stabilizer capacity": "2200VA",
		"Refrigerant": "V 0101- R134aV 0201- R134a",
		"Reversible Door": "No",
		"Shelf Material": "Steel",
		"Temperature Control": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Beverage Cooler",
		"Vegetable Crisper": "No",
		"Vegetable Crisper Cover": "No",
		"Weight/Kg - Net/Packing:": "78.5 / 85.15 ± 2 Kg",
		"Width": "549",
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
