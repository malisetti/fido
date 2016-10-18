package v1

import (
	"github.com/labstack/echo"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/ops"
	"gitlab.pramati.com/seshachalamm/fido/uaf/util"
	"net/http"
)

//RegRequest serves registration request
func RegRequest(c echo.Context) error {

	userName := c.Param("username")

	regReq := [1]msg.RegistrationRequest{}

	regReq[0] = getRegistrationRequest(userName, getAppID(), getAllowedAaids())

	return c.JSON(200, regReq)
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

func getAppID() string {
	return "v1/public/uaf/facets"
}

func getAllowedAaids() []string {
	allowedAaids := []string{"EBA0#0001", "0015#0001", "0012#0002", "0010#0001",
		"4e4e#0001", "5143#0001", "0011#0701", "0013#0001",
		"0014#0000", "0014#0001", "53EC#C002", "DAB8#8001",
		"DAB8#0011", "DAB8#8011", "5143#0111", "5143#0120",
		"4746#F816", "53EC#3801"}

	return allowedAaids
}

func getRegistrationRequest(username, appID string, acceptedAaids []string) msg.RegistrationRequest {

	notary := util.NotaryImpl{}
	return ops.CreateRegistrationRequest(username, appID, acceptedAaids, notary)
}
