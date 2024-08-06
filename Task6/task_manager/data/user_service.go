package data

import (
	"context"
	"os"
	"taskmanager/db"
	"taskmanager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var jwtSecret []byte

var userCollection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		// Handle error (optional)
		panic("Error loading .env file")
	}

	// Read JWT_SECRET from environment variables
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	userCollection = db.Client.Database("TaskManager").Collection("Users")

}

func RegisterUser(user *models.User) (interface{}, error) {
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	return insertResult.InsertedID, err
}
func GenerateToken(user *models.User) (string, error) {

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Minute * 40).Unix(),
	})
	jwtToken, err := token.SignedString(jwtSecret)
	return jwtToken, err

}
func GetUserByID(id string) (*models.User, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var result models.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}

// func ParseToken(tokenString string) (jwt.MapClaims, error) {
//     token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//         // Don't forget to validate the alg is what you expect:

//         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//             return nil, errors.New("unexpected signing method")
//         }
//         return jwtSecret, nil
//     })

//     if err != nil {
//         return nil, err
//     }

//     if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//         return claims, nil
//     } else {
//         return nil, errors.New("invalid token")
//     }
// }
