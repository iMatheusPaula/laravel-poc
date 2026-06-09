package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ProjectID      string
	SubscriptionID string
	SMTPHost       string
	SMTPPort       int
	SMTPUsername   string
	SMTPPassword   string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectID := os.Getenv("GCP_PROJECT_ID")
	subscriptionID := os.Getenv("GCP_SUBSCRIPTION_ID")
	SMTPHost := os.Getenv("SMTP_HOST")
	SMTPPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	SMTPUsername := os.Getenv("SMTP_USERNAME")
	SMTPPassword := os.Getenv("SMTP_PASSWORD")

	if err != nil {
		log.Fatal("SMTP_PORT invalid", err)
	}

	if projectID == "" || subscriptionID == "" {
		log.Fatal("GCP_PROJECT_ID and GCP_SUBSCRIPTION_ID must be set")
	}

	return Config{
		ProjectID:      projectID,
		SubscriptionID: subscriptionID,
		SMTPHost:       SMTPHost,
		SMTPPort:       SMTPPort,
		SMTPUsername:   SMTPUsername,
		SMTPPassword:   SMTPPassword,
	}
}
