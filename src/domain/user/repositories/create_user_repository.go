package repositories

import (
	"database/sql"
	"template-go/src/domain/user/models"
	"template-go/tool/logger"
)

func CreateUser(t *sql.Tx, request models.CreateUserPayload) (user models.ReturnCreateUserPayload, err error) {
	query := `INSERT INTO users (guid, name, email, password, is_super_admin, created_at, created_by) VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6) RETURNING guid, name, email`

	q := t.QueryRow(
		query,
		request.Name,
		request.Email,
		request.Password,
		request.IsSuperAdmin,
		request.CreatedAt,
		request.CreatedBy,
	)

	err = q.Scan(
		&user.GUID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		logger.LogPrintError("Error create user", err)
	}

	return
}
