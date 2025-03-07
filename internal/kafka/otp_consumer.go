package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vithsutra/vithsutra_email_service/config"
	"github.com/vithsutra/vithsutra_email_service/internal/email"
)

type OTPMessage struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func StartOTPConsumer() {
	if len(config.Config.Kafka.Topics) < 1 {
		log.Fatal("[ERROR] No topics found in config.yaml")
	}

	topic := config.Config.Kafka.Topics[0] // "email.otp"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: config.Config.Kafka.Brokers,
		GroupID: config.Config.Kafka.GroupID,
		Topic:   topic,
	})

	defer reader.Close()
	log.Printf("[INFO] Listening to OTP topic: %s", topic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("[ERROR] Error reading OTP message: %v", err)
			continue
		}

		var otpMessage OTPMessage
		err = json.Unmarshal(msg.Value, &otpMessage)
		if err != nil {
			log.Printf("[ERROR] Failed to parse OTP message: %v", err)
			continue
		}

		// Prepare email data
		emailData := map[string]string{
			"OTP": otpMessage.OTP,
		}

		// Send Email
		err = email.SendEmail(otpMessage.Email, "Your OTP Code", "templates/otp.html", emailData)
		if err != nil {
			log.Printf("[ERROR] Failed to send OTP email: %v", err)
		}
	}
}
