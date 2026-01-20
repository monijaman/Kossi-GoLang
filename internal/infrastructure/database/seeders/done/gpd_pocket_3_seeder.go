package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// GpdPocket3Seeder seeds specifications for GPD Pocket 3
type GpdPocket3Seeder struct {
	BaseSeeder
}

// NewGpdPocket3Seeder creates a new GPD Pocket 3 seeder
func NewGpdPocket3Seeder() *GpdPocket3Seeder {
	return &GpdPocket3Seeder{
		BaseSeeder: BaseSeeder{name: "GPD Pocket 3 Specifications"},
	}
}

func (gps3 *GpdPocket3Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"GPD":                          "জিপিডি",
		"Pocket 3":                     "পকেট ৩",
		"Windows 11 Home":              "উইন্ডোজ ১১ হোম",
		"Intel":                        "ইন্টেল",
		"Core i7-1195G7":               "কোর আই৭-১১৯৫জি৭",
		"11th Gen":                     "১১তম প্রজন্ম",
		"Intel Iris Xe":                "ইন্টেল আইরিস এক্সই",
		"1920 x 1200 pixels":           "১৯২০ x ১২০০ পিক্সেল",
		"8 inches":                     "৮ ইঞ্চি",
		"38.5 Wh":                      "৩৮.৫ ডব্লিউএইচ",
		"Aluminum":                     "অ্যালুমিনিয়াম",
		"Black":                        "কালো",
		"0.7 kg":                       "০.৭ কেজি",
		"2023":                         "২০২৩",
		"1 Year Warranty":              "১ বছর ওয়ারেন্টি",
		"1 TB":                         "১ টিবি",
		"Bluetooth 5.2":                "ব্লুটুথ ৫.২",
		"16 GB":                        "১৬ জিবি",
		"IPS":                          "আইপিএস",
		"60 Hz":                        "৬০ হার্জ",
		"Yes":                          "হ্যাঁ",
		"3.5mm Combo Jack":             "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 6E":                     "ওয়াই-ফাই ৬ই",
		"USB 3.2 Gen 2":                "ইউএসবি ৩.২ জেন ২",
		"Integrated":                   "ইন্টিগ্রেটেড",
		"1 TB SSD":                     "১ টিবি এসএসডি",
		"Intel Core i7":                "ইন্টেল কোর আই৭",
		"Single Fan":                   "সিঙ্গেল ফ্যান",
		"60":                           "৬০",
		"No":                           "না",
		"1 Year":                       "১ বছর",
		"Ultraportable":                "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                  "লিথিয়াম-আয়ন",
		"65W USB-C":                    "৬৫ওয়াট ইউএসবি-সি",
		"8-10 hours":                   "৮-১০ ঘন্টা",
		"IPS, 500 nits":                "আইপিএস, ৫০০ নিটস",
		"2.9 GHz base / 5.0 GHz boost": "২.৯ গিগাহার্টজ বেস / ৫.০ গিগাহার্টজ বুস্ট",
		"12 MB cache":                  "১২ এমবি ক্যাশ",
		"2MP IR":                       "২এমপি আইআর",
		"Stereo Speakers":              "স্টেরিও স্পিকার",
		"198 x 137 x 20 mm":            "১৯৮ x ১৩৭ x ২০ মিমি",
		"TPM 2.0":                      "টিপিএম ২.০",
		"4":                            "৪",
		"8":                            "৮",
		"3733 MHz":                     "৩৭৩৩ মেগাহার্টজ",
		"PCIe Gen4":                    "পিসিআই জেন৪",
		"Yes (M.2)":                    "হ্যাঁ (এম.২)",
		"Shared":                       "শেয়ার্ড",
		"English/Bangla":               "ইংরেজি/বাংলা",
		"Standard":                     "স্ট্যান্ডার্ড",
		"HDMI 2.0b":                    "এইচডিএমআই ২.০বি",
		"Thunderbolt 4, 3x USB 3.2":    "থান্ডারবোল্ট ৪, ৩x ইউএসবি ৩.২",
	}
}

// Seed implements the Seeder interface for GPD Pocket 3
func (gps3 *GpdPocket3Seeder) Seed(db *gorm.DB) error {
	productSlug := "gpd-pocket-3"
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
		"Model Name":              "Pocket 3",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Core i7-1195G7",
		"Processor Generation":    "11th Gen",
		"Graphics Card":           "Intel Iris Xe",
		"Screen Resolution":       "1920 x 1200 pixels",
		"Display Size":            "8 inches",
		"Battery Capacity":        "38.5 Wh",
		"Build Material":          "Aluminum",
		"Color":                   "Black",
		"Product Weight":          "0.7 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "1 TB",
		"Bluetooth Version":       "Bluetooth 5.2",
		"RAM":                     "16 GB",
		"Weight":                  "0.7 kg",
		"Display Type":            "IPS",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "8 inches",
		"Backlit Keyboard":        "Yes",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "16 GB",
		"Wifi Support":            "Wi-Fi 6E",
		"Usb Type":                "USB 3.2 Gen 2",
		"Battery":                 "38.5 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "1 TB SSD",
		"Cpu Type":                "Intel Core i7",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "65W USB-C",
		"Standby Time":            "8-10 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1920 x 1200 pixels",
		"Display Characteristics": "IPS, 500 nits",
		"Processor Speed":         "2.9 GHz base / 5.0 GHz boost",
		"Clock Feature":           "12 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "2MP IR",
		"Audio Quality":           "Stereo Speakers",
		"Sensors":                 "No",
		"Dimensions":              "198 x 137 x 20 mm",
		"Body Type":               "Aluminum",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Black",
		"Special Features":        "TPM 2.0",
		"Processor Cores":         "4",
		"Processor Threads":       "8",
		"RAM Speed":               "3733 MHz",
		"RAM Slots":               "Soldered",
		"RAM Expandable":          "No",
		"Storage Interface":       "PCIe Gen4",
		"Storage Expandable":      "Yes (M.2)",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "Yes",
		"Ethernet":                "Yes",
		"Thunderbolt Version":     "4",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English/Bangla",
		"Build Standard":          "Standard",
		"Touchscreen":             "Yes",
		"HDMI Ports":              "HDMI 2.0b",
		"USB Ports":               "Thunderbolt 4, 3x USB 3.2",
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
		if banglaValue, exists := gps3.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ GPD Pocket 3 specifications seeded successfully")
	return nil
}
