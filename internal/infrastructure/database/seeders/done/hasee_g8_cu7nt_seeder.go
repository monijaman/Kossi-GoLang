package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HaseeG8Cu7ntSeeder seeds specifications for Hasee G8-CU7NT
type HaseeG8Cu7ntSeeder struct {
	BaseSeeder
}

// NewHaseeG8Cu7ntSeeder creates a new Hasee G8-CU7NT seeder
func NewHaseeG8Cu7ntSeeder() *HaseeG8Cu7ntSeeder {
	return &HaseeG8Cu7ntSeeder{
		BaseSeeder: BaseSeeder{name: "Hasee G8-CU7NT Specifications"},
	}
}

func (hg8cu7nts *HaseeG8Cu7ntSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Hasee":                        "হাসি",
		"G8-CU7NT":                     "জি৮-সিইউ৭এনটি",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 7 7735HS":               "রাইজেন ৭ ৭৭৩৫এইচএস",
		"Zen 3+":                       "জেন ৩+",
		"NVIDIA GeForce RTX 4070":      "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৭০",
		"1920 x 1080 pixels":           "১৯২০ x ১০৮০ পিক্সেল",
		"15.6 inches":                  "১৫.৬ ইঞ্চি",
		"53 Wh":                        "৫৩ ডব্লিউএইচ",
		"Plastic":                      "প্লাস্টিক",
		"Black":                        "কালো",
		"2.3 kg":                       "২.৩ কেজি",
		"2023":                         "২০২৩",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"1 TB":                         "১ টিবি",
		"Bluetooth 5.2":                "ব্লুটুথ ৫.২",
		"16 GB":                        "১৬ জিবি",
		"IPS":                          "আইপিএস",
		"165 Hz":                       "১৬৫ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6":                      "ওয়াই-ফাই ৬",
		"USB 3.2 Gen 1":                "ইউএসবি ৩.২ জেন ১",
		"Dedicated":                    "ডেডিকেটেড",
		"1 TB SSD":                     "১ টিবি এসএসডি",
		"AMD Ryzen 7":                  "এএমডি রাইজেন ৭",
		"Dual Fan":                     "ডুয়াল ফ্যান",
		"165":                          "১৬৫",
		"No":                           "না",
		"1 Year":                       "১ বছর",
		"Gaming":                       "গেমিং",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"180W Adapter":                 "১৮০ওয়াট অ্যাডাপ্টার",
		"3-4 hours gaming":             "৩-৪ ঘন্টা গেমিং",
		"IPS, 300 nits":                "আইপিএস, ৩০০ নিটস",
		"2.7 GHz base / 4.7 GHz boost": "২.৭ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"16 MB cache":                  "১৬ এমবি ক্যাশ",
		"1MP":                          "১এমপি",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"360 x 250 x 25 mm":            "৩৬০ x ২৫০ x ২৫ মিমি",
		"8":                            "৮",
		"16":                           "১৬",
		"4800 MHz":                     "৪৮০০ মেগাহার্টজ",
		"Yes (up to 32GB)":             "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"PCIe Gen4":                    "পিসিআই জেন৪",
		"Yes (M.2)":                    "হ্যাঁ (এম.২)",
		"8 GB GDDR6":                   "৮ জিবি জিডিডিআর৬",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"3x USB 3.2 Gen 1, USB-C":      "৩x ইউএসবি ৩.২ জেন ১, ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for Hasee G8-CU7NT
func (hg8cu7nts *HaseeG8Cu7ntSeeder) Seed(db *gorm.DB) error {
	productSlug := "hasee-g8-cu7nt"
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
		"Model Name":              "G8-CU7NT",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 7 7735HS",
		"Processor Generation":    "Zen 3+",
		"Graphics Card":           "NVIDIA GeForce RTX 4070",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "15.6 inches",
		"Battery Capacity":        "53 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Black",
		"Product Weight":          "2.3 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "16 GB",
		"Weight":                  "2.3 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "165 Hz",
		"Screen Size":             "15.6 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6",
		"Usb Type":                "USB 3.2 Gen 1",
		"Battery":                 "53 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "AMD Ryzen 7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "165",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "180W Adapter",
		"Standby Time":            "3-4 hours gaming",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.7 GHz base / 4.7 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1MP",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "360 x 250 x 25 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "No",
		"Processor Cores":         "8",
		"Processor Threads":       "16",
		"RAM Speed":               "4800 MHz",
		"RAM Slots":               "2",
		"RAM Expandable":          "Yes (up to 32GB)",
		"Storage Interface":       "PCIe Gen4",
		"Storage Expandable":      "Yes (M.2)",
		"Graphics VRAM":           "8 GB GDDR6",
		"Display Touch Support":   "No",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "3x USB 3.2 Gen 1, USB-C",
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
		if banglaValue, exists := hg8cu7nts.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Hasee G8-CU7NT specifications seeded successfully")
	return nil
}
