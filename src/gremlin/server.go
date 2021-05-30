package gremlin

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/", hello)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
