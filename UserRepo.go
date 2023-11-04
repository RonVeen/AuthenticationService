package main

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	InsertUser  = "INSERT INTO Users VALUES ($1, $2, $3);"
	SelectUser  = "SELECT uuid, email, status FROM Users WHERE email=$1;"
	FindByUUID  = "SELECT uuid, email, status FROM USERS WHERE uuid=$1;"
	UpdateToken = "UPDATE Tokens SET token=$2, expires=$3 WHERE uuid=$1;"
	SelectToken = "SELECT token FROM Tokens WHERE uuid=$1;"
	InsertToken = "INSERT INTO Tokens VALUES($1, $2, $3);"
	FindByToken = "SELECT uuid, expires FROM Tokens where token=$1;"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (s *Server) FindByEmail(email string) (*User, error) {
	var user User
	if err := s.userRepo.db.Get(&user, SelectUser, email); err != nil {
		return &User{}, err
	}
	return &user, nil

}

func (s *Server) Create(uuid string, email string) error {
	if _, err := s.userRepo.db.Exec(InsertUser, uuid, email, Enabled); err != nil {
		return err
	}
	return nil
}

func (s *Server) UpdateToken(uuid string, token string, expires time.Time) error {
	var existing string
	err := s.userRepo.db.Get(&existing, SelectToken, uuid)
	if err != nil {
		if s.userRepo.db.Exec(InsertToken, uuid, token, expires); err != nil {
			return err
		}
	} else {
		if _, err := s.userRepo.db.Exec(UpdateToken, uuid, token, expires); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) FindUserByToken(token string) (*User, error) {
	type usertoken struct {
		Uuid    string
		Expires time.Time
	}
	var userToken usertoken
	if err := s.userRepo.db.Get(&userToken, FindByToken, token); err != nil {
		return nil, err
	}

	if time.Now().After(userToken.Expires) {
		return nil, errors.New("token expired")
	}

	var user User
	if err := s.userRepo.db.Get(&user, FindByUUID, userToken.Uuid); err != nil {
		return nil, err
	}

	return &user, nil
}
