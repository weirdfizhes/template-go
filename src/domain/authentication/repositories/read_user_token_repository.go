package repositories

import (
	"template-go/src/domain/authentication/models"
	"template-go/tool/logger"

	"github.com/jmoiron/sqlx"
)

func GetUserToken(db *sqlx.DB, userGUID string) (ut models.ReadUserTokenPayload, err error) {
	query := `SELECT guid, user_guid, access_token, refresh_token FROM user_tokens WHERE user_guid = $1`

	q := db.QueryRow(
		query,
		userGUID,
	)

	err = q.Scan(
		&ut.GUID,
		&ut.UserGUID,
		&ut.AccessToken,
		&ut.RefreshToken,
	)

	if err != nil {
		logger.LogPrintError("Error get user token", err)
	}

	return
}

func GetUserTokenFromRefreshToken(db *sqlx.DB, refreshToken string) (ut models.ReadUserTokenPayload, err error) {
	query := `SELECT guid, user_guid, access_token, refresh_token FROM user_tokens WHERE refresh_token = $1`

	q := db.QueryRow(
		query,
		refreshToken,
	)

	err = q.Scan(
		&ut.GUID,
		&ut.UserGUID,
		&ut.AccessToken,
		&ut.RefreshToken,
	)

	if err != nil {
		logger.LogPrintError("Error get user token from refresh token", err)
	}

	return
}
