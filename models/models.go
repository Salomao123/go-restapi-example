package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
