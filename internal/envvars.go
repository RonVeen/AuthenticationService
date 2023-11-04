package internal

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVars() {
	var err error
	if os.Getenv("APP_ENV") == "" {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load("./" + os.Getenv("APP_ENV") + ".env")
	}
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
