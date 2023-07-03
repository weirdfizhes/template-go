package services

import (
	"context"
	"database/sql"
	"template-go/src/domain/authentication/models"
	"template-go/src/domain/authentication/repositories"
	userRepository "template-go/src/domain/user/repositories"
	"template-go/src/handlers"
	"template-go/tool/hash"
)

func (s *AuthService) LoginUser(ctx context.Context, request models.UserLoginPayload) (token models.TokenPayload, err error) {
	user, err := userRepository.GetUserByEmail(s.mainDB, request.Email)
	if err != nil {
		return
	}

	err = hash.ComparePassword(request.Password, user.Password)
	if err != nil {
		return
	}

	token, err = handlers.CreateJWT(user.GUID, user.CreatedAt)
	if err != nil {
		return
	}

	ut, err := repositories.GetUserToken(s.mainDB, user.GUID)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if ut.GUID != "" {
		err = repositories.UpdateUserToken(s.mainDB, models.UpdateUserTokenPayload{
			UserGUID:     user.GUID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})

		if err != nil {
			return
		}
	} else {
		err = repositories.CreateUserToken(s.mainDB, models.CreateUserTokenPayload{
			UserGUID:     user.GUID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})

		if err != nil {
			return
		}
	}

	return
}
