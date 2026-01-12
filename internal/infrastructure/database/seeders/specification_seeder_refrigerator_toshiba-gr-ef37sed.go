package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorToshibaGrEf37sed seeds specifications/options for product 'toshiba-gr-ef37sed'
type SpecificationSeederRefrigeratorToshibaGrEf37sed struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorToshibaGrEf37sed creates a new seeder instance
func NewSpecificationSeederRefrigeratorToshibaGrEf37sed() *SpecificationSeederRefrigeratorToshibaGrEf37sed {
	return &SpecificationSeederRefrigeratorToshibaGrEf37sed{
		BaseSeeder: BaseSeeder{name: "Specifications for toshiba-gr-ef37sed"},
	}
}

func (s *SpecificationSeederRefrigeratorToshibaGrEf37sed) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Toshiba": "তোশিবা",
		"GR-EF37SED": "GR-EF37SED",
		"Japan": "জাপান",
		"Direct Cool": "ডিরেক্ট কুল",
		"No Frost": "নো ফ্রস্ট",
		"Yes": "হ্যাঁ",
		"No": "না",
		"Manual": "ম্যানুয়াল",
		"Automatic": "স্বয়ংক্রিয়",
		"Mechanical": "মেকানিক্যাল",
		"Digital": "ডিজিটাল",
		"R600a": "R600a",
		"220-240V~ and 50Hz": "২২০-২৪০V~ এবং ৫০Hz",
		"Inverter": "ইনভার্টার",
		"Conventional": "প্রচলিত",
		"Top Mount": "টপ মাউন্ট",
		"Bottom Mount": "বটম মাউন্ট",
		"Single Door": "সিঙ্গেল ডোর",
		"Double Door": "ডাবল ডোর",
		"Side by Side": "সাইড বাই সাইড",
		"Cross Door": "ক্রস ডোর",
		"French Door": "ফ্রেঞ্চ ডোর",
		"Multi-Door": "মাল্টি-ডোর",
		"Tempered Glass": "টেম্পার্ড গ্লাস",
		"Wire": "ওয়্যার",
		"Glass": "গ্লাস",
		"LED": "LED",
		"1 Year": "১ বছর",
		"2 Years": "২ বছর",
		"5 Years": "৫ বছর",
		"10 Years": "১০ বছর",
		"Copper": "তামা",
		"Smart": "স্মার্ট",
		"WiFi": "ওয়াইফাই",
		"N/A": "N/A",
		"N~ST": "N~ST",
		"RoHS Certified": "RoHS সার্টিফাইড",
		"Cyclopentane": "সাইক্লোপেন্টেন",
		"Recessed/ Grip": "রিসেসড/গ্রিপ",
		"GPPS/3": "GPPS/৩",
		"GPPS/4": "GPPS/৪",
		"HIPS/2": "HIPS/২",
		"HIPS/3": "HIPS/৩",
		"Wire/2": "Wire/২",
		"Wire/3": "Wire/৩",
		"Wire/4": "Wire/৪",
		"Tempered Glass/3": "টেম্পার্ড গ্লাস/৩",
		"Tempered Glass/4": "টেম্পার্ড গ্লাস/৪",
		"2": "২",
		"3": "৩",
		"4": "৪",
		"50": "৫০",
		"Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C": "ফ্রিজার ক্যাবিনেট -১৮°C এর কম, রেফ্রিজারেটর ক্যাবিনেট ০°C থেকে +৫°C",
		"Compressor: 10 Years, Spare Parts: 2 Years, After Sales Service: 5 Years": "কম্প্রেসর: ১০ বছর, খুচরা যন্ত্রাংশ: ২ বছর, বিক্রয়োত্তর সেবা: ৫ বছর",
	}
}

// Seed inserts specification records for the product identified by slug 'toshiba-gr-ef37sed'
func (s *SpecificationSeederRefrigeratorToshibaGrEf37sed) Seed(db *gorm.DB) error {
	productSlug := "toshiba-gr-ef37sed"
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
		"Origin":                            318,
		"Gross Weight":                      728,
		"Drawer":                            729,
		"Energy Rating":                     730,
		"Shelf (Material/No.)":              731,

	}

	specs := map[string]string{
		"Brand": "Toshiba",
		"Model Name": "GR-EF37SED",
		"Origin": "Japan",
		"Type": "No Frost",
		"Door Type": "Top Mount",
		"Compressor Type": "Inverter",
		"Cooling Technology": "No Frost",
		"Defrost Type": "Automatic",
		"Defrosting (Automatic/ Manual)": "Automatic",
		"Temperature Control": "Digital",
		"Shelf Material": "Tempered Glass",
		"Shelf (Material/No.)": "Tempered Glass/3",
		"Number of Shelves": "3",
		"Door Basket": "Yes",
		"Vegetable Crisper": "Yes",
		"Vegetable Crisper Cover": "Yes",
		"Egg Tray or Pocket": "Yes",
		"Interior Lamp": "Yes",
		"Ice Maker": "Yes",
		"Handle (Recessed/ Grip)": "Recessed/ Grip",
		"Lock": "No",
		"Reversible Door": "Yes",
		"Capillary": "Copper",
		"Polyurethane Foam Blowing Agent": "Cyclopentane",
		"Climate Type (SN, N, ST, T)": "N~ST",
		"Thermostat": "RoHS Certified",
		"Cooling Effect": "Freezer Cabinet Less than -18°C, Refrigerator Cabinet 0°C to +5°C",
		"Warranty": "2 Years",
		"Compressor Warranty (Years)": "10 Years",
		"Refrigerant": "R600a",
		"Voltage": "220-240V~ and 50Hz",
		"Frequency (Hz)": "50",
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
