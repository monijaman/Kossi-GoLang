package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AppleMacbookPro13M2Seeder seeds specifications for Apple MacBook Pro 13-inch M2
type AppleMacbookPro13M2Seeder struct {
	BaseSeeder
}

// NewAppleMacbookPro13M2Seeder creates a new Apple MacBook Pro 13-inch M2 seeder
func NewAppleMacbookPro13M2Seeder() *AppleMacbookPro13M2Seeder {
	return &AppleMacbookPro13M2Seeder{
		BaseSeeder: BaseSeeder{name: "Apple MacBook Pro 13-inch M2 Specifications"},
	}
}

func (amas *AppleMacbookPro13M2Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Apple":                         "অ্যাপল",
		"MacBook Pro 13-inch M2":        "ম্যাকবুক প্রো ১৩-ইঞ্চি এম২",
		"macOS Monterey":                "ম্যাকওএস মন্টেরে",
		"M2":                            "এম২",
		"Apple M2 GPU":                  "অ্যাপল এম২ জিপিইউ",
		"2560 x 1600 pixels":            "২৫৬০ x ১৬০০ পিক্সেল",
		"13.3 inches":                   "১৩.৩ ইঞ্চি",
		"58.2 Wh":                       "৫৮.২ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Space Gray":                    "স্পেস গ্রে",
		"1.4 kg":                        "১.৪ কেজি",
		"2022":                          "২০২২",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"256 GB":                        "২৫৬ জিবি",
		"Bluetooth 5.0":                 "ব্লুটুথ ৫.০",
		"8 GB":                          "৮ জিবি",
		"Retina":                        "রেটিনা",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"No":                            "না",
		"Wi-Fi 6":                       "ওয়াই-ফাই ৬",
		"Thunderbolt 3":                 "থান্ডারবোল্ট ৩",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"256 GB SSD":                    "২৫৬ জিবি এসএসডি",
		"Apple M2":                      "অ্যাপল এম২",
		"Active":                        "অ্যাকটিভ",
		"60":                            "৬০",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Professional":                    "প্রফেশনাল",
		"Lithium-polymer":                 "লিথিয়াম-পলিমার",
		"MagSafe":                         "ম্যাগসেফ",
		"Up to 20 hours":                  "আপ টু ২০ ঘন্টা",
		"Retina, 500 nits":                "রেটিনা, ৫০০ নিটস",
		"8-core CPU, 10-core GPU":         "৮-কোর সিপিইউ, ১০-কোর জিপিইউ",
		"Up to 3.49 GHz":                  "আপ টু ৩.৪৯ গিগাহার্টজ",
		"720p FaceTime HD":                "৭২০পি ফেসটাইম এইচডি",
		"Three microphones, two speakers": "থ্রি মাইক্রোফোন, টু স্পিকার",
		"Ambient light sensor":            "অ্যাম্বিয়েন্ট লাইট সেন্সর",
		"11.97 x 8.36 x 0.61 inches":      "১১.৯৭ x ৮.৩৬ x ০.৬১ ইঞ্চি",
		"Space Gray, Silver":              "স্পেস গ্রে, সিলভার",
		"Touch Bar":                       "টাচ বার",
		// Major laptop specs translations
		"8":         "৮",
		"LPDDR5":    "এলপিডিডিআর৫",
		"Soldered":  "সোল্ডার্ড",
		"NVMe PCIe": "এনভিএমই পিসিআই",
		"Shared":    "শেয়ার্ড",

		"English/Bangla":   "ইংরেজি/বাংলা",
		"Premium":          "প্রিমিয়াম",
		"2x Thunderbolt 3": "২x থান্ডারবোল্ট ৩",
	}
}

// Seed implements the Seeder interface for Apple MacBook Pro 13-inch M2
func (amas *AppleMacbookPro13M2Seeder) Seed(db *gorm.DB) error {
	productSlug := "apple-macbook-pro-13-inch-m2"
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
		"Model Name":           "MacBook Pro 13-inch M2",
		"Operating System":     "macOS Monterey",
		"Processor Brand":      "Apple",
		"Processor Model":      "M2",
		"Processor Generation": "M2",
		"Chipset":              "Apple M2",
		"Graphics Card":        "Apple M2 GPU",
		"Screen Resolution":    "2560 x 1600 pixels",
		"Display Size":         "13.3 inches",
		"Battery Capacity":     "58.2 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Space Gray",
		"Product Weight":       "1.4 kg",
		"Release Year":         "2022",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "256 GB",
		"Bluetooth Version":    "Bluetooth 5.0",
		"RAM":                  "8 GB",
		"Weight":               "1.4 kg",
		"Display Type":         "Retina",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "13.3 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "Yes",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 6",
		"Usb Type":             "Thunderbolt 3",
		"Battery":              "58.2 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "256 GB SSD",
		"Cpu Type":             "Apple M2",
		"Cooling Technology":   "Active",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Professional",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "MagSafe",
		"Standby Time":            "Up to 20 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1600 pixels",
		"Display Characteristics": "Retina, 500 nits",
		"Processor Speed":         "8-core CPU, 10-core GPU",
		"Clock Feature":           "Up to 3.49 GHz",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p FaceTime HD",
		"Audio Quality":           "Three microphones, two speakers",
		"Sensors":                 "Ambient light sensor",
		"Dimensions":              "11.97 x 8.36 x 0.61 inches",
		"Body Type":               "Aluminum",
		"Cooling System":          "Active",
		"Available Colors":        "Space Gray, Silver",
		"Special Features":        "Touch Bar",
		// Major laptop specs added to database
		"Processor Cores":       "8",
		"Processor Threads":     "8",
		"RAM Speed":             "LPDDR5",
		"RAM Slots":             "Soldered",
		"RAM Expandable":        "No",
		"Storage Interface":     "NVMe PCIe",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "Thunderbolt 3",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "No",
		"USB Ports":   "2x Thunderbolt 3",
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

	log.Printf("✅ Apple MacBook Pro 13-inch M2 specifications seeded successfully")
	return nil
}
