package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "Users"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}

type Login struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Password string             `json:"password" bson:"password"`
}

type UserRepository interface {
	RegisterUser(user *User) error
	GetUserByID(id primitive.ObjectID) (*User, error)
}
type UserUsecase interface {
	RegisterUser(user *User) error
	LoginUser(user *User) (string, error)
	GetUserByID(id primitive.ObjectID) (*User, error)
}
