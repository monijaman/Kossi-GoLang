package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// AsusZenbook14OledSeeder seeds specifications for Asus ZenBook 14 OLED
type AsusZenbook14OledSeeder struct {
	BaseSeeder
}

// NewAsusZenbook14OledSeeder creates a new Asus ZenBook 14 OLED seeder
func NewAsusZenbook14OledSeeder() *AsusZenbook14OledSeeder {
	return &AsusZenbook14OledSeeder{
		BaseSeeder: BaseSeeder{name: "Asus ZenBook 14 OLED Specifications"},
	}
}

func (az *AsusZenbook14OledSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Asus":                          "এসুস",
		"ZenBook 14 OLED":               "জেনবুক ১৪ ওএলইডি",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i7-1355U":                 "কোর আই৭-১৩৫৫ইউ",
		"13th Gen":                      "১৩তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Iris Xe":                 "ইন্টেল আইরিস এক্সই",
		"2880 x 1800 pixels":            "২৮৮০ x ১৮০০ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"75 Wh":                         "৭৫ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Pine Grey":                     "পাইন গ্রে",
		"1.2 kg":                        "১.২ কেজি",
		"2023":                          "২০২৩",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"512 GB":                        "৫১২ জিবি",
		"Bluetooth 5.3":                 "ব্লুটুথ ৫.৩",
		"16 GB":                         "১৬ জিবি",
		"OLED":                          "ওএলইডি",
		"90 Hz":                         "৯০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E (802.11ax)":           "ওয়াই-ফাই ৬ই (৮০২.১১এক্স)",
		"USB 3.2 Gen 1":                 "ইউএসবি ৩.২ জেন ১",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"512 GB SSD":                    "৫১২ জিবি এসএসডি",
		"Intel Core i7":                 "ইন্টেল কোর আই৭",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"90":                            "৯০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		"Ultraportable":                 "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                   "লিথিয়াম-আয়ন",
		"65W USB-C":                     "৬৫ওয়াট ইউএসবি-সি",
		"10-12 hours":                   "১০-১২ ঘন্টা",
		"OLED, 400 nits":                "ওএলইডি, ৪০০ নিটস",
		"1.7 GHz base / 5.0 GHz boost":  "১.৭ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"12 MB cache":                   "১২ এমবি ক্যাশ",
		"1080p FHD":                     "১০৮০পি এফএইচডি",
		"Harman Kardon Speakers":         "হারম্যান কার্ডন স্পিকার",
		"313 x 220 x 15 mm":             "৩১৩ x ২২০ x ১৫ মিমি",
		"TPM 2.0":                       "টিপিএম ২.০",
		"10":                            "১০",
		"12":                            "১২",
		"4266 MHz":                      "৪২৬৬ মেগাহার্টজ",
		"Yes (up to 32GB)":              "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":                "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":                "হ্যাঁ (এম.২ স্লট)",
		"Shared":                        "শেয়ার্ড",
		"English/Bangla":                "ইংরেজি/বাংলা",
		"Standard":                      "স্ট্যান্ডার্ড",
		"HDMI 2.1":                      "এইচডিএমআই ২.১",
		"1x USB 3.2 Gen 1, 1x USB-C 3.2 Gen 2": "১x ইউএসবি ৩.২ জেন ১, ১x ইউএসবি-সি ৩.২ জেন ২",
		"Thunderbolt 4":                 "থান্ডারবোল্ট ৪",
	}
}

// Seed implements the Seeder interface for Asus ZenBook 14 OLED
func (az *AsusZenbook14OledSeeder) Seed(db *gorm.DB) error {
	productSlug := "asus-zenbook-14-oled"
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
		"Processor Cores":       786,
		"Processor Threads":     787,
		"RAM Speed":             788,
		"RAM Slots":             789,
		"RAM Expandable":        790,
		"Storage Interface":      791,
		"Storage Expandable":     792,
		"Graphics VRAM":         793,
		"Display Touch Support": 794,
		"Ethernet":              795,
		"Thunderbolt Version":   796,
		"SD Card Reader":         797,
		"Keyboard Language":     798,
		"Build Standard":        799,
		"Touchscreen": 129,
		"HDMI Ports":  167,
		"USB Ports":   173,
	}

	specs := map[string]string{
		"Brand":                "Asus",
		"Model Name":           "ZenBook 14 OLED",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core i7-1355U",
		"Processor Generation": "13th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel Iris Xe",
		"Screen Resolution":    "2880 x 1800 pixels",
		"Display Size":         "14 inches",
		"Battery Capacity":     "75 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Pine Grey",
		"Product Weight":       "1.2 kg",
		"Release Year":         "2023",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "512 GB",
		"Bluetooth Version":    "Bluetooth 5.3",
		"RAM":                  "16 GB",
		"Weight":               "1.2 kg",
		"Display Type":         "OLED",
		"Refresh Rate":         "90 Hz",
		"Screen Size":          "14 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "3.5mm Combo Jack",
		"Ram":                  "16 GB",
		"Wifi Support":         "Wi-Fi 6E (802.11ax)",
		"Usb Type":             "USB 3.2 Gen 1",
		"Battery":              "75 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "512 GB SSD",
		"Cpu Type":             "Intel Core i7",
		"Cooling Technology":   "Single Fan",
		"Frequency (Hz)":       "90",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2880 x 1800 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "1.7 GHz base / 5.0 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1080p FHD",
		"Audio Quality":           "Harman Kardon Speakers",
		"Sensors":                 "No",
		"Dimensions":              "313 x 220 x 15 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Pine Grey",
		"Special Features":        "TPM 2.0",
		"Processor Cores":       "10",
		"Processor Threads":     "12",
		"RAM Speed":             "4266 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 32GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":     "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "No",
		"Ethernet":              "No",
		"Thunderbolt Version":   "Thunderbolt 4",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Standard",
		"Touchscreen": "No",
		"HDMI Ports":  "HDMI 2.1",
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
		if banglaValue, exists := az.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Asus ZenBook 14 OLED specifications seeded successfully")
	return nil
}