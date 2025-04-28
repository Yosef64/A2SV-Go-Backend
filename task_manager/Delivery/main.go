package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/router"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	usecases "task_manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// secret keys
	jwtKey := "NEVER_GIVE_UP"

	// Defining collections
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

	db := client.Database("task_management_api")
	taskCollection := db.Collection("tasks")
	userCollection := db.Collection("users")
	// initialize services
	jwtService := Infrastructure.NewJWTService(jwtKey)
	passwordSerivce := Infrastructure.NewPasswordService()
	authMiddleware := Infrastructure.NewAuthMiddleware(jwtKey, jwtService)

	// initialize repositories
	userRepo := Repositories.NewUserRepository(userCollection, passwordSerivce)
	taskRepo := Repositories.NewTaskRepository(taskCollection)

	// initialize usecases
	userUsecase := usecases.NewUserUsecase(userRepo, jwtService, passwordSerivce)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// initialize controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	if err != nil {
		log.Fatal(err)
	}

	router := router.SetupRouter(userController, taskController, authMiddleware)
	router.Run(":8080")

}
