package ops

import (
	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/storage"
)

func Verify(response msg.AuthenticationResponse, serverData storage.StorageInterface) {

}

func ProcessResponse(resp msg.RegistrationResponse, int serverDataExpiryInMs, notary crypto.Notary) []storage.RegistrationRecord, error {

}
