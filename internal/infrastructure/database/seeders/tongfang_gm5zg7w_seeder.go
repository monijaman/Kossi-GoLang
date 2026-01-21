package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// TongfangGm5zg7wSeeder seeds specifications for Tongfang GM5ZG7W
type TongfangGm5zg7wSeeder struct {
	BaseSeeder
}

// NewTongfangGm5zg7wSeeder creates a new Tongfang GM5ZG7W seeder
func NewTongfangGm5zg7wSeeder() *TongfangGm5zg7wSeeder {
	return &TongfangGm5zg7wSeeder{
		BaseSeeder: BaseSeeder{name: "Tongfang GM5ZG7W Specifications"},
	}
}

func (tgs *TongfangGm5zg7wSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Tongfang":                             "টংফ্যাং",
		"GM5ZG7W":                              "জিএম৫জেডজি৭ডব্লিউ",
		"Windows 11 Home":                      "উইন্ডোজ ১১ হোম",
		"AMD":                                  "এএমডি",
		"Ryzen 7 7735HS":                       "রাইজেন ৭ ৭৭৩৫এইচএস",
		"7th Gen":                              "৭ম প্রজন্ম",
		"AMD Platform Controller Hub":          "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 4060":              "এনভিডিয়া গেফোর্স আরটিএক্স ৪০৬০",
		"1920 x 1080 pixels":                   "১৯২০ x ১০৮০ পিক্সেল",
		"15.6 inches":                          "১৫.৬ ইঞ্চি",
		"56 Wh":                                "৫৬ ডব্লিউএইচ",
		"Plastic":                              "প্লাস্টিক",
		"Black":                                "ব্ল্যাক",
		"2.1 kg":                               "২.১ কেজি",
		"2024":                                 "২০২৪",
		"1 Year International Warranty":        "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                               "৫১২ জিবি",
		"Bluetooth 5.2":                        "ব্লুটুথ ৫.২",
		"16 GB":                                "১৬ জিবি",
		"IPS":                                  "আইপিএস",
		"144 Hz":                               "১৪৪ হার্জ",
		"Yes":                                  "হ্যাঁ",
		"3.5mm Combo Jack":                     "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":                   "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                        "ইউএসবি ৩.২ জেন ১",
		"Dedicated":                            "ডেডিকেটেড",
		"512 GB SSD":                           "৫১২ জিবি এসএসডি",
		"AMD Ryzen 7":                          "এএমডি রাইজেন ৭",
		"Dual Fan":                             "ডুয়াল ফ্যান",
		"144":                                  "১৪৪",
		"No":                                   "না",
		"1 Year":                               "১ বছর",
		"Gaming":                               "গেমিং",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"180W":                                 "১৮০ওয়াট",
		"5-7 hours":                            "৫-৭ ঘন্টা",
		"IPS, 300 nits":                        "আইপিএস, ৩০০ নিটস",
		"2.0 GHz base / 4.7 GHz boost":         "২.০ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"16 MB cache":                          "১৬ এমবি ক্যাশ",
		"1080p HD":                             "১০৮০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"354 x 232 x 22 mm":                    "৩৫৪ x ২৩২ x ২২ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"8":                                    "৮",
		"16":                                   "১৬",
		"4800 MHz":                             "৪৮০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                     "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"8 GB GDDR6":                           "৮ জিবি জিডিডিআর৬",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"3x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "৩x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Tongfang GM5ZG7W
func (tgs *TongfangGm5zg7wSeeder) Seed(db *gorm.DB) error {
	productSlug := "tongfang-gm5zg7w"
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
		"Brand":                "Tongfang",
		"Model Name":           "GM5ZG7W",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "AMD",
		"Processor Model":      "Ryzen 7 7735HS",
		"Processor Generation": "7th Gen",
		"Chipset":              "AMD Platform Controller Hub",
		"Graphics Card":        "NVIDIA GeForce RTX 4060",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "15.6 inches",
		"Battery Capacity":     "56 Wh",
		"Build Material":       "Plastic",
		"Color":                "Black",
		"Product Weight":       "2.1 kg",
		"Release Year":         "2024",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "16 GB",
		"Weight":               "2.1 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "144 Hz",
		"Screen Size":          "15.6 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "56 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "AMD Ryzen 7",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "144",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "180W",
		"Standby Time":            "5-7 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.0 GHz base / 4.7 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "354 x 232 x 22 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "8",
		"Processor Threads":     "16",
		"RAM Speed":             "4800 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 64GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "8 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "3x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := tgs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Tongfang GM5ZG7W specifications seeded successfully")
	return nil
}
