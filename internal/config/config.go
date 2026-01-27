package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        string
	RedisUrl    string
	JwtKey      string
}

var required = []string{
	"DATABASE_URL",
	"REDIS_URL",
	"JWT_KEY"}

func Load() *Config {
	_ = godotenv.Load()

	for _, env := range required {
		if os.Getenv(env) == "" {
			log.Fatalf("%v environment variable is required", env)
		}
	}

	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		RedisUrl:    os.Getenv("REDIS_URL"),
		JwtKey:      os.Getenv("JWT_KEY"),
	}
}
