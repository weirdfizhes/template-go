package repositories

import (
	"fmt"
	"template-go/src/domain/user/models"
	"template-go/src/handlers"
	"template-go/tool/logger"

	"github.com/jmoiron/sqlx"
)

func GetUserByEmail(db *sqlx.DB, email string) (user models.GetUserPayload, err error) {
	query := `SELECT guid, name, email, password, created_at FROM users WHERE email=$1`

	q := db.QueryRow(
		query,
		email,
	)

	err = q.Scan(
		&user.GUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		logger.LogPrintError("Error get user by email", err)
	}

	return
}

func GetAllUsers(db *sqlx.DB, paginate handlers.Pagination, search string) (user []models.GetUserPayload, err error) {
	var (
		userSingle models.GetUserPayload
		likeQuery  string
		pagination string
	)

	if search != "" {
		likeQuery = fmt.Sprintf(`WHERE LOWER(users.name) LIKE LOWER('%s') OR LOWER(users.email) LIKE LOWER('%s')`, "%"+search+"%", "%"+search+"%")
	}

	if paginate.Page != 0 && paginate.Limit != 0 {
		pagination = fmt.Sprintf(`LIMIT %v OFFSET %v`, paginate.Limit, paginate.Offset)
	}

	query := fmt.Sprintf(`SELECT users.guid, users.name, users.email, users.password, users.created_at, users.is_super_admin, users.created_at, users.updated_at, users_created.name AS created_name, users_updated.name AS updated_name FROM users LEFT JOIN users users_created ON users.created_by=users_created.guid LEFT JOIN users users_updated ON users.updated_by=users_updated.guid %s ORDER BY users.created_at ASC %s`, likeQuery, pagination)

	q, err := db.Queryx(
		query,
	)
	if err != nil {
		logger.LogPrintError("Error get all user", err)
		return
	}

	for q.Next() {
		err = q.Scan(
			&userSingle.GUID,
			&userSingle.Name,
			&userSingle.Email,
			&userSingle.Password,
			&userSingle.CreatedAt,
			&userSingle.IsSuperAdmin,
			&userSingle.CreatedAt,
			&userSingle.UpdatedAt,
			&userSingle.CreatedName,
			&userSingle.UpdatedName,
		)

		if err != nil {
			logger.LogPrintError("Error scan db value for get all user", err)
			return
		}

		user = append(user, userSingle)
	}

	return
}

func GetUserById(db *sqlx.DB, id string) (user models.GetUserPayload, err error) {
	query := `SELECT users.guid, users.name, users.email, users.password, users.created_at, users.is_super_admin, users.created_at, users.updated_at, users_created.name AS created_name, users_updated.name AS updated_name FROM users LEFT JOIN users users_created ON users.created_by=users_created.guid LEFT JOIN users users_updated ON users.updated_by=users_updated.guid WHERE users.guid = $1`

	q, err := db.Queryx(
		query,
		id,
	)
	if err != nil {
		logger.LogPrintError("Error get user by id", err)
		return
	}

	for q.Next() {
		err = q.Scan(
			&user.GUID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.IsSuperAdmin,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedName,
			&user.UpdatedName,
		)

		if err != nil {
			logger.LogPrintError("Error scan db value for get user by id", err)
			return
		}
	}

	return
}

func GetCountUser(db *sqlx.DB, search string) (count int64, err error) {
	var (
		likeQuery string
	)

	if search != "" {
		likeQuery = fmt.Sprintf(`WHERE LOWER(users.name) LIKE LOWER('%s') OR LOWER(users.email) LIKE LOWER('%s')`, "%"+search+"%", "%"+search+"%")
	}

	query := fmt.Sprintf(`SELECT count(id) FROM users %s`, likeQuery)

	q, err := db.Queryx(
		query,
	)
	if err != nil {
		return
	}

	for q.Next() {
		err = q.Scan(
			&count,
		)

		if err != nil {
			return
		}
	}

	return
}
