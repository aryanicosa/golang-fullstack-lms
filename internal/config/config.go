package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    // Postgres configs
	DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    ServerPort string

	// Redis configs
    RedisHost     string
    RedisPort     string
    RedisPassword string
    RedisDB       int
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        JWTSecret:  os.Getenv("JWT_SECRET"),
        ServerPort: os.Getenv("SERVER_PORT"),

		RedisHost:     os.Getenv("REDIS_HOST"),
        RedisPort:     os.Getenv("REDIS_PORT"),
        RedisPassword: os.Getenv("REDIS_PASSWORD"),
        RedisDB:       0, // Default database
    }, nil
}