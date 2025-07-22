package seeders

import (
	"gorm.io/gorm"
)

// BrandCategorySeeder handles seeding brand-category relationships
type BrandCategorySeeder struct {
	BaseSeeder
}

// NewBrandCategorySeeder creates a new brand-category seeder
func NewBrandCategorySeeder() *BrandCategorySeeder {
	return &BrandCategorySeeder{
		BaseSeeder: BaseSeeder{name: "Brand Categories"},
	}
}

// Seed implements the Seeder interface
func (bcs *BrandCategorySeeder) Seed(db *gorm.DB) error {
	// Brand-category relationships from the PHP seeder
	brandCategoryMappings := map[string][]string{
		"Adidas": {"Footwear", "Clothing", "Sports Equipment"},
		"Airbus": {"Aerospace", "Industrial Equipment"},
		"Amazon": {"E-commerce", "Technology", "Books", "Electronics"},
		"Apple": {"Electronics", "Mobile Phones", "Laptops", "Technology"},
		"Audi": {"Automotive", "Luxury Cars"},
		"Baidu": {"Technology", "Search Engines"},
		"Barclays": {"Financial Services", "Banking"},
		"BMW": {"Automotive", "Luxury Cars", "Motorcycles"},
		"Bosch": {"Industrial Equipment", "Home Appliances", "Automotive"},
		"Bose": {"Electronics", "Audio Equipment"},
		"Bridgestone": {"Automotive", "Tires"},
		"Budweiser": {"Beverages", "Alcoholic Beverages", "Beer"},
		"Canon": {"Electronics", "Cameras", "Photography"},
		"Chanel": {"Fashion", "Luxury Goods", "Perfumes & Fragrances"},
		"Chevrolet": {"Automotive", "Cars"},
		"Cisco": {"Technology", "Networking", "Telecommunications"},
		"Coca-Cola": {"Beverages", "Soft Drinks"},
		"Colgate": {"Personal Care", "Oral Care"},
		"DHL": {"Logistics", "Shipping"},
		"Disney": {"Entertainment", "Media", "Theme Parks"},
		"Dove": {"Personal Care", "Beauty Products"},
		"Dropbox": {"Technology", "Cloud Storage"},
		"eBay": {"E-commerce", "Online Marketplace"},
		"Emirates": {"Airlines", "Travel"},
		"Facebook": {"Technology", "Social Media"},
		"Ford": {"Automotive", "Cars", "Trucks"},
		"Gatorade": {"Beverages", "Sports Drinks"},
		"Google": {"Technology", "Search Engines", "Software"},
		"Gucci": {"Fashion", "Luxury Goods", "Accessories"},
		"H&M": {"Fashion", "Clothing", "Retail"},
		"Heineken": {"Beverages", "Alcoholic Beverages", "Beer"},
		"Hewlett-Packard": {"Technology", "Computers", "Printers"},
		"Honda": {"Automotive", "Cars", "Motorcycles"},
		"Huawei": {"Technology", "Mobile Phones", "Telecommunications"},
		"IBM": {"Technology", "Software", "Consulting"},
		"Intel": {"Technology", "Computer Hardware", "Semiconductors"},
		"Jaguar": {"Automotive", "Luxury Cars"},
		"Johnson & Johnson": {"Healthcare", "Personal Care", "Pharmaceuticals"},
		"JPMorgan Chase": {"Financial Services", "Banking", "Investment"},
		"KFC": {"Food & Beverages", "Fast Food", "Restaurants"},
		"Kia": {"Automotive", "Cars"},
		"L'Oréal": {"Beauty Products", "Personal Care", "Cosmetics"},
		"Levi's": {"Clothing", "Fashion", "Denim"},
		"LG": {"Electronics", "Home Appliances", "Mobile Phones"},
		"Lexus": {"Automotive", "Luxury Cars"},
		"Louis Vuitton": {"Fashion", "Luxury Goods", "Accessories"},
		"Mastercard": {"Financial Services", "Payment Services"},
		"McDonald's": {"Food & Beverages", "Fast Food", "Restaurants"},
		"Mercedes-Benz": {"Automotive", "Luxury Cars"},
		"Microsoft": {"Technology", "Software", "Computer Hardware"},
		"Mitsubishi": {"Automotive", "Industrial Equipment"},
		"Nestlé": {"Food & Beverages", "Confectionery", "Nutrition"},
		"Nike": {"Footwear", "Clothing", "Sports Equipment"},
		"Nissan": {"Automotive", "Cars"},
		"Olay": {"Beauty Products", "Personal Care", "Skincare"},
		"Pepsi": {"Beverages", "Soft Drinks"},
		"Pfizer": {"Healthcare", "Pharmaceuticals"},
		"Philips": {"Electronics", "Healthcare", "Lighting"},
		"Porsche": {"Automotive", "Luxury Cars", "Sports Cars"},
		"Prada": {"Fashion", "Luxury Goods", "Accessories"},
		"Procter & Gamble": {"Personal Care", "Household Products"},
		"Puma": {"Footwear", "Clothing", "Sports Equipment"},
		"Red Bull": {"Beverages", "Energy Drinks", "Sports"},
		"Rolex": {"Watches", "Luxury Goods", "Jewelry"},
		"Samsung": {"Electronics", "Mobile Phones", "Home Appliances"},
		"Sony": {"Electronics", "Entertainment", "Gaming"},
		"Starbucks": {"Food & Beverages", "Coffee", "Restaurants"},
		"Subway": {"Food & Beverages", "Fast Food", "Restaurants"},
		"T-Mobile": {"Telecommunications", "Mobile Services"},
		"Tesla": {"Automotive", "Electric Vehicles", "Technology"},
		"The North Face": {"Clothing", "Sports Equipment", "Outdoor Gear"},
		"Toyota": {"Automotive", "Cars", "Hybrid Vehicles"},
		"Visa": {"Financial Services", "Payment Services"},
		"Vodafone": {"Telecommunications", "Mobile Services"},
		"Volvo": {"Automotive", "Cars", "Trucks"},
		"Walmart": {"Retail", "Grocery", "E-commerce"},
		"Whirlpool": {"Home Appliances", "Kitchen Appliances"},
		"Yahoo": {"Technology", "Internet Services"},
		"Yves Saint Laurent": {"Fashion", "Luxury Goods", "Beauty Products"},
		"Zara": {"Fashion", "Clothing", "Retail"},
		"Abercrombie & Fitch": {"Fashion", "Clothing", "Retail"},
		"Acer": {"Technology", "Computers", "Electronics"},
		"Allianz": {"Financial Services", "Insurance"},
		"American Express": {"Financial Services", "Payment Services"},
		"Aston Martin": {"Automotive", "Luxury Cars", "Sports Cars"},
		"Ben & Jerry's": {"Food & Beverages", "Ice Cream", "Desserts"},
		"Bulgari": {"Jewelry", "Luxury Goods", "Watches"},
		"Burberry": {"Fashion", "Luxury Goods", "Clothing"},
		"Cadillac": {"Automotive", "Luxury Cars"},
		"Carlsberg": {"Beverages", "Alcoholic Beverages", "Beer"},
		"Chivas Regal": {"Beverages", "Alcoholic Beverages", "Whiskey"},
		"Coach": {"Fashion", "Luxury Goods", "Accessories"},
		"Corona": {"Beverages", "Alcoholic Beverages", "Beer"},
		"Crocs": {"Footwear", "Casual Wear"},
		"Danone": {"Food & Beverages", "Dairy Products", "Nutrition"},
		"Diesel": {"Fashion", "Clothing", "Accessories"},
		"Dior": {"Fashion", "Luxury Goods", "Beauty Products"},
		"Dell": {"Technology", "Computers", "Electronics"},
		"Duracell": {"Electronics", "Batteries"},
		"Ferrari": {"Automotive", "Luxury Cars", "Sports Cars"},
		"Fendi": {"Fashion", "Luxury Goods", "Accessories"},
		"Fila": {"Footwear", "Clothing", "Sports Equipment"},
		"Fitbit": {"Electronics", "Fitness", "Wearables"},
		"GAP": {"Fashion", "Clothing", "Retail"},
		"Geox": {"Footwear", "Clothing"},
		"Glenfiddich": {"Beverages", "Alcoholic Beverages", "Whiskey"},
		"Havaianas": {"Footwear", "Flip-flops"},
		"Hugo Boss": {"Fashion", "Clothing", "Luxury Goods"},
		"Jack Daniel's": {"Beverages", "Alcoholic Beverages", "Whiskey"},
		"Jameson": {"Beverages", "Alcoholic Beverages", "Whiskey"},
		"Kering": {"Fashion", "Luxury Goods"},
		"Kiehl's": {"Beauty Products", "Personal Care", "Skincare"},
		"Lacoste": {"Fashion", "Clothing", "Sports Equipment"},
		"Lancôme": {"Beauty Products", "Cosmetics", "Skincare"},
		"Land Rover": {"Automotive", "SUVs", "Off-road Vehicles"},
		"Lindt": {"Food & Beverages", "Confectionery", "Chocolate"},
		"Lululemon": {"Clothing", "Sports Equipment", "Fitness"},
		"Maserati": {"Automotive", "Luxury Cars", "Sports Cars"},
		"Maui Jim": {"Accessories", "Sunglasses", "Eyewear"},
		"Michael Kors": {"Fashion", "Luxury Goods", "Accessories"},
		"Montblanc": {"Luxury Goods", "Writing Instruments", "Accessories"},
		"Moët & Chandon": {"Beverages", "Alcoholic Beverages", "Champagne"},
		"Nivea": {"Personal Care", "Beauty Products", "Skincare"},
		"Oreo": {"Food & Beverages", "Confectionery", "Cookies"},
		"Patek Philippe": {"Watches", "Luxury Goods", "Jewelry"},
		"Peugeot": {"Automotive", "Cars"},
		"Ray-Ban": {"Accessories", "Sunglasses", "Eyewear"},
		"Reebok": {"Footwear", "Clothing", "Sports Equipment"},
		"Ralph Lauren": {"Fashion", "Clothing", "Luxury Goods"},
		"Remy Martin": {"Beverages", "Alcoholic Beverages", "Cognac"},
		"Ritz-Carlton": {"Hotels", "Hospitality", "Luxury Services"},
		"Ryanair": {"Airlines", "Travel"},
		"Saint Laurent": {"Fashion", "Luxury Goods", "Accessories"},
		"Swarovski": {"Jewelry", "Luxury Goods", "Crystals"},
		"Tag Heuer": {"Watches", "Luxury Goods", "Sports"},
		"Tiffany & Co.": {"Jewelry", "Luxury Goods", "Accessories"},
		"Tommy Hilfiger": {"Fashion", "Clothing", "Accessories"},
		"Tory Burch": {"Fashion", "Luxury Goods", "Accessories"},
		"Under Armour": {"Footwear", "Clothing", "Sports Equipment"},
		"Vans": {"Footwear", "Clothing", "Skateboarding"},
		"Versace": {"Fashion", "Luxury Goods", "Accessories"},
		"Vichy": {"Beauty Products", "Personal Care", "Skincare"},
		"Vivienne Westwood": {"Fashion", "Luxury Goods", "Accessories"},
		"Wella": {"Beauty Products", "Hair Care"},
		"Yamaha": {"Motorcycles", "Musical Instruments", "Electronics"},
		"Zegna": {"Fashion", "Luxury Goods", "Clothing"},
		"Nokia": {"Technology", "Mobile Phones", "Telecommunications"},
		"Motorola": {"Technology", "Mobile Phones", "Telecommunications"},
		"HTC": {"Technology", "Mobile Phones", "Electronics"},
		"Cooper's": {"Beverages", "Alcoholic Beverages", "Beer"},
		"Hunter": {"Footwear", "Outdoor Gear", "Clothing"},
		"Nasir Group": {"Conglomerate", "Multiple Industries"},
	}

	for brandName, categoryNames := range brandCategoryMappings {
		// Get the brand
		brand, err := CreateOrFindBrand(db, brandName, GenerateSlug(brandName))
		if err != nil {
			return err
		}

		for _, categoryName := range categoryNames {
			// Get the category
			category, err := CreateOrFindCategory(db, categoryName, GenerateSlug(categoryName))
			if err != nil {
				return err
			}

			// Create the brand-category relationship
			if err := CreateBrandCategoryRelation(db, brand.ID, category.ID); err != nil {
				return err
			}
		}
	}

	return nil
}
