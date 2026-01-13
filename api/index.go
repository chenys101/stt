package handler

import (
	"log"
	"net/http"
	"backend/internal/config"
	"backend/internal/pkg/database"
	"backend/internal/route"
)

// Handler is the serverless entry point.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Load config (you might need to adjust the path for Vercel environment variables)
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	if err := database.Connect(cfg.Database.DSN); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	// Setup router
	router := route.SetupRouter()

	// Let the router handle the request
	router.ServeHTTP(w, r)
}
