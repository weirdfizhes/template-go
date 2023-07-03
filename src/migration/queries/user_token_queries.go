package queries

import (
	"database/sql"
)

// Query Table Up
func createUserTokenSequence(t *sql.Tx) (err error) {
	query := `
		CREATE SEQUENCE IF NOT EXISTS user_tokens_id_seq
    		INCREMENT 1
    		START 1
    		MINVALUE 1
    		MAXVALUE 9223372036854775807
    		CACHE 1;
	`

	_, err = t.Exec(query)

	return
}

func createUserTokenTable(t *sql.Tx) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS user_tokens
		(
			id bigint NOT NULL DEFAULT nextval('user_tokens_id_seq'::regclass),
			guid CHARACTER VARYING NOT NULL PRIMARY KEY,
			user_guid CHARACTER VARYING NOT NULL,
			access_token TEXT NOT NULL,
			refresh_token TEXT NOT NULL,
			created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp without time zone
		);
	`

	_, err = t.Exec(query)

	return
}

// Query Table Down
func dropUserTokenTable(t *sql.Tx) (err error) {
	query := `
		DROP TABLE IF EXISTS users;
	`
	_, err = t.Exec(query)

	return
}

func dropUserTokenSequence(t *sql.Tx) (err error) {
	query := `
		DROP SEQUENCE IF EXISTS users_seq;
	`

	_, err = t.Exec(query)

	return
}

// Migrate Up
func UserTokenMigrationUp(t *sql.Tx) (err error) {
	err = createUserTokenSequence(t)
	if err != nil {
		return
	}

	err = createUserTokenTable(t)
	if err != nil {
		return
	}

	return
}

// Migrate Down
func UserTokenMigrationDown(t *sql.Tx) (err error) {
	err = dropUserTokenTable(t)
	if err != nil {
		return
	}

	err = dropUserTokenSequence(t)
	if err != nil {
		return
	}

	return
}
