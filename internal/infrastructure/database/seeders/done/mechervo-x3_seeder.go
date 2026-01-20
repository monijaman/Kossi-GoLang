package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MechervoX3Seeder seeds specifications for Mechrevo X3
type MechervoX3Seeder struct {
	BaseSeeder
}

// NewMechervoX3Seeder creates a new Mechrevo X3 seeder
func NewMechervoX3Seeder() *MechervoX3Seeder {
	return &MechervoX3Seeder{
		BaseSeeder: BaseSeeder{name: "Mechrevo X3 Specifications"},
	}
}

func (mxs *MechervoX3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Mechrevo":                             "মেচরেভো",
		"X3":                                   "এক্স৩",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"AMD":                                  "এএমডি",
		"Ryzen 5 7640H":                        "রাইজেন ৫ ৭৬৪০এইচ",
		"Zen 4":                                "জেন ৪",
		"AMD X670":                             "এএমডি এক্স৬৭০",
		"NVIDIA GeForce RTX 4050":              "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৫০",
		"1920 x 1200 pixels":                   "১৯২০ x ১২০০ পিক্সেল",
		"14 inches":                            "১৪ ইঞ্চি",
		"65 Wh":                                "৬৫ ডব্লিউএইচ",
		"Aluminum":                             "অ্যালুমিনিয়াম",
		"Black":                                "কালো",
		"1.8 kg":                               "১.৮ কেজি",
		"2024":                                 "২০২৪",
		"1 Year International Warranty":        "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                               "৫১২ জিবি",
		"Bluetooth 5.2":                        "ব্লুটুথ ৫.২",
		"16 GB":                                "১৬ জিবি",
		"IPS":                                  "আইপিএস",
		"120 Hz":                               "১২০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                        "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                            "ডেডিকেটেড",
		"512 GB SSD":                           "৫১২ জিবি এসএসডি",
		"AMD Ryzen 5":                          "এএমডি রাইজেন ৫",
		"Dual Fan":                             "ডুয়াল ফ্যান",
		"120":                                  "১২০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Gaming":                               "গেমিং",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"150W":                                 "১৫০ওয়াট",
		"4-6 hours":                            "৪-৬ ঘন্টা",
		"IPS, 350 nits":                        "আইপিএস, ৩৫০ নিটস",
		"4.3 GHz base / 5.0 GHz boost":         "৪.৩ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"16 MB cache":                          "১৬ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"320 x 220 x 18 mm":                    "৩২০ x ২২০ x ১৮ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"6":                                    "৬",
		"12":                                   "১২",
		"5200 MHz":                             "৫২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"6 GB GDDR6":                           "৬ জিবি জিডিডিআর৬",
		"English":                              "ইংরেজি",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"2x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2": "২x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Mechrevo X3
func (mxs *MechervoX3Seeder) Seed(db *gorm.DB) error {
	productSlug := "mechrevo-x3"
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
		"Brand":                   "Mechrevo",
		"Model Name":              "X3",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 5 7640H",
		"Processor Generation":    "Zen 4",
		"Chipset":                 "AMD X670",
		"Graphics Card":           "NVIDIA GeForce RTX 4050",
		"Screen Resolution":       "1920 x 1200 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "65 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "1.8 kg",
		"Release Year":            "2024",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "16 GB",
		"Weight":                  "1.8 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "120 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "65 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "AMD Ryzen 5",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "120",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "150W",
		"Standby Time":            "4-6 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 350 nits",
		"Processor Speed":         "4.3 GHz base / 5.0 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "320 x 220 x 18 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "6",
		"Processor Threads":       "12",
		"RAM Speed":               "5200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "6 GB GDDR6",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
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
		if banglaValue, exists := mxs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Mechrevo X3 specifications seeded successfully")
	return nil
}
