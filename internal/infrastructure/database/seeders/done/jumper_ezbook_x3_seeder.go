package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// JumperEzbookX3Seeder seeds specifications for Jumper EZbook X3
type JumperEzbookX3Seeder struct {
	BaseSeeder
}

// NewJumperEzbookX3Seeder creates a new Jumper EZbook X3 seeder
func NewJumperEzbookX3Seeder() *JumperEzbookX3Seeder {
	return &JumperEzbookX3Seeder{
		BaseSeeder: BaseSeeder{name: "Jumper EZbook X3 Specifications"},
	}
}

func (jex *JumperEzbookX3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Jumper":                       "জাম্পার",
		"EZbook X3":                    "ইজিবুক এক্স৩",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 5 3500U":                "রাইজেন ৫ ৩৫০০ইউ",
		"2nd Gen":                      "২য় প্রজন্ম",
		"AMD Platform Controller Hub":  "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"AMD Radeon Vega 8":            "এএমডি রেডিয়ন ভেগা ৮",
		"1920 x 1080 pixels":           "১৯২০ x ১০৮০ পিক্সেল",
		"13.3 inches":                  "১৩.৩ ইঞ্চি",
		"50 Wh":                        "৫০ ডব্লিউএইচ",
		"Aluminum":                     "অ্যালুমিনিয়াম",
		"Silver":                       "সিলভার",
		"1.2 kg":                       "১.২ কেজি",
		"2023":                         "২০২৩",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"512 GB":                       "৫১২ জিবি",
		"Bluetooth 5.2":                "ব্লুটুথ ৫.২",
		"16 GB":                        "১৬ জিবি",
		"IPS":                          "আইপিএস",
		"60 Hz":                        "৬০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":           "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                "ইউএসবি ৩.২ জেন ১",
		"Integrated":                   "ইন্টিগ্রেটেড",
		"512 GB SSD":                   "৫১২ জিবি এসএসডি",
		"AMD Ryzen 5":                  "এএমডি রাইজেন ৫",
		"Single Fan":                   "সিঙ্গেল ফ্যান",
		"60":                           "৬০",
		"1 Year":                       "১ বছর",
		"Ultraportable":                "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"65W USB-C":                    "৬৫ওয়াট ইউএসবি-সি",
		"10-12 hours":                  "১০-১২ ঘন্টা",
		"IPS, 300 nits":                "আইপিএস, ৩০০ নিটস",
		"2.1 GHz base / 3.7 GHz boost": "২.১ গিগাহার্টজ বেস / ৩.৭ গিগাহার্টজ বুস্ট",
		"4 MB cache":                   "৪ এমবি ক্যাশ",
		"1080p FHD":                    "১০৮০পি এফএইচডি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"310 x 210 x 15 mm":            "৩১০ x ২১০ x ১৫ মিমি",
		"No":                           "না",
		"4":                            "৪",
		"8":                            "৮",
		"3200 MHz":                     "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":             "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":               "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":               "হ্যাঁ (এম.২ স্লট)",
		"Shared":                       "শেয়ার্ড",
		"English":                      "ইংরেজি",
		"Standard":                     "স্ট্যান্ডার্ড",
		"2x USB 3.2 Gen 1, 1x USB-C":   "২x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for Jumper EZbook X3
func (jex *JumperEzbookX3Seeder) Seed(db *gorm.DB) error {
	productSlug := "jumper-ezbook-x3"
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
		"Brand":                   "Jumper",
		"Model Name":              "EZbook X3",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 5 3500U",
		"Processor Generation":    "2nd Gen",
		"Chipset":                 "AMD Platform Controller Hub",
		"Graphics Card":           "AMD Radeon Vega 8",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "13.3 inches",
		"Battery Capacity":        "50 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Silver",
		"Product Weight":          "1.2 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "16 GB",
		"Weight":                  "1.2 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "13.3 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "50 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "AMD Ryzen 5",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.1 GHz base / 3.7 GHz boost",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "310 x 210 x 15 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Silver",
		"Special Features":        "No",
		"Processor Cores":         "4",
		"Processor Threads":       "8",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "No",
		"USB Ports":               "2x USB 3.2 Gen 1, 1x USB-C",
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
		if banglaValue, exists := jex.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Jumper EZbook X3 specifications seeded successfully")
	return nil
}
