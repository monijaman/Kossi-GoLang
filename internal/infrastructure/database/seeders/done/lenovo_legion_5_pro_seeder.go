package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LenovoLegion5ProSeeder seeds specifications for Lenovo Legion 5 Pro
type LenovoLegion5ProSeeder struct {
	BaseSeeder
}

// NewLenovoLegion5ProSeeder creates a new Lenovo Legion 5 Pro seeder
func NewLenovoLegion5ProSeeder() *LenovoLegion5ProSeeder {
	return &LenovoLegion5ProSeeder{
		BaseSeeder: BaseSeeder{name: "Lenovo Legion 5 Pro Specifications"},
	}
}

func (llp *LenovoLegion5ProSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Lenovo":                        "লেনোভো",
		"Legion 5 Pro":                  "লেজিওন ৫ প্রো",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"AMD":                           "এএমডি",
		"Ryzen 7 6800H":                 "রাইজেন ৭ ৬৮০০এইচ",
		"6th Gen":                       "৬ষ্ঠ প্রজন্ম",
		"AMD Platform Controller Hub":   "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 3060":       "এনভিডিয়া জিফোর্স আরটিএক্স ৩০৬০",
		"2560 x 1600 pixels":            "২৫৬০ x ১৬০০ পিক্সেল",
		"16 inches":                     "১৬ ইঞ্চি",
		"80 Wh":                         "৮০ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Black":                         "কালো",
		"2.45 kg":                       "২.৪৫ কেজি",
		"2022":                          "২০২২",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                          "১ টিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"16 GB":                         "১৬ জিবি",
		"IPS":                           "আইপিএস",
		"165 Hz":                        "১৬৫ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                     "ডেডিকেটেড",
		"1 TB SSD":                      "১ টিবি এসএসডি",
		"AMD Ryzen 7":                   "এএমডি রাইজেন ৭",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"165":                           "১৬৫",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Gaming":                               "গেমিং",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"230W USB-C":                           "২৩০ওয়াট ইউএসবি-সি",
		"6-8 hours":                            "৬-৮ ঘন্টা",
		"IPS, 350 nits":                        "আইপিএস, ৩৫০ নিটস",
		"3.2 GHz base / 4.7 GHz boost":         "৩.২ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"16 MB cache":                          "১৬ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"356 x 264 x 20 mm":                    "৩৫৬ x ২৬৪ x ২০ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"8":                                    "৮",
		"16":                                   "১৬",
		"3200 MHz":                             "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"6 GB GDDR6":                           "৬ জিবি জিডিডিআর৬",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"2x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2": "২x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Lenovo Legion 5 Pro
func (llp *LenovoLegion5ProSeeder) Seed(db *gorm.DB) error {
	productSlug := "lenovo-legion-5-pro"
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
		"Brand":                "Lenovo",
		"Model Name":           "Legion 5 Pro",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "AMD",
		"Processor Model":      "Ryzen 7 6800H",
		"Processor Generation": "6th Gen",
		"Chipset":              "AMD Platform Controller Hub",
		"Graphics Card":        "NVIDIA GeForce RTX 3060",
		"Screen Resolution":    "2560 x 1600 pixels",
		"Display Size":         "16 inches",
		"Battery Capacity":     "80 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Black",
		"Product Weight":       "2.45 kg",
		"Release Year":         "2022",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "1 TB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "16 GB",
		"Weight":               "2.45 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "165 Hz",
		"Screen Size":          "16 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 2",
		"Battery":              "80 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "1 TB SSD",
		"Cpu Type":             "AMD Ryzen 7",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "165",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "230W USB-C",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "IPS, 350 nits",
		"Processor Speed":         "3.2 GHz base / 4.7 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "356 x 264 x 20 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "8",
		"Processor Threads":     "16",
		"RAM Speed":             "3200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "6 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "Yes",
		"Thunderbolt Version":   "4",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "2x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := llp.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Lenovo Legion 5 Pro specifications seeded successfully")
	return nil
}
