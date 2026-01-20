package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// PanasonicToughbook55Seeder seeds specifications for Panasonic Toughbook 55
type PanasonicToughbook55Seeder struct {
	BaseSeeder
}

// NewPanasonicToughbook55Seeder creates a new Panasonic Toughbook 55 seeder
func NewPanasonicToughbook55Seeder() *PanasonicToughbook55Seeder {
	return &PanasonicToughbook55Seeder{
		BaseSeeder: BaseSeeder{name: "Panasonic Toughbook 55 Specifications"},
	}
}

func (pts *PanasonicToughbook55Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Panasonic":                        "প্যানাসনিক",
		"Toughbook 55":                     "টাফবুক ৫৫",
		"Windows 11 Pro":                   "উইন্ডোজ ১১ প্রো",
		"Intel":                            "ইন্টেল",
		"Core i5-12450H":                   "কোর আই৫-১২৪৫০এইচ",
		"12th Gen":                         "১২তম প্রজন্ম",
		"Intel Platform Controller Hub":    "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                    "ইন্টেল আইরিস এক্সই",
		"1920 x 1080 pixels":               "১৯২০ x ১০৮০ পিক্সেল",
		"15.6 inches":                      "১৫.৬ ইঞ্চি",
		"6-cell":                           "৬-সেল",
		"Magnesium Alloy":                  "ম্যাগনেসিয়াম অ্যালয়",
		"Black":                            "কালো",
		"2.5 kg":                           "২.৫ কেজি",
		"2023":                             "২০২৩",
		"3 Year Warranty":                  "৩ বছর ওয়ারেন্টি",
		"512 GB":                           "৫১২ জিবি",
		"Bluetooth 5.1":                    "ব্লুটুথ ৫.১",
		"16 GB":                            "১৬ জিবি",
		"IPS":                              "আইপিএস",
		"60 Hz":                            "৬০ হার্জ",
		"Yes":                              "হ্যাঁ",
		"3.5mm Combo Jack":                 "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":               "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                    "ইউএসবি ৩.২ জেন ১",
		"Integrated":                       "ইন্টিগ্রেটেড",
		"512 GB SSD":                       "৫১২ জিবি এসএসডি",
		"Intel Core i5":                    "ইন্টেল কোর আই৫",
		"Dual Fan":                         "ডুয়াল ফ্যান",
		"60":                               "৬০",
		"No":                               "না",
		"3 Year":                           "৩ বছর",
		"Rugged":                           "রাগড",
		"Lithium-ion":                      "লিথিয়াম-আয়ন",
		"90W AC Adapter":                   "৯০ওয়াট এসি অ্যাডাপ্টার",
		"8-10 hours":                       "৮-১০ ঘন্টা",
		"IPS, 300 nits":                    "আইপিএস, ৩০০ নিটস",
		"2.0 GHz base / 4.4 GHz boost":     "২.০ গিগাহার্টজ বেস / ৪.৪ গিগাহার্টজ বুস্ট",
		"12 MB cache":                      "১২ এমবি ক্যাশ",
		"1080p FHD":                        "১০৮০পি এফএইচডি",
		"Stereo Speakers with Dolby Audio": "ডলবি অডিও সহ স্টেরিও স্পিকার",
		"378 x 260 x 29.8 mm":              "৩৭৮ x ২৬০ x ২৯.৮ মিমি",
		"TPM 2.0, MIL-STD-810H":            "টিপিএম ২.০, এমআইএল-এসটিডি-৮১০এইচ",
		"8":                                "৮",
		"16":                               "১৬",
		"3200 MHz":                         "৩২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                 "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                   "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                   "হ্যাঁ (এম.২ স্লট)",
		"Shared":                           "শেয়ার্ড",
		"English":                          "ইংরেজি",

		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "২x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Panasonic Toughbook 55
func (pts *PanasonicToughbook55Seeder) Seed(db *gorm.DB) error {
	productSlug := "panasonic-toughbook-55"
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
		"Model Name":              "Toughbook 55",
		"Operating System":        "Windows 11 Pro",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i5-12450H",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "15.6 inches",
		"Battery Capacity":        "6-cell",
		"Build Material":          "Magnesium Alloy",
		"Color":                   "Black",
		"Product Weight":          "2.5 kg",
		"Release Year":            "2023",
		"Warranty":                "3 Year Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "16 GB",
		"Weight":                  "2.5 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "15.6 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "6-cell",
		"Gpu Type":                "Integrated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "Intel Core i5",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "3 Year",
		"Laptop Type":             "Rugged",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "90W AC Adapter",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.0 GHz base / 4.4 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers with Dolby Audio",
		"Sensors":                 "No",
		"Dimensions":              "378 x 260 x 29.8 mm",
		"Body Type":               "Rugged",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0, MIL-STD-810H",
		"Processor Cores":         "8",
		"Processor Threads":       "16",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Rugged",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.1",
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

	log.Printf("✅ Panasonic Toughbook 55 specifications seeded successfully")
	return nil
}
