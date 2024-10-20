package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config defines the values necessary to connect to the database
type Config struct {
	PublicHost string
	Port       string

	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

// Envs contains the initial configuration values for connecting to the database
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load() // This picks up whatever values you have in the .env file. NEVER COMMIT THE .env FILE, as mentioned in .gitignore. This has sensitive data

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "https://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "test"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
