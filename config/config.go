package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI      string
	DatabaseName  string
	ServerPort    string
	ServerTimeout time.Duration
}

func Load() (*Config, error) {
	// Load .env file if exists
	_ = godotenv.Load(".env")

	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DatabaseName:  getEnv("DATABASE_NAME", "pharmacy"),
		ServerPort:    getEnv("SERVER_PORT", "1323"),
		ServerTimeout: time.Second * 30,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
