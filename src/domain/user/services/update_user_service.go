package services

import (
	"context"
	"database/sql"
	"log"
	"template-go/src/domain/user/models"
	"template-go/src/domain/user/repositories"
	"template-go/tool/logger"
)

func (s *UserService) UpdateUser(ctx context.Context, req models.UpdateUserPayload) (user models.ReturnUpdateUserPayload, err error) {
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
		}
		log.Println("Rollbacked!")
	}()

	user, err = repositories.UpdateUser(t, req)
	if err != nil {
		return
	}

	if err = t.Commit(); err != nil {
		logger.LogPrintError("Error commit", err)
		return
	}

	return
}
