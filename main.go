package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"userapi/handlers"
	"userapi/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// @title User API
// @version 1.0
// @description A simple REST API for managing users
// @host localhost:8080
// @BasePath /
func main() {
	// Configure logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting User API service...")

	// Get database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbName := getEnv("DB_NAME", "userdb")
	dbPort := getEnv("DB_PORT", "3306")

	log.Printf("Database configuration: host=%s, port=%s, user=%s, database=%s", dbHost, dbPort, dbUser, dbName)

	// Create database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to database with retry logic
	var db *sql.DB
	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to connect to database (attempt %d/%d)...", i+1, maxRetries)
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v", err)
			db.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Successfully connected to database")
		break
	}

	if err != nil {
		log.Fatalf("Could not connect to database after %d attempts: %v", maxRetries, err)
	}
	defer func() {
		log.Println("Closing database connection...")
		db.Close()
	}()

	// Configure database connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Create handlers
	userRepo := repository.NewMySQLUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)
	pingHandler := handlers.NewPingHandler()

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/ping", pingHandler.Ping).Methods("GET")
	router.HandleFunc("/users", userHandler.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")
	router.HandleFunc("/users", userHandler.List).Methods("GET")

	// Add logging middleware
	router.Use(loggingMiddleware)

	// Serve static documentation
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	log.Printf("API documentation available at http://localhost:%s/docs/", port)
	log.Printf("Health check available at http://localhost:%s/ping", port)
	
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
} 