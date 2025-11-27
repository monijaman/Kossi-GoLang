package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get database connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	productID := 499

	// Query 1: Product basic info and translations
	fmt.Println("=== Product Info and Translations ===")
	type ProductTranslation struct {
		ProductID             int
		ProductName           string
		ProductSlug           string
		Locale                string
		TranslatedName        string
		TranslatedDescription string
	}

	var productTranslations []ProductTranslation
	db.Raw(`
		SELECT 
			p.id AS product_id,
			p.name AS product_name,
			p.slug AS product_slug,
			pt.locale,
			pt.name AS translated_name,
			pt.description AS translated_description
		FROM 
			products p
		LEFT JOIN 
			product_translations pt ON p.id = pt.product_id
		WHERE 
			p.id = ?
		ORDER BY 
			pt.locale
	`, productID).Scan(&productTranslations)

	for _, pt := range productTranslations {
		fmt.Printf("Product: %s (ID: %d, Slug: %s)\n", pt.ProductName, pt.ProductID, pt.ProductSlug)
		if pt.Locale != "" {
			fmt.Printf("  Locale: %s\n", pt.Locale)
			fmt.Printf("  Translated Name: %s\n", pt.TranslatedName)
			fmt.Printf("  Description: %s\n", pt.TranslatedDescription)
		}
		fmt.Println()
	}

	// Query 2: Specifications with translations
	fmt.Println("\n=== Specifications with Translations ===")
	type SpecTranslation struct {
		ProductID           int
		ProductName         string
		SpecID              int
		SpecificationKey    string
		SpecValueEnglish    string
		TranslationLocale   string
		SpecValueTranslated string
	}

	var specTranslations []SpecTranslation
	db.Raw(`
		SELECT 
			p.id AS product_id,
			p.name AS product_name,
			s.id AS spec_id,
			sk.specification_key,
			s.value AS spec_value_english,
			st.locale AS translation_locale,
			st.value AS spec_value_translated
		FROM 
			products p
		LEFT JOIN 
			specifications s ON p.id = s.product_id
		LEFT JOIN 
			specification_keys sk ON s.specification_key_id = sk.id
		LEFT JOIN 
			specification_translations st ON s.id = st.specification_id
		WHERE 
			p.id = ?
		ORDER BY 
			sk.specification_key, st.locale
	`, productID).Scan(&specTranslations)

	currentKey := ""
	for _, st := range specTranslations {
		if st.SpecificationKey != currentKey {
			currentKey = st.SpecificationKey
			fmt.Printf("\n%s: %s\n", st.SpecificationKey, st.SpecValueEnglish)
		}
		if st.TranslationLocale != "" {
			fmt.Printf("  [%s] %s\n", st.TranslationLocale, st.SpecValueTranslated)
		}
	}

	// Query 3: Count summary
	fmt.Println("\n\n=== Translation Summary ===")
	type TranslationSummary struct {
		ProductID           int
		ProductName         string
		TotalSpecifications int
		BengaliTranslations int
		HindiTranslations   int
	}

	var summary TranslationSummary
	db.Raw(`
		SELECT 
			p.id,
			p.name,
			COUNT(DISTINCT s.id) AS total_specifications,
			COUNT(DISTINCT CASE WHEN st.locale = 'bn' THEN st.id END) AS bengali_translations,
			COUNT(DISTINCT CASE WHEN st.locale = 'hi' THEN st.id END) AS hindi_translations
		FROM 
			products p
		LEFT JOIN 
			specifications s ON p.id = s.product_id
		LEFT JOIN 
			specification_translations st ON s.id = st.specification_id
		WHERE 
			p.id = ?
		GROUP BY 
			p.id, p.name
	`, productID).Scan(&summary)

	fmt.Printf("Product: %s (ID: %d)\n", summary.ProductName, summary.ProductID)
	fmt.Printf("Total Specifications: %d\n", summary.TotalSpecifications)
	fmt.Printf("Bengali Translations: %d\n", summary.BengaliTranslations)
	fmt.Printf("Hindi Translations: %d\n", summary.HindiTranslations)
}
