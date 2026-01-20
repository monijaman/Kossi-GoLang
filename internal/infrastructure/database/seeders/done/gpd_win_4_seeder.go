package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// GpdWin4Seeder seeds specifications for GPD Win 4
type GpdWin4Seeder struct {
	BaseSeeder
}

// NewGpdWin4Seeder creates a new GPD Win 4 seeder
func NewGpdWin4Seeder() *GpdWin4Seeder {
	return &GpdWin4Seeder{
		BaseSeeder: BaseSeeder{name: "GPD Win 4 Specifications"},
	}
}

func (gw4s *GpdWin4Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"GPD":                          "জিপিডি",
		"Win 4":                        "উইন ৪",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 7 7840U":                "রাইজেন ৭ ৭৮৪০ইউ",
		"Zen 4":                        "জেন ৪",
		"AMD Radeon 780M":              "এএমডি রেডিওন ৭৮০এম",
		"1280 x 720 pixels":            "১২৮০ x ৭২০ পিক্সেল",
		"6 inches":                     "৬ ইঞ্চি",
		"50 Wh":                        "৫০ ডব্লিউএইচ",
		"Aluminum":                     "অ্যালুমিনিয়াম",
		"Black":                        "কালো",
		"0.6 kg":                       "০.৬ কেজি",
		"2023":                         "২০২৩",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"1 TB":                         "১ টিবি",
		"Bluetooth 5.2":                "ব্লুটুথ ৫.২",
		"32 GB":                        "৩২ জিবি",
		"IPS":                          "আইপিএস",
		"60 Hz":                        "৬০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6":                      "ওয়াই-ফাই ৬",
		"USB 3.2 Gen 2":                "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                    "ডেডিকেটেড",
		"1 TB SSD":                     "১ টিবি এসএসডি",
		"AMD Ryzen 7":                  "এএমডি রাইজেন ৭",
		"Dual Fan":                     "ডুয়াল ফ্যান",
		"60":                           "৬০",
		"No":                           "না",
		"1 Year":                       "১ বছর",
		"Ultraportable Gaming":         "আল্ট্রাপোর্টেবল গেমিং",
		"Lithium-polymer":              "লিথিয়াম-পলিমার",
		"65W USB-C":                    "৬৫ওয়াট ইউএসবি-সি",
		"4-5 hours gaming":             "৪-৫ ঘন্টা গেমিং",
		"IPS, 400 nits":                "আইপিএস, ৪০০ নিটস",
		"2.7 GHz base / 5.1 GHz boost": "২.৭ গিগাহার্টজ বেস / ৫.১ গিগাহার্টজ বুস্ট",
		"16 MB cache":                  "১৬ এমবি ক্যাশ",
		"2MP":                          "২এমপি",
		"Dual Speakers":                "ডুয়াল স্পিকার",
		"180 x 90 x 20 mm":             "১৮০ x ৯০ x ২০ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"8":                            "৮",
		"16":                           "১৬",
		"6400 MHz":                     "৬৪০০ মেগাহার্টজ",
		"PCIe Gen4":                    "পিসিআই জেন৪",
		"Yes (M.2)":                    "হ্যাঁ (এম.২)",
		"Dedicated (RDNA3)":            "ডেডিকেটেড (আরডিএনএ৩)",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"USB4, USB 3.2 Gen2, USB-C":    "ইউএসবি৪, ইউএসবি ৩.২ জেন২, ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for GPD Win 4
func (gw4s *GpdWin4Seeder) Seed(db *gorm.DB) error {
	productSlug := "gpd-win-4"
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
		"Model Name":              "Win 4",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 7 7840U",
		"Processor Generation":    "Zen 4",
		"Graphics Card":           "AMD Radeon 780M",
		"Screen Resolution":       "1280 x 720 pixels",
		"Display Size":            "6 inches",
		"Battery Capacity":        "50 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "0.6 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "32 GB",
		"Weight":                  "0.6 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "6 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "32 GB",
		"Wifi Support":            "Wi-Fi 6",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "50 Wh",
		"Gpu Type":                "Dedicated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "AMD Ryzen 7",
		"Cooling Technology":      "Dual Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable Gaming",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "4-5 hours gaming",
		"Wireless Charging":       "No",
		"Resolution":              "1280 x 720 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.7 GHz base / 5.1 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "2MP",
		"Audio Quality":           "Dual Speakers",
		"Sensors":                 "No",
		"Dimensions":              "180 x 90 x 20 mm",
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
		"Storage Expandable":      "Yes (M.2)",
		"Graphics VRAM":           "Dedicated (RDNA3)",
		"Display Touch Support":   "Yes",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "Yes",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.1",
		"USB Ports":               "USB4, USB 3.2 Gen2, USB-C",
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
		if banglaValue, exists := gw4s.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ GPD Win 4 specifications seeded successfully")
	return nil
}
