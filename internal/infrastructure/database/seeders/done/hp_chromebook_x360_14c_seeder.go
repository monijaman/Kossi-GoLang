package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HpChromebookX36014cSeeder seeds specifications for HP Chromebook x360 14c
type HpChromebookX36014cSeeder struct {
	BaseSeeder
}

// NewHpChromebookX36014cSeeder creates a new HP Chromebook x360 14c seeder
func NewHpChromebookX36014cSeeder() *HpChromebookX36014cSeeder {
	return &HpChromebookX36014cSeeder{
		BaseSeeder: BaseSeeder{name: "HP Chromebook x360 14c Specifications"},
	}
}

func (hcs *HpChromebookX36014cSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HP":                            "এইচপি",
		"Chromebook x360 14c":           "ক্রোমবুক এক্স৩৬০ ১৪সি",
		"Chrome OS":                     "ক্রোম ওএস",
		"MediaTek":                      "মিডিয়াটেক",
		"Kompanio 520":                  "কম্পানিও ৫২০",
		"ARM Cortex-A76":                "আরএম কর্টেক্স-এ৭৬",
		"MediaTek SoC":                  "মিডিয়াটেক এসওসি",
		"Mali-G57 MC2":                  "মালি-জি৫৭ এমসি২",
		"1920 x 1200 pixels":            "১৯২০ x ১২০০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"45 Wh":                         "৪৫ ডব্লিউএইচ",
		"Plastic":                       "প্লাস্টিক",
		"Silver":                        "সিলভার",
		"1.35 kg":                       "১.৩৫ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"128 GB":                        "১২৮ জিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"128 GB eMMC":                   "১২৮ জিবি ইএমএমসি",
		"MediaTek Kompanio 520":         "মিডিয়াটেক কম্পানিও ৫২০",
		"Passive Cooling":               "প্যাসিভ কুলিং",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Chromebook":                           "ক্রোমবুক",
		"Lithium-polymer":                      "লিথিয়াম-পলিমার",
		"45W USB-C Adapter":                    "৪৫ওয়াট ইউএসবি-সি অ্যাডাপ্টার",
		"10-12 hours":                          "১০-১২ ঘন্টা",
		"IPS, 300 nits":                        "আইপিএস, ৩০০ নিটস",
		"2.2 GHz":                              "২.২ গিগাহার্টজ",
		"4 MB cache":                           "৪ এমবি ক্যাশ",
		"1080p HD":                             "১০৮০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"321.9 x 206.5 x 17.9 mm":              "৩২১.৯ x ২০৬.৫ x ১৭.৯ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"8":                                    "৮",
		"LPDDR4X":                              "এলপিডিডিআর৪এক্স",
		"eMMC":                                 "ইএমএমসি",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"2x USB-C 3.2 Gen 1, 1x USB 3.2 Gen 1": "২x ইউএসবি-সি ৩.২ জেন ১, ১x ইউএসবি ৩.২ জেন ১"}
}

// Seed implements the Seeder interface for HP Chromebook x360 14c
func (hcs *HpChromebookX36014cSeeder) Seed(db *gorm.DB) error {
	productSlug := "hp-chromebook-x360-14c"
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
		"Brand":                "HP",
		"Model Name":           "Chromebook x360 14c",
		"Operating System":     "Chrome OS",
		"Processor Brand":      "MediaTek",
		"Processor Model":      "Kompanio 520",
		"Processor Generation": "ARM Cortex-A76",
		"Chipset":              "MediaTek SoC",
		"Graphics Card":        "Mali-G57 MC2",
		"Screen Resolution":    "1920 x 1200 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "45 Wh",
		"Build Material":       "Plastic",
		"Color":                "Silver",
		"Product Weight":       "1.35 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "128 GB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "8 GB",
		"Weight":               "1.35 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "45 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "128 GB eMMC",
		"Cpu Type":             "MediaTek Kompanio 520",
		"Cooling Technology":   "Passive Cooling",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Chromebook",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "45W USB-C Adapter",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.2 GHz",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "321.9 x 206.5 x 17.9 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Passive Cooling",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "8",
		"Processor Threads":     "8",
		"RAM Speed":             "LPDDR4X",
		"RAM Slots":             "Soldered",
		"RAM Expandable":        "No",
		"Storage Interface":     "eMMC",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "Yes",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "Yes",
		"HDMI Ports":  "No",
		"USB Ports":   "2x USB-C 3.2 Gen 1, 1x USB 3.2 Gen 1",
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
		if banglaValue, exists := hcs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ HP Chromebook x360 14c specifications seeded successfully")
	return nil
}
