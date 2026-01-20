package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

type AcerAspireVeroSeeder struct {
	BaseSeeder
}

func NewAcerAspireVeroSeeder() *AcerAspireVeroSeeder {
	return &AcerAspireVeroSeeder{BaseSeeder: BaseSeeder{name: "Acer Aspire Vero Specifications"}}
}

func (aas *AcerAspireVeroSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"Acer": "এসার", "Aspire Vero": "অ্যাস্পায়ার ভেরো", "Eco-Friendly": "ইকো-ফ্রেন্ডলি",
		"Windows 11 Home": "উইন্ডোজ ১১ হোম", "Intel": "ইন্টেল", "Core i5-1235U": "কোর আই৫-১২৩৫ইউ",
		"12th Gen": "১২তম প্রজন্ম", "10": "১০", "12": "১২", "1.3 GHz": "১.৩ গিগাহার্টজ",
		"4.4 GHz": "৪.৪ গিগাহার্টজ", "12 MB": "১২ এমবি", "Integrated": "ইন্টিগ্রেটেড",
		"Intel Iris Xe": "ইন্টেল আইরিস এক্সই", "Shared": "শেয়ারড", "15.6 inches": "১৫.৬ ইঞ্চি",
		"1920 x 1080 pixels": "১৯২০ x ১০৮০ পিক্সেল", "IPS": "আইপিএস", "60 Hz": "৬০ হার্জ",
		"16:9": "১৬:৯", "250 nits": "২৫০ নিটস", "45% NTSC": "৪৫% এনটিএসসি", "No": "না",
		"Matte": "ম্যাট", "DDR4": "ডিডিআর৪", "16 GB": "১৬ জিবি", "3200 MHz": "৩২০০ মেগাহার্টজ",
		"2": "২", "Yes (up to 32GB)": "হ্যাঁ (৩২জিবি পর্যন্ত)", "SSD": "এসএসডি",
		"512 GB": "৫১২ জিবি", "NVMe PCIe Gen3": "এনভিএমই পিসিআই জেন৩",
		"Yes (M.2 slot)": "হ্যাঁ (এম.২ স্লট)", "None": "কোনটি নয়", "Standard": "স্ট্যান্ডার্ড",
		"Yes": "হ্যাঁ", "English/Bangla": "ইংরেজি/বাংলা", "Precision Touchpad": "প্রিসিশন টাচপ্যাড",
		"720p HD": "৭২০পি এইচডি", "Dual Microphone": "ডুয়াল মাইক্রোফোন",
		"Acer TrueHarmony": "এসার ট্রু হারমনি", "Stereo Speakers": "স্টেরিও স্পিকার",
		"Wi-Fi 6 (802.11ax)": "ওয়াই-ফাই ৬ (৮০২.১১এক্স)", "Bluetooth 5.2": "ব্লুটুথ ৫.২",
		"1x USB 2.0": "১x ইউএসবি ২.০", "2x USB 3.2 Gen 1": "২x ইউএসবি ৩.২ জেন ১",
		"1x USB-C (USB 3.2 Gen 1)": "১x ইউএসবি-সি (ইউএসবি ৩.২ জেন ১)",
		"HDMI 2.0":                 "এইচডিএমআই ২.০", "3.5mm Combo Jack": "৩.৫মিমি কম্বো জ্যাক",
		"56 Wh": "৫৬ ডব্লিউএইচ", "8-9 hours": "৮-৯ ঘন্টা", "65W": "৬৫ওয়াট",
		"Recycled Plastic": "রিসাইকেল প্লাস্টিক", "Volcanic Gray": "ভলক্যানিক গ্রে",
		"363.4 x 238.5 x 17.9 mm": "৩৬৩.৪ x ২৩৮.৫ x ১৭.৯ মিমি", "1.75 kg": "১.৭৫ কেজি",
		"Single Fan": "সিঙ্গেল ফ্যান", "TPM 2.0": "টিপিএম ২.০", "2024": "২০২৪",
		"1 Year International Warranty": "১ বছর আন্তর্জাতিক ওয়ারেন্টি",
	}
}

