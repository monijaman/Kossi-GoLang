package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type JBProduct struct {
	NameEN string `json:"name_en"`
}

func main() {
	// Load environment variables (prefer .env.production if present)
	if _, err := os.Stat(".env.production"); err == nil {
		_ = godotenv.Load(".env.production")
	}
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Build from components
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbSSLMode := os.Getenv("DB_SSLMODE")
		if dbHost == "" || dbUser == "" || dbName == "" {
			log.Fatal("DATABASE_URL not set and DB_HOST/DB_USER/DB_NAME are missing")
		}
		if dbPort == "" {
			dbPort = "5432"
		}
		if dbSSLMode == "" {
			dbSSLMode = "disable"
		}
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
	}

	fmt.Println("Using DB DSN (masked):", mask(dsn))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	// Load motorbikes.json
	jsonPath := "init-db/seeders/motorbikes.json"
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("failed to read %s: %v", jsonPath, err)
	}

	var products []JBProduct
	if err := json.Unmarshal(data, &products); err != nil {
		log.Fatalf("failed to parse json: %v", err)
	}

	jsonNames := make(map[string]struct{})
	for _, p := range products {
		n := strings.ToLower(strings.TrimSpace(p.NameEN))
		if n != "" {
			jsonNames[n] = struct{}{}
		}
	}

	// Query DB for product names that have at least one review
	rows, err := db.Query(`SELECT p.name FROM products p JOIN product_reviews pr ON pr.product_id = p.id GROUP BY p.id, p.name`)
	if err != nil {
		log.Fatalf("db query failed: %v", err)
	}
	defer rows.Close()

	reviewed := make(map[string]struct{})
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalf("scan failed: %v", err)
		}
		reviewed[strings.ToLower(strings.TrimSpace(name))] = struct{}{}
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("rows err: %v", err)
	}

	// Now list JSON product names missing reviews
	missing := []string{}
	for name := range jsonNames {
		if _, ok := reviewed[name]; !ok {
			missing = append(missing, name)
		}
	}

	fmt.Printf("\nTotal products in JSON: %d\n", len(jsonNames))
	fmt.Printf("Total products with reviews in DB: %d\n", len(reviewed))
	fmt.Printf("Products missing reviews: %d\n\n", len(missing))

	for _, n := range missing {
		fmt.Println(n)
	}
}

func mask(s string) string {
	if strings.HasPrefix(s, "postgres://") || strings.HasPrefix(s, "postgresql://") {
		// mask password
		parts := strings.Split(s, "@")
		if len(parts) >= 2 {
			return parts[0] + "@***"
		}
	}
	return "***masked***"
}
