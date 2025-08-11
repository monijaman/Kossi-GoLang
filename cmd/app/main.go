// main.go - Entry Point for the Auth Microservice
//
// This file is the main entry point for the authentication microservice.
// It is responsible for:
//   - Loading environment variables and configuration
//   - Connecting to the database and running migrations
//   - Initializing repositories and HTTP handlers
//   - Registering all HTTP routes (auth, user, health, etc.)
//   - Starting and gracefully shutting down the HTTP server
//
// This file wires together the infrastructure, domain, and interface layers according to Clean Architecture principles.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"kossti/internal/infrastructure/database"
	database_seeders "kossti/internal/infrastructure/database/seeders"
	handlerauth "kossti/internal/interface/handler/auth"
	handlerbrand "kossti/internal/interface/handler/brand"
	handlercategory "kossti/internal/interface/handler/category"
	handlerfeedback "kossti/internal/interface/handler/feedback"
	handlerformgenerator "kossti/internal/interface/handler/formgenerator"
	handlerproduct "kossti/internal/interface/handler/product"
	handlerproductreview "kossti/internal/interface/handler/productreview"
	handlerspecification "kossti/internal/interface/handler/specification"
	handleruser "kossti/internal/interface/handler/user"
	pgRepo "kossti/internal/interface/repository/postgres"
)

// findAvailablePort finds an available port starting from the given port
func findAvailablePort(startPort int) int {
	for port := startPort; port < startPort+100; port++ {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return port
		}
	}
	return startPort // fallback to original port
}

// corsMiddleware adds CORS headers to allow cross-origin requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// killProcessOnPort attempts to find and kill any process using the specified port
func killProcessOnPort(port int) {
	// On Windows, you can use netstat + taskkill, but for now we'll just log
	fmt.Printf("⚠️  Port %d appears to be in use. You may need to manually kill the process.\n", port)
	fmt.Printf("💡 Run this command to find the process: netstat -ano | findstr :%d\n", port)
	fmt.Printf("💡 Then kill it with: taskkill /PID <PID> /F\n")
}

