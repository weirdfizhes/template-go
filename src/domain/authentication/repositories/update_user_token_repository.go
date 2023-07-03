package repositories

import (
	"template-go/src/domain/authentication/models"
	"template-go/tool/logger"

	"github.com/jmoiron/sqlx"
)

func UpdateUserToken(db *sqlx.DB, request models.UpdateUserTokenPayload) (err error) {
	var guid string
	query := `UPDATE user_tokens SET access_token = $2, refresh_token = $3, updated_at = (now() at time zone 'UTC')::TIMESTAMP WHERE user_guid = $1 RETURNING guid`

	q := db.QueryRow(
		query,
		request.UserGUID,
		request.AccessToken,
		request.RefreshToken,
	)

	err = q.Scan(
		&guid,
	)

	if err != nil {
		logger.LogPrintError("Error update user token", err)
	}

	return
}
