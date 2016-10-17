package msg

type DeregistrationRequest struct {
	Header         OperationHeader           `json:"header"`
	Authenticators []DeregisterAuthenticator `json:"authenticators"`
}
