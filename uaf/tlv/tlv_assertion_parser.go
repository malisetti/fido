package tlv

import (
	"encoding/base64"
)

type TLVAssertionParser struct {
}

func (parser *TLVAssertionParser) ParseStr(base64OfRegResponse string) (Tags, error) {
	bytes, err := base64.StdEncoding.DecodeString(base64OfRegResponse)
	isReg := false
	tlv := TLVAssertionParser{}

	return tlv.ParseBytes(bytes, isReg)
}

func (parser *TLVAssertionParser) ParseBytes(bytes []byte, isReg bool) (Tags, error) {
	ret := Tags{}
	t := Tag{}

	return ret, nil
}
