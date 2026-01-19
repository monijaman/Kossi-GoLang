package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HonorMagicbookPro16Seeder seeds specifications for Honor MagicBook Pro 16
type HonorMagicbookPro16Seeder struct {
	BaseSeeder
}

// NewHonorMagicbookPro16Seeder creates a new Honor MagicBook Pro 16 seeder
func NewHonorMagicbookPro16Seeder() *HonorMagicbookPro16Seeder {
	return &HonorMagicbookPro16Seeder{
		BaseSeeder: BaseSeeder{name: "Honor MagicBook Pro 16 Specifications"},
	}
}

func (hms *HonorMagicbookPro16Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Honor":                         "অনার",
		"MagicBook Pro 16":              "ম্যাজিকবুক প্রো ১৬",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i7-1260P":                 "কোর আই৭-১২৬০পি",
		"12th Gen":                      "১২তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                 "ইন্টেল আইরিস এক্সই",
		"2560 x 1600 pixels":            "২৫৬০ x ১৬০০ পিক্সেল",
		"16 inches":                     "১৬ ইঞ্চি",
		"75 Wh":                         "৭৫ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Titanium Gray":                 "টাইটানিয়াম গ্রে",
		"1.7 kg":                        "১.৭ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                          "১ টিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"16 GB":                         "১৬ জিবি",
		"OLED":                          "ওএলইডি",
		"90 Hz":                         "৯০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"USB-C":                         "ইউএসবি-সি",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"1 TB SSD":                      "১ টিবি এসএসডি",
		"Intel Core i7":                 "ইন্টেল কোর আই৭",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"90":                            "৯০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		"Ultraportable":                 "আল্ট্রাপোর্টেবল",
		"Lithium-polymer":               "লিথিয়াম-পলিমার",
		"100W USB-C":                    "১০০ওয়াট ইউএসবি-সি",
		"10-12 hours":                   "১০-১২ ঘন্টা",
		"OLED, 400 nits":                "ওএলইডি, ৪০০ নিটস",
		"2.1 GHz base / 4.7 GHz boost":  "২.১ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"18 MB cache":                   "১৮ এমবি ক্যাশ",
		"1080p HD":                      "১০৮০পি এইচডি",
		"Quad Speakers":                 "কোয়াড স্পিকার",
		"356.5 x 249.5 x 15.9 mm":       "৩৫৬.৫ x ২৪৯.৫ x ১৫.৯ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"12":                            "১২",
		"16":                            "১৬",
		"4800 MHz":                      "৪৮০০ মেগাহার্টজ",
		"Yes (up to 64GB)":              "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                "হ্যাঁ (এম.২ স্লট)",
		"Shared":                        "শেয়ার্ড",
		"English/Bangla":                "ইংরেজি/বাংলা",
		"Premium":                       "প্রিমিয়াম",
		"HDMI 2.1":                      "এইচডিএমআই ২.১",
		"2x USB-C 3.2 Gen 2, 1x USB-A":  "২x ইউএসবি-সি ৩.২ জেন ২, ১x ইউএসবি-এ"}
}

// Seed implements the Seeder interface for Honor MagicBook Pro 16
func (hms *HonorMagicbookPro16Seeder) Seed(db *gorm.DB) error {
	productSlug := "honor-magicbook-pro-16"
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
		"Brand":                   "Honor",
		"Model Name":              "MagicBook Pro 16",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-1260P",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "2560 x 1600 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "75 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Titanium Gray",
		"Product Weight":          "1.7 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "16 GB",
		"Weight":                  "1.7 kg",
		"Display Type":            "OLED",
		"Refresh Rate":            "90 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "USB-C",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "75 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "90",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "100W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "2.1 GHz base / 4.7 GHz boost",
		"Clock Feature":           "18 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p HD",
		"Audio Quality":           "Quad Speakers",
		"Sensors":                 "No",
		"Dimensions":              "356.5 x 249.5 x 15.9 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Titanium Gray",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "12",
		"Processor Threads":       "16",
		"RAM Speed":               "4800 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Premium",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "2x USB-C 3.2 Gen 2, 1x USB-A",
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

	log.Printf("✅ Honor MagicBook Pro 16 specifications seeded successfully")
	return nil
}
