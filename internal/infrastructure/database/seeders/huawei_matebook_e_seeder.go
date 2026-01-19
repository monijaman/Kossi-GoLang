package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HuaweiMatebookESeeder seeds specifications for Huawei MateBook E
type HuaweiMatebookESeeder struct {
	BaseSeeder
}

// NewHuaweiMatebookESeeder creates a new Huawei MateBook E seeder
func NewHuaweiMatebookESeeder() *HuaweiMatebookESeeder {
	return &HuaweiMatebookESeeder{
		BaseSeeder: BaseSeeder{name: "Huawei MateBook E Specifications"},
	}
}

func (hms *HuaweiMatebookESeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Huawei":                        "হুয়াওয়ে",
		"MateBook E":                    "মেটবুক ই",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"AMD":                           "এএমডি",
		"Ryzen 5 3500U":                 "রাইজেন ৫ ৩৫০০ইউ",
		"2nd Gen":                       "২য় প্রজন্ম",
		"AMD Platform Controller Hub":   "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Radeon Vega 8":                 "রেডিওন ভেগা ৮",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"12.6 inches":                   "১২.৬ ইঞ্চি",
		"40 Wh":                         "৪০ ডব্লিউএইচ",
		"Plastic":                       "প্লাস্টিক",
		"Gray":                          "গ্রে",
		"1.3 kg":                        "১.৩ কেজি",
		"2020":                          "২০২০",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"256 GB":                        "২৫৬ জিবি",
		"Bluetooth 5.0":                 "ব্লুটুথ ৫.০",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":            "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"256 GB SSD":                    "২৫৬ জিবি এসএসডি",
		"AMD Ryzen 5":                   "এএমডি রাইজেন ৫",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		"Budget":                        "বাজেট",
		"Lithium-ion":                   "লিথিয়াম-আয়ন",
		"40W USB-C":                     "৪০ওয়াট ইউএসবি-সি",
		"8-10 hours":                    "৮-১০ ঘন্টা",
		"IPS, 250 nits":                 "আইপিএস, ২৫০ নিটস",
		"2.1 GHz base / 3.7 GHz boost":  "২.১ গিগাহার্টজ বেস / ৩.৭ গিগাহার্টজ বুস্ট",
		"4 MB cache":                    "৪ এমবি ক্যাশ",
		"720p HD":                       "৭২০পি এইচডি",
		"Stereo Speakers":               "স্টেরিও স্পিকার",
		"278.8 x 194.1 x 15.9 mm":       "২৭৮.৮ x ১৯৪.১ x ১৫.৯ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"4":                             "৪",
		"8":                             "৮",
		"2400 MHz":                      "২৪০০ মেগাহার্টজ",
		"RAM Expandable":                "No",
		"Storage Interface":             "NVMe PCIe Gen3",
		"Storage Expandable":            "No",
		"Graphics VRAM":                 "Shared",
		"Keyboard Language":             "English/Bangla",
		"Build Standard":                "Standard",
		"HDMI Ports":                    "No",
		"USB Ports":                     "1x USB-C 3.2 Gen 1, 1x USB-A",
		"Huawei Share":                  "হুয়াওয়ে শেয়ার",
	}
}

// Seed implements the Seeder interface for Huawei MateBook E
func (hms *HuaweiMatebookESeeder) Seed(db *gorm.DB) error {
	productSlug := "huawei-matebook-e"
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
		"Model Name":              "MateBook E",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 5 3500U",
		"Processor Generation":    "2nd Gen",
		"Chipset":                 "AMD Platform Controller Hub",
		"Graphics Card":           "Radeon Vega 8",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "12.6 inches",
		"Battery Capacity":        "40 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Gray",
		"Product Weight":          "1.3 kg",
		"Release Year":            "2020",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "256 GB",
		"Bluetooth Version":       "Bluetooth 5.0",
		"RAM":                     "8 GB",
		"Weight":                  "1.3 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "12.6 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "8 GB",
		"Wifi Support":            "Wi-Fi 5 (802.11ac)",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "40 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "256 GB SSD",
		"Cpu Type":                "AMD Ryzen 5",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Budget",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "40W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 250 nits",
		"Processor Speed":         "2.1 GHz base / 3.7 GHz boost",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "278.8 x 194.1 x 15.9 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Gray",
		"Special Features":        "Huawei Share, TPM 2.0",
		"Processor Cores":         "4",
		"Processor Threads":       "8",
		"RAM Speed":               "2400 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "No",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "No",
		"USB Ports":               "1x USB-C 3.2 Gen 1, 1x USB-A",
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

	log.Printf("✅ Huawei MateBook E specifications seeded successfully")
	return nil
}
