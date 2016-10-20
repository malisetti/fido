package storage

const DLM = "#"

type AuthenticatorRecord struct {
	AAID     string
	KeyID    string
	DeviceID string
	Username string
	Status   string
}

func (ar *AuthenticatorRecord) String() string {
	return ar.AAID + DLM + ar.KeyID
}
