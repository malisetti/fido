package uaf

//defines all the constants and types for uaf
const (
	FIDO_VERSION = "v1.0"

	Reg   = "Reg"
	Auth  = "Auth"
	Dereg = "Dereg"

	//Key Protection Types
	KEY_PROTECTION_SOFTWARE       = 0x01
	KEY_PROTECTION_HARDWARE       = 0x02
	KEY_PROTECTION_TEE            = 0x04
	KEY_PROTECTION_SECURE_ELEMENT = 0x08
	KEY_PROTECTION_REMOTE_HANDLE  = 0x10

	//User Verification Methods
	USER_VERIFY_PRESENCE    = 0x01
	USER_VERIFY_FINGERPRINT = 0x02
	USER_VERIFY_PASSCODE    = 0x04
	USER_VERIFY_VOICEPRINT  = 0x08
	USER_VERIFY_FACEPRINT   = 0x10
	USER_VERIFY_LOCATION    = 0x20
	USER_VERIFY_EYEPRINT    = 0x40
	USER_VERIFY_PATTERN     = 0x80
	USER_VERIFY_HANDPRINT   = 0x100
	USER_VERIFY_NONE        = 0x200
	USER_VERIFY_ALL         = 0x400

	//Matcher Protection Types
	MATCHER_PROTECTION_SOFTWARE = 0x01
	MATCHER_PROTECTION_TEE      = 0x02
	MATCHER_PROTECTION_ON_CHIP  = 0x04

	//Authenticator Attachment Hints
	ATTACHMENT_HINT_INTERNAL    = 0x01
	ATTACHMENT_HINT_EXTERNAL    = 0x02
	ATTACHMENT_HINT_WIRED       = 0x04
	ATTACHMENT_HINT_WIRELESS    = 0x08
	ATTACHMENT_HINT_NFC         = 0x10
	ATTACHMENT_HINT_BLUETOOTH   = 0x20
	ATTACHMENT_HINT_NETWORK     = 0x40
	ATTACHMENT_HINT_READY       = 0x80
	ATTACHMENT_HINT_WIFI_DIRECT = 0x100

	//Transaction Confirmation Display Types
	TRANSACTION_CONFIRMATION_DISPLAY_ANY                 = 0x01
	TRANSACTION_CONFIRMATION_DISPLAY_PRIVILEGED_SOFTWARE = 0x02
	TRANSACTION_CONFIRMATION_DISPLAY_TEE                 = 0x04
	TRANSACTION_CONFIRMATION_DISPLAY_HARDWARE            = 0x08
	TRANSACTION_CONFIRMATION_DISPLAY_REMOTE              = 0x10

	//Tags used for crypto algorithms and types
	UAF_ALG_SIGN_SECP256R1_ECDSA_SHA256_RAW = 0x01
	UAF_ALG_SIGN_SECP256R1_ECDSA_SHA256_DER = 0x02
	UAF_ALG_SIGN_RSASSA_PSS_SHA256_RAW      = 0x03
	UAF_ALG_SIGN_RSASSA_PSS_SHA256_DER      = 0x04
	UAF_ALG_SIGN_SECP256K1_ECDSA_SHA256_RAW = 0x05
	UAF_ALG_SIGN_SECP256K1_ECDSA_SHA256_DER = 0x06

	//Public Key Representation Formats
	UAF_ALG_KEY_ECC_X962_RAW     = 0x100
	UAF_ALG_KEY_ECC_X962_DER     = 0x101
	UAF_ALG_KEY_RSA_2048_PSS_RAW = 0x102
	UAF_ALG_KEY_RSA_2048_PSS_DER = 0x103

	//Predefined Tags
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
	TAG_EXTENSION_2                 = 0x3E12
	TAG_EXTENSION_ID                = 0x2E13
	TAG_EXTENSION_DATA              = 0x2E14
)

type DOMString string

type Operation DOMString

type AAID DOMString

type KeyID DOMString

