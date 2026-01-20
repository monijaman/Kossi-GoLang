package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SamsungGalaxyBook2ProSeeder seeds specifications for Samsung Galaxy Book2 Pro
type SamsungGalaxyBook2ProSeeder struct {
	BaseSeeder
}

// NewSamsungGalaxyBook2ProSeeder creates a new Samsung Galaxy Book2 Pro seeder
func NewSamsungGalaxyBook2ProSeeder() *SamsungGalaxyBook2ProSeeder {
	return &SamsungGalaxyBook2ProSeeder{
		BaseSeeder: BaseSeeder{name: "Samsung Galaxy Book2 Pro Specifications"},
	}
}

func (sgbps *SamsungGalaxyBook2ProSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Samsung":                              "স্যামসাং",
		"Galaxy Book2 Pro":                     "গ্যালাক্সি বুক২ প্রো",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"Intel":                                "ইন্টেল",
		"Core i5-1135G7":                       "কোর আই৫-১১৩৫জি৭",
		"11th Gen":                             "১১তম প্রজন্ম",
		"Intel Platform Controller Hub":        "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                        "ইন্টেল আইরিস এক্সই",
		"1920 x 1080 pixels":                   "১৯২০ x ১০৮০ পিক্সেল",
		"13.3 inches":                          "১৩.৩ ইঞ্চি",
		"63 Wh":                                "৬৩ ডব্লিউএইচ",
		"Aluminum":                             "অ্যালুমিনিয়াম",
		"Graphite":                             "গ্রাফাইট",
		"1.1 kg":                               "১.১ কেজি",
		"2021":                                 "২০২১",
		"1 Year Warranty":                      "১ বছর ওয়ারেন্টি",
		"256 GB":                               "২৫৬ জিবি",
		"Bluetooth 5.1":                        "ব্লুটুথ ৫.১",
		"8 GB":                                 "৮ জিবি",
		"AMOLED":                               "এএমওএলইডি",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                        "ইউএসবি ৩.২ জেন ২",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"256 GB SSD":                           "২৫৬ জিবি এসএসডি",
		"Intel Core i5":                        "ইন্টেল কোর আই৫",
		"Single Fan":                           "সিঙ্গেল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Premium":                              "প্রিমিয়াম",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"12-14 hours":                          "১২-১৪ ঘন্টা",
		"AMOLED, 400 nits":                     "এএমওএলইডি, ৪০০ নিটস",
		"2.4 GHz base / 4.2 GHz boost":         "২.৪ গিগাহার্টজ বেস / ৪.২ গিগাহার্টজ বুস্ট",
		"8 MB cache":                           "৮ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers with Dolby Atmos":     "ডলবি অ্যাটমস সহ স্টেরিও স্পিকার",
		"304 x 202 x 11.2 mm":                  "৩০৪ x ২০২ x ১১.২ মিমি",
		"TPM 2.0, S Pen Support":               "টিপিএম ২.০, এস পেন সাপোর্ট",
		"4":                                    "৪",
		"8":                                    "৮",
		"4266 MHz":                             "৪২৬৬ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Shared":                               "শেয়ার্ড",
		"English":                              "ইংরেজি",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Samsung Galaxy Book2 Pro
func (sgbps *SamsungGalaxyBook2ProSeeder) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-book2-pro"
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
		"Model Name":              "Galaxy Book2 Pro",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i5-1135G7",
		"Processor Generation":    "11th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "13.3 inches",
		"Battery Capacity":        "63 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Graphite",
		"Product Weight":          "1.1 kg",
		"Release Year":            "2021",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "256 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "8 GB",
		"Weight":                  "1.1 kg",
		"Display Type":            "AMOLED",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "13.3 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "63 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "256 GB SSD",
		"Cpu Type":                "Intel Core i5",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Premium",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "12-14 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "AMOLED, 400 nits",
		"Processor Speed":         "2.4 GHz base / 4.2 GHz boost",
		"Clock Feature":           "8 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers with Dolby Atmos",
		"Sensors":                 "No",
		"Dimensions":              "304 x 202 x 11.2 mm",
		"Body Type":               "Premium",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Graphite",
		"Special Features":        "TPM 2.0, S Pen Support",
		"Processor Cores":         "4",
		"Processor Threads":       "8",
		"RAM Speed":               "4266 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "4",
		"SD Card Reader":          "microSD",
		"Keyboard Language":       "English",
		"Build Standard":          "Premium",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := sgbps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Samsung Galaxy Book2 Pro specifications seeded successfully")
	return nil
}
