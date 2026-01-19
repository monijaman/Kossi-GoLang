package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

// CameraSeeder seeds specifications for a camera
type CameraSeeder struct {
	BaseSeeder
}

// NewCameraSeeder creates a new camera seeder
func NewCameraSeeder() *CameraSeeder {
	return &CameraSeeder{
		BaseSeeder: BaseSeeder{name: "Camera Specifications"},
	}
}

func (cs *CameraSeeder) getBanglaTranslations() map[string]string {
	return map[string]string{
		// Camera Type & Basic Info
		"Mirrorless": "মিররলেস",
		"DSLR":       "ডিএসএলআর",
		"Full Frame": "ফুল ফ্রেম",
		"APS-C":      "এপিএস-সি",
		"CMOS":       "সিএমওএস",
		"BSI-CMOS":   "বিএসআই-সিএমওএস",

		// Megapixels
		"24.2 MP": "২৪.২ মেগাপিক্সেল",
		"45.7 MP": "৪৫.৭ মেগাপিক্সেল",

		// ISO Range
		"100-51200":            "১০০-৫১২০০",
		"50-204800 (Expanded)": "৫০-২০৪৮০০ (সম্প্রসারিত)",

		// Lens Mount
		"Sony E-Mount": "সনি ই-মাউন্ট",
		"Canon EF":     "ক্যানন ইএফ",
		"Nikon Z":      "নিকন জেড",

		// Aperture
		"f/1.4": "এফ/১.৪",
		"f/2.8": "এফ/২.৮",
		"f/4":   "এফ/৪",

		// Optical Zoom
		"10x":  "১০x",
		"5x":   "৫x",
		"None": "নেই",

		// Digital Zoom
		"2x": "২x",
		"4x": "৪x",

		// Image Stabilization
		"5-Axis In-Body": "৫-অক্ষ ইন-বডি",
		"Optical":        "অপটিক্যাল",
		"Yes":            "হ্যাঁ",
		"No":             "না",

		// Autofocus
		"693-point":       "৬৯৩-পয়েন্ট",
		"425-point":       "৪২৫-পয়েন্ট",
		"Hybrid AF":       "হাইব্রিড এএফ",
		"Phase Detection": "ফেজ ডিটেকশন",

		// Continuous Shooting
		"10 fps": "১০ এফপিএস",
		"20 fps": "২০ এফপিএস",

		// Shutter Speed
		"1/8000 - 30s": "১/৮০০০ - ৩০সে",
		"1/4000 - 30s": "১/৪০০০ - ৩০সে",

		// Video Resolution
		"4K @ 60fps":     "৪কে @ ৬০এফপিএস",
		"4K @ 30fps":     "৪কে @ ৩০এফপিএস",
		"1080p @ 120fps": "১০৮০পি @ ১২০এফপিএস",

		// Viewfinder
		"Electronic":     "ইলেকট্রনিক",
		"2.36M-dot OLED": "২.৩৬এম-ডট ওএলইডি",

		// LCD Screen
		"3.0 inch":     "৩.০ ইঞ্চি",
		"3.2 inch":     "৩.২ ইঞ্চি",
		"921k-dot":     "৯২১কে-ডট",
		"1.04M-dot":    "১.০৪এম-ডট",
		"Touchscreen":  "টাচস্ক্রিন",
		"Tilting":      "টিল্টিং",
		"Articulating": "আর্টিকুলেটিং",

		// Memory Cards
		"SD/SDHC/SDXC":     "এসডি/এসডিএইচসি/এসডিএক্সসি",
		"CFexpress Type A": "সিএফএক্সপ্রেস টাইপ এ",
		"Dual Slot":        "ডুয়াল স্লট",
		"Single Slot":      "সিঙ্গেল স্লট",

		// Connectivity
		"Wi-Fi":         "ওয়াই-ফাই",
		"Bluetooth 4.2": "ব্লুটুথ ৪.২",
		"Bluetooth 5.0": "ব্লুটুথ ৫.০",
		"NFC":           "এনএফসি",
		"USB Type-C":    "ইউএসবি টাইপ-সি",
		"Micro USB":     "মাইক্রো ইউএসবি",
		"HDMI Type D":   "এইচডিএমআই টাইপ ডি",
		"HDMI Type A":   "এইচডিএমআই টাইপ এ",

		// Battery
		"NP-FZ100":     "এনপি-এফজেড১০০",
		"LP-E6N":       "এলপি-ই৬এন",
		"2280 shots":   "২২৮০ শট",
		"1200 shots":   "১২০০ শট",
		"USB Charging": "ইউএসবি চার্জিং",

		// Build & Weather
		"Magnesium Alloy": "ম্যাগনেসিয়াম অ্যালয়",
		"Weather Sealed":  "ওয়েদার সিল্ড",

		// Weight & Dimensions
		"650g":                   "৬৫০গ্রাম",
		"599g":                   "৫৯৯গ্রাম",
		"126.9 x 95.6 x 73.7 mm": "১২৬.৯ x ৯৫.৬ x ৭৩.৭ মিমি",

		// Additional Features
		"Built-in Flash":     "বিল্ট-ইন ফ্ল্যাশ",
		"External Mic Input": "এক্সটার্নাল মাইক ইনপুট",
		"Headphone Jack":     "হেডফোন জ্যাক",
		"GPS":                "জিপিএস",
		"Intervalometer":     "ইন্টারভালোমিটার",
		"Focus Stacking":     "ফোকাস স্ট্যাকিং",
		"HDR":                "এইচডিআর",
		"RAW":                "র",
		"JPEG":               "জেপিইজি",
	}
}