type ServerChallenge DOMString

type Version struct {
	Major uint16 `json:"major"`
	Minor uint16 `json:"minor"`
}

type Extension struct {
	ID            string `json:"id"`
	Data          string `json:"data"`
	FailIfUnknown bool   `json:"fail_if_unknown"`
}

type OperationHeader struct {
	UPV        Version     `json:"upv"`
	Op         Operation   `json:"op"`
	AppID      string      `json:"appID"`
	ServerData string      `json:"serverData"`
	Exts       []Extension `json:"exts"`
}

type ChannelBinding struct {
	ServerEndPoint       DOMString `json:"serverEndPoint"`
	TlsServerCertificate DOMString `json:"tlsServerCertificate"`
	TlsUnique            DOMString `json:"tlsUnique"`
	CID_Pubkey           DOMString `json:"cid_pubkey"`
}

type FinalChallengeParams struct {
	AppID          DOMString       `json:"appID"`
	Challenge      ServerChallenge `json:"challenge"`
	FacetID        DOMString       `json:"facetID"`
	ChannelBinding ChannelBinding  `json:"channelBinding"`
}

type JwkKey struct {
	KTY DOMString `json:"kty"` //EC
	CRV DOMString `json:"crv"` //P-256
	X   DOMString `json:"x"`
	Y   DOMString `json:"y"`
}

type MatchCriteria struct {
	AAID                     []AAID      `json:"aaid"`
	VendorID                 []DOMString `json:"vendorID"`
	KeyIDs                   []KeyID     `json:"keyIDs"`
	UserVerification         int32       `json:"userVerification"`
	KeyProtection            int16       `json:"keyProtection"`
	MatcherProtection        int16       `json:"matcherProtection"`
	AttachmentHint           int32       `json:"attachmentHint"`
	TcDisplay                int16       `json:"tcDisplay"`
	AuthenticationAlgorithms []int16     `json:"authenticationAlgorithms"`
	AssertionSchemes         []DOMString `json:"assertionSchemes"`
	AssertionTypes           []int16     `json:"attestationTypes"`
	AuthenticatorVersion     int16       `json:"authenticatorVersion"`
	Exts                     []Extension `json:"exts"`
}

type Policy struct {
	Accepted   []MatchCriteria `json:"accepted"`
	Disallowed []MatchCriteria `json:"disallowed"`
}

type RegistrationRequest struct {
	Header    OperationHeader `json:"header"`
	Challenge ServerChallenge `json:"challenge"`
	Username  DOMString       `json:"username"`
	Policy    Policy          `json:"policy"`
}

type AuthenticatorRegistrationAssertion struct {
	AssertionScheme             DOMString                             `json:"assertionScheme"`
	Assertion                   DOMString                             `json:"assertion"`
	TcDisplayPNGCharacteristics []DisplayPNGCharacteristicsDescriptor `json:"tcDisplayPNGCharacteristics"`
	Exts                        []Extension                           `json:"exts"`
}

type RegistrationResponse struct {
	Header     OperationHeader                      `json:"header"`
	FcParams   DOMString                            `json:"fcParams"`
	Assertions []AuthenticatorRegistrationAssertion `json:"assertions"`
}

//Acceptory Descriptors

type CodeAccuracyDescriptor struct {
	Base          uint16 `json:"base"`
	MinLength     uint16 `json:"minLength"`
	MaxRetries    uint16 `json:"maxRetries"`
	BlockSlowdown uint16 `json:"blockSlowdown"`
}

type BiometricAccuracyDescriptor struct {
	FAR                  float32 `json:"FAR"`
	FRR                  float32 `json:"FRR"`
	EER                  float32 `json:"EER"`
	FAAR                 float32 `json:"FAAR"`
	MaxReferenceDataSets uint16  `json:"maxReferenceDataSets"`
	MaxRetries           uint16  `json:"maxRetries"`
	BlockSlowdown        uint16  `json:"blockSlowdown"`
}

