package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

const hmacSecret = "HMAC-is-just-one-way"

type NotaryImpl struct {
}

func (n NotaryImpl) Sign(dataToSign string) string {
	key := []byte(hmacSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(dataToSign))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (n NotaryImpl) Verify(dataToSign string, signature string) bool {
	key := []byte(hmacSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(dataToSign))

	b, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}

	return hmac.Equal(b, h.Sum(nil))
}
