package data

import (
	"context"
	"log"
	"taskmanager/db"
	"taskmanager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection

func init() {
	taskCollection = db.Client.Database("TaskManager").Collection("Tasks")
}

func GetAllTasks(userID, role string) []models.Task {
	var tasks []models.Task
	filter := bson.D{{Key: "userid", Value: userID}}

	if role == "admin" {
		filter = bson.D{}
	}
	cursor, err := taskCollection.Find(context.Background(), filter)
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

func GetTaskByID(id, userid, role string) (*models.Task, error) {
	filter := bson.D{{Key: "id", Value: id}, {Key: "userid", Value: userid}}
	if role == "admin" {
		filter = bson.D{{Key: "id", Value: id}}
	}
	var result models.Task
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&result)
	return &result, err
}
func UpdateTask(id, userid, role string, newTask *models.Task) {
	filter := bson.D{{Key: "id", Value: id}, {Key: "userid", Value: userid}}
	if role == "admin" {
		filter = bson.D{{Key: "id", Value: id}}
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: newTask.Title},
			{Key: "description", Value: newTask.Description},
			{Key: "duedate", Value: newTask.DueDate},
			{Key: "status", Value: newTask.Status},
		}},
	}
	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
func AddTask(task *models.Task) (interface{}, error) {
	insertResult, err := taskCollection.InsertOne(context.TODO(), task)
	return insertResult.InsertedID, err
}

func DeleteTask(id, userid, role string) error {
	filter := bson.D{{Key: "id", Value: id}, {Key: "userid", Value: userid}}
	if role == "admin" {
		filter = bson.D{{Key: "id", Value: id}}
	}
	_, err := taskCollection.DeleteOne(context.TODO(), filter)

	return err
}
