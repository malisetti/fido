package tlv

func ReadUAFV1_UINT16(bytes []byte) (int, error) {
	a := bytes.readUnsignedByte()
	b := bytes.readUnsignedByte()

	return a + b*256
}

func EncodeInt(id int) []byte {
	bytes := []byte{}
	bytes[0] = (byte)(id & 0x00ff)
	bytes[1] = (byte)((id & 0xff00) >> 8)

	return bytes
}
