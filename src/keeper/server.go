package keeper

import (
	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {
	e.File("/", "public/index.html")
	e.GET("/gremlins", GetGremlins())
}
