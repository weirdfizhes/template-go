package models

import (
	"template-go/tool/constants"
	"template-go/tool/hash"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type UpdateUserPayload struct {
	GUID         string    `json:"id"`
	Name         string    `json:"name" valid:"required"`
	Email        string    `json:"email" valid:"required"`
	Password     string    `json:"password"`
	IsSuperAdmin bool      `json:"is_super_admin"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
}

type ReturnUpdateUserPayload struct {
	GUID  string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (payload *UpdateUserPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(constants.ErrBadRequest, "bad request: %s", err.Error())
		return
	}

	return
}

func (payload *UpdateUserPayload) ToEntity(id string) (data UpdateUserPayload) {
	data = UpdateUserPayload{
		GUID:         id,
		Name:         payload.Name,
		Email:        payload.Email,
		IsSuperAdmin: payload.IsSuperAdmin,
		UpdatedAt:    time.Now(),
		UpdatedBy:    payload.UpdatedBy,
	}

	if payload.Password != "" {
		data.Password = hash.HashPassword(payload.Password)
	}

	return
}
