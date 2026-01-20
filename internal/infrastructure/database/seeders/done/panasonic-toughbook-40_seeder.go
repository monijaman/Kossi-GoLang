package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// PanasonicToughbook40Seeder seeds specifications for Panasonic Toughbook 40
type PanasonicToughbook40Seeder struct {
	BaseSeeder
}

// NewPanasonicToughbook40Seeder creates a new Panasonic Toughbook 40 seeder
func NewPanasonicToughbook40Seeder() *PanasonicToughbook40Seeder {
	return &PanasonicToughbook40Seeder{
		BaseSeeder: BaseSeeder{name: "Panasonic Toughbook 40 Specifications"},
	}
}

func (pts *PanasonicToughbook40Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":                            "প্যানাসনিক",
		"Toughbook 40":                         "টাফবুক ৪০",
		"Windows 11 Pro":                       "উইন্ডোজ ১১ প্রো",
		"Intel":                                "ইন্টেল",
		"Core i5-1135G7":                       "কোর আই৫-১১৩৫জি৭",
		"11th Gen":                             "১১তম প্রজন্ম",
		"Intel Platform Controller Hub":        "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                        "ইন্টেল আইরিস এক্সই",
		"1920 x 1200 pixels":                   "১৯২০ x ১২০০ পিক্সেল",
		"14 inches":                            "১৪ ইঞ্চি",
		"4-cell":                               "৪-সেল",
		"Magnesium Alloy":                      "ম্যাগনেসিয়াম অ্যালয়",
		"Black":                                "কালো",
		"1.8 kg":                               "১.৮ কেজি",
		"2022":                                 "২০২২",
		"3 Year Warranty":                      "৩ বছর ওয়ারেন্টি",
		"256 GB":                               "২৫৬ জিবি",
		"Bluetooth 5.1":                        "ব্লুটুথ ৫.১",
		"8 GB":                                 "৮ জিবি",
		"IPS":                                  "আইপিএস",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                        "ইউএসবি ৩.২ জেন ১",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"256 GB SSD":                           "২৫৬ জিবি এসএসডি",
		"Intel Core i5":                        "ইন্টেল কোর আই৫",
		"Dual Fan":                             "ডুয়াল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"3 Year":                               "৩ বছর",
		"Rugged":                               "রাগড",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"65W AC Adapter":                       "৬৫ওয়াট এসি অ্যাডাপ্টার",
		"10-12 hours":                          "১০-১২ ঘন্টা",
		"IPS, 400 nits":                        "আইপিএস, ৪০০ নিটস",
		"2.4 GHz base / 4.2 GHz boost":         "২.৪ গিগাহার্টজ বেস / ৪.২ গিগাহার্টজ বুস্ট",
		"8 MB cache":                           "৮ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"330 x 230 x 24.9 mm":                  "৩৩০ x ২৩০ x ২৪.৯ মিমি",
		"TPM 2.0, MIL-STD-810G":                "টিপিএম ২.০, এমআইএল-এসটিডি-৮১০জি",
		"4":                                    "৪",
		"8":                                    "৮",
		"3200 MHz":                             "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"Shared":                               "শেয়ার্ড",
		"English":                              "ইংরেজি",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "২x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Panasonic Toughbook 40
func (pts *PanasonicToughbook40Seeder) Seed(db *gorm.DB) error {
	productSlug := "panasonic-toughbook-40"
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
		"Model Name":              "Toughbook 40",
		"Operating System":        "Windows 11 Pro",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i5-1135G7",
		"Processor Generation":    "11th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "1920 x 1200 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "4-cell",
		"Build Material":          "Magnesium Alloy",
		"Color":                   "Black",
		"Product Weight":          "1.8 kg",
		"Release Year":            "2022",
		"Warranty":                "3 Year Warranty",
		"Storage Capacity":        "256 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "8 GB",
		"Weight":                  "1.8 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "4-cell",
		"Gpu Type":                "Integrated",
		"Storage":                 "256 GB SSD",
		"Cpu Type":                "Intel Core i5",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "3 Year",
		"Laptop Type":             "Rugged",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W AC Adapter",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.4 GHz base / 4.2 GHz boost",
		"Clock Feature":           "8 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "330 x 230 x 24.9 mm",
		"Body Type":               "Rugged",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0, MIL-STD-810G",
		"Processor Cores":         "4",
		"Processor Threads":       "8",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Rugged",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2",
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

	log.Printf("✅ Panasonic Toughbook 40 specifications seeded successfully")
	return nil
}
