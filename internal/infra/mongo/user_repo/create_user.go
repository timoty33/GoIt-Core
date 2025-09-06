package user_repo

import (
	"context"

	"github.com/timoty33/goit-core/internal/domain/entities"
	e "github.com/timoty33/goit-core/internal/domain/errors"
	"github.com/timoty33/goit-core/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (u *UserRepository) Create(ctx context.Context, user *entities.User) e.GoItError {
	logger.Log(
		zapcore.InfoLevel,
		"Creating user",
		zap.String("locate", "internal/infra/mongo/user_repo/create_user.go"),
	)

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		logger.Log(
			zapcore.ErrorLevel,
			"Error when creating user",
			zap.String("locate", "internal/infra/mongo/user_repo/create_user.go"),
			zap.Error(err),
		)

		return e.New("ERRO DB", "Error when creating user", err)
	}

	logger.Log(
		zapcore.ErrorLevel,
		"User created with success",
		zap.String("locate", "internal/infra/mongo/user_repo/create_user.go"),
	)
	return e.New("OK DB", "Usuário inserido com sucesso", nil)
}
