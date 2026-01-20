package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

type AcerAspire7Seeder struct {
	BaseSeeder
}

func NewAcerAspire7Seeder() *AcerAspire7Seeder {
	return &AcerAspire7Seeder{
		BaseSeeder: BaseSeeder{name: "Acer Aspire 7 Specifications"},
	}
}

func (aas *AcerAspire7Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Acer": "এসার", "Aspire 7": "অ্যাস্পায়ার ৭", "Performance": "পারফরম্যান্স",
		"Windows 11 Home": "উইন্ডোজ ১১ হোম", "Intel": "ইন্টেল",
		"Core i7-12650H": "কোর আই৭-১২৬৫০এইচ", "12th Gen": "১২তম প্রজন্ম",
		"10": "১০", "16": "১৬", "2.3 GHz": "২.৩ গিগাহার্টজ",
		"4.7 GHz": "৪.৭ গিগাহার্টজ", "24 MB": "২৪ এমবি",
		"Dedicated": "ডেডিকেটেড", "NVIDIA": "এনভিডিয়া",
		"NVIDIA GeForce RTX 3050": "এনভিডিয়া জিফোর্স আরটিএক্স ৩০৫০",
		"4 GB GDDR6":              "৪ জিবি জিডিডিআর৬", "15.6 inches": "১৫.৬ ইঞ্চি",
		"1920 x 1080 pixels": "১৯২০ x ১০৮০ পিক্সেল", "IPS": "আইপিএস",
		"144Hz": "১৪৪ হার্জ", "16:9": "১৬:৯", "300 nits": "৩০০ নিটস",
		"72% NTSC": "৭২% এনটিএসসি", "No": "না", "Matte": "ম্যাট",
		"DDR4": "ডিডিআর৪", "16 GB": "১৬ জিবি", "3200 MHz": "৩২০০ মেগাহার্টজ",
		"2": "২", "Yes (up to 32GB)": "হ্যাঁ (৩২জিবি পর্যন্ত)",
		"SSD": "এসএসডি", "512 GB": "৫১২ জিবি",
		"NVMe PCIe Gen4": "এনভিএমই পিসিআই জেন৪",
		"Yes (M.2 slot)": "হ্যাঁ (এম.২ স্লট)", "None": "কোনটি নয়",
		"Standard": "স্ট্যান্ডার্ড", "Yes": "হ্যাঁ",
		"English/Bangla": "ইংরেজি/বাংলা", "Precision Touchpad": "প্রিসিশন টাচপ্যাড",
		"720p HD": "৭২০পি এইচডি", "Dual Microphone": "ডুয়াল মাইক্রোফোন",
		"DTS Audio": "ডিটিএস অডিও", "Stereo Speakers": "স্টেরিও স্পিকার",
		"Wi-Fi 6 (802.11ax)": "ওয়াই-ফাই ৬ (৮০২.১১এক্স)",
		"Bluetooth 5.2":      "ব্লুটুথ ৫.২", "Yes (RJ-45)": "হ্যাঁ (আরজে-৪৫)",
		"2x USB 3.2 Gen 1":         "২x ইউএসবি ৩.২ জেন ১",
		"1x USB-C (USB 3.2 Gen 2)": "১x ইউএসবি-সি (ইউএসবি ৩.২ জেন ২)",
		"HDMI 2.1":                 "এইচডিএমআই ২.১", "3.5mm Combo Jack": "৩.৫মিমি কম্বো জ্যাক",
		"57 Wh": "৫৭ ডব্লিউএইচ", "6-7 hours": "৬-৭ ঘন্টা",
		"135W": "১৩৫ওয়াট", "Aluminum & Plastic": "অ্যালুমিনিয়াম ও প্লাস্টিক",
		"Charcoal Black": "চারকোল ব্ল্যাক", "363.4 x 254.5 x 22.9 mm": "৩৬৩.৪ x ২৫৪.৫ x ২২.৯ মিমি",
		"2.15 kg": "২.১৫ কেজি", "Dual Fan": "ডুয়াল ফ্যান",
		"TPM 2.0": "টিপিএম ২.০", "2024": "২০২৪",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
	}
}

func (aas *AcerAspire7Seeder) Seed(db *gorm.DB) error {
	productSlug := "acer-aspire-7"
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
		"Brand": "Acer", "Model Name": "Aspire 7", "Laptop Type": "Performance",
		"Operating System": "Windows 11 Home", "Processor Brand": "Intel",
		"Processor Model": "Core i7-12650H", "Processor Generation": "12th Gen",
		"Processor Cores": "10", "Processor Threads": "16",
		"Processor Base Clock": "2.3 GHz", "Processor Boost Clock": "4.7 GHz",
		"Processor Cache": "24 MB", "Chipset": "Intel Platform Controller Hub",
		"Graphics Type": "Dedicated", "Graphics Brand": "NVIDIA",
		"Graphics Card": "NVIDIA GeForce RTX 3050", "Graphics VRAM": "4 GB GDDR6",
		"Display Size": "15.6 inches", "Screen Resolution": "1920 x 1080 pixels",
		"Display Panel Type": "IPS", "Display Refresh Rate": "144Hz",
		"Display Aspect Ratio": "16:9", "Display Brightness": "300 nits",
		"Display Color Gamut": "72% NTSC", "Display Touch Support": "No",
		"Display Surface": "Matte", "RAM Type": "DDR4", "RAM Capacity": "16 GB",
		"RAM Speed": "3200 MHz", "RAM Slots": "2",
		"RAM Expandable": "Yes (up to 32GB)", "Storage Type": "SSD",
		"Storage Capacity": "512 GB", "Storage Interface": "NVMe PCIe Gen4",
		"Storage Expandable": "Yes (M.2 slot)", "Secondary Storage Type": "None",
		"Secondary Storage Capacity": "None", "Keyboard Type": "Standard",
		"Backlit Keyboard": "Yes", "Keyboard Language": "English/Bangla",
		"Touchpad Type": "Precision Touchpad", "Fingerprint Reader": "No",
		"Webcam Resolution": "720p HD", "Webcam Privacy Shutter": "No",
		"Microphone Type": "Dual Microphone", "Speaker Brand": "DTS Audio",
		"Audio Features": "Stereo Speakers", "WiFi Standard": "Wi-Fi 6 (802.11ax)",
		"Bluetooth Version": "Bluetooth 5.2", "Ethernet": "Yes (RJ-45)",
		"USB A Ports": "2x USB 3.2 Gen 1", "USB C Ports": "1x USB-C (USB 3.2 Gen 2)",
		"Thunderbolt Version": "No", "HDMI Version": "HDMI 2.1",
		"DisplayPort Support": "No", "SD Card Reader": "No",
		"Audio Jack": "3.5mm Combo Jack", "Battery Capacity": "57 Wh",
		"Battery Life": "6-7 hours", "Fast Charging": "No",
		"Power Adapter Wattage": "135W", "USB-C Charging": "No",
		"Body Material": "Aluminum & Plastic", "Color": "Charcoal Black",
		"Product Dimensions": "363.4 x 254.5 x 22.9 mm",
		"Product Weight":     "2.15 kg", "Build Standard": "Standard",
		"Cooling System": "Dual Fan", "Security Chip": "TPM 2.0",
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

		if banglaValue, exists := aas.getBanglaTranslations()[value]; exists {
			translation := models.SpecificationTranslationModel{
				SpecificationID: spec.ID, Locale: "bn", Value: banglaValue,
			}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("❌ Failed to create translation for key %s: %v", key, err)
			}
		}
	}

	log.Printf("✅ Acer Aspire 7 specifications seeded successfully")
	return nil
}
