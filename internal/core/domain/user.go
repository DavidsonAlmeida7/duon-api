package domain

import (
	"errors"
	"time"
)

// Erros de domínio permitem que os adapters saibam exatamente o que falhou
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserInvalidEmail  = errors.New("invalid email address")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrUserInvalidEmail
	}

	// Adicione outras validações de domínio aqui

	return nil
}

//func (u *User) Activate() {
//	u.Active = true
//}
//
//func (u *User) Deactivate() {
//	u.Active = false
//}
