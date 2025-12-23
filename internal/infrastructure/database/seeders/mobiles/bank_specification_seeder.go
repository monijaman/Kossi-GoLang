package seeders

import (
	"encoding/json"
	"kossti/internal/infrastructure/database/models"
	"os"
	"strings"

	"gorm.io/gorm"
)

// BankSpecificationSeeder handles seeding bank specifications from bankspec.json
type BankSpecificationSeeder struct {
	BaseSeeder
}

// NewBankSpecificationSeeder creates a new bank specification seeder
func NewBankSpecificationSeeder() *BankSpecificationSeeder {
	return &BankSpecificationSeeder{
		BaseSeeder: BaseSeeder{name: "Bank Specifications"},
	}
}

// BankSpecData represents bank data from bankspec.json
type BankSpecData struct {
	BankName      string `json:"Bank Name"`
	Brand         string `json:"Brand"`
	BankType      string `json:"Type"`
	SwiftCode     string `json:"SWIFT Code"`
	Headquarters  string `json:"Headquarters"`
	Website       string `json:"Website"`
	PhoneNumber   string `json:"Customer Care Phone"`
	Established   string `json:"Established Year"`
	TotalBranches string `json:"Total Branches"`
	TotalATMs     string `json:"Total ATMs"`
	TotalAgents   string `json:"Total Agents"`
	MobileAppName string `json:"Mobile Banking App Name"`
	Chairman      string `json:"Chairman"`
	CEO           string `json:"Managing Director / CEO"`
}

// Seed implements the Seeder interface
func (bss *BankSpecificationSeeder) Seed(db *gorm.DB) error {
	// Load bankspec.json
	bankSpecs, err := loadBankSpecifications()
	if err != nil {
		return err
	}

	// Add fallback specs for banks not in bankspec.json
	fallbackSpecs := map[string]map[string]string{
		"shimanto-bank": {
			"Website":                "https://www.shimantobkash.com",
			"Bank Type":              "Private Commercial Bank",
			"Established Year":       "2009",
			"Customer Service Phone": "16477",
		},
		"woori-bank": {
			"Website":                "https://www.wooribank.com.bd",
			"Bank Type":              "Foreign Bank",
			"Established Year":       "2006",
			"Customer Service Phone": "16477",
			"SWIFT Code":             "WOORIBDD",
		},
	}

	// Add specifications for each bank from bankspec.json
	for _, bankSpec := range bankSpecs {
		// Normalize brand slug
		brandSlug := normalizeSlug(bankSpec.Brand)

		// Find product by brand slug (product seeder should have created them)
		var product models.ProductModel
		if err := db.Where("slug = ?", brandSlug).First(&product).Error; err != nil {
			continue // Skip if product not found
		}

		// Add specifications for this bank
		specs := extractSpecifications(bankSpec)
		for specKey, specValue := range specs {
			if specValue == "" {
				continue
			}

			// Find or create specification key
			var specKeyModel models.SpecificationKeyModel
			result := db.Where("specification_key = ?", specKey).First(&specKeyModel)

			if result.Error == gorm.ErrRecordNotFound {
				// Create new specification key
				specKeyModel = models.SpecificationKeyModel{
					SpecificationKey: specKey,
				}
				if err := db.Create(&specKeyModel).Error; err != nil {
					continue
				}
			}

			// Check if specification already exists
			var existingSpec models.SpecificationModel
			specResult := db.Where("product_id = ? AND specification_key_id = ?", product.ID, specKeyModel.ID).First(&existingSpec)

			if specResult.Error == gorm.ErrRecordNotFound {
				// Create specification
				spec := models.SpecificationModel{
					ProductID:          product.ID,
					SpecificationKeyID: specKeyModel.ID,
					Value:              specValue,
					Status:             1,
				}

				if err := db.Create(&spec).Error; err != nil {
					continue
				}
			}
		}
	}

	// Add fallback specifications for banks not in bankspec.json
	for bankSlug, specs := range fallbackSpecs {
		// Find product by slug
		var product models.ProductModel
		if err := db.Where("slug = ?", bankSlug).First(&product).Error; err != nil {
			continue // Skip if product not found
		}

		// Add each specification
		for specKey, specValue := range specs {
			if specValue == "" {
				continue
			}

			// Find or create specification key
			var specKeyModel models.SpecificationKeyModel
			result := db.Where("specification_key = ?", specKey).First(&specKeyModel)

			if result.Error == gorm.ErrRecordNotFound {
				// Create new specification key
				specKeyModel = models.SpecificationKeyModel{
					SpecificationKey: specKey,
				}
				if err := db.Create(&specKeyModel).Error; err != nil {
					continue
				}
			}

			// Check if specification already exists
			var existingSpec models.SpecificationModel
			specResult := db.Where("product_id = ? AND specification_key_id = ?", product.ID, specKeyModel.ID).First(&existingSpec)

			if specResult.Error == gorm.ErrRecordNotFound {
				// Create specification
				spec := models.SpecificationModel{
					ProductID:          product.ID,
					SpecificationKeyID: specKeyModel.ID,
					Value:              specValue,
					Status:             1,
				}

				if err := db.Create(&spec).Error; err != nil {
					continue
				}
			}
		}
	}

	return nil
}

