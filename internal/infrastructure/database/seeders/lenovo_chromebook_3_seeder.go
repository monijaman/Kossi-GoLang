package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LenovoChromebook3Seeder seeds specifications for Lenovo Chromebook 3
type LenovoChromebook3Seeder struct {
	BaseSeeder
}

// NewLenovoChromebook3Seeder creates a new Lenovo Chromebook 3 seeder
func NewLenovoChromebook3Seeder() *LenovoChromebook3Seeder {
	return &LenovoChromebook3Seeder{
		BaseSeeder: BaseSeeder{name: "Lenovo Chromebook 3 Specifications"},
	}
}

func (lc3 *LenovoChromebook3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Lenovo":                           "লেনোভো",
		"Chromebook 3":                     "ক্রোমবুক ৩",
		"Chrome OS":                        "ক্রোম ওএস",
		"MediaTek":                         "মিডিয়াটেক",
		"MT8183":                           "এমটি৮১৮৩",
		"MediaTek Platform Controller Hub": "মিডিয়াটেক প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Integrated Graphics":              "ইন্টিগ্রেটেড গ্রাফিক্স",
		"1366 x 768 pixels":                "১৩৬৬ x ৭৬৮ পিক্সেল",
		"11.6 inches":                      "১১.৬ ইঞ্চি",
		"37 Wh":                            "৩৭ ডব্লিউএইচ",
		"Plastic":                          "প্লাস্টিক",
		"Black":                            "কালো",
		"1.15 kg":                          "১.১৫ কেজি",
		"2023":                             "২০২৩",
		"1 Year International Warranty":    "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"32 GB":                            "৩২ জিবি",
		"Bluetooth 4.2":                    "ব্লুটুথ ৪.২",
		"4 GB":                             "৪ জিবি",
		"TN":                               "টিএন",
		"60 Hz":                            "৬০ হার্জ",
		"No":                               "না",
		"3.5mm Combo Jack":                 "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":               "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.2 Gen 1":                    "ইউএসবি ৩.২ জেন ১",
		"Integrated":                       "ইন্টিগ্রেটেড",
		"32 GB eMMC":                       "৩২ জিবি ইএমএমসি",
		"MediaTek MT8183":                  "মিডিয়াটেক এমটি৮১৮৩",
		"Passive Cooling":                  "প্যাসিভ কুলিং",
		"60":                               "৬০",
		"1 Year":                           "১ বছর",
		// New translations for added specs
		"Entry Level Chromebook":               "এন্ট্রি লেভেল ক্রোমবুক",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"45W USB-C":                            "৪৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"TN, 220 nits":                         "টিএন, ২২০ নিটস",
		"2.0 GHz":                              "২.০ গিগাহার্টজ",
		"1 MB cache":                           "১ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Mono Speaker":                         "মনো স্পিকার",
		"286 x 200 x 18 mm":                    "২৮৬ x ২০০ x ১৮ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"4":                                    "৪",
		"LPDDR3":                               "এলপিডিডিআর৩",
		"eMMC":                                 "ইএমএমসি",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 1.4":                             "এইচডিএমআই ১.৪",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 1": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ১",
	}
}

// Seed implements the Seeder interface for Lenovo Chromebook 3
func (lc3 *LenovoChromebook3Seeder) Seed(db *gorm.DB) error {
	productSlug := "lenovo-chromebook-3"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}
		return err
	}
	productID := prod.ID

	// Mapping of specification keys to their IDs in the database
	existingKeyMapping := map[string]uint{
		"Brand":                310,
		"Model Name":           316,
		"Operating System":     49,
		"Processor Brand":      204,
		"Processor Model":      206,
		"Processor Generation": 205,
		"Chipset":              19,
		"Graphics Card":        200,
		"Screen Resolution":    207,
		"Display Size":         683,
		"Battery Capacity":     11,
		"Build Material":       14,
		"Color":                318,
		"Product Weight":       697,
		"Release Year":         317,
		"Warranty":             323,
		"Storage Capacity":     71,
		"Bluetooth Version":    13,
		"RAM":                  60,
		"Weight":               80,
		"Display Type":         27,
		"Refresh Rate":         61,
		"Screen Size":          66,
		"Backlit Keyboard":     199,
		"Audio Jack":           682,
		"Ram":                  684,
		"Wifi Support":         687,
		"Usb Type":             688,
		"Battery":              689,
		"Gpu Type":             691,
		"Storage":              693,
		"Cpu Type":             696,
		"Cooling Technology":   698,
		"Frequency (Hz)":       704,
		"App Control":          705,
		"Warranty Period":      706,
		// New additions from specs.sql
		"Laptop Type":             24,
		"Battery Type":            12,
		"Charging Speed":          18,
		"Standby Time":            70,
		"Wireless Charging":       82,
		"Resolution":              62,
		"Display Characteristics": 26,
		"Processor Speed":         57,
		"Clock Feature":           20,
		"3.5mm Audio Jack":        2,
		"Camera Features":         16,
		"Audio Quality":           9,
		"Sensors":                 67,
		"Dimensions":              25,
		"Body Type":               89,
		"Cooling System":          95,
		"Available Colors":        10,
		"Special Features":        69,
		// Major laptop specs added to database
		"Processor Cores":       786,
		"Processor Threads":     787,
		"RAM Speed":             788,
		"RAM Slots":             789,
		"RAM Expandable":        790,
		"Storage Interface":     791,
		"Storage Expandable":    792,
		"Graphics VRAM":         793,
		"Display Touch Support": 794,
		"Ethernet":              795,
		"Thunderbolt Version":   796,
		"SD Card Reader":        797,
		"Keyboard Language":     798,
		"Build Standard":        799,
		// Additional available keys
		"Touchscreen": 129,
		"HDMI Ports":  167,
		"USB Ports":   173,
	}

	specs := map[string]string{
		"Brand":                "Lenovo",
		"Model Name":           "Chromebook 3",
		"Operating System":     "Chrome OS",
		"Processor Brand":      "MediaTek",
		"Processor Model":      "MT8183",
		"Processor Generation": "MT8183",
		"Chipset":              "MediaTek Platform Controller Hub",
		"Graphics Card":        "Integrated Graphics",
		"Screen Resolution":    "1366 x 768 pixels",
		"Display Size":         "11.6 inches",
		"Battery Capacity":     "37 Wh",
		"Build Material":       "Plastic",
		"Color":                "Black",
		"Product Weight":       "1.15 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "32 GB",
		"Bluetooth Version":    "Bluetooth 4.2",
		"RAM":                  "4 GB",
		"Weight":               "1.15 kg",
		"Display Type":         "TN",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "11.6 inches",
		"Backlit Keyboard":     "No",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "4 GB",
		"Wifi Support":         "Wi-Fi 5 (802.11ac)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "37 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "32 GB eMMC",
		"Cpu Type":             "MediaTek MT8183",
		"Cooling Technology":   "Passive Cooling",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Entry Level Chromebook",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1366 x 768 pixels",
		"Display Characteristics": "TN, 220 nits",
		"Processor Speed":         "2.0 GHz",
		"Clock Feature":           "1 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Mono Speaker",
		"Sensors":                 "No",
		"Dimensions":              "286 x 200 x 18 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Passive Cooling",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "4",
		"Processor Threads":     "4",
		"RAM Speed":             "LPDDR3",
		"RAM Slots":             "1",
		"RAM Expandable":        "No",
		"Storage Interface":     "eMMC",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "Yes",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 1.4",
		"USB Ports":   "1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 1",
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
		if banglaValue, exists := lc3.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Lenovo Chromebook 3 specifications seeded successfully")
	return nil
}
