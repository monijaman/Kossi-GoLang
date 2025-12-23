package seeders

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// CategorySeeder handles seeding categories
type CategorySeeder struct {
	BaseSeeder
}

// NewCategorySeeder creates a new category seeder
func NewCategorySeeder() *CategorySeeder {
	return &CategorySeeder{
		BaseSeeder: BaseSeeder{name: "Categories"},
	}
}

// Seed implements the Seeder interface
func (cs *CategorySeeder) Seed(db *gorm.DB) error {
	categories := []struct {
		Name string
		Slug string
	}{
		{"Accessories", "accessories"},
		{"Aerospace", "aerospace"},
		{"Airline", "airline"},
		{"Alcoholic Beverages", "alcoholic-beverages"},
		{"Audio Equipment", "audio-equipment"},
		{"Automotive", "automotive"},
		{"Baby Products", "baby-products"},
		{"Bags & Luggage", "bags-luggage"},
		{"Bakery", "bakery"},
		{"Banking", "banking"},
		{"Batteries", "batteries"},
		{"Beauty Products", "beauty-products"},
		{"Beer", "beer"},
		{"Beverages", "beverages"},
		{"Books", "books"},
		{"Bulbs", "bulbs"},
		{"Cameras", "cameras"},
		{"Cars", "cars"},
		{"Casual Wear", "casual-wear"},
		{"Champagne", "champagne"},
		{"Clothing", "clothing"},
		{"Cloud Storage", "cloud-storage"},
		{"Coffee", "coffee"},
		{"Cognac", "cognac"},
		{"Computers", "computers"},
		{"Computer Hardware", "computer-hardware"},
		{"Confectionery", "confectionery"},
		{"Conglomerate", "conglomerate"},
		{"Consulting", "consulting"},
		{"Cookies", "cookies"},
		{"Cosmetics", "cosmetics"},
		{"Craft Supplies", "craft-supplies"},
		{"Crystal", "crystal"},
		{"Dairy Products", "dairy-products"},
		{"Denim", "denim"},
		{"Electric Vehicles", "electric-vehicles"},
		{"Electronics", "electronics"},
		{"Energy Drinks", "energy-drinks"},
		{"Entertainment", "entertainment"},
		{"Eyewear", "eyewear"},
		{"Fashion", "fashion"},
		{"Fast Food", "fast-food"},
		{"Financial Services", "financial-services"},
		{"Fitness", "fitness"},
		{"Fitness Equipment", "fitness-equipment"},
		{"Flip-Flops", "flip-flops"},
		{"Food & Beverage", "food-beverage"},
		{"Footwear", "footwear"},
		{"Fragrances & Scents", "fragrances-scents"},
		{"Furniture", "furniture"},
		{"Gaming", "gaming"},
		{"Garden & Outdoor", "garden-outdoor"},
		{"Glass", "glass"},
		{"Groceries", "groceries"},
		{"Hair Care", "hair-care"},
		{"Health & Beauty", "health-beauty"},
		{"Healthcare", "healthcare"},
		{"Home Appliances", "home-appliances"},
		{"Hospitality", "hospitality"},
		{"Hotels", "hotels"},
		{"Household Goods", "household-goods"},
		{"Hybrid Vehicles", "hybrid-vehicles"},
		{"Ice Cream", "ice-cream"},
		{"Industrial Equipment", "industrial-equipment"},
		{"Insurance", "insurance"},
		{"Internet Services", "internet-services"},
		{"Investment", "investment"},
		{"Jewelry", "jewelry"},
		{"Kitchen Appliances", "kitchen-appliances"},
		{"Kitchenware", "kitchenware"},
		{"Laptops", "laptops"},
		{"Light", "light"},
		{"Logistics", "logistics"},
		{"Luxury Cars", "luxury-cars"},
		{"Luxury Goods", "luxury-goods"},
		{"Luxury Services", "luxury-services"},
		{"Media", "media"},
		{"Melamine", "melamine"},
		{"Mobile Phones", "mobile-phones"},
		{"Mobile Services", "mobile-services"},
		{"Motorcycles", "motorcycles"},
		{"Multiple Industries", "multiple-industries"},
		{"Music Instruments", "music-instruments"},
		{"Musical Instruments", "musical-instruments"},
		{"Networking", "networking"},
		{"Nutrition", "nutrition"},
		{"Office Supplies", "office-supplies"},
		{"Online Marketplace", "online-marketplace"},
		{"Oral Care", "oral-care"},
		{"Outdoor Gear", "outdoor-gear"},
		{"Packaging", "packaging"},
		{"Payment Services", "payment-services"},
		{"Perfumes", "perfumes"},
		{"Personal Care", "personal-care"},
		{"Pet Supplies", "pet-supplies"},
		{"Pharmaceuticals", "pharmaceuticals"},
		{"Photography", "photography"},
		{"Printers", "printers"},
		{"Printing", "printing"},
		{"Refrigerators", "refrigerators"},
		{"Restaurants", "restaurants"},
		{"Retail", "retail"},
		{"Search Engine", "search-engine"},
		{"Semiconductors", "semiconductors"},
		{"Shipping", "shipping"},
		{"Skateboarding", "skateboarding"},
		{"Skincare", "skincare"},
		{"Social Media", "social-media"},
		{"Soft Drinks", "soft-drinks"},
		{"Software", "software"},
		{"Sports", "sports"},
		{"Sports Cars", "sports-cars"},
		{"Sports Drinks", "sports-drinks"},
		{"Sports Equipment", "sports-equipment"},
		{"Stationery", "stationery"},
		{"Sweets", "sweets"},
		{"Technology", "technology"},
		{"Telecommunications", "telecommunications"},
		{"Theme Parks", "theme-parks"},
		{"Tires", "tires"},
		{"Toys", "toys"},
		{"Travel", "travel"},
		{"Trucks", "trucks"},
		{"TV", "tv"},
		{"Watches", "watches"},
		{"Wearables", "wearables"},
		{"Whisky", "whisky"},
		{"Writing Instruments", "writing-instruments"},
	}

	for _, cat := range categories {
		category := &entities.Category{
			Name: cat.Name,
			Slug: cat.Slug,
		}

		var categoryModel models.CategoryModel
		// Check if category already exists
		if err := db.Where("name = ?", cat.Name).First(&categoryModel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new category
				categoryModel.FromEntity(category)
				if err := db.Create(&categoryModel).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
