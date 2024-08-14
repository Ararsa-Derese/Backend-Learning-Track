package usecase

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (lu *UserUsecase) RegisterUser(user *domain.User) error {
	return lu.userRepository.RegisterUser(user)
}

func (lu *UserUsecase) LoginUser(user *domain.User) (string ,error) {
	jwttoken , error := infrastructure.GenerateToken(user)
	return jwttoken , error
	
}

func (lu *UserUsecase) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	return lu.userRepository.GetUserByID(id)
}

