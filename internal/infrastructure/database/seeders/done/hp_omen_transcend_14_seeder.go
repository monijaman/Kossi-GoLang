package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HpOmenTranscend14Seeder seeds specifications for HP Omen Transcend 14
type HpOmenTranscend14Seeder struct {
	BaseSeeder
}

// NewHpOmenTranscend14Seeder creates a new HP Omen Transcend 14 seeder
func NewHpOmenTranscend14Seeder() *HpOmenTranscend14Seeder {
	return &HpOmenTranscend14Seeder{
		BaseSeeder: BaseSeeder{name: "HP Omen Transcend 14 Specifications"},
	}
}

func (hots *HpOmenTranscend14Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HP":                            "এইচপি",
		"Omen Transcend 14":             "ওমেন ট্রান্সেন্ড ১৪",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i7-13700H":                "কোর আই৭-১৩৭০০এইচ",
		"13th Gen":                      "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 4060":       "এনভিডিয়া জিফোর্স আরটিএক্স ৪০৬০",
		"1920 x 1200 pixels":            "১৯২০ x ১২০০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"73 Wh":                         "৭৩ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Shadow Black":                  "শ্যাডো ব্ল্যাক",
		"1.8 kg":                        "১.৮ কেজি",
		"2024":                          "২০২৪",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                          "১ টিবি",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"16 GB":                         "১৬ জিবি",
		"IPS":                           "আইপিএস",
		"165 Hz":                        "১৬৫ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Dedicated":                     "ডেডিকেটেড",
		"1 TB SSD":                      "১ টিবি এসএসডি",
		"Intel Core i7":                 "ইন্টেল কোর আই৭",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"165":                           "১৬৫",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"High-End Gaming":              "হাই-এন্ড গেমিং",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"200W AC Adapter":              "২০০ওয়াট এসি অ্যাডাপ্টার",
		"4-6 hours":                    "৪-৬ ঘন্টা",
		"IPS, 400 nits":                "আইপিএস, ৪০০ নিটস",
		"2.4 GHz base / 5.0 GHz boost": "২.৪ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"24 MB cache":                  "২৪ এমবি ক্যাশ",
		"1080p FHD":                    "১০৮০পি এফএইচডি",
		"B&O Speakers":                 "বি&ও স্পিকার",
		"313 x 227 x 19.9 mm":          "৩১৩ x ২২৭ x ১৯.৯ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"14":                           "১৪",
		"20":                           "২০",
		"5600 MHz":                     "৫৬০০ মেগাহার্টজ",
		"Yes (up to 32GB)":             "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":               "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":               "হ্যাঁ (এম.২ স্লট)",
		"8 GB GDDR6":                   "৮ জিবি জিডিডিআর৬",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Gaming":                       "গেমিং",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2, 1x Thunderbolt 4": "২x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২, ১x থান্ডারবোল্ট ৪"}
}

// Seed implements the Seeder interface for HP Omen Transcend 14
func (hots *HpOmenTranscend14Seeder) Seed(db *gorm.DB) error {
	productSlug := "hp-omen-transcend-14"
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
		"Brand":                "HP",
		"Model Name":           "Omen Transcend 14",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i7-13700H",
		"Processor Generation": "13th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "NVIDIA GeForce RTX 4060",
		"Screen Resolution":    "1920 x 1200 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "73 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Shadow Black",
		"Product Weight":       "1.8 kg",
		"Release Year":         "2024",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "1 TB",
		"Bluetooth Version":    "Bluetooth 5.3",
		"RAM":                  "16 GB",
		"Weight":               "1.8 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "165 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "73 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "1 TB SSD",
		"Cpu Type":             "Intel Core i7",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "165",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "High-End Gaming",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "200W AC Adapter",
		"Standby Time":            "4-6 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.4 GHz base / 5.0 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "B&O Speakers",
		"Sensors":                 "No",
		"Dimensions":              "313 x 227 x 19.9 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Shadow Black",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "14",
		"Processor Threads":     "20",
		"RAM Speed":             "5600 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "8 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "Thunderbolt 4",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Gaming",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "2x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2, 1x Thunderbolt 4",
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
		if banglaValue, exists := hots.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ HP Omen Transcend 14 specifications seeded successfully")
	return nil
}
