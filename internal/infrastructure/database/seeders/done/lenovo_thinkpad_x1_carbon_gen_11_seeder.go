package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LenovoThinkpadX1CarbonGen11Seeder seeds specifications for Lenovo ThinkPad X1 Carbon Gen 11
type LenovoThinkpadX1CarbonGen11Seeder struct {
	BaseSeeder
}

// NewLenovoThinkpadX1CarbonGen11Seeder creates a new Lenovo ThinkPad X1 Carbon Gen 11 seeder
func NewLenovoThinkpadX1CarbonGen11Seeder() *LenovoThinkpadX1CarbonGen11Seeder {
	return &LenovoThinkpadX1CarbonGen11Seeder{
		BaseSeeder: BaseSeeder{name: "Lenovo ThinkPad X1 Carbon Gen 11 Specifications"},
	}
}

func (lts *LenovoThinkpadX1CarbonGen11Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Lenovo":                        "লেনোভো",
		"ThinkPad X1 Carbon Gen 11":     "থিঙ্কপ্যাড এক্স১ কার্বন জেন ১১",
		"Windows 11 Pro":                "উইন্ডোজ ১১ প্রো",
		"Intel":                         "ইন্টেল",
		"Core Ultra 7 165U":             "কোর আল্ট্রা ৭ ১৬৫ইউ",
		"14th Gen":                      "১৪তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Arc Graphics":            "ইন্টেল আর্ক গ্রাফিক্স",
		"2880 x 1800 pixels":            "২৮৮০ x ১৮০০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"57 Wh":                         "৫৭ ডব্লিউএইচ",
		"Carbon Fiber":                  "কার্বন ফাইবার",
		"Black":                         "কালো",
		"1.12 kg":                       "১.১২ কেজি",
		"2024":                          "২০২৪",
		"3 Year Onsite Warranty":        "৩ বছর অনসাইট ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"16 GB":                         "১৬ জিবি",
		"OLED":                          "ওএলইডি",
		"90 Hz":                         "৯০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 4.0":                       "ইউএসবি ৪.০",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"Intel Core Ultra 7":            "ইন্টেল কোর আল্ট্রা ৭",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"90":                            "৯০",
		"No":                            "না",
		"3 Years":                       "৩ বছর",
		"Ultrabook":                     "আল্ট্রাবুক",
		"Lithium-polymer":               "লিথিয়াম-পলিমার",
		"65W USB-C":                     "৬৫ওয়াট ইউএসবি-সি",
		"Up to 12 hours":                "১২ ঘন্টা পর্যন্ত",
		"OLED, 400 nits":                "ওএলইডি, ৪০০ নিটস",
		"1.7 GHz base / 4.9 GHz boost":  "১.৭ গিগাহার্টজ বেস / ৪.৯ গিগাহার্টজ বুস্ট",
		"12 MB cache":                   "১২ এমবি ক্যাশ",
		"1080p FHD IR":                  "১০৮০পি এফএইচডি আইআর",
		"Dolby Atmos Speakers":          "ডলবি অ্যাটমস স্পিকার",
		"312.4 x 220.7 x 14.95 mm":      "৩১২.৪ x ২২০.৭ x ১৪.৯৫ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"14":                            "১৪",
		"16":                            "১৬",
		"7467 MHz":                      "৭৪৬৭ মেগাহার্টজ",
		"Yes (up to 64GB)":              "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                "হ্যাঁ (এম.২ স্লট)",
		"Shared":                        "শেয়ার্ড",
		"English/Bangla":                "ইংরেজি/বাংলা",
		"Military Grade":                "মিলিটারি গ্রেড",
		"HDMI 2.1":                      "এইচডিএমআই ২.১",
		"2x Thunderbolt 4, 1x USB-C":    "২x থান্ডারবোল্ট ৪, ১x ইউএসবি-সি",
		"Touchscreen":                   "টাচস্ক্রিন",
		"Thunderbolt 4":                 "থান্ডারবোল্ট ৪",
		"Match-on-chip fingerprint":     "ম্যাচ-অন-চিপ ফিঙ্গারপ্রিন্ট",
		"IR camera":                     "আইআর ক্যামেরা",
	}
}

// Seed implements the Seeder interface for Lenovo ThinkPad X1 Carbon Gen 11
func (lts *LenovoThinkpadX1CarbonGen11Seeder) Seed(db *gorm.DB) error {
	productSlug := "lenovo-thinkpad-x1-carbon-gen-11"
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
		"Model Name":              "ThinkPad X1 Carbon Gen 11",
		"Operating System":        "Windows 11 Pro",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core Ultra 7 165U",
		"Processor Generation":    "14th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Arc Graphics",
		"Screen Resolution":       "2880 x 1800 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "57 Wh",
		"Build Material":          "Carbon Fiber",
		"Color":                   "Black",
		"Product Weight":          "1.12 kg",
		"Release Year":            "2024",
		"Warranty":                "3 Year Onsite Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.3",
		"RAM":                     "16 GB",
		"Weight":                  "1.12 kg",
		"Display Type":            "OLED",
		"Refresh Rate":            "90 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 4.0",
		"Battery":                 "57 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "Intel Core Ultra 7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "90",
		"App Control":             "No",
		"Warranty Period":         "3 Years",
		"Laptop Type":             "Ultrabook",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "Up to 12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2880 x 1800 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "1.7 GHz base / 4.9 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD IR",
		"Audio Quality":           "Dolby Atmos Speakers",
		"Sensors":                 "Match-on-chip fingerprint",
		"Dimensions":              "312.4 x 220.7 x 14.95 mm",
		"Body Type":               "Carbon Fiber",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0, IR camera",
		"Processor Cores":         "14",
		"Processor Threads":       "16",
		"RAM Speed":               "7467 MHz",
		"RAM Slots":               "Soldered",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "Thunderbolt 4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Military Grade",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "2x Thunderbolt 4, 1x USB-C",
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

	log.Printf("✅ Lenovo ThinkPad X1 Carbon Gen 11 specifications seeded successfully")
	return nil
}
