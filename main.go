package main

import (
	"log"

	"github.com/vithsutra/vithsutra_email_service/config"
	"github.com/vithsutra/vithsutra_email_service/internal/kafka"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[ERROR] Failed to load config: %v", err)
	}

	go kafka.StartOTPConsumer()
	go kafka.StartWelcomeConsumer()

	select {} // Keep the application running
}
