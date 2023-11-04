package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ronveen/AuthenticationService/internal"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "fueling"
	password = "s3cr3t"
	dbname   = "fueling"
)

func main() {
	loadEnvironment()

	db := internal.SetupDatabase()
	server := internal.NewServer(
		internal.NewUserRepo(db),
		internal.NewSecretsRepo(db),
		gin.Default(),
	)
	server.Run()
}

func loadEnvironment() {
	var err error
	if os.Getenv("APP_ENV") == "" {
		err = godotenv.Load("./envs/.env")
	} else {
		err = godotenv.Load("./envs/" + os.Getenv("APP_ENV") + ".env")
	}
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
