package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// RazerBlade14Seeder seeds specifications for Razer Blade 14
type RazerBlade14Seeder struct {
	BaseSeeder
}

// NewRazerBlade14Seeder creates a new Razer Blade 14 seeder
func NewRazerBlade14Seeder() *RazerBlade14Seeder {
	return &RazerBlade14Seeder{
		BaseSeeder: BaseSeeder{name: "Razer Blade 14 Specifications"},
	}
}

func (rbs *RazerBlade14Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Razer":                                  "রেজার",
		"Blade 14":                               "ব্লেড ১৪",
		"Windows 11 Home":                        "উইন্ডোজ ১১ হোম",
		"Intel":                                  "ইন্টেল",
		"Core i7-12800H":                         "কোর আই৭-১২৮০০এইচ",
		"12th Gen":                               "১২তম প্রজন্ম",
		"Intel Platform Controller Hub":          "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 4060":                "এনভিডিয়া জিফোর্স আরটিএক্স ৪০৬০",
		"2560 x 1600 pixels":                     "২৫৬০ x ১৬০০ পিক্সেল",
		"14 inches":                              "১৪ ইঞ্চি",
		"61.6 Wh":                                "৬১.৬ ডব্লিউএইচ",
		"Aluminum":                               "অ্যালুমিনিয়াম",
		"Black":                                  "কালো",
		"1.78 kg":                                "১.৭৮ কেজি",
		"2023":                                   "২০২৩",
		"1 Year Warranty":                        "১ বছর ওয়ারেন্টি",
		"512 GB":                                 "৫১২ জিবি",
		"Bluetooth 5.3":                          "ব্লুটুথ ৫.৩",
		"16 GB":                                  "১৬ জিবি",
		"IPS":                                    "আইপিএস",
		"165 Hz":                                 "১৬৫ হার্জ",
		"Yes":                                    "হ্যাঁ",
		"3.5mm Combo Jack":                       "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":                    "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                          "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                              "ডেডিকেটেড",
		"512 GB SSD":                             "৫১২ জিবি এসএসডি",
		"Intel Core i7":                          "ইন্টেল কোর আই৭",
		"Dual Fan":                               "ডুয়াল ফ্যান",
		"165":                                    "১৬৫",
		"No":                                     "না",
		"1 Year":                                 "১ বছর",
		"Gaming":                                 "গেমিং",
		"Lithium-ion":                            "লিথিয়াম-আয়ন",
		"230W AC Adapter":                        "২৩০ওয়াট এসি অ্যাডাপ্টার",
		"6-8 hours":                              "৬-৮ ঘন্টা",
		"IPS, 500 nits":                          "আইপিএস, ৫০০ নিটস",
		"2.4 GHz base / 4.8 GHz boost":           "২.৪ গিগাহার্টজ বেস / ৪.৮ গিগাহার্টজ বুস্ট",
		"24 MB cache":                            "২৪ এমবি ক্যাশ",
		"1080p FHD":                              "১০৮০পি এফএইচডি",
		"Stereo Speakers with THX Spatial Audio": "থেক্স স্পেশাল অডিও সহ স্টেরিও স্পিকার",
		"310 x 220 x 16.8 mm":                    "৩১০ x ২২০ x ১৬.৮ মিমি",
		"TPM 2.0, Chroma RGB Keyboard":           "টিপিএম ২.০, ক্রোমা আরজিবি কীবোর্ড",
		"14":                                     "১৪",
		"20":                                     "২০",
		"4800 MHz":                               "৪৮০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                       "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                         "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                         "হ্যাঁ (এম.২ স্লট)",
		"8 GB GDDR6":                             "৮ জিবি জিডিডিআর৬",
		"English":                                "ইংরেজি",
		"HDMI 2.1":                               "এইচডিএমআই ২.১",
		"2x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2":   "২x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Razer Blade 14
func (rbs *RazerBlade14Seeder) Seed(db *gorm.DB) error {
	productSlug := "razer-blade-14"
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
		"Brand":                   "Razer",
		"Model Name":              "Blade 14",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-12800H",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "NVIDIA GeForce RTX 4060",
		"Screen Resolution":       "2560 x 1600 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "61.6 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "1.78 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.3",
		"RAM":                     "16 GB",
		"Weight":                  "1.78 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "165 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "61.6 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "165",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "230W AC Adapter",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "IPS, 500 nits",
		"Processor Speed":         "2.4 GHz base / 4.8 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers with THX Spatial Audio",
		"Sensors":                 "No",
		"Dimensions":              "310 x 220 x 16.8 mm",
		"Body Type":               "Gaming",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0, Chroma RGB Keyboard",
		"Processor Cores":         "14",
		"Processor Threads":       "20",
		"RAM Speed":               "4800 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "8 GB GDDR6",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Gaming",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "2x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := rbs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Razer Blade 14 specifications seeded successfully")
	return nil
}
