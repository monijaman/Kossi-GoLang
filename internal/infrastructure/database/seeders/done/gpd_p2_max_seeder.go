package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// GpdP2MaxSeeder seeds specifications for GPD P2 Max
type GpdP2MaxSeeder struct {
	BaseSeeder
}

// NewGpdP2MaxSeeder creates a new GPD P2 Max seeder
func NewGpdP2MaxSeeder() *GpdP2MaxSeeder {
	return &GpdP2MaxSeeder{
		BaseSeeder: BaseSeeder{name: "GPD P2 Max Specifications"},
	}
}

func (gp2ms *GpdP2MaxSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"GPD":                          "জিপিডি",
		"P2 Max":                       "পি২ ম্যাক্স",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"AMD":                          "এএমডি",
		"Ryzen 7 6800U":                "রাইজেন ৭ ৬৮০০ইউ",
		"Zen 3+":                       "জেন ৩+",
		"AMD Radeon 680M":              "এএমডি রেডিওন ৬৮০এম",
		"1920 x 1080 pixels":           "১৯২০ x ১০৮০ পিক্সেল",
		"8 inches":                     "৮ ইঞ্চি",
		"46 Wh":                        "৪৬ ডব্লিউএইচ",
		"Plastic":                      "প্লাস্টিক",
		"Black":                        "কালো",
		"0.8 kg":                       "০.৮ কেজি",
		"2022":                         "২০২২",
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
		"6-8 hours gaming":             "৬-৮ ঘন্টা গেমিং",
		"IPS, 400 nits":                "আইপিএস, ৪০০ নিটস",
		"2.8 GHz base / 4.7 GHz boost": "২.৮ গিগাহার্টজ বেস / ৪.৭ গিগাহার্টজ বুস্ট",
		"16 MB cache":                  "১৬ এমবি ক্যাশ",
		"1MP":                          "১এমপি",
		"Dual Speakers":                "ডুয়াল স্পিকার",
		"210 x 120 x 25 mm":            "২১০ x ১২০ x ২৫ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"8":                            "৮",
		"16":                           "১৬",
		"6400 MHz":                     "৬৪০০ মেগাহার্টজ",
		"PCIe Gen4":                    "পিসিআই জেন৪",
		"Yes (M.2)":                    "হ্যাঁ (এম.২)",
		"Dedicated (RDNA2)":            "ডেডিকেটেড (আরডিএনএ২)",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.1":                     "এইচডিএমআই ২.১",
		"USB4, USB 3.2 Gen2, USB-C":    "ইউএসবি৪, ইউএসবি ৩.২ জেন২, ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for GPD P2 Max
func (gp2ms *GpdP2MaxSeeder) Seed(db *gorm.DB) error {
	productSlug := "gpd-p2-max"
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
		"Model Name":              "P2 Max",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "AMD",
		"Processor Model":         "Ryzen 7 6800U",
		"Processor Generation":    "Zen 3+",
		"Graphics Card":           "AMD Radeon 680M",
		"Screen Resolution":       "1920 x 1080 pixels",
		"Display Size":            "8 inches",
		"Battery Capacity":        "46 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Black",
		"Product Weight":          "0.8 kg",
		"Release Year":            "2022",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "32 GB",
		"Weight":                  "0.8 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "8 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "32 GB",
		"Wifi Support":            "Wi-Fi 6",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "46 Wh",
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
		"Standby Time":            "6-8 hours gaming",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1080 pixels",
		"Display Characteristics": "IPS, 400 nits",
		"Processor Speed":         "2.8 GHz base / 4.7 GHz boost",
		"Clock Feature":           "16 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "1MP",
		"Audio Quality":           "Dual Speakers",
		"Sensors":                 "No",
		"Dimensions":              "210 x 120 x 25 mm",
		"Body Type":               "Plastic",
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
		"Graphics VRAM":           "Dedicated (RDNA2)",
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
		if banglaValue, exists := gp2ms.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ GPD P2 Max specifications seeded successfully")
	return nil
}
