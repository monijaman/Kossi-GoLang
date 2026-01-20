package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// GetacB360Seeder seeds specifications for Getac B360
type GetacB360Seeder struct {
	BaseSeeder
}

// NewGetacB360Seeder creates a new Getac B360 seeder
func NewGetacB360Seeder() *GetacB360Seeder {
	return &GetacB360Seeder{
		BaseSeeder: BaseSeeder{name: "Getac B360 Specifications"},
	}
}

func (gbs *GetacB360Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Getac":                         "গেটাক",
		"B360":                          "বি৩৬০",
		"Windows 10 Pro":                "উইন্ডোজ ১০ প্রো",
		"Intel":                         "ইন্টেল",
		"Celeron 6305":                  "সেলেরন ৬৩০৫",
		"11th Gen":                      "১১তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD":                     "ইন্টেল ইউএইচডি",
		"1920 x 1080 pixels":            "১৯২০ x ১০৮০ পিক্সেল",
		"11.6 inches":                   "১১.৬ ইঞ্চি",
		"37 Wh":                         "৩৭ ডব্লিউএইচ",
		"Magnesium Alloy":               "ম্যাগনেসিয়াম অ্যালয়",
		"Black":                         "কালো",
		"1.5 kg":                        "১.৫ কেজি",
		"2021":                          "২০২১",
		"3 Year Warranty":               "৩ বছর ওয়ারেন্টি",
		"128 GB":                        "১২৮ জিবি",
		"Bluetooth 5.0":                 "ব্লুটুথ ৫.০",
		"8 GB":                          "৮ জিবি",
		"IPS":                           "আইপিএস",
		"60 Hz":                         "৬০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"128 GB SSD":                    "১২৮ জিবি এসএসডি",
		"Intel Celeron":                 "ইন্টেল সেলেরন",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"No":                            "না",
		"3 Years":                       "৩ বছর",
		// New translations for added specs
		"Rugged Tablet":                        "রাগড ট্যাবলেট",
		"Lithium-ion":                          "লিথিয়াম-আয়ন",
		"45W USB-C":                            "৪৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                           "৮-১০ ঘন্টা",
		"IPS, 1000 nits":                       "আইপিএস, ১০০০ নিটস",
		"1.8 GHz base / 3.0 GHz boost":         "১.৮ গিগাহার্টজ বেস / ৩.০ গিগাহার্টজ বুস্ট",
		"4 MB cache":                           "৪ এমবি ক্যাশ",
		"720p HD":                              "৭২০পি এইচডি",
		"Stereo Speakers":                      "স্টেরিও স্পিকার",
		"300 x 200 x 25 mm":                    "৩০০ x ২০০ x ২৫ মিমি",
		"MIL-STD-810G":                         "এমআইএল-এসটিডি-৮১০জি",
		"2":                                    "২",
		"3200 MHz":                             "৩২০০ মেগাহার্টজ",
		"Yes (up to 32GB)":                     "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen3":                       "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"Shared":                               "শেয়ার্ড",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Rugged":                               "রাগড",
		"HDMI 2.0":                             "এইচডিএমআই ২.০",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
	}
}

// Seed implements the Seeder interface for Getac B360
func (gbs *GetacB360Seeder) Seed(db *gorm.DB) error {
	productSlug := "getac-b360"
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
		"Brand":                "Getac",
		"Model Name":           "B360",
		"Operating System":     "Windows 10 Pro",
		"Processor Brand":      "Intel",
		"Processor Model":      "Celeron 6305",
		"Processor Generation": "11th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel UHD",
		"Screen Resolution":    "1920 x 1080 pixels",
		"Display Size":         "11.6 inches",
		"Battery Capacity":     "37 Wh",
		"Build Material":       "Magnesium Alloy",
		"Color":                "Black",
		"Product Weight":       "1.5 kg",
		"Release Year":         "2021",
		"Warranty":             "3 Year Warranty",
		"Storage Capacity":     "128 GB",
		"Bluetooth Version":    "Bluetooth 5.0",
		"RAM":                  "8 GB",
		"Weight":               "1.5 kg",
		"Display Type":         "IPS",
		"Refresh Rate":         "60 Hz",
		"Screen Size":          "11.6 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "8 GB",
		"Wifi Support":         "Wi-Fi 6 (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "37 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "128 GB SSD",
		"Cpu Type":             "Intel Celeron",
		"Cooling Technology":   "Single Fan",
		"Frequency (Hz)":       "60",
		"App Control":          "No",
		"Warranty Period":      "3 Years",
		// New specs from specs.sql
		"Laptop Type":             "Rugged Tablet",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 1000 nits",
		"Processor Speed":         "1.8 GHz base / 3.0 GHz boost",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "300 x 200 x 25 mm",
		"Body Type":               "Magnesium Alloy",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Black",
		"Special Features":        "MIL-STD-810G",
		// Major laptop specs added to database
		"Processor Cores":       "2",
		"Processor Threads":     "2",
		"RAM Speed":             "3200 MHz",
		"RAM Slots":             "1",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen3",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "Yes",
		"Ethernet":              "Yes",
		"Thunderbolt Version":   "No",
		"SD Card Reader":        "Yes",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Rugged",
		// Additional available keys
		"Touchscreen": "Yes",
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
		if banglaValue, exists := gbs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Getac B360 specifications seeded successfully")
	return nil
}
