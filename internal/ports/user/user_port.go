package user

import (
	"context"

	"github.com/timoty33/goit-core/internal/domain/entities"
	e "github.com/timoty33/goit-core/internal/domain/errors"
	ur "github.com/timoty33/goit-core/internal/infra/mongo/user_repo"
)

type UserPortStruct struct {
	repo ur.UserRepository
}

var UserPort UserPortStruct

func init() {
	UserPort = UserPortStruct{
		repo: ur.UserRepo,
	}
}

func (u *UserPortStruct) Create(ctx context.Context, user *entities.User) e.GoItError {
	return u.repo.Create(ctx, user)
}
