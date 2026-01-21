package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LgGram17Seeder seeds specifications for LG Gram 17
type LgGram17Seeder struct {
	BaseSeeder
}

// NewLgGram17Seeder creates a new LG Gram 17 seeder
func NewLgGram17Seeder() *LgGram17Seeder {
	return &LgGram17Seeder{
		BaseSeeder: BaseSeeder{name: "LG Gram 17 Specifications"},
	}
}

func (lgs *LgGram17Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"LG":                            "এলজি",
		"Gram 17":                       "গ্রাম ১৭",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i7-1360P":                 "কোর আই৭-১৩৬০পি",
		"13th Gen":                      "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                 "ইন্টেল আইরিস এক্সই",
		"2560 x 1600 pixels":            "২৫৬০ x ১৬০০ পিক্সেল",
		"17 inches":                     "১৭ ইঞ্চি",
		"80 Wh":                         "৮০ ডব্লিউএইচ",
		"Magnesium-Lithium Alloy":       "ম্যাগনেসিয়াম-লিথিয়াম অ্যালয়",
		"Silver":                        "সিলভার",
		"1.35 kg":                       "১.৩৫ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.1":                 "ব্লুটুথ ৫.১",
		"16 GB":                         "১৬ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"Intel Core i7":                 "ইন্টেল কোর আই৭",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		"Ultrabook":                     "আল্ট্রাবুক",
		"Lithium-polymer":               "লিথিয়াম-পলিমার",
		"65W USB-C":                     "৬৫ওয়াট ইউএসবি-সি",
		"Up to 22 hours":                "২২ ঘন্টা পর্যন্ত",
		"IPS, 350 nits":                 "আইপিএস, ৩৫০ নিটস",
		"2.2 GHz base / 5.0 GHz boost":  "২.২ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"18 MB cache":                   "১৮ এমবি ক্যাশ",
		"1080p FHD":                     "১০৮০পি এফএইচডি",
		"Stereo Speakers":               "স্টেরিও স্পিকার",
		"380 x 260.3 x 17.4 mm":         "৩৮০ x ২৬০.৩ x ১৭.৪ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"12":                            "১২",
		"16":                            "১৬",
		"5200 MHz":                      "৫২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":              "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                "হ্যাঁ (এম.২ স্লট)",
		"Shared":                        "শেয়ার্ড",
		"English/Bangla":                "ইংরেজি/বাংলা",
		"Standard":                      "স্ট্যান্ডার্ড",
		"HDMI 2.0":                      "এইচডিএমআই ২.০",
		"1x USB-C 3.2 Gen 2, 2x USB-A":  "১x ইউএসবি-সি ৩.২ জেন ২, ২x ইউএসবি-এ",
		"Thunderbolt 4":                 "থান্ডারবোল্ট ৪",
		"Fingerprint sensor":            "ফিঙ্গারপ্রিন্ট সেন্সর",
	}
}

// Seed implements the Seeder interface for LG Gram 17
func (lgs *LgGram17Seeder) Seed(db *gorm.DB) error {
	productSlug := "lg-gram-17"
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
		"Brand":                   "LG",
		"Model Name":              "Gram 17",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-1360P",
		"Processor Generation":    "13th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "2560 x 1600 pixels",
		"Display Size":            "17 inches",
		"Battery Capacity":        "80 Wh",
		"Build Material":          "Magnesium-Lithium Alloy",
		"Color":                   "Silver",
		"Product Weight":          "1.35 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "512 GB",
		"Bluetooth Version":       "Bluetooth 5.1",
		"RAM":                     "16 GB",
		"Weight":                  "1.35 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "17 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6 (802.11ax)",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "80 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "512 GB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultrabook",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "Up to 22 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "IPS, 350 nits",
		"Processor Speed":         "2.2 GHz base / 5.0 GHz boost",
		"Clock Feature":           "18 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "Fingerprint sensor",
		"Dimensions":              "380 x 260.3 x 17.4 mm",
		"Body Type":               "Magnesium-Lithium Alloy",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "12",
		"Processor Threads":       "16",
		"RAM Speed":               "5200 MHz",
		"RAM Slots":               "Soldered",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "Thunderbolt 4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.0",
		"USB Ports":               "1x USB-C 3.2 Gen 2, 2x USB-A",
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
		if banglaValue, exists := lgs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ LG Gram 17 specifications seeded successfully")
	return nil
}
