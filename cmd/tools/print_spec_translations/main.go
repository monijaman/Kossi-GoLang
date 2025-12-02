package main

import (
	"flag"
	"fmt"
	"log"

	"kossti/internal/infrastructure/database"

	"gorm.io/gorm"
)

func main() {
	productID := flag.Uint("product", 583, "product id to fetch specification translations for")
	flag.Parse()

	cfg := database.NewDatabaseConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := printSpecTranslations(db, uint(*productID)); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func printSpecTranslations(db *gorm.DB, productID uint) error {
	type Row struct {
		ID                 uint
		SpecificationID    uint
		Locale             string
		Value              string
		ProductID          uint
		SpecValue          string
		SpecificationKeyID uint
	}

	var rows []Row

	tx := db.Table("specification_translations st").
		Select("st.id, st.specification_id, st.locale, st.value, s.product_id, s.value as spec_value, s.specification_key_id").
		Joins("join specifications s on s.id = st.specification_id").
		Where("s.product_id = ?", productID).
		Scan(&rows)

	if tx.Error != nil {
		return tx.Error
	}

	if len(rows) == 0 {
		fmt.Printf("No specification translations found for product id %d\n", productID)
		return nil
	}

	for _, r := range rows {
		fmt.Printf("ID: %d, SpecID: %d, Locale: %s, Value: %s, ProductID: %d, SpecValue: %s, SpecKeyID: %d\n",
			r.ID, r.SpecificationID, r.Locale, r.Value, r.ProductID, r.SpecValue, r.SpecificationKeyID)
	}

	return nil
}
