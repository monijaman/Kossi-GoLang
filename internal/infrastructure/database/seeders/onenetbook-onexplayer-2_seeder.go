package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// OnenetbookOnexplayer2Seeder seeds specifications for One-Netbook OneXPlayer 2
type OnenetbookOnexplayer2Seeder struct {
	BaseSeeder
}

// NewOnenetbookOnexplayer2Seeder creates a new One-Netbook OneXPlayer 2 seeder
func NewOnenetbookOnexplayer2Seeder() *OnenetbookOnexplayer2Seeder {
	return &OnenetbookOnexplayer2Seeder{
		BaseSeeder: BaseSeeder{name: "One-Netbook OneXPlayer 2 Specifications"},
	}
}

func (ooss *OnenetbookOnexplayer2Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"One-Netbook":                            "ওয়ান-নেটবুক",
		"OneXPlayer 2":                           "ওয়ানএক্সপ্লেয়ার ২",
		"Windows 11 Home":                        "উইন্ডোজ ১১ হোম",
		"AMD":                                    "এএমডি",
		"Ryzen 7 6800U":                          "রাইজেন ৭ ৬৮০০ইউ",
		"Zen 3+":                                 "জেন ৩+",
		"AMD Renoir":                             "এএমডি রেনয়ার",
		"NVIDIA GeForce RTX 4060":                "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৬০",
		"1920 x 1200 pixels":                     "১৯২০ x ১২০০ পিক্সেল",
		"8.4 inches":                             "৮.৪ ইঞ্চি",
		"47 Wh":                                  "৪৭ ডব্লিউএইচ",
		"Plastic/Metal":                          "প্লাস্টিক/মেটাল",
		"Black":                                  "কালো",
		"0.97 kg":                                "০.৯৭ কেজি",
		"2023":                                   "২০২৩",
		"1 Year International Warranty":          "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                                 "৫১২ জিবি",
		"Bluetooth 5.2":                          "ব্লুটুথ ৫.২",
		"16 GB":                                  "১৬ জিবি",
		"IPS":                                    "আইপিএস",
		"60 Hz":                                  "৬০ হার্জ",
		"Yes":                                    "হ্যাঁ",
		"3.5mm Combo Jack":                       "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                     "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                          "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                              "ডেডিকেটেড",
		"512 GB SSD":                             "৫১২ জিবি এসএসডি",
		"AMD Ryzen 7":                            "এএমডি রাইজেন ৭",
		"Fan Cooling":                            "ফ্যান কুলিং",
		"60":                                     "৬০",
		"No":                                     "না",
		"1 Year":                                 "১ বছর",
		"Handheld Gaming":                        "হ্যান্ডহেল্ড গেমিং",
		"Lithium-ion":                            "লিথিয়াম-আয়ন",
		"65W":                                    "৬৫ওয়াট",
		"4-6 hours":                              "৪-৬ ঘন্টা",
		"IPS, 400 nits":                          "আইপিএস, ৪০০ নিটস",
		"2.7 GHz base / 4.7 GHz boost":           "২.৭ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"16 MB cache":                            "১৬ এমবি ক্যাশ",
		"720p HD":                                "৭২০পি এইচডি",
		"Stereo Speakers":                        "স্টেরিও স্পিকার",
		"210 x 130 x 25 mm":                      "২১০ x ১৩০ x ২৫ মিমি",
		"TPM 2.0":                                "টিপিএম ২.০",
		"8":                                      "৮",
		"16":                                     "১৬",
		"6400 MHz":                               "৬৪০০ মেগাহার্টজ",
		"NVMe PCIe Gen4":                         "এনভিএমই পিসিআই জেন৪",
		"8 GB GDDR6":                             "৮ জিবি জিডিডিআর৬",
		"English":                                "ইংরেজি",
		"1x USB-C 3.2 Gen 2, 1x USB-C 3.2 Gen 2": "১x ইউএসবি-সি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for One-Netbook OneXPlayer 2
func (ooss *OnenetbookOnexplayer2Seeder) Seed(db *gorm.DB) error {
	productSlug := "onenetbook-onexplayer-2"
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
		"Brand":                   "One-Netbook",
		"Model Name":              "OneXPlayer 2",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 7 6800U",
		"Processor Generation":    "Zen 3+",
		"Chipset":                 "AMD Renoir",
		"Graphics Card":           "NVIDIA GeForce RTX 4060",
		"Screen Resolution":       "1920 x 1200 pixels",
		"Display Size":            "8.4 inches",
		"Battery Capacity":        "47 Wh",
		"Build Material":          "Plastic/Metal",
		"Color":                   "Black",
		"Product Weight":          "0.97 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "16 GB",
		"Weight":                  "0.97 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "8.4 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "47 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "AMD Ryzen 7",
		"Cooling Technology":      "Fan Cooling",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Handheld Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W",
		"Standby Time":            "4-6 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.7 GHz base / 4.7 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "210 x 130 x 25 mm",
		"Body Type":               "Plastic/Metal",
		"Cooling System":          "Fan Cooling",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "8",
		"Processor Threads":       "16",
		"RAM Speed":               "6400 MHz",
		"RAM Slots":               "No",
		"RAM Expandable":          "No",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "8 GB GDDR6",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "No",
		"USB Ports":               "1x USB-C 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := ooss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ One-Netbook OneXPlayer 2 specifications seeded successfully")
	return nil
}
