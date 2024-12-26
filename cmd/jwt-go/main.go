package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/thecodingmontana/jwt-go/internal/routes"
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

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)

	routes.RegisterRoutes(router)

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
