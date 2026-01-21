package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BankBrandSeeder handles seeding bank brands
type BankBrandSeeder struct {
	BaseSeeder
}

// NewBankBrandSeeder creates a new bank brand seeder
func NewBankBrandSeeder() *BankBrandSeeder {
	return &BankBrandSeeder{
		BaseSeeder: BaseSeeder{name: "Bank Brands"},
	}
}

// Seed implements the Seeder interface
func (bbs *BankBrandSeeder) Seed(db *gorm.DB) error {
	// All Bangladeshi banks (46 total) - Removed foreign banks: Citibank, Commercial Bank of Ceylon, Habib Bank, HSBC Bank, National Bank of Pakistan, State Bank of India
	brands := []struct {
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
		{Name: "Islami Bank", Slug: "islami-bank"},
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

	for _, brand := range brands {
		var existingBrand models.BrandModel
		result := db.Where("slug = ?", brand.Slug).First(&existingBrand)

		if result.Error == gorm.ErrRecordNotFound {
			newBrand := &models.BrandModel{
				Name:   brand.Name,
				Slug:   brand.Slug,
				Status: 1,
			}

			if err := db.Create(newBrand).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
