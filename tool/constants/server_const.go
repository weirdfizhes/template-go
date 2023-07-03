package constants

import (
	"time"

	"github.com/labstack/echo/v4"
)

const (
	DefaultGraceTimeout       = 10 * time.Second
	DefaultHeaderToken        = echo.HeaderAuthorization
	DefaultHeaderRefreshToken = "Refresh-Token"
)
