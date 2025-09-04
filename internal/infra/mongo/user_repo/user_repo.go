package user_repo

import (
	"github.com/timoty33/goit-core/internal/infra/mongo"

	m "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *m.Collection
}

var UserRepo UserRepository

func init() {
	UserRepo = UserRepository{
		collection: mongo.DB.Collection("users"),
	}
}

