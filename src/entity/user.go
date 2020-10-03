package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// NewUser create a new user
func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Prepare is responsible for setting the default values
func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = string(password)

	return nil
}
