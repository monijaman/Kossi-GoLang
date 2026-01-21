package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// DellXPS159530Seeder seeds specifications for Dell XPS 15 9530
type DellXPS159530Seeder struct {
	BaseSeeder
}

// NewDellXPS159530Seeder creates a new Dell XPS 15 9530 seeder
func NewDellXPS159530Seeder() *DellXPS159530Seeder {
	return &DellXPS159530Seeder{
		BaseSeeder: BaseSeeder{name: "Dell XPS 15 9530 Specifications"},
	}
}

func (dxps *DellXPS159530Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Dell":                          "ডেল",
		"XPS 15 9530":                   "এক্সপিএস ১৫ ৯৫৩০",
		"Windows 11 Pro":                "উইন্ডোজ ১১ প্রো",
		"Intel":                         "ইন্টেল",
		"Core i9-13900H":                "কোর আই৯-১৩৯০০এইচ",
		"13th Gen":                      "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"NVIDIA GeForce RTX 4070":       "এনভিডিয়া জিফোর্স আরটিএক্স ৪০৭০",
		"3456 x 2160 pixels":            "৩৪৫৬ x ২১৬০ পিক্সেল",
		"15.6 inches":                   "১৫.৬ ইঞ্চি",
		"86 Wh":                         "৮৬ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Platinum":                      "প্লাটিনাম",
		"1.96 kg":                       "১.৯৬ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                          "১ টেরাবাইট",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"32 GB":                         "৩২ জিবি",
		"OLED":                          "ওএলইডি",
		"120 Hz":                        "১২০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 2":                 "ইউএসবি ৩.২ জেন ২",
		"Dedicated":                     "ডেডিকেটেড",
		"1 TB SSD":                      "১ টেরাবাইট এসএসডি",
		"Intel Core i9":                 "ইন্টেল কোর আই৯",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"120":                           "১২০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Premium Ultraportable":                "প্রিমিয়াম আল্ট্রাপোর্টেবল",
		"Lithium-polymer":                      "লিথিয়াম-পলিমার",
		"130W USB-C":                           "১৩০ওয়াট ইউএসবি-সি",
		"6-8 hours":                            "৬-৮ ঘন্টা",
		"OLED, 400 nits":                       "ওএলইডি, ৪০০ নিটস",
		"2.6 GHz base / 5.4 GHz boost":         "২.৬ গিগাহার্টজ বেস / ৫.৪ গিগাহার্টজ বুস্ট",
		"24 MB cache":                          "২৪ এমবি ক্যাশ",
		"1080p FHD":                            "১০৮০পি এফএইচডি",
		"Quad Speakers":                        "কোয়াড স্পিকার",
		"344.0 x 230.1 x 15.0 mm":              "৩৪৪.০ x ২৩০.১ x ১৫.০ মিমি",
		"TPM 2.0":                              "টিপিএম ২.০",
		"14":                                   "১৪",
		"20":                                   "২০",
		"5200 MHz":                             "৫২০০ মেগাহার্টজ",
		"Yes (up to 64GB)":                     "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                       "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                       "হ্যাঁ (এম.২ স্লট)",
		"8 GB GDDR6":                           "৮ জিবি জিডিডিআর৬",
		"English/Bangla":                       "ইংরেজি/বাংলা",
		"Premium":                              "প্রিমিয়াম",
		"HDMI 2.1":                             "এইচডিএমআই ২.১",
		"2x Thunderbolt 4, 1x USB-C 3.2 Gen 2": "২x থান্ডারবোল্ট ৪, ১x ইউএসবি-সি ৩.২ জেন ২"}
}

// Seed implements the Seeder interface for Dell XPS 15 9530
func (dxps *DellXPS159530Seeder) Seed(db *gorm.DB) error {
	productSlug := "dell-xps-15-9530"
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
		"Brand":                "Dell",
		"Model Name":           "XPS 15 9530",
		"Operating System":     "Windows 11 Pro",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i9-13900H",
		"Processor Generation": "13th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "NVIDIA GeForce RTX 4070",
		"Screen Resolution":    "3456 x 2160 pixels",
		"Display Size":         "15.6 inches",
		"Battery Capacity":     "86 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Platinum",
		"Product Weight":       "1.96 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "1 TB",
		"Bluetooth Version":    "Bluetooth 5.3",
		"RAM":                  "32 GB",
		"Weight":               "1.96 kg",
		"Display Type":         "OLED",
		"Refresh Rate":         "120 Hz",
		"Screen Size":          "15.6 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "32 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 2",
		"Battery":              "86 Wh",
		"Gpu Type":             "Dedicated",
		"Storage":              "1 TB SSD",
		"Cpu Type":             "Intel Core i9",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "120",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Premium Ultraportable",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "130W USB-C",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "3456 x 2160 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "2.6 GHz base / 5.4 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Quad Speakers",
		"Sensors":                 "Fingerprint Reader",
		"Dimensions":              "344.0 x 230.1 x 15.0 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Platinum",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "14",
		"Processor Threads":     "20",
		"RAM Speed":             "5200 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 64GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "8 GB GDDR6",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "4",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "2x Thunderbolt 4, 1x USB-C 3.2 Gen 2",
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
		if banglaValue, exists := dxps.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Dell XPS 15 9530 specifications seeded successfully")
	return nil
}
