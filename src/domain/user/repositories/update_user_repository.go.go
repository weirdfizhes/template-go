package repositories

import (
	"database/sql"
	"template-go/src/domain/user/models"
	"template-go/tool/logger"
)

func UpdateUser(t *sql.Tx, req models.UpdateUserPayload) (user models.ReturnUpdateUserPayload, err error) {
	query := `UPDATE users SET name = $2, email = $3, updated_at = $4, updated_by = $5 WHERE guid = $1 RETURNING guid, name, email`

	q := t.QueryRow(
		query,
		req.GUID,
		req.Name,
		req.Email,
		req.UpdatedAt,
		req.UpdatedBy,
	)

	err = q.Scan(
		&user.GUID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		logger.LogPrintError("Error update user", err)
		return
	}

	if req.Password != "" {
		query = `UPDATE users SET password = $2, updated_at = $3, updated_by = $4 WHERE guid = $1 RETURNING guid`
		q = t.QueryRow(
			query,
			req.GUID,
			req.Password,
			req.UpdatedAt,
			req.UpdatedBy,
		)

		err = q.Scan(
			&user.GUID,
		)

		logger.LogPrintError("Error update user password", err)
	}

	return
}
