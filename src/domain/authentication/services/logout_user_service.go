package services

import (
	"context"
	"template-go/src/domain/authentication/models"
	"template-go/src/domain/authentication/repositories"
	userModels "template-go/src/domain/user/models"
)

func (s *AuthService) LogoutUser(ctx context.Context, user userModels.GetUserPayload) (err error) {
	err = repositories.UpdateUserToken(s.mainDB, models.UpdateUserTokenPayload{
		UserGUID:     user.GUID,
		AccessToken:  "",
		RefreshToken: "",
	})

	return
}
