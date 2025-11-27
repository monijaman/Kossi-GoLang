package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Product struct {
	NameEN string `json:"name_en"`
}

func main() {
	// Read motorbikes.json
	jsonPath := "init-db/seeders/motorbikes.json"
	jsonData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("❌ Error reading %s: %v\n", jsonPath, err)
		os.Exit(1)
	}

	var products []Product
	if err := json.Unmarshal(jsonData, &products); err != nil {
		fmt.Printf("❌ Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Create a map of product names for quick lookup
	productMap := make(map[string]bool)
	for _, p := range products {
		productMap[strings.ToLower(p.NameEN)] = true
	}

	fmt.Printf("📦 Loaded %d products from JSON\n\n", len(products))

	// Find all seeder files
	seederDir := "internal/infrastructure/database/seeders"
	files, err := filepath.Glob(filepath.Join(seederDir, "*_seeder.go"))
	if err != nil {
		fmt.Printf("❌ Error finding seeder files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📄 Found %d seeder files\n\n", len(files))

	// Pattern to match product search in seeder files
	// Looking for: db.Where("name = ?", "Product Name").First(&product)
	productSearchPattern := regexp.MustCompile(`db\.Where\("name\s*=\s*\?",\s*"([^"]+)"\)`)

	missingProducts := make(map[string][]string) // product name -> list of seeder files
	foundCount := 0
	missingCount := 0

	// Check each seeder file
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		matches := productSearchPattern.FindAllStringSubmatch(string(content), -1)
		for _, match := range matches {
			if len(match) > 1 {
				productName := match[1]
				productNameLower := strings.ToLower(productName)

				if !productMap[productNameLower] {
					missingProducts[productName] = append(missingProducts[productName], filepath.Base(file))
					missingCount++
				} else {
					foundCount++
				}
			}
		}
	}

	fmt.Printf("✅ Valid products: %d\n", foundCount)
	fmt.Printf("❌ Missing products: %d\n\n", missingCount)

	if len(missingProducts) > 0 {
		fmt.Println("❌ Products referenced in seeders but NOT found in motorbikes.json:")
		fmt.Println(strings.Repeat("=", 80))

		for productName, seederFiles := range missingProducts {
			fmt.Printf("\n📦 Product: \"%s\"\n", productName)
			fmt.Printf("   Used in %d seeder file(s):\n", len(seederFiles))
			for _, file := range seederFiles {
				fmt.Printf("   - %s\n", file)
			}
		}

		// Now extract the seeder function names to comment out
		fmt.Println("\n\n" + strings.Repeat("=", 80))
		fmt.Println("🔧 SEEDERS TO COMMENT OUT IN seeder.go:")
		fmt.Println(strings.Repeat("=", 80))

		seederFunctionPattern := regexp.MustCompile(`func \(s \*(\w+)\) Seed`)

		for productName, seederFiles := range missingProducts {
			for _, seederFile := range seederFiles {
				fullPath := filepath.Join(seederDir, seederFile)
				content, err := ioutil.ReadFile(fullPath)
				if err != nil {
					continue
				}

				matches := seederFunctionPattern.FindStringSubmatch(string(content))
				if len(matches) > 1 {
					structName := matches[1]
					fmt.Printf("// manager.AddSeeder(New%s()) // Product not found: %s\n", structName, productName)
				}
			}
		}
	}

	fmt.Println("\n\n✅ Validation complete!")
}
