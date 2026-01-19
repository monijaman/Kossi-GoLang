package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MobileOperatorFormGeneratorSeeder handles seeding form generator data for Mobile Operators category
type MobileOperatorFormGeneratorSeeder struct {
	BaseSeeder
}

// NewMobileOperatorFormGeneratorSeeder creates a new mobile operator form generator seeder
func NewMobileOperatorFormGeneratorSeeder() *MobileOperatorFormGeneratorSeeder {
	return &MobileOperatorFormGeneratorSeeder{
		BaseSeeder: BaseSeeder{name: "Mobile Operator Form Generator"},
	}
}

// Seed implements the Seeder interface
func (mofgs *MobileOperatorFormGeneratorSeeder) Seed(db *gorm.DB) error {
	// Mobile operator specification keys for form generation
	mobileOperatorSpecificationKeys := []string{
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

	// Get all specification key IDs for mobile operators
	var specKeyIDs []uint
	for _, keyName := range mobileOperatorSpecificationKeys {
		var specKey models.SpecificationKeyModel
		if err := db.Where("specification_key = ?", keyName).First(&specKey).Error; err == nil {
			specKeyIDs = append(specKeyIDs, specKey.ID)
		}
	}

	// Convert IDs to JSON
	specKeyIDsJSON, err := json.Marshal(specKeyIDs)
	if err != nil {
		return err
	}

	// Find Mobile Operator category (ID: 80)
	var mobileCategory models.CategoryModel
	if err := db.Where("id = ?", 80).First(&mobileCategory).Error; err != nil {
		// If not found by ID, try by slug
		if err := db.Where("slug = ?", "mobile-operators").First(&mobileCategory).Error; err != nil {
			return nil // Skip if category not found
		}
	}

	// Check if form generator entry already exists for mobile operator category
	var existingForm models.FormGeneratorModel
	result := db.Where("category_id = ?", mobileCategory.ID).First(&existingForm)

	if result.Error == gorm.ErrRecordNotFound {
		// Create form generator entry
		formGenerator := &models.FormGeneratorModel{
			CategoryID:      mobileCategory.ID,
			SpecificationID: string(specKeyIDsJSON),
			Status:          "active",
		}

		if err := db.Create(formGenerator).Error; err != nil {
			return err
		}
	}

	return nil
}
