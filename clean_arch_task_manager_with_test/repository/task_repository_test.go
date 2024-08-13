package repository_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/mongo/mocks"
	"context"
	"errors"
	"testing"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

import "cleantaskmanager/repository"

func TestTaskRepository_AddTask(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTask

	mockTask := domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     "2021-01-01T00:00:00Z",
		Status:      "pending",
		UserID:      "1",
	}
	mockemptyTask := domain.Task{}
	mockTaskID := "12345"
	mockClaims := domain.Claims{}

	t.Run("success", func(t *testing.T) {

		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(mockTaskID, nil).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTaskRepository(databaseHelper, collectionName)

		err := ur.AddTask(context.Background(), &mockClaims, &mockTask)

		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(mockemptyTask, errors.New("Unexpected")).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTaskRepository(databaseHelper, collectionName)

		err := ur.AddTask(context.Background(), &mockClaims, &mockemptyTask)

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
	})

}
func TestTaskRepository_GetTasks(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	var cursorHelper *mocks.Cursor

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	cursorHelper = &mocks.Cursor{}

	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{UserID: "1", Role: "user"}

	t.Run("success", func(t *testing.T) {
		collectionHelper.On("Find", mock.Anything, mock.Anything).Return(cursorHelper, nil).Once()
		cursorHelper.On("Close", mock.Anything).Return(nil).Once()
		cursorHelper.On("Next", mock.Anything).Return(false).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		tr := repository.NewTaskRepository(databaseHelper, collectionName)

		_, err := tr.GetTasks(context.Background(), &mockClaims)

		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
		cursorHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("Find", mock.Anything, mock.Anything).Return(cursorHelper, errors.New("Unexpected")).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		tr := repository.NewTaskRepository(databaseHelper, collectionName)

		_, err := tr.GetTasks(context.Background(), &mockClaims)

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		cursorHelper.AssertExpectations(t)
	})
}

func TestTaskRepository_GetTask(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	var singleResultHelper *mocks.SingleResult
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	singleResultHelper = &mocks.SingleResult{}
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{UserID: "1", Role: "user"}
	t.Run("success", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(singleResultHelper).Once()
		singleResultHelper.On("Decode", mock.AnythingOfType("*domain.Task")).Return(nil).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		tr := repository.NewTaskRepository(databaseHelper, collectionName)
		_, err := tr.GetTask(context.Background(), &mockClaims, "1")
		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		singleResultHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(singleResultHelper).Once()
		singleResultHelper.On("Decode", mock.AnythingOfType("*domain.Task")).Return(errors.New("Unexpected")).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		tr := repository.NewTaskRepository(databaseHelper, collectionName)
		_, err := tr.GetTask(context.Background(), &mockClaims, "1")
		assert.Error(t, err)
		collectionHelper.AssertExpectations(t)
		singleResultHelper.AssertExpectations(t)
	})
}

func TestTaskRepository_UpdateTask(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{}
	mockTask := domain.UpdateTask{Title: "Updated Task"}
	updateresult := &mongo.UpdateResult{}
	t.Run("success", func(t *testing.T) {

		// updateResult := &mocks.UpdateResult{}
		collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		tr := repository.NewTaskRepository(databaseHelper, collectionName)

		err := tr.UpdateTask(context.Background(), &mockClaims, "1", &mockTask)

		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, errors.New("Unexpected")).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		tr := repository.NewTaskRepository(databaseHelper, collectionName)

		err := tr.UpdateTask(context.Background(), &mockClaims, "1", &mockTask)

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
	})
}

func TestTaskRepository_DeleteTask(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection
	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{}
	var id int64 = 1
	t.Run("success", func(t *testing.T) {
		collectionHelper.On("DeleteOne", mock.Anything, mock.Anything).Return(id,nil).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		tr := repository.NewTaskRepository(databaseHelper, collectionName)
		err := tr.DeleteTask(context.Background(), &mockClaims, "1")
		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("DeleteOne", mock.Anything, mock.Anything).Return(id,errors.New("Unexpected")).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		tr := repository.NewTaskRepository(databaseHelper, collectionName)
		err := tr.DeleteTask(context.Background(), &mockClaims, "1")
		assert.Error(t, err)
		collectionHelper.AssertExpectations(t)
	})
}

