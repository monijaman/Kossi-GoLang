package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter seeds specifications/options for product 'marcel-mcg-c0t-ddge-xx-inverter'
type SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter struct {
	BaseSeeder
}

// NewSpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter creates a new seeder instance
func NewSpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter() *SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter {
	return &SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter{
		BaseSeeder: BaseSeeder{name: "Specifications for marcel-mcg-c0t-ddge-xx-inverter"},
	}
}

func (s *SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Marcel":                          "মার্সেল",
		"marcel-mcg-c0t-ddge-xx-inverter": "মার্সেল-mcg-c0t-ddge-xx-inverter",
		"MCG-C0T-DDGE-XX-INVERTER":        "MCG-C0T-DDGE-XX-INVERTER",
		"Direct Cool":                     "ডাইরেক্ট কুল",
		"300 Ltr.":                        "৩০০ লিটার",
		"220-240V/ 50Hz":                  "২২০-২৪০ভি/ ৫০হার্টজ",
		"50Hz":                            "৫০হার্টজ",
		"RSCR, CSIR":                      "আরএসসিআর, সিএসআইআর",
		"Mechanical":                      "মেকানিক্যাল",
		"Manual":                          "ম্যানুয়াল",
		"R600a":                           "আর৬০০এ",
		"Wire":                            "ওয়্যার",
		"1":                               "১",
		"No":                              "না",
		"1210 x 680 x 860 mm":             "১২১০ x ৬৮০ x ৮৬০ মিমি",
		"53 ± 2 Kg":                       "৫৩ ± ২ কেজি",
		"145V-260V":                       "১৪৫ভি-২৬০ভি",
		"26/54/81 (20ft/40ft/40HQ)":       "২৬/৫৪/৮১ (২০এফটি/৪০এফটি/৪০এইচকিউ)",
		"Embossed Aluminium Sheet":        "এমবসড অ্যালুমিনিয়াম শীট",
		"Replacement Guarantee: 1 Year, Main Parts (Compressor): 12 Years, Door: 3 Years, Spare Parts: 4 Years, After Sales Service: 5 Years": "রিপ্লেসমেন্ট গ্যারান্টি: ১ বছর, মূল অংশ (কম্প্রেসার): ১২ বছর, দরজা: ৩ বছর, স্পেয়ার পার্টস: ৪ বছর, আফটার সেলস সার্ভিস: ৫ বছর",
		"12": "১২",
		"Lock: No, Interior Lamp: No, Handle: Yes, Condenser: Steel, Capillary: Copper, Polyurethane foam blowing agent Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology], Exterior Material: Painted Steel (PCM), Interior Material: Embossed Aluminium Sheet, Shelf: Wire/1, Basket: Wire/1, Sliding Glass: Yes, Loading quantity: 26/54/81 (20ft/40ft/40HQ), Climatic Type: N~ST, Wide voltage range: 145V-260V, Cooling Effect: Freezer Cabinet Less than -18°C, Recommended voltage stabilizer capacity: 2000VA or More, Warranty Note: This warranty does not cover the following cases: 1. Any damage due to accident, electrical fault, natural causes, negligence or improper installation. 2. Any damage or failure caused by unauthorized modification or alteration. 3. Products with original serial numbers that have been removed, distorted or cannot be readily recognized.": "লক: না, ইন্টেরিয়র ল্যাম্প: না, হ্যান্ডেল: হ্যাঁ, কনডেনসার: স্টিল, ক্যাপিলারি: কপার, পলিউরেথেন ফোম ব্লোয়িং এজেন্ট সাইক্লোপেন্টেন [ইকো-ফ্রেন্ডলি (১০০% সিএফসি এবং এইচসিএফসি ফ্রি) গ্রিন টেকনোলজি], এক্সটেরিয়র ম্যাটেরিয়াল: পেইন্টেড স্টিল (পিসিএম), ইন্টেরিয়র ম্যাটেরিয়াল: এমবসড অ্যালুমিনিয়াম শীট, শেল্ফ: ওয়্যার/১, বাস্কেট: ওয়্যার/১, স্লাইডিং গ্লাস: হ্যাঁ, লোডিং কোয়ান্টিটি: ২৬/৫৪/৮১ (২০এফটি/৪০এফটি/৪০এইচকিউ), ক্লাইমেটিক টাইপ: এন~এসটি, ওয়াইড ভোল্টেজ রেঞ্জ: ১৪৫ভি-২৬০ভি, কুলিং ইফেক্ট: ফ্রিজার ক্যাবিনেট লেস থ্যান -১৮°C, রেকমেন্ডেড ভোল্টেজ স্ট্যাবিলাইজার ক্যাপাসিটি: ২০০০ভিএ বা মোর, ওয়ারেন্টি নোট: এই ওয়ারেন্টি নিম্নলিখিত ক্ষেত্রে কভার করে না: ১. দুর্ঘটনা, বৈদ্যুতিক ত্রুটি, প্রাকৃতিক কারণ, অবহেলা বা অনুপযুক্ত ইনস্টলেশনের কারণে কোনো ক্ষতি। ২. অননুমোদিত পরিবর্তন বা পরিবর্ধনের কারণে কোনো ক্ষতি বা ব্যর্থতা। ৩. মূল সিরিয়াল নম্বর যেগুলি সরানো, বিকৃত বা সহজেই চেনা যায় না এমন পণ্য।",
		"Foaming Door": "Foaming Door",
		"220~240 / 50": "২২০~২৪০ / ৫০",
		"Steel": "স্টীল",
		"Yes": "হ্যাঁ",
		"53 ± 2 (Kg)": "৫৩ ± ২ ( কেজি)",
		"Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]": "Cyclopentene [Eco-friendly (১০০% CFC & HCFC Free) Green Technology]",
		"1280": "১২৮০",
		"Copper": "কপার",
		"V 23.01- 154 V 24.01- 235 V 25.01- 147 V 26.01- 147 V 27.01- 118": "V ২৩.০১- ১৫৪ V ২৪.০১- ২৩৫ V ২৫.০১- ১৪৭ V ২৬.০১- ১৪৭ V ২৭.০১- ১১৮",
		"RoHS Certified": "RoHS Certified",
		"V 23.01- R134a V 24.01- R134a V 25.01- R600a V 26.01- R600a V 27.01- R600a": "V ২৩.০১- R১৩৪a V ২৪.০১- R১৩৪a V ২৫.০১- R৬০০a V ২৬.০১- R৬০০a V ২৭.০১- R৬০০a",
		"N~ST": "N~ST",
		"Painted Steel (PCM)": "Painted Steel (PCM)",
		"V 21.01 V 22.01 V 23.01 V 24.01 V 25.01": "V ২১.০১ V ২২.০১ V ২৩.০১ V ২৪.০১ V ২৫.০১",
		"60 ± 2 (Kg)": "৬০ ± ২ ( কেজি)",
		"Wire/1": "Wire/১",
		"V 23.01- RSCR V 24.01- CSIR V 25.01- RSCR V 26.01- RSCR V 27.01- RSCR": "V ২৩.০১- RSCR V ২৪.০১- CSIR V ২৫.০১- RSCR V ২৬.০১- RSCR V ২৭.০১- RSCR",
		"725": "৭২৫",
		"890": "৮৯০",
		"V 26.01, 150V-260V V 27.01, 145V-260V": "V ২৬.০১, ১৫০V-২৬০V V ২৭.০১, ১৪৫V-২৬০V",
		"Freezer Cabinet Less than -18℃": "Freezer Cabinet Less than -১৮℃",
		"26/54/81": "২৬/৫৪/৮১",

	}
}

