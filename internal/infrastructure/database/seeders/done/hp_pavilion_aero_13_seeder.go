package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HpPavilionAero13Seeder seeds specifications for HP Pavilion Aero 13
type HpPavilionAero13Seeder struct {
	BaseSeeder
}

// NewHpPavilionAero13Seeder creates a new HP Pavilion Aero 13 seeder
func NewHpPavilionAero13Seeder() *HpPavilionAero13Seeder {
	return &HpPavilionAero13Seeder{
		BaseSeeder: BaseSeeder{name: "HP Pavilion Aero 13 Specifications"},
	}
}

func (hps *HpPavilionAero13Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HP":                            "এইচপি",
		"Pavilion Aero 13":              "প্যাভিলিয়ন এয়ারো ১৩",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"AMD":                           "এএমডি",
		"Ryzen 5 7520U":                 "রাইজেন ৫ ৭৫২০ইউ",
		"7th Gen":                       "৭ম প্রজন্ম",
		"AMD Platform Controller Hub":   "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"AMD Radeon":                    "এএমডি রেডিয়ন",
		"1920 x 1200 pixels":            "১৯২০ x ১২০০ পিক্সেল",
		"13.3 inches":                   "১৩.৩ ইঞ্চি",
		"43 Wh":                         "৪৩ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Natural Silver":                "ন্যাচারাল সিলভার",
		"0.99 kg":                       "০.৯৯ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"AMD Ryzen 5":                   "এএমডি রাইজেন ৫",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Ultralight":                     "আল্ট্রালাইট",
		"Lithium-polymer":                "লিথিয়াম-পলিমার",
		"45W USB-C":                      "৪৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                     "৮-১০ ঘন্টা",
		"IPS, 400 nits":                  "আইপিএস, ৪০০ নিটস",
		"2.3 GHz base / 4.0 GHz boost":   "২.৩ গিগাহার্টজ বেস / ৪.০ গিগাহার্টজ বুস্ট",
		"16 MB cache":                    "১৬ এমবি ক্যাশ",
		"1080p FHD":                      "১০৮০পি এফএইচডি",
		"Stereo Speakers":                "স্টেরিও স্পিকার",
		"297 x 209 x 14.5 mm":            "২৯৭ x ২০৯ x ১৪.৫ মিমি",
		"TPM 2.0":                        "টিপিএম ২.০",
		"6":                              "৬",
		"12":                             "১২",
		"5500 MHz":                       "৫৫০০ মেগাহার্টজ",
		"Yes (up to 32GB)":               "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                 "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":                 "হ্যাঁ (এম.২ স্লট)",
		"Shared":                         "শেয়ার্ড",
		"English/Bangla":                 "ইংরেজি/বাংলা",
		"Premium":                        "প্রিমিয়াম",
		"HDMI 1.4":                       "এইচডিএমআই ১.৪",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২",
	}
}

// Seed implements the Seeder interface for HP Pavilion Aero 13
func (hps *HpPavilionAero13Seeder) Seed(db *gorm.DB) error {
	productSlug := "hp-pavilion-aero-13"
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
		"Model Name":           "Pavilion Aero 13",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "AMD",
		"Processor Model":      "Ryzen 5 7520U",
		"Processor Generation": "7th Gen",
		"Chipset":              "AMD Platform Controller Hub",
		"Graphics Card":        "AMD Radeon",
		"Screen Resolution":    "1920 x 1200 pixels",
		"Display Size":         "13.3 inches",
		"Battery Capacity":     "43 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Natural Silver",
		"Product Weight":       "0.99 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "8 GB",
		"Weight":               "0.99 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "13.3 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "43 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "AMD Ryzen 5",
		"Cooling Technology":   "Single Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Ultralight",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "45W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.3 GHz base / 4.0 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "297 x 209 x 14.5 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Natural Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "6",
		"Processor Threads":     "12",
		"RAM Speed":             "5500 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen3",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 1.4",
		"USB Ports":   "1x USB 3.2 Gen 1, 1x USB-C 3.2",
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
		if banglaValue, exists := hps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ HP Pavilion Aero 13 specifications seeded successfully")
	return nil
}
