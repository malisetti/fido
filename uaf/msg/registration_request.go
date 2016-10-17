package msg

type RegistrationRequest struct {
	Header    OperationHeader `json:"header"`
	Challenge ServerChallenge `json:"challenge"`
	Username  DOMString       `json:"username"`
	Policy    Policy          `json:"policy"`
}
