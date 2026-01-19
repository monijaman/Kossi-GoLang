package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// HpSpectreFoldableSeeder seeds specifications for HP Spectre Foldable
type HpSpectreFoldableSeeder struct {
	BaseSeeder
}

// NewHpSpectreFoldableSeeder creates a new HP Spectre Foldable seeder
func NewHpSpectreFoldableSeeder() *HpSpectreFoldableSeeder {
	return &HpSpectreFoldableSeeder{
		BaseSeeder: BaseSeeder{name: "HP Spectre Foldable Specifications"},
	}
}

func (hss *HpSpectreFoldableSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"HP":                            "এইচপি",
		"Spectre Foldable":              "স্পেক্টর ফোল্ডেবল",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core Ultra 7 155H":             "কোর আল্ট্রা ৭ ১৫৫এইচ",
		"14th Gen":                      "১৪তম প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel Arc":                     "ইন্টেল আর্ক",
		"2560 x 1536 pixels":            "২৫৬০ x ১৫৩৬ পিক্সেল",
		"17.3 inches":                   "১৭.৩ ইঞ্চি",
		"70 Wh":                         "৭০ ডব্লিউএইচ",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Natural Silver":                "ন্যাচারাল সিলভার",
		"1.5 kg":                        "১.৫ কেজি",
		"2024":                          "২০২৪",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
		"1 TB":                          "১ টিবি",
		"Bluetooth 5.4":                 "ব্লুটুথ ৫.৪",
		"32 GB":                         "৩২ জিবি",
		"OLED":                          "ওএলইডি",
		"120 Hz":                        "১২০ হার্জ",
		"Yes":                           "হ্যাঁ",
		"USB-C":                         "ইউএসবি-সি",
		"Wi-Fi 7 (802.11be)":            "ওয়াই-ফাই ৭ (৮০২.১১বিই)",
		"USB 4.0 Gen 3":                 "ইউএসবি ৪.০ জেন ৩",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"1 TB SSD":                      "১ টিবি এসএসডি",
		"Intel Core Ultra 7":            "ইন্টেল কোর আল্ট্রা ৭",
		"Dual Fan":                      "ডুয়াল ফ্যান",
		"120":                           "১২০",
		"No":                            "না",
		"1 Year":                        "১ বছর",
		// New translations for added specs
		"Premium Foldable":             "প্রিমিয়াম ফোল্ডেবল",
		"Lithium-polymer":              "লিথিয়াম-পলিমার",
		"100W USB-C":                   "১০০ওয়াট ইউএসবি-সি",
		"10-12 hours":                  "১০-১২ ঘন্টা",
		"OLED, 400 nits":               "ওএলইডি, ৪০০ নিটস",
		"1.8 GHz base / 4.8 GHz boost": "১.৮ গিগাহার্টজ বেস / ৪.৮ গিগাহার্টজ বুস্ট",
		"24 MB cache":                  "২৪ এমবি ক্যাশ",
		"5MP HD":                       "৫এমপি এইচডি",
		"Quad Speakers":                "কোয়াড স্পিকার",
		"297 x 234 x 10.5 mm":          "২৯৭ x ২৩৪ x ১০.৫ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"16":                           "১৬",
		"22":                           "২২",
		"5600 MHz":                     "৫৬০০ মেগাহার্টজ",
		"Yes (up to 64GB)":             "হ্যাঁ (৬৪জিবি পর্যন্ত)",
		"NVMe PCIe Gen4":               "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)":               "হ্যাঁ (এম.২ স্লট)",
		"Shared":                       "শেয়ার্ড",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Premium":                      "প্রিমিয়াম",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"2x USB-C 4.0 Gen 3, 1x USB-A": "২x ইউএসবি-সি ৪.০ জেন ৩, ১x ইউএসবি-এ",
	}
}

// Seed implements the Seeder interface for HP Spectre Foldable
func (hss *HpSpectreFoldableSeeder) Seed(db *gorm.DB) error {
	productSlug := "hp-spectre-foldable"
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
		"Brand":                "HP",
		"Model Name":           "Spectre Foldable",
		"Operating System":     "Windows 11 Home",
		"Processor Brand":      "Intel",
		"Processor Model":      "Core Ultra 7 155H",
		"Processor Generation": "14th Gen",
		"Chipset":              "Intel Platform Controller Hub",
		"Graphics Card":        "Intel Arc",
		"Screen Resolution":    "2560 x 1536 pixels",
		"Display Size":         "17.3 inches",
		"Battery Capacity":     "70 Wh",
		"Build Material":       "Aluminum",
		"Color":                "Natural Silver",
		"Product Weight":       "1.5 kg",
		"Release Year":         "2024",
		"Warranty":             "1 Year International Warranty",
		"Storage Capacity":     "1 TB",
		"Bluetooth Version":    "Bluetooth 5.4",
		"RAM":                  "32 GB",
		"Weight":               "1.5 kg",
		"Display Type":         "OLED",
		"Refresh Rate":         "120 Hz",
		"Screen Size":          "17.3 inches",
		"Backlit Keyboard":     "Yes",
		"Audio Jack":           "USB-C",
		"Ram":                  "32 GB",
		"Wifi Support":         "Wi-Fi 7 (802.11be)",
		"Usb Type":             "USB 4.0 Gen 3",
		"Battery":              "70 Wh",
		"Gpu Type":             "Integrated",
		"Storage":              "1 TB SSD",
		"Cpu Type":             "Intel Core Ultra 7",
		"Cooling Technology":   "Dual Fan",
		"Frequency (Hz)":       "120",
		"App Control":          "No",
		"Warranty Period":      "1 Year",
		// New specs from specs.sql
		"Laptop Type":             "Premium Foldable",
		"Battery Type":            "Lithium-polymer",
		"Charging Speed":          "100W USB-C",
		"Standby Time":            "10-12 hours",
		"Wireless Charging":       "No",
		"Resolution":              "2560 x 1536 pixels",
		"Display Characteristics": "OLED, 400 nits",
		"Processor Speed":         "1.8 GHz base / 4.8 GHz boost",
		"Clock Feature":           "24 MB cache",
		"3.5mm Audio Jack":        "No",
		"Camera Features":         "5MP HD",
		"Audio Quality":           "Quad Speakers",
		"Sensors":                 "No",
		"Dimensions":              "297 x 234 x 10.5 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Dual Fan",
		"Available Colors":        "Natural Silver",
		"Special Features":        "TPM 2.0",
		// Major laptop specs added to database
		"Processor Cores":       "16",
		"Processor Threads":     "22",
		"RAM Speed":             "5600 MHz",
		"RAM Slots":             "2",
		"RAM Expandable":        "Yes (up to 64GB)",
		"Storage Interface":     "NVMe PCIe Gen4",
		"Storage Expandable":    "Yes (M.2 slot)",
		"Graphics VRAM":         "Shared",
		"Display Touch Support": "Yes",
		"Ethernet":              "No",
		"Thunderbolt Version":   "4",
		"SD Card Reader":        "No",
		"Keyboard Language":     "English/Bangla",
		"Build Standard":        "Premium",
		// Additional available keys
		"Touchscreen": "Yes",
		"HDMI Ports":  "HDMI 2.1",
		"USB Ports":   "2x USB-C 4.0 Gen 3, 1x USB-A",
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
		if banglaValue, exists := hss.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ HP Spectre Foldable specifications seeded successfully")
	return nil
}
