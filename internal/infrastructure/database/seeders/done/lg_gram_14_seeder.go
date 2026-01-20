package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LGGram14Seeder seeds specifications for LG Gram 14
type LGGram14Seeder struct {
	BaseSeeder
}

// NewLGGram14Seeder creates a new LG Gram 14 seeder
func NewLGGram14Seeder() *LGGram14Seeder {
	return &LGGram14Seeder{
		BaseSeeder: BaseSeeder{name: "LG Gram 14 Specifications"},
	}
}

func (lg14 *LGGram14Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                            "এলজি",
		"Gram 14":                       "গ্রাম ১৪",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i5-1340P":                 "কোর আই৫-১৩৪০পি",
		"13th Generation":               "১৩তম জেনারেশন",
		"Intel Evo Platform":            "ইন্টেল ইভো প্ল্যাটফর্ম",
		"Intel Iris Xe Graphics":        "ইন্টেল আইরিস এক্সই গ্রাফিক্স",
		"1920 x 1200 pixels":            "১৯২০ x ১২০০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"72 Wh":                         "৭২ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Dark Silver":                   "ডার্ক সিলভার",
		"0.99 kg":                       "০.৯৯ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.1":                 "ব্লুটুথ ৫.১",
		"16 GB":                         "১৬ জিবি",
		"OLED":                          "ওএলইডি",
		"90 Hz":                         "৯০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"Intel Core i5-1340P":           "ইন্টেল কোর আই৫-১৩৪০পি",
		"Passive Cooling":               "প্যাসিভ কুলিং",
		"90":                            "৯০",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Ultraportable":                        "আল্ট্রাপোর্টেবল",
		"Lithium-polymer":                      "লিথিয়াম-পলিমার",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"16-20 hours":                          "১৬-২০ ঘন্টা",
		"OLED, 350 nits":                       "ওএলইডি, ৩৫০ নিটস",
		"1.9 GHz (up to 4.6 GHz)":              "১.৯ গিগাহার্টজ (আপ টু ৪.৬ গিগাহার্টজ)",
		"12 MB cache":                          "১২ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Dolby Atmos Speakers":                 "ডলবি অ্যাটমস স্পিকার",
		"312 x 213 x 16.8 mm":                  "৩১২ x ২১৩ x ১৬.৮ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"12":                                   "১২",
		"16":                                   "১৬",
		"LPDDR5":                               "এলপিডিডিআর৫",
		"No":                                   "না",
		"PCIe 4.0":                             "পিসিআইই ৪.০",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Military Grade":                       "মিলিটারি গ্রেড",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"1x USB 3.2 Gen 2, 2x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ২, ২x ইউএসবি-সি ৩.২ জেন ২",
		"Thunderbolt 4":                        "থান্ডারবোল্ট ৪",
	}
}

// Seed implements the Seeder interface for LG Gram 14
func (lg14 *LGGram14Seeder) Seed(db *gorm.DB) error {
	productSlug := "lg-gram-14"
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
		"Brand":                "LG",
		"Model Name":           "Gram 14",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i5-1340P",
		"Processor Generation": "13th Generation",
		"Chipset":              "Intel Evo Platform",
		"Graphics Card":        "Intel Iris Xe Graphics",
		"Screen Resolution":    "1920 x 1200 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "72 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Dark Silver",
		"Product Weight":       "0.99 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.1",
		"RAM":                  "16 GB",
		"Weight":               "0.99 kg",
		"Display Type":         "OLED",
		"Refresh Rate":         "90 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 2",
		"Battery":              "72 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "Intel Core i5-1340P",
		"Cooling Technology":   "Passive Cooling",
		"Frequency (Hz)":       "90",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "16-20 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "OLED, 350 nits",
		"Processor Speed":         "1.9 GHz (up to 4.6 GHz)",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Dolby Atmos Speakers",
		"Sensors":                 "No",
		"Dimensions":              "312 x 213 x 16.8 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Passive Cooling",
		"Available Colors":        "Dark Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "12",
		"Processor Threads":     "16",
		"RAM Speed":             "LPDDR5",
		"RAM Slots":             "Soldered",
		"RAM Expandable":        "No",
		"Storage Interface":     "PCIe 4.0",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "Thunderbolt 4",
		"SD Card Reader":        "Yes",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Military Grade",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "1x USB 3.2 Gen 2, 2x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := lg14.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ LG Gram 14 specifications seeded successfully")
	return nil
}
