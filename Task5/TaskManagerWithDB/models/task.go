package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	ID          string `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	DueDate     string `validate:"required"`
	Status      string `validate:"required"`
}

func (t *Task) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

