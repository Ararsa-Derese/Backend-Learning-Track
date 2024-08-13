package usecase_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"cleantaskmanager/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUserUsecase_GetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userUsecase := usecase.NewUserUsecase(mockRepo)
	user := domain.User{Username: "testuser", Password: "password"}
	mockRepo.On("GetUserByID", user.ID).Return(&user, nil)
	result, err := userUsecase.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)

}


func TestUserUsecase_RegisterUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userUsecase := usecase.NewUserUsecase(mockRepo)
	user := domain.User{Username: "testuser", Password: "password"}
	mockRepo.On("RegisterUser", &user).Return(nil)
	err := userUsecase.RegisterUser(&user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
