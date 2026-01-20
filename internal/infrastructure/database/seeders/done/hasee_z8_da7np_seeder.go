package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HaseeZ8Da7npSeeder seeds specifications for Hasee Z8-DA7NP
type HaseeZ8Da7npSeeder struct {
	BaseSeeder
}

// NewHaseeZ8Da7npSeeder creates a new Hasee Z8-DA7NP seeder
func NewHaseeZ8Da7npSeeder() *HaseeZ8Da7npSeeder {
	return &HaseeZ8Da7npSeeder{
		BaseSeeder: BaseSeeder{name: "Hasee Z8-DA7NP Specifications"},
	}
}

func (hz8da7nps *HaseeZ8Da7npSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hasee":                        "হাসি",
		"Z8-DA7NP":                     "জেড৮-ডিএ৭এনপি",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 9 7945HX":               "রাইজেন ৯ ৭৯৪৫এইচএক্স",
		"Zen 4":                        "জেন ৪",
		"NVIDIA GeForce RTX 4080":      "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৮০",
		"2560 x 1440 pixels":           "২৫৬০ x ১৪৪০ পিক্সেল",
		"16 inches":                    "১৬ ইঞ্চি",
		"90 Wh":                        "৯০ ডব্লিউএইচ",
		"Aluminum":                     "অ্যালুমিনিয়াম",
		"Black":                        "কালো",
		"2.5 kg":                       "২.৫ কেজি",
		"2024":                         "২০২৪",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"2 TB":                         "২ টিবি",
		"Bluetooth 5.3":                "ব্লুটুথ ৫.৩",
		"32 GB":                        "৩২ জিবি",
		"IPS":                          "আইপিএস",
		"240 Hz":                       "২৪০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E":                     "ওয়াই-ফাই ৬ই",
		"USB 3.2 Gen 2":                "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                    "ডেডিকেটেড",
		"2 TB SSD":                     "২ টিবি এসএসডি",
		"AMD Ryzen 9":                  "এএমডি রাইজেন ৯",
		"Dual Fan":                     "ডুয়াল ফ্যান",
		"240":                          "২৪০",
		"No":                           "না",
		"1 Year":                       "১ বছর",
		"Gaming":                       "গেমিং",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"240W Adapter":                 "২৪০ওয়াট অ্যাডাপ্টার",
		"2-3 hours gaming":             "২-৩ ঘন্টা গেমিং",
		"IPS, 500 nits":                "আইপিএস, ৫০০ নিটস",
		"2.5 GHz base / 5.4 GHz boost": "২.৫ গিগাহার্টজ বেস / ৫.৪ গিগাহার্টজ বুস্ট",
		"64 MB cache":                  "৬৪ এমবি ক্যাশ",
		"2MP":                          "২এমপি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"370 x 260 x 25 mm":            "৩৭০ x ২৬০ x ২৫ মিমি",
		"16":                           "১৬",
		"32":                           "৩২",
		"5600 MHz":                     "৫৬০০ মেগাহার্টজ",
		"Yes (up to 64GB)":             "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"PCIe Gen5":                    "পিসিআই জেন৫",
		"Yes (M.2)":                    "হ্যাঁ (এম.২)",
		"16 GB GDDR6X":                 "১৬ জিবি জিডিডিআর৬এক্স",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"3x USB 3.2 Gen 2, USB-C":      "৩x ইউএসবি ৩.২ জেন ২, ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for Hasee Z8-DA7NP
func (hz8da7nps *HaseeZ8Da7npSeeder) Seed(db *gorm.DB) error {
	productSlug := "hasee-z8-da7np"
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
		"Brand":                   "Hasee",
		"Model Name":              "Z8-DA7NP",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 9 7945HX",
		"Processor Generation":    "Zen 4",
		"Graphics Card":           "NVIDIA GeForce RTX 4080",
		"Screen Resolution":       "2560 x 1440 pixels",
		"Display Size":            "16 inches",
		"Battery Capacity":        "90 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "2.5 kg",
		"Release Year":            "2024",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "2 TB",
		"Bluetooth Version":       "Bluetooth 5.3",
		"RAM":                     "32 GB",
		"Weight":                  "2.5 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "240 Hz",
		"Screen Size":             "16 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "32 GB",
		"Wifi Support":            "Wi-Fi 6E",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "90 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "2 TB SSD",
		"Cpu Type":                "AMD Ryzen 9",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "240",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "240W Adapter",
		"Standby Time":            "2-3 hours gaming",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1440 pixels",
		"Display Characteristics": "IPS, 500 nits",
		"Processor Speed":         "2.5 GHz base / 5.4 GHz boost",
		"Clock Feature":           "64 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "2MP",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "370 x 260 x 25 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "No",
		"Processor Cores":         "16",
		"Processor Threads":       "32",
		"RAM Speed":               "5600 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 64GB)",
		"Storage Interface":       "PCIe Gen5",
		"Storage Expandable":      "Yes (M.2)",
		"Graphics VRAM":           "16 GB GDDR6X",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "3x USB 3.2 Gen 2, USB-C",
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
		if banglaValue, exists := hz8da7nps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Hasee Z8-DA7NP specifications seeded successfully")
	return nil
}
