package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// FrameworkLaptop13AMDRyzen7040Seeder seeds specifications for Framework Laptop 13 (AMD Ryzen 7040)
type FrameworkLaptop13AMDRyzen7040Seeder struct {
	BaseSeeder
}

// NewFrameworkLaptop13AMDRyzen7040Seeder creates a new Framework Laptop 13 (AMD Ryzen 7040) seeder
func NewFrameworkLaptop13AMDRyzen7040Seeder() *FrameworkLaptop13AMDRyzen7040Seeder {
	return &FrameworkLaptop13AMDRyzen7040Seeder{
		BaseSeeder: BaseSeeder{name: "Framework Laptop 13 (AMD Ryzen 7040) Specifications"},
	}
}

func (fls *FrameworkLaptop13AMDRyzen7040Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Framework":                            "ফ্রেমওয়ার্ক",
		"Framework Laptop 13 (AMD Ryzen 7040)": "ফ্রেমওয়ার্ক ল্যাপটপ ১৩ (এএমডি রাইজেন ৭০৪০)",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"AMD":                                  "এএমডি",
		"Ryzen 7 7840U":                        "রাইজেন ৭ ৭৮৪০ইউ",
		"AMD Ryzen 7000":                       "এএমডি রাইজেন ৭০০০",
		"AMD Radeon 780M":                      "এএমডি রেডিওন ৭৮০এম",
		"2256 x 1504 pixels":                   "২২৫৬ x ১৫০৪ পিক্সেল",
		"13.5 inches":                          "১৩.৫ ইঞ্চি",
		"55 Wh":                                "৫৫ ডব্লিউএইচ",
		"Aluminum":                             "অ্যালুমিনিয়াম",
		"Gray":                                 "গ্রে",
		"1.3 kg":                               "১.৩ কেজি",
		"2023":                                 "২০২৩",
		"1 Year International Warranty":        "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                               "৫১২ জিবি",
		"Bluetooth 5.3":                        "ব্লুটুথ ৫.৩",
		"16 GB":                                "১৬ জিবি",
		"IPS":                                  "আইপিএস",
		"60 Hz":                                "৬০ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":                  "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                        "ইউএসবি ৩.২ জেন ২",
		"Integrated":                           "ইন্টিগ্রেটেড",
		"512 GB SSD":                           "৫১২ জিবি এসএসডি",
		"AMD Ryzen 7":                          "এএমডি রাইজেন ৭",
		"Dual Fan":                             "ডুয়াল ফ্যান",
		"60":                                   "৬০",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		// New translations for added specs
		"Modular":                                "মডুলার",
		"Lithium-polymer":                        "লিথিয়াম-পলিমার",
		"60W USB-C":                              "৬০ওয়াট ইউএসবি-সি",
		"8-10 hours":                             "৮-১০ ঘন্টা",
		"IPS, 400 nits":                          "আইপিএস, ৪০০ নিটস",
		"2.3 GHz base / 4.9 GHz boost":           "২.৩ গিগাহার্টজ বেস / ৪.৯ গিগাহার্টজ বুস্ট",
		"16 MB cache":                            "১৬ এমবি ক্যাশ",
		"1080p FHD":                              "১০৮০পি এফএইচডি",
		"Stereo Speakers":                        "স্টেরিও স্পিকার",
		"297.0 x 229.0 x 15.85 mm":               "২৯৭.০ x ২২৯.০ x ১৫.৮৫ মিমি",
		"TPM 2.0":                                "টিপিএম ২.০",
		"8":                                      "৮",
		"16":                                     "১৬",
		"5200 MHz":                               "৫২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                       "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                         "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                         "হ্যাঁ (এম.২ স্লট)",
		"Shared":                                 "শেয়ার্ড",
		"English/Bangla":                         "ইংরেজি/বাংলা",
		"Premium":                                "প্রিমিয়াম",
		"HDMI 2.0":                               "এইচডিএমআই ২.০",
		"2x USB-C 3.2 Gen 2, 1x USB-A 3.2 Gen 2": "২x ইউএসবি-সি ৩.২ জেন ২, ১x ইউএসবি-এ ৩.২ জেন ২"}
}

// Seed implements the Seeder interface for Framework Laptop 13 (AMD Ryzen 7040)
func (fls *FrameworkLaptop13AMDRyzen7040Seeder) Seed(db *gorm.DB) error {
	productSlug := "framework-laptop-13-amd-ryzen-7040"
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
		"Brand":                "Framework",
		"Model Name":           "Framework Laptop 13 (AMD Ryzen 7040)",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "AMD",
		"Processor Model":      "Ryzen 7 7840U",
		"Processor Generation": "AMD Ryzen 7000",
		"Graphics Card":        "AMD Radeon 780M",
		"Screen Resolution":    "2256 x 1504 pixels",
		"Display Size":         "13.5 inches",
		"Battery Capacity":     "55 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Gray",
		"Product Weight":       "1.3 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.3",
		"RAM":                  "16 GB",
		"Weight":               "1.3 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "13.5 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 2",
		"Battery":              "55 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "AMD Ryzen 7",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Modular",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "60W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2256 x 1504 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.3 GHz base / 4.9 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "Fingerprint Reader",
		"Dimensions":              "297.0 x 229.0 x 15.85 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Gray",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "8",
		"Processor Threads":     "16",
		"RAM Speed":             "5200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 64GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "Expansion Card",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "Expansion Card",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.0",
		"USB Ports":   "2x USB-C 3.2 Gen 2, 1x USB-A 3.2 Gen 2",
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
		if banglaValue, exists := fls.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Framework Laptop 13 (AMD Ryzen 7040) specifications seeded successfully")
	return nil
}
