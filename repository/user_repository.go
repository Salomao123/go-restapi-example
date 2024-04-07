package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Salomao123/go-restapi-example/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection}
}

func (repo *UserRepository) Create(ctx context.Context, user models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	res, err := repo.collection.InsertOne(ctx, user)

	if err != nil {
		log.Fatal(err.Error())
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	fmt.Println(user.ID)
	return nil
}

func (repo *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	cur, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		panic(err.Error())
	}

	var users []models.User

	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, user)
	}

	return users, nil
}
