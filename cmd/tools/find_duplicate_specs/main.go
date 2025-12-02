package main

import (
	"flag"
	"fmt"
	"log"

	"kossti/internal/infrastructure/database"

	"gorm.io/gorm"
)

func main() {
	listRows := flag.Bool("list", false, "also list individual duplicate row IDs")
	flag.Parse()

	cfg := database.NewDatabaseConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := findDuplicates(db, *listRows); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func findDuplicates(db *gorm.DB, listRows bool) error {
	type Dup struct {
		ProductID          uint
		SpecificationKeyID uint
		Value              string
		Count              int
	}

	var dups []Dup

	// Find groups with more than 1 occurrence
	q := `SELECT product_id, specification_key_id, value, COUNT(*) as count
          FROM specifications
          GROUP BY product_id, specification_key_id, value
          HAVING COUNT(*) > 1
          ORDER BY count DESC;`

	if tx := db.Raw(q).Scan(&dups); tx.Error != nil {
		return tx.Error
	}

	if len(dups) == 0 {
		fmt.Println("No duplicate specification groups found.")
		return nil
	}

	fmt.Printf("Found %d duplicate groups:\n", len(dups))
	for _, d := range dups {
		fmt.Printf("ProductID=%d, SpecKeyID=%d, Value='%s' -> Count=%d\n", d.ProductID, d.SpecificationKeyID, d.Value, d.Count)
		if listRows {
			type Row struct {
				ID uint
			}
			var rows []Row
			rr := db.Raw(`SELECT id FROM specifications WHERE product_id = ? AND specification_key_id = ? AND value = ? ORDER BY id`, d.ProductID, d.SpecificationKeyID, d.Value).Scan(&rows)
			if rr.Error != nil {
				return rr.Error
			}
			fmt.Printf("  Row IDs: ")
			for i, r := range rows {
				if i > 0 {
					fmt.Printf(", %d", r.ID)
				} else {
					fmt.Printf("%d", r.ID)
				}
			}
			fmt.Println()
		}
	}

	return nil
}
