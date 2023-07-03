package repositories

import (
	"database/sql"
	"template-go/src/domain/user/models"
	"template-go/tool/logger"
)

func DeleteUser(t *sql.Tx, id string) (user models.ReturnDeleteUserPayload, err error) {
	query := `DELETE FROM users WHERE guid=$1 RETURNING guid, name`

	q := t.QueryRow(
		query,
		id,
	)

	err = q.Scan(
		&user.GUID,
		&user.Name,
	)
	if err != nil {
		logger.LogPrintError("Error delete user", err)
	}

	return
}
