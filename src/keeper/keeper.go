package keeper

import (
	"fmt"
	"gremlins/src/registry"
	"net/http"

	"github.com/labstack/echo"
)

func GetGremlins() echo.HandlerFunc {
	return func(c echo.Context) error {
		var gremlins []string
		gremlins, err := registry.GetGremlins()
		if err != nil {
			fmt.Println(err)
			return err
		}
		return c.JSON(http.StatusOK, gremlins)
	}
}
