package user_repo

import (
	"github.com/timoty33/goit-core/internal/infra/mongo"

	m "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *m.Collection
}

var UserRepo UserRepository // user repo global

// ==== create userRepo constructor ====
func NewUserRepository(collection *m.Collection) UserRepository {
	return UserRepository{collection: collection}
}

func init() {
	UserRepo = NewUserRepository(mongo.DB.Collection("users"))
}
