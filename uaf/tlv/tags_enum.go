package tlv

import "errors"

const (
	UAF_CMD_STATUS_ERR_UNKNOWN      = 0x01
	TAG_UAFV1_REG_ASSERTION         = 0x3E01
	TAG_UAFV1_AUTH_ASSERTION        = 0x3E02
	TAG_UAFV1_KRD                   = 0x3E03
	TAG_UAFV1_SIGNED_DATA           = 0x3E04
	TAG_ATTESTATION_CERT            = 0x2E05
	TAG_SIGNATURE                   = 0x2E06
	TAG_ATTESTATION_BASIC_FULL      = 0x3E07
	TAG_ATTESTATION_BASIC_SURROGATE = 0x3E08
	TAG_KEYID                       = 0x2E09
	TAG_FINAL_CHALLENGE             = 0x2E0A
	TAG_AAID                        = 0x2E0B
	TAG_PUB_KEY                     = 0x2E0C
	TAG_COUNTERS                    = 0x2E0D
	TAG_ASSERTION_INFO              = 0x2E0E
	TAG_AUTHENTICATOR_NONCE         = 0x2E0F
	TAG_TRANSACTION_CONTENT_HASH    = 0x2E10
	TAG_EXTENSION                   = 0x3E11
	TAG_EXTENSION_NON_CRITICAL      = 0x3E12
	TAG_EXTENSION_ID                = 0x2E13
	TAG_EXTENSION_DATA              = 0x2E14
)

var (
	names = map[int]string{
		TAG_UAFV1_REG_ASSERTION:         "TAG_UAFV1_REG_ASSERTION",
		TAG_UAFV1_AUTH_ASSERTION:        "TAG_UAFV1_AUTH_ASSERTION",
		TAG_UAFV1_KRD:                   "TAG_UAFV1_KRD",
		TAG_UAFV1_SIGNED_DATA:           "TAG_UAFV1_SIGNED_DATA",
		TAG_ATTESTATION_BASIC_FULL:      "TAG_ATTESTATION_BASIC_FULL",
		TAG_ATTESTATION_BASIC_SURROGATE: "TAG_ATTESTATION_BASIC_SURROGATE",
		TAG_ATTESTATION_CERT:            "TAG_ATTESTATION_CERT",
		TAG_SIGNATURE:                   "TAG_SIGNATURE",
		TAG_KEYID:                       "TAG_KEYID",
		TAG_FINAL_CHALLENGE:             "TAG_FINAL_CHALLENGE",
		TAG_AAID:                        "TAG_AAID",
		TAG_PUB_KEY:                     "TAG_PUB_KEY",
		TAG_COUNTERS:                    "TAG_COUNTERS",
		TAG_ASSERTION_INFO:              "TAG_ASSERTION_INFO",
		TAG_AUTHENTICATOR_NONCE:         "TAG_AUTHENTICATOR_NONCE",
		TAG_TRANSACTION_CONTENT_HASH:    "TAG_TRANSACTION_CONTENT_HASH",
	}
)

type TagsEnum struct {
}

func values() []int {
	return []int{
		UAF_CMD_STATUS_ERR_UNKNOWN,
		TAG_UAFV1_REG_ASSERTION,
		TAG_UAFV1_AUTH_ASSERTION,
		TAG_UAFV1_KRD,
		TAG_UAFV1_SIGNED_DATA,
		TAG_ATTESTATION_CERT,
		TAG_SIGNATURE,
		TAG_ATTESTATION_BASIC_FULL,
		TAG_ATTESTATION_BASIC_SURROGATE,
		TAG_KEYID,
		TAG_FINAL_CHALLENGE,
		TAG_AAID,
		TAG_PUB_KEY,
		TAG_COUNTERS,
		TAG_ASSERTION_INFO,
		TAG_AUTHENTICATOR_NONCE,
		TAG_TRANSACTION_CONTENT_HASH,
		TAG_EXTENSION,
		TAG_EXTENSION_NON_CRITICAL,
		TAG_EXTENSION_ID,
		TAG_EXTENSION_DATA,
	}
}

func (tagsEnum *TagsEnum) Get(id int) (int, error) {
	for _, tag := range values() {
		if tag == id {
			return tag, nil
		}
	}

	return 0, errors.New("Tag not found.")
}

func (tagsEnum *TagsEnum) GetName(id int) string {
	if val, ok := names[id]; ok {
		return val
	}
	return "TAG_UNKNOWN"
}
