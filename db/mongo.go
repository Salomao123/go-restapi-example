package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to MongoDB!")
	return client, err
}
