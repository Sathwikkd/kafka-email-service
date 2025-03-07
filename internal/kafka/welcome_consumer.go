package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vithsutra/vithsutra_email_service/config"
	"github.com/vithsutra/vithsutra_email_service/internal/email"
)

type WelcomeMessage struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Platform string `json:"platform"`
}

func StartWelcomeConsumer() {
	if len(config.Config.Kafka.Topics) < 2 {
		log.Fatal("[ERROR] No welcome email topic found in config.yaml")
	}

	topic := config.Config.Kafka.Topics[1] // "email.welcome"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: config.Config.Kafka.Brokers,
		GroupID: config.Config.Kafka.GroupID,
		Topic:   topic,
		// Remove StartOffset to let Kafka track last committed offset
		CommitInterval: 0, // Disable auto-commit
	})

	defer reader.Close()
	log.Printf("[INFO] Listening to Welcome topic: %s", topic)

	for {
		msg, err := reader.FetchMessage(context.Background()) // Use FetchMessage instead of ReadMessage
		if err != nil {
			log.Printf("[ERROR] Error reading Welcome message: %v", err)
			continue
		}

		log.Printf("[DEBUG] Raw Welcome Message: %s", string(msg.Value))

		var welcomeMessage WelcomeMessage
		err = json.Unmarshal(msg.Value, &welcomeMessage)
		if err != nil {
			log.Printf("[ERROR] Failed to parse Welcome message: %v", err)
			continue
		}

		// Prepare email data
		emailData := map[string]string{
			"Username": welcomeMessage.Username,
			"Platform": welcomeMessage.Platform,
		}

		// Send Email
		err = email.SendEmail(welcomeMessage.Email, "Welcome to VithSutra!", "templates/welcome.html", emailData)
		if err != nil {
			log.Printf("[ERROR] Failed to send Welcome email: %v", err)
		}

		// ✅ Manually Commit Offset after processing
		err = reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Printf("[ERROR] Failed to commit offset: %v", err)
		} else {
			log.Printf("[INFO] Successfully processed and committed offset: %d", msg.Offset)
		}
	}
}
