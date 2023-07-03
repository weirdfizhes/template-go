package services

import (
	"context"
	"database/sql"
	"template-go/src/domain/authentication/models"
	"template-go/src/domain/authentication/repositories"
	userRepository "template-go/src/domain/user/repositories"
	"template-go/src/handlers"
)

func (s *AuthService) RefreshTokenUser(ctx context.Context, refreshToken string) (token models.TokenPayload, err error) {
	ut, err := repositories.GetUserTokenFromRefreshToken(s.mainDB, refreshToken)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	user, err := userRepository.GetUserById(s.mainDB, ut.UserGUID)
	if err != nil {
		return
	}

	token, err = handlers.CreateJWT(user.GUID, user.CreatedAt)
	if err != nil {
		return
	}

	if ut.GUID != "" {
		err = repositories.UpdateUserToken(s.mainDB, models.UpdateUserTokenPayload{
			UserGUID:     ut.UserGUID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})

		if err != nil {
			return
		}
	} else {
		err = repositories.CreateUserToken(s.mainDB, models.CreateUserTokenPayload{
			UserGUID:     ut.UserGUID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})

		if err != nil {
			return
		}
	}

	return
}
