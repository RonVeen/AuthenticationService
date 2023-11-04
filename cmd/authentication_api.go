package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ronveen/AuthenticationService/internal"
)

func main() {
	internal.LoadEnvVars()

	db := internal.SetupDatabase()
	server := internal.NewServer(
		internal.NewUserRepo(db),
		internal.NewSecretsRepo(db),
		gin.Default(),
	)
	server.Execute()
}