type PatternAccuracyDescriptor struct {
	MinComplexity uint32 `json:"minComplexity"`
	MaxRetries    uint16 `json:"maxRetries"`
	BlockSlowdown uint16 `json:"blockSlowdown"`
}

type VerificationMethodDescriptor struct {
	UserVerification uint32                      `json:"userVerification"`
	CaDesc           CodeAccuracyDescriptor      `json:"caDesc"`
	BaDesc           BiometricAccuracyDescriptor `json:"baDesc"`
	PaDesc           PatternAccuracyDescriptor   `json:"paDesc"`
}

type VerificationMethodANDCombinations []VerificationMethodDescriptor

type RgbPalletteEntry struct {
	R uint16 `json:"r"`
	B uint16 `json:"g"`
	G uint16 `json:"b"`
}

type DisplayPNGCharacteristicsDescriptor struct {
	Width       uint32             `json:"width"`
	Height      uint32             `json:"height"`
	BitDepth    []byte             `json:"bitDepth"`
	ColorType   []byte             `json:"colorType"`
	Compression []byte             `json:"compression"`
	Filter      []byte             `json:"filter"`
	Interlace   []byte             `json:"interlace"`
	Plte        []RgbPalletteEntry `json:"plte"`
}

type MetadataStatement struct {
	Aaid                        AAID                                  `json:"aaid"`
	Description                 DOMString                             `json:"description"`
	AuthenticatorVersion        uint16                                `json:"authenticatorVersion"`
	Upv                         Version                               `json:"upv"`
	AssertionScheme             DOMString                             `json:"assertionScheme"`
	AuthenticationAlgorithm     uint16                                `json:"authenticationAlgorithm"`
	PublicKeyAlgAndEncoding     uint16                                `json:"publicKeyAlgAndEncoding"`
	AttestationTypes            uint16                                `json:"attestationTypes"`
	UserVerificationDetails     []VerificationMethodANDCombinations   `json:"userVerificationDetails"`
	KeyProtection               uint16                                `json:"keyProtection"`
	MatcherProtection           uint16                                `json:"matcherProtection"`
	AttachmentHint              uint32                                `json:"attachmentHint"`
	IsSecondFactorOnly          bool                                  `json:"isSecondFactorOnly"`
	TcDisplay                   uint16                                `json:"tcDisplay"`
	TcDisplayContentType        DOMString                             `json:"tcDisplayContentType"`
	TcDisplayPNGCharacteristics []DisplayPNGCharacteristicsDescriptor `json:"tcDisplayPNGCharacteristics"`
	AttestationRootCertificates []DOMString                           `json:"attestationRootCertificates"`
	Icon                        DOMString                             `json:"icon"`
}

type Transaction struct {
	ContentType                 DOMString                           `json:"contentType"`
	Content                     DOMString                           `json:"content"`
	TcDisplayPNGCharacteristics DisplayPNGCharacteristicsDescriptor `json:"tcDisplayPNGCharacteristics"`
}

type AuthenticationRequest struct {
	Header      OperationHeader `json:"header"`
	Challenge   ServerChallenge `json:"challenge"`
	Transaction []Transaction   `json:"transaction"`
	Policy      Policy          `json:"policy"`
}

type AuthenticatorSignAssertion struct {
	AssertionScheme DOMString   `json:"assertionScheme"`
	Assertion       DOMString   `json:"assertion"`
	Exts            []Extension `json:"exts"`
}

type AuthenticationResponse struct {
	Header     OperationHeader              `json:"header"`
	FcParams   DOMString                    `json:"fcParams"`
	Assertions []AuthenticatorSignAssertion `json:"assertions"`
}

type DeregisterAuthenticator struct {
	Aaid  AAID  `json:"aaid"`
	KeyID KeyID `json:"keyID"`
}

type DeregistrationRequest struct {
	Header         OperationHeader           `json:"header"`
	Authenticators []DeregisterAuthenticator `json:"authenticators"`
}
