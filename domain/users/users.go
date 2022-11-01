package users

import (
	"time"
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
	Insert(user *User) error
	GetByEmail(user *User, email string) error
	Update(user *User) error
}
type UserUsecase interface {
	Insert(user *User) error
	GetByEmail(user *User, email string) error
	Update(user *User) error
}
