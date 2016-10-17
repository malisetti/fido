package msg

type RegistrationResponse struct {
	Header     OperationHeader                      `json:"header"`
	FcParams   DOMString                            `json:"fcParams"`
	Assertions []AuthenticatorRegistrationAssertion `json:"assertions"`
}
