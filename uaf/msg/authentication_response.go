package msg

type AuthenticationResponse struct {
	Header     OperationHeader              `json:"header"`
	FcParams   DOMString                    `json:"fcParams"`
	Assertions []AuthenticatorSignAssertion `json:"assertions"`
}
