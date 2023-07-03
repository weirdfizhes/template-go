package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"template-go/tool/constants"
	echohttp "template-go/tool/echo_http"
	"template-go/tool/logger"

	"template-go/src/domain/user/models"
	userModels "template-go/src/domain/user/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

var (
	UserParsed userModels.GetUserPayload
	ResJWT     *jwt.Token
)

// SetJwtMiddlewares is a middleware to handle jwt validation
func SetJwtMiddlewares(g *echo.Group) {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS512.Name,
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
	}))
	middleware.ErrJWTMissing.Code = http.StatusUnauthorized
	middleware.ErrJWTMissing.Message = "You're unauthorized"
}

// ParseJWT is a middleware function to parsing issuer data
func ParseJWT(g *echo.Group, db *sqlx.DB) {
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get header token
			authHeader := c.Request().Header.Get(constants.DefaultHeaderToken)
			token := strings.Fields(authHeader)
			if len(token) < 2 && token[1] == "" {
				return echohttp.ResponseData(c, http.StatusUnauthorized, "Error get header token", nil, errors.New("header token not found"))
			}

			// Parse jwt decode
			claims := jwt.MapClaims{}
			res, err := jwt.ParseWithClaims(token[1], claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				logger.LogPrintError("Error parse jwt token", err)
				return echohttp.ResponseData(c, http.StatusUnauthorized, "Error parse jwt token", nil, err)
			}

			// Check decoded data if it's valid or not
			resMap, ok := res.Claims.(jwt.MapClaims)
			if !ok && !res.Valid {
				err = errors.New("error validate jwt parse")
				logger.LogPrintError("Invalid token jwt", err)
				return echohttp.ResponseData(c, http.StatusUnauthorized, "Error parse jwt token", nil, err)
			}

			// Sync decoded data with database
			if issuer := resMap["iss"]; issuer != nil {
				user, err := getUser(db, issuer.(string))
				if err != nil {
					return echohttp.ResponseData(c, http.StatusUnauthorized, "Error middleware get user by id", nil, err)
				}

				if createdAt := resMap["created_at"]; createdAt != nil {
					t, err := time.Parse(time.RFC3339, createdAt.(string))
					if err != nil {
						return echohttp.ResponseData(c, http.StatusUnauthorized, "Error middleware parse created at", nil, err)
					}

					if user.CreatedAt.UTC() != t {
						return echohttp.ResponseData(c, http.StatusUnauthorized, "Created at didn't same with user created at data", nil, errors.New("created at is different"))
					}
				} else {
					return echohttp.ResponseData(c, http.StatusUnauthorized, "Created at didn't same with user created at data", nil, errors.New("created at is null"))
				}

				fmt.Println(user)

				UserParsed.GUID = user.GUID
				UserParsed.IsSuperAdmin = user.IsSuperAdmin
			}

			return next(c)
		}
	})
}

// ValidateRefreshToken is a middleware function to
func ValidateRefreshToken(g *echo.Group) {
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// Get header token
			refreshHeader := c.Request().Header.Get(constants.DefaultHeaderRefreshToken)

			// Parse jwt decode
			claims := jwt.MapClaims{}
			ResJWT, err = jwt.ParseWithClaims(refreshHeader, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				logger.LogPrintError("Error validate refresh token", err)
				return echohttp.ResponseData(c, http.StatusUnauthorized, "Error validate refresh token", nil, errors.Errorf("refresh token is invalid"))
			}

			return next(c)
		}
	})
}

func getUser(db *sqlx.DB, id string) (user models.GetUserPayload, err error) {
	query := `SELECT guid, is_super_admin, created_at FROM users WHERE guid = $1`

	q, err := db.Queryx(
		query,
		id,
	)
	if err != nil {
		logger.LogPrintError("Error get user by id", err)
		return
	}

	for q.Next() {
		err = q.Scan(
			&user.GUID,
			&user.IsSuperAdmin,
			&user.CreatedAt,
		)

		if err != nil {
			logger.LogPrintError("Error scan db value for get user by id", err)
			return
		}
	}

	return
}
