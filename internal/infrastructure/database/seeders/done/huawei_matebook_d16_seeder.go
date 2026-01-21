package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HuaweiMatebookD16Seeder seeds specifications for Huawei MateBook D16
type HuaweiMatebookD16Seeder struct {
	BaseSeeder
}

// NewHuaweiMatebookD16Seeder creates a new Huawei MateBook D16 seeder
func NewHuaweiMatebookD16Seeder() *HuaweiMatebookD16Seeder {
	return &HuaweiMatebookD16Seeder{
		BaseSeeder: BaseSeeder{name: "Huawei MateBook D16 Specifications"},
	}
}

func (hms *HuaweiMatebookD16Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Huawei":                               "হুয়াওয়ে",
		"MateBook D16":                         "মেটবুক ডি১৬",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"AMD":                                  "এএমডি",
		"Ryzen 5 5600H":                        "রাইজেন ৫ ৫৬০০এইচ",
		"5th Gen":                              "৫ম প্রজন্ম",
		"AMD Platform Controller Hub":          "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Radeon Graphics":                      "রেডিওন গ্রাফিক্স",
		"1920 x 1200 pixels":                   "১৯২০ x ১২০০ পিক্সেল",
		"16 inches":                            "১৬ ইঞ্চি",
		"60 Wh":                                "৬০ ডব্লিউএইচ",
		"Plastic":                              "প্লাস্টিক",
		"Mystic Silver":                        "মিস্টিক সিলভার",
		"1.7 kg":                               "১.৭ কেজি",
		"2022":                                 "২০২২",
		"1 Year International Warranty":        "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                               "৫১২ জিবি",
		"Bluetooth 5.1":                        "ব্লুটুথ ৫.১",
		"8 GB":                                 "৮ জিবি",
		"IPS":                                  "আইপিএস",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                        "ইউএসবি ৩.২ জেন ১",
		"Dedicated":                            "ডেডিকেটেড",
		"512 GB SSD":                           "৫১২ জিবি এসএসডি",
		"AMD Ryzen 5":                          "এএমডি রাইজেন ৫",
		"Single Fan":                           "সিঙ্গেল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Ultraportable":                        "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"IPS, 300 nits":                        "আইপিএস, ৩০০ নিটস",
		"3.3 GHz base / 4.2 GHz boost":         "৩.৩ গিগাহার্টজ বেস / ৪.২ গিগাহার্টজ বুস্ট",
		"16 MB cache":                          "১৬ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"356.9 x 249.1 x 18.4 mm":              "৩৫৬.৯ x ২৪৯.১ x ১৮.৪ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"6":                                    "৬",
		"12":                                   "১২",
		"3200 MHz":                             "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"4 GB":                                 "৪ জিবি",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
		"Huawei Share":                         "হুয়াওয়ে শেয়ার",
	}
}

// Seed implements the Seeder interface for Huawei MateBook D16
func (hms *HuaweiMatebookD16Seeder) Seed(db *gorm.DB) error {
	productSlug := "huawei-matebook-d16"
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
		"Brand":                   "Huawei",
		"Model Name":              "MateBook D16",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 5 5600H",
		"Processor Generation":    "5th Gen",
		"Chipset":                 "AMD Platform Controller Hub",
		"Graphics Card":           "Radeon Graphics",
		"Screen Resolution":       "1920 x 1200 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "60 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Mystic Silver",
		"Product Weight":          "1.7 kg",
		"Release Year":            "2022",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "8 GB",
		"Weight":                  "1.7 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "60 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "AMD Ryzen 5",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "3.3 GHz base / 4.2 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "356.9 x 249.1 x 18.4 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Mystic Silver",
		"Special Features":        "Huawei Share, TPM 2.0",
		"Processor Cores":         "6",
		"Processor Threads":       "12",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "4 GB",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := hms.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Huawei MateBook D16 specifications seeded successfully")
	return nil
}
