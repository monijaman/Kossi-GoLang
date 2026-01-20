package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// DellPrecision5680Seeder seeds specifications for Dell Precision 5680
type DellPrecision5680Seeder struct {
	BaseSeeder
}

// NewDellPrecision5680Seeder creates a new Dell Precision 5680 seeder
func NewDellPrecision5680Seeder() *DellPrecision5680Seeder {
	return &DellPrecision5680Seeder{
		BaseSeeder: BaseSeeder{name: "Dell Precision 5680 Specifications"},
	}
}

func (dps *DellPrecision5680Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Dell":                          "ডেল",
		"Precision 5680":                "প্রিসিশন ৫৬৮০",
		"Windows 11 Pro":                "উইন্ডোজ ১১ প্রো",
		"Intel":                         "ইন্টেল",
		"Core i9-13950HX":               "কোর আই৯-১৩৯৫০এইচএক্স",
		"13th Gen":                      "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA RTX A4500":              "এনভিডিয়া আরটিএক্স এ৪৫০০",
		"1920 x 1200 pixels":            "১৯২০ x ১২০০ পিক্সেল",
		"16 inches":                     "১৬ ইঞ্চি",
		"90 Wh":                         "৯০ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Gray":                          "গ্রে",
		"2.5 kg":                        "২.৫ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                          "১ টেরাবাইট",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"64 GB":                         "৬৪ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                     "ডেডিকেটেড",
		"1 TB SSD":                      "১ টেরাবাইট এসএসডি",
		"Intel Core i9":                 "ইন্টেল কোর আই৯",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Workstation":                          "ওয়ার্কস্টেশন",
		"Lithium-polymer":                      "লিথিয়াম-পলিমার",
		"130W USB-C":                           "১৩০ওয়াট ইউএসবি-সি",
		"4-6 hours":                            "৪-৬ ঘন্টা",
		"IPS, 400 nits":                        "আইপিএস, ৪০০ নিটস",
		"2.2 GHz base / 5.5 GHz boost":         "২.২ গিগাহার্টজ বেস / ৫.৫ গিগাহার্টজ বুস্ট",
		"36 MB cache":                          "৩৬ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"363.0 x 254.0 x 19.0 mm":              "৩৬৩.০ x ২৫৪.০ x ১৯.০ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"24":                                   "২৪",
		"32":                                   "৩২",
		"5600 MHz":                             "৫৬০০ মেগাহার্টজ",
		"Yes (up to 128GB)":                    "হ্যাঁ (১২৮জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"16 GB GDDR6":                          "১৬ জিবি জিডিডিআর৬",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Premium":                              "প্রিমিয়াম",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"2x Thunderbolt 4, 1x USB-C 3.2 Gen 2": "২x থান্ডারবোল্ট ৪, ১x ইউএসবি-সি ৩.২ জেন ২"}
}

// Seed implements the Seeder interface for Dell Precision 5680
func (dps *DellPrecision5680Seeder) Seed(db *gorm.DB) error {
	productSlug := "dell-precision-5680"
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
		"Brand":                "Dell",
		"Model Name":           "Precision 5680",
		"Operating System":     "Windows 11 Pro",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i9-13950HX",
		"Processor Generation": "13th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "NVIDIA RTX A4500",
		"Screen Resolution":    "1920 x 1200 pixels",
		"Display Size":         "16 inches",
		"Battery Capacity":     "90 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Gray",
		"Product Weight":       "2.5 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "1 TB",
		"Bluetooth Version":    "Bluetooth 5.3",
		"RAM":                  "64 GB",
		"Weight":               "2.5 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "16 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "64 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 2",
		"Battery":              "90 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "1 TB SSD",
		"Cpu Type":             "Intel Core i9",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Workstation",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "130W USB-C",
		"Standby Time":            "4-6 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.2 GHz base / 5.5 GHz boost",
		"Clock Feature":           "36 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "Fingerprint Reader",
		"Dimensions":              "363.0 x 254.0 x 19.0 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Gray",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "24",
		"Processor Threads":     "32",
		"RAM Speed":             "5600 MHz",
		"RAM Slots":             "4",
		"RAM Expandable":        "Yes (up to 128GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "16 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "Yes",
		"Thunderbolt Version":   "4",
		"SD Card Reader":        "Yes",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "2x Thunderbolt 4, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := dps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Dell Precision 5680 specifications seeded successfully")
	return nil
}
