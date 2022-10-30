package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Routes :nodoc:
func Routes(route *echo.Echo) {
	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Init HTTP!")
	})

}
