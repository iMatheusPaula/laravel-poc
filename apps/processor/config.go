package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ProjectID string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectID := os.Getenv("GCP_PROJECT_ID")

	if projectID == "" {
		log.Fatal("PROJECT_ID must be set")
	}

	return Config{
		ProjectID: projectID,
	}
}
