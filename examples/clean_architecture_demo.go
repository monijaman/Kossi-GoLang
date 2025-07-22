// Package main demonstrates how Clean Architecture solves the GORM coupling problem
package main

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"
	"fmt"
)

func main() {
	fmt.Println("🎯 Clean Architecture GORM Decoupling Demo")
	fmt.Println("==========================================")

	// 1. Domain Entity (Framework Independent)
	fmt.Println("\n1. Creating Domain Entity (No GORM dependency):")
	user := &entities.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	fmt.Printf("   Domain User: %+v\n", user)

	// 2. Infrastructure Model (GORM Specific)
	fmt.Println("\n2. Converting to Infrastructure Model (GORM specific):")
	userModel := &models.UserModel{}
	userModel.FromEntity(user)
	fmt.Printf("   GORM Model: %+v\n", userModel)

	// 3. Back to Domain Entity
	fmt.Println("\n3. Converting back to Domain Entity:")
	convertedUser := userModel.ToEntity()
	fmt.Printf("   Converted User: %+v\n", convertedUser)

	// 4. Demonstrate Brand/Category example
	fmt.Println("\n4. Brand-Category Relationship Example:")
	brandCategory := &entities.BrandCategory{
		ID:         1,
		BrandID:    10,
		CategoryID: 20,
	}
	
	brandCategoryModel := &models.BrandCategoryModel{}
	brandCategoryModel.FromEntity(brandCategory)
	fmt.Printf("   Original Entity: %+v\n", brandCategory)
	fmt.Printf("   GORM Model: %+v\n", brandCategoryModel)
	fmt.Printf("   Back to Entity: %+v\n", brandCategoryModel.ToEntity())

	fmt.Println("\n✅ SUCCESS: Business Logic is completely decoupled from GORM!")
	fmt.Println("   - Domain entities have zero GORM dependencies")
	fmt.Println("   - Infrastructure models handle GORM concerns")
	fmt.Println("   - Easy to switch ORMs by changing only infrastructure layer")
	fmt.Println("   - Business logic remains unchanged when switching databases")
}
