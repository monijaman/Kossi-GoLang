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
		Name string
		Slug string
	}{
		{"Adidas", "adidas"},
		{"Airbus", "airbus"},
		{"Amazon", "amazon"},
		{"Apple", "apple"},
		{"Audi", "audi"},
		{"Baidu", "baidu"},
		{"Barclays", "barclays"},
		{"BMW", "bmw"},
		{"Bosch", "bosch"},
		{"Bose", "bose"},
		{"Bridgestone", "bridgestone"},
		{"Budweiser", "budweiser"},
		{"Canon", "canon"},
		{"Chanel", "chanel"},
		{"Chevrolet", "chevrolet"},
		{"Cisco", "cisco"},
		{"Coca-Cola", "coca-cola"},
		{"Colgate", "colgate"},
		{"DHL", "dhl"},
		{"Disney", "disney"},
		{"Dove", "dove"},
		{"Dropbox", "dropbox"},
		{"eBay", "ebay"},
		{"Emirates", "emirates"},
		{"Facebook", "facebook"},
		{"Ford", "ford"},
		{"Gatorade", "gatorade"},
		{"Google", "google"},
		{"Gucci", "gucci"},
		{"H&M", "h-m"},
		{"Heineken", "heineken"},
		{"Hewlett-Packard", "hewlett-packard"},
		{"Honda", "honda"},
		{"Huawei", "huawei"},
		{"IBM", "ibm"},
		{"Intel", "intel"},
		{"Jaguar", "jaguar"},
		{"Johnson & Johnson", "johnson-johnson"},
		{"JPMorgan Chase", "jpmorgan-chase"},
		{"KFC", "kfc"},
		{"Kia", "kia"},
		{"L'Oréal", "l-oreal"},
		{"Levi's", "levi-s"},
		{"LG", "lg"},
		{"Lexus", "lexus"},
		{"Louis Vuitton", "louis-vuitton"},
		{"Mastercard", "mastercard"},
		{"McDonald's", "mcdonald-s"},
		{"Mercedes-Benz", "mercedes-benz"},
		{"Microsoft", "microsoft"},
		{"Mitsubishi", "mitsubishi"},
		{"Nestlé", "nestle"},
		{"Nike", "nike"},
		{"Nissan", "nissan"},
		{"Olay", "olay"},
		{"Pepsi", "pepsi"},
		{"Pfizer", "pfizer"},
		{"Philips", "philips"},
		{"Porsche", "porsche"},
		{"Prada", "prada"},
		{"Procter & Gamble", "procter-gamble"},
		{"Puma", "puma"},
		{"Red Bull", "red-bull"},
		{"Rolex", "rolex"},
		{"Samsung", "samsung"},
		{"Sony", "sony"},
		{"Starbucks", "starbucks"},
		{"Subway", "subway"},
		{"T-Mobile", "t-mobile"},
		{"Tesla", "tesla"},
		{"The North Face", "the-north-face"},
		{"Toyota", "toyota"},
		{"Visa", "visa"},
		{"Vodafone", "vodafone"},
		{"Volvo", "volvo"},
		{"Walmart", "walmart"},
		{"Whirlpool", "whirlpool"},
		{"Yahoo", "yahoo"},
		{"Yves Saint Laurent", "yves-saint-laurent"},
		{"Zara", "zara"},
		{"Abercrombie & Fitch", "abercrombie-fitch"},
		{"Acer", "acer"},
		{"Allianz", "allianz"},
		{"American Express", "american-express"},
		{"Aston Martin", "aston-martin"},
		{"Ben & Jerry's", "ben-jerry-s"},
		{"Bulgari", "bulgari"},
		{"Burberry", "burberry"},
		{"Cadillac", "cadillac"},
		{"Carlsberg", "carlsberg"},
		{"Chivas Regal", "chivas-regal"},
		{"Coach", "coach"},
		{"Corona", "corona"},
		{"Crocs", "crocs"},
		{"Danone", "danone"},
		{"Diesel", "diesel"},
		{"Dior", "dior"},
		{"Dell", "dell"},
		{"Duracell", "duracell"},
		{"Ferrari", "ferrari"},
		{"Fendi", "fendi"},
		{"Fila", "fila"},
		{"Fitbit", "fitbit"},
		{"GAP", "gap"},
		{"Geox", "geox"},
		{"Glenfiddich", "glenfiddich"},
		{"Havaianas", "havaianas"},
		{"Hugo Boss", "hugo-boss"},
		{"Jack Daniel's", "jack-daniel-s"},
		{"Jameson", "jameson"},
		{"Kering", "kering"},
		{"Kiehl's", "kiehl-s"},
		{"Lacoste", "lacoste"},
		{"Lancôme", "lancome"},
		{"Land Rover", "land-rover"},
		{"Lindt", "lindt"},
		{"Lululemon", "lululemon"},
		{"Maserati", "maserati"},
		{"Maui Jim", "maui-jim"},
		{"Michael Kors", "michael-kors"},
		{"Montblanc", "montblanc"},
		{"Moët & Chandon", "moet-chandon"},
		{"Nivea", "nivea"},
		{"Oreo", "oreo"},
		{"Patek Philippe", "patek-philippe"},
		{"Peugeot", "peugeot"},
		{"Ray-Ban", "ray-ban"},
		{"Reebok", "reebok"},
		{"Ralph Lauren", "ralph-lauren"},
		{"Remy Martin", "remy-martin"},
		{"Ritz-Carlton", "ritz-carlton"},
		{"Ryanair", "ryanair"},
		{"Saint Laurent", "saint-laurent"},
		{"Swarovski", "swarovski"},
		{"Tag Heuer", "tag-heuer"},
		{"Tiffany & Co.", "tiffany-co"},
		{"Tommy Hilfiger", "tommy-hilfiger"},
		{"Tory Burch", "tory-burch"},
		{"Under Armour", "under-armour"},
		{"Vans", "vans"},
		{"Versace", "versace"},
		{"Vichy", "vichy"},
		{"Vivienne Westwood", "vivienne-westwood"},
		{"Wella", "wella"},
		{"Yamaha", "yamaha"},
		{"Zegna", "zegna"},
		{"Nokia", "nokia"},
		{"Motorola", "motorola"},
		{"HTC", "htc"},
		{"Cooper's", "cooper-s"},
		{"Hunter", "hunter"},
		{"Nasir Group", "nasir-group"},
	}

	for _, brand := range brands {
		brandEntity := &entities.Brand{
			Name: brand.Name,
			Slug: brand.Slug,
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