func main() {
	// Load the appropriate .env file based on GO_ENV
	env := os.Getenv("GO_ENV")
	if env == "production" {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Println("No .env.production file found or error loading .env.production file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("No .env file found or error loading .env file")
		}
	}

	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	fmt.Println("Auth microservice starting...")
	fmt.Println("Database URL:", dbURL)
	fmt.Println("JWT Secret:", jwtSecret)
	fmt.Println("Kafka Broker:", kafkaBroker)

	// Validate required environment variables
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	fmt.Println("Attempting to connect to the database and migrate schema...")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("GORM database connection failed: %v", err)
	}

	// Test database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	fmt.Println("Database connection successful!")

	// Initialize migration manager and run all migrations
	migrationManager := database.NewMigrationManager(db)

	// Run complete database setup
	fmt.Println("Running database migrations...")
	if err := migrationManager.Setup(); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// --- Seed data only if not already seeded ---
	var userCount int64
	err = db.Table("users").Count(&userCount).Error
	if err != nil {
		log.Fatalf("Failed to check users table for seeding: %v", err)
	}
	if userCount == 0 {
		fmt.Println("No users found, running initial data seeder...")
		if err := database_seeders.SetupAllSeeders(db).RunAll(); err != nil {
			log.Fatalf("Database seeding failed: %v", err)
		}
		fmt.Println("Database seeding complete!")
	} else {
		fmt.Println("Seed data already exists, skipping seeder.")
	}

	// Create repository instance with GORM DB
	userRepo := pgRepo.NewPostgresUserRepo(db)
	refreshTokenRepo := pgRepo.NewPostgresRefreshTokenRepo(db)
	productRepo := pgRepo.NewPostgresProductRepo(db)
	imageRepo := pgRepo.NewPostgresImageRepo(db)
	categoryRepo := pgRepo.NewPostgresCategoryRepo(db)
	brandRepo := pgRepo.NewPostgresBrandRepo(db)
	specificationRepo := pgRepo.NewPostgresSpecificationRepo(db)
	specificationKeyRepo := pgRepo.NewPostgresSpecificationKeyRepo(db)
	productReviewRepo := pgRepo.NewProductReviewRepository(db)
	formGeneratorRepo := pgRepo.NewFormGeneratorRepository(db)
	feedbackRepo := pgRepo.NewFeedbackRepository(db)

	fmt.Println("Database migration complete! Setting up HTTP routes...")

	// Create a new HTTP mux for better route handling
	mux := http.NewServeMux()

	// Register grouped routes
	handlerauth.RegisterAuthRoutes(mux, userRepo, refreshTokenRepo)
	handleruser.RegisterUserRoutes(mux, userRepo)
	handlerproduct.RegisterProductRoutes(mux, productRepo, imageRepo, categoryRepo, brandRepo)
	handlercategory.RegisterCategoryRoutes(mux, categoryRepo)
	handlerbrand.RegisterBrandRoutes(mux, brandRepo)
	handlerspecification.RegisterSpecificationRoutes(mux, specificationRepo, specificationKeyRepo)
	handlerproductreview.RegisterProductReviewRoutes(mux, productReviewRepo)
	handlerformgenerator.RegisterRoutes(mux, formGeneratorRepo)
	handlerfeedback.RegisterRoutes(mux, feedbackRepo)

	// Add health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "healthy",
			"service": "kossti",
		})
	})

	fmt.Println("\nAPI Endpoints:")
	// Authentication endpoints
	fmt.Println("POST http://localhost:8080/register")
	fmt.Println("POST http://localhost:8080/login")
	fmt.Println("POST http://localhost:8080/v1/refresh-token")
	fmt.Println("POST http://localhost:8080/v1/registration")
	fmt.Println("POST http://localhost:8080/v1/login")
	fmt.Println("POST http://localhost:8080/v1/forgot-password")
	fmt.Println("POST http://localhost:8080/v1/reset-password")
	fmt.Println("POST http://localhost:8080/v1/logout")
	fmt.Println("GET  http://localhost:8080/v1/check-token")
	fmt.Println("GET  http://localhost:8080/v1/search_users")
	fmt.Println("GET  http://localhost:8080/v1/profile")
	fmt.Println("POST http://localhost:8080/v1/profile")
	fmt.Println("GET  http://localhost:8080/v1/checkrole")
	// User endpoints
	fmt.Println("GET  http://localhost:8080/api/users")
	fmt.Println("GET  http://localhost:8080/api/users/all")
	fmt.Println("GET  http://localhost:8080/api/users/search")
	fmt.Println("GET  http://localhost:8080/api/users/count")
	fmt.Println("GET  http://localhost:8080/api/users/{id}")
	// Product endpoints
	fmt.Println("GET  http://localhost:8080/api/products")
	fmt.Println("POST http://localhost:8080/api/products")
	fmt.Println("GET  http://localhost:8080/api/products/{id}")
	fmt.Println("PATCH http://localhost:8080/api/products/{id}")
	fmt.Println("GET  http://localhost:8080/api/products-by-slug/{slug}")
	fmt.Println("GET  http://localhost:8080/api/popular-products")
	fmt.Println("POST http://localhost:8080/api/products/{id}/increment-views")
	// Category endpoints
	fmt.Println("GET  http://localhost:8080/api/categories")
	fmt.Println("POST http://localhost:8080/api/categories")
	fmt.Println("GET  http://localhost:8080/api/categories/{id}")
	fmt.Println("PUT  http://localhost:8080/api/categories/{id}")
	fmt.Println("DELETE http://localhost:8080/api/categories/{id}")
	fmt.Println("GET  http://localhost:8080/api/wide-categories")
	fmt.Println("POST http://localhost:8080/api/category-translation")
	fmt.Println("GET  http://localhost:8080/api/category-translation/{id}")
	fmt.Println("POST http://localhost:8080/api/category-brands")
	fmt.Println("GET  http://localhost:8080/api/category-brands")
	fmt.Println("POST http://localhost:8080/api/category-status/{id}")
	// Brand endpoints
	fmt.Println("GET  http://localhost:8080/api/brands")
	fmt.Println("POST http://localhost:8080/api/brands")
	fmt.Println("GET  http://localhost:8080/api/brands/{id}")
	fmt.Println("PUT  http://localhost:8080/api/brands/{id}")
	fmt.Println("DELETE http://localhost:8080/api/brands/{id}")
	fmt.Println("GET  http://localhost:8080/api/specifications/{id}")
	fmt.Println("PUT  http://localhost:8080/api/specifications/{id}")
	fmt.Println("DELETE http://localhost:8080/api/specifications/{id}")
	fmt.Println("GET  http://localhost:8080/api/specificationsearch")
	fmt.Println("POST http://localhost:8080/api/spec_translation")
	fmt.Println("GET  http://localhost:8080/api/spec_translation/{id}")
	fmt.Println("GET  http://localhost:8080/api/get-public-spec/{id}")
	fmt.Println("GET  http://localhost:8080/api/speckey")
	fmt.Println("POST http://localhost:8080/api/speckey")
	fmt.Println("GET  http://localhost:8080/api/speckey/{id}")
	fmt.Println("POST http://localhost:8080/api/specremove/{id}")
	fmt.Println("GET  http://localhost:8080/api/speckey-translation")
	fmt.Println("POST http://localhost:8080/api/speckey-translation")
	// FormGenerator endpoints
	fmt.Println("POST http://localhost:8080/api/formgenerator")
	fmt.Println("GET  http://localhost:8080/api/formgenerator/{id}")
	fmt.Println("PUT  http://localhost:8080/api/formgenerator/{id}")
	fmt.Println("GET  http://localhost:8080/api/catgory-specs/{categoryId}")
	fmt.Println("")
	fmt.Println("📋 Feedback Endpoints:")
	fmt.Println("GET  http://localhost:8080/api/feedbacks")
	fmt.Println("POST http://localhost:8080/api/feedbacks")
	fmt.Println("GET  http://localhost:8080/api/feedbacks/{id}")
	fmt.Println("PUT  http://localhost:8080/api/feedbacks/{id}")
	fmt.Println("DEL  http://localhost:8080/api/feedbacks/{id}")
	fmt.Println("GET  http://localhost:8080/api/feedbacks/{id}/translations")
	fmt.Println("GET  http://localhost:8080/health")

	// Determine which port to use
	preferredPort := 8080
	availablePort := findAvailablePort(preferredPort)

	if availablePort != preferredPort {
		fmt.Printf("⚠️   Port %d is  in use, using port %d instead\n", preferredPort, availablePort)
	}

	serverAddr := fmt.Sprintf(":%d", availablePort)

	// Create HTTP server with CORS middleware
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for interrupt signal to terminate gracefully
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		fmt.Printf("\n🚀 Server is running on http://localhost:%d\n", availablePort)
		fmt.Println("Press Ctrl+C to stop the server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			if err.Error() == "listen tcp "+serverAddr+": bind: Only one usage of each socket address (protocol/network address/port) is normally permitted." {
				killProcessOnPort(availablePort)
			}
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	fmt.Println("\n🛑 Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	fmt.Println("✅ Server stopped gracefully")
}
