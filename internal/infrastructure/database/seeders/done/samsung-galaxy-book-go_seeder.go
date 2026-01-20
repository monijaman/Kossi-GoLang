package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SamsungGalaxyBookGoSeeder seeds specifications for Samsung Galaxy Book Go
type SamsungGalaxyBookGoSeeder struct {
	BaseSeeder
}

// NewSamsungGalaxyBookGoSeeder creates a new Samsung Galaxy Book Go seeder
func NewSamsungGalaxyBookGoSeeder() *SamsungGalaxyBookGoSeeder {
	return &SamsungGalaxyBookGoSeeder{
		BaseSeeder: BaseSeeder{name: "Samsung Galaxy Book Go Specifications"},
	}
}

func (sgbgs *SamsungGalaxyBookGoSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Samsung":                              "স্যামসাং",
		"Galaxy Book Go":                       "গ্যালাক্সি বুক গো",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"Qualcomm":                             "কোয়ালকম",
		"Snapdragon 7c Gen 2":                  "স্ন্যাপড্রাগন ৭সি জেন ২",
		"7th Gen":                              "৭ম প্রজন্ম",
		"Qualcomm Snapdragon":                  "কোয়ালকম স্ন্যাপড্রাগন",
		"Qualcomm Adreno":                      "কোয়ালকম অ্যাড্রেনো",
		"1920 x 1080 pixels":                   "১৯২০ x ১০৮০ পিক্সেল",
		"14 inches":                            "১৪ ইঞ্চি",
		"42.3 Wh":                              "৪২.৩ ডব্লিউএইচ",
		"Plastic":                              "প্লাস্টিক",
		"Silver":                               "সিলভার",
		"1.38 kg":                              "১.৩৮ কেজি",
		"2021":                                 "২০২১",
		"1 Year Warranty":                      "১ বছর ওয়ারেন্টি",
		"128 GB":                               "১২৮ জিবি",
		"Bluetooth 5.1":                        "ব্লুটুথ ৫.১",
		"4 GB":                                 "৪ জিবি",
		"IPS":                                  "আইপিএস",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                        "ইউএসবি ৩.২ জেন ১",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"128 GB eMMC":                          "১২৮ জিবি ইএমএমসি",
		"Qualcomm Snapdragon 7c":               "কোয়ালকম স্ন্যাপড্রাগন ৭সি",
		"Passive Cooling":                      "প্যাসিভ কুলিং",
		"60":                                   "৬০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Budget":                               "বাজেট",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"25W USB-C":                            "২৫ওয়াট ইউএসবি-সি",
		"10-12 hours":                          "১০-১২ ঘন্টা",
		"IPS, 250 nits":                        "আইপিএস, ২৫০ নিটস",
		"2.55 GHz":                             "২.৫৫ গিগাহার্টজ",
		"4 MB cache":                           "৪ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"324 x 225.8 x 15.5 mm":                "৩২৪ x ২২৫.৮ x ১৫.৫ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"8":                                    "৮",
		"2133 MHz":                             "২১৩৩ মেগাহার্টজ",
		"eMMC":                                 "ইএমএমসি",
		"Shared":                               "শেয়ার্ড",
		"English":                              "ইংরেজি",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 1": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ১",
	}
}

// Seed implements the Seeder interface for Samsung Galaxy Book Go
func (sgbgs *SamsungGalaxyBookGoSeeder) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-book-go"
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
		"Brand":                   310,
		"Model Name":              316,
		"Operating System":        49,
		"Processor Brand":         204,
		"Processor Model":         206,
		"Processor Generation":    205,
		"Chipset":                 19,
		"Graphics Card":           200,
		"Screen Resolution":       207,
		"Display Size":            683,
		"Battery Capacity":        11,
		"Build Material":          14,
		"Color":                   318,
		"Product Weight":          697,
		"Release Year":            317,
		"Warranty":                323,
		"Storage Capacity":        71,
		"Bluetooth Version":       13,
		"RAM":                     60,
		"Weight":                  80,
		"Display Type":            27,
		"Refresh Rate":            61,
		"Screen Size":             66,
		"Backlit Keyboard":        199,
		"Audio Jack":              682,
		"Ram":                     684,
		"Wifi Support":            687,
		"Usb Type":                688,
		"Battery":                 689,
		"Gpu Type":                691,
		"Storage":                 693,
		"Cpu Type":                696,
		"Cooling Technology":      698,
		"Frequency (Hz)":          704,
		"App Control":             705,
		"Warranty Period":         706,
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
		"Processor Cores":         786,
		"Processor Threads":       787,
		"RAM Speed":               788,
		"RAM Slots":               789,
		"RAM Expandable":          790,
		"Storage Interface":       791,
		"Storage Expandable":      792,
		"Graphics VRAM":           793,
		"Display Touch Support":   794,
		"Ethernet":                795,
		"Thunderbolt Version":     796,
		"SD Card Reader":          797,
		"Keyboard Language":       798,
		"Build Standard":          799,
		"Touchscreen":             129,
		"HDMI Ports":              167,
		"USB Ports":               173,
	}

	specs := map[string]string{
		"Brand":                   "Samsung",
		"Model Name":              "Galaxy Book Go",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Qualcomm",
		"Processor Model":         "Snapdragon 7c Gen 2",
		"Processor Generation":    "7th Gen",
		"Chipset":                 "Qualcomm Snapdragon",
		"Graphics Card":           "Qualcomm Adreno",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "42.3 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Silver",
		"Product Weight":          "1.38 kg",
		"Release Year":            "2021",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "128 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "4 GB",
		"Weight":                  "1.38 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "4 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "42.3 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "128 GB eMMC",
		"Cpu Type":                "Qualcomm Snapdragon 7c",
		"Cooling Technology":      "Passive Cooling",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Budget",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "25W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 250 nits",
		"Processor Speed":         "2.55 GHz",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "324 x 225.8 x 15.5 mm",
		"Body Type":               "Budget",
		"Cooling System":          "Passive Cooling",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "8",
		"Processor Threads":       "8",
		"RAM Speed":               "2133 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "No",
		"Storage Interface":       "eMMC",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "microSD",
		"Keyboard Language":       "English",
		"Build Standard":          "Budget",
		"Touchscreen":             "No",
		"HDMI Ports":              "No",
		"USB Ports":               "1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 1",
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
		if banglaValue, exists := sgbgs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Samsung Galaxy Book Go specifications seeded successfully")
	return nil
}
