package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"template-go/src/domain/user/models"
	"template-go/src/domain/user/repositories"
	"template-go/tool/logger"
)

func (s *UserService) CreateUser(ctx context.Context, request models.CreateUserPayload) (user models.ReturnCreateUserPayload, err error) {
	t, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			if rollBackErr := t.Rollback(); rollBackErr != nil {
				logger.LogPrintError("Error rollback", rollBackErr)
				return
			}
			log.Println("Rollbacked!")
		}
	}()

	exist, err := repositories.GetUserByEmail(s.mainDB, request.Email)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if exist.Email != "" {
		return user, errors.New("email has been taken")
	}

	user, err = repositories.CreateUser(t, request)
	if err != nil {
		return
	}

	if err = t.Commit(); err != nil {
		logger.LogPrintError("Error commit", err)
		return
	}

	return
}
