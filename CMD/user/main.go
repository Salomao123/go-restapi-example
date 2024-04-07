package main

import (
	"github.com/Salomao123/go-restapi-example/controller"
	"github.com/Salomao123/go-restapi-example/db"
	"github.com/Salomao123/go-restapi-example/repository"
	"github.com/Salomao123/go-restapi-example/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	router := gin.Default()

	collection := handleConnection()

	userRepository := repository.NewUserRepository(collection)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router.GET("/users", userController.FindAllUsers)
	router.POST("/users", userController.CreateUser)

	router.Run(":8000")
}

func handleConnection() *mongo.Collection {
	client, err := db.Connect("mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err.Error())
	}
	database := client.Database("user")
	collection := database.Collection("user")
	return collection
}
