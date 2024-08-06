package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB!")
		log.Fatal(err)

	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Error connecting to MongoDB!")
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Client = client
}
