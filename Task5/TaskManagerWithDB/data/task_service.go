package data

import (
	"context"
	"fmt"
	"log"
	"taskmanagerdb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database("TaskManager").Collection("Tasks")
}

func GetAllTasks() []models.Task {
	var tasks []models.Task
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var result models.Task
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}
func UpdateTask(id string, newTask *models.Task) {
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: newTask.Title},
			{Key: "description", Value: newTask.Description},
			{Key: "duedate", Value: newTask.DueDate},
			{Key: "status", Value: newTask.Status},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
func AddTask(task *models.Task) (interface{}, error) {
	insertResult, err := collection.InsertOne(context.TODO(), task)
	return insertResult.InsertedID, err
}

func DeleteTask(id string) error {
	filter := bson.D{{Key: "id", Value: id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	fmt.Println(err)
	return err
}
