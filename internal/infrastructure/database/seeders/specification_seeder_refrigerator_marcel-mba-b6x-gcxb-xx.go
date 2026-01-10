package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx seeds specifications/options for product 'marcel-mba-b6x-gcxb-xx'
type SpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx() *SpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx {
	return &SpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mba-b6x-gcxb-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                         "মার্সেল",
		"marcel-mba-b6x-gcxb-xx":         "মার্সেল-এমবিএ-বি৬এক্স-জিসিএক্সবি-এক্সএক্স",
		"MBA-B6X-GCXB-XX":                "এমবিএ-বি৬এক্স-জিসিএক্সবি-এক্সএক্স",
		"Glass Door":                     "গ্লাস দরজা",
		"168 Liters":                     "১৬৮ লিটার",
		"0 Liters":                       "০ লিটার",
		"3 Star":                         "৩ তারা",
		"3":                              "৩",
		"180 kWh":                        "১৮০ কিলোওয়াট ঘণ্টা",
		"480 x 520 x 850 mm (W x D x H)": "৪৮০ x ৫২০ x ৮৫০ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"32 kg":                          "৩২ কেজি",
		"Black Glass":                    "ব্ল্যাক গ্লাস",
		"Reciprocating":                  "রেসিপ্রোকেটিং",
		"Direct Cool":                    "ডাইরেক্ট কুল",
		"Manual":                         "ম্যানুয়াল",
		"Mechanical":                     "মেকানিক্যাল",
		"Wire Shelves":                   "ওয়্যার শেল্ফ",
		"2":                              "২",
		"No":                             "না",
		"38 dB":                          "৩৮ ডেসিবেল",
		"220V":                           "২২০ ভোল্ট",
		"50":                             "৫০",
		"1 Year Comprehensive + 5 Years on Compressor": "১ বছর কমপ্রিহেনসিভ + কম্প্রেসারে ৫ বছর",
		"5":          "৫",
		"R134a":      "আর১৩৪এ",
		"180 Liters": "১৮০ লিটার",
		"Glass door design, Interior LED light, Adjustable shelves, Lock & key": "গ্লাস দরজা ডিজাইন, ইন্টেরিয়র LED লাইট, অ্যাডজাস্টেবল শেল্ফ, লক এবং কী",
		// Add more translations as needed
		"Continuously maintained from 0.0°C and 7.2°C with a cabinet average below 3.2°C.": "Continuously maintained from ০.০°C and ৭.২°C with a cabinet average below ৩.২°C.",
		"1873":               "১৮৭৩",
		"IPX3":               "IPX৩",
		"254 Ltr":            "২৫৪ Ltr",
		"Steel":              "স্টীল",
		"Yes":                "হ্যাঁ",
		"57/ 57/ 27":         "৫৭/ ৫৭/ ২৭",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"Recessed/ Grip":     "Recessed/ Grip",
		"88":                 "৮৮",
		"Optional":           "Optional",
		"Electronic":         "ইলেকট্রনিক",
		"RSCR":               "RSCR",
		"R600a":              "R৬০০a",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"260 Ltr.":                            "২৬০ লিটার",
		"Copper":                              "কপার",
		"660":                                 "৬৬০",
		"Beverage Cooler":                     "Beverage Cooler",
		"Automatic":                           "স্বয়ংক্রিয়",
		"Steel tube":                          "Steel tube",
		"N ~ T":                               "N ~ T",
		"78.5 / 85.15 ± 2 Kg":                 "৭৮.৫ / ৮৫.১৫ ± ২ কেজি",
		"620":                                 "৬২০",
		"Recommended for Commercial use only": "Recommended for Commercial use only",
	
		"CycloPentane[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "CycloPentane[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]": "Cyclopentene[পরিবেশ বান্ধব (১০০% CFC &HCFC মুক্ত) সবুজ প্রযুক্তি]",
		"Freezer Cabinet Less than -180CRefrigerator Cabinet 00Cto +50C": "ফ্রিজার ক্যাবিনেট -১৮০C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০০C থেকে +৫০C",
		"Freezer Cabinet Less than -18℃Refrigerator Cabinet 0℃ to +5℃": "ফ্রিজার ক্যাবিনেট -১৮℃ এর কম, রেফ্রিজারেটর ক্যাবিনেট ০℃ থেকে +৫℃",
		"N ~ ST": "N ~ ST",
		"N~ST": "N~ST",
		"PVC/1": "PVC/১",
		"PVC/2": "PVC/২",
		"PVC/3": "PVC/৩",
		"RSIR": "RSIR",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Wire/1": "Wire/১",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Yes (ABS/ PS)": "হ্যাঁ (ABS/ PS)",
		"Yes (Plastic)": "হ্যাঁ (প্লাস্টিক)",
		"Yes/1": "হ্যাঁ/১",
		"Yes/2": "হ্যাঁ/২",
	
		"580": "৫৮০",}
}

// Seed inserts specification records for the product identified by slug 'marcel-mba-b6x-gcxb-xx'
func (s *SpecificationSeederRefrigeratorMarcelMbaB6xGcxbXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mba-b6x-gcxb-xx"
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
		"Application Type": "Recommended for Commercial use only",
		"CB/Safety Certificate (IEC 60335-1 & 60335-2-89)": "Yes",
		"Can Storage Dispenser": "No",
		"Capillary": "Copper",
		"Climate Class": "N ~ T",
		"Compressor Input Power (Watt)": "88",
		"Compressor Type": "RSCR",
		"Condenser": "Steel tube",
		"Cooling Efect": "Continuously maintained from 0.0°C and 7.2°C with a cabinet average below 3.2°C.",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Deodorzier": "Optional",
		"Depth/mm": "580",
		"Door Basket": "No",
		"Egg Tray or Pocket": "No",
		"Gross Volume": "260 Ltr.",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Height/mm": "1837",
		"Interior Lamp": "Yes",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "57/ 57/ 27",
		"Lock": "Yes",
		"Net Volume": "254 Ltr",
		"Polyurethane foam blowing agent": "Cyclopentene[Eco-friendly (100% CFC &HCFC Free) Green Technology]",
		"Product IP Rating": "IPX3",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Refrigerant": "R600a",
		"Reversible Door": "No",
		"Shelf (Material/ No.)": "Steel",
		"Temperature Control": "Electronic",
		"Type": "Beverage Cooler",
		"Vegetable Crisper": "No",
		"Vegetable Crisper Cover": "No",
		"Weight/Kg - Net/Packing:": "78.5 / 85.15 ± 2 Kg",
		"Width/mm": "549",
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
