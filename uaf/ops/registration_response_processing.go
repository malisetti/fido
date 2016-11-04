package ops

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"

	"github.com/mseshachalam/fido/uaf/crypto"
	"github.com/mseshachalam/fido/uaf/msg"
	"github.com/mseshachalam/fido/uaf/storage"
	"github.com/mseshachalam/fido/uaf/tlv"
)

type RegistrationResponseProcessing struct {
	CertificateValidator crypto.CertificateValidator
}

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

	parser := tlv.TlvAssertionParser{Base64OfRegResponse: "", IsReg: false}
	tags, err := parser.Parse(string(authenticatorRegistrationAssertion.Assertion))
}

func verifyAttestationSignature(tags tlv.Tags, record *storage.RegistrationRecord) error {
	ts := tags.GetTags()
	t, ok := ts[tlv.TAG_ATTESTATION_CERT]
	if !ok {
		return errors.New("exception")
	}

	certBytes := t.Value
	record.AttestCert = base64.URLEncoding.EncodeToString(certBytes)
	krd, ok := ts[tlv.TAG_UAFV1_KRD]
	if !ok {
		return errors.New("exception")
	}
	signature, ok := ts[tlv.TAG_SIGNATURE]
	if !ok {
		return errors.New("exception")
	}

	signedBytes := make([]byte, len(krd.Value)+4)
	bs := []byte(tlv.EncodeInt(krd.ID))
	copy(signedBytes, bs[:2])
	bsl := []byte(tlv.EncodeInt(krd.Length))
	copy(signedBytes, bsl[:2])
	copy(signedBytes, krd.Value[:len(krd.Value)])

	record.AttestDataToSign = base64.URLEncoding.EncodeToString(signedBytes)
	record.AttestSignature = base64.URLEncoding.EncodeToString(signature.Value)
	record.AttestVerifiedStatus = "FAILED_VALIDATION_ATTEMPT"

	if certificateValidator.validate(certBytes, signedBytes,
		signature.value) {
		record.attestVerifiedStatus = "VALID"
	} else {
		record.attestVerifiedStatus = "NOT_VERIFIED"
	}

	return nil
}
