package internal

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"strings"
)

const pepper = ""

const MinimumPasswordLength = 10
const RequiresAlpha = true
const RequiresNumberic = true
const RequiresSpecial = true
const RequiresMixedCase = true

const Alphas = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const Specials = "!@#$%&*()[]<>?/.,-+=~`_;:|\\"
const Numerics = "1234567890"

type SecretsRepository interface {
	StoreSecret(uuid uuid.UUID, password string) error
	Matches(uuid uuid.UUID, password string) bool
}

func validateEmail(email string) error {
	if len(strings.TrimSpace(email)) != 0 {
		_, err := mail.ParseAddress(email)
		if err == nil {
			return nil
		}
	}
	return errors.New("Invalid e-mail address")
}

func validatePassword(password string) error {
	if matchesMinimumLengthRule(password) {
		if matchesMixedCaseRule(password) {
			if matchesAlphaRule(password) {
				if matchesNumericRule(password) {
					if matchesSpecialRule(password) {
						return nil
					}
				}
			}
		}
	}
	return errors.New("Password does not match criteria")
}

func matchesMinimumLengthRule(password string) bool {
	return len(strings.TrimSpace(password)) >= MinimumPasswordLength
}

func matchesMixedCaseRule(password string) bool {
	if RequiresMixedCase {
		return strings.ToUpper(password) != strings.ToLower(password)
	}
	return true
}

func matchesAlphaRule(password string) bool {
	if RequiresAlpha {
		return contains(password, Alphas)
	}
	return true
}

func matchesNumericRule(password string) bool {
	if RequiresNumberic {
		return contains(password, Numerics)
	}
	return true
}

func matchesSpecialRule(password string) bool {
	if RequiresSpecial {
		return contains(password, Specials)
	}
	return true
}

func EncryptPassword(password string) (string, error) {
	base := password + pepper
	bytes, err := bcrypt.GenerateFromPassword([]byte(base), 1)
	return string(bytes), err
}

func Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func contains(subject string, required string) bool {
	for _, s := range subject {
		for _, r := range required {
			if s == r {
				return true
			}
		}
	}
	return false
}
