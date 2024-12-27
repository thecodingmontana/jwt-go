package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/thecodingmontana/jwt-go/internal/database/models"
	"github.com/thecodingmontana/jwt-go/internal/routes"
	"github.com/thecodingmontana/jwt-go/pkg/database"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Environmental Variables
	PORT, portFound := os.LookupEnv("PORT")

	if !portFound {
		log.Fatalf("PORT is missing in the .env file")
	}

	// Environmental Variables
	DATABASE_URL, dbURLFound := os.LookupEnv("DATABASE_URL")

	if !dbURLFound {
		log.Fatalf("DATABASE_URL is missing in the .env file")
	}

	// connect to Database
	conn := database.ConnectDB(DATABASE_URL)
	defer conn.Close()

	// initialize SQLC queries
	queries := models.New(conn)

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routes.RegisterRoutes(router, queries)

	// Server instance
	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("🚀 Server started at http://localhost:%s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("🚀 Failed to start the server!")
	}
}
