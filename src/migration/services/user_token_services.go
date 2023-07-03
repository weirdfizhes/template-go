package services

import (
	"context"
	"database/sql"
	"log"
	"template-go/src/migration/queries"
	"template-go/tool/constants"
)

func (s *MigrateService) UserTokenMigrateUp(ctx context.Context) (err error) {
	t, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			if rollBackErr := t.Rollback(); rollBackErr != nil {
				return
			}
		}
	}()

	err = queries.UserTokenMigrationUp(t)
	if err != nil {
		log.Printf("%s user", constants.MsgErrMigrateUp)
		return
	}

	if err = t.Commit(); err != nil {
		return
	}

	return
}

func (s *MigrateService) UserTokenMigrateDown(ctx context.Context) (err error) {
	t, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			if rollBackErr := t.Rollback(); rollBackErr != nil {
				return
			}
		}
	}()

	err = queries.UserTokenMigrationDown(t)
	if err != nil {
		log.Printf("%s user", constants.MsgErrMigrateUp)
		return
	}

	if err = t.Commit(); err != nil {
		return
	}

	return
}
