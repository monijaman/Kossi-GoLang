package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BankProductSeeder handles seeding bank products
type BankProductSeeder struct {
	BaseSeeder
}

// NewBankProductSeeder creates a new bank product seeder
func NewBankProductSeeder() *BankProductSeeder {
	return &BankProductSeeder{
		BaseSeeder: BaseSeeder{name: "Bank Products"},
	}
}

// Seed implements the Seeder interface
func (bps *BankProductSeeder) Seed(db *gorm.DB) error {
	// List of all Bangladeshi banks (52 total) - Removed foreign banks
	banks := []struct {
		Name string
		Slug string
	}{
		{Name: "AB Bank", Slug: "ab-bank"},
		{Name: "Agrani Bank", Slug: "agrani-bank"},
		{Name: "Al-Arafah Islami Bank", Slug: "al-arafah-islami-bank"},
		{Name: "Bangladesh Commerce Bank", Slug: "bangladesh-commerce-bank"},
		{Name: "Bangladesh Development Bank", Slug: "bangladesh-development-bank"},
		{Name: "Bangladesh Krishi Bank", Slug: "bangladesh-krishi-bank"},
		{Name: "Bank Alfalah", Slug: "bank-alfalah"},
		{Name: "Bank Asia", Slug: "bank-asia"},
		{Name: "BASIC Bank", Slug: "basic-bank"},
		{Name: "Bengal Commercial Bank", Slug: "bengal-bank"},
		{Name: "BRAC Bank", Slug: "brac-bank"},
		{Name: "Citizens Bank", Slug: "citizens-bank"},
		{Name: "City Bank", Slug: "city-bank"},
		{Name: "Community Bank", Slug: "community-bank"},
		{Name: "Dhaka Bank", Slug: "dhaka-bank"},
		{Name: "Dutch-Bangla Bank", Slug: "dutch-bangla-bank"},
		{Name: "Eastern Bank", Slug: "eastern-bank"},
		{Name: "EXIM Bank", Slug: "exim-bank"},
		{Name: "First Security Islami Bank", Slug: "first-security-islami-bank"},
		{Name: "Global Islami Bank", Slug: "global-islami-bank"},
		{Name: "ICB Islamic Bank", Slug: "icb-islamic-bank"},
		{Name: "IFIC Bank", Slug: "ific-bank"},
		{Name: "Islami Bank Bangladesh", Slug: "islami-bank-bangladesh"},
		{Name: "Jamuna Bank", Slug: "jamuna-bank"},
		{Name: "Janata Bank", Slug: "janata-bank"},
		{Name: "Meghna Bank", Slug: "meghna-bank"},
		{Name: "Mercantile Bank", Slug: "mercantile-bank"},
		{Name: "Midland Bank", Slug: "midland-bank"},
		{Name: "Modhumoti Bank", Slug: "modhumoti-bank"},
		{Name: "Mutual Trust Bank", Slug: "mutual-trust-bank"},
		{Name: "National Bank", Slug: "national-bank"},
		{Name: "NCC Bank", Slug: "ncc-bank"},
		{Name: "NRB Bank", Slug: "nrb-bank"},
		{Name: "NRBC Bank", Slug: "nrbc-bank"},
		{Name: "One Bank", Slug: "one-bank"},
		{Name: "Padma Bank", Slug: "padma-bank"},
		{Name: "Premier Bank", Slug: "premier-bank"},
		{Name: "Prime Bank", Slug: "prime-bank"},
		{Name: "Probashi Kallyan Bank", Slug: "probashi-kallyan-bank"},
		{Name: "Pubali Bank", Slug: "pubali-bank"},
		{Name: "Rajshahi Krishi Unnayan Bank", Slug: "rajshahi-krishi-unnayan-bank"},
		{Name: "Rupali Bank", Slug: "rupali-bank"},
		{Name: "SBAC Bank", Slug: "sbac-bank"},
		{Name: "Shahjalal Islami Bank", Slug: "shahjalal-islami-bank"},
		{Name: "Shimanto Bank", Slug: "shimanto-bank"},
		{Name: "Social Islami Bank", Slug: "social-islami-bank"},
		{Name: "Sonali Bank", Slug: "sonali-bank"},
		{Name: "Southeast Bank", Slug: "southeast-bank"},
		{Name: "Standard Bank", Slug: "standard-bank"},
		{Name: "Standard Chartered Bank", Slug: "standard-chartered-bank"},
		{Name: "Trust Bank", Slug: "trust-bank"},
		{Name: "Union Bank", Slug: "union-bank"},
		{Name: "United Commercial Bank", Slug: "united-commercial-bank"},
		{Name: "Uttara Bank", Slug: "uttara-bank"},
		{Name: "Woori Bank", Slug: "woori-bank"},
	}

	// Find Banking category
	var bankingCategory models.CategoryModel
	if err := db.Where("id = ?", 10).First(&bankingCategory).Error; err != nil {
		if err := db.Where("slug = ?", "banking").First(&bankingCategory).Error; err != nil {
			return err
		}
	}

	// Create products for each bank
	for _, bank := range banks {
		// Find brand by slug
		var brand models.BrandModel
		if err := db.Where("slug = ?", bank.Slug).First(&brand).Error; err != nil {
			continue // Skip if brand not found
		}

		// Check if product already exists
		var existingProduct models.ProductModel
		result := db.Where("brand_id = ? AND slug = ?", brand.ID, bank.Slug).First(&existingProduct)

		if result.Error == gorm.ErrRecordNotFound {
			// Create product
			description := bank.Name + " - Banking services and information"
			product := models.ProductModel{
				Name:        bank.Name,
				Slug:        bank.Slug,
				Description: &description,
				BrandID:     &brand.ID,
				CategoryID:  &bankingCategory.ID,
				Status:      1,
			}

			if err := db.Create(&product).Error; err != nil {
				continue
			}
		}
	}

	return nil
}
