package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores application configuration
var Config struct {
	Kafka struct {
		Brokers []string
		GroupID string
		Topics  []string
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		From     string
	}
	
}

// LoadConfig loads configuration from YAML file
func LoadConfig() error {
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/") // Looks inside "config/" directory
	viper.AddConfigPath(".")       // Also checks the root directory

	viper.AutomaticEnv() // Enables environment variable overrides

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] Failed to read config file: %v", err)
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("[ERROR] Failed to unmarshal config: %v", err)
		return err
	}

	// Debugging: Print loaded topics
	log.Printf("[DEBUG] Loaded Kafka Topics: %v", Config.Kafka.Topics)

	return nil
}
