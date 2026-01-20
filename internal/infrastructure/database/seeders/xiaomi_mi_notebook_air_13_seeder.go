package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// XiaomiMiNotebookAir13Seeder seeds specifications for Xiaomi Mi Notebook Air 13
type XiaomiMiNotebookAir13Seeder struct {
	BaseSeeder
}

// NewXiaomiMiNotebookAir13Seeder creates a new Xiaomi Mi Notebook Air 13 seeder
func NewXiaomiMiNotebookAir13Seeder() *XiaomiMiNotebookAir13Seeder {
	return &XiaomiMiNotebookAir13Seeder{
		BaseSeeder: BaseSeeder{name: "Xiaomi Mi Notebook Air 13 Specifications"},
	}
}

func (xmnba13s *XiaomiMiNotebookAir13Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Xiaomi":                        "শাওমি",
		"Mi Notebook Air 13":            "এমআই নোটবুক এয়ার ১৩",
		"Windows 10 Home":               "উইন্ডোজ ১০ হোম",
		"Intel":                         "ইন্টেল",
		"Core i5-8250U":                 "কোর আই৫-৮২৫০ইউ",
		"8th Gen":                       "৮ম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD Graphics 620":        "ইন্টেল ইউএইচডি গ্রাফিক্স ৬২০",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"13.3 inches":                   "১৩.৩ ইঞ্চি",
		"40 Wh":                         "৪০ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Silver":                        "সিলভার",
		"1.3 kg":                        "১.৩ কেজি",
		"2018":                          "২০১৮",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"256 GB":                        "২৫৬ জিবি",
		"Bluetooth 4.2":                 "ব্লুটুথ ৪.২",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":            "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.1 Gen 1":                 "ইউএসবি ৩.১ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"256 GB SSD":                    "২৫৬ জিবি এসএসডি",
		"Intel Core i5":                 "ইন্টেল কোর আই৫",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Ultraportable":                        "আল্ট্রাপোর্টেবল",
		"Lithium-polymer":                      "লিথিয়াম-পলিমার",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"10-12 hours":                          "১০-১২ ঘন্টা",
		"IPS, 300 nits":                        "আইপিএস, ৩০০ নিটস",
		"1.6 GHz base / 3.4 GHz boost":         "১.৬ গিগাহার্টজ বেস / ৩.৪ গিগাহার্টজ বুস্ট",
		"6 MB cache":                           "৬ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"309 x 210 x 14.8 mm":                  "৩০৯ x ২১০ x ১৪.৮ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"4":                                    "৪",
		"8":                                    "৮",
		"2400 MHz":                             "২৪০০ মেগাহার্টজ",
		"No (soldered)":                        "না (সোল্ডার্ড)",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 1.4":                             "এইচডিএমআই ১.৪",
		"1x USB 3.1 Gen 1, 1x USB-C 3.1 Gen 1": "১x ইউএসবি ৩.১ জেন ১, ১x ইউএসবি-সি ৩.১ জেন ১",
	}
}

// Seed implements the Seeder interface for Xiaomi Mi Notebook Air 13
func (xmnba13s *XiaomiMiNotebookAir13Seeder) Seed(db *gorm.DB) error {
	productSlug := "xiaomi-mi-notebook-air-13"
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
		"Brand":                "Xiaomi",
		"Model Name":           "Mi Notebook Air 13",
		"Operating System":     "Windows 10 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i5-8250U",
		"Processor Generation": "8th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel UHD Graphics 620",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "13.3 inches",
		"Battery Capacity":     "40 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Silver",
		"Product Weight":       "1.3 kg",
		"Release Year":         "2018",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "256 GB",
		"Bluetooth Version":    "Bluetooth 4.2",
		"RAM":                  "8 GB",
		"Weight":               "1.3 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "13.3 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 5 (802.11ac)",
		"Usb Type":             "USB 3.1 Gen 1",
		"Battery":              "40 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "256 GB SSD",
		"Cpu Type":             "Intel Core i5",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "1.6 GHz base / 3.4 GHz boost",
		"Clock Feature":           "6 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "309 x 210 x 14.8 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "4",
		"Processor Threads":     "8",
		"RAM Speed":             "2400 MHz",
		"RAM Slots":             "1",
		"RAM Expandable":        "No (soldered)",
		"Storage Interface":     "NVMe PCIe Gen3",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 1.4",
		"USB Ports":   "1x USB 3.1 Gen 1, 1x USB-C 3.1 Gen 1",
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
		if banglaValue, exists := xmnba13s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Xiaomi Mi Notebook Air 13 specifications seeded successfully")
	return nil
}
