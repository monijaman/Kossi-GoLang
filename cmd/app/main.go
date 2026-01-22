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
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	database_migrations "kossti/internal/infrastructure/database/migrations"
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

// connectWithRetry attempts to open a GORM DB connection with retries and
// exponential backoff. This helps when the remote database is starting up
// or a proxy transiently closes connections (common on managed providers).
func connectWithRetry(dsn string, maxAttempts int, baseWait time.Duration) (*gorm.DB, error) {
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// verify with a ping
			sqlDB, err := db.DB()
			if err != nil {
				lastErr = err
				_ = sqlDB // just to keep linter happy if nil
			} else {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()
				if err := sqlDB.PingContext(ctx); err == nil {
					return db, nil
				} else {
					lastErr = err
					// Close the underlying DB to avoid leaking connections
					_ = sqlDB.Close()
				}
			}
		} else {
			lastErr = err
		}

		wait := baseWait * time.Duration(1<<uint(attempt-1))
		if wait > 30*time.Second {
			wait = 30 * time.Second
		}
		log.Printf("Database connect attempt %d/%d failed: %v. Retrying in %s...", attempt, maxAttempts, lastErr, wait)
		time.Sleep(wait)
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, lastErr)
}

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
		// Allow specific origin when present so credentials can be used safely in browsers
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			// Ensure caches vary by Origin
			w.Header().Add("Vary", "Origin")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, X-Country-Code")
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
	// Load environment files. Priority:
	// 1) If GO_ENV=production -> load .env.production
	// 2) Else if .env.production exists -> load it and set GO_ENV=production
	// 3) Else load .env (development)
	if os.Getenv("GO_ENV") == "production" {
		if err := godotenv.Load(".env.production"); err != nil {
			log.Println("warning: failed loading .env.production:", err)
		}
	} else {
		if _, err := os.Stat(".env.production"); err == nil {
			// .env.production exists - prefer it for local production-like runs
			if err := godotenv.Load(".env.production"); err != nil {
				log.Println("warning: failed loading .env.production:", err)
			} else {
				// mark runtime as production when loading .env.production
				os.Setenv("GO_ENV", "production")
			}
		} else {
			if err := godotenv.Load(".env"); err != nil {
				log.Println("warning: no .env file found or error loading .env file")
			}
		}
	}

	// Compose DB connection string: prefer DATABASE_URL, otherwise build from DB_* variables
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	if dbURL == "" {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		name := os.Getenv("DB_NAME")
		ssl := os.Getenv("DB_SSLMODE")
		if ssl == "" {
			ssl = "disable"
		}
		if host != "" && port != "" && user != "" && name != "" {
			dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC", host, user, pass, name, port, ssl)
			log.Println("info: built DB DSN from DB_* environment variables")
		}
	}

	fmt.Println("Auth microservice starting...")
	fmt.Printf("PORT from env: %s\n", os.Getenv("PORT"))
	fmt.Println("Database URL: configured")
	fmt.Println("Kafka Broker:", kafkaBroker)

	// Validate required environment variables
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required (or provide DB_HOST/DB_USER/DB_NAME etc.)")
	}
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	fmt.Println("Starting initialization...")

	// Initialize database connection variable
	var db *gorm.DB
	dbReady := make(chan bool, 1)

	// Connect to database asynchronously
	go func() {
		fmt.Println("Attempting to connect to the database...")
		var err error
		db, err = connectWithRetry(dbURL, 12, 2*time.Second)
		if err != nil {
			log.Printf("ERROR: Database connection failed: %v", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("ERROR: Failed to get underlying sql.DB: %v", err)
			return
		}

		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)

		fmt.Println("Database connection successful!")

		// Run migrations and seeders
		fmt.Println("Running specification translation verification...")
		if err := database_migrations.TranslateEnglishSpecifications(db); err != nil {
			log.Printf("Warning: Specification translation failed: %v", err)
		}

		fmt.Println("Converting specifications to Bengali...")
		if err := database_migrations.ConvertSpecificationsAfter9833ToBengali(db); err != nil {
			log.Printf("Warning: Specifications conversion failed: %v", err)
		}

		var userCount int64
		err = db.Table("users").Count(&userCount).Error
		if err != nil {
			log.Printf("Failed to check users table: %v", err)
		} else if userCount == 0 {
			fmt.Println("Running initial data seeder...")
			if err := database_seeders.SetupAllSeeders(db).RunAll(); err != nil {
				log.Printf("Seeding failed: %v", err)
			}
			fmt.Println("Seeding complete!")
		}

		dbReady <- true
	}()

	routesReady := make(chan bool, 1)

	// Initialize router
	mux := http.NewServeMux()

	// Register a basic health check route immediately.
	// This is crucial for deployment platforms like Railway/Heroku
	// which need a fast-responding endpoint to confirm the server started.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // Explicitly set 200 OK
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Asynchronously register other routes that depend on the database
	go func() {
		// Wait for the database to be ready
		<-dbReady
		if db == nil {
			log.Println("ERROR: Database connection failed, routes will not be available")
			return
		}

		fmt.Println("Initializing repositories...")
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

		fmt.Println("Registering API routes...")
		handlerauth.RegisterAuthRoutes(mux, userRepo, refreshTokenRepo)
		handleruser.RegisterUserRoutes(mux, userRepo)
		handlerproduct.RegisterProductRoutes(mux, productRepo, imageRepo, categoryRepo, brandRepo, productReviewRepo)
		handlercategory.RegisterCategoryRoutes(mux, categoryRepo)
		handlerbrand.RegisterBrandRoutes(mux, brandRepo)
		handlerspecification.RegisterSpecificationRoutes(mux, specificationRepo, specificationKeyRepo, productRepo, formGeneratorRepo)
		handlerproductreview.RegisterProductReviewRoutes(mux, productReviewRepo, productRepo, imageRepo)
		handlerformgenerator.RegisterRoutes(mux, formGeneratorRepo)
		handlerfeedback.RegisterRoutes(mux, feedbackRepo)

		fmt.Println("[STARTUP] All application routes have been registered.")
		routesReady <- true
	}()

	// Determine port for the server
	preferredPort := 8080
	portEnv := os.Getenv("PORT")
	fmt.Printf("[STARTUP] PORT env var: '%s'\n", portEnv)
	if portEnv != "" {
		if p, err := strconv.Atoi(portEnv); err == nil {
			preferredPort = p
			fmt.Printf("[STARTUP] Using PORT from env: %d\n", p)
		} else {
			fmt.Printf("[STARTUP] Failed to parse PORT: %v\n", err)
		}
	}
	availablePort := preferredPort

	if availablePort != preferredPort {
		fmt.Printf("⚠️   Port %d is in use, using port %d instead\n", preferredPort, availablePort)
	}

	serverAddr := fmt.Sprintf("0.0.0.0:%d", availablePort)
	fmt.Printf("[STARTUP] Server will bind to: %s\n", serverAddr)

	// Create HTTP server with CORS middleware
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Channel to listen for interrupt signal to terminate gracefully
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	serverReady := make(chan bool, 1)
	go func() {
		fmt.Printf("[STARTUP] Starting ListenAndServe on %s\n", serverAddr)
		serverReady <- true
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("ERROR: ListenAndServe failed: %v", err)
		}
	}()

	// Wait for server to be ready before proceeding
	<-serverReady
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("[STARTUP] Server ready and accepting connections\n")

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
