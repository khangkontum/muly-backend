package domain

import (
	"errors"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  Password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type Password struct {
	Plaintext *string
	Hash      []byte
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, user *User) (User, error)
}
type UserUsecase interface {
	Insert(c *gin.Context, user *User) error
	GetByEmail(c *gin.Context, email string) error
	Update(c *gin.Context, user *User) error
}

func (p *Password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Plaintext = &plaintextPassword
	p.Hash = hash

	return nil
}

func (p *Password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
