package main

import (
	"context"
	"fmt"
	"kossti/internal/infrastructure/database/models"
	postgresRepo "kossti/internal/interface/repository/postgres"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dbURL := "postgres://root:root@localhost:5428/kossti?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create repository
	productRepo := postgresRepo.NewPostgresProductRepo(db)

	// Get product by ID
	product, err := productRepo.GetByID(context.Background(), 3)
	if err != nil {
		log.Fatal("Failed to get product:", err)
	}

	fmt.Printf("Product from repository:\n")
	fmt.Printf("ID: %d\n", product.ID)
	fmt.Printf("Name: %s\n", product.Name)
	fmt.Printf("Status: %v\n", product.Status)
	fmt.Printf("Priority: %d\n", product.Priority)
	fmt.Printf("CategoryID: %v\n", product.CategoryID)
	fmt.Printf("BrandID: %v\n", product.BrandID)
	fmt.Printf("CreatedAt: %v\n", product.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", product.UpdatedAt)

	// Also check the raw model
	var productModel models.ProductModel
	if err := db.First(&productModel, 3).Error; err != nil {
		log.Fatal("Failed to get product model:", err)
	}

	fmt.Printf("\nProduct model from GORM:\n")
	fmt.Printf("ID: %d\n", productModel.ID)
	fmt.Printf("Name: %s\n", productModel.Name)
	fmt.Printf("Status: %v\n", productModel.Status)
	fmt.Printf("Priority: %d\n", productModel.Priority)
	fmt.Printf("CreatedAt: %v\n", productModel.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", productModel.UpdatedAt)

	// Test ToEntity conversion
	convertedProduct := productModel.ToEntity()
	fmt.Printf("\nConverted product:\n")
	fmt.Printf("ID: %d\n", convertedProduct.ID)
	fmt.Printf("Name: %s\n", convertedProduct.Name)
	fmt.Printf("Status: %v\n", convertedProduct.Status)
	fmt.Printf("Priority: %d\n", convertedProduct.Priority)
}
