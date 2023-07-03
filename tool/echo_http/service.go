package echohttp

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
	mainDB *sqlx.DB
}

func NewService(
	mainDB *sqlx.DB,
) *Service {
	return &Service{
		mainDB: mainDB,
	}
}

func (s *Service) GetDB() *sqlx.DB {
	return s.mainDB
}
