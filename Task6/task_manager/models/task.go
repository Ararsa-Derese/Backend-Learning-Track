package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID          string `validate:"-"` // "-" tag means this field is excluded from validation
	Title       string `validate:"required"`
	Description string `validate:"required"`
	DueDate     string `validate:"required"`
	Status      string `validate:"required"`
	UserID      string `validate:"required"`
}

func (t *Task) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// NewTask creates a new task with a unique ID for a given user.
func NewTask(userID, title, description, dueDate, status string) *Task {
	// Generate a new UUID for the task.
	taskUUID := uuid.New().String()

	// Create a unique ID by combining the user's ID and the task's UUID.
	uniqueTaskID := fmt.Sprintf("%s-%s", userID, taskUUID)

	return &Task{
		ID:          uniqueTaskID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
		UserID:      userID,
	}
}
