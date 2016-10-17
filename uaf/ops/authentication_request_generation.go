package ops

import (
	"crypto/rand"
	"encoding/base64"
	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
)

func CreateRegistrationRequest(userName string, notary crypto.Notary) msg.RegistrationRequest {
	challenge := generateChallenge()
	serverDataString := generateServerData(username, challenge, notary)

	return createRegistrationRequest(username, serverDataString, challenge)
}

func generateChallenge() string {
	b := make([]byte, 16)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}
