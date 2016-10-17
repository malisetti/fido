package msg

type JwkKey struct {
	KTY DOMString `json:"kty"` //EC
	CRV DOMString `json:"crv"` //P-256
	X   DOMString `json:"x"`
	Y   DOMString `json:"y"`
}
