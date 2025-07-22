// Package config shows how to wire dependencies using Clean Architecture principles.
// This demonstrates how easy it is to switch between different implementations.
package config

import (
	"kossti/internal/infrastructure/database"
	"kossti/internal/interface/repository/postgres"

	"gorm.io/gorm"
)

// ServiceContainer holds all dependencies
type ServiceContainer struct {
	// Database services
	MigrationManager *database.MigrationManager
	
	// Repositories (ready to use)
	UserRepository *postgres.PostgresUserRepo
}

// NewServiceContainerWithGORM creates services using GORM
func NewServiceContainerWithGORM(db *gorm.DB) *ServiceContainer {
	// Infrastructure layer - GORM implementations
	migrationManager := database.NewMigrationManager(db)
	userRepo := postgres.NewPostgresUserRepo(db)

	return &ServiceContainer{
		MigrationManager: migrationManager,
		UserRepository:   userRepo,
	}
}

// Example usage in your application:
//
// func main() {
//     db := connectToDatabase()
//     container := config.NewServiceContainerWithGORM(db)
//     
//     // Run migrations
//     if err := container.MigrationManager.Setup(); err != nil {
//         log.Fatal("Migration failed:", err)
//     }
//     
//     // Use repository for database operations
//     users, err := container.UserRepository.List(ctx, 10, 0)
//     if err != nil {
//         log.Fatal("Failed to list users:", err)
//     }
// }

// Example of how you could switch to MongoDB:
// func NewServiceContainerWithMongoDB(db *mongo.Database) *ServiceContainer {
//     // Infrastructure layer - MongoDB implementations
//     userRepo := mongo.NewMongoUserRepo(db)
//     migrationManager := mongo.NewMongoMigrationManager(db)
//
//     return &ServiceContainer{
//         MigrationManager: migrationManager,
//         UserRepository:   userRepo,
//     }
// }

// The key insight: Your business logic and domain entities
// are completely independent of the database implementation!
// 
// Benefits of this approach:
// 1. Easy to test - mock the repository interfaces
// 2. Easy to switch databases - just change the repository implementation
// 3. Business logic is database-agnostic
// 4. Clean separation of concerns
// 5. Follows SOLID principles