// loadBankSpecifications loads bank data from bankspec.json
func loadBankSpecifications() ([]BankSpecData, error) {
	var specs []BankSpecData

	// Try loading from init-db/seeders
	jsonFile, err := os.ReadFile("init-db/seeders/bankspec.json")
	if err != nil {
		// Try from different path
		jsonFile, err = os.ReadFile("./init-db/seeders/bankspec.json")
		if err != nil {
			return specs, err
		}
	}

	err = json.Unmarshal(jsonFile, &specs)
	if err != nil {
		return specs, err
	}

	return specs, nil
}

// extractSpecifications extracts relevant specifications from bank data
func extractSpecifications(bank BankSpecData) map[string]string {
	specs := make(map[string]string)

	// Add specifications with values (skip empty ones)
	if bank.SwiftCode != "" {
		specs["SWIFT Code"] = bank.SwiftCode
	}
	if bank.Website != "" {
		specs["Website"] = bank.Website
	}
	if bank.BankType != "" {
		specs["Bank Type"] = bank.BankType
	}
	if bank.Headquarters != "" {
		specs["Headquarters"] = bank.Headquarters
	}
	if bank.PhoneNumber != "" {
		specs["Customer Service Phone"] = bank.PhoneNumber
	}
	if bank.Established != "" {
		specs["Established Year"] = bank.Established
	}
	if bank.TotalBranches != "" {
		specs["Total Branches"] = bank.TotalBranches
	}
	if bank.TotalATMs != "" {
		specs["Total ATMs"] = bank.TotalATMs
	}
	if bank.TotalAgents != "" {
		specs["Total Agents"] = bank.TotalAgents
	}
	if bank.MobileAppName != "" {
		specs["Mobile Banking App"] = bank.MobileAppName
	}
	if bank.Chairman != "" {
		specs["Chairman"] = bank.Chairman
	}
	if bank.CEO != "" {
		specs["CEO"] = bank.CEO
	}

	return specs
}

// normalizeSlug converts a bank name to a consistent slug format
func normalizeSlug(name string) string {
	// Map common brand variations to standard slugs
	name = strings.TrimSpace(name)

	// Abbreviation mappings
	mapping := map[string]string{
		"BRAC Bank":                    "brac-bank",
		"DBBL":                         "dutch-bangla-bank",
		"Dutch‑Bangla Bank":            "dutch-bangla-bank",
		"Dutch-Bangla Bank":            "dutch-bangla-bank",
		"Standard Chartered":           "standard-chartered-bank",
		"Standard Chartered Bank":      "standard-chartered-bank",
		"City Bank":                    "city-bank",
		"The City Bank":                "city-bank",
		"EBL":                          "eastern-bank",
		"Eastern Bank":                 "eastern-bank",
		"Bank Asia":                    "bank-asia",
		"IFIC Bank":                    "ific-bank",
		"IFIC":                         "ific-bank",
		"Trust Bank":                   "trust-bank",
		"Dhaka Bank":                   "dhaka-bank",
		"Islami Bank":                  "islami-bank",
		"Islami Bank Bangladesh":       "islami-bank",
		"IBBL":                         "islami-bank",
		"Agrani Bank":                  "agrani-bank",
		"Sonali Bank":                  "sonali-bank",
		"Rupali Bank":                  "rupali-bank",
		"Pubali Bank":                  "pubali-bank",
		"Bangladesh Krishi Bank":       "bangladesh-krishi-bank",
		"BKB":                          "bangladesh-krishi-bank",
		"Bangladesh Development Bank":  "bangladesh-development-bank",
		"BDBL":                         "bangladesh-development-bank",
		"Rajshahi Krishi Unnayan Bank": "rajshahi-krishi-unnayan-bank",
		"RKUB":                         "rajshahi-krishi-unnayan-bank",
		"Prime Bank":                   "prime-bank",
		"Premier Bank":                 "premier-bank",
		"Mutual Trust Bank":            "mutual-trust-bank",
		"MTB":                          "mutual-trust-bank",
		"National Bank":                "national-bank",
		"Jamuna Bank":                  "jamuna-bank",
		"Meghna Bank":                  "meghna-bank",
		"Midland Bank":                 "midland-bank",
		"NRB Bank":                     "nrb-bank",
		"One Bank":                     "one-bank",
		"South Bangla Bank":            "sbac-bank",
		"SBAC Bank":                    "sbac-bank",
		"SBAC":                         "sbac-bank",
		"Southeast Bank":               "southeast-bank",
		"Union Bank":                   "union-bank",
		"Bengal Commercial Bank":       "bengal-bank",
		"EXIM Bank":                    "exim-bank",
		"IFC Bank":                     "ifc-bank",
		"Shahjalal Islami Bank":        "shahjalal-islami-bank",
		"SJIBL":                        "shahjalal-islami-bank",
		"Social Islami Bank":           "social-islami-bank",
		"SIBL":                         "social-islami-bank",
		"First Security Islami Bank":   "first-security-islami-bank",
		"FSIBL":                        "first-security-islami-bank",
		"Al-Arafah Islami Bank":        "al-arafah-islami-bank",
	}

	if slug, exists := mapping[name]; exists {
		return slug
	}

	// Fallback to generic slug generation
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "/", "-")
	return slug
}
