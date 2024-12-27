package handlers

import "github.com/thecodingmontana/jwt-go/internal/database/models"

type APIConfig struct {
	DB *models.Queries
}

func NewAPIConfig(queries *models.Queries) *APIConfig {
	return &APIConfig{
		DB: queries,
	}
}
