package ops

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"

	"github.com/mseshachalam/fido/uaf/crypto"
	"github.com/mseshachalam/fido/uaf/msg"
	"github.com/mseshachalam/fido/uaf/storage"
)

func Verify(response msg.AuthenticationResponse, serverData storage.StorageInterface) {

}

func ProcessResponse(response msg.RegistrationResponse, serverDataExpiryInMs int, notary crypto.Notary) ([]storage.RegistrationRecord, error) {
	checkAssertions(response)
	records := make([]storage.RegistrationRecord, len(response.Assertions))
	checkVersion(response.Header.UPV)
	checkServerData(response.Header.ServerData, records)
	fcp, err := getFcp(response)

	for i, record := range records {
		records[i] = processAssertions(response.Assertions[i], &record)
	}

	return records, err
}

func checkAssertions(response msg.RegistrationResponse) error {
	if response.Assertions != nil && len(response.Assertions) > 0 {
		return nil
	}

	return errors.New("Missing assertions in registration response")
}

func checkVersion(upv msg.Version) {

}

func checkServerData(serverData string, records []storage.RegistrationRecord) {

}

func getFcp(response msg.RegistrationResponse) (msg.FinalChallengeParams, error) {
	fcParams := msg.FinalChallengeParams{}
	by, err := base64.StdEncoding.DecodeString(string(response.FcParams))
	if err != nil {
		return fcParams, err
	}

	b := bytes.Buffer{}
	b.Write(by)

	d := gob.NewDecoder(&b)
	err = d.Decode(&fcParams)

	return fcParams, err
}

func processAssertions(authenticatorRegistrationAssertion msg.AuthenticatorRegistrationAssertion, record *storage.RegistrationRecord) storage.RegistrationRecord {
	if record == nil {
		record = &storage.RegistrationRecord{}
		record.Status = "INVALID_USERNAME"
	}

	parser := tlv.TlvAssertionParser{}
}
