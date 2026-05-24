package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ProjectID      string
	SubscriptionID string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectID := os.Getenv("GCP_PROJECT_ID")
	subscriptionID := os.Getenv("GCP_SUBSCRIPTION_ID")

	if projectID == "" || subscriptionID == "" {
		log.Fatal("GCP_PROJECT_ID and GCP_SUBSCRIPTION_ID must be set")
	}

	return Config{
		ProjectID:      projectID,
		SubscriptionID: subscriptionID,
	}
}
