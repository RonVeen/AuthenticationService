package internal

import (
	"github.com/jmoiron/sqlx"
)

type SecretsRepo struct {
	db *sqlx.DB
}

func NewSecretsRepo(db *sqlx.DB) *SecretsRepo {
	return &SecretsRepo{db: db}
}

func (s *server) StoreSecret(uuid string, password string) error {
	encrypted, err := EncryptPassword(password)
	if err != nil {
		return err
	}
	s.secretsRepo.db.Exec("INSERT INTO Secrets values ($1, $2)", uuid, encrypted)
	return nil
}

func (s *server) Matches(uuid string, password string) bool {
	var hash string
	if err := s.secretsRepo.db.Get(&hash, "SELECT hash FROM Secrets WHERE uuid=$1", uuid); err != nil {
		return false
	}

	if err := Compare(hash, password); err != nil {
		return false
	}

	return true
}