// Seed implements the Seeder interface for Camera
func (cs *CameraSeeder) Seed(db *gorm.DB) error {
	// Replace with actual product slug - you'll need to create the product first
	productSlug := "sony-a7-iii" // Change this to your actual camera product slug
	var prod models.ProductModel
	if err := db.Where("slug = ?", productSlug).First(&prod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("⚠️  Product not found: %s", productSlug)
			return nil
		}
		return err
	}
	productID := prod.ID

	// Mapping of specification keys to their IDs from specs.sql
	existingKeyMapping := map[string]uint{
		// Camera specific specs from specs.sql
		"Aperture":                184,
		"Auto Focus":              185,
		"Digital Zoom":            186,
		"Flash Type":              187,
		"Image Stabilization":     188,
		"ISO Range":               189,
		"Lens Mount":              190,
		"Maximum Aperture":        191,
		"Megapixels":              192,
		"Optical Zoom":            193,
		"Sensor Size":             194,
		"Sensor Type":             195,
		"Shutter Speed":           196,
		"Video Format":            197,
		"Viewfinder Type":         198,
		"Camera Features":         16,
		"Video Recording":         77,
		"Camera Video Resolution": 692,

		// Display & Screen
		"Screen Size":         66,
		"Touchscreen":         129,
		"Display Type":        27,
		"Screen Resolution":   207,
		"Touchscreen Display": 209,

		// Connectivity
		"Bluetooth Version": 13,
		"USB Type":          76,
		"NFC Support":       46,
		"Connectivity":      357,
		"HDMI Ports":        167,
		"Wi-Fi Support":     387,

		// Power & Battery
		"Battery Type":      12,
		"Battery Capacity":  11,
		"Wireless Charging": 82,
		"Fast Charging":     553,
		"Charging Speed":    18,

		// Physical Specs
		"Weight":             80,
		"Dimensions":         25,
		"Product Dimensions": 319,
		"Product Weight":     697,
		"Build Material":     14,
		"Water Resistance":   78,

		// Common specs
		"Brand":            310,
		"Model Name":       316,
		"Model Number":     317,
		"Color":            311,
		"Available Colors": 10,
		"Condition":        312,
		"Manufacturer":     315,
		"Warranty":         323,

		// Additional useful specs
		"Special Features":    69,
		"Included Components": 314,
		"Product Description": 318,
		"Device Type":         24,
	}

	// Example specifications for Sony A7 III (you can modify based on your camera)
	specs := map[string]string{
		// Basic Info
		"Brand":        "Sony",
		"Model Name":   "Alpha 7 III",
		"Model Number": "ILCE-7M3",
		"Device Type":  "Mirrorless Camera",
		"Color":        "Black",
		"Manufacturer": "Sony Corporation",
		"Warranty":     "2 Years International Warranty",
		"Condition":    "Brand New",

		// Camera Sensor & Image
		"Sensor Type": "BSI-CMOS (Exmor R)",
		"Sensor Size": "Full Frame (35.6 x 23.8 mm)",
		"Megapixels":  "24.2 MP",
		"ISO Range":   "100-51200 (Expandable to 50-204800)",

		// Lens & Optics
		"Lens Mount":       "Sony E-Mount",
		"Maximum Aperture": "Lens Dependent",
		"Aperture":         "Lens Dependent",
		"Optical Zoom":     "Lens Dependent",
		"Digital Zoom":     "2x (Clear Image Zoom)",

		// Focus & Stabilization
		"Image Stabilization": "5-Axis In-Body (5.5 stops)",
		"Auto Focus":          "693-point Hybrid AF (Phase Detection + Contrast)",

		// Shooting
		"Shutter Speed": "1/8000 - 30s (Bulb mode available)",

		// Video Capabilities
		"Video Format":            "XAVC S, AVCHD, MP4",
		"Video Recording":         "4K (3840 x 2160) @ 30p/24p, Full HD @ 120p",
		"Camera Video Resolution": "4K UHD (3840 x 2160)",
		"Camera Features":         "Eye AF, Animal Eye AF, Real-time Tracking, Silent Shooting, Dual Card Slots",

		// Viewfinder & Display
		"Viewfinder Type":     "Electronic (OLED, 2.36M-dot)",
		"Screen Size":         "3.0 inch",
		"Screen Resolution":   "921,600 dots",
		"Touchscreen":         "Yes",
		"Touchscreen Display": "Tilting Touchscreen",
		"Display Type":        "TFT LCD",

		// Flash
		"Flash Type": "External (Hot Shoe for Flash)",

		// Connectivity
		"USB Type":          "USB Type-C (USB 3.2 Gen 1)",
		"Bluetooth Version": "Bluetooth 4.2",
		"NFC Support":       "Yes",
		"Wi-Fi Support":     "Yes (802.11ac 2.4GHz/5GHz)",
		"HDMI Ports":        "1x HDMI Type D (Micro)",
		"Connectivity":      "USB-C, Wi-Fi, Bluetooth, NFC, HDMI, 3.5mm Mic, 3.5mm Headphone",

		// Battery & Power
		"Battery Type":     "NP-FZ100 (Rechargeable Lithium-Ion)",
		"Battery Capacity": "2280 mAh (7.2V, 16.4Wh)",
		"Charging Speed":   "USB Charging Supported",

		// Build & Physical
		"Weight":             "650g (body only)",
		"Product Weight":     "650g",
		"Dimensions":         "126.9 x 95.6 x 73.7 mm",
		"Product Dimensions": "126.9 x 95.6 x 73.7 mm (W x H x D)",
		"Build Material":     "Magnesium Alloy",
		"Water Resistance":   "Weather Sealed (Dust & Moisture Resistant)",

		// Additional
		"Special Features":    "Dual SD Card Slots (UHS-II), USB Charging, Pixel Shift Multi Shooting, 5-axis IBIS",
		"Included Components": "Camera Body, NP-FZ100 Battery, BC-QZ1 Charger, USB Cable, Shoulder Strap, Body Cap, Shoe Cap, Eyepiece Cup",
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
		if banglaValue, exists := cs.getBanglaTranslations()[value]; exists {
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

	log.Printf("✅ Camera specifications seeded successfully")
	return nil
}
