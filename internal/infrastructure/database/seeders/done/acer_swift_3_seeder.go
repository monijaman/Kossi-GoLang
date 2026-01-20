package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

type AcerSwift3Seeder struct {
	BaseSeeder
}

func NewAcerSwift3Seeder() *AcerSwift3Seeder {
	return &AcerSwift3Seeder{
		BaseSeeder: BaseSeeder{name: "Acer Swift 3 Specifications"},
	}
}

func (ass *AcerSwift3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Acer":                          "এসার",
		"Swift 3":                       "সুইফট ৩",
		"Ultraportable":                 "আল্ট্রাপোর্টেবল",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Core i5-12450H":                "কোর আই৫-১২৪৫০এইচ",
		"12th Gen":                      "১২তম প্রজন্ম",
		"8":                             "৮",
		"12":                            "১২",
		"2.0 GHz":                       "২.০ গিগাহার্টজ",
		"4.4 GHz":                       "৪.৪ গিগাহার্টজ",
		"12 MB":                         "১২ এমবি",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"Intel Iris Xe Graphics":        "ইন্টেল আইরিস এক্সই গ্রাফিক্স",
		"Shared":                        "শেয়ার্ড",
		"14 inches":                     "১৪ ইঞ্চি",
		"1920 x 1200 pixels":            "১৯২০ x ১২০০ পিক্সেল",
		"IPS":                           "আইপিএস",
		"60Hz":                          "৬০ হার্জ",
		"16:10":                         "১৬:১০",
		"400 nits":                      "৪০০ নিটস",
		"100% sRGB":                     "১০০% এসআরজিবি",
		"No":                            "না",
		"Glossy":                        "গ্লসি",
		"DDR4":                          "ডিডিআর৪",
		"8 GB":                          "৮ জিবি",
		"3200 MHz":                      "৩২০০ মেগাহার্টজ",
		"2":                             "২",
		"Yes (up to 32GB)":              "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"SSD":                           "এসএসডি",
		"512 GB":                        "৫১২ জিবি",
		"NVMe PCIe Gen3":                "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)":                "হ্যাঁ (এম.২ স্লট)",
		"None":                          "কোনটি নয়",
		"Standard":                      "স্ট্যান্ডার্ড",
		"Yes":                           "হ্যাঁ",
		"English/Bangla":                "ইংরেজি/বাংলা",
		"Precision Touchpad":            "প্রিসিশন টাচপ্যাড",
		"1080p FHD":                     "১০৮০পি এফএইচডি",
		"Privacy Shutter":               "প্রাইভেসি শাটার",
		"Dual Microphone":               "ডুয়াল মাইক্রোফোন",
		"Acer TrueHarmony":              "এসার ট্রু হারমোনি",
		"Stereo Speakers":               "স্টেরিও স্পিকার",
		"Wi-Fi 6 (802.11ax)":            "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"Bluetooth 5.2":                 "ব্লুটুথ ৫.২",
		"1x USB 3.2 Gen 1":              "১x ইউএসবি ৩.২ জেন ১",
		"1x USB-C (USB 3.2 Gen 2)":      "১x ইউএসবি-সি (ইউএসবি ৩.২ জেন ২)",
		"HDMI 2.0":                      "এইচডিএমআই ২.০",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"50 Wh":                         "৫০ ডব্লিউএইচ",
		"8-10 hours":                    "৮-১০ ঘন্টা",
		"65W":                           "৬৫ওয়াট",
		"Aluminum":                      "অ্যালুমিনিয়াম",
		"Silver":                        "সিলভার",
		"312.2 x 217.9 x 15.9 mm":       "৩১২.২ x ২১৭.৯ x ১৫.৯ মিমি",
		"1.2 kg":                        "১.২ কেজি",
		"Premium":                       "প্রিমিয়াম",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"TPM 2.0":                       "টিপিএম ২.০",
		"2024":                          "২০২৪",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
	}
}

