package msg

type DeregisterAuthenticator struct {
	Aaid  AAID  `json:"aaid"`
	KeyID KeyID `json:"keyID"`
}
