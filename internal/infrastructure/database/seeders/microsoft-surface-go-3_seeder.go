package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MicrosoftSurfaceGo3Seeder seeds specifications for Microsoft Surface Go 3
type MicrosoftSurfaceGo3Seeder struct {
	BaseSeeder
}

// NewMicrosoftSurfaceGo3Seeder creates a new Microsoft Surface Go 3 seeder
func NewMicrosoftSurfaceGo3Seeder() *MicrosoftSurfaceGo3Seeder {
	return &MicrosoftSurfaceGo3Seeder{
		BaseSeeder: BaseSeeder{name: "Microsoft Surface Go 3 Specifications"},
	}
}

func (msgs *MicrosoftSurfaceGo3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Microsoft":                              "মাইক্রোসফ্ট",
		"Surface Go 3":                           "সারফেস গো ৩",
		"Windows 11 Home":                        "উইন্ডোজ ১১ হোম",
		"Intel":                                  "ইন্টেল",
		"Core i3-10100Y":                         "কোর আই৩-১০১০০ওয়াই",
		"10th Gen":                               "১০তম প্রজন্ম",
		"Intel Platform Controller Hub":          "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD Graphics":                     "ইন্টেল ইউএইচডি গ্রাফিক্স",
		"1920 x 1280 pixels":                     "১৯২০ x ১২৮০ পিক্সেল",
		"10.5 inches":                            "১০.৫ ইঞ্চি",
		"26.8 Wh":                                "২৬.৮ ডব্লিউএইচ",
		"Aluminum":                               "অ্যালুমিনিয়াম",
		"Platinum":                               "প্লাটিনাম",
		"0.54 kg":                                "০.৫৪ কেজি",
		"2021":                                   "২০২১",
		"1 Year International Warranty":          "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"128 GB":                                 "১২৮ জিবি",
		"Bluetooth 5.0":                          "ব্লুটুথ ৫.০",
		"8 GB":                                   "৮ জিবি",
		"IPS":                                    "আইপিএস",
		"60 Hz":                                  "৬০ হার্জ",
		"Yes":                                    "হ্যাঁ",
		"3.5mm Combo Jack":                       "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                     "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                          "ইউএসবি ৩.২ জেন ১",
		"Integrated":                             "ইন্টিগ্রেটেড",
		"128 GB SSD":                             "১২৮ জিবি এসএসডি",
		"Intel Core i3":                          "ইন্টেল কোর আই৩",
		"Passive":                                "প্যাসিভ",
		"60":                                     "৬০",
		"No":                                     "না",
		"1 Year":                                 "১ বছর",
		"2-in-1 Tablet":                          "২-ইন-১ ট্যাবলেট",
		"Lithium-ion":                            "লিথিয়াম-আয়ন",
		"39W Surface Connect":                    "৩৯ওয়াট সারফেস কানেক্ট",
		"13.5 hours":                             "১৩.৫ ঘন্টা",
		"IPS, 400 nits":                          "আইপিএস, ৪০০ নিটস",
		"1.3 GHz base / 3.9 GHz boost":           "১.৩ গিগাহার্টজ বেস / ৩.৯ গিগাহার্টজ বুস্ট",
		"6 MB cache":                             "৬ এমবি ক্যাশ",
		"720p HD":                                "৭২০পি এইচডি",
		"Stereo Speakers":                        "স্টেরিও স্পিকার",
		"245 x 175 x 8.3 mm":                     "২৪৫ x ১৭৫ x ৮.৩ মিমি",
		"TPM 2.0":                                "টিপিএম ২.০",
		"2":                                      "২",
		"4":                                      "৪",
		"3200 MHz":                               "৩২০০ মেগাহার্টজ",
		"NVMe PCIe Gen3":                         "এনভিএমই পিসিআই জেন৩",
		"Shared":                                 "শেয়ার্ড",
		"1x USB-C 3.2 Gen 1, 1x Surface Connect": "১x ইউএসবি-সি ৩.২ জেন ১, ১x সারফেস কানেক্ট",
	}
}

// Seed implements the Seeder interface for Microsoft Surface Go 3
func (msgs *MicrosoftSurfaceGo3Seeder) Seed(db *gorm.DB) error {
	productSlug := "microsoft-surface-go-3"
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
		"Brand":                   "Microsoft",
		"Model Name":              "Surface Go 3",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i3-10100Y",
		"Processor Generation":    "10th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel UHD Graphics",
		"Screen Resolution":       "1920 x 1280 pixels",
		"Display Size":            "10.5 inches",
		"Battery Capacity":        "26.8 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Platinum",
		"Product Weight":          "0.54 kg",
		"Release Year":            "2021",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "128 GB",
		"Bluetooth Version":       "Bluetooth 5.0",
		"RAM":                     "8 GB",
		"Weight":                  "0.54 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "10.5 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "26.8 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "128 GB SSD",
		"Cpu Type":                "Intel Core i3",
		"Cooling Technology":      "Passive",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "2-in-1 Tablet",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "39W Surface Connect",
		"Standby Time":            "13.5 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1280 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "1.3 GHz base / 3.9 GHz boost",
		"Clock Feature":           "6 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "245 x 175 x 8.3 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Passive",
		"Available Colors":        "Platinum",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "2",
		"Processor Threads":       "4",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "No",
		"RAM Expandable":          "No",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "No",
		"USB Ports":               "1x USB-C 3.2 Gen 1, 1x Surface Connect",
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
		if banglaValue, exists := msgs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Microsoft Surface Go 3 specifications seeded successfully")
	return nil
}
