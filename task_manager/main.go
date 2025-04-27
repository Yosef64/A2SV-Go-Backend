package main

import (
	"context"
	"log"
	"task_manager/data"
	"task_manager/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	data.TaskCollection = client.Database("task_management_api").Collection("tasks")
	data.UserCollection = client.Database("task_management_api").Collection("users")
	router.SetupRouter().Run(":8080")

}
