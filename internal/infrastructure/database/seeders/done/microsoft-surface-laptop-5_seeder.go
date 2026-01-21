package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MicrosoftSurfaceLaptop5Seeder seeds specifications for Microsoft Surface Laptop 5
type MicrosoftSurfaceLaptop5Seeder struct {
	BaseSeeder
}

// NewMicrosoftSurfaceLaptop5Seeder creates a new Microsoft Surface Laptop 5 seeder
func NewMicrosoftSurfaceLaptop5Seeder() *MicrosoftSurfaceLaptop5Seeder {
	return &MicrosoftSurfaceLaptop5Seeder{
		BaseSeeder: BaseSeeder{name: "Microsoft Surface Laptop 5 Specifications"},
	}
}

func (mss *MicrosoftSurfaceLaptop5Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Microsoft":                              "মাইক্রোসফ্ট",
		"Surface Laptop 5":                       "সারফেস ল্যাপটপ ৫",
		"Windows 11 Home":                        "উইন্ডোজ ১১ হোম",
		"Intel":                                  "ইন্টেল",
		"Core i7-1265U":                          "কোর আই৭-১২৬৫ইউ",
		"12th Gen":                               "১২তম প্রজন্ম",
		"Intel Platform Controller Hub":          "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                          "ইন্টেল আইরিস এক্সই",
		"2256 x 1504 pixels":                     "২২৫৬ x ১৫০৪ পিক্সেল",
		"13.5 inches":                            "১৩.৫ ইঞ্চি",
		"47.5 Wh":                                "৪৭.৫ ডব্লিউএইচ",
		"Aluminum":                               "অ্যালুমিনিয়াম",
		"Platinum":                               "প্লাটিনাম",
		"1.27 kg":                                "১.২৭ কেজি",
		"2022":                                   "২০২২",
		"1 Year International Warranty":          "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                                 "৫১২ জিবি",
		"Bluetooth 5.1":                          "ব্লুটুথ ৫.১",
		"16 GB":                                  "১৬ জিবি",
		"IPS":                                    "আইপিএস",
		"60 Hz":                                  "৬০ হার্জ",
		"Yes":                                    "হ্যাঁ",
		"3.5mm Combo Jack":                       "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                     "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                          "ইউএসবি ৩.২ জেন ২",
		"Integrated":                             "ইন্টিগ্রেটেড",
		"512 GB SSD":                             "৫১২ জিবি এসএসডি",
		"Intel Core i7":                          "ইন্টেল কোর আই৭",
		"Single Fan":                             "সিঙ্গেল ফ্যান",
		"60":                                     "৬০",
		"No":                                     "না",
		"1 Year":                                 "১ বছর",
		"Ultrabook":                              "আল্ট্রাবুক",
		"Lithium-ion":                            "লিথিয়াম-আয়ন",
		"65W Surface Connect":                    "৬৫ওয়াট সারফেস কানেক্ট",
		"18 hours":                               "১৮ ঘন্টা",
		"IPS, 400 nits":                          "আইপিএস, ৪০০ নিটস",
		"1.8 GHz base / 4.8 GHz boost":           "১.৮ গিগাহার্টজ বেস / ৪.৮ গিগাহার্টজ বুস্ট",
		"12 MB cache":                            "১২ এমবি ক্যাশ",
		"720p HD":                                "৭২০পি এইচডি",
		"Omnisonic Speakers":                     "ওমনিসনিক স্পিকার",
		"308 x 223 x 14.5 mm":                    "৩০৮ x ২২৩ x ১৪.৫ মিমি",
		"TPM 2.0":                                "টিপিএম ২.০",
		"10":                                     "১০",
		"12":                                     "১২",
		"4266 MHz":                               "৪২৬৬ মেগাহার্টজ",
		"NVMe PCIe Gen4":                         "এনভিএমই পিসিআই জেন৪",
		"Shared":                                 "শেয়ার্ড",
		"English":                                "ইংরেজি",
		"1x USB-C 3.2 Gen 2, 1x Surface Connect": "১x ইউএসবি-সি ৩.২ জেন ২, ১x সারফেস কানেক্ট",
	}
}

// Seed implements the Seeder interface for Microsoft Surface Laptop 5
func (mss *MicrosoftSurfaceLaptop5Seeder) Seed(db *gorm.DB) error {
	productSlug := "microsoft-surface-laptop-5"
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
		"Model Name":              "Surface Laptop 5",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-1265U",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "2256 x 1504 pixels",
		"Display Size":            "13.5 inches",
		"Battery Capacity":        "47.5 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Platinum",
		"Product Weight":          "1.27 kg",
		"Release Year":            "2022",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "16 GB",
		"Weight":                  "1.27 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "13.5 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "47.5 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultrabook",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W Surface Connect",
		"Standby Time":            "18 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2256 x 1504 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "1.8 GHz base / 4.8 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Omnisonic Speakers",
		"Sensors":                 "No",
		"Dimensions":              "308 x 223 x 14.5 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Platinum",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "10",
		"Processor Threads":       "12",
		"RAM Speed":               "4266 MHz",
		"RAM Slots":               "No",
		"RAM Expandable":          "No",
		"Storage Interface":       "NVMe PCIe Gen4",
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
		"USB Ports":               "1x USB-C 3.2 Gen 2, 1x Surface Connect",
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
		if banglaValue, exists := mss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Microsoft Surface Laptop 5 specifications seeded successfully")
	return nil
}
