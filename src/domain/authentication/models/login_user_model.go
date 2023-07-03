package models

import (
	"template-go/tool/constants"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type UserLoginPayload struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type TokenPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (payload *UserLoginPayload) Validate() (err error) {
	// Validate Payload
	if _, err = govalidator.ValidateStruct(payload); err != nil {
		err = errors.Wrapf(constants.ErrBadRequest, "bad request: %s", err.Error())
		return
	}

	return
}