func (ass *AcerSwift3Seeder) Seed(db *gorm.DB) error {
	productSlug := "acer-swift-3"
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}
		return err
	}

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
		"Brand": "Acer", "Model Name": "Swift 3", "Laptop Type": "Ultraportable",
		"Operating System": "Windows 11 Home", "Processor Brand": "Intel",
		"Processor Model": "Core i5-12450H", "Processor Generation": "12th Gen",
		"Processor Cores": "8", "Processor Threads": "12",
		"Processor Base Clock": "2.0 GHz", "Processor Boost Clock": "4.4 GHz",
		"Processor Cache": "12 MB", "Chipset": "Intel Platform Controller Hub",
		"Graphics Type": "Integrated", "Graphics Brand": "Intel",
		"Graphics Card": "Intel Iris Xe Graphics", "Graphics VRAM": "Shared",
		"Display Size": "14 inches", "Screen Resolution": "1920 x 1200 pixels",
		"Display Panel Type": "IPS", "Display Refresh Rate": "60Hz",
		"Display Aspect Ratio": "16:10", "Display Brightness": "400 nits",
		"Display Color Gamut": "100% sRGB", "Display Touch Support": "No",
		"Display Surface": "Glossy", "RAM Type": "DDR4", "RAM Capacity": "8 GB",
		"RAM Speed": "3200 MHz", "RAM Slots": "2",
		"RAM Expandable": "Yes (up to 32GB)", "Storage Type": "SSD",
		"Storage Capacity": "512 GB", "Storage Interface": "NVMe PCIe Gen3",
		"Storage Expandable": "Yes (M.2 slot)", "Secondary Storage Type": "None",
		"Secondary Storage Capacity": "None", "Keyboard Type": "Standard",
		"Backlit Keyboard": "Yes", "Keyboard Language": "English/Bangla",
		"Touchpad Type": "Precision Touchpad", "Fingerprint Reader": "No",
		"Webcam Resolution": "1080p FHD", "Webcam Privacy Shutter": "Privacy Shutter",
		"Microphone Type": "Dual Microphone", "Speaker Brand": "Acer TrueHarmony",
		"Audio Features": "Stereo Speakers", "WiFi Standard": "Wi-Fi 6 (802.11ax)",
		"Bluetooth Version": "Bluetooth 5.2", "Ethernet": "No",
		"USB A Ports": "1x USB 3.2 Gen 1", "USB C Ports": "1x USB-C (USB 3.2 Gen 2)",
		"Thunderbolt Version": "No", "HDMI Version": "HDMI 2.0",
		"DisplayPort Support": "No", "SD Card Reader": "No",
		"Audio Jack": "3.5mm Combo Jack", "Battery Capacity": "50 Wh",
		"Battery Life": "8-10 hours", "Fast Charging": "No",
		"Power Adapter Wattage": "65W", "USB-C Charging": "Yes",
		"Body Material": "Aluminum", "Color": "Silver",
		"Product Dimensions": "312.2 x 217.9 x 15.9 mm",
		"Product Weight":     "1.2 kg", "Build Standard": "Premium",
		"Cooling System": "Single Fan", "Security Chip": "TPM 2.0",
		"Release Year": "2024", "Warranty": "1 Year International Warranty",
	}

	for key, value := range specs {
		keyID, exists := existingKeyMapping[key]
		if !exists {
			log.Printf("⚠️  Specification key not found: %s", key)
			continue
		}

		spec := models.SpecificationModel{
			ProductID: prod.ID, SpecificationKeyID: keyID, Value: value,
		}

		var existingSpec models.SpecificationModel
		if err := db.Where("product_id = ? AND specification_key_id = ?", prod.ID, keyID).First(&existingSpec).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
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

		if banglaValue, exists := ass.getBanglaTranslations()[value]; exists {
			translation := models.SpecificationTranslationModel{
				SpecificationID: spec.ID, Locale: "bn", Value: banglaValue,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("❌ Failed to create translation for key %s: %v", key, err)
			}
		}
	}

	log.Printf("✅ Acer Swift 3 specifications seeded successfully")
	return nil
}
