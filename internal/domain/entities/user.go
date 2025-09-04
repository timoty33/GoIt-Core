package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Photo     string             `bson:"photo"`
}
