package app

import (
	"os"
)

type Config struct {
	ServerPort string
	RedisHost  string
	RedisPort  string
}

func GetEnv(key, dfValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return dfValue
	}
	return value
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort: GetEnv("SERVER_PORT", "8000"),
		RedisHost:  GetEnv("REDIS_HOST", "localhost"),
		RedisPort:  GetEnv("REDIS_PORT", "6379"),
	}

	return cfg
}
