package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// GpdWinMax2Seeder seeds specifications for GPD Win Max 2
type GpdWinMax2Seeder struct {
	BaseSeeder
}

// NewGpdWinMax2Seeder creates a new GPD Win Max 2 seeder
func NewGpdWinMax2Seeder() *GpdWinMax2Seeder {
	return &GpdWinMax2Seeder{
		BaseSeeder: BaseSeeder{name: "GPD Win Max 2 Specifications"},
	}
}

func (gwm2s *GpdWinMax2Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"GPD":                          "জিপিডি",
		"Win Max 2":                    "উইন ম্যাক্স ২",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 7 8840U":                "রাইজেন ৭ ৮৮৪০ইউ",
		"Zen 4":                        "জেন ৪",
		"AMD Radeon 780M":              "এএমডি রেডিওন ৭৮০এম",
		"1920 x 1200 pixels":           "১৯২০ x ১২০০ পিক্সেল",
		"10.1 inches":                  "১০.১ ইঞ্চি",
		"67 Wh":                        "৬৭ ডব্লিউএইচ",
		"Aluminum":                     "অ্যালুমিনিয়াম",
		"Black":                        "কালো",
		"0.97 kg":                      "০.৯৭ কেজি",
		"2024":                         "২০২৪",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"2 TB":                         "২ টিবি",
		"Bluetooth 5.2":                "ব্লুটুথ ৫.২",
		"64 GB":                        "৬৪ জিবি",
		"IPS":                          "আইপিএস",
		"60 Hz":                        "৬০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6":                      "ওয়াই-ফাই ৬",
		"USB 3.2 Gen 2":                "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                    "ডেডিকেটেড",
		"2 TB SSD":                     "২ টিবি এসএসডি",
		"AMD Ryzen 7":                  "এএমডি রাইজেন ৭",
		"Dual Fan":                     "ডুয়াল ফ্যান",
		"60":                           "৬০",
		"No":                           "না",
		"1 Year":                       "১ বছর",
		"Ultraportable Gaming":         "আল্ট্রাপোর্টেবল গেমিং",
		"Lithium-polymer":              "লিথিয়াম-পলিমার",
		"100W USB-C":                   "১০০ওয়াট ইউএসবি-সি",
		"5-6 hours gaming":             "৫-৬ ঘন্টা গেমিং",
		"IPS, 400 nits":                "আইপিএস, ৪০০ নিটস",
		"2.3 GHz base / 5.1 GHz boost": "২.৩ গিগাহার্টজ বেস / ৫.১ গিগাহার্টজ বুস্ট",
		"24 MB cache":                  "২৪ এমবি ক্যাশ",
		"2MP":                          "২এমপি",
		"4 Speakers":                   "৪ স্পিকার",
		"260 x 115 x 25 mm":            "২৬০ x ১১৫ x ২৫ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"8":                            "৮",
		"16":                           "১৬",
		"6400 MHz":                     "৬৪০০ মেগাহার্টজ",
		"Yes (up to 64GB)":             "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"PCIe Gen4":                    "পিসিআই জেন৪",
		"Yes (M.2 2280/2230)":          "হ্যাঁ (এম.২ ২২৮০/২২৩০)",
		"Dedicated (RDNA3)":            "ডেডিকেটেড (আরডিএনএ৩)",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"USB4, 2x USB 3.2 Gen2, USB-C": "ইউএসবি৪, ২x ইউএসবি ৩.২ জেন২, ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for GPD Win Max 2
func (gwm2s *GpdWinMax2Seeder) Seed(db *gorm.DB) error {
	productSlug := "gpd-win-max-2"
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
		"Brand":                   "GPD",
		"Model Name":              "Win Max 2",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 7 8840U",
		"Processor Generation":    "Zen 4",
		"Graphics Card":           "AMD Radeon 780M",
		"Screen Resolution":       "1920 x 1200 pixels",
		"Display Size":            "10.1 inches",
		"Battery Capacity":        "67 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "0.97 kg",
		"Release Year":            "2024",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "2 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "64 GB",
		"Weight":                  "0.97 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "10.1 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "64 GB",
		"Wifi Support":            "Wi-Fi 6",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "67 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "2 TB SSD",
		"Cpu Type":                "AMD Ryzen 7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable Gaming",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "100W USB-C",
		"Standby Time":            "5-6 hours gaming",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.3 GHz base / 5.1 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "2MP",
		"Audio Quality":           "4 Speakers",
		"Sensors":                 "No",
		"Dimensions":              "260 x 115 x 25 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "8",
		"Processor Threads":       "16",
		"RAM Speed":               "6400 MHz",
		"RAM Slots":               "Soldered",
		"RAM Expandable":          "No",
		"Storage Interface":       "PCIe Gen4",
		"Storage Expandable":      "Yes (M.2 2280/2230)",
		"Graphics VRAM":           "Dedicated (RDNA3)",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "Yes",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "USB4, 2x USB 3.2 Gen2, USB-C",
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
		if banglaValue, exists := gwm2s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ GPD Win Max 2 specifications seeded successfully")
	return nil
}
