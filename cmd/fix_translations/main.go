package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgresql://postgres:wqxrKmVOvJeQORpcsGhWOGiRuQdJDLBU@switchyard.proxy.rlwy.net:58014/railway?sslmode=require"
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("connect:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("ping:", err)
	}

	// spec_id -> Bengali value for product 341
	translations := map[int]string{
		63933: "কমিউনিটি ব্যাংক বাংলাদেশ পিএলসি",
		11128: "২০১৯",
		63934: "পুলিশ প্লাজা কনকর্ড (লেভেল ১০, টাওয়ার ২), প্লট ২, রোড ১৪৪, গুলশান ১, ঢাকা ১২১২, বাংলাদেশ",
		11125: "CBBDBDHD",
		63935: "--",
		63936: "বাণিজ্যিক ব্যাংক লাইসেন্স",
		63937: "বেসরকারি",
		63938: "আব্দুল আওয়াল মিন্টু",
		63939: "মো. শফিকুল ইসলাম",
		63940: "৫০",
		63941: "১০০",
		63942: "২০০",
		63943: "টেমেনোস টি২৪",
		63944: "হ্যাঁ",
		63945: "হ্যাঁ",
		63946: "কমিউনিটি ব্যাংক মোবাইল",
		63947: "হ্যাঁ",
		63948: "হ্যাঁ",
		63949: "হ্যাঁ",
		63950: "হ্যাঁ",
		63951: "হ্যাঁ",
		63952: "স্থায়ী আমানত, মাসিক আয় প্রকল্প, সঞ্চয়ী হিসাব",
		63953: "ব্যক্তিগত ঋণ, গৃহ ঋণ, গাড়ি ঋণ",
		63954: "হ্যাঁ",
		63955: "হ্যাঁ",
		63956: "হ্যাঁ",
		63957: "হ্যাঁ",
		63958: "হ্যাঁ",
		63959: "হ্যাঁ",
		63960: "হ্যাঁ",
		63961: "হ্যাঁ",
		63962: "হ্যাঁ",
		63963: "হ্যাঁ",
		63964: "হ্যাঁ",
		63965: "হ্যাঁ",
		63966: "হ্যাঁ",
		63967: "হ্যাঁ",
		63968: "হ্যাঁ",
		63969: "০৯৬৭৭৭১৬৭০৭",
		63970: "contact@communitybankbd.info",
		63971: "facebook.com/communitybankbd",
		63972: "পুলিশ প্লাজা কনকর্ড (লেভেল ১০, টাওয়ার ২), প্লট ২, রোড ১৪৪, গুলশান ১, ঢাকা ১২১২, বাংলাদেশ",
		63973: "২৪/৭",
		63974: "রবিবার থেকে বৃহস্পতিবার",
		63975: "৫০,০০০ টাকা",
		63976: "১,০০,০০০ টাকা",
		63977: "হ্যাঁ",
		63978: "হ্যাঁ",
		11126: "www.communitybankbd.com",
		11127: "বেসরকারি বাণিজ্যিক ব্যাংক",
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("begin tx:", err)
	}

	stmt, err := tx.Prepare(`INSERT INTO specification_translations (specification_id, locale, value) VALUES ($1, 'bn', $2)`)
	if err != nil {
		tx.Rollback()
		log.Fatal("prepare:", err)
	}
	defer stmt.Close()

	count := 0
	for specID, value := range translations {
		if _, err := stmt.Exec(specID, value); err != nil {
			tx.Rollback()
			log.Fatalf("insert spec_id=%d: %v", specID, err)
		}
		count++
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("commit:", err)
	}

	fmt.Printf("✓ Inserted %d Bengali translations for product 341\n", count)
}
