package msg

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
