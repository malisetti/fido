package tlv

import (
	"bytes"
)

type ByteInputStream struct {
	Data []byte
}

func (bsi *ByteInputStream) Read(numberOfBytes int) []byte {
	readBytes := bytes.Buffer{}
	readBytes.Write(bsi.Data)
}
