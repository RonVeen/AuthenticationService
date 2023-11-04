package internal

import (
	"github.com/gin-gonic/gin"
	"os"
)

type server struct {
	userRepo    *UserRepo
	secretsRepo *SecretsRepo
	Router      *gin.Engine
}

func NewServer(ur *UserRepo, sr *SecretsRepo, r *gin.Engine) *server {
	return &server{
		userRepo:    ur,
		secretsRepo: sr,
		Router:      r,
	}
}

func (s *server) Run() {
	s.Routes()
	if os.Getenv("PORT") == "" {
		s.Router.Run()
	} else {
		s.Router.Run(":" + os.Getenv("PORT"))

	}
}
