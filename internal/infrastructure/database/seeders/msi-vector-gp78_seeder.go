package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// MsiVectorGp78Seeder seeds specifications for MSI Vector GP78
type MsiVectorGp78Seeder struct {
	BaseSeeder
}

// NewMsiVectorGp78Seeder creates a new MSI Vector GP78 seeder
func NewMsiVectorGp78Seeder() *MsiVectorGp78Seeder {
	return &MsiVectorGp78Seeder{
		BaseSeeder: BaseSeeder{name: "MSI Vector GP78 Specifications"},
	}
}

func (mvgs *MsiVectorGp78Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"MSI":                           "এমএসআই",
		"Vector GP78":                   "ভেক্টর জিপি৭৮",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i9-13980HX":               "কোর আই৯-১৩৯৮০এইচএক্স",
		"13th Gen":                      "১৩তম প্রজন্ম",
		"Intel HM770":                   "ইন্টেল এইচএমএম৭৭০",
		"NVIDIA GeForce RTX 4090":       "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৯০",
		"2560 x 1440 pixels":            "২৫৬০ x ১৪৪০ পিক্সেল",
		"17 inches":                     "১৭ ইঞ্চি",
		"99.9 Wh":                       "৯৯.৯ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Black":                         "কালো",
		"3.2 kg":                        "৩.২ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"2 TB":                          "২ টেরাবাইট",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"64 GB":                         "৬৪ জিবি",
		"IPS":                           "আইপিএস",
		"240 Hz":                        "২৪০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                     "ডেডিকেটেড",
		"2 TB SSD":                      "২ টেরাবাইট এসএসডি",
		"Intel Core i9":                 "ইন্টেল কোর আই৯",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"240":                           "২৪০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		"High-End Gaming":               "হাই-এন্ড গেমিং",
		"Lithium-ion":                   "লিথিয়াম-আয়ন",
		"330W":                          "৩৩০ওয়াট",
		"4-6 hours":                     "৪-৬ ঘন্টা",
		"IPS, 1000 nits":                "আইপিএস, ১০০০ নিটস",
		"2.0 GHz base / 5.6 GHz boost":  "২.০ গিগাহার্টজ বেস / ৫.৬ গিগাহার্টজ বুস্ট",
		"36 MB cache":                   "৩৬ এমবি ক্যাশ",
		"1080p FHD":                     "১০৮০পি এফএইচডি",
		"Nahimic Audio":                 "নাহিমিক অডিও",
		"397 x 284 x 23 mm":             "৩৯৭ x ২৮৪ x ২৩ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"17":                            "১৭",
		"32":                            "৩২",
		"5600 MHz":                      "৫৬০০ মেগাহার্টজ",
		"Yes (up to 128GB)":             "হ্যাঁ (১২৮জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slots)":               "হ্যাঁ (এম.২ স্লটস)",
		"16 GB GDDR6":                   "১৬ জিবি জিডিডিআর৬",
		"English":                       "ইংরেজি",
		"HDMI 2.1":                      "এইচডিএমআই ২.১",
		"3x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2, 1x Thunderbolt 4": "৩x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২, ১x থান্ডারবোল্ট ৪",
	}
}

// Seed implements the Seeder interface for MSI Vector GP78
func (mvgs *MsiVectorGp78Seeder) Seed(db *gorm.DB) error {
	productSlug := "msi-vector-gp78"
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
		"Brand":                   "MSI",
		"Model Name":              "Vector GP78",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i9-13980HX",
		"Processor Generation":    "13th Gen",
		"Chipset":                 "Intel HM770",
		"Graphics Card":           "NVIDIA GeForce RTX 4090",
		"Screen Resolution":       "2560 x 1440 pixels",
		"Display Size":            "17 inches",
		"Battery Capacity":        "99.9 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "3.2 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "2 TB",
		"Bluetooth Version":       "Bluetooth 5.3",
		"RAM":                     "64 GB",
		"Weight":                  "3.2 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "240 Hz",
		"Screen Size":             "17 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "64 GB",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "99.9 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "2 TB SSD",
		"Cpu Type":                "Intel Core i9",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "240",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "High-End Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "330W",
		"Standby Time":            "4-6 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1440 pixels",
		"Display Characteristics": "IPS, 1000 nits",
		"Processor Speed":         "2.0 GHz base / 5.6 GHz boost",
		"Clock Feature":           "36 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Nahimic Audio",
		"Sensors":                 "No",
		"Dimensions":              "397 x 284 x 23 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "24",
		"Processor Threads":       "32",
		"RAM Speed":               "5600 MHz",
		"RAM Slots":               "4",
		"RAM Expandable":          "Yes (up to 128GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slots)",
		"Graphics VRAM":           "16 GB GDDR6",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "3x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2, 1x Thunderbolt 4",
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
		if banglaValue, exists := mvgs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ MSI Vector GP78 specifications seeded successfully")
	return nil
}
