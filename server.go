package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	port        string
	router      *gin.Engine
	userRepo    *UserRepo
	secretsRepo *SecretsRepo
}
