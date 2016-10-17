package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

//RegRequest serves registration request
func RegRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//RegResponse serves registration response
func RegResponse(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//AuthRequest serves auth response
func AuthRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//AuthResponse serves auth response
func AuthResponse(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//DeregRequest server deregistration request
func DeregRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
