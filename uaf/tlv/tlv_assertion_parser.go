package tlv

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
)

func ParseStr(base64OfRegResponse string) (Tags, error) {
	b, err := base64.StdEncoding.DecodeString(base64OfRegResponse)
	if err != nil {
		tags := Tags{}
		return tags, err
	}
	buf := bytes.NewReader(b)
	isReg := false
	tags, err := ParseBytes(buf, isReg)

	return tags, err
}

func ParseBytes(bytes *bytes.Reader, isReg bool) (Tags, error) {
	ret := Tags{}
	var t Tag

	for bytes.Len() > 0 {
		t = Tag{}

		id := []byte{}
		bytes.Read(id)
		t.ID = int(binary.BigEndian.Uint16(id))
		length := bytes.Read
		t.Length = int(binary.BigEndian.Uint16())
		switch t.ID {
		case TAG_UAFV1_AUTH_ASSERTION:
			addTagAndValue(bytes, ret, t)
			addSubTags(isReg, ret, t)

		case TAG_UAFV1_SIGNED_DATA:
			addTagAndValue(bytes, ret, t)
			addSubTags(isReg, ret, t)
		case TAG_UAFV1_REG_ASSERTION:
			isReg := true
			addTagAndValue(bytes, ret, t)
			addSubTags(isReg, ret, t)
		case TAG_UAFV1_KRD:
			ret.Add(t)
			addTagAndValue(bytes, ret, t)
			addSubTags(isReg, ret, t)
		case TAG_AAID:
			addTagAndValue(bytes, ret, t)
		case TAG_ASSERTION_INFO:
			// 2 - Vendor assigned authenticator version.
			// 1 - Authentication Mode indicating whether user
			// explicitly verified or not and indicating if there is a
			// transaction content or not.
			// 2 - Signature algorithm and encoding format.
			if isReg {
				t.Value = bytes.read(7)
			} else {
				t.Value = bytes.read(5)
			}
			ret.Add(t)
		case TAG_AUTHENTICATOR_NONCE:
			addTagAndValue(bytes, ret, t)
		case TAG_FINAL_CHALLENGE:
			addTagAndValue(bytes, ret, t)
		case TAG_TRANSACTION_CONTENT_HASH:
			if t.Length > 0 {
				addTagAndValue(bytes, ret, t)
			} else {
				// Length of Transaction Content Hash. This length is 0
				// if AuthenticationMode == 0x01, i.e. authentication,
				// not transaction confirmation.
				ret.Add(t)
			}
		case TAG_KEYID:
			addTagAndValue(bytes, ret, t)
		case TAG_COUNTERS:
			// Indicates how many times this authenticator has performed
			// signatures in the past
			if isReg {
				t.value = bytes.read(8)
			} else {
				t.value = bytes.read(4)
			}
			ret.add(t)
		case TAG_PUB_KEY:
			addTagAndValue(bytes, ret, t)
		case TAG_ATTESTATION_BASIC_FULL:
			ret.Add(t)
		case TAG_SIGNATURE:
			addTagAndValue(bytes, ret, t)
		case TAG_ATTESTATION_CERT:
			addTagAndValue(bytes, ret, t)
		case TAG_ATTESTATION_BASIC_SURROGATE:
			ret.add(t)
		default:
			t.StatusID = TagsEnum.UAF_CMD_STATUS_ERR_UNKNOWN
			t.Value = bytes.readAll()
			ret.add(t)
		}
	}
}

func addSubTags(isReg bool, ret Tags, t Tag) error {
	//ret.addAll(parse(new ByteInputStream(t.value), isReg));
}

func addTagAndValue(bytes []byte, ret Tags, t Tag) {
	//t.Value = bytes.read(t.length)
	//ret.Add(t)
}
