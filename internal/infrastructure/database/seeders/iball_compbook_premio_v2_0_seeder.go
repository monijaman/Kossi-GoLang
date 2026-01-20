package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// IballCompbookPremioV20Seeder seeds specifications for iBall CompBook Premio v2.0
type IballCompbookPremioV20Seeder struct {
	BaseSeeder
}

// NewIballCompbookPremioV20Seeder creates a new iBall CompBook Premio v2.0 seeder
func NewIballCompbookPremioV20Seeder() *IballCompbookPremioV20Seeder {
	return &IballCompbookPremioV20Seeder{
		BaseSeeder: BaseSeeder{name: "iBall CompBook Premio v2.0 Specifications"},
	}
}

func (icp *IballCompbookPremioV20Seeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		"iBall":                         "ইবল",
		"CompBook Premio v2.0":          "কম্পবুক প্রেমিও ভি২.০",
		"Windows 11 Home":               "উইন্ডোজ ১১ হোম",
		"Intel":                         "ইন্টেল",
		"Celeron N4020":                 "সেলেরন এন৪০২০",
		"4th Gen":                       "৪র্থ প্রজন্ম",
		"Intel Platform Controller Hub": "ইন্টেল প্ল্যাটফর্ম কন্ট্রোলার হাব",
		"Intel UHD Graphics 600":        "ইন্টেল ইউএইচডি গ্রাফিক্স ৬০০",
		"1366 x 768 pixels":             "১৩৬৬ x ৭৬৮ পিক্সেল",
		"14 inches":                     "১৪ ইঞ্চি",
		"37 Wh":                         "৩৭ ডব্লিউএইচ",
		"Plastic":                       "প্লাস্টিক",
		"Black":                         "কালো",
		"1.3 kg":                        "১.৩ কেজি",
		"2023":                          "২০২৩",
		"1 Year Warranty":               "১ বছর ওয়ারেন্টি",
		"128 GB":                        "১২৮ জিবি",
		"Bluetooth 5.0":                 "ব্লুটুথ ৫.০",
		"4 GB":                          "৪ জিবি",
		"TN":                            "টিএন",
		"60 Hz":                         "৬০ হার্জ",
		"No":                            "না",
		"3.5mm Combo Jack":              "৩.৫মিমি কম্বো জ্যাক",
		"Wi-Fi 5 (802.11ac)":            "ওয়াই-ফাই ৫ (৮০২.১১এসি)",
		"USB 3.0":                       "ইউএসবি ৩.০",
		"Integrated":                    "ইন্টিগ্রেটেড",
		"128 GB SSD":                    "১২৮ জিবি এসএসডি",
		"Intel Celeron":                 "ইন্টেল সেলেরন",
		"Single Fan":                    "সিঙ্গেল ফ্যান",
		"60":                            "৬০",
		"1 Year":                        "১ বছর",
		"Ultraportable":                 "আল্ট্রাপোর্টেবল",
		"Lithium-ion":                   "লিথিয়াম-আয়ন",
		"45W Adapter":                   "৪৫ওয়াট অ্যাডাপ্টার",
		"6-8 hours":                     "৬-৮ ঘন্টা",
		"TN, 220 nits":                  "টিএন, ২২০ নিটস",
		"1.1 GHz base / 2.8 GHz boost":  "১.১ গিগাহার্টজ বেস / ২.৮ গিগাহার্টজ বুস্ট",
		"4 MB cache":                    "৪ এমবি ক্যাশ",
		"720p HD":                       "৭২০পি এইচডি",
		"Mono Speaker":                  "মোনো স্পিকার",
		"327 x 228 x 19.5 mm":           "৩২৭ x ২২৮ x ১৯.৫ মিমি",
		"2":                             "২",
		"2400 MHz":                      "২৪০০ মেগাহার্টজ",
		"SATA":                          "সাটা",
		"Shared":                        "শেয়ার্ড",
		"English":                       "ইংরেজি",
		"Standard":                      "স্ট্যান্ডার্ড",
		"1x USB 3.0, 1x USB-C":          "১x ইউএসবি ৩.০, ১x ইউএসবি-সি",
	}
}

// Seed implements the Seeder interface for iBall CompBook Premio v2.0
func (icp *IballCompbookPremioV20Seeder) Seed(db *gorm.DB) error {
	productSlug := "iball-compbook-premio-v2-0"
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
		"Chipset":                 19,
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
		"Brand":                   "iBall",
		"Model Name":              "CompBook Premio v2.0",
		"Operating System":        "Windows 11 Home",
		"Processor Brand":         "Intel",
		"Processor Model":         "Celeron N4020",
		"Processor Generation":    "4th Gen",
		"Chipset":                 "Intel Platform Controller Hub",
		"Graphics Card":           "Intel UHD Graphics 600",
		"Screen Resolution":       "1366 x 768 pixels",
		"Display Size":            "14 inches",
		"Battery Capacity":        "37 Wh",
		"Build Material":          "Plastic",
		"Color":                   "Black",
		"Product Weight":          "1.3 kg",
		"Release Year":            "2023",
		"Warranty":                "1 Year Warranty",
		"Storage Capacity":        "128 GB",
		"Bluetooth Version":       "Bluetooth 5.0",
		"RAM":                     "4 GB",
		"Weight":                  "1.3 kg",
		"Display Type":            "TN",
		"Refresh Rate":            "60 Hz",
		"Screen Size":             "14 inches",
		"Backlit Keyboard":        "No",
		"Audio Jack":              "3.5mm Combo Jack",
		"Ram":                     "4 GB",
		"Wifi Support":            "Wi-Fi 5 (802.11ac)",
		"Usb Type":                "USB 3.0",
		"Battery":                 "37 Wh",
		"Gpu Type":                "Integrated",
		"Storage":                 "128 GB SSD",
		"Cpu Type":                "Intel Celeron",
		"Cooling Technology":      "Single Fan",
		"Frequency (Hz)":          "60",
		"App Control":             "No",
		"Warranty Period":         "1 Year",
		"Laptop Type":             "Ultraportable",
		"Battery Type":            "Lithium-ion",
		"Charging Speed":          "45W Adapter",
		"Standby Time":            "6-8 hours",
		"Wireless Charging":       "No",
		"Resolution":              "1366 x 768 pixels",
		"Display Characteristics": "TN, 220 nits",
		"Processor Speed":         "1.1 GHz base / 2.8 GHz boost",
		"Clock Feature":           "4 MB cache",
		"3.5mm Audio Jack":        "Yes",
		"Camera Features":         "720p HD",
		"Audio Quality":           "Mono Speaker",
		"Sensors":                 "No",
		"Dimensions":              "327 x 228 x 19.5 mm",
		"Body Type":               "Plastic",
		"Cooling System":          "Single Fan",
		"Available Colors":        "Black",
		"Special Features":        "No",
		"Processor Cores":         "2",
		"Processor Threads":       "2",
		"RAM Speed":               "2400 MHz",
		"RAM Slots":               "1",
		"RAM Expandable":          "No",
		"Storage Interface":       "SATA",
		"Storage Expandable":      "No",
		"Graphics VRAM":           "Shared",
		"Display Touch Support":   "No",
		"Ethernet":                "No",
		"Thunderbolt Version":     "No",
		"SD Card Reader":          "No",
		"Keyboard Language":       "English",
		"Build Standard":          "Standard",
		"Touchscreen":             "No",
		"HDMI Ports":              "No",
		"USB Ports":               "1x USB 3.0, 1x USB-C",
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
		if banglaValue, exists := icp.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ iBall CompBook Premio v2.0 specifications seeded successfully")
	return nil
}
