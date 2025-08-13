package seeders

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BrandSeeder handles seeding brands
type BrandSeeder struct {
	BaseSeeder
}

// NewBrandSeeder creates a new brand seeder
func NewBrandSeeder() *BrandSeeder {
	return &BrandSeeder{
		BaseSeeder: BaseSeeder{name: "Brands"},
	}
}

// Seed implements the Seeder interface
func (bs *BrandSeeder) Seed(db *gorm.DB) error {
	brands := []struct {
		Name   string
		Slug   string
		Status int
	}{
		{"Adidas", "adidas", 1},
		{"Airbus", "airbus", 1},
		{"Amazon", "amazon", 1},
		{"Apple", "apple", 1},
		{"Audi", "audi", 1},
		{"Baidu", "baidu", 1},
		{"Barclays", "barclays", 1},
		{"BMW", "bmw", 1},
		{"Bosch", "bosch", 1},
		{"Bose", "bose", 1},
		{"Bridgestone", "bridgestone", 1},
		{"Budweiser", "budweiser", 1},
		{"Canon", "canon", 1},
		{"Chanel", "chanel", 1},
		{"Chevrolet", "chevrolet", 1},
		{"Cisco", "cisco", 1},
		{"Coca-Cola", "coca-cola", 1},
		{"Colgate", "colgate", 1},
		{"DHL", "dhl", 1},
		{"Disney", "disney", 1},
		{"Dove", "dove", 1},
		{"Dropbox", "dropbox", 1},
		{"eBay", "ebay", 1},
		{"Emirates", "emirates", 1},
		{"Facebook", "facebook", 1},
		{"Ford", "ford", 1},
		{"Gatorade", "gatorade", 1},
		{"Google", "google", 1},
		{"Gucci", "gucci", 1},
		{"H&M", "h-m", 1},
		{"Heineken", "heineken", 1},
		{"Hewlett-Packard", "hewlett-packard", 1},
		{"Honda", "honda", 1},
		{"Huawei", "huawei", 1},
		{"IBM", "ibm", 1},
		{"Intel", "intel", 1},
		{"Jaguar", "jaguar", 1},
		{"Johnson & Johnson", "johnson-johnson", 1},
		{"JPMorgan Chase", "jpmorgan-chase", 1},
		{"KFC", "kfc", 1},
		{"Kia", "kia", 1},
		{"L'Oréal", "l-oreal", 1},
		{"Levi's", "levi-s", 1},
		{"LG", "lg", 1},
		{"Lexus", "lexus", 1},
		{"Louis Vuitton", "louis-vuitton", 1},
		{"Mastercard", "mastercard", 1},
		{"McDonald's", "mcdonald-s", 1},
		{"Mercedes-Benz", "mercedes-benz", 1},
		{"Microsoft", "microsoft", 1},
		{"Mitsubishi", "mitsubishi", 1},
		{"Nestlé", "nestle", 1},
		{"Nike", "nike", 1},
		{"Nissan", "nissan", 1},
		{"Olay", "olay", 1},
		{"Pepsi", "pepsi", 1},
		{"Pfizer", "pfizer", 1},
		{"Philips", "philips", 1},
		{"Porsche", "porsche", 1},
		{"Prada", "prada", 1},
		{"Procter & Gamble", "procter-gamble", 1},
		{"Puma", "puma", 1},
		{"Red Bull", "red-bull", 1},
		{"Rolex", "rolex", 1},
		{"Samsung", "samsung", 1},
		{"Sony", "sony", 1},
		{"Starbucks", "starbucks", 1},
		{"Subway", "subway", 1},
		{"T-Mobile", "t-mobile", 1},
		{"Tesla", "tesla", 1},
		{"The North Face", "the-north-face", 1},
		{"Toyota", "toyota", 1},
		{"Visa", "visa", 1},
		{"Vodafone", "vodafone", 1},
		{"Volvo", "volvo", 1},
		{"Walmart", "walmart", 1},
		{"Whirlpool", "whirlpool", 1},
		{"Yahoo", "yahoo", 1},
		{"Yves Saint Laurent", "yves-saint-laurent", 1},
		{"Zara", "zara", 1},
		{"Abercrombie & Fitch", "abercrombie-fitch", 1},
		{"Acer", "acer", 1},
		{"Allianz", "allianz", 1},
		{"American Express", "american-express", 1},
		{"Aston Martin", "aston-martin", 1},
		{"Ben & Jerry's", "ben-jerry-s", 1},
		{"Bulgari", "bulgari", 1},
		{"Burberry", "burberry", 1},
		{"Cadillac", "cadillac", 1},
		{"Carlsberg", "carlsberg", 1},
		{"Chivas Regal", "chivas-regal", 1},
		{"Coach", "coach", 1},
		{"Corona", "corona", 1},
		{"Crocs", "crocs", 1},
		{"Danone", "danone", 1},
		{"Diesel", "diesel", 1},
		{"Dior", "dior", 1},
		{"Dell", "dell", 1},
		{"Duracell", "duracell", 1},
		{"Ferrari", "ferrari", 1},
		{"Fendi", "fendi", 1},
		{"Fila", "fila", 1},
		{"Fitbit", "fitbit", 1},
		{"GAP", "gap", 1},
		{"Geox", "geox", 1},
		{"Glenfiddich", "glenfiddich", 1},
		{"Havaianas", "havaianas", 1},
		{"Hugo Boss", "hugo-boss", 1},
		{"Jack Daniel's", "jack-daniel-s", 1},
		{"Jameson", "jameson", 1},
		{"Kering", "kering", 1},
		{"Kiehl's", "kiehl-s", 1},
		{"Lacoste", "lacoste", 1},
		{"Lancôme", "lancome", 1},
		{"Land Rover", "land-rover", 1},
		{"Lindt", "lindt", 1},
		{"Lululemon", "lululemon", 1},
		{"Maserati", "maserati", 1},
		{"Maui Jim", "maui-jim", 1},
		{"Michael Kors", "michael-kors", 1},
		{"Montblanc", "montblanc", 1},
		{"Moët & Chandon", "moet-chandon", 1},
		{"Nivea", "nivea", 1},
		{"Oreo", "oreo", 1},
		{"Patek Philippe", "patek-philippe", 1},
		{"Peugeot", "peugeot", 1},
		{"Ray-Ban", "ray-ban", 1},
		{"Reebok", "reebok", 1},
		{"Ralph Lauren", "ralph-lauren", 1},
		{"Remy Martin", "remy-martin", 1},
		{"Ritz-Carlton", "ritz-carlton", 1},
		{"Ryanair", "ryanair", 1},
		{"Saint Laurent", "saint-laurent", 1},
		{"Swarovski", "swarovski", 1},
		{"Tag Heuer", "tag-heuer", 1},
		{"Tiffany & Co.", "tiffany-co", 1},
		{"Tommy Hilfiger", "tommy-hilfiger", 1},
		{"Tory Burch", "tory-burch", 1},
		{"Under Armour", "under-armour", 1},
		{"Vans", "vans", 1},
		{"Versace", "versace", 1},
		{"Vichy", "vichy", 1},
		{"Vivienne Westwood", "vivienne-westwood", 1},
		{"Wella", "wella", 1},
		{"Yamaha", "yamaha", 1},
		{"Zegna", "zegna", 1},
		{"Nokia", "nokia", 1},
		{"Motorola", "motorola", 1},
		{"HTC", "htc", 1},
		{"Cooper's", "cooper-s", 1},
		{"Hunter", "hunter", 1},
		{"Nasir Group", "nasir-group", 1},
	}

	for _, brand := range brands {
		brandEntity := &entities.Brand{
			Name:   brand.Name,
			Slug:   brand.Slug,
			Status: brand.Status,
		}

		var brandModel models.BrandModel
		// Check if brand already exists
		if err := db.Where("name = ?", brand.Name).First(&brandModel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new brand
				brandModel.FromEntity(brandEntity)
				if err := db.Create(&brandModel).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
