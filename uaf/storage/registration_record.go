package storage

type RegistrationRecord struct {
	Authenticator               AuthenticatorRecord
	PublicKey                   string
	SignCounter                 string
	AuthenticatorVersion        string
	TcDisplayPNGCharacteristics string
	Username                    string
	UserID                      string
	DeviceID                    string
	TimeStamp                   string
	Status                      string
	AttestCert                  string
	AttestDataToSign            string
	AttestSignature             string
	AttestVerifiedStatus        string
}
