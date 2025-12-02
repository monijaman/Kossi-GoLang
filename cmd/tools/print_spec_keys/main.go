package main

import (
	"flag"
	"fmt"
	"log"

	"kossti/internal/infrastructure/database"

	"gorm.io/gorm"
)

func main() {
	listAll := flag.Bool("all", true, "list all specification_keys rows")
	listDup := flag.Bool("dups", true, "list duplicated specification_key values")
	flag.Parse()

	cfg := database.NewDatabaseConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if *listAll {
		if err := printAllKeys(db); err != nil {
			log.Fatalf("error printing all keys: %v", err)
		}
	}

	if *listDup {
		if err := printDuplicates(db); err != nil {
			log.Fatalf("error printing duplicates: %v", err)
		}
	}
}

func printAllKeys(db *gorm.DB) error {
	type KeyRow struct {
		ID               uint
		SpecificationKey string
	}
	var rows []KeyRow
	if tx := db.Table("specification_keys").Select("id, specification_key").Order("id").Scan(&rows); tx.Error != nil {
		return tx.Error
	}
	fmt.Println("--- specification_keys rows ---")
	for _, r := range rows {
		fmt.Printf("ID=%d  Key='%s'\n", r.ID, r.SpecificationKey)
	}
	fmt.Printf("Total keys: %d\n", len(rows))
	return nil
}

func printDuplicates(db *gorm.DB) error {
	type Dup struct {
		SpecificationKey string
		Count            int
	}
	var dups []Dup
	q := `SELECT specification_key, COUNT(*) as count
          FROM specification_keys
          GROUP BY specification_key
          HAVING COUNT(*) > 1
          ORDER BY count DESC;`
	if tx := db.Raw(q).Scan(&dups); tx.Error != nil {
		return tx.Error
	}
	fmt.Println("--- duplicate specification_key values ---")
	if len(dups) == 0 {
		fmt.Println("No duplicate specification_key values found.")
		return nil
	}
	for _, d := range dups {
		fmt.Printf("Key='%s'  Count=%d\n", d.SpecificationKey, d.Count)
	}
	return nil
}
