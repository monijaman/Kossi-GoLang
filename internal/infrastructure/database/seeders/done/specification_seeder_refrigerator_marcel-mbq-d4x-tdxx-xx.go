package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx seeds specifications/options for product 'marcel-mbq-d4x-tdxx-xx'
type SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx() *SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx {
	return &SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mbq-d4x-tdxx-xx"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mbq-d4x-tdxx-xx":          "মার্সেল-mbq-d4x-tdxx-xx",
		"MBQ-D4X-TDXX-XX":                 "MBQ-D4X-TDXX-XX",
		"Glass Door":                      "গ্লাস ডোর",
		"396 Liters":                      "৩৯৬ লিটার",
		"0 Liters":                        "০ লিটার",
		"440 Liters":                      "৪৪০ লিটার",
		"643 x 738 x 1834 mm (W x D x H)": "৬৪৩ x ৭৩৮ x ১৮৩৪ মিমি (প্রস্থ x গভীরতা x উচ্চতা)",
		"117.25 kg":                       "১১৭.২৫ কেজি",
		"Black":                           "কালো",
		"V 0101-CSR":                      "V 0101-CSR",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"Automatic":                       "অটোমেটিক",
		"Mechanical":                      "মেকানিক্যাল",
		"Steel":                           "স্টিল",
		"1":                               "১",
		"No":                              "না",
		"42 dB":                           "৪২ ডেসিবেল",
		"220-240V":                        "২২০-২৪০ ভোল্ট",
		"50":                              "৫০",
		"2 Years":                         "২ বছর",
		"5":                               "৫",
		"R134a":                           "R134a",
		"Beverage Cooler, Commercial Use, External Condenser, Lock, Interior Lamp": "বেভারেজ কুলার, কমার্শিয়াল ব্যবহার, এক্সটার্নাল কন্ডেনসার, লক, ইন্টেরিয়র ল্যাম্প",
		// Add more translations as needed
		"12W": "১২W",
		"Continuously maintained from 0.0°C and 7.2°C with a cabinet average below 3.2°C.": "Continuously maintained from ০.০°C and ৭.২°C with a cabinet average below ৩.২°C.",
		"1000VA (If built in), Otherwise: 2100VA": "১০০০VA (If built in), Otherwise: ২১০০VA",
		"54/ 54/ 24": "৫৪/ ৫৪/ ২৪",
		"Yes": "হ্যাঁ",
		"1885": "১৮৮৫",
		"220-240V~ and 50Hz": "২২০-২৪০V~ and ৫০Hz",
		"Recessed/ Grip": "Recessed/ Grip",
		"117.25 /125.8 ± 0.5Kg": "১১৭.২৫ /১২৫.৮ ± ০.৫ কেজি",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"V 0101-260": "V ০১০১-২৬০",
		"741": "৭৪১",
		"Copper": "কপার",
		"RoHS Certified": "RoHS Certified",
		"1.40A": "১.৪০A",
		"IP24": "IP২৪",
		"665": "৬৬৫",
		"Beverage Cooler": "Beverage Cooler",
		"440 Ltr.": "৪৪০ লিটার",
		"V 0101- R134a": "V ০১০১- R১৩৪a",
		"External condenser copper coated Steel": "External condenser copper coated Steel",
		"Recommended for Commercial use only": "Recommended for Commercial use only",
		"N ~ T and 5 or 7 (IEC 60335-2-89)": "N ~ T and ৫ or ৭ (IEC ৬০৩৩৫-২-৮৯)",
		"396 Ltr": "৩৯৬ Ltr",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mbq-d4x-tdxx-xx'
func (s *SpecificationSeederRefrigeratorMarcelMbqD4xTdxxXx) Seed(db *gorm.DB) error {
	productSlug := "marcel-mbq-d4x-tdxx-xx"
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
		"Brand":                       310,
		"Model Name":                  316,
		"Door Type":                   142,
		"Capacity":                    138,
		"Refrigerator Capacity":       156,
		"Freezer Capacity":            146,
		"Energy Efficiency Rating":    143,
		"Energy Star Rating":          144,
		"Annual Energy Consumption":   137,
		"Dimensions":                  25,
		"Weight":                      80,
		"Color":                       311,
		"Compressor Type":             139,
		"Cooling Technology":          698,
		"Defrost Type":                141,
		"Temperature Control":         158,
		"Shelf Material":              699,
		"Number of Shelves":           154,
		"Door Bins":                   700,
		"Crisper Drawers":             701,
		"Ice Maker":                   702,
		"Water Dispenser":             703,
		"Noise Level":                 150,
		"Voltage":                     160,
		"Frequency (Hz)":              704,
		"App Control":                 705,
		"Voice Assistant Support":     385,
		"Warranty":                    323,
		"Compressor Warranty (Years)": 707,
		"Refrigerant":                 708,
		"Gross Volume":                709,
		"Net Volume":                  710,
		"Special Features":            69,
	}

	
    		
    specs := map[string]string{
		"Type": "Beverage Cooler",
		"Gross Volume": "440 Ltr.",
		"Net Volume": "396 Ltr",
		"Climatic Type (SN, N, ST, T)": "N ~ T and 5 or 7 (IEC 60335-2-89)",
		"Rated Operating Voltage and Frequency": "220-240V~ and 50Hz",
		"Rated Current": "1.40A",
		"Rated Power of LED Lamp": "12W",
		"Compressor Input Power (Watt)": "V 0101-260",
		"Compressor Type": "V 0101-CSR",
		"Cooling Effect": "Continuously maintained from 0.0°C and 7.2°C with a cabinet average below 3.2°C.",
		"Temperature Control": "Mechanical",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Reversible Door": "No",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "Yes",
		"Refrigerant": "V 0101- R134a",
		"External Condenser": "External condenser copper coated Steel",
		"Thermostat": "RoHS Certified",
		"Capillary": "Copper",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Recommended voltage stabilizer capacity": "1000VA (If built in), Otherwise: 2100VA",
		"Application Type": "Recommended for Commercial use only",
		"Product IP Rating": "IP24",
		"Shelf (Material/ No.)": "Steel",
		"Door Basket": "No",
		"Interior Lamp": "Yes",
		"Vegetable Crisper": "No",
		"Vegetable Crisper Cover": "No",
		"Egg Tray or Pocket": "No",
		"Can Storage Dispenser": "No",
		"Deodorzier": "No",
		"Width/mm": "665",
		"Depth/mm": "741",
		"Height/mm": "1885",
		"Weight/Kg - Net/Packing:": "117.25 /125.8 ± 0.5Kg",
		"Loading Capacity- 40HQ/ 40Ft/ 20Ft": "54/ 54/ 24",
		"CB/Safety Certificate (IEC 60335-1 & 60335-2-89)": "Yes",
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
