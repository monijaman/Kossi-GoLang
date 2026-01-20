package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LenovoThinkpadX1ExtremeGen5Seeder seeds specifications for Lenovo ThinkPad X1 Extreme Gen 5
type LenovoThinkpadX1ExtremeGen5Seeder struct {
	BaseSeeder
}

// NewLenovoThinkpadX1ExtremeGen5Seeder creates a new Lenovo ThinkPad X1 Extreme Gen 5 seeder
func NewLenovoThinkpadX1ExtremeGen5Seeder() *LenovoThinkpadX1ExtremeGen5Seeder {
	return &LenovoThinkpadX1ExtremeGen5Seeder{
		BaseSeeder: BaseSeeder{name: "Lenovo ThinkPad X1 Extreme Gen 5 Specifications"},
	}
}

func (lts *LenovoThinkpadX1ExtremeGen5Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Lenovo":                        "লেনোভো",
		"ThinkPad X1 Extreme Gen 5":     "থিঙ্কপ্যাড এক্স১ এক্সট্রিম জেন ৫",
		"Windows 11 Pro":                "উইন্ডোজ ১১ প্রো",
		"Intel":                         "ইন্টেল",
		"Core i9-12900HX":               "কোর আই৯-১২৯০০এইচএক্স",
		"12th Gen":                      "১২তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 3050 Ti":    "এনভিডিয়া জিফোর্স আরটিএক্স ৩০৫০ টাই",
		"3840 x 2400 pixels":            "৩৮৪০ x ২৪০০ পিক্সেল",
		"16 inches":                     "১৬ ইঞ্চি",
		"90 Wh":                         "৯০ ডব্লিউএইচ",
		"Magnesium Alloy":               "ম্যাগনেসিয়াম অ্যালয়",
		"Black":                         "কালো",
		"1.87 kg":                       "১.৮৭ কেজি",
		"2023":                          "২০২৩",
		"3 Year Onsite Warranty":        "৩ বছর অনসাইট ওয়ারেন্টি",
		"1 TB":                          "১ টিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"32 GB":                         "৩২ জিবি",
		"OLED":                          "ওএলইডি",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                     "ডেডিকেটেড",
		"1 TB SSD":                      "১ টিবি এসএসডি",
		"Intel Core i9":                 "ইন্টেল কোর আই৯",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"3 Years":                       "৩ বছর",
		"High-Performance Laptop":       "হাই-পারফরম্যান্স ল্যাপটপ",
		"Lithium-polymer":               "লিথিয়াম-পলিমার",
		"135W USB-C":                    "১৩৫ওয়াট ইউএসবি-সি",
		"Up to 8 hours":                 "৮ ঘন্টা পর্যন্ত",
		"OLED, 400 nits":                "ওএলইডি, ৪০০ নিটস",
		"2.3 GHz base / 5.0 GHz boost":  "২.৩ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"30 MB cache":                   "৩০ এমবি ক্যাশ",
		"1080p FHD IR":                  "১০৮০পি এফএইচডি আইআর",
		"Dolby Atmos Speakers":          "ডলবি অ্যাটমস স্পিকার",
		"356.7 x 251.6 x 16.9 mm":       "৩৫৬.৭ x ২৫১.৬ x ১৬.৯ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"24":                            "২৪",
		"32":                            "৩২",
		"4800 MHz":                      "৪৮০০ মেগাহার্টজ",
		"Yes (up to 128GB)":             "হ্যাঁ (১২৮জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                "হ্যাঁ (এম.২ স্লট)",
		"4 GB":                          "৪ জিবি",
		"English/Bangla":                "ইংরেজি/বাংলা",
		"Military Grade":                "মিলিটারি গ্রেড",
		"HDMI 2.1":                      "এইচডিএমআই ২.১",
		"2x Thunderbolt 4, 2x USB-A":    "২x থান্ডারবোল্ট ৪, ২x ইউএসবি-এ",
		"Touchscreen":                   "টাচস্ক্রিন",
		"Thunderbolt 4":                 "থান্ডারবোল্ট ৪",
		"Match-on-chip fingerprint":     "ম্যাচ-অন-চিপ ফিঙ্গারপ্রিন্ট",
		"IR camera":                     "আইআর ক্যামেরা",
		"Ethernet (RJ-45)":              "ইথারনেট (আরজে-৪৫)",
	}
}

// Seed implements the Seeder interface for Lenovo ThinkPad X1 Extreme Gen 5
func (lts *LenovoThinkpadX1ExtremeGen5Seeder) Seed(db *gorm.DB) error {
	productSlug := "lenovo-thinkpad-x1-extreme-gen-5"
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
		"Brand":                   "Lenovo",
		"Model Name":              "ThinkPad X1 Extreme Gen 5",
		"Operating System":        "Windows 11 Pro",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i9-12900HX",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "NVIDIA GeForce RTX 3050 Ti",
		"Screen Resolution":       "3840 x 2400 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "90 Wh",
		"Build Material":          "Magnesium Alloy",
		"Color":                   "Black",
		"Product Weight":          "1.87 kg",
		"Release Year":            "2023",
		"Warranty":                "3 Year Onsite Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "32 GB",
		"Weight":                  "1.87 kg",
		"Display Type":            "OLED",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "32 GB",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "90 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "Intel Core i9",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "3 Years",
		"Laptop Type":             "High-Performance Laptop",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "135W USB-C",
		"Standby Time":            "Up to 8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "3840 x 2400 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "2.3 GHz base / 5.0 GHz boost",
		"Clock Feature":           "30 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD IR",
		"Audio Quality":           "Dolby Atmos Speakers",
		"Sensors":                 "Match-on-chip fingerprint",
		"Dimensions":              "356.7 x 251.6 x 16.9 mm",
		"Body Type":               "Magnesium Alloy",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0, IR camera",
		"Processor Cores":         "24",
		"Processor Threads":       "32",
		"RAM Speed":               "4800 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 128GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "4 GB",
		"Display Touch Support":   "Yes",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "Thunderbolt 4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Military Grade",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "2x Thunderbolt 4, 2x USB-A",
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
		if banglaValue, exists := lts.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Lenovo ThinkPad X1 Extreme Gen 5 specifications seeded successfully")
	return nil
}
