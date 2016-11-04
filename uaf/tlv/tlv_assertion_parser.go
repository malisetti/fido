package tlv

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
)

type TlvAssertionParser struct {
	Base64OfRegResponse string
	IsReg               bool
}

func (tlvAssertParser *TlvAssertionParser) Parse(base64OfRegResponse string) (Tags, error) {
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
	var err error
	var t Tag

	for bytes.Len() > 0 {
		t = Tag{}

		id := []byte{}
		bytes.Read(id)
		t.ID = int(binary.BigEndian.Uint16(id))
		length := binary.BigEndian.Uint16(id)
		t.Length = int(length)
		switch t.ID {
		case TAG_UAFV1_AUTH_ASSERTION:
			addTagAndValue(bytes, ret, t)
			err = addSubTags(isReg, ret, t)

		case TAG_UAFV1_SIGNED_DATA:
			addTagAndValue(bytes, ret, t)
			err = addSubTags(isReg, ret, t)
		case TAG_UAFV1_REG_ASSERTION:
			isReg := true
			addTagAndValue(bytes, ret, t)
			err = addSubTags(isReg, ret, t)
		case TAG_UAFV1_KRD:
			ret.Add(t)
			addTagAndValue(bytes, ret, t)
			err = addSubTags(isReg, ret, t)
		case TAG_AAID:
			addTagAndValue(bytes, ret, t)
		case TAG_ASSERTION_INFO:
			// 2 - Vendor assigned authenticator version.
			// 1 - Authentication Mode indicating whether user
			// explicitly verified or not and indicating if there is a
			// transaction content or not.
			// 2 - Signature algorithm and encoding format.
			if isReg {
				val := make([]byte, 7)
				_, err = bytes.Read(val)
				t.Value = val
			} else {
				val := make([]byte, 5)
				_, err = bytes.Read(val)
				t.Value = val
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
				val := make([]byte, 8)
				_, err = bytes.Read(val)
				t.Value = val
			} else {
				val := make([]byte, 4)
				_, err = bytes.Read(val)
				t.Value = val
			}
			ret.Add(t)
		case TAG_PUB_KEY:
			addTagAndValue(bytes, ret, t)
		case TAG_ATTESTATION_BASIC_FULL:
			ret.Add(t)
		case TAG_SIGNATURE:
			addTagAndValue(bytes, ret, t)
		case TAG_ATTESTATION_CERT:
			addTagAndValue(bytes, ret, t)
		case TAG_ATTESTATION_BASIC_SURROGATE:
			ret.Add(t)
		default:
			t.StatusID = UAF_CMD_STATUS_ERR_UNKNOWN
			val := []byte{}
			bytes.Read(val)
			t.Value = val
			ret.Add(t)
		}
	}

	return ret, err
}

func addSubTags(isReg bool, ret Tags, t Tag) error {
	buf := bytes.NewReader(t.Value)

	tags, err := ParseBytes(buf, isReg)

	if err != nil {
		return err
	}

	ret.AddAll(tags)

	return nil
}

func addTagAndValue(bytes *bytes.Reader, ret Tags, t Tag) {
	val := make([]byte, t.Length)
	bytes.Read(val)
	t.Value = val

	ret.Add(t)
}
