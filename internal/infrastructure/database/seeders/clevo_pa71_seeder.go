package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// ClevoPA71Seeder seeds specifications for Clevo PA71
type ClevoPA71Seeder struct {
	BaseSeeder
}

// NewClevoPA71Seeder creates a new Clevo PA71 seeder
func NewClevoPA71Seeder() *ClevoPA71Seeder {
	return &ClevoPA71Seeder{
		BaseSeeder: BaseSeeder{name: "Clevo PA71 Specifications"},
	}
}

func (cps *ClevoPA71Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Clevo":                        "ক্লেভো",
		"PA71":                         "পিএ৭১",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"Intel":                        "ইন্টেল",
		"Core i7-11800H":               "কোর আই৭-১১৮০০এইচ",
		"11th Gen":                     "১১তম প্রজন্ম",
		"NVIDIA GeForce GTX 1650":      "এনভিডিয়া জিফোর্স জিটিএক্স ১৬৫০",
		"1920 x 1080 pixels":           "১৯২০ x ১০৮০ পিক্সেল",
		"17.3 inches":                  "১৭.৩ ইঞ্চি",
		"65 Wh":                        "৬৫ ডব্লিউএইচ",
		"Plastic/Metal":                "প্লাস্টিক/মেটাল",
		"Black":                        "কালো",
		"2.8 kg":                       "২.৮ কেজি",
		"2022":                         "২০২২",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"512 GB":                       "৫১২ জিবি",
		"Bluetooth 5.2":                "ব্লুটুথ ৫.২",
		"16 GB":                        "১৬ জিবি",
		"IPS":                          "আইপিএস",
		"144 Hz":                       "১৪৪ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":           "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2":                      "ইউএসবি ৩.২",
		"Dedicated":                    "ডেডিকেটেড",
		"512 GB SSD":                   "৫১২ জিবি এসএসডি",
		"Intel Core i7":                "ইন্টেল কোর আই৭",
		"Dual Fan":                     "ডুয়াল ফ্যান",
		"144":                          "১৪৪",
		"1 Year":                       "১ বছর",
		"Gaming":                       "গেমিং",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"180W Adapter":                 "১৮০ওয়াট অ্যাডাপ্টার",
		"4-6 hours":                    "৪-৬ ঘন্টা",
		"IPS, 300 nits":                "আইপিএস, ৩০০ নিটস",
		"2.3 GHz base / 4.6 GHz boost": "২.৩ গিগাহার্টজ বেস / ৪.৬ গিগাহার্টজ বুস্ট",
		"24 MB cache":                  "২৪ এমবি ক্যাশ",
		"1080p HD":                     "১০৮০পি এইচডি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"395 x 265 x 25 mm":            "৩৯৫ x ২৬৫ x ২৫ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"8":                            "৮",
		"16":                           "১৬",
		"3200 MHz":                     "৩২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":             "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":               "এনভিএমই পিসিআই জেন৪",
		"Yes (2x M.2 slots)":           "হ্যাঁ (২x এম.২ স্লট)",
		"4 GB GDDR6":                   "৪ জিবি জিডিডিআর৬",
		"No":                           "না",
		"English":                      "ইংরেজি",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"3x USB 3.2, 1x USB-C":         "৩x ইউএসবি ৩.২, ১x ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for Clevo PA71
func (cps *ClevoPA71Seeder) Seed(db *gorm.DB) error {
	productSlug := "clevo-pa71"
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
		"Model Name":           "PA71",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i7-11800H",
		"Processor Generation": "11th Gen",
		"Chipset":              "NVIDIA GeForce GTX 1650",
		"Graphics Card":        "NVIDIA GeForce GTX 1650",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "17.3 inches",
		"Battery Capacity":     "65 Wh",
		"Build Material":       "Plastic/Metal",
		"Color":                "Black",
		"Product Weight":       "2.8 kg",
		"Release Year":         "2022",
		"Warranty":             "1 Year Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "16 GB",
		"Weight":               "2.8 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "144 Hz",
		"Screen Size":          "17.3 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2",
		"Battery":              "65 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "Intel Core i7",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "144",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "180W Adapter",
		"Standby Time":            "4-6 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.3 GHz base / 4.6 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "395 x 265 x 25 mm",
		"Body Type":               "Plastic/Metal",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "8",
		"Processor Threads":     "16",
		"RAM Speed":             "3200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 64GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (2x M.2 slots)",
		"Graphics VRAM":         "4 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "Yes",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "3x USB 3.2, 1x USB-C",
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
		if banglaValue, exists := cps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Clevo PA71 specifications seeded successfully")
	return nil
}
