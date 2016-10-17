package msg

type AuthenticatorRegistrationAssertion struct {
	AssertionScheme             DOMString                             `json:"assertionScheme"`
	Assertion                   DOMString                             `json:"assertion"`
	TcDisplayPNGCharacteristics []DisplayPNGCharacteristicsDescriptor `json:"tcDisplayPNGCharacteristics"`
	Exts                        []Extension                           `json:"exts"`
}
