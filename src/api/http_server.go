package api

import (
	"net/http"
	"template-go/src/middlewares"
	echohttp "template-go/tool/echo_http"
	vald "template-go/tool/validator"

	authApp "template-go/src/domain/authentication/applications"
	userApp "template-go/src/domain/user/applications"
	migrateApp "template-go/src/migration/applications"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Routes(s *echohttp.Service) *echo.Echo {
	e := echo.New()

	// validator
	validator := validator.New()
	e.Validator = vald.NewValidator(validator)

	// banner
	e.HideBanner = true

	// handle middleware
	middlewares.LoggerMiddleware(e)
	middlewares.RecoverMiddleware(e)
	middlewares.CorsMiddleware(e)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Backend is running!",
		})
	})

	api := e.Group("/api")

	// domain routes
	migrateApp.Migrate(s, api)
	authApp.AddRouteAuthApp(s, api)
	userApp.AddRouteUserApp(s, api)

	return e
}
