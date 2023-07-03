package handlers

import (
	"os"
	"template-go/src/domain/authentication/models"
	"template-go/tool/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtClaims struct {
	Created_at time.Time `json:"created_at"`

	jwt.StandardClaims
}

func CreateJWT(guid string, created_at time.Time) (token models.TokenPayload, err error) {
	var (
		secret []byte    = []byte(os.Getenv("JWT_SECRET"))
		claims jwtClaims = jwtClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
				Issuer:    guid,
			},
			Created_at: created_at,
		}
	)

	// Create Access Token
	rawAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token.AccessToken, err = rawAccessToken.SignedString(secret)
	if err != nil {
		logger.LogPrintError("Error signed access token", err)
		return
	}

	// Create Refresh Token
	rawRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    guid,
	})
	token.RefreshToken, err = rawRefreshToken.SignedString(secret)
	if err != nil {
		logger.LogPrintError("Error signed refresh token", err)
		return
	}
	return
}
