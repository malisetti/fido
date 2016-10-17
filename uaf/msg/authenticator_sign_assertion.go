package msg

type AuthenticatorSignAssertion struct {
	AssertionScheme DOMString   `json:"assertionScheme"`
	Assertion       DOMString   `json:"assertion"`
	Exts            []Extension `json:"exts"`
}
