package tlv

import ()

type ByteInputStream struct {
	Data []byte
}

func (bsi *ByteInputStream) Read(numberOfBytes int) ([]byte, error) {
	return bsi.Data[:numberOfBytes], nil
}

func (bsi *ByteInputStream) ReadAll() ([]byte, error) {

}

func (bsi *ByteInputStream) ReadInteger() (int, error) {

}

func (bsi *ByteInputStream) ReadSigned() (byte, error) {

}

func (bsi *ByteInputStream) ReadUnsigned() (int, error) {

}
