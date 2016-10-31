package main

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	rv1 "gitlab.pramati.com/seshachalamm/fido/router/v1"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, msg.FIDO_VERSION)
	})

	v := e.Group("/v1/public")

	//Registration
	v.GET("/regRequest/:username", rv1.RegRequest)

	v.POST("/regResponse", rv1.RegResponse)

	//Authentication
	v.GET("/authRequest", rv1.AuthRequest)

	v.POST("/authResponse", rv1.AuthResponse)

	//Deregistration
	v.POST("/deregRequest", rv1.DeregRequest)

	std := standard.New(":1323")
	std.SetHandler(e)
	gracehttp.Serve(std.Server)
}
