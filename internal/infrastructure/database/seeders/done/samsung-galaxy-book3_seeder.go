package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// SamsungGalaxyBook3Seeder seeds specifications for Samsung Galaxy Book3
type SamsungGalaxyBook3Seeder struct {
	BaseSeeder
}

// NewSamsungGalaxyBook3Seeder creates a new Samsung Galaxy Book3 seeder
func NewSamsungGalaxyBook3Seeder() *SamsungGalaxyBook3Seeder {
	return &SamsungGalaxyBook3Seeder{
		BaseSeeder: BaseSeeder{name: "Samsung Galaxy Book3 Specifications"},
	}
}

func (sgbs *SamsungGalaxyBook3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Samsung":                       "স্যামসাং",
		"Galaxy Book3":                  "গ্যালাক্সি বুক৩",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i3-1215U":                 "কোর আই৩-১২১৫ইউ",
		"12th Gen":                      "১২তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD Graphics":            "ইন্টেল ইউএইচডি গ্রাফিক্স",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"15.6 inches":                   "১৫.৬ ইঞ্চি",
		"54 Wh":                         "৫৪ ডব্লিউএইচ",
		"Plastic":                       "প্লাস্টিক",
		"Silver":                        "সিলভার",
		"1.6 kg":                        "১.৬ কেজি",
		"2022":                          "২০২২",
		"1 Year Warranty":               "১ বছর ওয়ারেন্টি",
		"256 GB":                        "২৫৬ জিবি",
		"Bluetooth 5.1":                 "ব্লুটুথ ৫.১",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"256 GB SSD":                    "২৫৬ জিবি এসএসডি",
		"Intel Core i3":                 "ইন্টেল কোর আই৩",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		"Budget":                        "বাজেট",
		"Lithium-ion":                   "লিথিয়াম-আয়ন",
		"45W USB-C":                     "৪৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                    "৮-১০ ঘন্টা",
		"IPS, 250 nits":                 "আইপিএস, ২৫০ নিটস",
		"0.9 GHz base / 4.4 GHz boost":  "০.৯ গিগাহার্টজ বেস / ৪.৪ গিগাহার্টজ বুস্ট",
		"10 MB cache":                   "১০ এমবি ক্যাশ",
		"720p HD":                       "৭২০পি এইচডি",
		"Stereo Speakers":               "স্টেরিও স্পিকার",
		"356 x 229 x 15.4 mm":           "৩৫৬ x ২২৯ x ১৫.৪ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"6":                             "৬",
		"8":                             "৮",
		"3200 MHz":                      "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":              "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                "এনভিএমই পিসিআই জেন৩",

		"Shared":  "শেয়ার্ড",
		"English": "ইংরেজি",

		"HDMI 1.4":                             "এইচডিএমআই ১.৪",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 1": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ১",
	}
}

// Seed implements the Seeder interface for Samsung Galaxy Book3
func (sgbs *SamsungGalaxyBook3Seeder) Seed(db *gorm.DB) error {
	productSlug := "samsung-galaxy-book3"
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
		"Model Name":              "Galaxy Book3",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i3-1215U",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel UHD Graphics",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "15.6 inches",
		"Battery Capacity":        "54 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Silver",
		"Product Weight":          "1.6 kg",
		"Release Year":            "2022",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "256 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "8 GB",
		"Weight":                  "1.6 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "15.6 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "54 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "256 GB SSD",
		"Cpu Type":                "Intel Core i3",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Budget",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 250 nits",
		"Processor Speed":         "0.9 GHz base / 4.4 GHz boost",
		"Clock Feature":           "10 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "356 x 229 x 15.4 mm",
		"Body Type":               "Budget",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "6",
		"Processor Threads":       "8",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "microSD",
		"Keyboard Language":       "English",
		"Build Standard":          "Budget",
		"Touchscreen":             "No",
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
		if banglaValue, exists := sgbs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Samsung Galaxy Book3 specifications seeded successfully")
	return nil
}
