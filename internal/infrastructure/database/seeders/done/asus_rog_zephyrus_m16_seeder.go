package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AsusRogZephyrusM16Seeder seeds specifications for Asus ROG Zephyrus M16
type AsusRogZephyrusM16Seeder struct {
	BaseSeeder
}

// NewAsusRogZephyrusM16Seeder creates a new Asus ROG Zephyrus M16 seeder
func NewAsusRogZephyrusM16Seeder() *AsusRogZephyrusM16Seeder {
	return &AsusRogZephyrusM16Seeder{
		BaseSeeder: BaseSeeder{name: "Asus ROG Zephyrus M16 Specifications"},
	}
}

func (arz *AsusRogZephyrusM16Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Asus":                                 "এসুস",
		"ROG Zephyrus M16":                     "আরওজি জেফাইরাস এম১৬",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"Intel":                                "ইন্টেল",
		"Core i9-13900H":                       "কোর আই৯-১৩৯০০এইচ",
		"13th Gen":                             "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub":        "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 4060":              "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৬০",
		"2560 x 1600 pixels":                   "২৫৬০ x ১৬০০ পিক্সেল",
		"16 inches":                            "১৬ ইঞ্চি",
		"90 Wh":                                "৯০ ডব্লিউএইচ",
		"Aluminum":                             "অ্যালুমিনিয়াম",
		"Black":                                "ব্ল্যাক",
		"2.1 kg":                               "২.১ কেজি",
		"2023":                                 "২০২৩",
		"1 Year International Warranty":        "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                                 "১ টেরাবাইট",
		"Bluetooth 5.2":                        "ব্লুটুথ ৫.২",
		"32 GB":                                "৩২ জিবি",
		"IPS":                                  "আইপিএস",
		"165 Hz":                               "১৬৫ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                        "ইউএসবি ৩.২ জেন ১",
		"Dedicated":                            "ডেডিকেটেড",
		"1 TB SSD":                             "১ টেরাবাইট এসএসডি",
		"Intel Core i9":                        "ইন্টেল কোর আই৯",
		"Dual Fan":                             "ডুয়াল ফ্যান",
		"165":                                  "১৬৫",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Gaming":                               "গেমিং",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"240W":                                 "২৪০ওয়াট",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"IPS, 500 nits":                        "আইপিএস, ৫০০ নিটস",
		"2.6 GHz base / 5.4 GHz boost":         "২.৬ গিগাহার্টজ বেস / ৫.৪ গিগাহার্টজ বুস্ট",
		"24 MB cache":                          "২৪ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Dolby Atmos Speakers":                 "ডলবি অ্যাটমস স্পিকার",
		"354 x 248 x 19 mm":                    "৩৫৪ x ২৪৮ x ১৯ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"14":                                   "১৪",
		"20":                                   "২০",
		"5200 MHz":                             "৫২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                     "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"8 GB GDDR6":                           "৮ জিবি জিডিডিআর৬",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "২x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
		"Thunderbolt 4":                        "থান্ডারবোল্ট ৪",
	}
}

// Seed implements the Seeder interface for Asus ROG Zephyrus M16
func (arz *AsusRogZephyrusM16Seeder) Seed(db *gorm.DB) error {
	productSlug := "asus-rog-zephyrus-m16"
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
		"Brand":                   "Asus",
		"Model Name":              "ROG Zephyrus M16",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i9-13900H",
		"Processor Generation":    "13th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "NVIDIA GeForce RTX 4060",
		"Screen Resolution":       "2560 x 1600 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "90 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "2.1 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "32 GB",
		"Weight":                  "2.1 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "165 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "32 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "90 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "Intel Core i9",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "165",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "240W",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "IPS, 500 nits",
		"Processor Speed":         "2.6 GHz base / 5.4 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Dolby Atmos Speakers",
		"Sensors":                 "No",
		"Dimensions":              "354 x 248 x 19 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "14",
		"Processor Threads":       "20",
		"RAM Speed":               "5200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "8 GB GDDR6",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "Thunderbolt 4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
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
		if banglaValue, exists := arz.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Asus ROG Zephyrus M16 specifications seeded successfully")
	return nil
}
