package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddress  string
	SentryEnabled bool
	ServerPort    string
}

var instance *Config
var once sync.Once

func Init() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	once.Do(func() {
		instance = &Config{
			RedisAddress:  os.Getenv("REDIS_ADDRESS"),
			SentryEnabled: os.Getenv("SENTRY_ENABLED") == "true",
			ServerPort:    os.Getenv("SERVER_PORT"),
		}
	})
}

func GetInstance() *Config {
	return instance
}
