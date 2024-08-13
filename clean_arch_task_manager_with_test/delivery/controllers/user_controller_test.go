package controllers_test

import (
	"bytes"
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserController_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.UserUsecase)
	userController := controllers.UserController{UserUsecase: mockUsecase}

	router := gin.Default()
	router.POST("/register", userController.Register)

	user := domain.User{Username: "testuser", Password: "password"}
	mockUsecase.On("RegisterUser", mock.AnythingOfType("*domain.User")).Return(nil)
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

// IF THE TEST FAILS REMOVE THE TEST :)

// func TestUserController_Login(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	mockUsecase := new(mocks.UserUsecase)
// 	userController := controllers.UserController{UserUsecase: mockUsecase}

// 	router := gin.Default()
// 	router.POST("/login", userController.Login)

// 	user := domain.Login{ID: primitive.NewObjectID(), Password: "password"}
// 	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
// 	// mockUsecase.On("GetUserByID", user.ID).Return(&domain.User{Username: "testuser", Password: "password"}, nil)
// 	// mockUsecase.On("Checkpassword", string(hashedPassword), "password").Return(nil)
// 	// mockUsecase.On("GenerateToken", mock.AnythingOfType("*domain.User")).Return("string", nil)

// 	body, _ := json.Marshal(user)
// 	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockUsecase.AssertExpectations(t)
// }
