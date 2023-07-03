package services

import "github.com/jmoiron/sqlx"

type MigrateService struct {
	mainDB *sqlx.DB
}

func NewMigrateService(
	mainDB *sqlx.DB,
) *MigrateService {
	return &MigrateService{
		mainDB: mainDB,
	}
}
