package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

// func GetEnvFallback(name string, fallback string) string {
// 	geted := os.Getenv(name)
// 	if geted == "" {
// 		return fallback
// 	}
// 	return geted
// }
