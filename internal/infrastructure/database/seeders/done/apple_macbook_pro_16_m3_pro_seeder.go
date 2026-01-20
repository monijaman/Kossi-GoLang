package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AppleMacbookPro16M3ProSeeder seeds specifications for Apple MacBook Pro 16-inch M3 Pro
type AppleMacbookPro16M3ProSeeder struct {
	BaseSeeder
}

// NewAppleMacbookPro16M3ProSeeder creates a new Apple MacBook Pro 16-inch M3 Pro seeder
func NewAppleMacbookPro16M3ProSeeder() *AppleMacbookPro16M3ProSeeder {
	return &AppleMacbookPro16M3ProSeeder{
		BaseSeeder: BaseSeeder{name: "Apple MacBook Pro 16-inch M3 Pro Specifications"},
	}
}

func (amas *AppleMacbookPro16M3ProSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Apple":                         "অ্যাপল",
		"MacBook Pro 16-inch M3 Pro":    "ম্যাকবুক প্রো ১৬-ইঞ্চি এম৩ প্রো",
		"macOS Sonoma":                  "ম্যাকওএস সোনোমা",
		"M3 Pro":                        "এম৩ প্রো",
		"Apple M3 Pro GPU":              "অ্যাপল এম৩ প্রো জিপিইউ",
		"3456 x 2234 pixels":            "৩৪৫৬ x ২২৩৪ পিক্সেল",
		"16.2 inches":                   "১৬.২ ইঞ্চি",
		"100 Wh":                        "১০০ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Space Gray":                    "স্পেস গ্রে",
		"2.15 kg":                       "২.১৫ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"18 GB":                         "১৮ জিবি",
		"Liquid Retina XDR":             "লিকুইড রেটিনা এক্সডিআর",
		"120 Hz":                        "১২০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"No":                            "না",
		"Wi-Fi 6E":                      "ওয়াই-ফাই ৬ই",
		"Thunderbolt 4":                 "থান্ডারবোল্ট ৪",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"Apple M3 Pro":                  "অ্যাপল এম৩ প্রো",
		"Active":                        "অ্যাকটিভ",
		"120":                           "১২০",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Professional":                    "প্রফেশনাল",
		"Lithium-polymer":                 "লিথিয়াম-পলিমার",
		"MagSafe":                         "ম্যাগসেফ",
		"Up to 22 hours":                  "আপ টু ২২ ঘন্টা",
		"Liquid Retina XDR, 1600 nits":    "লিকুইড রেটিনা এক্সডিআর, ১৬০০ নিটস",
		"12-core CPU, 18-core GPU":        "১২-কোর সিপিইউ, ১৮-কোর জিপিইউ",
		"Up to 4.06 GHz":                  "আপ টু ৪.০৬ গিগাহার্টজ",
		"1080p FaceTime HD":               "১০৮০পি ফেসটাইম এইচডি",
		"Three microphones, six speakers": "থ্রি মাইক্রোফোন, সিক্স স্পিকার",
		"Ambient light sensor":            "অ্যাম্বিয়েন্ট লাইট সেন্সর",
		"14.01 x 9.77 x 0.66 inches":      "১৪.০১ x ৯.৭৭ x ০.৬৬ ইঞ্চি",
		"Space Gray, Silver":              "স্পেস গ্রে, সিলভার",
		"Touch ID":                        "টাচ আইডি",
		// Major laptop specs translations
		"12":        "১২",
		"LPDDR5":    "এলপিডিডিআর৫",
		"Soldered":  "সোল্ডার্ড",
		"NVMe PCIe": "এনভিএমই পিসিআই",
		"Shared":    "শেয়ার্ড",

		"English/Bangla":   "ইংরেজি/বাংলা",
		"Premium":          "প্রিমিয়াম",
		"3x Thunderbolt 4": "৩x থান্ডারবোল্ট ৪",
	}
}

// Seed implements the Seeder interface for Apple MacBook Pro 16-inch M3 Pro
func (amas *AppleMacbookPro16M3ProSeeder) Seed(db *gorm.DB) error {
	productSlug := "apple-macbook-pro-16-inch-m3-pro"
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
		"Brand":                "Apple",
		"Model Name":           "MacBook Pro 16-inch M3 Pro",
		"Operating System":     "macOS Sonoma",
		"Processor Brand":      "Apple",
		"Processor Model":      "M3 Pro",
		"Processor Generation": "M3 Pro",
		"Chipset":              "Apple M3 Pro",
		"Graphics Card":        "Apple M3 Pro GPU",
		"Screen Resolution":    "3456 x 2234 pixels",
		"Display Size":         "16.2 inches",
		"Battery Capacity":     "100 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Space Gray",
		"Product Weight":       "2.15 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.3",
		"RAM":                  "18 GB",
		"Weight":               "2.15 kg",
		"Display Type":         "Liquid Retina XDR",
		"Refresh Rate":         "120 Hz",
		"Screen Size":          "16.2 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "Yes",
		"Ram":                  "18 GB",
		"Wifi Support":         "Wi-Fi 6E",
		"Usb Type":             "Thunderbolt 4",
		"Battery":              "100 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "Apple M3 Pro",
		"Cooling Technology":   "Active",
		"Frequency (Hz)":       "120",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Professional",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "MagSafe",
		"Standby Time":            "Up to 22 hours",
		"Wireless Charging":       "No",
		"Resolution":              "3456 x 2234 pixels",
		"Display Characteristics": "Liquid Retina XDR, 1600 nits",
		"Processor Speed":         "12-core CPU, 18-core GPU",
		"Clock Feature":           "Up to 4.06 GHz",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FaceTime HD",
		"Audio Quality":           "Three microphones, six speakers",
		"Sensors":                 "Ambient light sensor",
		"Dimensions":              "14.01 x 9.77 x 0.66 inches",
		"Body Type":               "Aluminum",
		"Cooling System":          "Active",
		"Available Colors":        "Space Gray, Silver",
		"Special Features":        "Touch ID",
		// Major laptop specs added to database
		"Processor Cores":       "12",
		"Processor Threads":     "12",
		"RAM Speed":             "LPDDR5",
		"RAM Slots":             "Soldered",
		"RAM Expandable":        "No",
		"Storage Interface":     "NVMe PCIe",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "Thunderbolt 4",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "No",
		"USB Ports":   "3x Thunderbolt 4",
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
		if banglaValue, exists := amas.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Apple MacBook Pro 16-inch M3 Pro specifications seeded successfully")
	return nil
}
