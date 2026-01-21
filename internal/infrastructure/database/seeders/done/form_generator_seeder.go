package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// FormGeneratorSeeder handles seeding form generator data for Banking category
type FormGeneratorSeeder struct {
	BaseSeeder
}

// NewFormGeneratorSeeder creates a new form generator seeder
func NewFormGeneratorSeeder() *FormGeneratorSeeder {
	return &FormGeneratorSeeder{
		BaseSeeder: BaseSeeder{name: "Form Generator"},
	}
}

// Seed implements the Seeder interface
func (fgs *FormGeneratorSeeder) Seed(db *gorm.DB) error {
	// Banking category specification keys
	bankingSpecificationKeys := []string{
		"Bank Name",
		"Established Year",
		"Headquarters",
		"SWIFT Code",
		"Routing Number",
		"License Type",
		"Ownership",
		"Chairman",
		"Managing Director / CEO",
		"Total Branches",
		"Total ATMs",
		"Total Agents",
		"Core Banking System",
		"Internet Banking",
		"Mobile Banking App",
		"Mobile Banking App Name",
		"SMS Banking",
		"Phone Banking",
		"Debit Card",
		"Credit Card",
		"International Card Support",
		"Deposit Schemes",
		"Loan Schemes",
		"Islamic Banking Window",
		"Foreign Exchange Service",
		"Remittance Service",
		"Corporate Banking",
		"SME Banking",
		"Agricultural Loan Schemes",
		"Student Banking",
		"Women Banking",
		"Agent Banking Service",
		"Digital Wallet",
		"UPI / QR Payment Support",
		"ATM Network Partnership",
		"Visa / Mastercard / UnionPay Support",
		"Government Payment Support",
		"Utility Bill Payment Support",
		"Customer Care Phone",
		"Customer Care Email",
		"Website",
		"Facebook Page",
		"Head Office Address",
		"Helpline Availability",
		"Working Days",
		"Transaction Limit (Daily ATM)",
		"Transaction Limit (Daily App)",
		"Foreign Currency Account Support",
		"Nagad / bKash / Rocket / Upay Linked",
	}

	// Get all specification key IDs for banking
	var specKeyIDs []uint
	for _, keyName := range bankingSpecificationKeys {
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

	// Find Banking category (ID: 10)
	var bankingCategory models.CategoryModel
	if err := db.Where("id = ?", 10).First(&bankingCategory).Error; err != nil {
		// If not found, try by slug
		if err := db.Where("slug = ?", "banking").First(&bankingCategory).Error; err != nil {
			return err
		}
	}

	// Check if form generator entry already exists for banking category
	var existingForm models.FormGeneratorModel
	result := db.Where("category_id = ?", bankingCategory.ID).First(&existingForm)

	if result.Error == gorm.ErrRecordNotFound {
		// Create form generator entry
		formGenerator := &models.FormGeneratorModel{
			CategoryID:      bankingCategory.ID,
			SpecificationID: string(specKeyIDsJSON),
			Status:          "active",
		}

		if err := db.Create(formGenerator).Error; err != nil {
			return err
		}
	}

	return nil
}
