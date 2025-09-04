package user_repo

import (
	"context"

	"github.com/timoty33/goit-core/internal/domain/entities"
	e "github.com/timoty33/goit-core/internal/domain/errors"
)

func (u *UserRepository) Create(ctx context.Context, user *entities.User) e.GoItError {
	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return e.New("ERRO DB", "Erro ao inserir usuário", err)
	}

	return e.New("OK DB", "Usuário inserido com sucesso", nil)
}
