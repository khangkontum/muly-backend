package user

import (
	"crypto/sha256"
	"errors"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type password struct {
	plaintext *string
	hash      [32]byte
}

type UserRepository interface {
	Set(plaintextPassword string) error
	Matches(plaintextPassword string) (bool, error)
}

func (p *password) Set(plaintextPassword string) error {
	hash := sha256.Sum256([]byte(plaintextPassword))

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {

	hash := sha256.Sum256([]byte(plaintextPassword))
	if hash != p.hash {
		return false, errors.New("Password Mismatched")
	}
	return true, nil
}
