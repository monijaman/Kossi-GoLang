package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AsusChromebookFlipC434Seeder seeds specifications for Asus Chromebook Flip C434
type AsusChromebookFlipC434Seeder struct {
	BaseSeeder
}

// NewAsusChromebookFlipC434Seeder creates a new Asus Chromebook Flip C434 seeder
func NewAsusChromebookFlipC434Seeder() *AsusChromebookFlipC434Seeder {
	return &AsusChromebookFlipC434Seeder{
		BaseSeeder: BaseSeeder{name: "Asus Chromebook Flip C434 Specifications"},
	}
}

func (ac *AsusChromebookFlipC434Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Asus":                                 "এসুস",
		"Chromebook Flip C434":                 "ক্রোমবুক ফ্লিপ সি৪৩৪",
		"Chrome OS":                            "ক্রোম ওএস",
		"Intel":                                "ইন্টেল",
		"Celeron N4020":                        "সেলেরন এন৪০২০",
		"10th Gen":                             "১০তম প্রজন্ম",
		"Intel Platform Controller Hub":        "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD Graphics 600":               "ইন্টেল ইউএইচডি গ্রাফিক্স ৬০০",
		"1366 x 768 pixels":                    "১৩৬৬ x ৭৬৮ পিক্সেল",
		"14 inches":                            "১৪ ইঞ্চি",
		"42 Wh":                                "৪২ ডব্লিউএইচ",
		"Plastic":                              "প্লাস্টিক",
		"Silver":                               "সিলভার",
		"1.3 kg":                               "১.৩ কেজি",
		"2021":                                 "২০২১",
		"1 Year International Warranty":        "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"64 GB":                                "৬৪ জিবি",
		"Bluetooth 4.2":                        "ব্লুটুথ ৪.২",
		"4 GB":                                 "৪ জিবি",
		"TN":                                   "টিএন",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":                   "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.2 Gen 1":                        "ইউএসবি ৩.২ জেন ১",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"64 GB eMMC":                           "৬৪ জিবি ইএমএমসি",
		"Intel Celeron":                        "ইন্টেল সেলেরন",
		"Single Fan":                           "সিঙ্গেল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Convertible":                          "কনভার্টিবল",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"45W":                                  "৪৫ওয়াট",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"TN, 250 nits":                         "টিএন, ২৫০ নিটস",
		"1.1 GHz base / 2.8 GHz boost":         "১.১ গিগাহার্টজ বেস / ২.৮ গিগাহার্টজ বুস্ট",
		"4 MB cache":                           "৪ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"326 x 216 x 16 mm":                    "৩২৬ x ২১৬ x ১৬ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"2":                                    "২",
		"2400 MHz":                             "২৪০০ মেগাহার্টজ",
		"eMMC":                                 "ইএমএমসি",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 1.4":                             "এইচডিএমআই ১.৪",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 1": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ১",
	}
}

// Seed implements the Seeder interface for Asus Chromebook Flip C434
func (ac *AsusChromebookFlipC434Seeder) Seed(db *gorm.DB) error {
	productSlug := "asus-chromebook-flip-c434"
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
		"Brand":                   "Asus",
		"Model Name":              "Chromebook Flip C434",
		"Operating System":        "Chrome OS",
		"Processor Brand":         "Intel",
		"Processor Model":         "Celeron N4020",
		"Processor Generation":    "10th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel UHD Graphics 600",
		"Screen Resolution":       "1366 x 768 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "42 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Silver",
		"Product Weight":          "1.3 kg",
		"Release Year":            "2021",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "64 GB",
		"Bluetooth Version":       "Bluetooth 4.2",
		"RAM":                     "4 GB",
		"Weight":                  "1.3 kg",
		"Display Type":            "TN",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "4 GB",
		"Wifi Support":            "Wi-Fi 5 (802.11ac)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "42 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "64 GB eMMC",
		"Cpu Type":                "Intel Celeron",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Convertible",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1366 x 768 pixels",
		"Display Characteristics": "TN, 250 nits",
		"Processor Speed":         "1.1 GHz base / 2.8 GHz boost",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "326 x 216 x 16 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "2",
		"Processor Threads":       "2",
		"RAM Speed":               "2400 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "No",
		"Storage Interface":       "eMMC",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 1.4",
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
		if banglaValue, exists := ac.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Asus Chromebook Flip C434 specifications seeded successfully")
	return nil
}
