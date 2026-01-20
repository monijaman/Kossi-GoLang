package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// LenovoIdeapadSlim5Seeder seeds specifications for Lenovo IdeaPad Slim 5
type LenovoIdeapadSlim5Seeder struct {
	BaseSeeder
}

// NewLenovoIdeapadSlim5Seeder creates a new Lenovo IdeaPad Slim 5 seeder
func NewLenovoIdeapadSlim5Seeder() *LenovoIdeapadSlim5Seeder {
	return &LenovoIdeapadSlim5Seeder{
		BaseSeeder: BaseSeeder{name: "Lenovo IdeaPad Slim 5 Specifications"},
	}
}

func (lis *LenovoIdeapadSlim5Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Lenovo":                        "লেনোভো",
		"IdeaPad Slim 5":                "আইডিয়াপ্যাড স্লিম ৫",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"AMD":                           "এএমডি",
		"Ryzen 5 7520U":                 "রাইজেন ৫ ৭৫২০ইউ",
		"7th Gen":                       "৭ম প্রজন্ম",
		"AMD Platform Controller Hub":   "এএমডি প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"AMD Radeon Graphics":           "এএমডি রেডিয়ন গ্রাফিক্স",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"56 Wh":                         "৫৬ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Gray":                          "গ্রে",
		"1.46 kg":                       "১.৪৬ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"16 GB":                         "১৬ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"AMD Ryzen 5":                   "এএমডি রাইজেন ৫",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Ultraportable":                        "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"IPS, 300 nits":                        "আইপিএস, ৩০০ নিটস",
		"2.8 GHz base / 4.3 GHz boost":         "২.৮ গিগাহার্টজ বেস / ৪.৩ গিগাহার্টজ বুস্ট",
		"16 MB cache":                          "১৬ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"313 x 220 x 16 mm":                    "৩১৩ x ২২০ x ১৬ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"6":                                    "৬",
		"12":                                   "১২",
		"3200 MHz":                             "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Standard":                             "স্ট্যান্ডার্ড",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Lenovo IdeaPad Slim 5
func (lis *LenovoIdeapadSlim5Seeder) Seed(db *gorm.DB) error {
	productSlug := "lenovo-ideapad-slim-5"
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
		"Brand":                "Lenovo",
		"Model Name":           "IdeaPad Slim 5",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "AMD",
		"Processor Model":      "Ryzen 5 7520U",
		"Processor Generation": "7th Gen",
		"Chipset":              "AMD Platform Controller Hub",
		"Graphics Card":        "AMD Radeon Graphics",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "56 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Gray",
		"Product Weight":       "1.46 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.2",
		"RAM":                  "16 GB",
		"Weight":               "1.46 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "56 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "AMD Ryzen 5",
		"Cooling Technology":   "Single Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "2.8 GHz base / 4.3 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "313 x 220 x 16 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Gray",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "6",
		"Processor Threads":     "12",
		"RAM Speed":             "3200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen3",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.0",
		"USB Ports":   "1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := lis.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Lenovo IdeaPad Slim 5 specifications seeded successfully")
	return nil
}
