package tlv

func EncodeInt(id int) []byte {
	bytes := []byte{}
	bytes[0] = (byte)(id & 0x00ff)
	bytes[1] = (byte)((id & 0xff00) >> 8)

	return bytes
}
