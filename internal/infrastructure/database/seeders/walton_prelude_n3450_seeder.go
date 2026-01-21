package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// WaltonPreludeN3450Seeder seeds specifications for Walton Prelude N3450
type WaltonPreludeN3450Seeder struct {
	BaseSeeder
}

// NewWaltonPreludeN3450Seeder creates a new Walton Prelude N3450 seeder
func NewWaltonPreludeN3450Seeder() *WaltonPreludeN3450Seeder {
	return &WaltonPreludeN3450Seeder{
		BaseSeeder: BaseSeeder{name: "Walton Prelude N3450 Specifications"},
	}
}

func (wps *WaltonPreludeN3450Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Walton":                        "ওয়ালটন",
		"Prelude N3450":                 "প্রেলুড এন৩৪৫০",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Celeron N3450":                 "সেলেরন এন৩৪৫০",
		"8th Gen":                       "৮ম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel HD Graphics":             "ইন্টেল এইচডি গ্রাফিক্স",
		"1366 x 768 pixels":             "১৩৬৬ x ৭৬৮ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"38 Wh":                         "৩৮ ডব্লিউএইচ",
		"Plastic":                       "প্লাস্টিক",
		"Black":                         "কালো",
		"1.3 kg":                        "১.৩ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"64 GB":                         "৬৪ জিবি",
		"Bluetooth 4.2":                 "ব্লুটুথ ৪.২",
		"4 GB":                          "৪ জিবি",
		"TN":                            "টিএন",
		"60 Hz":                         "৬০ হার্জ",
		"No":                            "না",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":            "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.0":                       "ইউএসবি ৩.০",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"64 GB eMMC":                    "৬৪ জিবি ইএমএমসি",
		"Intel Celeron":                 "ইন্টেল সেলেরন",
		"Passive Cooling":               "প্যাসিভ কুলিং",
		"60":                            "৬০",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Budget":                       "বাজেট",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"45W Adapter":                  "৪৫ওয়াট অ্যাডাপ্টার",
		"6-8 hours":                    "৬-৮ ঘন্টা",
		"TN, 200 nits":                 "টিএন, ২০০ নিটস",
		"1.1 GHz base / 2.2 GHz boost": "১.১ গিগাহার্টজ বেস / ২.২ গিগাহার্টজ বুস্ট",
		"2 MB cache":                   "২ এমবি ক্যাশ",
		"Mono Speakers":                "মোনো স্পিকার",
		"330 x 230 x 20 mm":            "৩৩০ x ২৩০ x ২০ মিমি",
		"Basic":                        "বেসিক", // Major laptop specs translations
		"4":                            "৪",
		"2133 MHz":                     "২১৩৩ মেগাহার্টজ",

		"eMMC":                   "ইএমএমসি",
		"Shared":                 "শেয়ার্ড",
		"English/Bangla":         "ইংরেজি/বাংলা",
		"Standard":               "স্ট্যান্ডার্ড",
		"1x USB 3.0, 1x USB 2.0": "১x ইউএসবি ৩.০, ১x ইউএসবি ২.০"}
}

// Seed implements the Seeder interface for Walton Prelude N3450
func (wps *WaltonPreludeN3450Seeder) Seed(db *gorm.DB) error {
	productSlug := "walton-prelude-n3450"
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
		"Brand":                "Walton",
		"Model Name":           "Prelude N3450",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Celeron N3450",
		"Processor Generation": "8th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel HD Graphics",
		"Screen Resolution":    "1366 x 768 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "38 Wh",
		"Build Material":       "Plastic",
		"Color":                "Black",
		"Product Weight":       "1.3 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "64 GB",
		"Bluetooth Version":    "Bluetooth 4.2",
		"RAM":                  "4 GB",
		"Weight":               "1.3 kg",
		"Display Type":         "TN",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "No",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "4 GB",
		"Wifi Support":         "Wi-Fi 5 (802.11ac)",
		"Usb Type":             "USB 3.0",
		"Battery":              "38 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "64 GB eMMC",
		"Cpu Type":             "Intel Celeron",
		"Cooling Technology":   "Passive Cooling",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Budget",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W Adapter",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1366 x 768 pixels",
		"Display Characteristics": "TN, 200 nits",
		"Processor Speed":         "1.1 GHz base / 2.2 GHz boost",
		"Clock Feature":           "2 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "VGA",
		"Audio Quality":           "Mono Speakers",
		"Sensors":                 "No",
		"Dimensions":              "330 x 230 x 20 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Passive Cooling",
		"Available Colors":        "Black",
		"Special Features":        "Basic",
		// Major laptop specs added to database
		"Processor Cores":       "4",
		"Processor Threads":     "4",
		"RAM Speed":             "2133 MHz",
		"RAM Slots":             "1",
		"RAM Expandable":        "No",
		"Storage Interface":     "eMMC",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "VGA",
		"USB Ports":   "1x USB 3.0, 1x USB 2.0",
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
		if banglaValue, exists := wps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Walton Prelude N3450 specifications seeded successfully")
	return nil
}