func (aas *AcerAspireVeroSeeder) Seed(db *gorm.DB) error {
	productSlug := "acer-aspire-vero"
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
		"Brand": "Acer", "Model Name": "Aspire Vero", "Laptop Type": "Eco-Friendly",
		"Operating System": "Windows 11 Home", "Processor Brand": "Intel", "Processor Model": "Core i5-1235U",
		"Processor Generation": "12th Gen", "Processor Cores": "10", "Processor Threads": "12",
		"Processor Base Clock": "1.3 GHz", "Processor Boost Clock": "4.4 GHz", "Processor Cache": "12 MB",
		"Chipset": "Intel Platform Controller Hub", "Graphics Type": "Integrated", "Graphics Brand": "Intel",
		"Graphics Card": "Intel Iris Xe", "Graphics VRAM": "Shared", "Display Size": "15.6 inches",
		"Screen Resolution": "1920 x 1080 pixels", "Display Panel Type": "IPS", "Display Refresh Rate": "60 Hz",
		"Display Aspect Ratio": "16:9", "Display Brightness": "250 nits", "Display Color Gamut": "45% NTSC",
		"Display Touch Support": "No", "Display Surface": "Matte", "RAM Type": "DDR4", "RAM Capacity": "16 GB",
		"RAM Speed": "3200 MHz", "RAM Slots": "2", "RAM Expandable": "Yes (up to 32GB)", "Storage Type": "SSD",
		"Storage Capacity": "512 GB", "Storage Interface": "NVMe PCIe Gen3", "Storage Expandable": "Yes (M.2 slot)",
		"Secondary Storage Type": "None", "Secondary Storage Capacity": "None", "Keyboard Type": "Standard",
		"Backlit Keyboard": "Yes", "Keyboard Language": "English/Bangla", "Touchpad Type": "Precision Touchpad",
		"Fingerprint Reader": "No", "Webcam Resolution": "720p HD", "Webcam Privacy Shutter": "No",
		"Microphone Type": "Dual Microphone", "Speaker Brand": "Acer TrueHarmony", "Audio Features": "Stereo Speakers",
		"WiFi Standard": "Wi-Fi 6 (802.11ax)", "Bluetooth Version": "Bluetooth 5.2", "Ethernet": "No",
		"USB A Ports": "1x USB 2.0", "USB C Ports": "1x USB-C (USB 3.2 Gen 1)", "Thunderbolt Version": "No",
		"HDMI Version": "HDMI 2.0", "DisplayPort Support": "No", "SD Card Reader": "No",
		"Audio Jack": "3.5mm Combo Jack", "Battery Capacity": "56 Wh", "Battery Life": "8-9 hours",
		"Fast Charging": "No", "Power Adapter Wattage": "65W", "USB-C Charging": "No",
		"Body Material": "Recycled Plastic", "Color": "Volcanic Gray", "Product Dimensions": "363.4 x 238.5 x 17.9 mm",
		"Product Weight": "1.75 kg", "Build Standard": "Standard", "Cooling System": "Single Fan",
		"Security Chip": "TPM 2.0", "Release Year": "2024", "Warranty": "1 Year International Warranty",
	}

	for key, value := range specs {
		keyID, exists := existingKeyMapping[key]
		if !exists {
			log.Printf("⚠️  Specification key not found: %s", key)
			continue
		}

		spec := models.SpecificationModel{ProductID: prod.ID, SpecificationKeyID: keyID, Value: value}
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
			translation := models.SpecificationTranslationModel{SpecificationID: spec.ID, Locale: "bn", Value: banglaValue}
			if err := db.Create(&translation).Error; err != nil {
				log.Printf("❌ Failed to create translation for key %s: %v", key, err)
			}
		}
	}

	log.Printf("✅ Acer Aspire Vero specifications seeded successfully")
	return nil
}