// Seed inserts specification records for the product identified by slug 'marcel-mcg-c0t-ddge-xx-inverter'
func (s *SpecificationSeederRefrigeratorMarcelMcgC0tDdgeXxInverter) Seed(db *gorm.DB) error {
	productSlug := "marcel-mcg-c0t-ddge-xx-inverter"
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
		"Basket": "Wire/1",
		"Capillary": "Copper",
		"Caster": "No",
		"Climatic Type (SN, N, ST, T)": "N~ST",
		"Compressor Input Power (Watt)": "V 23.01- 154V 24.01- 235V 25.01- 147V 26.01- 147V 27.01- 118",
		"Compressor Type": "V 23.01- RSCRV 24.01- CSIRV 25.01- RSCRV 26.01- RSCRV 27.01- RSCR",
		"Condenser": "Steel",
		"Cooling Efect": "Freezer Cabinet Less than -18℃",
		"Defrosting (Automatic/ Manual)": "Manual",
		"Depth/mm": "680",
		"Drawer": "No",
		"Exterior Material": "Painted Steel (PCM)",
		"Gross Volume": "300 Ltr.",
		"Gross Weight": "60 ± 2 (Kg)",
		"Handle (Recessed/ Grip)": "Yes",
		"Height/mm": "860",
		"Ice/ cold water Dispenser": "No",
		"Interior Lamp": "No",
		"Interior Material": "Embossed Aluminium Sheet",
		"Loading quantity (20ft/40ft.40HQ)": "26/54/81",
		"Lock": "No",
		"Net Volume": "300 Ltr.",
		"Net Weight": "53 ± 2 (Kg)",
		"Polyurethane foam blowing agent": "Cyclopentene [Eco-friendly (100% CFC & HCFC Free) Green Technology]",
		"Rated Operating Voltage and Frequency": "220~240 / 50",
		"Recommended voltage stabilizer capacity (2000VA or More)": "V 21.01V 22.01V 23.01V 24.01V 25.01",
		"Refrigerant": "V 23.01- R134aV 24.01- R134aV 25.01- R600aV 26.01- R600aV 27.01- R600a",
		"Reversible Door": "Foaming Door",
		"Shelf (Material/ No)": "Wire/1",
		"Sliding Glass": "No",
		"Temperature Control (Electronic/  Mechanical)": "Mechanical",
		"Thermostat": "RoHS Certified",
		"Type": "Direct Cool",
		"Wide voltage range": "V 26.01, 150V-260VV 27.01, 145V-260V",
		"Width/mm": "1210",
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
