package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ClevoNH55Seeder seeds specifications for Clevo NH55
type ClevoNH55Seeder struct {
	BaseSeeder
}

// NewClevoNH55Seeder creates a new Clevo NH55 seeder
func NewClevoNH55Seeder() *ClevoNH55Seeder {
	return &ClevoNH55Seeder{
		BaseSeeder: BaseSeeder{name: "Clevo NH55 Specifications"},
	}
}

func (cns *ClevoNH55Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Clevo":                        "ক্লেভো",
		"NH55":                         "এনএইচ৫৫",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"Intel":                        "ইন্টেল",
		"Core i5-11300H":               "কোর আই৫-১১৩০০এইচ",
		"11th Gen":                     "১১তম প্রজন্ম",
		"NVIDIA GeForce GTX 1650":      "এনভিডিয়া জিফোর্স জিটিএক্স ১৬৫০",
		"1920 x 1080 pixels":           "১৯২০ x ১০৮০ পিক্সেল",
		"15.6 inches":                  "১৫.৬ ইঞ্চি",
		"52 Wh":                        "৫২ ডব্লিউএইচ",
		"Plastic/Metal":                "প্লাস্টিক/মেটাল",
		"Black":                        "কালো",
		"2.1 kg":                       "২.১ কেজি",
		"2021":                         "২০২১",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"256 GB":                       "২৫৬ জিবি",
		"Bluetooth 5.1":                "ব্লুটুথ ৫.১",
		"8 GB":                         "৮ জিবি",
		"IPS":                          "আইপিএস",
		"60 Hz":                        "৬০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":           "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2":                      "ইউএসবি ৩.২",
		"Dedicated":                    "ডেডিকেটেড",
		"256 GB SSD":                   "২৫৬ জিবি এসএসডি",
		"Intel Core i5":                "ইন্টেল কোর আই৫",
		"Dual Fan":                     "ডুয়াল ফ্যান",
		"60":                           "৬০",
		"1 Year":                       "১ বছর",
		"Gaming":                       "গেমিং",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"120W Adapter":                 "১২০ওয়াট অ্যাডাপ্টার",
		"6-8 hours":                    "৬-৮ ঘন্টা",
		"IPS, 250 nits":                "আইপিএস, ২৫০ নিটস",
		"2.6 GHz base / 4.4 GHz boost": "২.৬ গিগাহার্টজ বেস / ৪.৪ গিগাহার্টজ বুস্ট",
		"8 MB cache":                   "৮ এমবি ক্যাশ",
		"720p HD":                      "৭২০পি এইচডি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"365 x 255 x 21 mm":            "৩৬৫ x ২৫৫ x ২১ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"4":                            "৪",
		"8":                            "৮",
		"3200 MHz":                     "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":             "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":               "এনভিএমই পিসিআই জেন৩",
		"Yes (1x M.2 slot)":            "হ্যাঁ (১x এম.২ স্লট)",
		"4 GB GDDR6":                   "৪ জিবি জিডিডিআর৬",
		"No":                           "না",
		"English":                      "ইংরেজি",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.0":                     "এইচডিএমআই ২.০",
		"2x USB 3.2, 1x USB-C":         "২x ইউএসবি ৩.২, ১x ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for Clevo NH55
func (cns *ClevoNH55Seeder) Seed(db *gorm.DB) error {
	productSlug := "clevo-nh55"
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
		"Brand":                "Clevo",
		"Model Name":           "NH55",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i5-11300H",
		"Processor Generation": "11th Gen",
		"Chipset":              "NVIDIA GeForce GTX 1650",
		"Graphics Card":        "NVIDIA GeForce GTX 1650",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "15.6 inches",
		"Battery Capacity":     "52 Wh",
		"Build Material":       "Plastic/Metal",
		"Color":                "Black",
		"Product Weight":       "2.1 kg",
		"Release Year":         "2021",
		"Warranty":             "1 Year Warranty",
		"Storage Capacity":     "256 GB",
		"Bluetooth Version":    "Bluetooth 5.1",
		"RAM":                  "8 GB",
		"Weight":               "2.1 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "15.6 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2",
		"Battery":              "52 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "256 GB SSD",
		"Cpu Type":             "Intel Core i5",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "120W Adapter",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 250 nits",
		"Processor Speed":         "2.6 GHz base / 4.4 GHz boost",
		"Clock Feature":           "8 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "365 x 255 x 21 mm",
		"Body Type":               "Plastic/Metal",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "4",
		"Processor Threads":     "8",
		"RAM Speed":             "3200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen3",
		"Storage Expandable":    "Yes (1x M.2 slot)",
		"Graphics VRAM":         "4 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "Yes",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.0",
		"USB Ports":   "2x USB 3.2, 1x USB-C",
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
		if banglaValue, exists := cns.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Clevo NH55 specifications seeded successfully")
	return nil
}
