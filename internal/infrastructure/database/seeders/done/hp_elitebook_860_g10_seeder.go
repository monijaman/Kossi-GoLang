package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HpElitebook860G10Seeder seeds specifications for HP EliteBook 860 G10
type HpElitebook860G10Seeder struct {
	BaseSeeder
}

// NewHpElitebook860G10Seeder creates a new HP EliteBook 860 G10 seeder
func NewHpElitebook860G10Seeder() *HpElitebook860G10Seeder {
	return &HpElitebook860G10Seeder{
		BaseSeeder: BaseSeeder{name: "HP EliteBook 860 G10 Specifications"},
	}
}

func (hebs *HpElitebook860G10Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HP":                            "এইচপি",
		"EliteBook 860 G10":             "এলিটবুক ৮৬০ জি১০",
		"Windows 11 Pro":                "উইন্ডোজ ১১ প্রো",
		"Intel":                         "ইন্টেল",
		"Core i7-1265U":                 "কোর আই৭-১২৬৫ইউ",
		"12th Gen":                      "১২তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                 "ইন্টেল আইরিস এক্সই",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"16 inches":                     "১৬ ইঞ্চি",
		"68 Wh":                         "৬৮ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Silver":                        "সিলভার",
		"1.5 kg":                        "১.৫ কেজি",
		"2023":                          "২০২৩",
		"3 Year International Warranty": "৩ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"16 GB":                         "১৬ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"Intel Core i7":                 "ইন্টেল কোর আই৭",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"3 Year":                        "৩ বছর",
		// New translations for added specs
		"Business":                     "বিজনেস",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"65W USB-C":                    "৬৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                   "৮-১০ ঘন্টা",
		"IPS, 400 nits":                "আইপিএস, ৪০০ নিটস",
		"1.8 GHz base / 4.8 GHz boost": "১.৮ গিগাহার্টজ বেস / ৪.৮ গিগাহার্টজ বুস্ট",
		"12 MB cache":                  "১২ এমবি ক্যাশ",
		"1080p FHD":                    "১০৮০পি এফএইচডি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"356.9 x 248.2 x 16.9 mm":      "৩৫৬.৯ x ২৪৮.২ x ১৬.৯ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"10":                           "১০",
		"12":                           "১২",
		"5200 MHz":                     "৫২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":             "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":               "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":               "হ্যাঁ (এম.২ স্লট)",
		"Shared":                       "শেয়ার্ড",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"HDMI 2.0":                     "এইচডিএমআই ২.০",
		"2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2, 1x Thunderbolt 4": "২x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২, ১x থান্ডারবোল্ট ৪"}
}

// Seed implements the Seeder interface for HP EliteBook 860 G10
func (hebs *HpElitebook860G10Seeder) Seed(db *gorm.DB) error {
	productSlug := "hp-elitebook-860-g10"
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
		"Brand":                310,
		"Model Name":           316,
		"Operating System":     49,
		"Processor Brand":      204,
		"Processor Model":      206,
		"Processor Generation": 205,
		"Chipset":              19,
		"Graphics Card":        200,
		"Screen Resolution":    207,
		"Display Size":         683,
		"Battery Capacity":     11,
		"Build Material":       14,
		"Color":                318,
		"Product Weight":       697,
		"Release Year":         317,
		"Warranty":             323,
		"Storage Capacity":     71,
		"Bluetooth Version":    13,
		"RAM":                  60,
		"Weight":               80,
		"Display Type":         27,
		"Refresh Rate":         61,
		"Screen Size":          66,
		"Backlit Keyboard":     199,
		"Audio Jack":           682,
		"Ram":                  684,
		"Wifi Support":         687,
		"Usb Type":             688,
		"Battery":              689,
		"Gpu Type":             691,
		"Storage":              693,
		"Cpu Type":             696,
		"Cooling Technology":   698,
		"Frequency (Hz)":       704,
		"App Control":          705,
		"Warranty Period":      706,
		// New additions from specs.sql
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
		// Major laptop specs added to database
		"Processor Cores":       786,
		"Processor Threads":     787,
		"RAM Speed":             788,
		"RAM Slots":             789,
		"RAM Expandable":        790,
		"Storage Interface":     791,
		"Storage Expandable":    792,
		"Graphics VRAM":         793,
		"Display Touch Support": 794,
		"Ethernet":              795,
		"Thunderbolt Version":   796,
		"SD Card Reader":        797,
		"Keyboard Language":     798,
		"Build Standard":        799,
		// Additional available keys
		"Touchscreen": 129,
		"HDMI Ports":  167,
		"USB Ports":   173,
	}

	specs := map[string]string{
		"Brand":                "HP",
		"Model Name":           "EliteBook 860 G10",
		"Operating System":     "Windows 11 Pro",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i7-1265U",
		"Processor Generation": "12th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel Iris Xe",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "16 inches",
		"Battery Capacity":     "68 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Silver",
		"Product Weight":       "1.5 kg",
		"Release Year":         "2023",
		"Warranty":             "3 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "16 GB",
		"Weight":               "1.5 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "16 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "68 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "Intel Core i7",
		"Cooling Technology":   "Single Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "3 Year",
		// New specs from specs.sql
		"Laptop Type":             "Business",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "1.8 GHz base / 4.8 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "356.9 x 248.2 x 16.9 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "10",
		"Processor Threads":     "12",
		"RAM Speed":             "5200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 64GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "Yes",
		"Thunderbolt Version":   "Thunderbolt 4",
		"SD Card Reader":        "Yes",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Business",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.0",
		"USB Ports":   "2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2, 1x Thunderbolt 4",
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
		if banglaValue, exists := hebs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ HP EliteBook 860 G10 specifications seeded successfully")
	return nil
}
