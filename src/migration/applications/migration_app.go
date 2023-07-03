package application

import (
	"net/http"
	"os"

	"template-go/src/migration/services"
	"template-go/tool/constants"
	echohttp "template-go/tool/echo_http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func Migrate(s *echohttp.Service, e *echo.Group) {
	svc := services.NewMigrateService(s.GetDB())

	m := e.Group("/migrate")
	m.POST("/up", up(svc))
	m.POST("/down", down(svc))
}

func up(svc *services.MigrateService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		key := echohttp.GetHeaderKeyValue(ctx.Request().Header, os.Getenv("MIGRATION_HEADER"))
		if key != os.Getenv("MIGRATION_KEY") {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgDiffMigrationKey, nil, errors.Wrapf(constants.ErrInvalidToken, constants.MsgBadToken))
		}

		err := svc.UserMigrateUp(ctx.Request().Context())
		if err != nil {
			echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error migrate up user", nil, err)
		}

		err = svc.UserTokenMigrateUp(ctx.Request().Context())
		if err != nil {
			echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error migrate up user token", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful to migrate up", nil, nil)
	}
}

func down(svc *services.MigrateService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		key := echohttp.GetHeaderKeyValue(ctx.Request().Header, os.Getenv("MIGRATION_HEADER"))
		if key != os.Getenv("MIGRATION_KEY") {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgDiffMigrationKey, nil, errors.Wrapf(constants.ErrInvalidToken, constants.MsgBadToken))
		}

		err := svc.UserMigrateDown(ctx.Request().Context())
		if err != nil {
			echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error migrate down user", nil, err)
		}

		err = svc.UserTokenMigrateDown(ctx.Request().Context())
		if err != nil {
			echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error migrate down user token", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful to migrate down", nil, nil)
	}
}
