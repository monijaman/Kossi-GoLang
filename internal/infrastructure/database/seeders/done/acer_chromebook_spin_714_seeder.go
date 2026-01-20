package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AcerChromebookSpin714Seeder seeds specifications for Acer Chromebook Spin 714
type AcerChromebookSpin714Seeder struct {
	BaseSeeder
}

// NewAcerChromebookSpin714Seeder creates a new Acer Chromebook Spin 714 seeder
func NewAcerChromebookSpin714Seeder() *AcerChromebookSpin714Seeder {
	return &AcerChromebookSpin714Seeder{
		BaseSeeder: BaseSeeder{name: "Acer Chromebook Spin 714 Specifications"},
	}
}

func (acss *AcerChromebookSpin714Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Acer":                          "এসার",
		"Chromebook Spin 714":           "ক্রোমবুক স্পিন ৭১৪",
		"Chrome OS":                     "ক্রোম ওএস",
		"Intel":                         "ইন্টেল",
		"Celeron N5100":                 "সেলেরন এন৫১০০",
		"11th Gen":                      "১১তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD Graphics":            "ইন্টেল ইউএইচডি গ্রাফিক্স",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"54 Wh":                         "৫৪ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Silver":                        "সিলভার",
		"1.4 kg":                        "১.৪ কেজি",
		"2024":                          "২০২৪",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"128 GB":                        "১২৮ জিবি",
		"Bluetooth 5.1":                 "ব্লুটুথ ৫.১",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"128 GB SSD":                    "১২৮ জিবি এসএসডি",
		"Intel Celeron":                 "ইন্টেল সেলেরন",
		"Passive":                       "প্যাসিভ",
		"60":                            "৬০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Convertible Chromebook":               "কনভার্টিবল ক্রোমবুক",
		"Lithium-polymer":                      "লিথিয়াম-পলিমার",
		"65W USB-C":                            "৬৫ওয়াট ইউএসবি-সি",
		"10-12 hours":                          "১০-১২ ঘন্টা",
		"IPS, 300 nits":                        "আইপিএস, ৩০০ নিটস",
		"1.1 GHz base / 2.8 GHz boost":         "১.১ গিগাহার্টজ বেস / ২.৮ গিগাহার্টজ বুস্ট",
		"4 MB cache":                           "৪ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"321 x 214 x 15.9 mm":                  "৩২১ x ২১৪ x ১৫.৯ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"4":                                    "৪",
		"2933 MHz":                             "২৯৩৩ মেগাহার্টজ",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Premium":                              "প্রিমিয়াম",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ২, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Acer Chromebook Spin 714
func (acss *AcerChromebookSpin714Seeder) Seed(db *gorm.DB) error {
	productSlug := "acer-chromebook-spin-714"
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
		"Brand":                "Acer",
		"Model Name":           "Chromebook Spin 714",
		"Operating System":     "Chrome OS",
		"Processor Brand":      "Intel",
		"Processor Model":      "Celeron N5100",
		"Processor Generation": "11th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel UHD Graphics",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "54 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Silver",
		"Product Weight":       "1.4 kg",
		"Release Year":         "2024",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "128 GB",
		"Bluetooth Version":    "Bluetooth 5.1",
		"RAM":                  "8 GB",
		"Weight":               "1.4 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 2",
		"Battery":              "54 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "128 GB SSD",
		"Cpu Type":             "Intel Celeron",
		"Cooling Technology":   "Passive",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Convertible Chromebook",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 300 nits",
		"Processor Speed":         "1.1 GHz base / 2.8 GHz boost",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "321 x 214 x 15.9 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Passive",
		"Available Colors":        "Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "4",
		"Processor Threads":     "4",
		"RAM Speed":             "2933 MHz",
		"RAM Slots":             "1",
		"RAM Expandable":        "No",
		"Storage Interface":     "NVMe PCIe Gen3",
		"Storage Expandable":    "No",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "Yes",
		"Ethernet":              "No",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "Yes",
		"HDMI Ports":  "HDMI 2.0",
		"USB Ports":   "1x USB 3.2 Gen 2, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := acss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Acer Chromebook Spin 714 specifications seeded successfully")
	return nil
}
