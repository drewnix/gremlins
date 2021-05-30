package keeper

import (
	"fmt"
	"gremlins/src/registry"
	"net/http"

	"github.com/labstack/echo"
)

func GetGremlins() echo.HandlerFunc {
	return func(c echo.Context) error {
		if gremlinProvider, err := registry.GetProvider(registry.GremlinService); err == nil {
			fmt.Printf("Gremlin service found at: %v\n", gremlinProvider)
		}
		return c.JSON(http.StatusOK, "gremlins")
	}
}
