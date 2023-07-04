package models

import (
	"database/sql"

	"time"
)

type GetUserPayload struct {
	ID           int            `json:"id"`
	GUID         string         `json:"guid"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	IsSuperAdmin bool           `json:"is_super_admin"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `json:"created_by"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	UpdatedBy    string         `json:"updated_by"`
	CreatedName  sql.NullString `json:"created_name"`
	UpdatedName  sql.NullString `json:"updated_name"`
}

type readUser struct {
	GUID         string     `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	IsSuperAdmin bool       `json:"is_super_admin"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    string     `json:"created_by"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UpdatedBy    *string    `json:"updated_by"`
}

func ToPayloadUserSingle(user GetUserPayload) (payload readUser) {
	payload = readUser{
		GUID:         user.GUID,
		Name:         user.Name,
		Email:        user.Email,
		IsSuperAdmin: user.IsSuperAdmin,
		CreatedAt:    user.CreatedAt,
	}

	if user.UpdatedAt.Valid {
		payload.UpdatedAt = &user.UpdatedAt.Time
	}

	if user.CreatedName.Valid {
		payload.CreatedBy = user.CreatedName.String
	}

	if user.UpdatedName.Valid {
		payload.UpdatedBy = &user.UpdatedName.String
	}

	return
}

func ToPayloadUserArray(listUser []GetUserPayload) (payload []*readUser) {
	payload = make([]*readUser, len(listUser))

	for i := range listUser {
		payload[i] = new(readUser)
		data := ToPayloadUserSingle(listUser[i])
		payload[i] = &data
	}

	return
}
