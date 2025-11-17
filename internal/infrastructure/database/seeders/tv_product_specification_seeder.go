package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVProductSpecificationSeeder handles seeding TV products with specifications
type TVProductSpecificationSeeder struct {
	BaseSeeder
}

// NewTVProductSpecificationSeeder creates a new TV product specification seeder
func NewTVProductSpecificationSeeder() *TVProductSpecificationSeeder {
	return &TVProductSpecificationSeeder{
		BaseSeeder: BaseSeeder{name: "TV Product Specifications"},
	}
}

// Seed implements the Seeder interface
func (tvpss *TVProductSpecificationSeeder) Seed(db *gorm.DB) error {
	// TV brand data with sample models and specifications
	tvBrands := []struct {
		slug   string
		name   string
		models []tvProductModel
	}{
		{
			slug: "samsung",
			name: "Samsung",
			models: []tvProductModel{
				{
					name: "Samsung 55\" 4K Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "55\"",
						"Resolution":             "4K (3840x2160)",
						"Panel Type":             "QLED",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "800",
						"Smart TV OS":            "Tizen",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "20W",
						"Dolby Atmos Support":    "Yes",
						"Price (BDT)":            "55000",
					},
				},
				{
					name: "Samsung 65\" 8K Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "65\"",
						"Resolution":             "8K (7680x4320)",
						"Panel Type":             "QLED",
						"Refresh Rate":           "120Hz",
						"Brightness (Nits)":      "1000",
						"Smart TV OS":            "Tizen",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "4",
						"USB Ports":              "3",
						"Audio Output (Watts)":   "40W",
						"Dolby Atmos Support":    "Yes",
						"Price (BDT)":            "150000",
					},
				},
			},
		},
		{
			slug: "lg",
			name: "LG",
			models: []tvProductModel{
				{
					name: "LG 55\" OLED TV",
					specs: map[string]interface{}{
						"Screen Size":            "55\"",
						"Resolution":             "4K (3840x2160)",
						"Panel Type":             "OLED",
						"Refresh Rate":           "120Hz",
						"Brightness (Nits)":      "800",
						"Smart TV OS":            "webOS",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "4",
						"USB Ports":              "3",
						"Audio Output (Watts)":   "40W",
						"Dolby Atmos Support":    "Yes",
						"Price (BDT)":            "80000",
					},
				},
				{
					name: "LG 43\" HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "300",
						"Smart TV OS":            "webOS",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "10W",
						"Price (BDT)":            "25000",
					},
				},
			},
		},
		{
			slug: "sony",
			name: "Sony",
			models: []tvProductModel{
				{
					name: "Sony Bravia 55\" 4K TV",
					specs: map[string]interface{}{
						"Screen Size":            "55\"",
						"Resolution":             "4K (3840x2160)",
						"Panel Type":             "LCD",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "700",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "20W",
						"Price (BDT)":            "60000",
					},
				},
			},
		},
		{
			slug: "xiaomi",
			name: "Xiaomi",
			models: []tvProductModel{
				{
					name: "Xiaomi 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "400",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "18000",
					},
				},
				{
					name: "Xiaomi 55\" 4K Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "55\"",
						"Resolution":             "4K (3840x2160)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "500",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "16W",
						"Price (BDT)":            "35000",
					},
				},
			},
		},
		{
			slug: "walton",
			name: "Walton",
			models: []tvProductModel{
				{
					name: "Walton 32\" HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "HD (1366x768)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "300",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "12000",
					},
				},
				{
					name: "Walton 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "350",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "10W",
						"Price (BDT)":            "16000",
					},
				},
			},
		},
		{
			slug: "hisense",
			name: "Hisense",
			models: []tvProductModel{
				{
					name: "Hisense 50\" 4K Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "50\"",
						"Resolution":             "4K (3840x2160)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "450",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "20W",
						"Price (BDT)":            "40000",
					},
				},
			},
		},
		{
			slug: "tcl",
			name: "TCL",
			models: []tvProductModel{
				{
					name: "TCL 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "350",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "12W",
						"Price (BDT)":            "15000",
					},
				},
			},
		},
		{
			slug: "haier",
			name: "Haier",
			models: []tvProductModel{
				{
					name: "Haier 32\" HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "HD (1366x768)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "280",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "11000",
					},
				},
			},
		},
		{
			slug: "panasonic",
			name: "Panasonic",
			models: []tvProductModel{
				{
					name: "Panasonic 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "400",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "20W",
						"Price (BDT)":            "22000",
					},
				},
			},
		},
		{
			slug: "singer",
			name: "Singer",
			models: []tvProductModel{
				{
					name: "Singer 32\" HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "HD (1366x768)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Brightness (Nits)":      "300",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "10000",
					},
				},
			},
		},
		{
			slug: "vision",
			name: "Vision",
			models: []tvProductModel{
				{
					name: "Vision 32\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "10W",
						"Price (BDT)":            "13000",
					},
				},
			},
		},
		{
			slug: "minister",
			name: "Minister",
			models: []tvProductModel{
				{
					name: "Minister 24\" HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "24\"",
						"Resolution":             "HD (1366x768)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "6W",
						"Price (BDT)":            "8000",
					},
				},
			},
		},
		{
			slug: "rangs",
			name: "Rangs",
			models: []tvProductModel{
				{
					name: "Rangs 32\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "12500",
					},
				},
				{
					name: "Rangs 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "16W",
						"Price (BDT)":            "20000",
					},
				},
			},
		},
		{
			slug: "myOne",
			name: "MyOne",
			models: []tvProductModel{
				{
					name: "MyOne 32\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "13500",
					},
				},
			},
		},
		{
			slug: "toshiba",
			name: "Toshiba",
			models: []tvProductModel{
				{
					name: "Toshiba 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "16W",
						"Price (BDT)":            "19000",
					},
				},
				{
					name: "Toshiba 50\" 4K Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "50\"",
						"Resolution":             "4K (3840x2160)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "3",
						"USB Ports":              "2",
						"Audio Output (Watts)":   "20W",
						"Price (BDT)":            "42000",
					},
				},
			},
		},
		{
			slug: "sharp",
			name: "Sharp",
			models: []tvProductModel{
				{
					name: "Sharp 43\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "43\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "12W",
						"Price (BDT)":            "17000",
					},
				},
			},
		},
		{
			slug: "konka",
			name: "Konka",
			models: []tvProductModel{
				{
					name: "Konka 32\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "32\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "10W",
						"Price (BDT)":            "14000",
					},
				},
			},
		},
		{
			slug: "eco+",
			name: "ECO+",
			models: []tvProductModel{
				{
					name: "ECO+ 24\" HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "24\"",
						"Resolution":             "HD (1366x768)",
						"Panel Type":             "VA",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "6W",
						"Price (BDT)":            "7500",
					},
				},
			},
		},
		{
			slug: "haiko",
			name: "Haiko",
			models: []tvProductModel{
				{
					name: "Haiko 28\" Full HD Smart TV",
					specs: map[string]interface{}{
						"Screen Size":            "28\"",
						"Resolution":             "Full HD (1920x1080)",
						"Panel Type":             "IPS",
						"Refresh Rate":           "60Hz",
						"Smart TV OS":            "Android TV",
						"WiFi Connectivity":      "Yes",
						"Bluetooth Connectivity": "Yes",
						"HDMI Ports":             "2",
						"USB Ports":              "1",
						"Audio Output (Watts)":   "8W",
						"Price (BDT)":            "11500",
					},
				},
			},
		},
	}

	// Get TV category
	var tvCategory models.CategoryModel
	if err := db.Where("id = ?", 124).First(&tvCategory).Error; err != nil {
		return nil // Skip if category not found
	}

	// Seed products for each brand
	for _, brandData := range tvBrands {
		// Find brand by slug
		var brand models.BrandModel
		if err := db.Where("slug = ?", brandData.slug).First(&brand).Error; err != nil {
			continue // Skip if brand not found
		}

		// Create products for this brand
		for _, prodModel := range brandData.models {
			// Check if product already exists
			var existingProduct models.ProductModel
			result := db.Where("name = ? AND brand_id = ?", prodModel.name, brand.ID).First(&existingProduct)

			if result.Error == gorm.ErrRecordNotFound {
				// Create new product
				categoryID := tvCategory.ID
				brandID := brand.ID
				product := models.ProductModel{
					CategoryID: &categoryID,
					BrandID:    &brandID,
					Name:       prodModel.name,
					Slug:       generateSlug(prodModel.name),
					Status:     1,
				}

				if err := db.Create(&product).Error; err != nil {
					continue
				}

				// Add specifications to product
				for specKey, specValue := range prodModel.specs {
					// Find specification key by name
					var specKeyModel models.SpecificationKeyModel
					if err := db.Where("specification_key = ?", specKey).First(&specKeyModel).Error; err != nil {
						continue
					}

					// Create product specification
					productSpec := models.SpecificationModel{
						ProductID:          product.ID,
						SpecificationKeyID: specKeyModel.ID,
						Value:              toString(specValue),
						Status:             1,
					}

					db.Create(&productSpec)
				}
			}
		}
	}

	return nil
}

// Helper structs
type tvProductModel struct {
	name  string
	specs map[string]interface{}
}

// Helper function to convert interface to string
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	default:
		// All spec values are passed as strings in the specs map
		return ""
	}
}

// Helper function to generate slug from name
func generateSlug(name string) string {
	// Simple slug generation - can be improved with external library
	slug := ""
	for _, ch := range name {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			if ch >= 'A' && ch <= 'Z' {
				slug += string(ch + 32)
			} else {
				slug += string(ch)
			}
		} else if ch == ' ' {
			slug += "-"
		}
	}
	return slug
}
