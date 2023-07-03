package queries

import (
	"database/sql"
	"template-go/tool/uuid"
	"time"
)

// Query Table Up
func createUserSequence(t *sql.Tx) (err error) {
	query := `
		CREATE SEQUENCE IF NOT EXISTS users_seq
    		INCREMENT 1
    		START 1
    		MINVALUE 1
    		MAXVALUE 9223372036854775807
    		CACHE 1;
	`

	_, err = t.Exec(query)

	return
}

func createUserTable(t *sql.Tx) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS users
		(
			id bigint NOT NULL DEFAULT nextval('users_seq'::regclass),
			guid CHARACTER VARYING NOT NULL PRIMARY KEY,
			name CHARACTER VARYING NOT NULL,
			email CHARACTER VARYING UNIQUE NOT NULL,
			password CHARACTER VARYING NOT NULL,
			is_super_admin BOOLEAN NOT NULL DEFAULT FALSE,
			created_at timestamp without time zone NOT NULL,
			created_by CHARACTER VARYING NOT NULL,
			updated_at timestamp without time zone,
			updated_by CHARACTER VARYING
		);
	`

	_, err = t.Exec(query)

	return
}

// Query Table Down
func dropUserTable(t *sql.Tx) (err error) {
	query := `
		DROP TABLE IF EXISTS users;
	`
	_, err = t.Exec(query)

	return
}

func dropUserSequence(t *sql.Tx) (err error) {
	query := `
		DROP SEQUENCE IF EXISTS users_seq;
	`

	_, err = t.Exec(query)

	return
}

// Query Seed User Table
func seedUserTable(t *sql.Tx) (err error) {
	query := `
		INSERT INTO users (guid, name, email, password, is_super_admin, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING;
	`

	_, err = t.Exec(query, uuid.GenerateGUID(), "Admin", "admin@gmail.com", "$2a$10$89l83wbZ.X9JnAenCPYBBOCWbmC14LzY1pnRai7sNa8XRbZDQrL2C", true, time.Now(), "system")

	return
}

// Migrate Up
func UserMigrationUp(t *sql.Tx) (err error) {
	err = createUserSequence(t)
	if err != nil {
		return
	}

	err = createUserTable(t)
	if err != nil {
		return
	}

	err = seedUserTable(t)
	if err != nil {
		return
	}

	return
}

// Migrate Down
func UserMigrationDown(t *sql.Tx) (err error) {
	err = dropUserTable(t)
	if err != nil {
		return
	}

	err = dropUserSequence(t)
	if err != nil {
		return
	}

	return
}
