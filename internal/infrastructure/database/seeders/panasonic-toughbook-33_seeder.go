package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// PanasonicToughbook33Seeder seeds specifications for Panasonic Toughbook 33
type PanasonicToughbook33Seeder struct {
	BaseSeeder
}

// NewPanasonicToughbook33Seeder creates a new Panasonic Toughbook 33 seeder
func NewPanasonicToughbook33Seeder() *PanasonicToughbook33Seeder {
	return &PanasonicToughbook33Seeder{
		BaseSeeder: BaseSeeder{name: "Panasonic Toughbook 33 Specifications"},
	}
}

func (pts *PanasonicToughbook33Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":                            "প্যানাসনিক",
		"Toughbook 33":                         "টাফবুক ৩৩",
		"Windows 10 Pro":                       "উইন্ডোজ ১০ প্রো",
		"Intel":                                "ইন্টেল",
		"Core i5-7300U":                        "কোর আই৫-৭৩০০ইউ",
		"7th Gen":                              "৭ম প্রজন্ম",
		"Intel Platform Controller Hub":        "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel HD Graphics 620":                "ইন্টেল এইচডি গ্রাফিক্স ৬২০",
		"1920 x 1080 pixels":                   "১৯২০ x ১০৮০ পিক্সেল",
		"12.5 inches":                          "১২.৫ ইঞ্চি",
		"3-cell":                               "৩-সেল",
		"Magnesium Alloy":                      "ম্যাগনেসিয়াম অ্যালয়",
		"Black":                                "কালো",
		"1.4 kg":                               "১.৪ কেজি",
		"2018":                                 "২০১৮",
		"3 Year Warranty":                      "৩ বছর ওয়ারেন্টি",
		"256 GB":                               "২৫৬ জিবি",
		"Bluetooth 4.2":                        "ব্লুটুথ ৪.২",
		"8 GB":                                 "৮ জিবি",
		"IPS":                                  "আইপিএস",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":                   "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.1 Gen 1":                        "ইউএসবি ৩.১ জেন ১",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"256 GB SSD":                           "২৫৬ জিবি এসএসডি",
		"Intel Core i5":                        "ইন্টেল কোর আই৫",
		"Single Fan":                           "সিঙ্গেল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"3 Year":                               "৩ বছর",
		"Ultra-Rugged":                         "আল্ট্রা-রাগড",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"45W AC Adapter":                       "৪৫ওয়াট এসি অ্যাডাপ্টার",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"IPS, 1000 nits":                       "আইপিএস, ১০০০ নিটস",
		"2.6 GHz base / 3.5 GHz boost":         "২.৬ গিগাহার্টজ বেস / ৩.৫ গিগাহার্টজ বুস্ট",
		"3 MB cache":                           "৩ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Mono Speaker":                         "মনো স্পিকার",
		"304 x 216 x 29.9 mm":                  "৩০৪ x ২১৬ x ২৯.৯ মিমি",
		"TPM 1.2, MIL-STD-810G":                "টিপিএম ১.২, এমআইএল-এসটিডি-৮১০জি",
		"2":                                    "২",
		"4":                                    "৪",
		"2133 MHz":                             "২১৩৩ মেগাহার্টজ",
		"SATA III":                             "সাটা তৃতীয়",
		"Shared":                               "শেয়ার্ড",
		"English":                              "ইংরেজি",
		"HDMI 1.4":                             "এইচডিএমআই ১.৪",
		"2x USB 3.1 Gen 1, 1x USB-C 3.1 Gen 1": "২x ইউএসবি ৩.১ জেন ১, ১x ইউএসবি-সি ৩.১ জেন ১",
	}
}

// Seed implements the Seeder interface for Panasonic Toughbook 33
func (pts *PanasonicToughbook33Seeder) Seed(db *gorm.DB) error {
	productSlug := "panasonic-toughbook-33"
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
		"Brand":                   "Panasonic",
		"Model Name":              "Toughbook 33",
		"Operating System":        "Windows 10 Pro",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i5-7300U",
		"Processor Generation":    "7th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel HD Graphics 620",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "12.5 inches",
		"Battery Capacity":        "3-cell",
		"Build Material":          "Magnesium Alloy",
		"Color":                   "Black",
		"Product Weight":          "1.4 kg",
		"Release Year":            "2018",
		"Warranty":                "3 Year Warranty",
		"Storage Capacity":        "256 GB",
		"Bluetooth Version":       "Bluetooth 4.2",
		"RAM":                     "8 GB",
		"Weight":                  "1.4 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "12.5 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 5 (802.11ac)",
		"Usb Type":                "USB 3.1 Gen 1",
		"Battery":                 "3-cell",
		"Gpu Type":                "Integrated",
		"Storage":                 "256 GB SSD",
		"Cpu Type":                "Intel Core i5",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "3 Year",
		"Laptop Type":             "Ultra-Rugged",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W AC Adapter",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 1000 nits",
		"Processor Speed":         "2.6 GHz base / 3.5 GHz boost",
		"Clock Feature":           "3 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Mono Speaker",
		"Sensors":                 "No",
		"Dimensions":              "304 x 216 x 29.9 mm",
		"Body Type":               "Ultra-Rugged",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 1.2, MIL-STD-810G",
		"Processor Cores":         "2",
		"Processor Threads":       "4",
		"RAM Speed":               "2133 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "No",
		"Storage Interface":       "SATA III",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Ultra-Rugged",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 1.4",
		"USB Ports":               "2x USB 3.1 Gen 1, 1x USB-C 3.1 Gen 1",
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
		if banglaValue, exists := pts.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Panasonic Toughbook 33 specifications seeded successfully")
	return nil
}
