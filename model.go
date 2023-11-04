package main

import "github.com/google/uuid"

type UserStatus string

const (
	Enabled              UserStatus = "Enabled"
	AwaitingConfirmation            = "AwaitingConfirmation"
	TemporarySuspended              = "TemporarySuspended"
	PermanentlySuspended            = "PermanentlySuspended"
	Blocked                         = "Blocked"
	Closed                          = "Closed"
)

type User struct {
	Email  string
	Uuid   string
	Status UserStatus
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Create(uuid uuid.UUID, email string)
	UpdateToken(uuid string, token string) bool
}
