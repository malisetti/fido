package v1

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/util"
)

//RegRequest serves registration request
func RegRequest(c echo.Context) error {
	userName := c.Param("username")
	regReq := [1]msg.RegistrationRequest{}
	fetchRequest := util.FetchRequest{AppID: getAppID(), AllowedAAIDs: getAllowedAaids()}
	regReq[0] = fetchRequest.GetRegistrationRequest(userName)

	return c.JSON(http.StatusOK, regReq)
}

//RegResponse serves registration response
func RegResponse(c echo.Context) error {
	fromJSON := []msg.RegistrationResponse{}
	err := c.Bind(&fromJSON)
	if err != nil {
		return err
	}
	registrationResponse := fromJSON[0]
	result := util.ProcessRegResponse(registrationResponse)

	return c.JSON(http.StatusOK, result)
}

//AuthRequest serves auth response
func AuthRequest(c echo.Context) error {
	ret := [1]msg.AuthenticationRequest{}
	fetchRequest := util.FetchRequest{AppID: getAppID(), AllowedAAIDs: getAllowedAaids()}
	ret[0] = fetchRequest.GetAuthenticationRequest()

	return c.JSON(http.StatusOK, ret)
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
