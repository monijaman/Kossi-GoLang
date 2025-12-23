package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MobileOperatorSpecificationSeeder handles seeding mobile operator specifications and specification keys
type MobileOperatorSpecificationSeeder struct {
	BaseSeeder
}

// NewMobileOperatorSpecificationSeeder creates a new mobile operator specification seeder
func NewMobileOperatorSpecificationSeeder() *MobileOperatorSpecificationSeeder {
	return &MobileOperatorSpecificationSeeder{
		BaseSeeder: BaseSeeder{name: "Mobile Operator Specifications"},
	}
}

// Seed implements the Seeder interface
func (moss *MobileOperatorSpecificationSeeder) Seed(db *gorm.DB) error {
	// Specification keys for mobile operators
	specificationKeys := []string{
		"Operator Name",
		"Established Year",
		"Headquarters Location",
		"Network Type",
		"Coverage Area",
		"Total Subscribers",
		"Total Base Stations",
		"2G Service Available",
		"3G Service Available",
		"4G/LTE Service Available",
		"5G Service Available",
		"Voice Call Service",
		"SMS Service",
		"Data Service",
		"MMS Service",
		"Internet Browsing",
		"Video Call Support",
		"Roaming Service",
		"International Roaming",
		"Mobile Money Service",
		"Mobile Money Partner",
		"USSD Support",
		"Prepaid Plans Available",
		"Postpaid Plans Available",
		"Customer Care Phone",
		"Customer Care Email",
		"Official Website",
		"Mobile App Name",
		"Android App Available",
		"iOS App Available",
		"Head Office Address",
		"Helpline Hours",
		"Complaint Resolution Time",
		"Network Speed (Average)",
		"Customer Satisfaction Rating",
		"CEO / MD Name",
		"Chairman Name",
		"Number of Employees",
		"Network Infrastructure Partners",
		"Technology Provider",
		"ISO Certifications",
		"Corporate Social Responsibility Programs",
		"Service Availability SLA",
		"Call Drop Rate",
		"Voice Quality Rating",
		"Data Speed Rating",
		"Network Reliability",
		"Emergency Services Support",
		"Disaster Recovery Support",
	}

	// Create or find specification keys
	for _, keyName := range specificationKeys {
		var existingKey models.SpecificationKeyModel
		result := db.Where("specification_key = ?", keyName).First(&existingKey)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new specification key
			specKey := models.SpecificationKeyModel{
				SpecificationKey: keyName,
			}
			if err := db.Create(&specKey).Error; err != nil {
				continue
			}
		}
	}

	// Mobile operator data for specifications
	operatorsData := map[string]map[string]string{
		"banglalink": {
			"Operator Name":                "Banglalink",
			"Established Year":             "1989",
			"Headquarters Location":        "Dhaka, Bangladesh",
			"Network Type":                 "2G, 3G, 4G LTE",
			"Total Subscribers":            "35+ Million",
			"Total Base Stations":          "13,000+",
			"2G Service Available":         "Yes",
			"3G Service Available":         "Yes",
			"4G/LTE Service Available":     "Yes",
			"5G Service Available":         "No",
			"Voice Call Service":           "Yes",
			"SMS Service":                  "Yes",
			"Data Service":                 "Yes",
			"Internet Browsing":            "Yes",
			"Roaming Service":              "Yes",
			"International Roaming":        "Yes",
			"Mobile Money Service":         "Banglalink Money",
			"Prepaid Plans Available":      "Yes",
			"Postpaid Plans Available":     "Yes",
			"Customer Care Phone":          "16000",
			"Official Website":             "https://www.banglalink.com.bd",
			"Mobile App Name":              "BL Live",
			"Android App Available":        "Yes",
			"iOS App Available":            "Yes",
			"Customer Satisfaction Rating": "Good",
			"Network Reliability":          "Excellent",
		},
		"grameenphone": {
			"Operator Name":                "Grameenphone",
			"Established Year":             "1997",
			"Headquarters Location":        "Dhaka, Bangladesh",
			"Network Type":                 "2G, 3G, 4G LTE",
			"Total Subscribers":            "80+ Million",
			"Total Base Stations":          "19,000+",
			"2G Service Available":         "Yes",
			"3G Service Available":         "Yes",
			"4G/LTE Service Available":     "Yes",
			"5G Service Available":         "No",
			"Voice Call Service":           "Yes",
			"SMS Service":                  "Yes",
			"Data Service":                 "Yes",
			"Internet Browsing":            "Yes",
			"Roaming Service":              "Yes",
			"International Roaming":        "Yes",
			"Mobile Money Service":         "bKash",
			"Prepaid Plans Available":      "Yes",
			"Postpaid Plans Available":     "Yes",
			"Customer Care Phone":          "121",
			"Official Website":             "https://www.grameenphone.com",
			"Mobile App Name":              "MyGP",
			"Android App Available":        "Yes",
			"iOS App Available":            "Yes",
			"Customer Satisfaction Rating": "Excellent",
			"Network Reliability":          "Excellent",
		},
		"robi": {
			"Operator Name":                "Robi",
			"Established Year":             "2004",
			"Headquarters Location":        "Dhaka, Bangladesh",
			"Network Type":                 "2G, 3G, 4G LTE",
			"Total Subscribers":            "50+ Million",
			"Total Base Stations":          "15,000+",
			"2G Service Available":         "Yes",
			"3G Service Available":         "Yes",
			"4G/LTE Service Available":     "Yes",
			"5G Service Available":         "No",
			"Voice Call Service":           "Yes",
			"SMS Service":                  "Yes",
			"Data Service":                 "Yes",
			"Internet Browsing":            "Yes",
			"Roaming Service":              "Yes",
			"International Roaming":        "Yes",
			"Mobile Money Service":         "Robi Wallet",
			"Prepaid Plans Available":      "Yes",
			"Postpaid Plans Available":     "Yes",
			"Customer Care Phone":          "18600",
			"Official Website":             "https://www.robi.com.bd",
			"Mobile App Name":              "My Robi",
			"Android App Available":        "Yes",
			"iOS App Available":            "Yes",
			"Customer Satisfaction Rating": "Good",
			"Network Reliability":          "Good",
		},
		"teletalk": {
			"Operator Name":                "Teletalk",
			"Established Year":             "2004",
			"Headquarters Location":        "Dhaka, Bangladesh",
			"Network Type":                 "2G, 3G, 4G LTE",
			"Total Subscribers":            "10+ Million",
			"Total Base Stations":          "3,000+",
			"2G Service Available":         "Yes",
			"3G Service Available":         "Yes",
			"4G/LTE Service Available":     "Yes",
			"5G Service Available":         "No",
			"Voice Call Service":           "Yes",
			"SMS Service":                  "Yes",
			"Data Service":                 "Yes",
			"Internet Browsing":            "Yes",
			"Roaming Service":              "Yes",
			"International Roaming":        "Yes",
			"Mobile Money Service":         "Teletalk Money",
			"Prepaid Plans Available":      "Yes",
			"Postpaid Plans Available":     "Yes",
			"Customer Care Phone":          "1500",
			"Official Website":             "https://www.teletalk.com.bd",
			"Mobile App Name":              "My Teletalk",
			"Android App Available":        "Yes",
			"iOS App Available":            "Yes",
			"Customer Satisfaction Rating": "Fair",
			"Network Reliability":          "Fair",
		},
		"airtel": {
			"Operator Name":                "Airtel",
			"Established Year":             "2009",
			"Headquarters Location":        "Dhaka, Bangladesh",
			"Network Type":                 "2G, 3G, 4G LTE",
			"Total Subscribers":            "20+ Million",
			"Total Base Stations":          "5,000+",
			"2G Service Available":         "Yes",
			"3G Service Available":         "Yes",
			"4G/LTE Service Available":     "Yes",
			"5G Service Available":         "No",
			"Voice Call Service":           "Yes",
			"SMS Service":                  "Yes",
			"Data Service":                 "Yes",
			"Internet Browsing":            "Yes",
			"Roaming Service":              "Yes",
			"International Roaming":        "Yes",
			"Mobile Money Service":         "Airtel Money",
			"Prepaid Plans Available":      "Yes",
			"Postpaid Plans Available":     "Yes",
			"Customer Care Phone":          "121",
			"Official Website":             "https://www.airtel.com.bd",
			"Mobile App Name":              "My Airtel",
			"Android App Available":        "Yes",
			"iOS App Available":            "Yes",
			"Customer Satisfaction Rating": "Good",
			"Network Reliability":          "Good",
		},
	}

	// Find Mobile Operator category (ID: 80)
	var mobileCategory models.CategoryModel
	if err := db.Where("id = ?", 80).First(&mobileCategory).Error; err != nil {
		// If not found by ID, try by slug
		if err := db.Where("slug = ?", "mobile-operators").First(&mobileCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	// Create specifications for each operator
	for operatorSlug, specs := range operatorsData {
		// Find product by slug
		var product models.ProductModel
		if err := db.Where("slug = ?", operatorSlug).First(&product).Error; err != nil {
			continue // Skip if product not found
		}

		// Add specifications for this operator
		for specKey, specValue := range specs {
			if specValue == "" {
				continue
			}

			// Find specification key
			var specKeyModel models.SpecificationKeyModel
			if err := db.Where("specification_key = ?", specKey).First(&specKeyModel).Error; err != nil {
				continue
			}

			// Check if specification already exists
			var existingSpec models.SpecificationModel
			result := db.Where("product_id = ? AND specification_key_id = ?", product.ID, specKeyModel.ID).First(&existingSpec)

			if result.Error == gorm.ErrRecordNotFound {
				// Create new specification
				specification := models.SpecificationModel{
					ProductID:          product.ID,
					SpecificationKeyID: specKeyModel.ID,
					Value:              specValue,
				}

				if err := db.Create(&specification).Error; err != nil {
					continue
				}
			}
		}
	}

	return nil
}
