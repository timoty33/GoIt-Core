package repositories

import (
	"context"

	e "github.com/timoty33/goit-core/internal/domain/errors"

	"github.com/timoty33/goit-core/internal/domain/entities"
)

type UserInterface interface {
	Create(ctx context.Context, user *entities.User) e.GoItError
}
