package services

import (
	"context"
	"database/sql"
	"log"
	"template-go/src/migration/queries"
	"template-go/tool/constants"
)

func (s *MigrateService) UserMigrateUp(ctx context.Context) (err error) {
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

	err = queries.UserMigrationUp(t)
	if err != nil {
		log.Printf("%s user", constants.MsgErrMigrateUp)
		return
	}

	if err = t.Commit(); err != nil {
		return
	}

	return
}

func (s *MigrateService) UserMigrateDown(ctx context.Context) (err error) {
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

	err = queries.UserMigrationDown(t)
	if err != nil {
		log.Printf("%s user", constants.MsgErrMigrateUp)
		return
	}

	if err = t.Commit(); err != nil {
		return
	}

	return
}
