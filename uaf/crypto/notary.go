package crypto

type Notary interface {
	Sign(dataToSign string) string

	Verify(dataToSign string, signature string) bool
}
