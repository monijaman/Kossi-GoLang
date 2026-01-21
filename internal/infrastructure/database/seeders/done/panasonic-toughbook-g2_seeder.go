package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// PanasonicToughbookG2Seeder seeds specifications for Panasonic Toughbook G2
type PanasonicToughbookG2Seeder struct {
	BaseSeeder
}

// NewPanasonicToughbookG2Seeder creates a new Panasonic Toughbook G2 seeder
func NewPanasonicToughbookG2Seeder() *PanasonicToughbookG2Seeder {
	return &PanasonicToughbookG2Seeder{
		BaseSeeder: BaseSeeder{name: "Panasonic Toughbook G2 Specifications"},
	}
}

func (pts *PanasonicToughbookG2Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":                     "প্যানাসনিক",
		"Toughbook G2":                  "টাফবুক জি২",
		"Windows 10 Pro":                "উইন্ডোজ ১০ প্রো",
		"Intel":                         "ইন্টেল",
		"Core i5-4300U":                 "কোর আই৫-৪৩০০ইউ",
		"4th Gen":                       "৪র্থ প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel HD Graphics 4400":        "ইন্টেল এইচডি গ্রাফিক্স ৪৪০০",
		"1600 x 900 pixels":             "১৬০০ x ৯০০ পিক্সেল",
		"12.5 inches":                   "১২.৫ ইঞ্চি",
		"4-cell":                        "৪-সেল",
		"Magnesium Alloy":               "ম্যাগনেসিয়াম অ্যালয়",
		"Black":                         "কালো",
		"1.5 kg":                        "১.৫ কেজি",
		"2013":                          "২০১৩",
		"3 Year Warranty":               "৩ বছর ওয়ারেন্টি",
		"128 GB":                        "১২৮ জিবি",
		"Bluetooth 4.0":                 "ব্লুটুথ ৪.০",
		"8 GB":                          "৮ জিবি",
		"TN":                            "টিএন",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 4 (802.11n)":             "ওয়াই-ফাই ৪ (৮০২.১১এন)",
		"USB 3.0":                       "ইউএসবি ৩.০",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"128 GB SSD":                    "১২৮ জিবি এসএসডি",
		"Intel Core i5":                 "ইন্টেল কোর আই৫",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"3 Year":                        "৩ বছর",
		"Rugged":                        "রাগড",
		"Lithium-ion":                   "লিথিয়াম-আয়ন",
		"45W AC Adapter":                "৪৫ওয়াট এসি অ্যাডাপ্টার",
		"6-8 hours":                     "৬-৮ ঘন্টা",
		"TN, 400 nits":                  "টিএন, ৪০০ নিটস",
		"1.9 GHz base / 2.9 GHz boost":  "১.৯ গিগাহার্টজ বেস / ২.৯ গিগাহার্টজ বুস্ট",
		"3 MB cache":                    "৩ এমবি ক্যাশ",
		"VGA":                           "ভিজিএ",
		"Mono Speaker":                  "মনো স্পিকার",
		"304 x 216 x 29.9 mm":           "৩০৪ x ২১৬ x ২৯.৯ মিমি",
		"TPM 1.2, MIL-STD-810F":         "টিপিএম ১.২, এমআইএল-এসটিডি-৮১০এফ",
		"2":                             "২",
		"4":                             "৪",
		"1600 MHz":                      "১৬০০ মেগাহার্টজ",
		"SATA II":                       "সাটা দ্বিতীয়",
		"Shared":                        "শেয়ার্ড",
		"English":                       "ইংরেজি",
		"2x USB 3.0, 1x USB 2.0":        "২x ইউএসবি ৩.০, ১x ইউএসবি ২.০",
	}
}

// Seed implements the Seeder interface for Panasonic Toughbook G2
func (pts *PanasonicToughbookG2Seeder) Seed(db *gorm.DB) error {
	productSlug := "panasonic-toughbook-g2"
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
		"Model Name":              "Toughbook G2",
		"Operating System":        "Windows 10 Pro",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i5-4300U",
		"Processor Generation":    "4th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel HD Graphics 4400",
		"Screen Resolution":       "1600 x 900 pixels",
		"Display Size":            "12.5 inches",
		"Battery Capacity":        "4-cell",
		"Build Material":          "Magnesium Alloy",
		"Color":                   "Black",
		"Product Weight":          "1.5 kg",
		"Release Year":            "2013",
		"Warranty":                "3 Year Warranty",
		"Storage Capacity":        "128 GB",
		"Bluetooth Version":       "Bluetooth 4.0",
		"RAM":                     "8 GB",
		"Weight":                  "1.5 kg",
		"Display Type":            "TN",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "12.5 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 4 (802.11n)",
		"Usb Type":                "USB 3.0",
		"Battery":                 "4-cell",
		"Gpu Type":                "Integrated",
		"Storage":                 "128 GB SSD",
		"Cpu Type":                "Intel Core i5",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "3 Year",
		"Laptop Type":             "Rugged",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W AC Adapter",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1600 x 900 pixels",
		"Display Characteristics": "TN, 400 nits",
		"Processor Speed":         "1.9 GHz base / 2.9 GHz boost",
		"Clock Feature":           "3 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "VGA",
		"Audio Quality":           "Mono Speaker",
		"Sensors":                 "No",
		"Dimensions":              "304 x 216 x 29.9 mm",
		"Body Type":               "Rugged",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 1.2, MIL-STD-810F",
		"Processor Cores":         "2",
		"Processor Threads":       "4",
		"RAM Speed":               "1600 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "No",
		"Storage Interface":       "SATA II",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Rugged",
		"Touchscreen":             "No",
		"HDMI Ports":              "VGA",
		"USB Ports":               "2x USB 3.0, 1x USB 2.0",
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

	log.Printf("✅ Panasonic Toughbook G2 specifications seeded successfully")
	return nil
}
