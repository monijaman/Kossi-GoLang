package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// TeclastF7PlusSeeder seeds specifications for Teclast F7 Plus
type TeclastF7PlusSeeder struct {
	BaseSeeder
}

// NewTeclastF7PlusSeeder creates a new Teclast F7 Plus seeder
func NewTeclastF7PlusSeeder() *TeclastF7PlusSeeder {
	return &TeclastF7PlusSeeder{
		BaseSeeder: BaseSeeder{name: "Teclast F7 Plus Specifications"},
	}
}

func (tfps *TeclastF7PlusSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Teclast":                      "টেকলাস্ট",
		"F7 Plus":                      "এফ৭ প্লাস",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 7 5700U":                "রাইজেন ৭ ৫৭০০ইউ",
		"5th Gen":                      "৫ম প্রজন্ম",
		"AMD SoC":                      "এএমডি এসওসি",
		"AMD Radeon Graphics":          "এএমডি রেডিওন গ্রাফিক্স",
		"1920 x 1080 pixels":           "১৯২০ x ১০৮০ পিক্সেল",
		"15.6 inches":                  "১৫.৬ ইঞ্চি",
		"6000 mAh":                     "৬০০০ এমএএইচ",
		"Plastic":                      "প্লাস্টিক",
		"Silver":                       "সিলভার",
		"1.8 kg":                       "১.৮ কেজি",
		"2023":                         "২০২৩",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"512 GB":                       "৫১২ জিবি",
		"Bluetooth 5.0":                "ব্লুটুথ ৫.০",
		"16 GB":                        "১৬ জিবি",
		"IPS":                          "আইপিএস",
		"60 Hz":                        "৬০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6":                      "ওয়াই-ফাই ৬",
		"USB 3.0":                      "ইউএসবি ৩.০",
		"Integrated":                   "ইন্টিগ্রেটেড",
		"512 GB SSD":                   "৫১২ জিবি এসএসডি",
		"AMD Ryzen 7":                  "এএমডি রাইজেন ৭",
		"Single Fan":                   "সিঙ্গেল ফ্যান",
		"60":                           "৬০",
		"No":                           "না",
		"1 Year":                       "১ বছর",
		"Ultraportable":                "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"65W USB-C":                    "৬৫ওয়াট ইউএসবি-সি",
		"7-9 hours":                    "৭-৯ ঘন্টা",
		"IPS, 250 nits":                "আইপিএস, ২৫০ নিটস",
		"1.8 GHz base / 4.3 GHz boost": "১.৮ গিগাহার্টজ বেস / ৪.৩ গিগাহার্টজ বুস্ট",
		"16 MB cache":                  "১৬ এমবি ক্যাশ",
		"720p HD":                      "৭২০পি এইচডি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"358 x 242 x 18 mm":            "৩৫৮ x ২৪২ x ১৮ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"8":                            "৮",
		"16":                           "১৬",
		"3200 MHz":                     "৩২০০ মেগাহার্টজ",
		"2":                            "২",
		"Yes (up to 64GB)":             "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":               "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":               "হ্যাঁ (এম.২ স্লট)",
		"Shared":                       "শেয়ার্ড",
		"English":                      "ইংরেজি",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.0":                     "এইচডিএমআই ২.০",
		"2x USB 3.0, 1x USB-C":         "২x ইউএসবি ৩.০, ১x ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for Teclast F7 Plus
func (tfps *TeclastF7PlusSeeder) Seed(db *gorm.DB) error {
	productSlug := "teclast-f7-plus"
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
		"Brand":                   "Teclast",
		"Model Name":              "F7 Plus",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 7 5700U",
		"Processor Generation":    "5th Gen",
		"Chipset":                 "AMD SoC",
		"Graphics Card":           "AMD Radeon Graphics",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "15.6 inches",
		"Battery Capacity":        "6000 mAh",
		"Build Material":          "Plastic",
		"Color":                   "Silver",
		"Product Weight":          "1.8 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.0",
		"RAM":                     "16 GB",
		"Weight":                  "1.8 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "15.6 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6",
		"Usb Type":                "USB 3.0",
		"Battery":                 "6000 mAh",
		"Gpu Type":                "Integrated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "AMD Ryzen 7",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "7-9 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 250 nits",
		"Processor Speed":         "1.8 GHz base / 4.3 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "358 x 242 x 18 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "8",
		"Processor Threads":       "16",
		"RAM Speed":               "3200 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen3",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "2x USB 3.0, 1x USB-C",
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
		if banglaValue, exists := tfps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Teclast F7 Plus specifications seeded successfully")
	return nil
}
