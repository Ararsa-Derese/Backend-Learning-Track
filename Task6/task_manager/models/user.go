package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID       string `validate:"-"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// NewUser creates a new user with a unique ID.
func NewUser(username, password, role string) *User {
	userUUID := uuid.New().String()
	// Create a unique ID by combining the user's ID and the task's UUID.
	uniqueUserID := fmt.Sprintf("%s-%s", username, userUUID)
	return &User{
		ID:       uniqueUserID,
		Username: username,
		Password: password,
		Role:     role,
	}
}
