package crypto

type CertificateValidator interface {
	Validate(cert string, signedData string, signature string) (bool, error)

	ValidateBytes(certBytes []byte, signedDataBytes []byte, signature []byte) (bool, error)
}
