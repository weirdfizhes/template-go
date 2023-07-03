package repositories

import (
	"template-go/src/domain/authentication/models"
	"template-go/tool/logger"

	"github.com/jmoiron/sqlx"
)

func CreateUserToken(db *sqlx.DB, request models.CreateUserTokenPayload) (err error) {
	var guid string
	query := `INSERT INTO user_tokens (guid, user_guid, access_token, refresh_token, created_at) VALUES (gen_random_uuid(), $1, $2, $3, (now() at time zone 'UTC')::TIMESTAMP) RETURNING guid`

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
		logger.LogPrintError("Error create user token", err)
	}

	return
}
