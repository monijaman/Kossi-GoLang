package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HpSpectreX36016Seeder seeds specifications for HP Spectre x360 16
type HpSpectreX36016Seeder struct {
	BaseSeeder
}

// NewHpSpectreX36016Seeder creates a new HP Spectre x360 16 seeder
func NewHpSpectreX36016Seeder() *HpSpectreX36016Seeder {
	return &HpSpectreX36016Seeder{
		BaseSeeder: BaseSeeder{name: "HP Spectre x360 16 Specifications"},
	}
}

func (hss *HpSpectreX36016Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HP":                             "এইচপি",
		"Spectre x360 16":                "স্পেক্টর এক্স৩৬০ ১৬",
		"Windows 11 Home":                "উইন্ডোজ ১১ হোম",
		"Intel":                          "ইন্টেল",
		"Core i7-12700H":                 "কোর আই৭-১২৭০০এইচ",
		"12th Gen":                       "১২তম প্রজন্ম",
		"Intel Platform Controller Hub":  "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                  "ইন্টেল আইরিস এক্সই",
		"2560 x 1600 pixels":             "২৫৬০ x ১৬০০ পিক্সেল",
		"16 inches":                      "১৬ ইঞ্চি",
		"96 Wh":                          "৯৬ ডব্লিউএইচ",
		"Aluminum":                       "অ্যালুমিনিয়াম",
		"Nocturne Blue":                  "নকটার্ন ব্লু",
		"2.05 kg":                        "২.০৫ কেজি",
		"2023":                           "২০২৩",
		"1 Year International Warranty":  "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                           "১ টিবি",
		"Bluetooth 5.3":                  "ব্লুটুথ ৫.৩",
		"32 GB":                          "৩২ জিবি",
		"OLED":                           "ওএলইডি",
		"120 Hz":                         "১২০ হার্জ",
		"Yes":                            "হ্যাঁ",
		"USB-C":                          "ইউএসবি-সি",
		"Wi-Fi 6E (802.11ax)":            "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 4.0 Gen 3":                  "ইউএসবি ৪.০ জেন ৩",
		"Integrated":                     "ইন্টিগ্রেটেড",
		"1 TB SSD":                       "১ টিবি এসএসডি",
		"Intel Core i7":                  "ইন্টেল কোর আই৭",
		"Dual Fan":                       "ডুয়াল ফ্যান",
		"120":                            "১২০",
		"No":                             "না",
		"1 Year":                         "১ বছর",
		"2-in-1 Convertible":             "২-ইন-১ কনভার্টিবল",
		"Lithium-polymer":                "লিথিয়াম-পলিমার",
		"140W USB-C":                     "১৪০ওয়াট ইউএসবি-সি",
		"12-14 hours":                    "১২-১৪ ঘন্টা",
		"OLED, 400 nits":                 "ওএলইডি, ৪০০ নিটস",
		"2.3 GHz base / 4.7 GHz boost":   "২.৩ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"24 MB cache":                    "২৪ এমবি ক্যাশ",
		"5MP IR":                         "৫এমপি আইআর",
		"Quad Speakers":                  "কোয়াড স্পিকার",
		"359.2 x 251 x 17.8 mm":          "৩৫৯.২ x ২৫১ x ১৭.৮ মিমি",
		"HP Sure Sense":                  "এইচপি শিউর সেন্স",
		"14":                             "১৪",
		"20":                             "২০",
		"4800 MHz":                       "৪৮০০ মেগাহার্টজ",
		"Yes (up to 64GB)":               "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                 "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                 "হ্যাঁ (এম.২ স্লট)",
		"Shared":                         "শেয়ার্ড",
		"English/Bangla":                 "ইংরেজি/বাংলা",
		"Premium":                        "প্রিমিয়াম",
		"HDMI 2.1":                       "এইচডিএমআই ২.১",
		"2x Thunderbolt 4, 1x USB-C 3.2": "২x থান্ডারবোল্ট ৪, ১x ইউএসবি-সি ৩.২"}
}

// Seed implements the Seeder interface for HP Spectre x360 16
func (hss *HpSpectreX36016Seeder) Seed(db *gorm.DB) error {
	productSlug := "hp-spectre-x360-16"
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
		"Brand":                   "HP",
		"Model Name":              "Spectre x360 16",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-12700H",
		"Processor Generation":    "12th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "2560 x 1600 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "96 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Nocturne Blue",
		"Product Weight":          "2.05 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year International Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.3",
		"RAM":                     "32 GB",
		"Weight":                  "2.05 kg",
		"Display Type":            "OLED",
		"Refresh Rate":            "120 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "USB-C",
		"Ram":                     "32 GB",
		"Wifi Support":            "Wi-Fi 6E (802.11ax)",
		"Usb Type":                "USB 4.0 Gen 3",
		"Battery":                 "96 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "120",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "2-in-1 Convertible",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "140W USB-C",
		"Standby Time":            "12-14 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "2.3 GHz base / 4.7 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "5MP IR",
		"Audio Quality":           "Quad Speakers",
		"Sensors":                 "No",
		"Dimensions":              "359.2 x 251 x 17.8 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Nocturne Blue",
		"Special Features":        "HP Sure Sense",
		"Processor Cores":         "14",
		"Processor Threads":       "20",
		"RAM Speed":               "4800 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "NVMe PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 slot)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Premium",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "2x Thunderbolt 4, 1x USB-C 3.2",
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
		if banglaValue, exists := hss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ HP Spectre x360 16 specifications seeded successfully")
	return nil
}
