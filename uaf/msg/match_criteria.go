package msg

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
