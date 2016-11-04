package crypto

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"
)

type CertificateValidatorImpl struct {
}

func (cvi CertificateValidatorImpl) Validate(cert string, signedData string, signature string) (bool, error) {
	var err error
	certBytes, err := base64.StdEncoding.DecodeString(cert)
	signedDataBytes, err := base64.StdEncoding.DecodeString(signedData)
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return false, err
	}

	return cvi.ValidateBytes(certBytes, signedDataBytes, signatureBytes)
}

func (cvi CertificateValidatorImpl) ValidateBytes(certBytes []byte, signedDataBytes []byte, signature []byte) (bool, error) {
	x509Certificate, err := x509.ParseDERCRL(certBytes)
	if err != nil {
		return false, err
	}

	sigAlgOID := x509Certificate.SignatureAlgorithm.Algorithm.String()

	if strings.Contains(sigAlgOID, "RSA") {
		h := sha256.New()
		h.Write(signedDataBytes)
		d := h.Sum(nil)

		//do something here

		return err == nil, err
	}
}
