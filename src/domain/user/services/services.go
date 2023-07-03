package services

import "github.com/jmoiron/sqlx"

type UserService struct {
	mainDB *sqlx.DB
}

func NewUserService(
	mainDB *sqlx.DB,
) *UserService {
	return &UserService{
		mainDB: mainDB,
	}
}
