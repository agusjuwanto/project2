package config

import "os"

type Config struct {
    AppName string
    Env     string
    Port    string
}

func Load() Config {
    return Config{
        AppName: getEnv("APP_NAME", "gin-rest-api"),
        Env:     getEnv("APP_ENV", "development"),
        Port:    getEnv("APP_PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
