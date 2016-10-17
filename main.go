package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	rv1 "gitlab.pramati.com/seshachalamm/fido/router/v1"
	"gitlab.pramati.com/seshachalamm/fido/uaf"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, uaf.FIDO_VERSION)
	})

	v := e.Group("/v1/public")

	//Registration
	v.GET("/regRequest/:username", rv1.RegRequest)

	v.POST("/regResponse", rv1.RegResponse)

	//Authentication
	v.GET("/authRequest", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v.POST("/authResponse", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//Deregistration
	v.POST("/deregRequest", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Run(standard.New(":1323"))
}
