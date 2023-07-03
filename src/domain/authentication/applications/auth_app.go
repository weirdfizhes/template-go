package applications

import (
	"net/http"
	"template-go/src/domain/authentication/models"
	"template-go/src/domain/authentication/services"
	"template-go/src/middlewares"
	"template-go/tool/constants"
	echohttp "template-go/tool/echo_http"

	"github.com/labstack/echo/v4"
)

func AddRouteAuthApp(s *echohttp.Service, e *echo.Group) {
	svc := services.NewAuthService(s.GetDB())

	auth := e.Group("/auth")
	auth.POST("/login", loginUser(svc))

	authLogout := auth.Group("")
	middlewares.SetJwtMiddlewares(authLogout)
	middlewares.ParseJWT(authLogout, s.GetDB())
	authLogout.POST("/logout", logoutUser(svc))

	authRefresh := auth.Group("")
	middlewares.ValidateRefreshToken(authRefresh)
	authRefresh.POST("/refresh", refreshToken(svc))
}

func loginUser(svc *services.AuthService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request models.UserLoginPayload
		)

		if err := ctx.Bind(&request); err != nil {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrBind, nil, err)
		}

		// Validate request
		if err := request.Validate(); err != nil {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidStruct, nil, err)
		}

		res, err := svc.LoginUser(ctx.Request().Context(), request)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error login user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful login user", res, nil)
	}
}

func logoutUser(svc *services.AuthService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := svc.LogoutUser(ctx.Request().Context(), middlewares.UserParsed)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error logout user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful logout user", nil, nil)
	}
}

func refreshToken(svc *services.AuthService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := svc.RefreshTokenUser(ctx.Request().Context(), middlewares.ResJWT.Raw)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error refresh token user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful refresh token user", res, nil)
	}
}
