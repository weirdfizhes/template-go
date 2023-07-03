package services

import "github.com/jmoiron/sqlx"

type AuthService struct {
	mainDB *sqlx.DB
}

func NewAuthService(
	mainDB *sqlx.DB,
) *AuthService {
	return &AuthService{
		mainDB: mainDB,
	}
}
