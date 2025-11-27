package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var batch40Products = []string{
	"TVS Apache RTR 160 4V Fi",
	"TVS XL 100 Comfort",
	"TVS Apache RTR 160 2V ABS",
	"Honda CB Shine SP",
	"Honda X-Blade 160 ABS",
	"TVS Apache RTR 160 2V Single Disc",
	"Yamaha FZS V3 ABS Vintage Edition",
	"Honda Hornet 2.0",
	"New TVS Apache RTR 160 4V",
	"Yamaha FZS V3 ABS BS4",
	"Honda Dream 110",
	"Yamaha FZS V4",
	"Suzuki Gixxer Monotone",
	"Honda SP160 (Single Disc)",
	"Royal Enfield Bullet 350",
	"KTM RC 125 2022",
	"Yamaha FZS V3 BS6",
	"Yamaha FZ-X",
	"TVS Apache RTR 160 2V Refresh (X-Connect)",
	"Yamaha Saluto 125",
	"Honda CB150R Streetfire",
	"TVS Radeon",
	"Royal Enfield Meteor 350 (Supernova)",
	"TVS Max 125",
	"Honda Shine 100",
	"TVS Stryker 125",
	"TVS Wego",
	"Royal Enfield Hunter 350 (Rebel Blue/Rebel Red/Black)",
	"Yamaha FZ25",
	"Yamaha Ray ZR Street Rally",
	"New TVS Apache RTR 160 4V (Single Channel ABS)",
	"TVS Max Semi Trail 125",
	"Vespa VXL 125 (CBS)",
	"Bajaj Platina 110 H Gear",
	"Suzuki GS150R",
	"Bajaj Pulsar N160 Twin Disc Carburetor",
	"TVS Raider 125",
	"Honda Super Cub C125 ABS",
	"Royal Enfield Meteor 350 (Stellar)",
	"Honda X-Blade 160",
	"TVS XL 100 i-Touch",
	"Honda Livo 110 Disc CBS",
	"Bajaj Discover 110 Disc",
	"Bajaj Platina ES",
	"KTM Duke 125 European Edition",
	"Royal Enfield Classic 350 (Halcyon Green/Black)",
	"KTM Duke 125 Indian",
	"Yamaha R15 V3 Indonesia",
	"Yamaha FZS Fi Deluxe",
	"Vespa VXL 125",
	"Suzuki Gixxer Carb Disc",
	"Bajaj Pulsar F250 (Red)",
	"Royal Enfield Classic 350 Dark (Gun Metal Grey/Stealth Black)",
	"TVS Rockz 125",
	"Honda Vario 160 ABS",
	"R15 V3 Indian Version Dual ABS",
	"Yamaha Ray ZR 125 Fi",
	"Bajaj Platina 110 ES",
	"Pulsar 150 Twin Disc ABS",
	"Vespa VXL 150 (Yellow)",
	"Honda X-Blade 160 Fi ABS",
	"TVS Metro Plus",
	"Royal Enfield Classic 350 (Chrome Red/Bronze)",
}

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// First, delete translations
	fmt.Println("Deleting product review translations...")
	deleteTransQuery := `
		DELETE FROM product_review_translations 
		WHERE product_review_id IN (
			SELECT pr.id FROM product_reviews pr
			INNER JOIN products p ON p.id = pr.product_id
			WHERE p.name = $1 AND pr.user_id = 1
		)`

	transCount := 0
	for _, productName := range batch40Products {
		result, err := db.Exec(deleteTransQuery, productName)
		if err != nil {
			log.Printf("Error deleting translations for %s: %v", productName, err)
			continue
		}
		rows, _ := result.RowsAffected()
		if rows > 0 {
			transCount++
			fmt.Printf("   ✅ Deleted translation for %s\n", productName)
		}
	}

	// Then delete reviews
	fmt.Println("\nDeleting product reviews...")
	deleteReviewQuery := `
		DELETE FROM product_reviews 
		WHERE product_id IN (
			SELECT id FROM products WHERE name = $1
		) AND user_id = 1`

	reviewCount := 0
	for _, productName := range batch40Products {
		result, err := db.Exec(deleteReviewQuery, productName)
		if err != nil {
			log.Printf("Error deleting review for %s: %v", productName, err)
			continue
		}
		rows, _ := result.RowsAffected()
		if rows > 0 {
			reviewCount++
			fmt.Printf("   ✅ Deleted review for %s\n", productName)
		}
	}

	fmt.Printf("\n%s\n", "=================================================================")
	fmt.Printf("✅ Deleted %d translations and %d reviews\n", transCount, reviewCount)
	fmt.Println("=================================================================")
	fmt.Println("\nYou can now reseed with detailed content:")
	fmt.Println("  go run ./cmd/migrate/main.go -seed")
}
