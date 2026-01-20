package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// System76LemurProSeeder seeds specifications for System76 Lemur Pro
type System76LemurProSeeder struct {
	BaseSeeder
}

// NewSystem76LemurProSeeder creates a new System76 Lemur Pro seeder
func NewSystem76LemurProSeeder() *System76LemurProSeeder {
	return &System76LemurProSeeder{
		BaseSeeder: BaseSeeder{name: "System76 Lemur Pro Specifications"},
	}
}

func (slps *System76LemurProSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"System76":                             "সিস্টেম৭৬",
		"Lemur Pro":                            "লেমুর প্রো",
		"Pop!_OS 24.04 LTS":                    "পপ!_ওএস ২৪.০৪ এলটিএস",
		"Intel":                                "ইন্টেল",
		"Core Ultra 7 155U":                    "কোর আল্ট্রা ৭ ১৫৫ইউ",
		"14th Gen":                             "১৪তম প্রজন্ম",
		"Intel Iris Xe":                        "ইন্টেল আইরিস এক্সই",
		"1800 x 1200 pixels":                   "১৮০০ x ১২০০ পিক্সেল",
		"14 inches":                            "১৪ ইঞ্চি",
		"50 Wh":                                "৫০ ডব্লিউএইচ",
		"Aluminum":                             "অ্যালুমিনিয়াম",
		"Gray":                                 "গ্রে",
		"2.2 kg":                               "২.২ কেজি",
		"2024":                                 "২০২৪",
		"1 Year Warranty":                      "১ বছর ওয়ারেন্টি",
		"1 TB":                                 "১ টিবি",
		"Bluetooth 5.3":                        "ব্লুটুথ ৫.৩",
		"8 GB":                                 "৮ জিবি",
		"IPS":                                  "আইপিএস",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":                  "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                        "ইউএসবি ৩.২ জেন ২",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"1 TB NVMe SSD":                        "১ টিবি এনভিএমই এসএসডি",
		"Intel Core Ultra 7":                   "ইন্টেল কোর আল্ট্রা ৭",
		"Single Fan":                           "সিঙ্গেল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Ultraportable":                        "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"Up to 14 hours":                       "১৪ ঘন্টা পর্যন্ত",
		"IPS, 400 nits":                        "আইপিএস, ৪০০ নিটস",
		"0.8 GHz base / 4.8 GHz boost":         "০.৮ গিগাহার্টজ বেস / ৪.৮ গিগাহার্টজ বুস্ট",
		"12 MB cache":                          "১২ এমবি ক্যাশ",
		"1080p HD":                             "১০৮০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"TPM 2.0":                              "টিপিএম ২.০",
		"12":                                   "১২",
		"14":                                   "১৪",
		"5600 MHz":                             "৫৬০০ মেগাহার্টজ",
		"Yes (up to 56GB)":                     "হ্যাঁ (৫৬জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
		"Touchscreen":                          "না",
		"HDMI Ports":                           "HDMI 2.0",
		"USB Ports":                            "1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
	}
}

// Seed implements the Seeder interface for System76 Lemur Pro
func (slps *System76LemurProSeeder) Seed(db *gorm.DB) error {
	productSlug := "system76-lemur-pro"
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
		"Display Type":            27,
		"Refresh Rate":            61,
		"Backlit Keyboard":        199,
		"Audio Jack":              682,
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
		"Brand":                   "System76",
		"Model Name":              "Lemur Pro",
		"Operating System":        "Pop!_OS 24.04 LTS",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core Ultra 7 155U",
		"Processor Generation":    "14th Gen",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "1800 x 1200 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "50 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Gray",
		"Product Weight":          "2.2 kg",
		"Release Year":            "2024",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.3",
		"RAM":                     "8 GB",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "50 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "1 TB NVMe SSD",
		"Cpu Type":                "Intel Core Ultra 7",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "Up to 14 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1800 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "0.8 GHz base / 4.8 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Gray",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "12",
		"Processor Threads":       "14",
		"RAM Speed":               "5600 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 56GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := slps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ System76 Lemur Pro specifications seeded successfully")
	return nil
}
