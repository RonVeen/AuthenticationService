package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	db := setupDatabase()
	server := Server{
		userRepo:    NewUserRepo(db),
		secretsRepo: NewSecretsRepo(db),
		router:      gin.Default(),
	}
	server.routes()
	server.router.Run()
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
