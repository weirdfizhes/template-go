package models

import (
	"template-go/tool/constants"
	"template-go/tool/hash"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type CreateUserPayload struct {
	GUID         string    `json:"guid"`
	Name         string    `json:"name" valid:"required"`
	Email        string    `json:"email" valid:"required"`
	Password     string    `json:"password" valid:"required"`
	IsSuperAdmin bool      `json:"is_super_admin"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
}

type ReturnCreateUserPayload struct {
	GUID  string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (payload *CreateUserPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(constants.ErrBadRequest, "bad request: %s", err.Error())
		return
	}

	return
}

func (payload *CreateUserPayload) ToEntity(user GetUserPayload) (data CreateUserPayload) {
	data = CreateUserPayload{
		Name:         payload.Name,
		Email:        payload.Email,
		IsSuperAdmin: payload.IsSuperAdmin,
		Password:     hash.HashPassword(payload.Password),
		CreatedAt:    time.Now(),
		CreatedBy:    user.GUID,
	}

	return
}
