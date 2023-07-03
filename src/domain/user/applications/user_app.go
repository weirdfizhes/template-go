package applications

import (
	"errors"
	"net/http"
	"template-go/src/domain/user/models"
	"template-go/src/domain/user/services"
	"template-go/src/handlers"
	"template-go/src/middlewares"
	"template-go/tool/constants"
	echohttp "template-go/tool/echo_http"

	"github.com/labstack/echo/v4"
)

func AddRouteUserApp(s *echohttp.Service, e *echo.Group) {
	svc := services.NewUserService(s.GetDB())

	user := e.Group("/user")
	middlewares.SetJwtMiddlewares(user)
	middlewares.ParseJWT(user, s.GetDB())
	user.POST("", createUser(svc))
	user.GET("", getAllUsers(svc))
	user.GET("/:id", getUser(svc))
	user.PUT("/:id", updateUser(svc))
	user.DELETE("/:id", deleteUser(svc))
}

func createUser(svc *services.UserService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request models.CreateUserPayload
		)

		if !middlewares.UserParsed.IsSuperAdmin {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidSA, nil, errors.New("create user must using super admin account"))
		}

		if err := ctx.Bind(&request); err != nil {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrBind, nil, err)
		}

		request = request.ToEntity(middlewares.UserParsed)

		// Validate request
		if err := request.Validate(); err != nil {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidStruct, nil, err)
		}

		user, err := svc.CreateUser(ctx.Request().Context(), request)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error create user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful create user", user, nil)
	}
}

func updateUser(svc *services.UserService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			request models.UpdateUserPayload
			id      string
		)

		if id = ctx.Param("id"); id != "" && !middlewares.UserParsed.IsSuperAdmin {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidSA, nil, errors.New("cannot update other user except super admin"))
		} else {
			id = middlewares.UserParsed.GUID
		}

		if err := ctx.Bind(&request); err != nil {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrBind, nil, err)
		}

		// Validate request
		if err := request.Validate(); err != nil {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidStruct, nil, err)
		}

		user, err := svc.UpdateUser(ctx.Request().Context(), request.ToEntity(id))
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error update user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful update user", user, nil)
	}
}

func deleteUser(svc *services.UserService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			id string
		)

		if !middlewares.UserParsed.IsSuperAdmin {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidSA, nil, errors.New("create user must using super admin account"))
		}

		id = ctx.Param("id")
		if id == "" {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgInvalidParam, nil, errors.New("id cannot be null"))
		}

		user, err := svc.DeleteUser(ctx.Request().Context(), id)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error delete user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful delete user", user, nil)
	}
}

func getAllUsers(svc *services.UserService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		search := ctx.QueryParam("search")

		paginate, err := handlers.PaginationCountHandler(ctx)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error parse pagination int", nil, err)
		}

		if !middlewares.UserParsed.IsSuperAdmin {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidSA, nil, errors.New("create user must using super admin account"))
		}

		count, user, err := svc.GetAllUsers(ctx.Request().Context(), paginate, search)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error get all user", nil, err)
		}

		return echohttp.ResponsePagination(ctx, http.StatusOK, "Successful get all users", models.ToPayloadUserArray(user), nil, paginate.Page, paginate.Limit, count)
	}
}

func getUser(svc *services.UserService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var (
			id string
		)

		if !middlewares.UserParsed.IsSuperAdmin {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgErrValidSA, nil, errors.New("create user must using super admin account"))
		}

		id = ctx.Param("id")
		if id == "" {
			return echohttp.ResponseData(ctx, http.StatusBadRequest, constants.MsgInvalidParam, nil, errors.New("id cannot be null"))
		}

		user, err := svc.GetUser(ctx.Request().Context(), id)
		if err != nil {
			return echohttp.ResponseData(ctx, http.StatusInternalServerError, "Error get user", nil, err)
		}

		return echohttp.ResponseData(ctx, http.StatusOK, "Successful get user", models.ToPayloadUserSingle(user), nil)
	}
}
