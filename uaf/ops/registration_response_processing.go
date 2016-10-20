package ops

import (
	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/storage"
	"gitlab.pramati.com/seshachalamm/fido/uaf/util"
)

func Verify(response msg.AuthenticationResponse, serverData util.Storage) {

}

func ProcessResponse(resp msg.RegistrationResponse, int serverDataExpiryInMs, notary crypto.Notary) []storage.RegistrationRecord, error {

}
