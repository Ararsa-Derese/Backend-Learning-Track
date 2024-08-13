package usecase_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"cleantaskmanager/usecase"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskUseCase_AddTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	taskUseCase := usecase.NewTaskUsecase(mockRepo)

	tasks := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "pending",
		UserID:      primitive.NewObjectID(),
	}
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	mockRepo.On("AddTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("*domain.Task")).Return(nil)
	err := taskUseCase.AddTask(context.Background(), &claims, &tasks)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaskUseCase_GetTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	taskUseCase := usecase.NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID()
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	mockRepo.On("GetTask", mock.Anything,mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(&domain.Task{}, nil)
	task, err := taskUseCase.GetTask(context.Background(),&claims, taskID)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	mockRepo.AssertExpectations(t)
}

func TestTaskUseCase_GetTasks(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	taskUseCase := usecase.NewTaskUsecase(mockRepo)

	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	mockRepo.On("GetTasks", mock.Anything, mock.AnythingOfType("*domain.Claims")).Return([]domain.Task{}, nil)
	tasks, err := taskUseCase.GetTasks(context.Background(), &claims)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
	mockRepo.AssertExpectations(t)
}

func TestTaskUseCase_UpdateTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	taskUseCase := usecase.NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID()
	task := domain.UpdateTask{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     "2021-01-01T00:00:00Z",
		Status:      "pending",
	}
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	mockRepo.On("UpdateTask", mock.Anything, mock.AnythingOfType("*domain.Claims"),mock.AnythingOfType("primitive.ObjectID"),mock.AnythingOfType("*domain.UpdateTask")).Return(nil)
	err := taskUseCase.UpdateTask(context.Background(),&claims, taskID,&task)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaskUseCase_DeleteTask(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	taskUseCase := usecase.NewTaskUsecase(mockRepo)
	taskID := primitive.NewObjectID()
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	mockRepo.On("DeleteTask", mock.Anything,mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(nil)
	err := taskUseCase.DeleteTask(context.Background(), &claims,taskID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}