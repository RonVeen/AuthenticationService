package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"time"
)

const TokenLifetime = 60 * 60 * 24 // Tokens life 24 hours

type emailAndPassword struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (s *Server) handleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials emailAndPassword
		c.Bind(&credentials)

		user, err := s.FindByEmail(credentials.Email)
		if err != nil {
			c.JSON(401, gin.H{"Reason": "User/Password combination invalid"})
			return
		}

		if user.Status != Enabled {
			c.JSON(401, gin.H{"Reason": "User not enabled"})
			return
		}

		if s.Matches(user.Uuid, credentials.Password) == false {
			c.JSON(401, gin.H{"Reason": "User/Password combination invalid"})
			return
		}

		token, err := randomHex(30)
		if err != nil {
			c.JSON(401, gin.H{"Status": "Failed to generate token"})
		}

		var expires = time.Now().Add(time.Duration(time.Second * TokenLifetime))
		s.UpdateToken(user.Uuid, token, expires)

		c.JSON(200, gin.H{"Status": token})
	}
}

func (s *Server) handleCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials emailAndPassword
		c.Bind(&credentials)

		if err := validateEmail(credentials.Email); err != nil {
			c.JSON(404, gin.H{"Reason": err.Error()})
			return
		}

		if err := validatePassword(credentials.Password); err != nil {
			c.JSON(404, gin.H{"Reason": err.Error()})
			return
		}

		uuid := uuid2.New().String()
		if err := s.StoreSecret(uuid, credentials.Password); err != nil {
			c.JSON(401, gin.H{"Reason": err.Error()})
		}

		if err := s.Create(uuid, credentials.Email); err != nil {
			c.JSON(401, gin.H{"Reason": err.Error()})
		}

		c.Status(201)

	}
}

func (s *Server) handleToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		user, err := s.FindUserByToken(token)
		if err != nil {
			c.JSON(401, gin.H{"Reason": "Invalid or expired token"})
			return
		}

		if user.Status != Enabled {
			c.JSON(403, gin.H{"Reason": "Account is not enabled"})
			return
		}

		c.JSON(200, &user)
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
