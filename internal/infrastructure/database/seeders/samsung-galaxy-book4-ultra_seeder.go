package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SamsungGalaxyBook4UltraSeeder seeds specifications for Samsung Galaxy Book4 Ultra
type SamsungGalaxyBook4UltraSeeder struct {
	BaseSeeder
}

// NewSamsungGalaxyBook4UltraSeeder creates a new Samsung Galaxy Book4 Ultra seeder
func NewSamsungGalaxyBook4UltraSeeder() *SamsungGalaxyBook4UltraSeeder {
	return &SamsungGalaxyBook4UltraSeeder{
		BaseSeeder: BaseSeeder{name: "Samsung Galaxy Book4 Ultra Specifications"},
	}
}

func (sgbus *SamsungGalaxyBook4UltraSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Samsung":                              "স্যামসাং",
		"Galaxy Book4 Ultra":                   "গ্যালাক্সি বুক৪ আল্ট্রা",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"Intel":                                "ইন্টেল",
		"Core i7-13700H":                       "কোর আই৭-১৩৭০০এইচ",
		"13th Gen":                             "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub":        "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 4050":              "এনভিডিয়া জিফোর্স আরটিএক্স ৪০৫০",
		"2880 x 1800 pixels":                   "২৮৮০ x ১৮০০ পিক্সেল",
		"16 inches":                            "১৬ ইঞ্চি",
		"76 Wh":                                "৭৬ ডব্লিউএইচ",
		"Aluminum":                             "অ্যালুমিনিয়াম",
		"Silver":                               "সিলভার",
		"1.6 kg":                               "১.৬ কেজি",
		"2023":                                 "২০২৩",
		"1 Year Warranty":                      "১ বছর ওয়ারেন্টি",
		"512 GB":                               "৫১২ জিবি",
		"Bluetooth 5.1":                        "ব্লুটুথ ৫.১",
		"16 GB":                                "১৬ জিবি",
		"AMOLED":                               "এএমওএলইডি",
		"120 Hz":                               "১২০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":                  "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                        "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                            "ডেডিকেটেড",
		"512 GB SSD":                           "৫১২ জিবি এসএসডি",
		"Intel Core i7":                        "ইন্টেল কোর আই৭",
		"Dual Fan":                             "ডুয়াল ফ্যান",
		"120":                                  "১২০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Premium":                              "প্রিমিয়াম",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"10-12 hours":                          "১০-১২ ঘন্টা",
		"AMOLED, 400 nits":                     "এএমওএলইডি, ৪০০ নিটস",
		"2.4 GHz base / 5.0 GHz boost":         "২.৪ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"24 MB cache":                          "২৪ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Stereo Speakers with Dolby Atmos":     "ডলবি অ্যাটমস সহ স্টেরিও স্পিকার",
		"355 x 250 x 12.5 mm":                  "৩৫৫ x ২৫০ x ১২.৫ মিমি",
		"TPM 2.0, S Pen Support":               "টিপিএম ২.০, এস পেন সাপোর্ট",
		"14":                                   "১৪",
		"20":                                   "২০",
		"5200 MHz":                             "৫২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                     "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"6 GB GDDR6":                           "৬ জিবি জিডিডিআর৬",
		"English":                              "ইংরেজি",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 2, 2x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ২, ২x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Samsung Galaxy Book4 Ultra
func (sgbus *SamsungGalaxyBook4UltraSeeder) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-book4-ultra"
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
		"Model Name":              "Galaxy Book4 Ultra",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-13700H",
		"Processor Generation":    "13th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "NVIDIA GeForce RTX 4050",
		"Screen Resolution":       "2880 x 1800 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "76 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Silver",
		"Product Weight":          "1.6 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "16 GB",
		"Weight":                  "1.6 kg",
		"Display Type":            "AMOLED",
		"Refresh Rate":            "120 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "76 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "120",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Premium",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2880 x 1800 pixels",
		"Display Characteristics": "AMOLED, 400 nits",
		"Processor Speed":         "2.4 GHz base / 5.0 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers with Dolby Atmos",
		"Sensors":                 "No",
		"Dimensions":              "355 x 250 x 12.5 mm",
		"Body Type":               "Premium",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0, S Pen Support",
		"Processor Cores":         "14",
		"Processor Threads":       "20",
		"RAM Speed":               "5200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "6 GB GDDR6",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "4",
		"SD Card Reader":          "microSD",
		"Keyboard Language":       "English",
		"Build Standard":          "Premium",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "1x USB 3.2 Gen 2, 2x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := sgbus.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Samsung Galaxy Book4 Ultra specifications seeded successfully")
	return nil
}
