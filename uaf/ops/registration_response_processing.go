package ops

import (
	"errors"

	"gitlab.pramati.com/seshachalamm/fido/uaf/crypto"
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/storage"
)

func Verify(response msg.AuthenticationResponse, serverData storage.StorageInterface) {

}

func ProcessResponse(resp msg.RegistrationResponse, serverDataExpiryInMs int, notary crypto.Notary) ([]storage.RegistrationRecord, error) {
	return []storage.RegistrationRecord{}, errors.New("")
}
