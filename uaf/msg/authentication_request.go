package msg

type AuthenticationRequest struct {
	Header      OperationHeader `json:"header"`
	Challenge   ServerChallenge `json:"challenge"`
	Transaction []Transaction   `json:"transaction"`
	Policy      Policy          `json:"policy"`
}
