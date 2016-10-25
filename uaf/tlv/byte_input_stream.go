package tlv

import ()

type ByteInputStream struct {
	Data []byte
}

func (bsi *ByteInputStream) Read(numberOfBytes int) []byte {
}
